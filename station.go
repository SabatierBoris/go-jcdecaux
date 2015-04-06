package jcdecaux

var (
	StationsList = Hyperlink("vls/v1/stations")
	StationsListFor  = Hyperlink("vls/v1/stations?contract={contract}")
	StationInfo = Hyperlink("vls/v1/stations/{station_number}?contract={contract}")
)

func (c *Client) Stations() (stations *StationsService){
	stations = &StationsService{client: c}
	return
}

type StationsService struct {
	client *Client
}

func (r *StationsService) All() (stations []Station, result *Result){
	url, _ := StationsList.Expand(M{})
	result = r.client.get(url, &stations)
	return
}

func (r *StationsService) AllFor(contract_name string) (stations []Station, result *Result){
	url, _ := StationsListFor.Expand(M{"contract": contract_name})
	result = r.client.get(url, &stations)
	return
}

func (r *StationsService) InformationsFor(station_number int, contract_name string) (station Station, result *Result){
	url, _ := StationInfo.Expand(M{"station_number": station_number, "contract": contract_name})
	result = r.client.get(url, &station)
	return
}

