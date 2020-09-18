package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/config"
	"github.com/elauso/go-rest-api/src/net/elau/gorestapi/model"
)

func main() {

	model.InitDB(os.Getenv("DATASOURCE_URL"))

	r := mux.NewRouter()

	pr := config.CreateProductRoute()
	r.HandleFunc("/products", pr.List).Methods("GET")
	r.HandleFunc("/products/{productId}", pr.Get).Methods("GET")
	r.HandleFunc("/products", pr.Create).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logFile := os.Getenv("LOG_FILE_LOCATION")
	if logFile != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    500,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		})
	}

	go func() {
		log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down...")
	os.Exit(0)
}
