package requester

import (
	"context"
	"net/http"
)

type Requester interface {
	Get(url string) (*http.Response, error)
}

type requester struct {
	ctx context.Context
}

func NewRequester(ctx context.Context) Requester {
	return &requester{
		ctx: ctx,
	}
}

func (r requester) Get(url string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(
		r.ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(request)
}
