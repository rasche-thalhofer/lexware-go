package lexware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rasche-thalhofer/lexware-go/types"
)

type countriesClient struct{ client *Client }

func (c *countriesClient) List(ctx context.Context) ([]types.Country, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/countries", nil)
	if err != nil {
		return nil, err
	}
	var countries []types.Country
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return countries, nil
}

type profileClient struct{ client *Client }

func (c *profileClient) Get(ctx context.Context) (*types.Profile, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/profile", nil)
	if err != nil {
		return nil, err
	}
	var profile types.Profile
	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &profile, nil
}

type paymentsClient struct{ client *Client }

func (c *paymentsClient) Get(ctx context.Context, id string) (*types.Payment, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/payments/"+id, nil)
	if err != nil {
		return nil, err
	}
	var payment types.Payment
	if err := json.Unmarshal(body, &payment); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &payment, nil
}

type paymentConditionsClient struct{ client *Client }

func (c *paymentConditionsClient) List(ctx context.Context) ([]types.PaymentCondition, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/payment-conditions", nil)
	if err != nil {
		return nil, err
	}
	var conditions []types.PaymentCondition
	if err := json.Unmarshal(body, &conditions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return conditions, nil
}

type postingCategoriesClient struct{ client *Client }

func (c *postingCategoriesClient) List(ctx context.Context) ([]types.PostingCategory, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/posting-categories", nil)
	if err != nil {
		return nil, err
	}
	var categories []types.PostingCategory
	if err := json.Unmarshal(body, &categories); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return categories, nil
}

type printLayoutsClient struct{ client *Client }

func (c *printLayoutsClient) List(ctx context.Context) ([]types.PrintLayout, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/print-layouts", nil)
	if err != nil {
		return nil, err
	}
	var layouts []types.PrintLayout
	if err := json.Unmarshal(body, &layouts); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return layouts, nil
}

type recurringTemplatesClient struct{ client *Client }

func (c *recurringTemplatesClient) Get(ctx context.Context, id string) (*types.RecurringTemplate, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/recurring-templates/"+id, nil)
	if err != nil {
		return nil, err
	}
	var template types.RecurringTemplate
	if err := json.Unmarshal(body, &template); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &template, nil
}

func (c *recurringTemplatesClient) List(ctx context.Context, opts *types.ListOptions) (*types.Page[types.RecurringTemplate], error) {
	params := make(map[string]string)
	addPagination(params, opts)
	body, err := c.client.doRequest(ctx, "GET", "/v1/recurring-templates"+buildQueryString(params), nil)
	if err != nil {
		return nil, err
	}
	var page types.Page[types.RecurringTemplate]
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &page, nil
}
