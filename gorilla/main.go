package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webservices/gorilla/handlers"
)

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	router := mux.NewRouter()

	productRouter := router.PathPrefix("/products").Subrouter()

	getRouter := productRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("", ph.GetProducts)
	getRouter.HandleFunc("/{id}", ph.GetProduct)

	putRouter := productRouter.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := productRouter.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	server := &http.Server{
		Addr:    ":8099",
		Handler: router,
	}

	go func() {
		log.Println("Server started on port 8080")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server shut down gracefully")
}
