package handler

import (
	"encoding/json"
	"net/http"

	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/configs"
	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/infra/entity"
	"github.com/go-chi/chi/v5"
)

func GetTemperatureOfLocation(w http.ResponseWriter, r *http.Request) {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	cep := chi.URLParam(r, "cep")
	isValidCep := entity.IsValidCep(cep)

	if !isValidCep {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid zipcode"})
		return
	}
	location, err := entity.GetCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot find zipcode"})
		return
	}
	temperature, err := entity.GetLocationTemperature(location.City, configs.ApiKey)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot find temperature for: " + location.City})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}
