package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type user struct {
	Name string
	Email string
}

func main()  {

	templates = template.Must(template.ParseGlob("*.html"))	

	// creating a uri route
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := user{"Leo", "leo@leo.com"}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	// up server
	fmt.Println("Server up at :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}