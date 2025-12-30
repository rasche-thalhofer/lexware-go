// Package types provides type definitions for the Lexware API.
package types

import "time"

// Voucher represents a bookkeeping voucher in Lexware.
type Voucher struct {
	ID                   string        `json:"id,omitempty"`
	OrganizationID       string        `json:"organizationId,omitempty"`
	Type                 VoucherType   `json:"type,omitempty"`
	VoucherStatus        VoucherStatus `json:"voucherStatus,omitempty"`
	VoucherNumber        string        `json:"voucherNumber,omitempty"`
	VoucherDate          time.Time     `json:"voucherDate,omitempty"`
	ShippingDate         *time.Time    `json:"shippingDate,omitempty"`
	DueDate              *time.Time    `json:"dueDate,omitempty"`
	TotalGrossAmount     float64       `json:"totalGrossAmount,omitempty"`
	TotalTaxAmount       float64       `json:"totalTaxAmount,omitempty"`
	TaxType              TaxType       `json:"taxType,omitempty"`
	UseCollectiveContact bool          `json:"useCollectiveContact,omitempty"`
	Remark               string        `json:"remark,omitempty"`
	VoucherItems         []VoucherItem `json:"voucherItems,omitempty"`
	Files                []VoucherFile `json:"files,omitempty"`
	CreatedDate          time.Time     `json:"createdDate,omitempty"`
	UpdatedDate          time.Time     `json:"updatedDate,omitempty"`
	Version              int           `json:"version,omitempty"`
}

// VoucherType represents the type of a bookkeeping voucher.
type VoucherType string

const (
	VoucherTypeSalesInvoice       VoucherType = "salesinvoice"
	VoucherTypeSalesCreditNote    VoucherType = "salescreditnote"
	VoucherTypePurchaseInvoice    VoucherType = "purchaseinvoice"
	VoucherTypePurchaseCreditNote VoucherType = "purchasecreditnote"
	VoucherTypeInvoice            VoucherType = "invoice"
	VoucherTypeCreditNote         VoucherType = "creditnote"
	VoucherTypeOrderConfirmation  VoucherType = "orderconfirmation"
	VoucherTypeQuotation          VoucherType = "quotation"
	VoucherTypeDeliveryNote       VoucherType = "deliverynote"
	VoucherTypeDownPaymentInvoice VoucherType = "downpaymentinvoice"
)

// VoucherItem represents an item in a voucher.
type VoucherItem struct {
	Amount         float64 `json:"amount,omitempty"`
	TaxAmount      float64 `json:"taxAmount,omitempty"`
	TaxRatePercent float64 `json:"taxRatePercent,omitempty"`
	CategoryID     string  `json:"categoryId,omitempty"`
}

// VoucherFile represents a file attached to a voucher.
type VoucherFile struct {
	ID string `json:"id,omitempty"`
}

// VoucherCreateRequest represents the request body for creating a voucher.
type VoucherCreateRequest struct {
	Type                 VoucherType   `json:"type"`
	VoucherNumber        string        `json:"voucherNumber,omitempty"`
	VoucherDate          time.Time     `json:"voucherDate"`
	ShippingDate         *time.Time    `json:"shippingDate,omitempty"`
	DueDate              *time.Time    `json:"dueDate,omitempty"`
	TotalGrossAmount     float64       `json:"totalGrossAmount"`
	TotalTaxAmount       float64       `json:"totalTaxAmount"`
	TaxType              TaxType       `json:"taxType"`
	UseCollectiveContact bool          `json:"useCollectiveContact,omitempty"`
	ContactID            string        `json:"contactId,omitempty"`
	Remark               string        `json:"remark,omitempty"`
	VoucherItems         []VoucherItem `json:"voucherItems"`
}

// VoucherUpdateRequest represents the request body for updating a voucher.
type VoucherUpdateRequest struct {
	VoucherNumber        string        `json:"voucherNumber,omitempty"`
	VoucherDate          time.Time     `json:"voucherDate"`
	ShippingDate         *time.Time    `json:"shippingDate,omitempty"`
	DueDate              *time.Time    `json:"dueDate,omitempty"`
	TotalGrossAmount     float64       `json:"totalGrossAmount"`
	TotalTaxAmount       float64       `json:"totalTaxAmount"`
	TaxType              TaxType       `json:"taxType"`
	UseCollectiveContact bool          `json:"useCollectiveContact,omitempty"`
	ContactID            string        `json:"contactId,omitempty"`
	Remark               string        `json:"remark,omitempty"`
	VoucherItems         []VoucherItem `json:"voucherItems"`
	Version              int           `json:"version"`
}

// VoucherFilterOptions specifies the optional parameters for filtering vouchers.
type VoucherFilterOptions struct {
	VoucherNumber string
	VoucherStatus VoucherStatus
	ContactID     string
}
