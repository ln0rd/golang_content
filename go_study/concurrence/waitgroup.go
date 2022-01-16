package main

import (
	"fmt"
	"sync"
	"time"
)

// O que acontece aqui é que criando um waitGroup eu garanto que minhas duas rotinas vão acabar no mesmo tempo, então elas
// executatam seus steps um apos o outro.

// uma maneira de sincronizar as GoRoutines

func main()  {
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)

	go func ()  {
		write("competing [1]")
		waitGroup.Done()
	}()

	go func ()  {
		write("competing [2]")	
	}()

	go func ()  {
		write("competing [2]")	
	}()

	go func ()  {
		write("competing [2]")	
	}()

	waitGroup.Wait()
}

func write(text string)  {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}