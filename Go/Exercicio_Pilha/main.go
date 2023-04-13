package main

import (
	"fmt"
	"os"
)

func CheckExpressionVector(Expression string) bool {
	var P PilhaVetor
	P.init()
	for _, i := range Expression {
		char := string(i)
		fmt.Print(char)
		if char == "(" {
			if P.EstaCheia() {
				fmt.Println("\nerror: Pilha cheia")
				os.Exit(1)
			}
			P.add(char)
		}
		if char == ")" {
			if P.EstaVazia() {
				return false
			}
			P.Remove()
		}
	}
	if P.EstaVazia() {
		return true
	} else {
		return false
	}
}

func CheckExpressionSlice(Expression string) bool {
	var P PilhaSlice
	P.init()
	for _, i := range Expression {
		char := string(i)
		fmt.Print(char)
		if char == "(" {
			P.add(char)
		}
		if char == ")" {
			if P.EstaVazia() {
				return false
			}
			P.Remove()
		}
	}
	if P.EstaVazia() {
		return true
	} else {
		return false
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args)
		fmt.Println("Erro: nenhum argumento fornecido")
		return
	}
	Expression := os.Args[1]
	ce := CheckExpressionVector(Expression)
	fmt.Print("\nResultado da expressão usando vetor: ")
	if ce {
		fmt.Println("correta")
	} else {
		fmt.Println("incorreta")
	}

	ce = CheckExpressionSlice(Expression)
	fmt.Print("\nResultado da expressão usando Slice: ")
	if ce {
		fmt.Println("correta")
	} else {
		fmt.Println("incorreta")
	}

}
