package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"io/ioutil"
	"strconv"
	"github.com/boltdb/bolt"
	"fmt"
	"log"
	"strings"
)

func ViewObjectKeysGet(w http.ResponseWriter, _ *http.Request) {
	keys := make([]string, 0)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("objects"))
		if b == nil {
			return nil
		}

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			if !strings.HasSuffix(string(k), "-content_type") { // <- very ugly, I know
				keys = append(keys, "\""+string(k)+"\"")
			}
		}
		return nil
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	w.Write([]byte(strings.Join(keys, ", ")))
	w.Write([]byte("]"))
}

func ViewObjectGet(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")
	if !validateKey(objectKey) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := db.View(func(tx *bolt.Tx) error {
		contentTypeKey := objectKey + "-content_type"

		b := tx.Bucket([]byte("objects"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		value := b.Get([]byte(objectKey))
		if len(value) == 0 {
			return fmt.Errorf("key not found")
		}

		contentType := b.Get([]byte(contentTypeKey))
		w.Header().Set("Content-Type", string(contentType))
		w.Write([]byte(value))
		return nil
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func ViewObjectPut(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")
	if !validateKey(objectKey) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentSize, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil || contentSize > 100000000 {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}
	if contentSize == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(r.Header.Get("Content-Type")) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Content-Type header required"))
		return
	}

	postData, _ := ioutil.ReadAll(r.Body)

	err = db.Update(func(tx *bolt.Tx) error {
		contentTypeKey := objectKey + "-content_type"

		b, err := tx.CreateBucketIfNotExists([]byte("objects"))
		if err != nil {
			return err
		}

		err = b.Put([]byte(objectKey), []byte(postData))
		if err != nil {
			return err
		}

		return b.Put([]byte(contentTypeKey), []byte(r.Header.Get("Content-Type")))
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func ViewObjectDelete(w http.ResponseWriter, r *http.Request) {
	objectKey := chi.URLParam(r, "objectKey")
	if !validateKey(objectKey) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := db.Update(func(tx *bolt.Tx) error {
		contentTypeKey := objectKey + "-content_type"

		b := tx.Bucket([]byte("objects"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}

		value := b.Get([]byte(objectKey))
		if len(value) == 0 {
			return fmt.Errorf("key not found")
		}

		err := b.Delete([]byte(objectKey))
		if err == nil {
			err = b.Delete([]byte(contentTypeKey))
		}
		return err
	})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
