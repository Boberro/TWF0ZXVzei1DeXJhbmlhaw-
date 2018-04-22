package main

import (
	"github.com/boltdb/bolt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

func setup() {
	dbConnect()
	defer dbClose()

	var getGopher = func() []byte {
		response, err := http.Get("https://golang.org/doc/gopher/frontpage.png")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			return contents
		}
		return nil
	}

	type MyObject struct {
		Key         string
		Value       []uint8
		ContentType []uint8
	}

	var myObjects = []*MyObject{
		{"demo1", []uint8("Ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn"), []uint8("text/plain")},
		{"demo2", []uint8("{\"glossary\":{\"title\":\"example glossary\",\"GlossDiv\":{\"title\":\"S\",\"GlossList\":{\"GlossEntry\":{\"ID\":\"SGML\",\"SortAs\":\"SGML\",\"GlossTerm\":\"Standard Generalized Markup Language\",\"Acronym\":\"SGML\",\"Abbrev\":\"ISO 8879:1986\",\"GlossDef\":{\"para\":\"A meta-markup language, used to create markup languages such as DocBook.\",\"GlossSeeAlso\":[\"GML\",\"XML\"]},\"GlossSee\":\"markup\"}}}}}"), []uint8("application/json")},
		{"demo3", getGopher(), []uint8("image/png")},
	}

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("objects"))
		if err != nil {
			return err
		}

		for _, obj := range myObjects {
			contentTypeKey := obj.Key + "-content_type"

			err = b.Put([]byte(obj.Key), obj.Value)
			if err != nil {
				return err
			}

			err = b.Put([]byte(contentTypeKey), obj.ContentType)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
