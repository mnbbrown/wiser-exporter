package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// room
	roomTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_temperature",
		Help: "The temperature of a room",
	}, []string{"id", "name"})

	roomHumidity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "wiser_room_humidity",
		Help: "The humidity of a room",
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
)
