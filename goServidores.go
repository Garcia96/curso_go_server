package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	canal := make(chan string)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	i := 0
	for {
		if i > 2 {
			break
		}
		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	//for i := 0; i < len(servidores); i++ {
	//	fmt.Println(<-canal)
	//}

	final := time.Since(inicio)
	fmt.Println(final)
}

func revisarServidor(servidor string, c chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, ": No esta disponible")
		c <- servidor + ": No esta disponible"
	} else {
		fmt.Println("Online")
		c <- servidor + " Online"
	}
}
