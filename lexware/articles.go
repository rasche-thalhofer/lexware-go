package lexware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rasche-thalhofer/lexware-go/types"
)

type articlesClient struct{ client *Client }

func (c *articlesClient) Create(ctx context.Context, article *types.ArticleCreateRequest) (*types.ActionResult, error) {
	body, err := c.client.doRequest(ctx, "POST", "/v1/articles", article)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *articlesClient) Get(ctx context.Context, id string) (*types.Article, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/articles/"+id, nil)
	if err != nil {
		return nil, err
	}
	var article types.Article
	if err := json.Unmarshal(body, &article); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &article, nil
}

func (c *articlesClient) Update(ctx context.Context, id string, article *types.ArticleUpdateRequest) (*types.ActionResult, error) {
	body, err := c.client.doRequest(ctx, "PUT", "/v1/articles/"+id, article)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *articlesClient) Delete(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "DELETE", "/v1/articles/"+id, nil)
	return err
}

func (c *articlesClient) List(ctx context.Context, opts *types.ListOptions, filter *types.ArticleFilterOptions) (*types.Page[types.Article], error) {
	params := make(map[string]string)
	addPagination(params, opts)
	if filter != nil {
		if filter.ArticleNumber != "" {
			params["articleNumber"] = filter.ArticleNumber
		}
		if filter.GTIN != "" {
			params["gtin"] = filter.GTIN
		}
		if filter.Type != "" {
			params["type"] = string(filter.Type)
		}
	}
	body, err := c.client.doRequest(ctx, "GET", "/v1/articles"+buildQueryString(params), nil)
	if err != nil {
		return nil, err
	}
	var page types.Page[types.Article]
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &page, nil
}
