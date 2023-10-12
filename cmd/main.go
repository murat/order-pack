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
	"order-pack/internal/pack"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := database.New("test.db")
	if err != nil {
		log.Fatalf("could not open database, %v", err)
	}

	if err := db.Conn.AutoMigrate(&pack.Pack{}); err != nil {
		log.Fatalf("could not migrate database, %v", err)
	}

	packSvc := pack.NewService(db)
	apiSrv := api.NewApi(packSvc)

	r := mux.NewRouter()
	r.HandleFunc("/", apiSrv.RootHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", apiSrv.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/packs", apiSrv.GetPackagesHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.LoggingHandler(os.Stdout, r),
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
