package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {

	// creating a uri route
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HI"))
		fmt.Println(r)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HI users"))
		fmt.Println(r)
	})

	// up server
	log.Fatal(http.ListenAndServe(":5000", nil))
}