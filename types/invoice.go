// Package types provides type definitions for the Lexware API.
package types

import "time"

// Invoice represents an invoice in Lexware.
type Invoice struct {
	ID                        string                    `json:"id,omitempty"`
	OrganizationID            string                    `json:"organizationId,omitempty"`
	CreatedDate               time.Time                 `json:"createdDate,omitempty"`
	UpdatedDate               time.Time                 `json:"updatedDate,omitempty"`
	Version                   int                       `json:"version,omitempty"`
	Language                  string                    `json:"language,omitempty"`
	Archived                  bool                      `json:"archived,omitempty"`
	VoucherStatus             VoucherStatus             `json:"voucherStatus,omitempty"`
	VoucherNumber             string                    `json:"voucherNumber,omitempty"`
	VoucherDate               time.Time                 `json:"voucherDate,omitempty"`
	DueDate                   *time.Time                `json:"dueDate,omitempty"`
	Address                   *Address                  `json:"address,omitempty"`
	XRechnung                 *XRechnung                `json:"xRechnung,omitempty"`
	ElectronicDocumentProfile ElectronicDocumentProfile `json:"electronicDocumentProfile,omitempty"`
	LineItems                 []LineItem                `json:"lineItems,omitempty"`
	TotalPrice                *TotalPrice               `json:"totalPrice,omitempty"`
	TaxAmounts                []TaxAmount               `json:"taxAmounts,omitempty"`
	TaxConditions             *TaxConditions            `json:"taxConditions,omitempty"`
	PaymentConditions         *PaymentConditions        `json:"paymentConditions,omitempty"`
	ShippingConditions        *ShippingConditions       `json:"shippingConditions,omitempty"`
	ClosingInvoice            bool                      `json:"closingInvoice,omitempty"`
	ClaimedGrossAmount        *float64                  `json:"claimedGrossAmount,omitempty"`
	DownPaymentDeductions     []DownPaymentDeduction    `json:"downPaymentDeductions,omitempty"`
	RecurringTemplateID       *string                   `json:"recurringTemplateId,omitempty"`
	RelatedVouchers           []RelatedVoucher          `json:"relatedVouchers,omitempty"`
	PrintLayoutID             string                    `json:"printLayoutId,omitempty"`
	Title                     string                    `json:"title,omitempty"`
	Introduction              string                    `json:"introduction,omitempty"`
	Remark                    string                    `json:"remark,omitempty"`
	Files                     *Files                    `json:"files,omitempty"`
}

// DownPaymentDeduction represents a down payment deduction.
type DownPaymentDeduction struct {
	ID                  string  `json:"id,omitempty"`
	VoucherType         string  `json:"voucherType,omitempty"`
	Title               string  `json:"title,omitempty"`
	VoucherNumber       string  `json:"voucherNumber,omitempty"`
	VoucherDate         string  `json:"voucherDate,omitempty"`
	ReceivedGrossAmount float64 `json:"receivedGrossAmount,omitempty"`
	ReceivedNetAmount   float64 `json:"receivedNetAmount,omitempty"`
	ReceivedTaxAmount   float64 `json:"receivedTaxAmount,omitempty"`
	TaxRatePercentage   float64 `json:"taxRatePercentage,omitempty"`
}

// InvoiceCreateRequest represents the request body for creating an invoice.
type InvoiceCreateRequest struct {
	VoucherDate        time.Time           `json:"voucherDate"`
	Address            *Address            `json:"address"`
	LineItems          []LineItem          `json:"lineItems"`
	TotalPrice         *TotalPrice         `json:"totalPrice,omitempty"`
	TaxConditions      *TaxConditions      `json:"taxConditions"`
	PaymentConditions  *PaymentConditions  `json:"paymentConditions,omitempty"`
	ShippingConditions *ShippingConditions `json:"shippingConditions,omitempty"`
	PrintLayoutID      string              `json:"printLayoutId,omitempty"`
	Title              string              `json:"title,omitempty"`
	Introduction       string              `json:"introduction,omitempty"`
	Remark             string              `json:"remark,omitempty"`
	Language           string              `json:"language,omitempty"`
	XRechnung          *XRechnung          `json:"xRechnung,omitempty"`
}

// InvoiceFinalizeAction represents the finalization action for an invoice.
type InvoiceFinalizeAction string

const (
	InvoiceFinalizeActionDraft    InvoiceFinalizeAction = ""
	InvoiceFinalizeActionFinalize InvoiceFinalizeAction = "finalize"
)
