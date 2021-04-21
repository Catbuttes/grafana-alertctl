# Grafana Alert Control

Control your grafana alerts by the command line.

## Installation

Install either by downloading one of the prebuilt binaries or by running `go get github.com/catbuttes/grafana-alertctl`

## Usage

```
Grafana Alert Control - Control Grafana alerts in bulk
  -disable
    	Disable alerts
  -enable
    	Enable alerts
  -force
    	Ignore active alerts when changing state
  -restore
    	Restore alerts
  -save
    	Save the alert state
  -statefile string
    	The file to save/restore state to/from (default "./alerts.gastate")
  -token string
    	The API token for the instance
  -url string
    	The URL of the Grafana Instance (default "https://localhost:3000")
```

### Saving Alerts

To save your alert settings locally, you can do the following

```
grafana-alertctl -url <your_grafana_instance> -token <your_api_token> -save
```

To specify where the save file, use 

```
grafana-alertctl -url <your_grafana_instance> -token <your_api_token> -statefile <file> -save
```

### Restoring Saved Alerts

To restore saved alert settings locally, you can do the following

```
grafana-alertctl -url <your_grafana_instance> -token <your_api_token> -restore
```

To specify where the save file to restore, use 

```
grafana-alertctl -url <your_grafana_instance> -token <your_api_token> -statefile <file> -restore
```

### Disabling Alerts

To disable alerts in the ok or unknown state (can be combined with the save option) use 

```
grafana-alertctl -url <your_grafana_instance> -token <your_api_token> -disable
```

To disregard the alert state and disable even active alerts, add `-force`

```
grafana-alertctl -url <your_grafana_instance> -token <your_api_token> -disable -force
```

### Enabling Alerts

To enable alerts use 

```
grafana-alertctl -url <your_grafana_instance> -token <your_api_token> -enable
```