// Package types provides type definitions for the Lexware API.
package types

import "time"

// EventSubscription represents an event subscription in Lexware.
type EventSubscription struct {
	SubscriptionID string    `json:"subscriptionId,omitempty"`
	OrganizationID string    `json:"organizationId,omitempty"`
	CreatedDate    time.Time `json:"createdDate,omitempty"`
	EventType      string    `json:"eventType,omitempty"`
	CallbackURL    string    `json:"callbackUrl,omitempty"`
}

// EventSubscriptionCreateRequest represents the request body for creating an event subscription.
type EventSubscriptionCreateRequest struct {
	EventType   string `json:"eventType"`
	CallbackURL string `json:"callbackUrl"`
}

// WebhookPayload represents the payload of a webhook callback.
type WebhookPayload struct {
	OrganizationID string    `json:"organizationId,omitempty"`
	EventType      string    `json:"eventType,omitempty"`
	ResourceID     string    `json:"resourceId,omitempty"`
	EventDate      time.Time `json:"eventDate,omitempty"`
}

// EventType constants for webhook events.
const (
	// Article events
	EventTypeArticleCreated = "article.created"
	EventTypeArticleChanged = "article.changed"
	EventTypeArticleDeleted = "article.deleted"

	// Contact events
	EventTypeContactCreated = "contact.created"
	EventTypeContactChanged = "contact.changed"
	EventTypeContactDeleted = "contact.deleted"

	// Credit note events
	EventTypeCreditNoteCreated       = "credit-note.created"
	EventTypeCreditNoteChanged       = "credit-note.changed"
	EventTypeCreditNoteDeleted       = "credit-note.deleted"
	EventTypeCreditNoteStatusChanged = "credit-note.status.changed"

	// Delivery note events
	EventTypeDeliveryNoteCreated       = "delivery-note.created"
	EventTypeDeliveryNoteChanged       = "delivery-note.changed"
	EventTypeDeliveryNoteDeleted       = "delivery-note.deleted"
	EventTypeDeliveryNoteStatusChanged = "delivery-note.status.changed"

	// Dunning events
	EventTypeDunningCreated = "dunning.created"
	EventTypeDunningChanged = "dunning.changed"
	EventTypeDunningDeleted = "dunning.deleted"

	// Down payment invoice events
	EventTypeDownPaymentInvoiceCreated       = "down-payment-invoice.created"
	EventTypeDownPaymentInvoiceChanged       = "down-payment-invoice.changed"
	EventTypeDownPaymentInvoiceDeleted       = "down-payment-invoice.deleted"
	EventTypeDownPaymentInvoiceStatusChanged = "down-payment-invoice.status.changed"

	// Invoice events
	EventTypeInvoiceCreated       = "invoice.created"
	EventTypeInvoiceChanged       = "invoice.changed"
	EventTypeInvoiceDeleted       = "invoice.deleted"
	EventTypeInvoiceStatusChanged = "invoice.status.changed"

	// Order confirmation events
	EventTypeOrderConfirmationCreated       = "order-confirmation.created"
	EventTypeOrderConfirmationChanged       = "order-confirmation.changed"
	EventTypeOrderConfirmationDeleted       = "order-confirmation.deleted"
	EventTypeOrderConfirmationStatusChanged = "order-confirmation.status.changed"

	// Payment events
	EventTypePaymentChanged = "payment.changed"

	// Quotation events
	EventTypeQuotationCreated       = "quotation.created"
	EventTypeQuotationChanged       = "quotation.changed"
	EventTypeQuotationDeleted       = "quotation.deleted"
	EventTypeQuotationStatusChanged = "quotation.status.changed"

	// Recurring template events
	EventTypeRecurringTemplateCreated = "recurring-template.created"
	EventTypeRecurringTemplateChanged = "recurring-template.changed"
	EventTypeRecurringTemplateDeleted = "recurring-template.deleted"

	// Token events
	EventTypeTokenRevoked = "token.revoked"

	// Voucher events
	EventTypeVoucherCreated       = "voucher.created"
	EventTypeVoucherChanged       = "voucher.changed"
	EventTypeVoucherDeleted       = "voucher.deleted"
	EventTypeVoucherStatusChanged = "voucher.status.changed"
)
