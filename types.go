package jcdecaux

type Contract struct {
	Name string `json:"name"`
	Commercial_name string `json:"commercial_name"`
	Country_code string `json:"country_code"`
	Cities []string `json:"cities"`
}

type Station struct {
	Number int `json:"number"`
	Contract_name string `json:"contract_name"`
	Name string `json:"name"`
	Address string `json:"address"`
	Position struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"position"`
	Banking bool `json:"banking"`
	Bonus bool `json:"bonus"`
	Status string `json:"status"`
	Bike_stands int `json:"bike_stands"`
	Available_bike_stands int `json:"available_bike_stands"`
	Available_bikes int `json:"available_bikes"`
	Last_update int `json:"last_update"`
}
