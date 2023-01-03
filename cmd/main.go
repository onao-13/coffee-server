package main

import (
	serv "coffee/internals/app"
	cfg "coffee/internals/config"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	config := cfg.UploadDevConfig()
	ctx, cancel := context.WithCancel(context.Background())

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)

	server := serv.NewServer(ctx, config)

	go func() {
		osCall := <-s
		log.Println("Server stoped: ", osCall)
		server.Shutdown()
		cancel()
	}()

	server.Connected()
}
