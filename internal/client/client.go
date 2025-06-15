package client

import (
	"log/slog"
	"net/http"
	"proxynum2/internal/config"
)

type ClientProxy struct {
	cfg    config.ClientProxyConfig
	logger *slog.Logger
	client http.Client
}

func New(cfg config.ClientProxyConfig, logger *slog.Logger) *ClientProxy {
	return &ClientProxy{
		client: http.Client{},
		logger: logger,
		cfg:    cfg,
	}
}

func (cl ClientProxy) Proxy(r *http.Request) (*http.Response, error) {

	proxyRequest, err := http.NewRequest(r.Method, cl.cfg.URL, r.Body)
	if err != nil {
		return nil, err
	}

	for key, value := range r.Header {

		if len(value) == 1 {
			proxyRequest.Header.Set(key, value[0])
			continue
		}

		for _, header := range value {
			proxyRequest.Header.Add(key, header)
		}
	}

	res, err := cl.client.Do(proxyRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}
