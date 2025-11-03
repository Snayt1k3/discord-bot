package adapters

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type DefaultHttpClient struct {
	Client *http.Client
}

func NewDefaultHttpClient() *DefaultHttpClient {
	return &DefaultHttpClient{Client: &http.Client{}}
}

func (d *DefaultHttpClient) Get(ctx context.Context, rawURL string, params map[string]string, headers map[string]string) (*http.Response, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return d.doRequest(ctx, http.MethodGet, u.String(), nil, headers)
}

func (d *DefaultHttpClient) Post(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error) {
	return d.doRequest(ctx, http.MethodPost, url, body, headers)
}

func (d *DefaultHttpClient) Patch(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error) {
	return d.doRequest(ctx, http.MethodPatch, url, body, headers)
}

func (d *DefaultHttpClient) Put(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error) {
	return d.doRequest(ctx, http.MethodPut, url, body, headers)
}

func (d *DefaultHttpClient) Delete(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error) {
	return d.doRequest(ctx, http.MethodDelete, url, body, headers)
}

func (d *DefaultHttpClient) doRequest(ctx context.Context, method, url string, body []byte, headers map[string]string) (*http.Response, error) {
	var reqBody io.Reader

	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)

	if err != nil {
		slog.Error("Error creating request", "err", err)
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return d.Client.Do(req)
}
