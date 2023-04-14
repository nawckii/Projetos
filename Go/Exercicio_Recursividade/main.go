package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Definir opções de linha de comando
	n := flag.Int("n", 0, "the number to calculate the Fibonacci sequence for")
	t := flag.String("t", "o", "tipo de calculo: (F)ibonacci, (f)atorial ou (o)utro")

	// Configurar a mensagem de ajuda
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	// Analisar a linha de comando
	flag.Parse()

	// Verificar se a flag foi fornecida
	if *n == 0 {
		flag.Usage()
		os.Exit(1)
	} else {
		if *t == "F" {
			// Calcular a sequência de Fibonacci
			fibonaccivar := fibonacci(*n)
			fmt.Println(fibonaccivar)
		} else if *t == "f" {
			// Calcular o fatorial
			fatorial := Fatorial(*n)
			fmt.Println(fatorial)
		} else {
			fmt.Println("Digite uma lista:")
			
		}

	}

}
