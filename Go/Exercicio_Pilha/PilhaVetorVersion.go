package main

import "fmt"

type PilhaVetor struct {
	PilhaVetor [20]string
	topo       int
	max        int
}

func (p *PilhaVetor) init() {
	p.topo = 0
	p.max = 20
}

func (p *PilhaVetor) add(valor string) {
	if p.EstaCheia() {
		fmt.Println("PilhaVetor cheia")
		return
	} else {
		p.PilhaVetor[p.topo] = valor
		p.topo++
	}
}

func (p *PilhaVetor) EstaVazia() bool {
	return p.topo == 0
}

func (p *PilhaVetor) EstaCheia() bool {
	return p.topo > p.max-1
}

func (p *PilhaVetor) Remove() string {
	if !p.EstaVazia() {
		p.topo--
		item := p.PilhaVetor[p.topo]
		return item
	} else {
		fmt.Println("Lista Vazia")
	}
	return ""
}

func (p *PilhaVetor) Mostra() {
	for i := 0; i < p.topo; i++ {
		fmt.Println(p.PilhaVetor[i])
	}
}
