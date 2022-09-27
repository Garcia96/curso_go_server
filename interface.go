package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) mover() string {
	return "soy un perro"
}

func (pez) mover() string {
	return "soy un pez"
}

func (pajaro) mover() string {
	return "soy un pájaro"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	perro := perro{}
	moverAnimal(perro)
	pez := pez{}
	moverAnimal(pez)
	pajaro := pajaro{}
	moverAnimal(pajaro)

}
