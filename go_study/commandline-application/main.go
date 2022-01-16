package main

import (
	"commandline/app"
	"log"
	"os"
)

func main()  {
	application := app.Generator()
	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}