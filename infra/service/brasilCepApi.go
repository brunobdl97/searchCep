package service

import (
	"encoding/json"
	"fmt"
	"github.com/searchCep/internal/dto/external"
	"net/http"
)

const BrasilApiPath = "https://brasilapi.com.br/api/cep/v1/%s"

type BasilCepApi struct {
	httpClient *http.Client
}

func NewBrasilCepApi(client *http.Client) *BasilCepApi {
	return &BasilCepApi{
		httpClient: client,
	}
}

func (b *BasilCepApi) GetCep(cep string) (interface{}, error) {
	formatedPath := fmt.Sprintf(BrasilApiPath, cep)
	resp, err := b.httpClient.Get(formatedPath)
	if err != nil {
		return nil, err
	}

	var data external.BrasilApiResponse
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
