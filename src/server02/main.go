// Package mainmain provides ...
package main

import (
	"server02/cache"
	"server02/http"
	"server02/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listhen()
	http.New(c).Listen()
}

