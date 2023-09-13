package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/viniciusmgaspar/desafio-multithreading/internal/infra/webServer/services"
)

type CepHandler struct {
	cepService *services.CepService
}

func NewCepHandler() *CepHandler {
	return &CepHandler{
		cepService: services.NewCepService(),
	}
}

func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := h.cepService.GetCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
