package main

import "fmt"

func generic(f interface{})  {
	fmt.Println(f)
}

func main()  {
	generic("string")
	generic(1)
	generic(true)


	maps := map[interface{}]interface{}{
		1: true,
		"dois": float32(100),
	}

	fmt.Println(maps)
}