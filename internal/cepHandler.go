package internal

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/searchCep/infra/service"
	"github.com/searchCep/internal/dto/external"
	"net/http"
	"time"
)

type CepHandler struct {
	client *http.Client
}

func NewCepHandler() *CepHandler {
	return &CepHandler{
		client: &http.Client{},
	}
}

func (h *CepHandler) Execute(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return
	}

	ch := make(chan interface{})
	errCh := make(chan error)

	go h.getByBrasilApi(cep, ch, errCh)
	go h.getByViaCep(cep, ch, errCh)

	select {
	case err := <-errCh:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case res := <-ch:
		fmt.Println(res)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
		return
	case <-time.After(1 * time.Second):
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		return
	}
}

func (h *CepHandler) getByBrasilApi(
	cep string,
	ch chan<- interface{},
	errCh chan<- error,
) {
	time.Sleep(1 * time.Second) // Simulating a delay
	getCep, err := service.NewBrasilCepApi(h.client).GetCep(cep)
	if err != nil {
		errCh <- err
		return
	}

	ch <- getCep.(external.BrasilApiResponse).ToResponse()
}

func (h *CepHandler) getByViaCep(
	cep string,
	ch chan<- interface{},
	errCh chan<- error,
) {
	time.Sleep(1 * time.Second) // Simulating a delay
	getCep, err := service.NewViaCepApi(h.client).GetCep(cep)
	if err != nil {
		errCh <- err
		return
	}

	ch <- getCep.(external.ViaCepResponse).ToResponse()
}
