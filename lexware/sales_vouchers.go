package lexware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rasche-thalhofer/lexware-go/types"
)

// Quotations
type quotationsClient struct{ client *Client }

func (c *quotationsClient) Create(ctx context.Context, quotation *types.QuotationCreateRequest, finalize bool) (*types.ActionResult, error) {
	path := "/v1/quotations"
	if finalize {
		path += "?finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, quotation)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *quotationsClient) Get(ctx context.Context, id string) (*types.Quotation, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/quotations/"+id, nil)
	if err != nil {
		return nil, err
	}
	var quotation types.Quotation
	if err := json.Unmarshal(body, &quotation); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &quotation, nil
}

func (c *quotationsClient) RenderDocument(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "GET", "/v1/quotations/"+id+"/document", nil)
	return err
}

func (c *quotationsClient) DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/quotations/"+id+"/document", nil, map[string]string{"Accept": "application/pdf"})
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

// Credit Notes
type creditNotesClient struct{ client *Client }

func (c *creditNotesClient) Create(ctx context.Context, creditNote *types.CreditNoteCreateRequest, finalize bool) (*types.ActionResult, error) {
	path := "/v1/credit-notes"
	if finalize {
		path += "?finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, creditNote)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *creditNotesClient) Get(ctx context.Context, id string) (*types.CreditNote, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/credit-notes/"+id, nil)
	if err != nil {
		return nil, err
	}
	var creditNote types.CreditNote
	if err := json.Unmarshal(body, &creditNote); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &creditNote, nil
}

func (c *creditNotesClient) Pursue(ctx context.Context, precedingSalesVoucherID string, creditNote *types.CreditNoteCreateRequest, finalize bool) (*types.ActionResult, error) {
	path := "/v1/credit-notes?precedingSalesVoucherId=" + precedingSalesVoucherID
	if finalize {
		path += "&finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, creditNote)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *creditNotesClient) RenderDocument(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "GET", "/v1/credit-notes/"+id+"/document", nil)
	return err
}

func (c *creditNotesClient) DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/credit-notes/"+id+"/document", nil, map[string]string{"Accept": "application/pdf"})
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

// Delivery Notes
type deliveryNotesClient struct{ client *Client }

func (c *deliveryNotesClient) Create(ctx context.Context, deliveryNote *types.DeliveryNoteCreateRequest, finalize bool) (*types.ActionResult, error) {
	path := "/v1/delivery-notes"
	if finalize {
		path += "?finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, deliveryNote)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *deliveryNotesClient) Get(ctx context.Context, id string) (*types.DeliveryNote, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/delivery-notes/"+id, nil)
	if err != nil {
		return nil, err
	}
	var deliveryNote types.DeliveryNote
	if err := json.Unmarshal(body, &deliveryNote); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &deliveryNote, nil
}

func (c *deliveryNotesClient) Pursue(ctx context.Context, precedingSalesVoucherID string, deliveryNote *types.DeliveryNoteCreateRequest, finalize bool) (*types.ActionResult, error) {
	path := "/v1/delivery-notes?precedingSalesVoucherId=" + precedingSalesVoucherID
	if finalize {
		path += "&finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, deliveryNote)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *deliveryNotesClient) RenderDocument(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "GET", "/v1/delivery-notes/"+id+"/document", nil)
	return err
}

func (c *deliveryNotesClient) DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/delivery-notes/"+id+"/document", nil, map[string]string{"Accept": "application/pdf"})
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

// Dunnings
type dunningsClient struct{ client *Client }

func (c *dunningsClient) Create(ctx context.Context, precedingSalesVoucherID string, dunning *types.DunningCreateRequest) (*types.ActionResult, error) {
	path := "/v1/dunnings?precedingSalesVoucherId=" + precedingSalesVoucherID
	body, err := c.client.doRequest(ctx, "POST", path, dunning)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *dunningsClient) Get(ctx context.Context, id string) (*types.Dunning, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/dunnings/"+id, nil)
	if err != nil {
		return nil, err
	}
	var dunning types.Dunning
	if err := json.Unmarshal(body, &dunning); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &dunning, nil
}

func (c *dunningsClient) Pursue(ctx context.Context, precedingSalesVoucherID string, dunning *types.DunningCreateRequest) (*types.ActionResult, error) {
	return c.Create(ctx, precedingSalesVoucherID, dunning)
}

func (c *dunningsClient) RenderDocument(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "GET", "/v1/dunnings/"+id+"/document", nil)
	return err
}

func (c *dunningsClient) DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/dunnings/"+id+"/document", nil, map[string]string{"Accept": "application/pdf"})
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

// Order Confirmations
type orderConfirmationsClient struct{ client *Client }

func (c *orderConfirmationsClient) Create(ctx context.Context, orderConfirmation *types.OrderConfirmation, finalize bool) (*types.ActionResult, error) {
	path := "/v1/order-confirmations"
	if finalize {
		path += "?finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, orderConfirmation)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *orderConfirmationsClient) Get(ctx context.Context, id string) (*types.OrderConfirmation, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/order-confirmations/"+id, nil)
	if err != nil {
		return nil, err
	}
	var orderConfirmation types.OrderConfirmation
	if err := json.Unmarshal(body, &orderConfirmation); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &orderConfirmation, nil
}

func (c *orderConfirmationsClient) Pursue(ctx context.Context, precedingSalesVoucherID string, orderConfirmation *types.OrderConfirmation, finalize bool) (*types.ActionResult, error) {
	path := "/v1/order-confirmations?precedingSalesVoucherId=" + precedingSalesVoucherID
	if finalize {
		path += "&finalize=true"
	}
	body, err := c.client.doRequest(ctx, "POST", path, orderConfirmation)
	if err != nil {
		return nil, err
	}
	var result types.ActionResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (c *orderConfirmationsClient) RenderDocument(ctx context.Context, id string) error {
	_, err := c.client.doRequest(ctx, "GET", "/v1/order-confirmations/"+id+"/document", nil)
	return err
}

func (c *orderConfirmationsClient) DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/order-confirmations/"+id+"/document", nil, map[string]string{"Accept": "application/pdf"})
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

// Down Payment Invoices
type downPaymentInvoicesClient struct{ client *Client }

func (c *downPaymentInvoicesClient) Get(ctx context.Context, id string) (*types.DownPaymentInvoice, error) {
	body, err := c.client.doRequest(ctx, "GET", "/v1/down-payment-invoices/"+id, nil)
	if err != nil {
		return nil, err
	}
	var invoice types.DownPaymentInvoice
	if err := json.Unmarshal(body, &invoice); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &invoice, nil
}

func (c *downPaymentInvoicesClient) DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error) {
	resp, err := c.client.doRequestRaw(ctx, "GET", "/v1/down-payment-invoices/"+id+"/document", nil, map[string]string{"Accept": "application/pdf"})
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
