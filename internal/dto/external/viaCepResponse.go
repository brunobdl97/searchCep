package external

import "github.com/searchCep/internal/dto/response"

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v ViaCepResponse) ToResponse() response.CepResponse {
	return response.CepResponse{
		Cep:          v.Cep,
		State:        v.Uf,
		City:         v.Localidade,
		Neighborhood: v.Bairro,
		Street:       v.Logradouro,
		Service:      "ViaCep",
	}
}
