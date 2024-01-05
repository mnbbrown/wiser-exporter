### wiser-exporter

configured via env variables:

```bash
export WISER_SECRET=$SECRET // wiser network secret key
export WISER_HUB_IP=192.168.0.126 // wiser hub IP adddress
export PORT=2112 // http server port
export INTERVAL=30s // how often to read
```

run the exporter:

```bash
go build -o wiser-exporter ./ && ./wiser-exporter
```

query the metrics:

```
# HELP wiser_channel_percent_demand The percent wiser is demanding from a channel
# TYPE wiser_channel_percent_demand gauge
wiser_channel_percent_demand{id="1",name="Channel-1"} 0
# HELP wiser_room_heating_rate The heating rate of the room
# TYPE wiser_room_heating_rate gauge
wiser_room_heating_rate{id="1",name="Living Room"} 1200
# HELP wiser_room_humidity The humidity of a room
# TYPE wiser_room_humidity gauge
wiser_room_humidity{id="1",name="Living Room"} 54
# HELP wiser_room_set_point The set point of a room
# TYPE wiser_room_set_point gauge
wiser_room_set_point{id="1",name="Living Room"} 19
# HELP wiser_room_temperature The temperature of a room
# TYPE wiser_room_temperature gauge
wiser_room_temperature{id="1",name="Living Room"} 19.8
```
