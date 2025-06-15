package handler

import (
	"io"
	"log/slog"
	"net/http"
	cl "proxynum2/internal/client"
)

type Handler struct {
	logger *slog.Logger
	client *cl.ClientProxy
}

func New(log *slog.Logger, client *cl.ClientProxy) *Handler {
	return &Handler{log, client}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func (ctr Handler) GetRequest(w http.ResponseWriter, r *http.Request) {

	ctr.logger.Debug("Get Request for Proxy")

	res, err := ctr.client.Proxy(r)
	if err != nil {
		http.Error(w, "Error ", http.StatusInternalServerError)
	}

	copyHeader(w.Header(), res.Header)
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}
