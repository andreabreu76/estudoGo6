package entities

/*
{
	"status":200,
	"ok":true,
	"code":"21931-390",
	"state":"RJ",
	"city":"Rio de Janeiro",
	"district":"Jardim Guanabara",
	"address":"Rua Rui Vaz Pinto",
	"statusText":"ok"
}
*/
type CepResponse struct {
	Status     string
	OK         string
	Code       string
	State      string
	City       string
	District   string
	Address    string
	StatusText string
}
