package api

import (
	"coffee/internals/app/controller"

	"github.com/gorilla/mux"
)

func CreateRoutes(coffeeController *controller.CoffeeController) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/api/v1/coffee/{id:[0-9]+}", coffeeController.FindById).Methods("GET")
	route.HandleFunc("/api/v1/coffee/add-new", coffeeController.AddNew).Methods("POST")
	route.HandleFunc("/api/v1/coffee/get-all", coffeeController.GetAll).Methods("GET")
	route.NotFoundHandler = route.NewRoute().HandlerFunc(controller.NotFound).GetHandler()
	return route
}
