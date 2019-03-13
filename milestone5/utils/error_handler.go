package utils

import (
	//"fmt"
	"strings"
	"encoding/json"
	"log"
	"net/http"
)

type OK interface {
	OK() error
}

func ReqDataDecode(r *http.Request, reqData OK) error {
	if err := json.NewDecoder(r.Body).Decode(reqData); err != nil {
		return err
	}
	return reqData.OK()
}

type ErrMissingField []string

func (errors ErrMissingField) Error() *string {
	if len(errors) != 0{
		errorsStr := "Fields: "+ strings.Join(errors[:], ", ") +" is required"
		return &errorsStr
	}
	return nil
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}