package controller

import (
	"errors"
	"net/http"
)

func NotFound(response http.ResponseWriter, request *http.Request) {
	WrapErrorWithStatus(response, errors.New("Not found"), http.StatusNotFound)
}
