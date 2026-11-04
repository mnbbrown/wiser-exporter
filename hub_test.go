package main

import (
	"encoding/json"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

// Shapes taken from a live hub (firmware 3.18.3), with identifiers replaced.
const devicesJSON = `{
  "Room": [
    {"id": 1, "Name": "Living Room", "SmartValveIds": [2], "WindowState": "Closed", "WindowDetectionActive": true, "OverrideType": "None", "ScheduledSetPoint": 190},
    {"id": 2, "Name": "Office", "SmartValveIds": [3], "WindowDetectionActive": true, "OverrideType": "Manual", "ScheduledSetPoint": 150}
  ],
  "Device": [
    {"id": 0, "ProductType": "Controller", "ReceptionOfController": {"Rssi": -85}},
    {"id": 2, "ProductType": "iTRV", "SerialNumber": "A", "BatteryLevel": "TwoThirds",
     "ReceptionOfController": {"Rssi": -82, "Lqi": 72}, "ReceptionOfDevice": {"Rssi": -76, "Lqi": 96}},
    {"id": 3, "ProductType": "iTRV", "SerialNumber": "B"}
  ],
  "SmartValve": [
    {"id": 2, "MeasuredTemperature": 194, "PercentageDemand": 0, "WindowState": "Open"},
    {"id": 3, "MeasuredTemperature": -32768}
  ]
}`

func TestRecordDevices(t *testing.T) {
	d := &DomainData{}
	if err := json.Unmarshal([]byte(devicesJSON), d); err != nil {
		t.Fatal(err)
	}
	recordDomain(d)

	// Only the reporting valve has link data; the controller is skipped.
	assert.Equal(t, 2, testutil.CollectAndCount(deviceRSSI))
	assert.Equal(t, -82.0, testutil.ToFloat64(deviceRSSI.WithLabelValues("2", "iTRV", "A", "hub_to_device")))
	assert.Equal(t, 96.0, testutil.ToFloat64(deviceLQI.WithLabelValues("2", "iTRV", "A", "device_to_hub")))
	assert.Equal(t, 1.0, testutil.ToFloat64(deviceBatteryLevel.WithLabelValues("2", "iTRV", "A", "twothirds")))

	// The unpaired valve reports no temperature, demand or window state.
	assert.Equal(t, 1, testutil.CollectAndCount(valveTemperature))
	assert.Equal(t, 19.4, testutil.ToFloat64(valveTemperature.WithLabelValues("2", "Living Room")))
	assert.Equal(t, 1, testutil.CollectAndCount(valveDemand))
	assert.Equal(t, 1.0, testutil.ToFloat64(valveWindowOpen.WithLabelValues("2", "Living Room")))

	assert.Equal(t, 0.0, testutil.ToFloat64(roomOverrideActive.WithLabelValues("1", "Living Room")))
	assert.Equal(t, 1.0, testutil.ToFloat64(roomOverrideActive.WithLabelValues("2", "Office")))
	assert.Equal(t, 19.0, testutil.ToFloat64(roomScheduledSetPoint.WithLabelValues("1", "Living Room")))

	// A device that drops off must not leave its old series behind.
	d.Device = d.Device[:1]
	recordDomain(d)
	assert.Equal(t, 0, testutil.CollectAndCount(deviceRSSI))
	assert.Equal(t, 0, testutil.CollectAndCount(deviceBatteryLevel))
}

func TestRecordHub(t *testing.T) {
	n := &NetworkData{}
	if err := json.Unmarshal([]byte(`{"Station": {"SSID": "home", "Channel": 6, "ConnectionStatus": "Connected", "RSSI": {"Current": -82, "Min": -86, "Max": -81}}}`), n); err != nil {
		t.Fatal(err)
	}
	recordNetwork(n)
	assert.Equal(t, -82.0, testutil.ToFloat64(hubWifiRSSI))
	assert.Equal(t, 1.0, testutil.ToFloat64(hubWifiConnected))

	s := &StatusData{}
	if err := json.Unmarshal([]byte(`{"uptime": 3634, "freeHeap": 67064, "lowestFreeHeap": 43176, "lastResetReason": "GENERAL_RST"}`), s); err != nil {
		t.Fatal(err)
	}
	recordStatus(s)
	assert.Equal(t, 3634.0, testutil.ToFloat64(hubUptime))
	assert.Equal(t, 1.0, testutil.ToFloat64(hubLastReset.WithLabelValues("GENERAL_RST")))
}
