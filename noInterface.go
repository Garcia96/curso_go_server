package main

import "fmt"

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) caminar() string {
	return "soy un perro"
}

func (pez) nadar() string {
	return "soy un pez"
}

func (pajaro) volar() string {
	return "soy un pájaro"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}
func moverPajaro(p pajaro) {
	fmt.Println(p.volar())
}

func moverPez(p pez) {
	fmt.Println(p.nadar())
}

func main() {
	perro := perro{}
	moverPerro(perro)
	pez := pez{}
	moverPez(pez)
	pajaro := pajaro{}
	moverPajaro(pajaro)

}
