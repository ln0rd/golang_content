package main

import (
	"fmt"
	"time"
)

func main()  {
	canal := multiplexar(write("Hi"), write("Hi 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canal1, canal2 <-chan string) <-chan string  {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2: 
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

func write(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprint("Valor recebido", text)
			time.Sleep(time.Microsecond * 5000)
		}
	}()

	return canal
}