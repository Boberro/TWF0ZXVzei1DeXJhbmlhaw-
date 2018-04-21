package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Route("/api", func(router chi.Router) {
		router.Route("/objects", func(router chi.Router) {
			router.Get("/", GetObjectKeys)

			router.Route("/{objectKey}", func(router chi.Router) {
				//router.Use(ArticleCtx)
				router.Get("/", GetObject)
				//router.Put("/", UpdateArticle)
				//router.Delete("/", DeleteArticle)
			})
		})
	})

	http.ListenAndServe(":8080", router)
}