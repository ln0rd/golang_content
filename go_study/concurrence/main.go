package main

import (
	"fmt"
	"time"
)

func main()  {
	go write("competing [1]") // goroutine
	write("competing [2]")
}

func write(text string)  {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}