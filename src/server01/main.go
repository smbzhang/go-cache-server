// Package mainmain provides ...
package main

import (
	"server01/cache"
	"server01/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}

