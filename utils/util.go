package util

import "strings"

type JsonError struct {
	Error string `json:"error"`
}

func NewJSONrror(err error) JsonError {
	jerr := JsonError{"generic error"}
	if err != nil {
		jerr.Error = err.Error()
	}
	return jerr
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
