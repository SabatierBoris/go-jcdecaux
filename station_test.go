package jcdecaux

import (
	"testing"
)

func searchStationByNumber(stations []Station, number int) (station *Station){
	for _,val := range stations{
		if val.Number == number {
			station = &val
			return
		}
	}
	station = nil
	return
}

func TestGetStationsList(t *testing.T){
	list,result := test_client.Stations().All()
	if result.HasError() {
		t.Error("Error during geting the stations list %q",result.Error())
	}else{
		cases := []struct{
			commercial string
			station int
			station_name string
		}{
			{"Paris", 31705, "31705 - CHAMPEAUX (BAGNOLET)"},
			{"Lyon", 2010, "02010 - CONFLUENCE DARSE"},
		}
		for _, val := range cases{
			station := searchStationByNumber(list, val.station)
			if station == nil {
				t.Errorf("Error: %q isn't in stations list", val.station_name)
			}else if station.Name != val.station_name {
				t.Errorf("Error: %q name isn't %q (%q)", val.station, val.station_name, station.Name)
			}
		}
	}
}

func TestGetStationsListFor(t *testing.T){
	cases := []struct{
		commercial string
		station int
		station_name string
	}{
		{"Paris", 31705, "31705 - CHAMPEAUX (BAGNOLET)"},
		{"Lyon", 2010, "02010 - CONFLUENCE DARSE"},
	}
	for _, val := range cases{
		list,result := test_client.Stations().AllFor(val.commercial)
		if result.HasError() {
			t.Errorf("Error during geting the stations list for %q: %q",val.commercial,result.Error())
		}else{
			station := searchStationByNumber(list, val.station)
			if station == nil {
				t.Errorf("Error: %q isn't in stations list", val.station_name)
			}else if station.Name != val.station_name {
				t.Errorf("Error: %q name isn't %q (%q)", val.station, val.station_name, station.Name)
			}
		}
	}
}

func TestGetStationInformation(t *testing.T){
	cases := []struct{
		commercial string
		station int
		station_name string
	}{
		{"Paris", 31705, "31705 - CHAMPEAUX (BAGNOLET)"},
		{"Lyon", 2010, "02010 - CONFLUENCE DARSE"},
	}
	for _, val := range cases {
		station,result := test_client.Stations().InformationsFor(val.station, val.commercial)
		if result.HasError() {
			t.Errorf("Error during geting the stations list for %q: %q",val.commercial,result.Error())
		}else if station.Name != val.station_name{
			t.Errorf("Error: %q name isn't %q (%q)", val.station, val.station_name, station.Name)
		}
	}
}
