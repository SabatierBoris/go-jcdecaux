package jcdecaux

var (
	ContractsList = Hyperlink("vls/v1/contracts")
)

func (c *Client) Contracts() (contracts *ContractsService){
	contracts = &ContractsService{client: c}
	return
}

type ContractsService struct {
	client *Client
}

func (r *ContractsService) All() (contracts []Contract, result *Result){
	url, _ := ContractsList.Expand(M{})
	result = r.client.get(url, &contracts)
	return
}
