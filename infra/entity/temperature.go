package entity

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/configs"
	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/infra/dto"
)

var (
	ErrorCannotFindTemperature = "cannot find temperature for this location: %s"
)

func GetLocationTemperature(location string) (temp dto.Temperature, err error) {

	weatherData, err := getWeatherApi(location)
	if err != nil {
		log.Println("Erro ao buscar temperatura:", err)
		return temp, fmt.Errorf(ErrorCannotFindTemperature, location)
	}

	return dto.Temperature{
		TempC: weatherData.Current.TempC,
		TempF: weatherData.Current.TempF,
		TempK: convertCelsiusToKelvin(weatherData.Current.TempC),
	}, nil
}

func getWeatherApi(city string) (data dto.WeatherData, err error) {

	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	locationEncode := url.QueryEscape(city)

	endpoint := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", configs.ApiKey, locationEncode)

	res, err := http.Get(endpoint)
	if err != nil {
		log.Println("Erro WeatherApi:", err)
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Erro na resposta da API: status code %d", res.StatusCode)
		return data, fmt.Errorf("erro da API: status code %d", res.StatusCode)
	}

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
