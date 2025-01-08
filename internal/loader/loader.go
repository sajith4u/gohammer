package loader

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func RunLoadTest(url string, concurrency, requests int) {
	var wg sync.WaitGroup
	start := time.Now()
	success := 0
	failed := 0

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requests/concurrency; j++ {
				resp, err := http.Get(url)
				if err != nil || resp.StatusCode >= 400 {
					failed++
					continue
				}
				success++
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Total: %d, Success: %d, Failed: %d, Time: %v\n", requests, success, failed, duration)
}
