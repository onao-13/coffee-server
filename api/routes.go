package api

import (
	"coffee/internals/app/controller"

	"github.com/gorilla/mux"
)

func CreateRoutes(coffeeController *controller.CoffeeController) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/coffee/{id:[0-9]+}", coffeeController.FindCoffeeById).Methods("GET")
	route.NotFoundHandler = route.NewRoute().HandlerFunc(controller.NotFound).GetHandler()
	return route
}
