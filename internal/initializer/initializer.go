package initializer

import (
	"context"
	"fmt"
	"sync"

	"github.com/93gad/URLCounterCLI/internal/searches"
)

type SearchResult struct {
	URL   string
	Count int
}

func InitializeAndRun(ctx context.Context, urls []string, searchString string, concurrency int) (map[string]int, error) {
	resultMap := make(map[string]int)
	resultChannel := make(chan SearchResult)
	done := make(chan struct{}) // Канал для отправки сигнала о завершении работы

	wg := sync.WaitGroup{}

	// Fan-out
	go FanOut(ctx, &wg, urls, searchString, concurrency, resultChannel)

	// Fan-in
	go FanIn(&wg, resultChannel, resultMap, done)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		return resultMap, nil
	}
}

func FanOut(ctx context.Context, wg *sync.WaitGroup, urls []string, searchString string, concurrency int, resultChannel chan<- SearchResult) {

	sem := make(chan struct{}, concurrency) // Новый канал размером 5

	for _, url := range urls {
		wg.Add(1)         // Счетчик для горутины
		sem <- struct{}{} // Запись в канал

		go func(url, searchString string) {
			defer wg.Done()          // Горутина завершилась
			defer func() { <-sem }() // Разрешение для новой горутины

			count, err := searches.CountOccurrences(ctx, url, searchString)
			if err != nil {
				fmt.Printf("Ошибка при выполнении запроса к %s: %v\n", url, err)
				return
			}
			resultChannel <- SearchResult{URL: url, Count: count} // Результат отправляется в канал
		}(url, searchString)
	}

	wg.Wait() // Блокирует выполнение горутины, пока счетчик wg не будет нулем
	close(resultChannel)
}

func FanIn(wg *sync.WaitGroup, resultChannel <-chan SearchResult, resultMap map[string]int, done chan<- struct{}) {
	defer close(done)

	for result := range resultChannel {
		resultMap[result.URL] = result.Count
	}

	wg.Wait()
}
