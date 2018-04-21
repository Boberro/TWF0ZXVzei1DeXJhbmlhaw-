package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	//router.Use(render.SetContentType(render.ContentTypeJSON))

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

	http.ListenAndServe(":8080", router)
}