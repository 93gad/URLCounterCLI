package main

import (
	"context"
	"fmt"
	"time"

	"github.com/93gad/URLCounterCLI/internal/flags"
	"github.com/93gad/URLCounterCLI/internal/initializer"
)

func main() {

	urls, searchString := flags.ParseFlags()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	concurrency := 5

	resultMap, err := initializer.InitializeAndRun(ctx, urls, searchString, concurrency)
	if err != nil {
		fmt.Printf("Ошибка при выполнении приложения: %v\n", err)
		return
	}

	for url, count := range resultMap {
		fmt.Printf("%s - %d\n", url, count)
	}
}
