// Package types provides type definitions for the Lexware API.
package types

import "time"

// Common types used across all endpoints

// Address represents an address in Lexware.
type Address struct {
	ContactID     *string `json:"contactId,omitempty"`
	Name          string  `json:"name,omitempty"`
	Supplement    string  `json:"supplement,omitempty"`
	Street        string  `json:"street,omitempty"`
	City          string  `json:"city,omitempty"`
	Zip           string  `json:"zip,omitempty"`
	CountryCode   string  `json:"countryCode,omitempty"`
	ContactPerson string  `json:"contactPerson,omitempty"`
}

// UnitPrice represents the unit price of an item.
type UnitPrice struct {
	Currency          string  `json:"currency,omitempty"`
	NetAmount         float64 `json:"netAmount,omitempty"`
	GrossAmount       float64 `json:"grossAmount,omitempty"`
	TaxRatePercentage float64 `json:"taxRatePercentage,omitempty"`
}

// TotalPrice represents the total price of a voucher.
type TotalPrice struct {
	Currency                string   `json:"currency,omitempty"`
	TotalNetAmount          float64  `json:"totalNetAmount,omitempty"`
	TotalGrossAmount        float64  `json:"totalGrossAmount,omitempty"`
	TotalTaxAmount          float64  `json:"totalTaxAmount,omitempty"`
	TotalDiscountAbsolute   *float64 `json:"totalDiscountAbsolute,omitempty"`
	TotalDiscountPercentage *float64 `json:"totalDiscountPercentage,omitempty"`
}

// TaxAmount represents the tax amount for a specific tax rate.
type TaxAmount struct {
	TaxRatePercentage float64 `json:"taxRatePercentage,omitempty"`
	TaxAmount         float64 `json:"taxAmount,omitempty"`
	NetAmount         float64 `json:"netAmount,omitempty"`
}

// TaxConditions represents the tax conditions of a voucher.
type TaxConditions struct {
	TaxType     string  `json:"taxType,omitempty"`
	TaxTypeNote *string `json:"taxTypeNote,omitempty"`
}

// PaymentConditions represents the payment conditions of a voucher.
type PaymentConditions struct {
	PaymentTermLabel          string                     `json:"paymentTermLabel,omitempty"`
	PaymentTermLabelTemplate  string                     `json:"paymentTermLabelTemplate,omitempty"`
	PaymentTermDuration       int                        `json:"paymentTermDuration,omitempty"`
	PaymentDiscountConditions *PaymentDiscountConditions `json:"paymentDiscountConditions,omitempty"`
}

// PaymentDiscountConditions represents payment discount conditions.
type PaymentDiscountConditions struct {
	DiscountPercentage float64 `json:"discountPercentage,omitempty"`
	DiscountRange      int     `json:"discountRange,omitempty"`
}

// ShippingConditions represents the shipping conditions of a voucher.
type ShippingConditions struct {
	ShippingDate    *time.Time `json:"shippingDate,omitempty"`
	ShippingEndDate *time.Time `json:"shippingEndDate,omitempty"`
	ShippingType    string     `json:"shippingType,omitempty"`
}

// RelatedVoucher represents a related voucher.
type RelatedVoucher struct {
	ID            string `json:"id,omitempty"`
	VoucherNumber string `json:"voucherNumber,omitempty"`
	VoucherType   string `json:"voucherType,omitempty"`
}

// Files represents the file references of a voucher.
type Files struct {
	DocumentFileID string `json:"documentFileId,omitempty"`
}

// LineItem represents a line item in a voucher.
type LineItem struct {
	ID                 *string    `json:"id,omitempty"`
	Type               string     `json:"type,omitempty"`
	Name               string     `json:"name,omitempty"`
	Description        string     `json:"description,omitempty"`
	Quantity           float64    `json:"quantity,omitempty"`
	UnitName           string     `json:"unitName,omitempty"`
	UnitPrice          *UnitPrice `json:"unitPrice,omitempty"`
	DiscountPercentage float64    `json:"discountPercentage,omitempty"`
	LineItemAmount     float64    `json:"lineItemAmount,omitempty"`
}

// XRechnung represents XRechnung related properties.
type XRechnung struct {
	BuyerReference string `json:"buyerReference,omitempty"`
}

// VoucherStatus represents the status of a voucher.
type VoucherStatus string

const (
	VoucherStatusDraft   VoucherStatus = "draft"
	VoucherStatusOpen    VoucherStatus = "open"
	VoucherStatusPaid    VoucherStatus = "paid"
	VoucherStatusPaidOff VoucherStatus = "paidoff"
	VoucherStatusVoided  VoucherStatus = "voided"
)

// TaxType represents the tax type of a voucher.
type TaxType string

const (
	TaxTypeNet                       TaxType = "net"
	TaxTypeGross                     TaxType = "gross"
	TaxTypeVatFree                   TaxType = "vatfree"
	TaxTypeIntraCommunitySupply      TaxType = "intraCommunitySupply"
	TaxTypeConstructionalServices    TaxType = "constructionalServices"
	TaxTypeExternalServices          TaxType = "externalServices"
	TaxTypeThirdPartyCountryService  TaxType = "thirdPartyCountryService"
	TaxTypeThirdPartyCountryDelivery TaxType = "thirdPartyCountryDelivery"
)

// LineItemType represents the type of a line item.
type LineItemType string

const (
	LineItemTypeService  LineItemType = "service"
	LineItemTypeMaterial LineItemType = "material"
	LineItemTypeCustom   LineItemType = "custom"
	LineItemTypeText     LineItemType = "text"
)

// ShippingType represents the shipping type.
type ShippingType string

const (
	ShippingTypeDelivery       ShippingType = "delivery"
	ShippingTypeDeliveryPeriod ShippingType = "deliveryperiod"
	ShippingTypeService        ShippingType = "service"
	ShippingTypeServicePeriod  ShippingType = "serviceperiod"
	ShippingTypeNone           ShippingType = "none"
)

// ElectronicDocumentProfile represents the electronic document profile.
type ElectronicDocumentProfile string

const (
	ElectronicDocumentProfileNone      ElectronicDocumentProfile = "NONE"
	ElectronicDocumentProfileEN16931   ElectronicDocumentProfile = "EN16931"
	ElectronicDocumentProfileXRechnung ElectronicDocumentProfile = "XRechnung"
)

// Page represents a paginated response.
type Page[T any] struct {
	Content          []T  `json:"content"`
	First            bool `json:"first"`
	Last             bool `json:"last"`
	TotalPages       int  `json:"totalPages"`
	TotalElements    int  `json:"totalElements"`
	NumberOfElements int  `json:"numberOfElements"`
	Size             int  `json:"size"`
	Number           int  `json:"number"`
}

// ActionResult represents the result of a create/update action.
type ActionResult struct {
	ID          string    `json:"id"`
	ResourceURI string    `json:"resourceUri"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Version     int       `json:"version"`
}

// ListOptions specifies the optional parameters for listing resources.
type ListOptions struct {
	Page int
	Size int
}
