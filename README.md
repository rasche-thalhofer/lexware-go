# lexware-go

A Go client library for the [Lexware API](https://api.lexware.io). This library provides a developer-friendly interface to interact with Lexware, the German accounting software.

## Features

- 🎯 **Type-safe API** - Fully typed request and response structures
- 🔧 **Interface-based design** - Follows Go best practices (consume interfaces, return structs)
- 🚀 **Easy initialization** - Simple client creation with just an API key
- 📦 **Kubernetes client-go style** - Familiar patterns for K8s developers
- 🔄 **Context support** - All methods accept `context.Context` for cancellation and timeouts
- ⚡ **Built-in rate limiting** - Global rate limiter (default: 2 req/s) to prevent API throttling

## Installation

```bash
go get github.com/rasche-thalhofer/lexware-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/rasche-thalhofer/lexware-go/lexware"
    "github.com/rasche-thalhofer/lexware-go/types"
)

func main() {
    // Create a new client with your API key
    client, err := lexware.NewClient("your-api-key")
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Get profile information
    profile, err := client.Profile().Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Organization: %s\n", profile.CompanyName)

    // List contacts
    contacts, err := client.Contacts().List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d contacts\n", len(contacts.Content))

    // Create an invoice
    invoice := &types.InvoiceCreateRequest{
        VoucherDate: time.Now(),
        Address: &types.Address{
            Name:        "Example Company",
            Street:      "Main Street 1",
            City:        "Berlin",
            Zip:         "10115",
            CountryCode: "DE",
        },
        LineItems: []types.LineItem{
            {
                Type:     "custom",
                Name:     "Consulting Services",
                Quantity: 10,
                UnitName: "hours",
                UnitPrice: &types.UnitPrice{
                    Currency:          "EUR",
                    NetAmount:         100.00,
                    GrossAmount:       119.00,
                    TaxRatePercentage: 19,
                },
            },
        },
        TaxConditions: &types.TaxConditions{
            TaxType: "net",
        },
    }

    result, err := client.Invoices().Create(ctx, invoice, false) // false = draft
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created invoice with ID: %s\n", result.ID)
}
```

## Client Configuration

### Simple Initialization

```go
client, err := lexware.NewClient("your-api-key")
```

### Custom Configuration

```go
client, err := lexware.NewClientWithConfig(lexware.Config{
    APIKey:  "your-api-key",
    BaseURL: "https://api.lexware.io", // Optional, defaults to production URL
    Timeout: 60 * time.Second,          // Optional, defaults to 30s
})
```

### With Custom HTTP Client

```go
httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns: 10,
    },
}

client, err := lexware.NewClientWithConfig(lexware.Config{
    APIKey:     "your-api-key",
    HTTPClient: httpClient,
})
```

### Rate Limiting

The client includes a global rate limiter that defaults to **2 requests per second**. This helps avoid hitting the Lexware API rate limits.

```go
// Default: 2 requests per second
client, err := lexware.NewClient("your-api-key")

// Custom rate limit: 5 requests per second
client, err := lexware.NewClientWithConfig(lexware.Config{
    APIKey:    "your-api-key",
    RateLimit: 5,
})

// Disable rate limiting (use with caution)
client, err := lexware.NewClientWithConfig(lexware.Config{
    APIKey:    "your-api-key",
    RateLimit: -1,
})
```

The rate limiter uses a token bucket algorithm and respects context cancellation. If a request is cancelled while waiting for rate limit capacity, the error will be returned immediately.

## Available Endpoints

The client provides access to all Lexware API endpoints:

```go
client.Articles()            // Articles management
client.Contacts()            // Contacts management
client.Countries()           // Countries list
client.CreditNotes()         // Credit notes management
client.DeliveryNotes()       // Delivery notes management
client.DownPaymentInvoices() // Down payment invoices (read-only)
client.Dunnings()            // Dunnings (payment reminders)
client.EventSubscriptions()  // Webhook subscriptions
client.Files()               // File upload/download
client.Invoices()            // Invoices management
client.OrderConfirmations()  // Order confirmations
client.Payments()            // Payment information
client.PaymentConditions()   // Payment conditions list
client.PostingCategories()   // Posting categories list
client.PrintLayouts()        // Print layouts list
client.Profile()             // Organization profile
client.Quotations()          // Quotations management
client.RecurringTemplates()  // Recurring templates
client.VoucherList()         // Voucher list with filters
client.Vouchers()            // Bookkeeping vouchers
```

## Examples

### Working with Contacts

```go
// Create a contact
contact := &types.ContactCreateRequest{
    Version: 0,
    Roles: &types.ContactRoles{
        Customer: &types.CustomerRole{},
    },
    Company: &types.Company{
        Name: "ACME Corp",
    },
    Addresses: &types.ContactAddresses{
        Billing: []types.ContactAddress{
            {
                Street:      "Main St 123",
                Zip:         "12345",
                City:        "Berlin",
                CountryCode: "DE",
            },
        },
    },
}

result, err := client.Contacts().Create(ctx, contact)

// Get a contact
contact, err := client.Contacts().Get(ctx, "contact-id")

// List contacts with filters
contacts, err := client.Contacts().List(ctx, 
    &types.ListOptions{Page: 0, Size: 25},
    &types.ContactFilterOptions{Customer: true},
)
```

### Working with Invoices

```go
// Create a draft invoice
result, err := client.Invoices().Create(ctx, invoiceReq, false)

// Create and finalize an invoice
result, err := client.Invoices().Create(ctx, invoiceReq, true)

// Get an invoice
invoice, err := client.Invoices().Get(ctx, "invoice-id")

// Download invoice PDF
pdfReader, err := client.Invoices().DownloadDocument(ctx, "invoice-id")
if err != nil {
    log.Fatal(err)
}
defer pdfReader.Close()

file, _ := os.Create("invoice.pdf")
defer file.Close()
io.Copy(file, pdfReader)
```

### Working with Event Subscriptions (Webhooks)

```go
// Subscribe to invoice events
subscription := &types.EventSubscriptionCreateRequest{
    EventType:   types.EventTypeInvoiceCreated,
    CallbackURL: "https://your-server.com/webhook",
}
result, err := client.EventSubscriptions().Create(ctx, subscription)

// List all subscriptions
subscriptions, err := client.EventSubscriptions().List(ctx)

// Delete a subscription
err := client.EventSubscriptions().Delete(ctx, "subscription-id")
```

### File Upload

```go
file, _ := os.Open("receipt.pdf")
defer file.Close()

result, err := client.Files().Upload(ctx, "receipt.pdf", file, types.FileUploadTypeVoucher)
fmt.Printf("Uploaded file ID: %s\n", result.ID)
```

## Error Handling

The library provides typed errors for better error handling:

```go
result, err := client.Invoices().Get(ctx, "invalid-id")
if err != nil {
    if apiErr, ok := err.(*lexware.APIError); ok {
        switch {
        case apiErr.IsNotFound():
            fmt.Println("Invoice not found")
        case apiErr.IsRateLimited():
            fmt.Println("Rate limited, please retry later")
        case apiErr.IsUnauthorized():
            fmt.Println("Invalid API key")
        default:
            fmt.Printf("API error (status %d): %s\n", apiErr.StatusCode, apiErr.Body)
        }
    }
}
```

## Pagination

Paginated endpoints accept `ListOptions`:

```go
// First page
page1, err := client.Contacts().List(ctx, &types.ListOptions{
    Page: 0,
    Size: 25,
}, nil)

fmt.Printf("Page %d of %d\n", page1.Number+1, page1.TotalPages)
fmt.Printf("Showing %d of %d total contacts\n", len(page1.Content), page1.TotalElements)

// Check if there are more pages
if !page1.Last {
    page2, err := client.Contacts().List(ctx, &types.ListOptions{
        Page: 1,
        Size: 25,
    }, nil)
}
```

## Rate Limiting

The Lexware API allows up to 2 requests per second. The library doesn't automatically handle rate limiting, but provides error detection:

```go
_, err := client.Invoices().Get(ctx, "id")
if apiErr, ok := err.(*lexware.APIError); ok && apiErr.IsRateLimited() {
    time.Sleep(time.Second)
    // Retry...
}
```

## API Documentation

For complete API documentation, refer to:
- [Lexware API Documentation](https://developers.lexware.io/docs/)
- [API Cookbooks (German)](https://developers.lexware.io/cookbooks/)

## License

MIT License - see [LICENSE](LICENSE) file.

