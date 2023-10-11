package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// В Go нельзя принудительно завершить горутину.
// Однако можно сделать горутину останавливаемой,
// прослушивая значение-сигнал через канал.
// Пакет context может использоваться для этой цели.
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("Goroutine started")
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				defer fmt.Println("Goroutine is completed")
				return
			default:
				// do some work
			}
		}
	}(ctx)
	fmt.Println("Let's wait a small amount of time")
	time.Sleep(3 * time.Second)
	fmt.Println("Send ending signal")
	cancel()
	wg.Wait()
}
