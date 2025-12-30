// Package lexware provides a Go client for the Lexware API.
package lexware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/rasche-thalhofer/lexware-go/types"
)

const (
	// DefaultBaseURL is the default base URL for the Lexware API.
	DefaultBaseURL = "https://api.lexware.io"

	// DefaultTimeout is the default timeout for HTTP requests.
	DefaultTimeout = 30 * time.Second
)

// Client is the main entry point to the Lexware API.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client

	articles            ArticlesInterface
	contacts            ContactsInterface
	countries           CountriesInterface
	creditNotes         CreditNotesInterface
	deliveryNotes       DeliveryNotesInterface
	downPaymentInvoices DownPaymentInvoicesInterface
	dunnings            DunningsInterface
	eventSubscriptions  EventSubscriptionsInterface
	files               FilesInterface
	invoices            InvoicesInterface
	orderConfirmations  OrderConfirmationsInterface
	payments            PaymentsInterface
	paymentConditions   PaymentConditionsInterface
	postingCategories   PostingCategoriesInterface
	printLayouts        PrintLayoutsInterface
	profile             ProfileInterface
	quotations          QuotationsInterface
	recurringTemplates  RecurringTemplatesInterface
	voucherList         VoucherListInterface
	vouchers            VouchersInterface
}

// Config holds configuration options for the client.
type Config struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
	Timeout    time.Duration
}

// NewClient creates a new Lexware API client with the given API key.
func NewClient(apiKey string) (*Client, error) {
	return NewClientWithConfig(Config{APIKey: apiKey})
}

// NewClientWithConfig creates a new Lexware API client with the given configuration.
func NewClientWithConfig(config Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		timeout := config.Timeout
		if timeout == 0 {
			timeout = DefaultTimeout
		}
		httpClient = &http.Client{Timeout: timeout}
	}

	client := &Client{
		baseURL:    baseURL,
		apiKey:     config.APIKey,
		httpClient: httpClient,
	}

	client.articles = &articlesClient{client: client}
	client.contacts = &contactsClient{client: client}
	client.countries = &countriesClient{client: client}
	client.creditNotes = &creditNotesClient{client: client}
	client.deliveryNotes = &deliveryNotesClient{client: client}
	client.downPaymentInvoices = &downPaymentInvoicesClient{client: client}
	client.dunnings = &dunningsClient{client: client}
	client.eventSubscriptions = &eventSubscriptionsClient{client: client}
	client.files = &filesClient{client: client}
	client.invoices = &invoicesClient{client: client}
	client.orderConfirmations = &orderConfirmationsClient{client: client}
	client.payments = &paymentsClient{client: client}
	client.paymentConditions = &paymentConditionsClient{client: client}
	client.postingCategories = &postingCategoriesClient{client: client}
	client.printLayouts = &printLayoutsClient{client: client}
	client.profile = &profileClient{client: client}
	client.quotations = &quotationsClient{client: client}
	client.recurringTemplates = &recurringTemplatesClient{client: client}
	client.voucherList = &voucherListClient{client: client}
	client.vouchers = &vouchersClient{client: client}

	return client, nil
}

func (c *Client) Articles() ArticlesInterface                       { return c.articles }
func (c *Client) Contacts() ContactsInterface                       { return c.contacts }
func (c *Client) Countries() CountriesInterface                     { return c.countries }
func (c *Client) CreditNotes() CreditNotesInterface                 { return c.creditNotes }
func (c *Client) DeliveryNotes() DeliveryNotesInterface             { return c.deliveryNotes }
func (c *Client) DownPaymentInvoices() DownPaymentInvoicesInterface { return c.downPaymentInvoices }
func (c *Client) Dunnings() DunningsInterface                       { return c.dunnings }
func (c *Client) EventSubscriptions() EventSubscriptionsInterface   { return c.eventSubscriptions }
func (c *Client) Files() FilesInterface                             { return c.files }
func (c *Client) Invoices() InvoicesInterface                       { return c.invoices }
func (c *Client) OrderConfirmations() OrderConfirmationsInterface   { return c.orderConfirmations }
func (c *Client) Payments() PaymentsInterface                       { return c.payments }
func (c *Client) PaymentConditions() PaymentConditionsInterface     { return c.paymentConditions }
func (c *Client) PostingCategories() PostingCategoriesInterface     { return c.postingCategories }
func (c *Client) PrintLayouts() PrintLayoutsInterface               { return c.printLayouts }
func (c *Client) Profile() ProfileInterface                         { return c.profile }
func (c *Client) Quotations() QuotationsInterface                   { return c.quotations }
func (c *Client) RecurringTemplates() RecurringTemplatesInterface   { return c.recurringTemplates }
func (c *Client) VoucherList() VoucherListInterface                 { return c.voucherList }
func (c *Client) Vouchers() VouchersInterface                       { return c.vouchers }

func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{StatusCode: resp.StatusCode, Body: string(respBody)}
	}

	return respBody, nil
}

func (c *Client) doRequestRaw(ctx context.Context, method, path string, body interface{}, headers map[string]string) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return c.httpClient.Do(req)
}

func buildQueryString(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	values := url.Values{}
	for k, v := range params {
		if v != "" {
			values.Set(k, v)
		}
	}
	if qs := values.Encode(); qs != "" {
		return "?" + qs
	}
	return ""
}

func addPagination(params map[string]string, opts *types.ListOptions) {
	if opts == nil {
		return
	}
	if opts.Page > 0 {
		params["page"] = strconv.Itoa(opts.Page)
	}
	if opts.Size > 0 {
		params["size"] = strconv.Itoa(opts.Size)
	}
}

// APIError represents an error returned by the Lexware API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Lexware API error (status %d): %s", e.StatusCode, e.Body)
}

func (e *APIError) IsNotFound() bool     { return e.StatusCode == http.StatusNotFound }
func (e *APIError) IsConflict() bool     { return e.StatusCode == http.StatusConflict }
func (e *APIError) IsRateLimited() bool  { return e.StatusCode == http.StatusTooManyRequests }
func (e *APIError) IsUnauthorized() bool { return e.StatusCode == http.StatusUnauthorized }
