package wpgo

import (
	"github.com/imroc/req/v3"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("wpgo")

type Client struct {
	*req.Client

	Post PostService
}

func NewClient(baseUrl string) *Client {
	client := &Client{Client: req.NewClient().SetBaseURL(baseUrl).SetCommonErrorResult(&ErrorResponse{})}

	client.Post = NewPostService(client)

	return client
}
