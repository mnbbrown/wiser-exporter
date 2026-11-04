package main

// OpenThermData is the hub's view of the OpenTherm link to the boiler,
// served from /data/v2/opentherm/. Temperatures, pressure and modulation are
// reported as integers scaled by 10 (e.g. 528 = 52.8°C, 10 = 1.0 bar).
type OpenThermData struct {
	Enabled                   bool   `json:"Enabled"`
	OperatingMode             string `json:"operatingMode"`
	ChFlowActiveUpperSetpoint int    `json:"chFlowActiveUpperSetpoint"`
	ChFlowActiveLowerSetpoint int    `json:"chFlowActiveLowerSetpoint"`
	Ch1FlowSetpoint           int    `json:"ch1FlowSetpoint"`
	Ch1FlowEnable             bool   `json:"ch1FlowEnable"`
	DhwFlowSetpoint           int    `json:"dhwFlowSetpoint"`
	DhwEnable                 bool   `json:"dhwEnable"`
	RoomTemperature           int    `json:"roomTemperature"`
	RoomSetpoint              int    `json:"roomSetpoint"`
	TrackedRoomID             int    `json:"TrackedRoomId"`
	Transmitter               struct {
		MessagesSentCount int `json:"messagesSentCount"`
	} `json:"transmitter"`
	Receiver struct {
		DecodingErrorsCount int `json:"decodingErrorsCount"`
	} `json:"receiver"`
	OperationalData struct {
		RelativeModulationLevel int `json:"RelativeModulationLevel"`
		MinimumModulationLevel  int `json:"MinimumModulationLevel"`
		ChPressureBar           int `json:"ChPressureBar"`
		DhwFlowRate             int `json:"DhwFlowRate"`
		Ch1FlowTemperature      int `json:"Ch1FlowTemperature"`
		ChReturnTemperature     int `json:"ChReturnTemperature"`
		MaximumCapacityKw       int `json:"MaximumCapacityKw"`
		SlaveStatus             int `json:"SlaveStatus"`
	} `json:"operationalData"`
}

// slaveStatusFlags names the bits of the OpenTherm slave status byte
// (data-id 0, low byte) as defined by the OpenTherm protocol spec.
var slaveStatusFlags = []struct {
	bit  uint
	name string
}{
	{0, "fault"},
	{1, "ch_active"},
	{2, "dhw_active"},
	{3, "flame"},
	{4, "cooling_active"},
	{5, "ch2_active"},
	{6, "diagnostic"},
}

func statusFlag(status int, bit uint) float64 {
	if status&(1<<bit) != 0 {
		return 1
	}
	return 0
}

func boolToFloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func recordOpenTherm(o *OpenThermData) {
	op := o.OperationalData
	otEnabled.Set(boolToFloat(o.Enabled))
	otFlowTemperature.Set(parseTemp(op.Ch1FlowTemperature))
	otReturnTemperature.Set(parseTemp(op.ChReturnTemperature))
	otPressure.Set(parseTemp(op.ChPressureBar))
	otModulation.Set(parseTemp(op.RelativeModulationLevel))
	otMinModulation.Set(float64(op.MinimumModulationLevel))
	otMaxCapacity.Set(float64(op.MaximumCapacityKw))
	otDhwFlowRate.Set(parseTemp(op.DhwFlowRate))
	otChFlowSetpoint.Set(parseTemp(o.Ch1FlowSetpoint))
	otChFlowEnabled.Set(boolToFloat(o.Ch1FlowEnable))
	otDhwSetpoint.Set(parseTemp(o.DhwFlowSetpoint))
	otDhwEnabled.Set(boolToFloat(o.DhwEnable))
	otMessagesSent.Set(float64(o.Transmitter.MessagesSentCount))
	otDecodingErrors.Set(float64(o.Receiver.DecodingErrorsCount))
	otSlaveStatus.Set(float64(op.SlaveStatus))
	for _, f := range slaveStatusFlags {
		otStatusFlag.WithLabelValues(f.name).Set(statusFlag(op.SlaveStatus, f.bit))
	}
}
