package client

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"proxynum2/internal/config"
	"proxynum2/internal/server/utils"
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
		return nil, fmt.Errorf("error Creating Request %s", err)
	}

	utils.CopyHeaders(r.Header, proxyRequest.Header)

	res, err := cl.client.Do(proxyRequest)
	if err != nil {
		return nil, fmt.Errorf("error Sending Request %s", err)
	}

	return res, nil
}
