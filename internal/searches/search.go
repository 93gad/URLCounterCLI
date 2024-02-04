package searches

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func CountOccurrences(ctx context.Context, url, searchString string) (int, error) {
	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return 0, ctx.Err() // Возвращает ошибку контекста, если контекст завершился
		default:
			return 0, err
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return strings.Count(string(body), searchString), nil
}
