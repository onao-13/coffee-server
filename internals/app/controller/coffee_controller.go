package controller

import (
	model "coffee/internals/app/models"
	service "coffee/internals/app/service"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CoffeeController struct {
	service *service.CoffeeService
}

func NewCoffeeController(service *service.CoffeeService) *CoffeeController {
	controller := new(CoffeeController)
	controller.service = service
	return controller
}

func (controller *CoffeeController) FindCoffeeById(response http.ResponseWriter, request *http.Request) {
	var coffee model.Coffee

	vars := mux.Vars(request)
	if vars["id"] == "" {
		WrapError(response, errors.New("id is empty"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	checkError(response, err)

	coffee, err = controller.service.FindCoffeeById(id)
	checkError(response, err)

	var message = map[string]interface{}{
		"status": "ok",
		"data":   coffee,
	}

	WrapOk(response, message)
}

func checkError(response http.ResponseWriter, err error) {
	if err != nil {
		log.Println("Parse error: ", err)
		WrapError(response, err)
		return
	}
}
