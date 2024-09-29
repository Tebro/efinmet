package vatsimdata

import (
	"encoding/json"
	"fmt"
	"net/http"
)



var dataUrl = "https://data.vatsim.net/v3/vatsim-data.json"

type DataFlightPlan struct {
	FlightRules         string `json:"flight_rules"`
	Aircraft            string `json:"aircraft_short"`
	Departure           string `json:"departure"`
	Arrival             string `json:"arrival"`
	Altitude            string `json:"altitude"`
	Deptime             string `json:"deptime"`
	Remarks             string `json:"remarks"`
	Route               string `json:"route"`
	AssignedTransponder string `json:"assigned_transponder"`
}

type DataPilot struct {
	Cid         int            `json:"cid"`
	Name        string         `json:"name"`
	Callsign    string         `json:"callsign"`
	Transponder string         `json:"transponder"`
	FlightPlan  DataFlightPlan `json:"flight_plan"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Altityde    int            `json:"altitude"`
	GroundSpeed int            `json:"groundspeed"`
}

type data struct {
	//general dataGeneral
	Pilots []DataPilot `json:"pilots"`
}

func PilotsForIcaoPrefix(prefix string, pilots []DataPilot) []DataPilot {
	res := []DataPilot{}
	prefixLen := len(prefix)
	for _, p := range pilots {
		if len(p.FlightPlan.Arrival) >= prefixLen && len(p.FlightPlan.Departure) >= prefixLen {
			startDep := p.FlightPlan.Departure[0:prefixLen]
			startArr := p.FlightPlan.Arrival[0:prefixLen]
			if startDep == prefix || startArr == prefix {
				res = append(res, p)
			}
		}
	}
	return res
}

func GetData() (*data, error) {
	res, err := http.Get(dataUrl)
	if err != nil {
		fmt.Printf("Could not get data: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	var d data
	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		fmt.Printf("Could not parse data: %v", err)
		return nil, err
	}
	return &d, nil
}
