package internal

type CepServiceInterface interface {
	GetCep(cep string) (interface{}, error)
}
