package client

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"proxynum2/internal/config"
)

type ClientProxy struct {
	cfg    config.ClientProxyConfig
	logger *slog.Logger
	client *http.Client
}

func New(cfg config.ClientProxyConfig, logger *slog.Logger, client *http.Client) *ClientProxy {
	return &ClientProxy{
		client: client,
		logger: logger,
		cfg:    cfg,
	}
}

func (cl ClientProxy) Proxy(r *http.Request) (*http.Response, error) {

	ctx, cancel := context.WithTimeout(context.Background(), cl.cfg.Timeout)
	defer cancel()

	proxyRequest, err := http.NewRequestWithContext(ctx, r.Method, cl.cfg.URL, r.Body)
	if err != nil {
		err := fmt.Errorf("error Creating Request %s", err)
		return nil, err
	}

	for key, value := range r.Header {
		for _, header := range value {
			proxyRequest.Header.Add(key, header)
		}
	}

	res, err := cl.client.Do(proxyRequest)
	if err != nil {
		err := fmt.Errorf("error Sending Request %s", err)
		return nil, err
	}

	return res, nil
}
