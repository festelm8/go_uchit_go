package utils

import (
	//"fmt"
	"strings"
	"encoding/json"
	"net/http"
)

type OK interface {
	OK() error
}

func ParseReqData(r *http.Request, reqData OK) error {
	if err := json.NewDecoder(r.Body).Decode(reqData); err != nil {
		return err
	}
	return reqData.OK()
}

type ErrMissingField struct{
	Fields []string
}

func (errors ErrMissingField) Error() string {
	return "Fields: "+ strings.Join(errors.Fields[:], ", ") +" is required"
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}