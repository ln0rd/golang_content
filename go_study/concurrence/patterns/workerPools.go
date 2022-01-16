package main

import "fmt"

func main()  {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	// go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	
	close(tasks)

	for i := 0; i < 45; i++ {
		results := <-results
		fmt.Println(results)
	}
}

// [<-chan] so recebe dados [chan<-] so envia dados
func worker(tasks <-chan int, results chan<- int)  {
	for task := range tasks {
		results <- fibonacci(task)
	}
}

func fibonacci(position int) int  {
	if position <= 1 {
		return position
	}

	return fibonacci(position - 2) + fibonacci(position - 1)
}