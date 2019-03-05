package utils

import (
	"log"
)

type ErrorResponse struct {
  Msg string `json:"msg"`
}

// checkError check errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}