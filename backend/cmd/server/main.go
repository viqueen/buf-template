package main

import (
	"api/go-sdk/todo/v1/todoV1connect"
	apitodov1 "backend/internal/api-todo-v1"
	"backend/internal/store"
	"backend/internal/store/gendb"
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"connectrpc.com/otelconnect"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	otelInterceptor, otelErr := otelconnect.NewInterceptor()
	if otelErr != nil {
		log.Fatalf("failed to initialise otel interceptor: %v", otelErr)
	}

	db, dbErr := store.InitStore()
	if dbErr != nil {
		log.Fatalf("failed to initialise db: %v", dbErr)
	}
	dataStore := gendb.New(db)

	todoService := apitodov1.NewTodoService(dataStore)
	todoPath, todoHandler := todoV1connect.NewTodoServiceHandler(todoService, connect.WithInterceptors(otelInterceptor))

	mux := http.NewServeMux()
	mux.Handle(todoPath, todoHandler)

	h2cMux := h2c.NewHandler(mux, &http2.Server{})
	log.Printf("starting server on :8080")
	err := http.ListenAndServe(":8080", withCORS(h2cMux))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
