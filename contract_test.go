package jcdecaux

import (
	"testing"
)

func searchContractByName(contracts []Contract, name string) (contract *Contract){
	for _,val := range contracts{
		if val.Name == name {
			contract = &val
			return
		}
	}
	contract = nil
	return
}


func TestGetContractsList(t *testing.T){
	list,result := test_client.Contracts().All()
	if result.HasError() {
		t.Error("Error during geting the contracts list %q",result.Error())
	}else{
		cases := []struct{
			name, commercial string
		}{
			{"Paris", "Velib"},
			{"Lyon", "Vélo'V"},
		}
		for _, val := range cases{
			contract := searchContractByName(list, val.name)
			if contract == nil {
				t.Errorf("error: %q isn't in contracts list", val.name)
			}else if contract.Commercial_name != val.commercial {
				t.Errorf("error: %q commercial name isn't %q (%q)", val.name, val.commercial, contract.Commercial_name)
			}
		}
	}
}
