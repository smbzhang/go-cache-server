// Package mainmain provides ...
package main

import (
	"server03/cache"
	"server03/http"
	"server03/tcp"
	"flag"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listhen()
	http.New(c).Listen()
}

