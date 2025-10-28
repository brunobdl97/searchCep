package external

import "github.com/searchCep/internal/dto/response"

type BrasilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (b BrasilApiResponse) ToResponse() response.CepResponse {
	return response.CepResponse{
		Cep:          b.Cep,
		State:        b.State,
		City:         b.City,
		Neighborhood: b.Neighborhood,
		Street:       b.Street,
		Service:      "BrasilApi",
	}
}
