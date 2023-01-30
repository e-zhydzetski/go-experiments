package tcpclose_test

import (
	"context"
	"fmt"
	"github.com/e-zhydzetski/go-experiments/tcpclose"
	"testing"
)

func Test(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	port, err := tcpclose.StartServer(ctx, "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	err = tcpclose.ConnectAndSendMsg(fmt.Sprintf("127.0.0.1:%d", port), "Test msg\n")
	if err != nil {
		t.Fatal(err)
	}
}
