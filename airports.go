package main

import (
	"encoding/json"
)

type airportsInfo struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type AirportsInfo map[string]airportsInfo

func parseAirportsMap() (AirportsInfo, error) {
	var topLevel map[string][]json.RawMessage
	err := json.Unmarshal([]byte(airportsJSON), &topLevel)
	if err != nil {
		return nil, err
	}

	var airportsMap AirportsInfo = make(AirportsInfo)

	for k, v := range topLevel {
		var info airportsInfo
		err := json.Unmarshal(v[0], &info)
		if err != nil {
			return nil, err
		}
		airportsMap[k] = info
	}

	return airportsMap, err
}

var airportsInfoCache AirportsInfo = make(AirportsInfo)

func GetAirportsInfo() (AirportsInfo, error) {
	if len(airportsInfoCache) == 0 {
		data, err := parseAirportsMap()
		if err != nil {
			return nil, err
		}
		airportsInfoCache = data
	}
	return airportsInfoCache, nil
}
