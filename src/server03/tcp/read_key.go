// Package tcp provides ...
package tcp

import (
	"bufio"
	"io"
)

func (s *Server) readKey(reader *bufio.Reader) (string, error) {
	klen, err := readLen(reader)
	if err != nil {
		return "", err
	}
	k := make([]byte, klen)
	_, err = io.ReadFull(reader, k)
	if err != nil {
		return "", err
	}
    return string(k), err
}

func (s *Server) readkeyAndValue(reader *bufio.Reader) ( string, []byte, error) {
	klen, err := readLen(reader)
	if err != nil {
		return "", nil, err
	}
	vlen, err := readLen(reader)
	if err != nil {
		return "", nil, err
	}
	k := make([]byte, klen)
	_, err = io.ReadFull(reader, k)
	if err != nil {
		return "", nil, err
	}
	v := make([]byte, vlen)
	_, err = io.ReadFull(reader, v)
	if err != nil {
		return "", nil, err
	}
	return string(k), v, err
}



