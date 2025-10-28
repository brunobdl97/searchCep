# searchCep

Small Go HTTP service that resolves Brazilian CEPs (postal codes). The handler calls both [BrasilAPI](https://brasilapi.com.br/) and [ViaCEP](https://viacep.com.br/) concurrently and returns the first successful response in a normalized JSON format.

## Prerequisites
- Go 1.20+ installed locally.
- Network access to brasilapi.com.br and viacep.com.br.

## Running the server
```bash
go run ./...
```
The service starts on `http://localhost:8080` and exposes a single route:

- `GET /cep/{cep}` — replace `{cep}` with an 8-digit CEP (digits only).

## Example request
Using curl:
```bash
curl http://localhost:8080/cep/30280040
```
Typical response:
```json
{
  "cep": "30280040",
  "state": "MG",
  "city": "Belo Horizonte",
  "neighborhood": "Santa Efigênia",
  "street": "Rua do Ouro",
  "service": "ViaCep"
}
```
The `service` field identifies which provider returned first.

## Using the HTTP client file
The `tests/generated-requests.http` file contains a ready-to-run request. Open it with an IDE or extension that supports `.http` files (JetBrains HTTP Client, VS Code REST Client, etc.) and execute the request once the server is running.

## Troubleshooting
- If you receive a timeout, there are some comments simulating a delay on each api call in the cepHandler file, just remove them.
- CEP must contain only digits; malformed CEPs return `400 Bad Request`.
