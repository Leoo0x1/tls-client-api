package api

import (
	"context"

	tls_client_cffi_src "github.com/bogdanfinn/tls-client/cffi_src"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/httpserver"
	"github.com/justtrackio/gosoline/pkg/log"
)

type FreeAllHandler struct {
	logger log.Logger
}

func NewFreeAllHandler(ctx context.Context, config cfg.Config, logger log.Logger) (gin.HandlerFunc, error) {
	handler := FreeAllHandler{
		logger: logger,
	}

	return httpserver.CreateHandler(handler), nil
}

func (fh FreeAllHandler) Handle(ctx context.Context, request *httpserver.Request) (*httpserver.Response, error) {
	tls_client_cffi_src.ClearSessionCache()

	out := tls_client_cffi_src.DestroyOutput{
		Id:      uuid.New().String(),
		Success: true,
	}

	return httpserver.NewJsonResponse(out), nil
}
