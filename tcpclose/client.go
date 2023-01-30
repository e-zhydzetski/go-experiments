package tcpclose

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func ConnectAndSendMsg(addr string, msg string) error {
	tcpAdr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return fmt.Errorf("resolve tcp: %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAdr)
	if err != nil {
		return fmt.Errorf("dial tcp: %v", err)
	}
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		log.Println("done", err)
		close(done) // signal the main goroutine
	}()
	_, _ = io.Copy(conn, strings.NewReader(msg))
	err = conn.CloseWrite()
	if err != nil {
		return fmt.Errorf("close write: %v", err)
	}
	<-done // wait for background goroutine to finish
	err = conn.CloseRead()
	if err != nil {
		return fmt.Errorf("close read: %v", err)
	}
	return nil
}
