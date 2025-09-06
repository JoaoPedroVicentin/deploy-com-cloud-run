package entity

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/infra/dto"
)

func GetLocationTemperature(location string, apyKey string) (temp dto.Temperature, err error) {

	weatherData, err := getWeatherApi(location, apyKey)
	if err != nil {
		log.Println("Erro ao buscar temperatura:", err)
		return temp, err
	}

	return dto.Temperature{
		TempC: weatherData.Current.TempC,
		TempF: weatherData.Current.TempF,
		TempK: convertCelsiusToKelvin(weatherData.Current.TempC),
	}, nil
}

func getWeatherApi(city string, apiKey string) (data dto.WeatherData, err error) {

	url := "http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + city + "&aqi=no"

	res, err := http.Get(url)
	if err != nil {
		log.Println("Erro WeatherApi:", err)
		return data, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler resposta WeatherApi:", err)
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Erro ao decodificar WeatherApi:", err)
		return data, err
	}

	return data, nil
}

func convertCelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}
