package controller

import (
	model "coffee/internals/app/models"
	service "coffee/internals/app/service"
	"encoding/json"
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

func (controller *CoffeeController) FindById(response http.ResponseWriter, request *http.Request) {
	var coffee model.Coffee

	vars := mux.Vars(request)
	if vars["id"] == "" {
		WrapError(response, errors.New("id is empty"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	checkError(response, err)

	coffee, err = controller.service.FindById(id)
	if err != nil {
		log.Println("Parse error: ", err)
		WrapError(response, err)
		return
	}

	var message = map[string]interface{}{
		"status": "ok",
		"data":   coffee,
	}

	WrapOk(response, message)
}

func (controller *CoffeeController) AddNew(response http.ResponseWriter, request *http.Request) {
	var coffee model.Coffee

	err := json.NewDecoder(request.Body).Decode(&coffee)
	checkError(response, err)

	err = controller.service.AddNew(coffee)
	checkError(response, err)
}

func (controller *CoffeeController) GetAll(response http.ResponseWriter, request *http.Request) {
	coffeeList := controller.service.GetAll()
	var message = map[string]interface{}{
		"status": "ok",
		"data":   coffeeList,
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
