// Package tcp provides ...
package tcp

import (
	"net"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)


func readLen(reader *bufio.Reader) (int, error) {
	tmp, e := reader.ReadString(' ')
	if e != nil {
		return 0, e
	}
	length, e := strconv.Atoi(strings.TrimSpace(tmp))
	if e != nil {
		return 0, e
	}
	return length, nil
}

func sendResponse(value []byte, err error, client net.Conn) error {
	if err != nil {
		errString := err.Error()
		res := fmt.Sprintf("-%d %s", len(errString), errString)
		_, e := client.Write([]byte(res))
		return e
	}
	vlen := fmt.Sprintf("%d ", len(value))
	_, e := client.Write(append([]byte(vlen), value...))
	return e
}
