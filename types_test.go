package main

import (
	"encoding/json"
	"os"
	"testing"

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
