package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"time"
	"github.com/Tebro/efinmet/pkg/utils"
	"github.com/Tebro/efinmet/pkg/vatsimdata"
)

var metarUrl = "https://www.ilmailusaa.fi/backend.php?{%22mode%22:%22metar%22,%22radius%22:%22100%22,%22points%22:[{%22_area%22:%221%22}]}"

type metar struct {
	P1  string `json:"p1"`
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func getMetars() (*map[string]string, error) {
	req, err := http.NewRequest("GET", metarUrl, nil)
	if err != nil {
		fmt.Printf("could not create metars request: %v", err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json, text/javascript, */*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Referer", "https://www.ilmailusaa.fi/weather-flightpath.html?location")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0")

	res, err := http.DefaultClient.Do(req)
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

func buildFieldsFromDataAndMetars(pilots []vatsimdata.DataPilot, metars *map[string]string) map[string]*field {
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
		if utils.FieldIcaoHasPrefix("EF", p.FlightPlan.Arrival) {
			if !hasField(p.FlightPlan.Arrival) {
				addField(p.FlightPlan.Arrival)
			}
			f := fields[p.FlightPlan.Arrival]
			f.NumArr += 1
		}
		if utils.FieldIcaoHasPrefix("EF", p.FlightPlan.Departure) {
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
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
		break
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
		break
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func distanceBetween(lat1, lon1, lat2, lon2 float64) float64 {
	r := 6371.0 // Radius of the Earth in km
	x := (lon2 - lon1) * math.Pi / 180 * math.Cos(((lat1+lat2)/2)*math.Pi/180)
	y := (lat2 - lat1) * math.Pi / 180
	return r * math.Sqrt(x*x+y*y)
}

func pilotsWithinRangeLimits(in []vatsimdata.DataPilot, airports AirportsInfo) (out []vatsimdata.DataPilot) {
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
		d, err := vatsimdata.GetData()
		if err != nil {
			fmt.Println("No data")
			return
		}

		pilotsForEfin := vatsimdata.PilotsForIcaoPrefix("EF", d.Pilots)
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
