package entity

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/infra/dto"
)

var (
	ErrorInvalidCep = errors.New("invalid zipcode")
	ErrorNotFound   = errors.New("zipcode not found")
)

func GetCep(cep string) (data dto.Address, err error) {
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	res, err := http.Get(url)
	if err != nil {
		log.Println("Erro ViaCEP:", err)
		return data, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler resposta ViaCEP:", err)
		return data, err
	}

	if strings.Contains(string(body), "erro") {
		log.Println("CEP não encontrado:", cep)
		return data, ErrorNotFound
	}

	var dataViaCep dto.ViaCep
	if err := json.Unmarshal(body, &dataViaCep); err != nil {
		log.Println("Erro ao decodificar ViaCEP:", err)
		return data, ErrorInvalidCep
	}

	return dto.Address{
		Cep:          dataViaCep.Cep,
		State:        dataViaCep.Uf,
		City:         dataViaCep.Localidade,
		Neighborhood: dataViaCep.Bairro,
		Street:       dataViaCep.Logradouro,
		API:          "ViaCEP",
	}, nil
}

func IsValidCep(cep string) error {
	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	if !regex.MatchString(cep) {
		return ErrorInvalidCep
	}
	return nil
}
