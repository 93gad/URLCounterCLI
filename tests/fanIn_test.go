package tests

import (
	"sync"
	"testing"
	"time"

	"github.com/93gad/URLCounterCLI/internal/initializer"
)

func TestFanIn(t *testing.T) {

	resultChannel := make(chan initializer.SearchResult, 2)
	done := make(chan struct{})
	resultMap := make(map[string]int)
	wg := &sync.WaitGroup{}

	// Имитирует отправку результатов в resultChannel
	go func() {
		resultChannel <- initializer.SearchResult{URL: "https://go.dev/tos", Count: 101}
		resultChannel <- initializer.SearchResult{URL: "https://go.dev/copyright", Count: 100}
		close(resultChannel)
	}()

	initializer.FanIn(wg, resultChannel, resultMap, done)

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Error("FanIn не завершил работу в течение периода ожидания")
	}

	expectedCount := 2
	if len(resultMap) != expectedCount {
		t.Errorf("Expected %d results, got %d", expectedCount, len(resultMap))
	}

	if resultMap["https://go.dev/tos"] != 101 {
		t.Errorf("Неправильное количество для URL-адреса https://go.dev/tos: %d", resultMap["https://go.dev/tos"])
	}
	if resultMap["https://go.dev/copyright"] != 100 {
		t.Errorf("Неправильное количество для URL-адреса https://go.dev/copyright: %d", resultMap["https://go.dev/copyright"])
	}
}
