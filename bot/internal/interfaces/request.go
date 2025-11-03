package interfaces

import (
	"context"
	"net/http"
)

type HttpClient interface {
	Get(ctx context.Context, url string, params map[string]string, headers map[string]string) (*http.Response, error)
	Post(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error)
	Patch(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error)
	Put(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error)
	Delete(ctx context.Context, url string, body []byte, headers map[string]string) (*http.Response, error)
}
