package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Set by goreleaser at build time.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// noReading is the value the hub reports for a temperature it doesn't have,
// e.g. a room whose sensor is offline or not yet paired (shows as -3276.8°C).
const noReading = -32768

type config struct {
	HubIP    string        `envconfig:"wiser_hub_ip" required:"true"`
	Secret   string        `envconfig:"wiser_secret" required:"true"`
	Port     int           `envconfig:"PORT" default:"2112"`
	Interval time.Duration `envconfig:"INTERVAL" default:"30s"`
}

func (c *config) url(path string) string {
	return fmt.Sprintf("http://%s/data/v2/%s", c.HubIP, path)
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

// getJSON fetches a hub endpoint and decodes the response into out.
func getJSON(c config, path string, out any) error {
	r, err := http.NewRequest("GET", c.url(path), nil)
	if err != nil {
		return err
	}
	r.Header.Set("SECRET", c.Secret)

	resp, err := httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from %s", resp.StatusCode, path)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func recordDomain(d *DomainData) {
	for _, room := range d.Room {
		labels := []string{fmt.Sprintf("%d", room.ID), room.Name}
		if room.CalculatedTemperature == noReading {
			roomTemperature.DeleteLabelValues(labels...)
		} else {
			roomTemperature.WithLabelValues(labels...).Set(parseTemp(room.CalculatedTemperature))
		}
		roomSetPoint.WithLabelValues(labels...).Set(parseTemp(room.CurrentSetPoint))
		roomHeatingRate.WithLabelValues(labels...).Set(float64(room.HeatingRate))
		roomDemand.WithLabelValues(labels...).Set(float64(room.PercentageDemand))

		roomScheduledSetPoint.WithLabelValues(labels...).Set(parseTemp(room.ScheduledSetPoint))
		roomOverrideActive.WithLabelValues(labels...).Set(boolToFloat(room.OverrideType != "" && room.OverrideType != "None"))
		roomWindowDetectionEnabled.WithLabelValues(labels...).Set(boolToFloat(room.WindowDetectionActive))
		if open, ok := windowOpen(room.WindowState); ok {
			roomWindowOpen.WithLabelValues(labels...).Set(open)
		} else {
			roomWindowOpen.DeleteLabelValues(labels...)
		}

		humidity := getHumidity(d, room.RoomStatID)
		if humidity != -1 {
			roomHumidity.WithLabelValues(labels...).Set(float64(humidity))
		}
	}

	for _, channel := range d.HeatingChannel {
		labels := []string{fmt.Sprintf("%d", channel.ID), channel.Name}
		channelPercentDemand.WithLabelValues(labels...).Set(float64(channel.PercentageDemand))
		channelRelayOn.WithLabelValues(labels...).Set(boolToFloat(channel.HeatingRelayState == "On"))
	}

	for _, device := range d.Device {
		if device.ProductType == "Controller" {
			continue
		}

		labels := []string{fmt.Sprintf("%d", device.ID), device.ProductType, device.SerialNumber}
		// Mains-powered and offline devices report no voltage at all.
		if device.BatteryVoltage == 0 {
			deviceBatteryVoltage.DeleteLabelValues(labels...)
		} else {
			deviceBatteryVoltage.WithLabelValues(labels...).Set(float64(device.BatteryVoltage))
		}
	}

	recordDevices(d)

	hubZigbeeNoSignalResets.Set(float64(d.Zigbee.NoSignalReset))
	hubZigbeeJPANCount.Set(float64(d.Zigbee.JPANCount))
	otConnected.Set(boolToFloat(d.System.OpenThermConnectionStatus == "Connected"))
}

// scrape reads one endpoint and records it, tracking failures per endpoint so
// a broken OpenTherm read doesn't stop room metrics (or the reverse).
func scrape[T any](c config, endpoint string, record func(*T)) {
	v := new(T)
	if err := getJSON(c, endpoint+"/", v); err != nil {
		slog.Warn("scrape_failed", "endpoint", endpoint, "err", err)
		scrapeErrors.WithLabelValues(endpoint).Inc()
		return
	}
	record(v)
	lastScrapeSuccess.WithLabelValues(endpoint).SetToCurrentTime()
}

func read(c config) {
	for {
		scrape(c, "domain", recordDomain)
		scrape(c, "opentherm", recordOpenTherm)
		scrape(c, "network", recordNetwork)
		scrape(c, "status", recordStatus)
		slog.Debug("read_complete")
		time.Sleep(c.Interval)
	}
}

func main() {
	c := config{}

	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	buildInfo.WithLabelValues(version, commit, date).Set(1)
	slog.Info("starting", "version", version, "commit", commit, "hub", c.HubIP)

	// start reading
	go read(c)

	http.Handle("/metrics", promhttp.Handler())
	addr := fmt.Sprintf(":%d", c.Port)
	slog.Info("listening", "address", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
