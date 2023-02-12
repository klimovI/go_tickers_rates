package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/klimovI/go_tickers_rates/server/handler"
	"github.com/klimovI/go_tickers_rates/server/service"
)

const port = "3001"

func StartServer() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handlers.Init(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Error starting server: %s", err.Error())
		}
	}()

	fmt.Println("Listening at port " + port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error shutting down: %s", err.Error())
	}

	os.Exit(0)
}
