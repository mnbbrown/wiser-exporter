### wiser-exporter

Uses the Drayton Wiser HubR local API (`/data/v2/domain/` and `/data/v2/opentherm/`) to export Prometheus metrics.

| Metric | Labels | Notes |
|---|---|---|
| `wiser_room_temperature` | id, name | °C. Omitted while the room has no reading |
| `wiser_room_set_point` | id, name | °C |
| `wiser_room_humidity` | id, name | %, rooms with a RoomStat only |
| `wiser_room_demand` | id, name | % |
| `wiser_room_heating_rate` | id, name | |
| `wiser_channel_percent_demand` | id, name | % |
| `wiser_channel_relay_on` | id, name | 1 when the heating relay is on |
| `wiser_device_battery_voltage` | id, product_type, serial_number | Omitted for mains-powered or offline devices |
| `wiser_opentherm_connected` | | 1 when the hub's OpenTherm link is up |
| `wiser_opentherm_enabled` | | |
| `wiser_opentherm_flow_temperature_celsius` | | |
| `wiser_opentherm_return_temperature_celsius` | | |
| `wiser_opentherm_ch_pressure_bar` | | |
| `wiser_opentherm_relative_modulation_percent` | | Holds its last value while the burner is off |
| `wiser_opentherm_minimum_modulation_percent` | | |
| `wiser_opentherm_maximum_capacity_kw` | | |
| `wiser_opentherm_ch_enabled` | | 1 while the hub is requesting CH |
| `wiser_opentherm_ch_flow_setpoint_celsius` | | Flow temperature the hub is requesting |
| `wiser_opentherm_dhw_enabled` | | |
| `wiser_opentherm_dhw_setpoint_celsius` | | |
| `wiser_opentherm_dhw_flow_rate_litres_per_minute` | | |
| `wiser_opentherm_messages_sent` | | Since hub boot |
| `wiser_opentherm_decoding_errors` | | Since hub boot |
| `wiser_opentherm_slave_status` | | Raw OpenTherm status byte |
| `wiser_opentherm_status_flag` | flag | `fault`, `ch_active`, `dhw_active`, `flame`, `cooling_active`, `ch2_active`, `diagnostic` |
| `wiser_scrape_errors_total` | endpoint | |
| `wiser_last_scrape_success_timestamp_seconds` | endpoint | |
| `wiser_exporter_build_info` | version, commit, date | |

#### Retrieving the secret

To use this exporter you need to find your network secret.

1. Press the "SETUP" button once on the front of the HubR.
2. Connect to its wifi network
3. Load `192.168.8.1/secret` in your browser
4. Copy the secret into an env variable or the defaults file (see below)
5. Press the "SETUP" button once again to return the HubR to normal mode.

#### Install on ubuntu

1. Download the latest deb from GitHub releases
2. `dpkg -i ./wiser-exporter_<version>_linux_amd64.deb` to install
3. Edit `/etc/default/wiser-exporter` and update secret + hub IP
4. Run `systemctl restart wiser-exporter`


#### Build from source

configured via env variables:

```bash
export WISER_SECRET=$SECRET // wiser network secret key
export WISER_HUB_IP=192.168.0.126 // wiser hub IP address
export PORT=2112 // http server port
export INTERVAL=30s // how often to read
```

run the exporter:

```bash
go build -o wiser-exporter ./ && ./wiser-exporter
```
