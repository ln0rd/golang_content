// encapsula uma go routine e retorna um canal

package main

import (
	"fmt"
	"time"
)

func main()  {
	canal := write("Hello World")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal )
	}
}

func write(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprint("Valor recebido", text)
			time.Sleep(time.Microsecond * 500)
		}
	}()

	return canal
}