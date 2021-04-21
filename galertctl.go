package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	instance  = flag.String("url", "https://localhost:3000", "The URL of the Grafana Instance")
	token     = flag.String("token", "", "The API token for the instance")
	stateFile = flag.String("statefile", "./alerts.gastate", "The file to save/restore state to/from")

	save    = flag.Bool("save", false, "Save the alert state")
	restore = flag.Bool("restore", false, "Restore alerts")

	disable = flag.Bool("disable", false, "Disable alerts")
	enable  = flag.Bool("enable", false, "Enable alerts")

	force = flag.Bool("force", false, "Ignore active alerts when changing state")
)

func HelpText() {
	fmt.Println("galertctl - Control Grafana alerts in bulk")

	flag.PrintDefaults()
}

func main() {
	flag.Usage = HelpText
	flag.Parse()

	if *save {
		state := getAlerts()
		saveState(state)
	}

	if *restore {
		state := loadState()
		_ = state

	}

}

func getAlerts() []GAState {

	alertsURL := *instance + "/api/alerts"

	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
	}

	request, err := http.NewRequest("GET", alertsURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", "Bearer "+*token)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseData ApiAlertsResponse
	json.Unmarshal(data, &responseData)

	GAStateData := make([]GAState, 0)

	for _, alert := range responseData {
		alertStateData := GAState{
			ID:          alert.ID,
			State:       alert.State,
			InstanceURL: *instance,
		}
		GAStateData = append(GAStateData, alertStateData)
	}

	//fmt.Printf("%+v", GAStateData)

	return GAStateData
}

func saveState(states []GAState) {

	var existingState = make([]GAState, 0)
	if _, err := os.Stat(*stateFile); err == nil {
		existingState = loadState()
	}

	existingState = append(existingState, states...)

	data, err := json.Marshal(existingState)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(*stateFile, data, 0644)
}

func loadState() []GAState {

	stateContents, err := ioutil.ReadFile(*stateFile)
	if err != nil {
		log.Fatal(err)
	}
	state := make([]GAState, 0)
	json.Unmarshal([]byte(stateContents), &state)

	return state
}

func setState(paused bool) {

}

type ApiAlertsResponse []struct {
	ID            int       `json:"id"`
	Dashboardid   int       `json:"dashboardId"`
	Dashboarduid  string    `json:"dashboardUid"`
	Dashboardslug string    `json:"dashboardSlug"`
	Panelid       int       `json:"panelId"`
	Name          string    `json:"name"`
	State         string    `json:"state"`
	Newstatedate  time.Time `json:"newStateDate"`
	Evaldate      time.Time `json:"evalDate"`
	Evaldata      struct {
	} `json:"evalData"`
	Executionerror string `json:"executionError"`
	URL            string `json:"url"`
}

type GAState struct {
	ID          int    `json:"id"`
	State       string `json:"state"`
	InstanceURL string `json:"instanceurl"`
}