package app

import (
	api "coffee/api"
	"coffee/internals/app/controller"
	data "coffee/internals/app/data"
	service "coffee/internals/app/service"
	config "coffee/internals/config"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"context"
	"log"
	"net/http"
)

type Server struct {
	cfg  config.Config
	ctx  context.Context
	srv  *http.Server
	pool *pgxpool.Pool
}

func NewServer(ctx context.Context, cfg config.Config) *Server {
	server := new(Server)
	server.ctx = ctx
	server.cfg = cfg
	return server
}

func (server *Server) Connected() {
	log.Println("Server starting")

	var err error
	server.pool, err = pgxpool.Connect(server.ctx, server.cfg.DbUrlConnection())

	if err != nil {
		log.Fatal("Database connection is failed")
	}

	defer server.pool.Close()

	coffeeStorage := data.NewCoffeeStorage(server.pool)

	coffeeService := service.NewCoffeeService(coffeeStorage)

	coffeeController := controller.NewCoffeeController(coffeeService)

	routes := api.CreateRoutes(coffeeController)

	server.srv = &http.Server{
		Addr:    ":" + server.cfg.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (server Server) Shutdown() {
	log.Printf("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server shutdown failed: ", err)
	}

	if err == http.ErrServerClosed {
		err = nil
	}
}
