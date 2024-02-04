package tests

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/93gad/URLCounterCLI/internal/initializer"
)

func TestFanOut(t *testing.T) {

	urls := []string{"https://go.dev/tos", "https://go.dev/copyright"}
	searchString := "go"
	concurrency := 2
	resultChannel := make(chan initializer.SearchResult, len(urls))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}

	initializer.FanOut(ctx, wg, urls, searchString, concurrency, resultChannel)

	timeout := time.After(5 * time.Second)

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	results := make(map[string]int)
	for {
		select {
		case <-timeout:
			t.Fatal("Время ожидания теста истекло")
		case result, ok := <-resultChannel:
			if !ok {

				resultChannel = nil
			} else {
				results[result.URL] = result.Count
			}
		}
		if resultChannel == nil {
			break
		}
	}

	if len(results) != len(urls) {
		t.Errorf("Ожидаемые результаты %d, получены %d", len(urls), len(results))
	}

	for url, count := range results {
		if count < 0 {
			t.Errorf("Недопустимое количество для URL-адреса %s: %d", url, count)
		}
	}
}
