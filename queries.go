package main

func getHumidity(d *DomainData, id int) int {
	for _, stat := range d.RoomStat {
		if stat.ID == id {
			return stat.MeasuredHumidity
		}
	}

	return -1
}
