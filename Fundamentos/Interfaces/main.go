package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Println("Area da forma:", f.area())
}

type forma interface {
	area() float64
}

func main() {
	fmt.Println("Interfaces")

	r := retangulo{10, 15}
	c := circulo{10}

	escreverArea(c)
	escreverArea(r)
}
