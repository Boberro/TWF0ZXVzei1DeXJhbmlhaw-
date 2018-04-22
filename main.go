package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"log"
	"os"
	"os/signal"
	"context"
	"flag"
)

var doSetup = flag.Bool("setup", false, "Generate database with demo data")

func main() {
	flag.Parse()

	if *doSetup {
		setup()
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// region DB
	dbConnect()
	defer dbClose()
	// endregion DB

	// region Router
	router := chi.NewRouter()
	router.Use(middleware.URLFormat)

	router.Route("/api", func(router chi.Router) {
		router.Route("/objects", func(router chi.Router) {
			router.Get("/", ViewObjectKeysGet)

			router.Route("/{objectKey}", func(router chi.Router) {
				//router.Use(ObjectCtx)
				router.Get("/", ViewObjectGet)
				router.Put("/", ViewObjectPut)
				router.Delete("/", ViewObjectDelete)
			})
		})
	})
	// endregion Router

	// region http server
	server := &http.Server{Addr: ":8080", Handler: router}

	go func() {
		log.Println("Listening on port :8080")

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop

	log.Println("\nShutting down the server...")

	server.Shutdown(context.Background())
	// endregion http server
}
