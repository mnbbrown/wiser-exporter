package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	buildInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_exporter_build_info",
		Help: "Always 1; labels carry the exporter's version, commit and build date",
	}, []string{"version", "commit", "date"})

	// room
	roomTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_temperature",
		Help: "The temperature of a room",
	}, []string{"id", "name"})

	roomHumidity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_humidity",
		Help: "The humidity of a room",
	}, []string{"id", "name"})

	roomDemand = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_demand",
		Help: "The percent (out of 100) the room is demanding",
	}, []string{"id", "name"})

	roomSetPoint = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_set_point",
		Help: "The set point of a room",
	}, []string{"id", "name"})

	roomHeatingRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_heating_rate",
		Help: "The heating rate of the room",
	}, []string{"id", "name"})

	// channel
	channelPercentDemand = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_channel_percent_demand",
		Help: "The percent wiser is demanding from a channel",
	}, []string{"id", "name"})

	channelRelayOn = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_channel_relay_on",
		Help: "1 if the heating relay for a channel is on",
	}, []string{"id", "name"})

	// devices
	deviceBatteryVoltage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_device_battery_voltage",
		Help: "The voltage of the battery reported by a device",
	}, []string{"id", "product_type", "serial_number"})

	// opentherm
	otConnected = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_connected",
		Help: "1 if the hub reports its OpenTherm link to the boiler as connected",
	})
	otEnabled = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_enabled",
		Help: "1 if OpenTherm control is enabled on the hub",
	})
	otFlowTemperature = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_flow_temperature_celsius",
		Help: "Boiler CH flow temperature",
	})
	otReturnTemperature = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_return_temperature_celsius",
		Help: "Boiler CH return temperature",
	})
	otPressure = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_ch_pressure_bar",
		Help: "Central heating water pressure",
	})
	otModulation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_relative_modulation_percent",
		Help: "Boiler relative modulation level. Holds its last value while the burner is off; use with the flame flag",
	})
	otMinModulation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_minimum_modulation_percent",
		Help: "Boiler minimum modulation level",
	})
	otMaxCapacity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_maximum_capacity_kw",
		Help: "Boiler maximum capacity",
	})
	otDhwFlowRate = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_dhw_flow_rate_litres_per_minute",
		Help: "Domestic hot water flow rate",
	})
	otChFlowSetpoint = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_ch_flow_setpoint_celsius",
		Help: "CH flow temperature the hub is requesting from the boiler",
	})
	otChFlowEnabled = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_ch_enabled",
		Help: "1 if the hub is requesting central heating from the boiler",
	})
	otDhwSetpoint = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_dhw_setpoint_celsius",
		Help: "DHW flow setpoint",
	})
	otDhwEnabled = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_dhw_enabled",
		Help: "1 if DHW is enabled",
	})
	otMessagesSent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_messages_sent",
		Help: "OpenTherm messages sent by the hub since it booted (monotonic until hub restart)",
	})
	otDecodingErrors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_decoding_errors",
		Help: "OpenTherm replies the hub failed to decode since it booted (monotonic until hub restart)",
	})
	otSlaveStatus = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_opentherm_slave_status",
		Help: "Raw OpenTherm slave status byte",
	})
	otStatusFlag = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_opentherm_status_flag",
		Help: "Decoded OpenTherm slave status flags (fault, ch_active, dhw_active, flame, ...)",
	}, []string{"flag"})

	// scrape health
	scrapeErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "wiser_scrape_errors_total",
		Help: "Failed requests to the hub, by endpoint",
	}, []string{"endpoint"})
	lastScrapeSuccess = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_last_scrape_success_timestamp_seconds",
		Help: "Unix time of the last successful read of each hub endpoint",
	}, []string{"endpoint"})
)

var (
	// room extras
	roomScheduledSetPoint = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_scheduled_set_point",
		Help: "The set point the room's schedule asks for right now (°C)",
	}, []string{"id", "name"})
	roomOverrideActive = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_override_active",
		Help: "1 while a boost or manual override is in effect",
	}, []string{"id", "name"})
	roomWindowDetectionEnabled = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_window_detection_enabled",
		Help: "1 if open-window detection is enabled for the room",
	}, []string{"id", "name"})
	roomWindowOpen = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_window_open",
		Help: "1 if the room is in open-window state (inferred from a rapid temperature drop)",
	}, []string{"id", "name"})

	// devices
	deviceRSSI = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_device_rssi_dbm",
		Help: "Zigbee signal strength; link is hub_to_device (as heard by the device) or device_to_hub",
	}, []string{"id", "product_type", "serial_number", "link"})
	deviceLQI = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_device_lqi",
		Help: "Zigbee link quality indicator (0-255, higher is better)",
	}, []string{"id", "product_type", "serial_number", "link"})
	deviceBatteryLevel = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_device_battery_level",
		Help: "Always 1; the level label is the hub's battery bucket (normal, twothirds, onethird, low, ...)",
	}, []string{"id", "product_type", "serial_number", "level"})

	// smart valves (iTRVs)
	valveTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_valve_temperature",
		Help: "Temperature measured at the radiator valve itself (°C)",
	}, []string{"id", "room"})
	valveDemand = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_valve_demand",
		Help: "How far the valve is asking to open (%)",
	}, []string{"id", "room"})
	valveWindowOpen = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_valve_window_open",
		Help: "1 if the valve has detected an open window",
	}, []string{"id", "room"})

	// hub
	hubWifiConnected = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_wifi_connected",
		Help: "1 if the hub reports its WiFi station as connected",
	})
	hubWifiRSSI = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_wifi_rssi_dbm",
		Help: "Signal strength of the hub's WiFi connection",
	})
	hubWifiConnectionFailures = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_wifi_connection_failures",
		Help: "WiFi connection failures since the hub booted",
	})
	hubWifiInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_hub_wifi_info",
		Help: "Always 1; labels carry the SSID and channel the hub is on",
	}, []string{"ssid", "channel"})
	hubUptime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_uptime_seconds",
		Help: "Seconds since the hub last booted",
	})
	hubFreeHeap = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_free_heap_bytes",
		Help: "Free heap memory on the hub",
	})
	hubLowestFreeHeap = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_lowest_free_heap_bytes",
		Help: "Lowest free heap seen since the hub booted",
	})
	hubLastReset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_hub_last_reset_info",
		Help: "Always 1; the reason label says why the hub last restarted",
	}, []string{"reason"})
	hubZigbeeNoSignalResets = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_zigbee_no_signal_resets",
		Help: "Hub Zigbee NoSignalReset counter (meaning undocumented; watch the trend)",
	})
	hubZigbeeJPANCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "wiser_hub_zigbee_jpan_count",
		Help: "Hub Zigbee JPANCount counter (meaning undocumented; watch the trend)",
	})
)
