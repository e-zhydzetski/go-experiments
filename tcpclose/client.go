package tcpclose

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func ConnectAndSendMsg(addr string, msg string) (string, error) {
	tcpAdr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return "", fmt.Errorf("resolve tcp: %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAdr)
	if err != nil {
		return "", fmt.Errorf("dial tcp: %v", err)
	}
	defer conn.Close()

	readDone := make(chan struct{})
	answer := bytes.NewBuffer(nil)
	go func() {
		_, err := io.Copy(answer, conn)
		log.Println("readDone, error:", err)
		close(readDone) // signal the main goroutine
	}()

	_, _ = io.Copy(conn, strings.NewReader(msg))
	_ = conn.CloseWrite()
	<-readDone
	_ = conn.CloseRead()

	return answer.String(), nil
}
