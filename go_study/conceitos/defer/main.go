package main

import "fmt"

func main() {
	// adia a executcao até o ultimo momento possivel
	defer fun1()
	fun2()

}

func fun1() {
	fmt.Println("executing fun1")
}

func fun2() {
	fmt.Println("executing fun2")
}
