// Package types provides type definitions for the Lexware API.
package types

// Contact represents a contact in Lexware.
type Contact struct {
	ID             string            `json:"id,omitempty"`
	OrganizationID string            `json:"organizationId,omitempty"`
	Version        int               `json:"version,omitempty"`
	Roles          *ContactRoles     `json:"roles,omitempty"`
	Company        *Company          `json:"company,omitempty"`
	Person         *Person           `json:"person,omitempty"`
	Addresses      *ContactAddresses `json:"addresses,omitempty"`
	XRechnung      *XRechnungContact `json:"xRechnung,omitempty"`
	EmailAddresses *EmailAddresses   `json:"emailAddresses,omitempty"`
	PhoneNumbers   *PhoneNumbers     `json:"phoneNumbers,omitempty"`
	Note           string            `json:"note,omitempty"`
	Archived       bool              `json:"archived,omitempty"`
}

// ContactRoles represents the roles of a contact.
type ContactRoles struct {
	Customer *CustomerRole `json:"customer,omitempty"`
	Vendor   *VendorRole   `json:"vendor,omitempty"`
}

// CustomerRole represents the customer role.
type CustomerRole struct {
	Number int `json:"number,omitempty"`
}

// VendorRole represents the vendor role.
type VendorRole struct {
	Number int `json:"number,omitempty"`
}

// Company represents company information.
type Company struct {
	Name                 string          `json:"name,omitempty"`
	TaxNumber            string          `json:"taxNumber,omitempty"`
	VATRegistrationID    string          `json:"vatRegistrationId,omitempty"`
	AllowTaxFreeInvoices bool            `json:"allowTaxFreeInvoices,omitempty"`
	ContactPersons       []ContactPerson `json:"contactPersons,omitempty"`
}

// Person represents a person contact.
type Person struct {
	Salutation string `json:"salutation,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
}

// ContactPerson represents a contact person within a company.
type ContactPerson struct {
	Salutation   string `json:"salutation,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	Primary      bool   `json:"primary,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
}

// ContactAddresses represents the addresses of a contact.
type ContactAddresses struct {
	Billing  []ContactAddress `json:"billing,omitempty"`
	Shipping []ContactAddress `json:"shipping,omitempty"`
}

// ContactAddress represents a contact address.
type ContactAddress struct {
	Supplement  string `json:"supplement,omitempty"`
	Street      string `json:"street,omitempty"`
	Zip         string `json:"zip,omitempty"`
	City        string `json:"city,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

// XRechnungContact represents XRechnung information for a contact.
type XRechnungContact struct {
	BuyerReference         string `json:"buyerReference,omitempty"`
	VendorNumberAtCustomer string `json:"vendorNumberAtCustomer,omitempty"`
}

// EmailAddresses represents email addresses.
type EmailAddresses struct {
	Business []string `json:"business,omitempty"`
	Office   []string `json:"office,omitempty"`
	Private  []string `json:"private,omitempty"`
	Other    []string `json:"other,omitempty"`
}

// PhoneNumbers represents phone numbers.
type PhoneNumbers struct {
	Business []string `json:"business,omitempty"`
	Office   []string `json:"office,omitempty"`
	Mobile   []string `json:"mobile,omitempty"`
	Private  []string `json:"private,omitempty"`
	Fax      []string `json:"fax,omitempty"`
	Other    []string `json:"other,omitempty"`
}

// ContactCreateRequest represents the request body for creating a contact.
type ContactCreateRequest struct {
	Version        int               `json:"version"`
	Roles          *ContactRoles     `json:"roles,omitempty"`
	Company        *Company          `json:"company,omitempty"`
	Person         *Person           `json:"person,omitempty"`
	Addresses      *ContactAddresses `json:"addresses,omitempty"`
	XRechnung      *XRechnungContact `json:"xRechnung,omitempty"`
	EmailAddresses *EmailAddresses   `json:"emailAddresses,omitempty"`
	PhoneNumbers   *PhoneNumbers     `json:"phoneNumbers,omitempty"`
	Note           string            `json:"note,omitempty"`
}

// ContactUpdateRequest represents the request body for updating a contact.
type ContactUpdateRequest struct {
	Version        int               `json:"version"`
	Roles          *ContactRoles     `json:"roles,omitempty"`
	Company        *Company          `json:"company,omitempty"`
	Person         *Person           `json:"person,omitempty"`
	Addresses      *ContactAddresses `json:"addresses,omitempty"`
	XRechnung      *XRechnungContact `json:"xRechnung,omitempty"`
	EmailAddresses *EmailAddresses   `json:"emailAddresses,omitempty"`
	PhoneNumbers   *PhoneNumbers     `json:"phoneNumbers,omitempty"`
	Note           string            `json:"note,omitempty"`
}

// ContactFilterOptions specifies the optional parameters for filtering contacts.
type ContactFilterOptions struct {
	Email    string
	Name     string
	Number   int
	Customer bool
	Vendor   bool
}
