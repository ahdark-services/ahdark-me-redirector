package blog

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/pkg/cache"
	"github.com/ahdark-services/ahdark-me-redirector/pkg/wpgo"
)

var tracer = otel.Tracer("service.blog")

type Service interface {
	GetPostUrl(ctx context.Context, id int) (string, error)
}

type service struct {
	fx.In
	WpGo  *wpgo.Client
	Cache cache.Driver
}

func NewService(svc service) Service {
	return &svc
}

const (
	postUrlCacheKeyFormat = "blog-post::%d"
)

// GetPostUrl implements Service.GetPostUrl
func (s *service) GetPostUrl(ctx context.Context, id int) (string, error) {
	ctx, span := tracer.Start(ctx, fmt.Sprintf("Get Post Url: %d", id))
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.Int("id", id),
		)
	}

	// Check cache first
	if url, ok := s.Cache.Get(ctx, fmt.Sprintf(postUrlCacheKeyFormat, id)); ok {
		return url.(string), nil
	}

	post, err := s.WpGo.Post.Get(ctx, id, wpgo.GetPostParams{})
	if err != nil {
		return "", err
	}

	_ = s.Cache.Set(ctx, fmt.Sprintf(postUrlCacheKeyFormat, id), post.Link, 0)

	return post.Link, nil
}
