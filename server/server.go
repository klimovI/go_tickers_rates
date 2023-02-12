package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/klimovI/go_tickers_rates/server/handler"
	"github.com/klimovI/go_tickers_rates/server/service"
)

const port = "3001"

func StartServer() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        handlers.Init(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}

	fmt.Println("Listening at port " + port)
}
