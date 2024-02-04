package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/93gad/URLCounterCLI/internal/searches"
)

func TestCountOccurrences(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("go go golang"))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := searches.CountOccurrences(ctx, server.URL, "go")
	if err != nil {
		t.Errorf("Неожиданная ошибка: %v", err)
	}

	expectedCount := 3
	if count != expectedCount {
		t.Errorf("Ожидаемое значение должно быть %d, получено %d", expectedCount, count)
	}
}

func TestCountOccurrences_HTTPError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := searches.CountOccurrences(ctx, "http://invalid-url", "go")

	if err == nil {
		t.Error("Ожидал ошибку, получил ноль")
	}
}
