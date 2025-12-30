// Package types provides type definitions for the Lexware API.
package types

import "time"

// Article represents an article in Lexware.
type Article struct {
	ID             string        `json:"id,omitempty"`
	OrganizationID string        `json:"organizationId,omitempty"`
	CreatedDate    time.Time     `json:"createdDate,omitempty"`
	UpdatedDate    time.Time     `json:"updatedDate,omitempty"`
	Archived       bool          `json:"archived,omitempty"`
	Title          string        `json:"title,omitempty"`
	Description    string        `json:"description,omitempty"`
	Type           ArticleType   `json:"type,omitempty"`
	ArticleNumber  string        `json:"articleNumber,omitempty"`
	GTIN           string        `json:"gtin,omitempty"`
	Note           string        `json:"note,omitempty"`
	UnitName       string        `json:"unitName,omitempty"`
	Price          *ArticlePrice `json:"price,omitempty"`
	Version        int           `json:"version,omitempty"`
}

// ArticlePrice represents the price of an article.
type ArticlePrice struct {
	NetPrice     float64 `json:"netPrice,omitempty"`
	GrossPrice   float64 `json:"grossPrice,omitempty"`
	LeadingPrice string  `json:"leadingPrice,omitempty"`
	TaxRate      float64 `json:"taxRate,omitempty"`
}

// ArticleType represents the type of an article.
type ArticleType string

const (
	ArticleTypeProduct ArticleType = "PRODUCT"
	ArticleTypeService ArticleType = "SERVICE"
)

// ArticleCreateRequest represents the request body for creating an article.
type ArticleCreateRequest struct {
	Title         string        `json:"title"`
	Description   string        `json:"description,omitempty"`
	Type          ArticleType   `json:"type"`
	ArticleNumber string        `json:"articleNumber,omitempty"`
	GTIN          string        `json:"gtin,omitempty"`
	Note          string        `json:"note,omitempty"`
	UnitName      string        `json:"unitName,omitempty"`
	Price         *ArticlePrice `json:"price,omitempty"`
}

// ArticleUpdateRequest represents the request body for updating an article.
type ArticleUpdateRequest struct {
	Title         string        `json:"title"`
	Description   string        `json:"description,omitempty"`
	Type          ArticleType   `json:"type"`
	ArticleNumber string        `json:"articleNumber,omitempty"`
	GTIN          string        `json:"gtin,omitempty"`
	Note          string        `json:"note,omitempty"`
	UnitName      string        `json:"unitName,omitempty"`
	Price         *ArticlePrice `json:"price,omitempty"`
	Version       int           `json:"version"`
}

// ArticleFilterOptions specifies the optional parameters for filtering articles.
type ArticleFilterOptions struct {
	ArticleNumber string
	GTIN          string
	Type          ArticleType
}
