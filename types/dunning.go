// Package types provides type definitions for the Lexware API.
package types

import "time"

// Dunning represents a dunning (payment reminder) in Lexware.
type Dunning struct {
	ID                        string                    `json:"id,omitempty"`
	OrganizationID            string                    `json:"organizationId,omitempty"`
	CreatedDate               time.Time                 `json:"createdDate,omitempty"`
	UpdatedDate               time.Time                 `json:"updatedDate,omitempty"`
	Version                   int                       `json:"version,omitempty"`
	Language                  string                    `json:"language,omitempty"`
	Archived                  bool                      `json:"archived,omitempty"`
	VoucherStatus             VoucherStatus             `json:"voucherStatus,omitempty"`
	VoucherDate               time.Time                 `json:"voucherDate,omitempty"`
	Address                   *Address                  `json:"address,omitempty"`
	ElectronicDocumentProfile ElectronicDocumentProfile `json:"electronicDocumentProfile,omitempty"`
	LineItems                 []LineItem                `json:"lineItems,omitempty"`
	TotalPrice                *TotalPrice               `json:"totalPrice,omitempty"`
	TaxAmounts                []TaxAmount               `json:"taxAmounts,omitempty"`
	TaxConditions             *TaxConditions            `json:"taxConditions,omitempty"`
	ShippingConditions        *ShippingConditions       `json:"shippingConditions,omitempty"`
	RelatedVouchers           []RelatedVoucher          `json:"relatedVouchers,omitempty"`
	PrintLayoutID             string                    `json:"printLayoutId,omitempty"`
	Title                     string                    `json:"title,omitempty"`
	Introduction              string                    `json:"introduction,omitempty"`
	Remark                    string                    `json:"remark,omitempty"`
	Files                     *Files                    `json:"files,omitempty"`
}

// DunningCreateRequest represents the request body for creating a dunning.
type DunningCreateRequest struct {
	VoucherDate        time.Time           `json:"voucherDate"`
	Address            *Address            `json:"address"`
	LineItems          []LineItem          `json:"lineItems"`
	TaxConditions      *TaxConditions      `json:"taxConditions"`
	ShippingConditions *ShippingConditions `json:"shippingConditions,omitempty"`
	PrintLayoutID      string              `json:"printLayoutId,omitempty"`
	Title              string              `json:"title,omitempty"`
	Introduction       string              `json:"introduction,omitempty"`
	Remark             string              `json:"remark,omitempty"`
	Language           string              `json:"language,omitempty"`
}
