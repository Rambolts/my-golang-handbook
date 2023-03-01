package main

import (
	"errors"
	"fmt"
)

type Conta struct {
	dono   string
	saldo  int
	divida int
}

func NewConta(nome string) *Conta {
	return &Conta{dono: nome, saldo: 0, divida: 20}
}

func (c *Conta) depositar(valor int) {
	c.saldo += valor
}

func (c *Conta) pagarDivida(valor int) (err error) {
	if valor > c.divida {
		valor = c.divida
	}
	if valor > c.saldo {
		return errors.New("Valor insuficiente")
	}
	c.divida -= valor
	c.saldo -= valor
	return nil
}

func main() {
	conta := NewConta("Jose")
	conta.depositar(50)
	err := conta.pagarDivida(25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conta)
}
