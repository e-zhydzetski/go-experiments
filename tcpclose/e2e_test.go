package tcpclose_test

import (
	"context"
	"fmt"
	"github.com/e-zhydzetski/go-experiments/tcpclose"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	port, err := tcpclose.StartServer(ctx, "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	const parallel = 10
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			answer, err := tcpclose.ConnectAndSendMsg(fmt.Sprintf("127.0.0.1:%d", port), "Test msg\n")
			if err != nil {
				t.Error(err)
			}
			expected := "\t TEST MSG\n\t Test msg\n\t test msg\n"
			if answer != expected {
				t.Errorf("invalid answer: expected '%s', got '%s'", expected, answer)
			}
		}()
	}

	wg.Wait()
}
