package services

import (
	"log"
	"time"

	"github.com/viniciusmgaspar/desafio-multithreading/pkg/services"
)

type CepService struct {
	ch1       chan services.ViaCepResponse
	ch2       chan services.ApiCepResponse
	viaCepSrv *services.ViaCepService
	apiCepSrv *services.ApiCepService
}

func NewCepService() *CepService {
	return &CepService{
		ch1:       make(chan services.ViaCepResponse),
		ch2:       make(chan services.ApiCepResponse),
		viaCepSrv: services.NewViaCepService(),
		apiCepSrv: services.NewApiCepService(),
	}
}

func (s *CepService) GetCep(cep string) (any, error) {
	go s.ViaCepTask(s.ch1, cep)
	go s.ApiCepTask(s.ch2, cep)
	select {
	case viaCepResponse := <-s.ch1:
		log.Printf("{ Api: %s, Response: %v }", "ViaCep", viaCepResponse)
		return viaCepResponse, nil
	case apiCepResponse := <-s.ch2:
		log.Printf("{ Api: %s, Response: %v }", "ApiCep", apiCepResponse)
		return apiCepResponse, nil
	case <-time.After(1 * time.Second):
		log.Print("Timeout")
		res := map[string]interface{}{
			"msg":    "Response timeout",
			"status": 408,
		}
		return res, nil
	}
}

func (s *CepService) ViaCepTask(c chan services.ViaCepResponse, cep string) {
	response, err := s.viaCepSrv.GetCep(cep)
	if err != nil {
		log.Print(err)
	}
	c <- *response
}

func (s *CepService) ApiCepTask(c chan services.ApiCepResponse, cep string) {
	response, err := s.apiCepSrv.GetCep(cep)
	if err != nil {
		log.Print(err)
	}

	c <- *response
}
