package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"week2/internal/model"
	"week2/internal/routes"
)

var port = flag.String("p", "8080", "http port")

const (
	ReadTimeout  = 60 * time.Second
	WriteTimeout = 60 * time.Second
)

func init() {
	model.NewDBEngine()
}

func main() {
	flag.Parse()
	router := routes.NewRouter()
	s := &http.Server{
		Addr:           ":" + (*port),
		Handler:        router,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
