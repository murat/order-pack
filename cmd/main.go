package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"order-pack/internal/api"
	"order-pack/internal/database"
	"order-pack/internal/product"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.New("./test.db")
	if err != nil {
		log.Fatalf("could not open database, %v", err)
	}

	if err := db.AutoMigrate(&product.Product{}); err != nil {
		log.Fatalf("could not migrate database, %v", err)
	}

	productSvc := product.NewService(db)
	apiSrv := api.NewApi(productSvc)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", apiSrv.RootHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", apiSrv.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/products", apiSrv.GetProductsHandler).Methods(http.MethodGet)
	r.HandleFunc("/products", apiSrv.CreateProductHandler).Methods(http.MethodPost)
	r.HandleFunc("/products/{id:[0-9]+}", apiSrv.GetProductHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.LoggingHandler(os.Stdout, api.ContentTypeMiddleware(r)),
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log.Println("[info] Server is starting...")
	go func() {
		<-quit
		log.Println("[info] Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("[error] Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Println("[info] Server is ready. Listening on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[error] could not start server, %v", err)
	}

	<-done
	log.Println("[info] Server stopped")
}
