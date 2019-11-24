// Package tcp provides ...
package tcp

import (
    "server03/cache"
	"net"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) Listhen() {
	server, err := net.Listen("tcp", ":12346")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go s.process(client)
	}
}
