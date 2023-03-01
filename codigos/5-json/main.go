package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 20}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"Numero":4,"Saldo":200}`)
	var novaConta Conta
	err = json.Unmarshal(jsonPuro, &novaConta)
	if err != nil {
		println(err)
	}
	println(novaConta.Numero, novaConta.Saldo)
}
