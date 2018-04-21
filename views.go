package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"io/ioutil"
	"strconv"
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
		w.Header().Set("Content-Type", obj.ContentType)
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

	contentSize, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil || contentSize > 100000000 {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	if len(r.Header.Get("Content-Type")) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Content-Type header required"))
		return
	}

	postData, _ := ioutil.ReadAll(r.Body)

	i, obj := getObject(myObjects, objectKey)
	if obj != nil {
		myObjects[i].Value = postData
		myObjects[i].ContentType = r.Header.Get("Content-Type")
	} else {
		myObjects = append(myObjects, &MyObject{objectKey, postData, r.Header.Get("Content-Type")})
	}
	w.Write([]byte(r.Header.Get("Content-Length")))
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
