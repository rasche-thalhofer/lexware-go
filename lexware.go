// Package lexwarego provides a Go client library for the Lexware API.
//
// This is the main entry point for the library. Import this package to use the Lexware API client.
//
// # Quick Start
//
//	import (
//	    "context"
//	    "log"
//
//	    lexware "github.com/rasche-thalhofer/lexware-go"
//	    "github.com/rasche-thalhofer/lexware-go/types"
//	)
//
//	func main() {
//	    client, err := lexware.NewClient("your-api-key")
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    profile, err := client.Profile().Get(context.Background())
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    log.Printf("Organization: %s", profile.CompanyName)
//	}
package lexwarego

import (
	"github.com/rasche-thalhofer/lexware-go/lexware"
	"github.com/rasche-thalhofer/lexware-go/types"
)

// Re-export main client types
type (
	Client   = lexware.Client
	Config   = lexware.Config
	APIError = lexware.APIError
)

// Re-export common types
type (
	Article                        = types.Article
	ArticleCreateRequest           = types.ArticleCreateRequest
	ArticleUpdateRequest           = types.ArticleUpdateRequest
	Contact                        = types.Contact
	ContactCreateRequest           = types.ContactCreateRequest
	ContactUpdateRequest           = types.ContactUpdateRequest
	Invoice                        = types.Invoice
	InvoiceCreateRequest           = types.InvoiceCreateRequest
	Quotation                      = types.Quotation
	QuotationCreateRequest         = types.QuotationCreateRequest
	CreditNote                     = types.CreditNote
	CreditNoteCreateRequest        = types.CreditNoteCreateRequest
	DeliveryNote                   = types.DeliveryNote
	DeliveryNoteCreateRequest      = types.DeliveryNoteCreateRequest
	Dunning                        = types.Dunning
	DunningCreateRequest           = types.DunningCreateRequest
	Voucher                        = types.Voucher
	VoucherCreateRequest           = types.VoucherCreateRequest
	VoucherUpdateRequest           = types.VoucherUpdateRequest
	EventSubscription              = types.EventSubscription
	EventSubscriptionCreateRequest = types.EventSubscriptionCreateRequest
	Profile                        = types.Profile
	Country                        = types.Country
	Payment                        = types.Payment
	PaymentCondition               = types.PaymentCondition
	PostingCategory                = types.PostingCategory
	PrintLayout                    = types.PrintLayout
	RecurringTemplate              = types.RecurringTemplate
	ListOptions                    = types.ListOptions
	ActionResult                   = types.ActionResult
	Address                        = types.Address
	LineItem                       = types.LineItem
	UnitPrice                      = types.UnitPrice
	TotalPrice                     = types.TotalPrice
	TaxConditions                  = types.TaxConditions
	PaymentConditions              = types.PaymentConditions
	ShippingConditions             = types.ShippingConditions
)

// NewClient creates a new Lexware API client with the given API key.
func NewClient(apiKey string) (*Client, error) {
	return lexware.NewClient(apiKey)
}

// NewClientWithConfig creates a new Lexware API client with the given configuration.
func NewClientWithConfig(config Config) (*Client, error) {
	return lexware.NewClientWithConfig(config)
}

// Constants
const (
	DefaultBaseURL = lexware.DefaultBaseURL
	DefaultTimeout = lexware.DefaultTimeout
)
