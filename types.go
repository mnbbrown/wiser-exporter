package main

type DomainData struct {
	System struct {
		PairingStatus              string `json:"PairingStatus"`
		ProvisionTokenSucceedCount int    `json:"provisionTokenSucceedCount"`
		ProvisionTokenRequestCount int    `json:"provisionTokenRequestCount"`
		TimeZoneOffset             int    `json:"TimeZoneOffset"`
		AutomaticDaylightSaving    bool   `json:"AutomaticDaylightSaving"`
		SystemMode                 string `json:"SystemMode"`
		FotaEnabled                bool   `json:"FotaEnabled"`
		ValveProtectionEnabled     bool   `json:"ValveProtectionEnabled"`
		AwayModeAffectsHotWater    bool   `json:"AwayModeAffectsHotWater"`
		AwayModeSetPointLimit      int    `json:"AwayModeSetPointLimit"`
		BoilerSettings             struct {
			ControlType     string `json:"ControlType"`
			FuelType        string `json:"FuelType"`
			CycleRate       string `json:"CycleRate"`
			OnOffHysteresis int    `json:"OnOffHysteresis"`
		} `json:"BoilerSettings"`
		CoolingModeDefaultSetpoint    int  `json:"CoolingModeDefaultSetpoint"`
		CoolingAwayModeSetpointLimit  int  `json:"CoolingAwayModeSetpointLimit"`
		ComfortModeEnabled            bool `json:"ComfortModeEnabled"`
		PreheatTimeLimit              int  `json:"PreheatTimeLimit"`
		DegradedModeSetpointThreshold int  `json:"DegradedModeSetpointThreshold"`
		GeoPosition                   struct {
			Latitude  float64 `json:"Latitude"`
			Longitude float64 `json:"Longitude"`
		} `json:"GeoPosition"`
		UfhOrphanModeOutput    string `json:"UfhOrphanModeOutput"`
		IsMigrated             bool   `json:"isMigrated"`
		UnixTime               int    `json:"UnixTime"`
		ActiveSystemVersion    string `json:"ActiveSystemVersion"`
		ZigbeePermitJoinActive bool   `json:"ZigbeePermitJoinActive"`
		BrandName              string `json:"BrandName"`
		CloudConnectionStatus  string `json:"CloudConnectionStatus"`
		ChipID                 string `json:"ChipId"`
		LocalDateAndTime       struct {
			Year  int    `json:"Year"`
			Month string `json:"Month"`
			Date  int    `json:"Date"`
			Day   string `json:"Day"`
			Time  int    `json:"Time"`
		} `json:"LocalDateAndTime"`
		HeatingButtonOverrideState  string `json:"HeatingButtonOverrideState"`
		HotWaterButtonOverrideState string `json:"HotWaterButtonOverrideState"`
		OpenThermConnectionStatus   string `json:"OpenThermConnectionStatus"`
		SunriseTimes                []int  `json:"SunriseTimes"`
		SunsetTimes                 []int  `json:"SunsetTimes"`
		IsTrialist                  bool   `json:"isTrialist"`
		IsProvisioned               bool   `json:"isProvisioned"`
		HardwareGeneration          int    `json:"HardwareGeneration"`
	} `json:"System"`
	Cloud struct {
		DetailedPublishing bool   `json:"DetailedPublishing"`
		WiserAPIHost       string `json:"WiserApiHost"`
		BootStrapAPIHost   string `json:"BootStrapApiHost"`
	} `json:"Cloud"`
	HeatingChannel []struct {
		ID                           int    `json:"id"`
		Name                         string `json:"Name"`
		RoomIds                      []int  `json:"RoomIds"`
		PercentageDemand             int    `json:"PercentageDemand"`
		DemandOnOffOutput            string `json:"DemandOnOffOutput"`
		HeatingRelayState            string `json:"HeatingRelayState"`
		IsSmartValvePreventingDemand bool   `json:"IsSmartValvePreventingDemand"`
	} `json:"HeatingChannel"`
	Room []struct {
		ID                      int    `json:"id"`
		ScheduleID              int    `json:"ScheduleId"`
		HeatingRate             int    `json:"HeatingRate"`
		RoomStatID              int    `json:"RoomStatId"`
		Name                    string `json:"Name"`
		Mode                    string `json:"Mode"`
		DemandType              string `json:"DemandType"`
		WindowDetectionActive   bool   `json:"WindowDetectionActive"`
		CalculatedTemperature   int    `json:"CalculatedTemperature"`
		CurrentSetPoint         int    `json:"CurrentSetPoint"`
		PercentageDemand        int    `json:"PercentageDemand"`
		ControlOutputState      string `json:"ControlOutputState"`
		SetpointOrigin          string `json:"SetpointOrigin"`
		DisplayedSetPoint       int    `json:"DisplayedSetPoint"`
		ScheduledSetPoint       int    `json:"ScheduledSetPoint"`
		AwayModeSuppressed      bool   `json:"AwayModeSuppressed"`
		RoundedAlexaTemperature int    `json:"RoundedAlexaTemperature"`
		EffectiveMode           string `json:"EffectiveMode"`
		PercentageDemandForItrv int    `json:"PercentageDemandForItrv"`
		ControlDirection        string `json:"ControlDirection"`
		HeatingType             string `json:"HeatingType"`
	} `json:"Room"`
	Device []struct {
		ID                      int    `json:"id"`
		NodeID                  int    `json:"NodeId"`
		ProductType             string `json:"ProductType"`
		ProductIdentifier       string `json:"ProductIdentifier"`
		ActiveFirmwareVersion   string `json:"ActiveFirmwareVersion"`
		ModelIdentifier         string `json:"ModelIdentifier"`
		DeviceLockEnabled       bool   `json:"DeviceLockEnabled"`
		DisplayedSignalStrength string `json:"DisplayedSignalStrength"`
		ReceptionOfController   struct {
			Rssi int `json:"Rssi"`
		} `json:"ReceptionOfController"`
		SerialNumber           string `json:"SerialNumber,omitempty"`
		ProductModel           string `json:"ProductModel,omitempty"`
		OtaImageQueryCount     int    `json:"OtaImageQueryCount,omitempty"`
		LastOtaImageQueryCount int    `json:"LastOtaImageQueryCount,omitempty"`
		ParentNodeID           int    `json:"ParentNodeId,omitempty"`
		OtaVersion             int    `json:"OtaVersion,omitempty"`
		OtaHardwareVersion     int    `json:"OtaHardwareVersion,omitempty"`
		BatteryVoltage         int    `json:"BatteryVoltage,omitempty"`
		BatteryLevel           string `json:"BatteryLevel,omitempty"`
		ReceptionOfDevice      struct {
			Rssi int `json:"Rssi"`
			Lqi  int `json:"Lqi"`
		} `json:"ReceptionOfDevice,omitempty"`
		PendingZigbeeMessageMask int    `json:"PendingZigbeeMessageMask,omitempty"`
		BindingsStatus           string `json:"BindingsStatus,omitempty"`
		ReportConfigStatus       string `json:"ReportConfigStatus,omitempty"`
	} `json:"Device"`
	Zigbee struct {
		JPANCount           int    `json:"JPANCount"`
		NetworkChannel      int    `json:"NetworkChannel"`
		ZigbeeModuleVersion string `json:"ZigbeeModuleVersion"`
		ZigbeeEUI           string `json:"ZigbeeEUI"`
	} `json:"Zigbee"`
	RoomStat []struct {
		ID                  int `json:"id"`
		SetPoint            int `json:"SetPoint"`
		MeasuredTemperature int `json:"MeasuredTemperature"`
		MeasuredHumidity    int `json:"MeasuredHumidity"`
	} `json:"RoomStat"`
	DeviceCapabilityMatrix struct {
		Roomstat           bool `json:"Roomstat"`
		ITRV               bool `json:"ITRV"`
		SmartPlug          bool `json:"SmartPlug"`
		UFH                bool `json:"UFH"`
		UFHFloorTempSensor bool `json:"UFHFloorTempSensor"`
		UFHDewSensor       bool `json:"UFHDewSensor"`
		HACT               bool `json:"HACT"`
		LACT               bool `json:"LACT"`
		Light              bool `json:"Light"`
		Shutter            bool `json:"Shutter"`
		LoadController     bool `json:"LoadController"`
	} `json:"DeviceCapabilityMatrix"`
	Schedule []struct {
		ID     int `json:"id"`
		Monday struct {
			SetPoints []struct {
				Time     int `json:"Time"`
				DegreesC int `json:"DegreesC"`
			} `json:"SetPoints"`
		} `json:"Monday"`
		Tuesday struct {
			SetPoints []struct {
				Time     int `json:"Time"`
				DegreesC int `json:"DegreesC"`
			} `json:"SetPoints"`
		} `json:"Tuesday"`
		Wednesday struct {
			SetPoints []struct {
				Time     int `json:"Time"`
				DegreesC int `json:"DegreesC"`
			} `json:"SetPoints"`
		} `json:"Wednesday"`
		Thursday struct {
			SetPoints []struct {
				Time     int `json:"Time"`
				DegreesC int `json:"DegreesC"`
			} `json:"SetPoints"`
		} `json:"Thursday"`
		Friday struct {
			SetPoints []struct {
				Time     int `json:"Time"`
				DegreesC int `json:"DegreesC"`
			} `json:"SetPoints"`
		} `json:"Friday"`
		Saturday struct {
			SetPoints []struct {
				Time     int `json:"Time"`
				DegreesC int `json:"DegreesC"`
			} `json:"SetPoints"`
		} `json:"Saturday"`
		Sunday struct {
			SetPoints []struct {
				Time     int `json:"Time"`
				DegreesC int `json:"DegreesC"`
			} `json:"SetPoints"`
		} `json:"Sunday"`
		Type string `json:"Type"`
	} `json:"Schedule"`
}
