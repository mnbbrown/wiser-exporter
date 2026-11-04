package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestOpenThermDataLoad(t *testing.T) {
	raw, err := os.ReadFile("fixtures/opentherm_data.json")
	if err != nil {
		t.Fatal(err)
	}

	o := &OpenThermData{}
	if err := json.Unmarshal(raw, o); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 528, o.OperationalData.Ch1FlowTemperature)
	assert.Equal(t, 10, o.OperationalData.ChPressureBar)
	assert.Equal(t, "OTplus", o.OperatingMode)

	recordOpenTherm(o)
	assert.Equal(t, 52.8, testutil.ToFloat64(otFlowTemperature))
	assert.Equal(t, 1.0, testutil.ToFloat64(otPressure))
	assert.Equal(t, 3200.0, testutil.ToFloat64(otMessagesSent))
}

func TestStatusFlags(t *testing.T) {
	// 10 = 0b1010: CH active + flame, observed live while the boiler was firing.
	o := &OpenThermData{}
	o.OperationalData.SlaveStatus = 10
	recordOpenTherm(o)

	want := map[string]float64{"fault": 0, "ch_active": 1, "dhw_active": 0, "flame": 1}
	for flag, v := range want {
		assert.Equal(t, v, testutil.ToFloat64(otStatusFlag.WithLabelValues(flag)), flag)
	}
}
