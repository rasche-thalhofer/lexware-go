package lexware

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rasche-thalhofer/lexware-go/types"
)

type contactsClient struct{ client *Client }

func (c *contactsClient) Create(ctx context.Context, contact *types.ContactCreateRequest) (*types.ActionResult, error) {
	body, err := c.client.doRequest(ctx, "POST", "/v1/contacts", contact)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *contactsClient) Get(ctx context.Context, id string) (*types.Contact, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/contacts/"+id, nil)
	if err != nil {
		return nil, err
	}
	var contact types.Contact
	if err := json.Unmarshal(body, &contact); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &contact, nil
}

func (c *contactsClient) Update(ctx context.Context, id string, contact *types.ContactUpdateRequest) (*types.ActionResult, error) {
	body, err := c.client.doRequest(ctx, "PUT", "/v1/contacts/"+id, contact)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *contactsClient) List(ctx context.Context, opts *types.ListOptions, filter *types.ContactFilterOptions) (*types.Page[types.Contact], error) {
	params := make(map[string]string)
	addPagination(params, opts)
	if filter != nil {
		if filter.Email != "" {
			params["email"] = filter.Email
		}
		if filter.Name != "" {
			params["name"] = filter.Name
		}
		if filter.Number != 0 {
			params["number"] = strconv.Itoa(filter.Number)
		}
		if filter.Customer {
			params["customer"] = "true"
		}
		if filter.Vendor {
			params["vendor"] = "true"
		}
	}
	body, err := c.client.doRequest(ctx, "GET", "/v1/contacts"+buildQueryString(params), nil)
	if err != nil {
		return nil, err
	}
	var page types.Page[types.Contact]
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &page, nil
}
