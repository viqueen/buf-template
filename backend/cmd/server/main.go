package main

import (
	"net/http"
	"time"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"connectrpc.com/otelconnect"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/viqueen/buf-template/api/go-sdk/todo/v1/todoV1connect"
	apitodov1 "github.com/viqueen/buf-template/backend/internal/api-todo-v1"
	"github.com/viqueen/buf-template/backend/internal/logger"
	"github.com/viqueen/buf-template/backend/internal/store"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	logger.Init()

	otelInterceptor, otelErr := otelconnect.NewInterceptor()
	if otelErr != nil {
		log.Fatal().Err(otelErr).Msg("failed to initialise otel interceptor")
	}

	db, dbErr := store.InitStore()
	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("failed to initialise db")
	}

	todoRepo := store.NewTodoRepository(db)

	todoService := apitodov1.NewTodoService(todoRepo)
	todoPath, todoHandler := todoV1connect.NewTodoServiceHandler(
		todoService,
		connect.WithInterceptors(otelInterceptor),
	)

	mux := http.NewServeMux()
	mux.Handle(todoPath, todoHandler)

	h2cMux := h2c.NewHandler(mux, &http2.Server{})

	log.Info().Msg("starting server on :8080")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      withCORS(h2cMux),
		ReadTimeout:  5 * time.Second,   //nolint: mnd
		WriteTimeout: 10 * time.Second,  //nolint: mnd
		IdleTimeout:  120 * time.Second, //nolint: mnd
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}

func withCORS(handler http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})

	return middleware.Handler(handler)
}
