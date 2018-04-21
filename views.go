package main

import (
	"net/http"
	"github.com/go-chi/chi"
)

func GetObjectKeys(w http.ResponseWriter, r *http.Request) {
	for _, obj := range myObjects {
		w.Write([]byte(obj.Key))
	}
}

func GetObject(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")

	for _, obj := range myObjects {
		if obj.Key == objectKey {
			w.Write([]byte(obj.Value))
		}
	}
}
