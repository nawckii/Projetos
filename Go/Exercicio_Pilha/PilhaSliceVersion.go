package main

import "fmt"

type PilhaSlice struct {
	PilhaSlice []string
}

func (p *PilhaSlice) init() {
}

func (p *PilhaSlice) add(valor string) {
	p.PilhaSlice = append(p.PilhaSlice, valor)
}

func (p *PilhaSlice) EstaVazia() bool {
	return len(p.PilhaSlice) == 0
}

func (p *PilhaSlice) Remove() string {
	if p.EstaVazia() {
		return ""
	}
	i := len(p.PilhaSlice) - 1
	res := p.PilhaSlice[i]
	p.PilhaSlice = p.PilhaSlice[:i]
	return res
}

func (p *PilhaSlice) Mostra() {
	for _, itens := range p.PilhaSlice {
		fmt.Println(itens)
	}
}
