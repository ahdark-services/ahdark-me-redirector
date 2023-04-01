package wpgo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/samber/lo"
	"go.opentelemetry.io/otel/attribute"

	"github.com/ahdark-services/ahdark-me-redirector/pkg/util"
)

type PostService interface {
	Get(ctx context.Context, id int, params GetPostParams) (*Post, error)
	List(ctx context.Context, params ListPostParams) (*PostList, error)
}

func NewPostService(client *Client) PostService {
	return &postServiceOp{client: client}
}

type postServiceOp struct {
	client *Client
}

func (op *postServiceOp) List(ctx context.Context, params ListPostParams) (*PostList, error) {
	ctx, span := tracer.Start(ctx, "WpGo Post List")
	defer span.End()

	query := map[string]string{
		"context":            string(params.Context),
		"page":               strconv.Itoa(params.Page),
		"per_page":           strconv.Itoa(params.PerPage),
		"search":             params.Search,
		"after":              params.After,
		"author":             util.JoinIntSlice(params.Author, ","),
		"author_exclude":     util.JoinIntSlice(params.AuthorExclude, ","),
		"before":             params.Before,
		"exclude":            util.JoinIntSlice(params.Exclude, ","),
		"include":            util.JoinIntSlice(params.Include, ","),
		"offset":             strconv.Itoa(params.Offset),
		"order":              string(params.Order),
		"orderby":            params.OrderBy,
		"slug":               params.Slug,
		"status":             params.Status,
		"tax_relation":       params.TaxRelation,
		"categories":         util.JoinIntSlice(params.Categories, ","),
		"categories_exclude": util.JoinIntSlice(params.CategoriesExclude, ","),
		"tags":               util.JoinIntSlice(params.Tags, ","),
		"tags_exclude":       util.JoinIntSlice(params.TagsExclude, ","),
		"sticky":             strconv.FormatBool(params.Sticky),
	}

	span.SetAttributes(lo.MapToSlice(query, func(key string, value string) attribute.KeyValue {
		return attribute.String(fmt.Sprintf("wp.query.%s", key), value)
	})...)

	resp, err := op.client.R().SetContext(ctx).
		SetQueryParams(query).
		SetSuccessResult(&PostList{}).
		Get("/wp-json/wp/v2/posts")
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("error listing posts: %op", resp.String())
	}

	return resp.SuccessResult().(*PostList), nil
}

func (op *postServiceOp) Get(ctx context.Context, id int, params GetPostParams) (*Post, error) {
	ctx, span := tracer.Start(ctx, fmt.Sprintf("WpGo Post Get: %d", id))
	defer span.End()

	query := map[string]string{
		"context":  string(params.Context),
		"password": params.Password,
	}

	span.SetAttributes(lo.MapToSlice(query, func(key string, value string) attribute.KeyValue {
		return attribute.String(fmt.Sprintf("wp.query.%s", key), value)
	})...)

	resp, err := op.client.R().SetContext(ctx).
		SetSuccessResult(&Post{}).
		Get(fmt.Sprintf("/wp-json/wp/v2/posts/%d", id))
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("error getting post: %s", resp.String())
	}

	return resp.SuccessResult().(*Post), nil
}
