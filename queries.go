package main

// getHumidity returns the humidity reported by a room's RoomStat, or -1 if
// the room has no RoomStat or it currently has no reading. While a RoomStat
// is offline (e.g. rejoining after a hub reboot) the hub reports 0 % humidity
// alongside the -3276.8 °C no-reading temperature; neither is a measurement.
func getHumidity(d *DomainData, id int) int {
	for _, stat := range d.RoomStat {
		if stat.ID == id {
			if stat.MeasuredTemperature == noReading || stat.MeasuredHumidity <= 0 {
				return -1
			}
			return stat.MeasuredHumidity
		}
	}

	return -1
}
