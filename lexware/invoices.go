package lexware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rasche-thalhofer/lexware-go/types"
)

type invoicesClient struct{ client *Client }

func (c *invoicesClient) Create(ctx context.Context, invoice *types.InvoiceCreateRequest, finalize bool) (*types.ActionResult, error) {
	path := "/v1/invoices"
	if finalize {
		path += "?finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, invoice)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *invoicesClient) Get(ctx context.Context, id string) (*types.Invoice, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/invoices/"+id, nil)
	if err != nil {
		return nil, err
	}
	var invoice types.Invoice
	if err := json.Unmarshal(body, &invoice); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &invoice, nil
}

func (c *invoicesClient) Pursue(ctx context.Context, precedingSalesVoucherID string, invoice *types.InvoiceCreateRequest, finalize bool) (*types.ActionResult, error) {
	path := "/v1/invoices?precedingSalesVoucherId=" + precedingSalesVoucherID
	if finalize {
		path += "&finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, invoice)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *invoicesClient) RenderDocument(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "GET", "/v1/invoices/"+id+"/document", nil)
	return err
}

func (c *invoicesClient) DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/invoices/"+id+"/document", nil, map[string]string{"Accept": "application/pdf"})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}
	return resp.Body, nil
}
