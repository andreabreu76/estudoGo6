package entities

type CepResponse struct {
	Status     int    `json:"status"`
	OK         bool   `json:"ok"`
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	StatusText string `json:"statusText"`
}
