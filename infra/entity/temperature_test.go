package entity

import "testing"

func TestShouldGetTemperatureOfLocation(t *testing.T) {
	location := "Sâo Paulo"
	temperature, err := GetLocationTemperature(location)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if temperature.TempC == 0 && temperature.TempF == 0 && temperature.TempK == 0 {
		t.Errorf("Expected valid temperature values, but got zeros")
	}
}

func TestShouldReturnErrorWhenLocationNotFound(t *testing.T) {
	location := "InvalidCityName"
	_, err := GetLocationTemperature(location)
	if err == nil {
		t.Errorf("Expected error for invalid location, but got none")
	}
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	celsius := 25.0
	expectedKelvin := 298.15
	kelvin := convertCelsiusToKelvin(celsius)
	if kelvin != expectedKelvin {
		t.Errorf("Expected %fK, but got %fK", expectedKelvin, kelvin)
	}
}
