package api

import (
	"coffee/internals/app/controller"

	"github.com/gorilla/mux"
)

func CreateRoutes(coffeeController *controller.CoffeeController) *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/coffee/{id:[0-9]+}", coffeeController.FindById).Methods("GET")
	route.HandleFunc("/coffee/add-new", coffeeController.AddNew).Methods("POST")
	route.HandleFunc("/coffee/get-all", coffeeController.GetAll).Methods("GET")
	route.NotFoundHandler = route.NewRoute().HandlerFunc(controller.NotFound).GetHandler()
	return route
}
