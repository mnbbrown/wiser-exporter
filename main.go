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

type config struct {
	HubIP    string        `envconfig:"wiser_hub_ip" required:"true"`
	Secret   string        `envconfig:"wiser_secret" required:"true"`
	Port     int           `envconfig:"PORT" default:"2112"`
	Interval time.Duration `envconfig:"INTERVAL" default:"30s"`
}

func (c *config) url(path string) string {
	return fmt.Sprintf("http://%s/data/domain/%s", c.HubIP, path)
}

func makeRequest(url string, secret string) (*http.Response, error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Tye", "application/json")
	r.Header.Set("SECRET", secret)

	resp, err := http.DefaultClient.Do(r)
	return resp, err
}

func read(c config) {
	for {
		resp, err := makeRequest(c.url(""), c.Secret)
		if err != nil {
			slog.Warn("failed_to_request", "err", err)
			time.Sleep(5 * time.Second)
			continue
		}

		d := &DomainData{}
		reader := json.NewDecoder(resp.Body)
		if err := reader.Decode(&d); err != nil {
			slog.Warn("failed_to_decode", "err", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, room := range d.Room {
			labels := []string{fmt.Sprintf("%d", room.ID), room.Name}
			roomTemperature.WithLabelValues(labels...).Set(parseTemp(room.CalculatedTemperature))
			roomSetPoint.WithLabelValues(labels...).Set(parseTemp(room.CurrentSetPoint))
			roomHeatingRate.WithLabelValues(labels...).Set(float64(room.HeatingRate))
			roomDemand.WithLabelValues(labels...).Set(float64(room.PercentageDemand))

			humidity := getHumidity(d, room.RoomStatID)
			if humidity != -1 {
				roomHumidity.WithLabelValues(labels...).Set(float64(humidity))
			}
		}

		for _, channel := range d.HeatingChannel {
			labels := []string{fmt.Sprintf("%d", channel.ID), channel.Name}
			channelPercentDemand.WithLabelValues(labels...).Set(float64(channel.PercentageDemand))
		}

		for _, device := range d.Device {
			if device.ProductType == "Controller" {
				continue
			}

			labels := []string{fmt.Sprintf("%d", device.ID), device.ProductType, device.SerialNumber}
			deviceBatteryVoltage.WithLabelValues(labels...).Set(float64(device.BatteryVoltage))
		}

		slog.Debug("read_successful")
		time.Sleep(c.Interval)
	}
}

func main() {
	c := config{}

	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// start reading
	go read(c)

	http.Handle("/metrics", promhttp.Handler())
	addr := fmt.Sprintf(":%d", c.Port)
	slog.Info("listening", "address", addr)
	http.ListenAndServe(addr, nil)
}
