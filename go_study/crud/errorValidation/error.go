package errorValidation

import (
	"log"
	"net/http"
)

func ErrorValidation(err error, message string)  {
	if err != nil {
		log.Fatal(message)
		return
	}
}

func ErrorResponseValidation(err error,  w http.ResponseWriter, message string)  {
	if err != nil {
		log.Fatal(message)
		w.Write([]byte(message))
		return 
	}
}

