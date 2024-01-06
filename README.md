### wiser-exporter

Uses the Drayton Wiser HubR API to make the following metrics available:

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

#### Retrieving the secret

To use this exporter you need to find your network secret.

1. Press the "SETUP" button once on the front of the HubR.
2. Connect to it's wifi network
3. Load `192.168.8.1/secret` in your browser
5. Copy the secret into an env variable or the defaults file (see below)
6. Press the "SETUP" button once again to return the HubR to normal mode.

#### Install on ubuntu

1. Download the latest deb from Github releases
2. `dpkg -i ./wiser-exporter.deb` to install
3. Edit `/etc/default/wiser-exporter` and update secret + hub IP
4. Run `systemctl restart wiser-exporter`


#### Build from source

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
