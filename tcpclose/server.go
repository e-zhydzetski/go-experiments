package tcpclose

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func StartServer(ctx context.Context, listenAddr string) (int, error) {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return 0, fmt.Errorf("listen: %v", err)
	}
	go func() {
		<-ctx.Done()
		_ = l.Close()
	}()
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Print(err)
				break
			}
			go handleConn(conn)
		}
	}()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func echo(c net.Conn, shout string, delay time.Duration) error {
	_, err := fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	if err != nil {
		return fmt.Errorf("send 1: %v", err)
	}
	time.Sleep(delay)
	_, err = fmt.Fprintln(c, "\t", shout)
	if err != nil {
		return fmt.Errorf("send 2: %v", err)
	}
	time.Sleep(delay)
	_, err = fmt.Fprintln(c, "\t", strings.ToLower(shout))
	if err != nil {
		return fmt.Errorf("send 3: %v", err)
	}
	return nil
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	if input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := echo(c, input.Text(), 1*time.Second)
			if err != nil {
				log.Println("echo error:", err)
			}
		}()
	}
	if err := input.Err(); err != nil {
		log.Println("scanner error:", err)
	}
	wg.Wait()
	defer c.Close()
}
