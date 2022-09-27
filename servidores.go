package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	final := time.Since(inicio)
	fmt.Println(final)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, ": No esta disponible")
	} else {
		fmt.Println("Online")
	}
}
