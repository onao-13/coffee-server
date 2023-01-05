package data

import (
	model "coffee/internals/app/models"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CoffeeStorage struct {
	pool *pgxpool.Pool
}

func NewCoffeeStorage(pool *pgxpool.Pool) *CoffeeStorage {
	storage := new(CoffeeStorage)
	storage.pool = pool
	return storage
}

func (storage *CoffeeStorage) FindById(id int64) (model.Coffee, error) {
	var coffee model.Coffee
	query := "SELECT * FROM coffee WHERE id := $1"

	err := pgxscan.Select(context.Background(), storage.pool, &coffee, query, id)

	if err != nil {
		log.Println("No coffees in database")
	}

	log.Print("Values: ", coffee.Name, coffee.Price)
	return coffee, err
}

func (storage *CoffeeStorage) AddNew(coffee model.Coffee) {
	query := "INSERT INTO coffee (name, price) VALUES($1, $2)"
	storage.pool.QueryRow(context.Background(), query, coffee.Name, coffee.Price)
	log.Print("Values: ", coffee.Name, coffee.Price)
}

func (storage *CoffeeStorage) GetAll() []model.Coffee {
	var coffeeList []model.Coffee

	query := "SELECT * FROM coffee"

	err := pgxscan.Select(context.Background(), storage.pool, &coffeeList, query)
	if err != nil {
		log.Println(err)
	}

	return coffeeList
}
