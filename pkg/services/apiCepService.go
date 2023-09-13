package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiCepResponse struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

type ApiCepService struct {
	baseUrl string
}

func NewApiCepService() *ApiCepService {
	return &ApiCepService{
		baseUrl: "https://cdn.apicep.com/file/apicep",
	}
}

func (s *ApiCepService) GetCep(cep string) (*ApiCepResponse, error) {
	cepParam := fmt.Sprintf("%s-%s", cep[:5], cep[5:])
	res, err := http.Get(fmt.Sprintf("%s/%s/json", s.baseUrl, cepParam))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	println(string(body))
	if err != nil {
		return nil, err
	}
	var resp ApiCepResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
