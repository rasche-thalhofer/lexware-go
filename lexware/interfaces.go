package lexware

import (
	"context"
	"io"

	"github.com/rasche-thalhofer/lexware-go/types"
)

// ArticlesInterface provides methods for managing articles.
type ArticlesInterface interface {
	Create(ctx context.Context, article *types.ArticleCreateRequest) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.Article, error)
	Update(ctx context.Context, id string, article *types.ArticleUpdateRequest) (*types.ActionResult, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, opts *types.ListOptions, filter *types.ArticleFilterOptions) (*types.Page[types.Article], error)
}

// ContactsInterface provides methods for managing contacts.
type ContactsInterface interface {
	Create(ctx context.Context, contact *types.ContactCreateRequest) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.Contact, error)
	Update(ctx context.Context, id string, contact *types.ContactUpdateRequest) (*types.ActionResult, error)
	List(ctx context.Context, opts *types.ListOptions, filter *types.ContactFilterOptions) (*types.Page[types.Contact], error)
}

// CountriesInterface provides methods for retrieving countries.
type CountriesInterface interface {
	List(ctx context.Context) ([]types.Country, error)
}

// CreditNotesInterface provides methods for managing credit notes.
type CreditNotesInterface interface {
	Create(ctx context.Context, creditNote *types.CreditNoteCreateRequest, finalize bool) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.CreditNote, error)
	Pursue(ctx context.Context, precedingSalesVoucherID string, creditNote *types.CreditNoteCreateRequest, finalize bool) (*types.ActionResult, error)
	RenderDocument(ctx context.Context, id string) error
	DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error)
}

// DeliveryNotesInterface provides methods for managing delivery notes.
type DeliveryNotesInterface interface {
	Create(ctx context.Context, deliveryNote *types.DeliveryNoteCreateRequest, finalize bool) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.DeliveryNote, error)
	Pursue(ctx context.Context, precedingSalesVoucherID string, deliveryNote *types.DeliveryNoteCreateRequest, finalize bool) (*types.ActionResult, error)
	RenderDocument(ctx context.Context, id string) error
	DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error)
}

// DownPaymentInvoicesInterface provides methods for down payment invoices (read-only).
type DownPaymentInvoicesInterface interface {
	Get(ctx context.Context, id string) (*types.DownPaymentInvoice, error)
	DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error)
}

// DunningsInterface provides methods for managing dunnings.
type DunningsInterface interface {
	Create(ctx context.Context, precedingSalesVoucherID string, dunning *types.DunningCreateRequest) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.Dunning, error)
	Pursue(ctx context.Context, precedingSalesVoucherID string, dunning *types.DunningCreateRequest) (*types.ActionResult, error)
	RenderDocument(ctx context.Context, id string) error
	DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error)
}

// EventSubscriptionsInterface provides methods for managing event subscriptions.
type EventSubscriptionsInterface interface {
	Create(ctx context.Context, subscription *types.EventSubscriptionCreateRequest) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.EventSubscription, error)
	List(ctx context.Context) ([]types.EventSubscription, error)
	Delete(ctx context.Context, id string) error
}

// FilesInterface provides methods for managing files.
type FilesInterface interface {
	Upload(ctx context.Context, filename string, content io.Reader, fileType types.FileUploadType) (*types.FileUploadResponse, error)
	Download(ctx context.Context, id string) (io.ReadCloser, error)
}

// InvoicesInterface provides methods for managing invoices.
type InvoicesInterface interface {
	Create(ctx context.Context, invoice *types.InvoiceCreateRequest, finalize bool) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.Invoice, error)
	Pursue(ctx context.Context, precedingSalesVoucherID string, invoice *types.InvoiceCreateRequest, finalize bool) (*types.ActionResult, error)
	RenderDocument(ctx context.Context, id string) error
	DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error)
}

// OrderConfirmationsInterface provides methods for managing order confirmations.
type OrderConfirmationsInterface interface {
	Create(ctx context.Context, orderConfirmation *types.OrderConfirmation, finalize bool) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.OrderConfirmation, error)
	Pursue(ctx context.Context, precedingSalesVoucherID string, orderConfirmation *types.OrderConfirmation, finalize bool) (*types.ActionResult, error)
	RenderDocument(ctx context.Context, id string) error
	DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error)
}

// PaymentsInterface provides methods for retrieving payment information.
type PaymentsInterface interface {
	Get(ctx context.Context, id string) (*types.Payment, error)
}

// PaymentConditionsInterface provides methods for retrieving payment conditions.
type PaymentConditionsInterface interface {
	List(ctx context.Context) ([]types.PaymentCondition, error)
}

// PostingCategoriesInterface provides methods for retrieving posting categories.
type PostingCategoriesInterface interface {
	List(ctx context.Context) ([]types.PostingCategory, error)
}

// PrintLayoutsInterface provides methods for retrieving print layouts.
type PrintLayoutsInterface interface {
	List(ctx context.Context) ([]types.PrintLayout, error)
}

// ProfileInterface provides methods for retrieving profile information.
type ProfileInterface interface {
	Get(ctx context.Context) (*types.Profile, error)
}

// QuotationsInterface provides methods for managing quotations.
type QuotationsInterface interface {
	Create(ctx context.Context, quotation *types.QuotationCreateRequest, finalize bool) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.Quotation, error)
	RenderDocument(ctx context.Context, id string) error
	DownloadDocument(ctx context.Context, id string) (io.ReadCloser, error)
}

// RecurringTemplatesInterface provides methods for recurring templates.
type RecurringTemplatesInterface interface {
	Get(ctx context.Context, id string) (*types.RecurringTemplate, error)
	List(ctx context.Context, opts *types.ListOptions) (*types.Page[types.RecurringTemplate], error)
}

// VoucherListInterface provides methods for retrieving the voucher list.
type VoucherListInterface interface {
	List(ctx context.Context, opts *types.ListOptions, filter *types.VoucherListFilterOptions) (*types.Page[types.VoucherListItem], error)
}

// VouchersInterface provides methods for managing vouchers.
type VouchersInterface interface {
	Create(ctx context.Context, voucher *types.VoucherCreateRequest) (*types.ActionResult, error)
	Get(ctx context.Context, id string) (*types.Voucher, error)
	Update(ctx context.Context, id string, voucher *types.VoucherUpdateRequest) (*types.ActionResult, error)
	List(ctx context.Context, opts *types.ListOptions, filter *types.VoucherFilterOptions) (*types.Page[types.Voucher], error)
	UploadFile(ctx context.Context, id string, filename string, content io.Reader) error
}
