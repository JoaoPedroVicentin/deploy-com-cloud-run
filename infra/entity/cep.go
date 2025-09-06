package entity

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/JoaoPedroVicentin/deploy-com-cloud-run/infra/dto"
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

	var dataViaCep dto.ViaCep
	if err := json.Unmarshal(body, &dataViaCep); err != nil {
		log.Println("Erro ao decodificar ViaCEP:", err)
		return data, err
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

func IsValidCep(cep string) bool {
	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return regex.MatchString(cep)
}
