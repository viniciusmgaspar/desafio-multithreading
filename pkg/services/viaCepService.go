package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCepResponse struct {
	Bairro      string `json:"bairro,omitempty"`
	CEP         string `json:"cep"`
	Complemento string `json:"complemento,omitempty"`
	DDD         string `json:"ddd"`
	GIA         string `json:"gia"`
	IBGE        string `json:"ibge"`
	Localidade  string `json:"localidade"`
	Logradouro  string `json:"logradouro,omitempty"`
	SIAFI       string `json:"siafi"`
	UF          string `json:"uf"`
}

type ViaCepService struct {
	baseUrl string
}

func NewViaCepService() *ViaCepService {
	return &ViaCepService{
		baseUrl: "http://viacep.com.br/ws",
	}
}

func (s *ViaCepService) GetCep(cep string) (*ViaCepResponse, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s/json", s.baseUrl, cep))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resp ViaCepResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
