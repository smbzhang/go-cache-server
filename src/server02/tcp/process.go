// Package tcp provides ...
package tcp

import (
	"net"
	"bufio"
	"io"
	"log"
)

func (s *Server) process(client net.Conn) {
	defer client.Close()
	// 在 net.Conn 上套了一层 bufio.Reader ，直接读取 conn，很可能客户端传来的数据不全，bufio会阻塞等待符合的数据。这里我感觉是处理tcp的粘包问题
	reader := bufio.NewReader(client)
	for {
		op, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				// 控制客户端是短链接
				log.Println("close client connection due to error: ", err)
				return
			}
		}
		if op == 'S' {
			err = s.set(client, reader)
		}else if op == 'G' {
			err = s.get(client, reader)
		}else if op == 'D' {
			err = s.del(client, reader)
		}else {
			log.Println("close client connection due to error: ", err)
			return
		}
		if err != nil {
			log.Println("close client connection due to error: ", err)
			return
		}
	}
}

func (s *Server) set(client net.Conn, reader *bufio.Reader) error {
	k, v, err := s.readkeyAndValue(reader)
	if err != nil {
		return err
	}
	err = s.Set(k, v)
	err = sendResponse(nil, err, client)
	return err
}

func (s *Server) get(client net.Conn, reader *bufio.Reader) error {
	k, err := s.readKey(reader)
	if err != nil {
		return err
	}
	v, err := s.Get(k)
	err = sendResponse(v, err, client)
	return err
}

func (s *Server) del(client net.Conn, reader *bufio.Reader) error {
	k, err := s.readKey(reader)
	if err != nil {
		return err
	}
	err = s.Del(k)
	err = sendResponse(nil, err, client)
	return err
}

