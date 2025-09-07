package entity

import (
	"testing"
)

func TestShouldGetCepAndReturnAddress(t *testing.T) {
	cep := "01001-000"
	address, err := GetCep(cep)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if address.Cep != "01001-000" {
		t.Errorf("Expected CEP '01001-000', but got %s", address.Cep)
	}
	if address.City != "São Paulo" {
		t.Errorf("Expected City 'São Paulo', but got %s", address.City)
	}
	if address.State != "SP" {
		t.Errorf("Expected State 'SP', but got %s", address.State)
	}
}

func TestShouldReturnErrorWhenCepNotFound(t *testing.T) {
	cep := "00000-000"
	_, err := GetCep(cep)
	if err == nil {
		t.Errorf("Expected error for invalid CEP, but got none")
	}
}

func TestIsValidCep(t *testing.T) {
	validCeps := []string{"12345-678", "12345678"}
	invalidCeps := []string{"1234-567", "1234567", "123456789", "12a45-678", "12345-67b"}

	for _, cep := range validCeps {
		err := IsValidCep(cep)
		if err != nil {
			t.Errorf("Expected CEP %s to be valid, but got invalid", cep)
		}
	}

	for _, cep := range invalidCeps {
		err := IsValidCep(cep)
		if err == nil {
			t.Errorf("Expected CEP %s to be invalid, but got valid", cep)
		}
	}
}
