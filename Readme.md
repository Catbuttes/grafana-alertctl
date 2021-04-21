# Grafana Alert Control

Control your grafana alerts by the command line.

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
galertctl.go -url <your_grafana_instance> -token <your_api_token> -save
```

To specify where the save file, use 

```
galertctl.go -url <your_grafana_instance> -token <your_api_token> -statefile <file> -save
```

### Restoring Saved Alerts

To restore saved alert settings locally, you can do the following

```
galertctl.go -url <your_grafana_instance> -token <your_api_token> -restore
```

To specify where the save file to restore, use 

```
galertctl.go -url <your_grafana_instance> -token <your_api_token> -statefile <file> -restore
```

### Disabling Alerts

To disable alerts in the ok or unknown state (can be combined with the save option) use 

```
galertctl.go -url <your_grafana_instance> -token <your_api_token> -disable
```

To disregard the alert state and disable even active alerts, add `-force`

```
galertctl.go -url <your_grafana_instance> -token <your_api_token> -disable -force
```

### Enabling Alerts

To enable alerts use 

```
galertctl.go -url <your_grafana_instance> -token <your_api_token> -enable
```