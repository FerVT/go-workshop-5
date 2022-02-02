package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	withCancel()
}

func withCancel() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	sleepAndTalk(ctx, 5*time.Second, "hello with cancel")
}

func withDuration() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	sleepAndTalk(ctx, 5*time.Second, "hello with timeout")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
