package main

import (
	"fmt"
	"introduction-test/address"
)

func main()  {
	kindAddress := address.KindOfAddress("avenida paulista")
	fmt.Println(kindAddress)
}