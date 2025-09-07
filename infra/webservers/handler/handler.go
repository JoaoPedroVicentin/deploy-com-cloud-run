package handler

import (
	"encoding/json"
	"net/http"

	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/infra/entity"
	"github.com/go-chi/chi/v5"
)

func GetTemperatureOfLocation(w http.ResponseWriter, r *http.Request) {

	cep := chi.URLParam(r, "cep")
	err := entity.IsValidCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	location, err := entity.GetCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	temperature, err := entity.GetLocationTemperature(location.City)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}
