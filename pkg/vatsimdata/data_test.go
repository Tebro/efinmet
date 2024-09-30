package vatsimdata

import (
	"reflect"
	"testing"
)

func TestPilotsForIcaoPrefix(t *testing.T) {
	tests := []struct {
		prefix string
		pilots []DataPilot
		want   []DataPilot
	}{
		{
			prefix: "KJFK",
			pilots: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "KJFK", Arrival: "KLAX"}},
				{FlightPlan: DataFlightPlan{Departure: "KLAX", Arrival: "KJFK"}},
				{FlightPlan: DataFlightPlan{Departure: "EGLL", Arrival: "KJFK"}},
			},
			want: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "KJFK", Arrival: "KLAX"}},
				{FlightPlan: DataFlightPlan{Departure: "KLAX", Arrival: "KJFK"}},
				{FlightPlan: DataFlightPlan{Departure: "EGLL", Arrival: "KJFK"}},
			},
		},
		{
			prefix: "KLAX",
			pilots: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "KJFK", Arrival: "KLAX"}},
				{FlightPlan: DataFlightPlan{Departure: "KLAX", Arrival: "KJFK"}},
				{FlightPlan: DataFlightPlan{Departure: "EGLL", Arrival: "KJFK"}},
			},
			want: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "KJFK", Arrival: "KLAX"}},
				{FlightPlan: DataFlightPlan{Departure: "KLAX", Arrival: "KJFK"}},
			},
		},
		{
			prefix: "EGLL",
			pilots: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "KJFK", Arrival: "KLAX"}},
				{FlightPlan: DataFlightPlan{Departure: "KLAX", Arrival: "KJFK"}},
				{FlightPlan: DataFlightPlan{Departure: "EGLL", Arrival: "KJFK"}},
			},
			want: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "EGLL", Arrival: "KJFK"}},
			},
		},
		{
			prefix: "EF",
			pilots: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "KJFK", Arrival: "EFHK"}},
				{FlightPlan: DataFlightPlan{Departure: "EFTU", Arrival: "EFRO"}},
				{FlightPlan: DataFlightPlan{Departure: "EGLL", Arrival: "KJFK"}},
			},
			want: []DataPilot{
				{FlightPlan: DataFlightPlan{Departure: "KJFK", Arrival: "EFHK"}},
				{FlightPlan: DataFlightPlan{Departure: "EFTU", Arrival: "EFRO"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			got := PilotsForIcaoPrefix(tt.prefix, tt.pilots)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PilotsForIcaoPrefix(%v, %v) = %v; want %v", tt.prefix, tt.pilots, got, tt.want)
			}
		})
	}
}
