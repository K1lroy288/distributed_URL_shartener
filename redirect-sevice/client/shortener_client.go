package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortenerClient interface {
	Resolve(ctx context.Context, shortCode string) (string, error)
}

type shortenerClient struct {
	baseURL string
	client  *http.Client
}

func NewShortenerClient(baseURL string) *shortenerClient {
	return &shortenerClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

type resolveResponse struct {
	LongURL string `json:"long_url"`
}

func (s *shortenerClient) Resolve(ctx context.Context, shortCode string) (string, error) {
	url := fmt.Sprintf("%s/short/%s", s.baseURL, shortCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("link not found")
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("shortener service error: %d", resp.StatusCode)
	}

	var res resolveResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res.LongURL, nil
}
