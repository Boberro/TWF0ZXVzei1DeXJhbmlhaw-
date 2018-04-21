package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"io/ioutil"
)

func ViewObjectKeysGet(w http.ResponseWriter, r *http.Request) {
	for _, obj := range myObjects {
		w.Write([]byte(obj.Key + ", "))
	}
}

func ViewObjectGet(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")
	if !validata_key(objectKey) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, obj := getObject(myObjects, objectKey)
	if obj != nil {
		w.Write([]byte(obj.Value))
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func ViewObjectPut(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")
	if !validata_key(objectKey) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request, _ := ioutil.ReadAll(r.Body)
	data := string(request)

	i, obj := getObject(myObjects, objectKey)
	if obj != nil {
		myObjects[i].Value = data
	} else {
		myObjects = append(myObjects, &MyObject{objectKey, data})
	}
}

func ViewObjectDelete(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")
	if !validata_key(objectKey) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i, obj := getObject(myObjects, objectKey)
	if obj != nil {
		myObjects = append(myObjects[:i], myObjects[i+1:]...)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
