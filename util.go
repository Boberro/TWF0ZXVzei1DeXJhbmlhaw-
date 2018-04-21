package main

import "regexp"

func validata_key(key string) bool {
	if len(key) > 100 {
		return false
	}

	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", key)
	if !match {
		return false
	}

	return true
}
