package data

import (
	model "coffee/internals/app/models"
)

type CoffeeStorage struct {
	// pool
}

func NewCoffeeStorage() *CoffeeStorage {
	storage := new(CoffeeStorage)
	return storage
}

func (data CoffeeStorage) FindCoffeeById(id int64) model.Coffee {
	coffeesList := [4]model.Coffee{
		model.New(1, "капучино", 3.25),
		model.New(2, "латте", 2.80),
		model.New(3, "эспрессо", 2.10),
		model.New(4, "колд брю", 1.73),
	}

	var searchedCoffee model.Coffee

	for _, coffee := range coffeesList {
		if coffee.Id == id {
			searchedCoffee = coffee
			break
		}
	}

	return searchedCoffee
}
