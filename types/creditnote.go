// Package types provides type definitions for the Lexware API.
package types

import "time"

// CreditNote represents a credit note in Lexware.
type CreditNote struct {
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
	Address                   *Address                  `json:"address,omitempty"`
	ElectronicDocumentProfile ElectronicDocumentProfile `json:"electronicDocumentProfile,omitempty"`
	LineItems                 []LineItem                `json:"lineItems,omitempty"`
	TotalPrice                *TotalPrice               `json:"totalPrice,omitempty"`
	TaxAmounts                []TaxAmount               `json:"taxAmounts,omitempty"`
	TaxConditions             *TaxConditions            `json:"taxConditions,omitempty"`
	RelatedVouchers           []RelatedVoucher          `json:"relatedVouchers,omitempty"`
	PrintLayoutID             string                    `json:"printLayoutId,omitempty"`
	Title                     string                    `json:"title,omitempty"`
	Introduction              string                    `json:"introduction,omitempty"`
	Remark                    string                    `json:"remark,omitempty"`
	Files                     *Files                    `json:"files,omitempty"`
}

// CreditNoteCreateRequest represents the request body for creating a credit note.
type CreditNoteCreateRequest struct {
	VoucherDate   time.Time      `json:"voucherDate"`
	Address       *Address       `json:"address"`
	LineItems     []LineItem     `json:"lineItems"`
	TotalPrice    *TotalPrice    `json:"totalPrice,omitempty"`
	TaxConditions *TaxConditions `json:"taxConditions"`
	PrintLayoutID string         `json:"printLayoutId,omitempty"`
	Title         string         `json:"title,omitempty"`
	Introduction  string         `json:"introduction,omitempty"`
	Remark        string         `json:"remark,omitempty"`
	Language      string         `json:"language,omitempty"`
}
