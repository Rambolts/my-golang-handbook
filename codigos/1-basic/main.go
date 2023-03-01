package main

// SUMMARY
//	1. Array 1
//	2. Array 2
//	3. Slices
//	4. Maps (dict)
//	5. Funções
//	6. Closures
//	7. Ponteiros
//	8. Generics
//  9. Constraints

import (
	"fmt"
)

var (
	a bool   = true
	b int    = 5
	c string = "Olamundo"
)

// CONSTRAINT (Variável com mais de um tipo)
type MyNumber int

type Number interface {
	~int | float64 // ~ serve para considerar outros tipos derivados do tipo em questão
}

func main() {
	fmt.Println("--------------------------------------")
	// ARRAY 1

	var meuArray [3]int
	meuArray[0] = 3
	meuArray[1] = 2
	meuArray[2] = 6
	fmt.Printf("Último item do array é %v\n", meuArray[len(meuArray)-1])
	fmt.Println("Meu Array 1")
	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d\n", i+1, v)
	}

	fmt.Println("--------------------------------------")
	// ARRAY 2

	meuArray2 := [3]string{"ola", "array", "dois"}
	fmt.Println("Meu Array 2")
	for i, v := range meuArray2 {
		fmt.Printf("O valor do índice %d é %v\n", i+1, v)
	}

	fmt.Println("--------------------------------------")
	// SLICES

	s := []int{10, 20, 30, 40, 50, 60, 70, 90, 101}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(s[6:]), cap(s[6:]), s[6:])
	fmt.Printf("len=%d cap=%d %v\n", len(s[4:6]), cap(s[4:6]), s[4:6])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4])+len(s[6:]), cap(append(s[:4], s[6:]...)), append(s[:4], s[6:]...))

	fmt.Println("--------------------------------------")
	// MAPS (DICT)

	salarios := map[string]int{"Gustavo": 3000, "Amanda": 2000, "Maria": 1500}
	salarios["Nilton"] = 6000

	for nome, salario := range salarios {
		fmt.Printf("O Salario de %s é %d\n", nome, salario)
	}

	fmt.Println("--------------------------------------")
	// FUNÇÕES VARIÁTICAS

	somatorio := sum(23, 43, 54, 12, 54)
	fmt.Println(somatorio)
	texto := join("ola", "mundo", "solitário")
	fmt.Println(texto)

	fmt.Println("--------------------------------------")
	// CLOSURES

	frase := func() string {
		return fmt.Sprintf("%s%d", join("O", "valor", "é"), sum(4, 6, 3)*2)
	}()
	fmt.Println(frase)

	fmt.Println("--------------------------------------")
	// PONTEIROS

	minhaVar1 := 10
	minhaVar2 := 20
	pointers(&minhaVar1, &minhaVar2)
	fmt.Println(minhaVar1)

	fmt.Println("--------------------------------------")
	// GENERICS

	m := map[string]int{"Gustavo": 1000, "Amanda": 2000}
	m2 := map[string]float64{"Gustavo": 1000.50, "Amanda": 2000.50}
	fmt.Println(soma(m))
	fmt.Println(soma(m2))

	fmt.Println(compara(5, 5))
	fmt.Println(compara("Gustavo", "gustavo"))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func join(palavras ...string) string {
	texto := ""
	for _, palavra := range palavras {
		texto += palavra + " "
	}
	return texto
}

func pointers(a, b *int) int {
	*a = 50
	return *a + *b
}

func soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}
