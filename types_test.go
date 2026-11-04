package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDomainDataLoad(t *testing.T) {
	raw, err := os.ReadFile("fixtures/domain_data.json")
	if err != nil {
		t.Fatal(err)
	}

	d := &DomainData{}
	if err := json.Unmarshal(raw, &d); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, d.Room[0].CalculatedTemperature, 200)
}

func TestNoReadingIsNotExported(t *testing.T) {
	raw, err := os.ReadFile("fixtures/domain_data.json")
	if err != nil {
		t.Fatal(err)
	}
	d := &DomainData{}
	if err := json.Unmarshal(raw, d); err != nil {
		t.Fatal(err)
	}

	recordDomain(d)
	before := testutil.CollectAndCount(roomTemperature)

	// The hub reports -3276.8°C for a room with no working sensor; that
	// room must drop out rather than be exported as a temperature.
	d.Room[0].CalculatedTemperature = noReading
	recordDomain(d)
	assert.Equal(t, before-1, testutil.CollectAndCount(roomTemperature))
}
