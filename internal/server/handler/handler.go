package handler

import (
	"io"
	"log/slog"
	"net/http"
	cl "proxynum2/internal/client"
	"proxynum2/internal/server/utils"
)

type Handler struct {
	logger *slog.Logger
	client *cl.ClientProxy
}

func New(log *slog.Logger, client *cl.ClientProxy) *Handler {
	return &Handler{log, client}
}

func (ctr Handler) GetRequest(w http.ResponseWriter, r *http.Request) {

	ctr.logger.Debug("Get Request for Proxy")

	res, err := ctr.client.Proxy(r)
	if err != nil {
		http.Error(w, "Error ", http.StatusInternalServerError)
	}

	utils.CopyHeaders(w.Header(), res.Header)
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}
