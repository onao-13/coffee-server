package controller

import (
	"encoding/json"
	"net/http"
)

func WrapError(response http.ResponseWriter, err error) {
	WrapErrorWithStatus(response, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(response http.ResponseWriter, err error, httpStatus int) {
	var message = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, _ := json.Marshal(message)
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("X-Content-Type-Options", "nosniff")
	response.WriteHeader(httpStatus)
	response.Write(res)
}

func WrapOk(response http.ResponseWriter, message map[string]interface{}) {
	res, _ := json.Marshal(message)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
