// Package types provides type definitions for the Lexware API.
package types

// Country represents a country in Lexware.
type Country struct {
	CountryCode       string `json:"countryCode,omitempty"`
	CountryNameDE     string `json:"countryNameDE,omitempty"`
	CountryNameEN     string `json:"countryNameEN,omitempty"`
	TaxClassification string `json:"taxClassification,omitempty"`
}

// TaxClassificationType represents country tax classification.
type TaxClassificationType string

const (
	TaxClassificationDomestic          TaxClassificationType = "domestic"
	TaxClassificationIntraCommunity    TaxClassificationType = "intraCommunity"
	TaxClassificationThirdPartyCountry TaxClassificationType = "thirdPartyCountry"
)

// Profile represents the user/organization profile in Lexware.
type Profile struct {
	OrganizationID         string       `json:"organizationId,omitempty"`
	CompanyName            string       `json:"companyName,omitempty"`
	Created                *LexwareDate `json:"created,omitempty"`
	ConnectionID           string       `json:"connectionId,omitempty"`
	TaxType                string       `json:"taxType,omitempty"`
	SmallBusiness          bool         `json:"smallBusiness,omitempty"`
	DistanceSalesPrinciple string       `json:"distanceSalesPrinciple,omitempty"`
}

// LexwareDate represents a date in Lexware format.
type LexwareDate struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

// Payment represents payment information for a voucher.
type Payment struct {
	OpenAmount    float64       `json:"openAmount,omitempty"`
	Currency      string        `json:"currency,omitempty"`
	PaymentStatus string        `json:"paymentStatus,omitempty"`
	VoucherType   string        `json:"voucherType,omitempty"`
	VoucherID     string        `json:"voucherId,omitempty"`
	VoucherNumber string        `json:"voucherNumber,omitempty"`
	VoucherDate   string        `json:"voucherDate,omitempty"`
	PaymentItems  []PaymentItem `json:"paymentItems,omitempty"`
}

// PaymentItem represents a payment item.
type PaymentItem struct {
	PaymentItemType string  `json:"paymentItemType,omitempty"`
	PostingDate     string  `json:"postingDate,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	Currency        string  `json:"currency,omitempty"`
}

// PaymentCondition represents a payment condition.
type PaymentCondition struct {
	ID                        string                     `json:"id,omitempty"`
	OrganizationID            string                     `json:"organizationId,omitempty"`
	PaymentTermLabelTemplate  string                     `json:"paymentTermLabelTemplate,omitempty"`
	PaymentTermDuration       int                        `json:"paymentTermDuration,omitempty"`
	PaymentDiscountConditions *PaymentDiscountConditions `json:"paymentDiscountConditions,omitempty"`
}

// PostingCategory represents a posting category.
type PostingCategory struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Type            string `json:"type,omitempty"`
	ContactRequired bool   `json:"contactRequired,omitempty"`
	SplitAllowed    bool   `json:"splitAllowed,omitempty"`
	GroupName       string `json:"groupName,omitempty"`
}

// PrintLayout represents a print layout.
type PrintLayout struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IsDefault bool   `json:"default,omitempty"`
}

// OrderConfirmation represents an order confirmation in Lexware.
type OrderConfirmation struct {
	ID                        string                    `json:"id,omitempty"`
	OrganizationID            string                    `json:"organizationId,omitempty"`
	CreatedDate               string                    `json:"createdDate,omitempty"`
	UpdatedDate               string                    `json:"updatedDate,omitempty"`
	Version                   int                       `json:"version,omitempty"`
	Language                  string                    `json:"language,omitempty"`
	Archived                  bool                      `json:"archived,omitempty"`
	VoucherStatus             VoucherStatus             `json:"voucherStatus,omitempty"`
	VoucherNumber             string                    `json:"voucherNumber,omitempty"`
	VoucherDate               string                    `json:"voucherDate,omitempty"`
	Address                   *Address                  `json:"address,omitempty"`
	ElectronicDocumentProfile ElectronicDocumentProfile `json:"electronicDocumentProfile,omitempty"`
	LineItems                 []LineItem                `json:"lineItems,omitempty"`
	TotalPrice                *TotalPrice               `json:"totalPrice,omitempty"`
	TaxAmounts                []TaxAmount               `json:"taxAmounts,omitempty"`
	TaxConditions             *TaxConditions            `json:"taxConditions,omitempty"`
	PaymentConditions         *PaymentConditions        `json:"paymentConditions,omitempty"`
	ShippingConditions        *ShippingConditions       `json:"shippingConditions,omitempty"`
	DeliveryTerms             string                    `json:"deliveryTerms,omitempty"`
	RelatedVouchers           []RelatedVoucher          `json:"relatedVouchers,omitempty"`
	PrintLayoutID             string                    `json:"printLayoutId,omitempty"`
	Title                     string                    `json:"title,omitempty"`
	Introduction              string                    `json:"introduction,omitempty"`
	Remark                    string                    `json:"remark,omitempty"`
	Files                     *Files                    `json:"files,omitempty"`
}

// DownPaymentInvoice represents a down payment invoice in Lexware.
type DownPaymentInvoice struct {
	ID                        string                    `json:"id,omitempty"`
	OrganizationID            string                    `json:"organizationId,omitempty"`
	CreatedDate               string                    `json:"createdDate,omitempty"`
	UpdatedDate               string                    `json:"updatedDate,omitempty"`
	Version                   int                       `json:"version,omitempty"`
	Language                  string                    `json:"language,omitempty"`
	Archived                  bool                      `json:"archived,omitempty"`
	VoucherStatus             VoucherStatus             `json:"voucherStatus,omitempty"`
	VoucherNumber             string                    `json:"voucherNumber,omitempty"`
	VoucherDate               string                    `json:"voucherDate,omitempty"`
	DueDate                   string                    `json:"dueDate,omitempty"`
	Address                   *Address                  `json:"address,omitempty"`
	ElectronicDocumentProfile ElectronicDocumentProfile `json:"electronicDocumentProfile,omitempty"`
	LineItems                 []LineItem                `json:"lineItems,omitempty"`
	TotalPrice                *TotalPrice               `json:"totalPrice,omitempty"`
	TaxAmounts                []TaxAmount               `json:"taxAmounts,omitempty"`
	TaxConditions             *TaxConditions            `json:"taxConditions,omitempty"`
	PaymentConditions         *PaymentConditions        `json:"paymentConditions,omitempty"`
	ShippingConditions        *ShippingConditions       `json:"shippingConditions,omitempty"`
	ClosingInvoiceID          *string                   `json:"closingInvoiceId,omitempty"`
	RelatedVouchers           []RelatedVoucher          `json:"relatedVouchers,omitempty"`
	PrintLayoutID             string                    `json:"printLayoutId,omitempty"`
	Title                     string                    `json:"title,omitempty"`
	Introduction              string                    `json:"introduction,omitempty"`
	Remark                    string                    `json:"remark,omitempty"`
	Files                     *Files                    `json:"files,omitempty"`
}

// RecurringTemplate represents a recurring template in Lexware.
type RecurringTemplate struct {
	ID                        string                     `json:"id,omitempty"`
	OrganizationID            string                     `json:"organizationId,omitempty"`
	CreatedDate               string                     `json:"createdDate,omitempty"`
	UpdatedDate               string                     `json:"updatedDate,omitempty"`
	Version                   int                        `json:"version,omitempty"`
	Language                  string                     `json:"language,omitempty"`
	Archived                  bool                       `json:"archived,omitempty"`
	Address                   *Address                   `json:"address,omitempty"`
	LineItems                 []LineItem                 `json:"lineItems,omitempty"`
	TotalPrice                *TotalPrice                `json:"totalPrice,omitempty"`
	TaxAmounts                []TaxAmount                `json:"taxAmounts,omitempty"`
	TaxConditions             *TaxConditions             `json:"taxConditions,omitempty"`
	PaymentConditions         *PaymentConditions         `json:"paymentConditions,omitempty"`
	ShippingConditions        *ShippingConditions        `json:"shippingConditions,omitempty"`
	PrintLayoutID             string                     `json:"printLayoutId,omitempty"`
	Title                     string                     `json:"title,omitempty"`
	Introduction              string                     `json:"introduction,omitempty"`
	Remark                    string                     `json:"remark,omitempty"`
	RecurringTemplateSettings *RecurringTemplateSettings `json:"recurringTemplateSettings,omitempty"`
}

// RecurringTemplateSettings represents settings for a recurring template.
type RecurringTemplateSettings struct {
	ID                        string `json:"id,omitempty"`
	StartDate                 string `json:"startDate,omitempty"`
	EndDate                   string `json:"endDate,omitempty"`
	Finalize                  bool   `json:"finalize,omitempty"`
	ShippingType              string `json:"shippingType,omitempty"`
	ExecutionInterval         string `json:"executionInterval,omitempty"`
	NextExecutionDate         string `json:"nextExecutionDate,omitempty"`
	LastExecutionFailed       bool   `json:"lastExecutionFailed,omitempty"`
	LastExecutionErrorMessage string `json:"lastExecutionErrorMessage,omitempty"`
}

// VoucherListItem represents an item in the voucherlist.
type VoucherListItem struct {
	ID            string        `json:"id,omitempty"`
	VoucherType   VoucherType   `json:"voucherType,omitempty"`
	VoucherStatus VoucherStatus `json:"voucherStatus,omitempty"`
	VoucherNumber string        `json:"voucherNumber,omitempty"`
	VoucherDate   string        `json:"voucherDate,omitempty"`
	UpdatedDate   string        `json:"updatedDate,omitempty"`
	DueDate       string        `json:"dueDate,omitempty"`
	ContactID     string        `json:"contactId,omitempty"`
	ContactName   string        `json:"contactName,omitempty"`
	TotalAmount   float64       `json:"totalAmount,omitempty"`
	OpenAmount    float64       `json:"openAmount,omitempty"`
	Currency      string        `json:"currency,omitempty"`
	Archived      bool          `json:"archived,omitempty"`
}

// VoucherListFilterOptions specifies the optional parameters for filtering the voucherlist.
type VoucherListFilterOptions struct {
	VoucherType     VoucherType
	VoucherStatus   VoucherStatus
	Archived        *bool
	ContactID       string
	VoucherDateFrom string
	VoucherDateTo   string
	CreatedDateFrom string
	CreatedDateTo   string
	UpdatedDateFrom string
	UpdatedDateTo   string
}

// FileUploadType represents the type of file upload.
type FileUploadType string

const (
	FileUploadTypeVoucher FileUploadType = "voucher"
)

// FileUploadResponse represents the response from a file upload.
type FileUploadResponse struct {
	ID string `json:"id,omitempty"`
}
