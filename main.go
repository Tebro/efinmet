package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"time"
)

var dataUrl = "https://data.vatsim.net/v3/vatsim-data.json"
var metarUrl = "https://www.ilmailusaa.fi/backend.php?{%22mode%22:%22metar%22,%22radius%22:%22100%22,%22points%22:[{%22_area%22:%221%22}]}"

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

func pilotsForIcaoPrefix(prefix string, pilots []DataPilot) []DataPilot {
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

func getData() (*data, error) {
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

type metar struct {
	P1  string `json:"p1"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func getMetars() (*map[string]string, error) {
	res, err := http.Get(metarUrl)
	if err != nil {
		fmt.Printf("Could not get metars: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	var d map[string]metar
	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		fmt.Printf("Could not parse metars: %v", err)
		return nil, err
	}

	result := map[string]string{}

	metarWord := "METAR"
	lenPrefix := len(metarWord)

	for k, v := range d {
		key := k[lenPrefix:]
		result[string(key)] = v.P1
	}

	return &result, nil
}

type field struct {
	ICAO   string
	NumDep int
	NumArr int
	Metar  string
}

func fieldIcaoHasPrefix(prefix string, fieldIcao string) bool {
	return len(fieldIcao) >= len(prefix) && fieldIcao[0:len(prefix)] == prefix
}

func buildFieldsFromDataAndMetars(pilots []DataPilot, metars *map[string]string) map[string]*field {
	fields := map[string]*field{}

	hasField := func(fieldName string) bool {
		_, ok := fields[fieldName]
		return ok
	}

	addField := func(fieldName string) {
		fields[fieldName] = &field{
			ICAO:  fieldName,
			Metar: (*metars)[fieldName],
		}
	}

	for _, p := range pilots {
		if fieldIcaoHasPrefix("EF", p.FlightPlan.Arrival) {
			if !hasField(p.FlightPlan.Arrival) {
				addField(p.FlightPlan.Arrival)
			}
			f := fields[p.FlightPlan.Arrival]
			f.NumArr += 1
		}
		if fieldIcaoHasPrefix("EF", p.FlightPlan.Departure) {
			if !hasField(p.FlightPlan.Departure) {
				addField(p.FlightPlan.Departure)
			}
			f := fields[p.FlightPlan.Departure]
			f.NumDep += 1
		}
	}

	return fields
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func distanceBetween(lat1, lon1, lat2, lon2 float64) float64 {
	r := 6371.0 // Radius of the Earth in km
	x := (lon2 - lon1) * math.Pi / 180 * math.Cos(((lat1+lat2)/2)*math.Pi/180)
	y := (lat2 - lat1) * math.Pi / 180
	return r * math.Sqrt(x*x+y*y)
}

func pilotsWithinRangeLimits(in []DataPilot, airports AirportsInfo) (out []DataPilot) {
	for _, pilot := range in {
		arrAirport, ok := airports[pilot.FlightPlan.Arrival]
		if ok { // Pilot arriving in Finland
			distance := distanceBetween(pilot.Latitude, pilot.Longitude, arrAirport.Lat, arrAirport.Lon)
			if distance <= 300*1.852 {
				out = append(out, pilot)
			}
		}
		depAirport, ok := airports[pilot.FlightPlan.Departure]
		if ok { // Pilot is departing from Finland
			distance := distanceBetween(pilot.Latitude, pilot.Longitude, depAirport.Lat, depAirport.Lon)
			if distance < 10*1.852 {
				out = append(out, pilot)
			}
		}
	}
	return
}

func main() {

	airports, err := GetAirportsInfo()
	if err != nil {
		fmt.Printf("Unable to read aiports: %v\n", err)
	}

	for {
		d, err := getData()
		if err != nil {
			fmt.Println("No data")
			return
		}

		pilotsForEfin := pilotsForIcaoPrefix("EF", d.Pilots)
		pilotsWithinRangeLimits := pilotsWithinRangeLimits(pilotsForEfin, airports)

		metars, err := getMetars()
		if err != nil {
			fmt.Println("No metars")
			return
		}

		result := buildFieldsFromDataAndMetars(pilotsWithinRangeLimits, metars)

		fields := make([]string, 0, len(result))
		for k := range result {
			fields = append(fields, k)
		}
		sort.Strings(fields)

		clear()
		fmt.Println("EFIN Met, relevant METARS")
		fmt.Println("Field: (dep, arr), metar")
		fmt.Println("--------------------------")
		for _, f := range fields {
			v := result[f]
			if v.Metar != "" {
				fmt.Printf("%s: (%d, %d) %s\n", f, v.NumDep, v.NumArr, v.Metar)
			}
		}
		time.Sleep(30 * time.Second)
	}
}
