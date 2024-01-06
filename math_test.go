package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundUp05(t *testing.T) {
	r := round(0.5)
	assert.Equal(t, r, 1)
}

func TestRoundUp06(t *testing.T) {
	r := round(0.6)
	assert.Equal(t, r, 1)
}

func TestRoundDown04(t *testing.T) {
	r := round(0.49999)
	assert.Equal(t, r, 0)
}

func TestFixedRounding(t *testing.T) {
	r := toFixed(0.49999, 2)
	assert.Equal(t, r, 0.50)
}

func TestFixedNoRounding(t *testing.T) {
	r := toFixed(0.414, 2)
	assert.Equal(t, r, 0.41)
}

func TestFixedNoPrecision(t *testing.T) {
	r := toFixed(41, 2)
	assert.Equal(t, r, 41.0)
}

func TestFixedNoPrecisionLarge(t *testing.T) {
	r := toFixed(413, 2)
	assert.Equal(t, r, 413.0)
}

func TestParseTemp(t *testing.T) {
	r := parseTemp(200)
	assert.Equal(t, r, 20.0)
}

func TestParseTempPrecision(t *testing.T) {
	r := parseTemp(194)
	assert.Equal(t, r, 19.4)
}

func TestParseTempPrecisionLarge(t *testing.T) {
	r := parseTemp(1944)
	assert.Equal(t, r, 194.4)
}

func TestParseTempNegative(t *testing.T) {
	r := parseTemp(-123)
	assert.Equal(t, r, -12.3)
}
