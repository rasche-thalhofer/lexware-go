package lexware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/rasche-thalhofer/lexware-go/types"
)

type vouchersClient struct{ client *Client }

func (c *vouchersClient) Create(ctx context.Context, voucher *types.VoucherCreateRequest) (*types.ActionResult, error) {
	body, err := c.client.doRequest(ctx, "POST", "/v1/vouchers", voucher)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *vouchersClient) Get(ctx context.Context, id string) (*types.Voucher, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/vouchers/"+id, nil)
	if err != nil {
		return nil, err
	}
	var voucher types.Voucher
	if err := json.Unmarshal(body, &voucher); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &voucher, nil
}

func (c *vouchersClient) Update(ctx context.Context, id string, voucher *types.VoucherUpdateRequest) (*types.ActionResult, error) {
	body, err := c.client.doRequest(ctx, "PUT", "/v1/vouchers/"+id, voucher)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *vouchersClient) List(ctx context.Context, opts *types.ListOptions, filter *types.VoucherFilterOptions) (*types.Page[types.Voucher], error) {
	params := make(map[string]string)
	addPagination(params, opts)
	if filter != nil {
		if filter.VoucherNumber != "" {
			params["voucherNumber"] = filter.VoucherNumber
		}
		if filter.VoucherStatus != "" {
			params["voucherStatus"] = string(filter.VoucherStatus)
		}
		if filter.ContactID != "" {
			params["contactId"] = filter.ContactID
		}
	}
	body, err := c.client.doRequest(ctx, "GET", "/v1/vouchers"+buildQueryString(params), nil)
	if err != nil {
		return nil, err
	}
	var page types.Page[types.Voucher]
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &page, nil
}

func (c *vouchersClient) UploadFile(ctx context.Context, id string, filename string, content io.Reader) error {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, content); err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", c.client.baseURL+"/v1/vouchers/"+id+"/files", &buf)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.client.apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return &APIError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}
	return nil
}

type voucherListClient struct{ client *Client }

func (c *voucherListClient) List(ctx context.Context, opts *types.ListOptions, filter *types.VoucherListFilterOptions) (*types.Page[types.VoucherListItem], error) {
	params := make(map[string]string)
	addPagination(params, opts)
	if filter != nil {
		if filter.VoucherType != "" {
			params["voucherType"] = string(filter.VoucherType)
		}
		if filter.VoucherStatus != "" {
			params["voucherStatus"] = string(filter.VoucherStatus)
		}
		if filter.Archived != nil {
			if *filter.Archived {
				params["archived"] = "true"
			} else {
				params["archived"] = "false"
			}
		}
		if filter.ContactID != "" {
			params["contactId"] = filter.ContactID
		}
		if filter.VoucherDateFrom != "" {
			params["voucherDateFrom"] = filter.VoucherDateFrom
		}
		if filter.VoucherDateTo != "" {
			params["voucherDateTo"] = filter.VoucherDateTo
		}
		if filter.CreatedDateFrom != "" {
			params["createdDateFrom"] = filter.CreatedDateFrom
		}
		if filter.CreatedDateTo != "" {
			params["createdDateTo"] = filter.CreatedDateTo
		}
		if filter.UpdatedDateFrom != "" {
			params["updatedDateFrom"] = filter.UpdatedDateFrom
		}
		if filter.UpdatedDateTo != "" {
			params["updatedDateTo"] = filter.UpdatedDateTo
		}
	}
	body, err := c.client.doRequest(ctx, "GET", "/v1/voucherlist"+buildQueryString(params), nil)
	if err != nil {
		return nil, err
	}
	var page types.Page[types.VoucherListItem]
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &page, nil
}
