package jcdecaux

import (
	"fmt"
	"testing"
	"os"
	"encoding/json"
)

var (
	test_client *Client
)

type Configuration struct {
	JCDecaux_APIKey string
}

func setup(){
	file, _ := os.Open("testing.config")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error :",err)
	}

	test_client = NewClient(configuration.JCDecaux_APIKey)
}

func teardown(){
}

func TestMain(m *testing.M){
	setup()
	retCode := m.Run()

	teardown()

	os.Exit(retCode)
}
