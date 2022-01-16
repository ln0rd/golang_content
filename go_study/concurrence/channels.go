package main

import (
	"fmt"
	"time"
)

// better way to synchronize goroutines

func main()  {
	channel := make(chan string)
	
	go write("Hi", channel)

	for {	
		msg, isOpen := <-channel
		if !isOpen {
			break
		}
		println(msg)
	}

	fmt.Println("End")

}


// adiciono o texto no canal
func write(text string, channel chan string)  {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	// preciso fechar para não criar um deadlock
	close(channel)
}