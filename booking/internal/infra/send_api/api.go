package send_api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/huytran2000-hcmus/bTaskee/booking/config"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/pkg/httpapi"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/repository"
)

var _ repository.SendRepository = (*SendAPI)(nil)

type SendAPI struct{}

func New() *SendAPI {
	return &SendAPI{}
}

func (p *SendAPI) SendTask(ctx context.Context, req model.SendTaskRequest) (model.SendTaskResponse, error) {
	httpReq := httpapi.HTTPRequest{
		Method: http.MethodPost,
		URL:    config.LoadEnv().SendTaskURI,
		Body:   req,
	}

	res, err := httpapi.SendHTTPRequest[model.SendTaskPayload](ctx, httpReq, httpapi.HTTPOptions{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		return res.Data, err
	}

	if res.Error != "" {
		return res.Data, errors.New(res.Error)
	}

	return res.Data, err
}
