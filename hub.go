package main

import (
	"fmt"
	"strings"
)

// NetworkData is /data/v2/network/: the hub's own WiFi connection.
type NetworkData struct {
	Station struct {
		SSID             string `json:"SSID"`
		Channel          int    `json:"Channel"`
		ConnectionStatus string `json:"ConnectionStatus"`
		RSSI             struct {
			Current int `json:"Current"`
			Min     int `json:"Min"`
			Max     int `json:"Max"`
		} `json:"RSSI"`
		ConnectionFailures int `json:"ConnectionFailures"`
	} `json:"Station"`
}

// StatusData is /data/v2/status/: hub uptime and memory.
type StatusData struct {
	Uptime          int    `json:"uptime"`
	FreeHeap        int    `json:"freeHeap"`
	LowestFreeHeap  int    `json:"lowestFreeHeap"`
	LastResetReason string `json:"lastResetReason"`
}

func recordNetwork(n *NetworkData) {
	st := n.Station
	hubWifiConnected.Set(boolToFloat(st.ConnectionStatus == "Connected"))
	hubWifiRSSI.Set(float64(st.RSSI.Current))
	hubWifiConnectionFailures.Set(float64(st.ConnectionFailures))
	hubWifiInfo.Reset()
	hubWifiInfo.WithLabelValues(st.SSID, fmt.Sprintf("%d", st.Channel)).Set(1)
}

func recordStatus(s *StatusData) {
	hubUptime.Set(float64(s.Uptime))
	hubFreeHeap.Set(float64(s.FreeHeap))
	hubLowestFreeHeap.Set(float64(s.LowestFreeHeap))
	hubLastReset.Reset()
	hubLastReset.WithLabelValues(s.LastResetReason).Set(1)
}

// windowOpen maps a WindowState to 0/1. ok is false when the device has no
// window state to report (e.g. a valve that isn't reporting yet).
func windowOpen(state string) (v float64, ok bool) {
	switch state {
	case "Open":
		return 1, true
	case "Closed":
		return 0, true
	}
	return 0, false
}

// recordDevices exports per-device link quality and battery level, and
// per-valve readings. Called from recordDomain.
func recordDevices(d *DomainData) {
	// Rebuilt every scrape so a device that drops off the network, or whose
	// battery level changes, doesn't leave a stale series behind.
	deviceRSSI.Reset()
	deviceLQI.Reset()
	deviceBatteryLevel.Reset()
	valveTemperature.Reset()
	valveDemand.Reset()
	valveWindowOpen.Reset()

	for _, device := range d.Device {
		if device.ProductType == "Controller" {
			continue
		}
		labels := []string{fmt.Sprintf("%d", device.ID), device.ProductType, device.SerialNumber}

		// ReceptionOfController: how well the device hears the hub.
		// ReceptionOfDevice: how well the hub hears the device.
		for _, link := range []struct {
			name string
			r    *Reception
		}{
			{"hub_to_device", device.ReceptionOfController},
			{"device_to_hub", device.ReceptionOfDevice},
		} {
			if link.r == nil {
				continue
			}
			deviceRSSI.WithLabelValues(append(labels, link.name)...).Set(float64(link.r.Rssi))
			if link.r.Lqi != 0 {
				deviceLQI.WithLabelValues(append(labels, link.name)...).Set(float64(link.r.Lqi))
			}
		}

		if device.BatteryLevel != "" {
			deviceBatteryLevel.WithLabelValues(append(labels, strings.ToLower(device.BatteryLevel))...).Set(1)
		}
	}

	roomForValve := map[int]string{}
	for _, room := range d.Room {
		for _, id := range room.SmartValveIDs {
			roomForValve[id] = room.Name
		}
	}

	for _, v := range d.SmartValve {
		labels := []string{fmt.Sprintf("%d", v.ID), roomForValve[v.ID]}
		if v.MeasuredTemperature != noReading {
			valveTemperature.WithLabelValues(labels...).Set(parseTemp(v.MeasuredTemperature))
		}
		if v.PercentageDemand != nil {
			valveDemand.WithLabelValues(labels...).Set(float64(*v.PercentageDemand))
		}
		if open, ok := windowOpen(v.WindowState); ok {
			valveWindowOpen.WithLabelValues(labels...).Set(open)
		}
	}
}
