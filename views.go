package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"io/ioutil"
)

func ViewObjectKeysGet(w http.ResponseWriter, r *http.Request) {
	for _, obj := range myObjects {
		w.Write([]byte(obj.Key))
	}
}

func ViewObjectGet(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")

	_, obj := getObject(myObjects, objectKey)
	if obj != nil {
		w.Write([]byte(obj.Value))
	}
}

func ViewObjectPut(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")
	request, _ := ioutil.ReadAll(r.Body)
	data := string(request)

	i, obj := getObject(myObjects, objectKey)
	if obj != nil {
		myObjects[i].Value = data
	} else {
		myObjects = append(myObjects, &MyObject{objectKey, data})
	}
}
