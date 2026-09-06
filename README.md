# MailAfrica Go SDK

A professional, open-source Go SDK for the [MailAfrica](https://mailafrica.online) user-facing API.

- **Module**: `github.com/MailAfrica/go-sdk`
- **Go version**: 1.25+
- **License**: MIT

## Installation

```bash
go get github.com/MailAfrica/go-sdk
```

## Quickstart

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MailAfrica/go-sdk"
)

func main() {
	ctx := context.Background()
	client := mailafrica.New(mailafrica.Config{
		BaseURL: "https://api.mailafrica.online",
		APIKey:  "MAIL_test_...",
	})

	resp, err := client.Register(ctx, mailafrica.RegisterRequest{
		Email:    strPtr("user@example.com"),
		Password: "secure-password",
		Name:     "Jane Doe",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Registered:", resp.User.ID)
}

func strPtr(s string) *string { return &s }
```

## Authentication

The SDK supports two authentication methods. Set **one** of them in `Config`:

| Field | Header sent | When to use |
|-------|-------------|-------------|
| `APIKey` | `X-API-Key: <key>` | Preferred for server-side and CLI usage |
| `JWT` | `Authorization: Bearer <jwt>` | For browser/SPA flows or when you already have a JWT |

```go
// API key (preferred)
client := mailafrica.New(mailafrica.Config{
    APIKey: "MAIL_...",
})

// JWT bearer
client := mailafrica.New(mailafrica.Config{
    JWT: "eyJhbGciOi...",
})
```

### Token Refresh

If you provide a `TokenRefresher` callback, the SDK will automatically call it on 401 responses, update the stored JWT, and retry the request once:

```go
client := mailafrica.New(mailafrica.Config{
    JWT:            initialToken,
    TokenRefresher: func(ctx context.Context) (string, error) {
        // Call your own /auth/refresh endpoint or token exchange logic
        newToken, err := refreshToken(ctx)
        return newToken, err
    },
})
```

The SDK does **not** store refresh tokens — the caller manages them.

## Error Handling

All API errors are returned as `*mailafrica.APIError`:

```go
msg, err := client.SendEmail(ctx, req)
if err != nil {
    var apiErr *mailafrica.APIError
    if errors.As(err, &apiErr) {
        fmt.Println("code:", apiErr.Code)
        fmt.Println("status:", apiErr.HTTPStatus)
        fmt.Println("request_id:", apiErr.RequestID)
    }

    if mailafrica.IsInsufficientBalance(err) {
        // handle low balance
    }
    if mailafrica.IsRateLimited(err) {
        // back off and retry
    }
    if mailafrica.IsNotFound(err) {
        // resource missing
    }
}
```

### Sentinel Errors

| Function | Backend Code |
|----------|-------------|
| `IsInsufficientBalance(err)` | `INSUFFICIENT_BALANCE` |
| `IsRateLimited(err)` | `RATE_LIMITED` |
| `IsNotVerified(err)` | `NOT_VERIFIED` |
| `IsAccountDisabled(err)` | `ACCOUNT_DISABLED` |
| `IsNotFound(err)` | `NOT_FOUND` |

## Resources

### Auth

```go
// Register
resp, err := client.Register(ctx, mailafrica.RegisterRequest{
    Email:    strPtr("user@example.com"),
    Password: "password",
    Name:     "Jane",
})

// Login
resp, err := client.Login(ctx, mailafrica.LoginRequest{
    Identifier: "user@example.com",
    Password:   "password",
})

// Google login
resp, err := client.GoogleLogin(ctx, "google-id-token")

// Refresh token
resp, err := client.Refresh(ctx, "refresh-token")

// Current user
user, err := client.Me(ctx)

// Update profile
user, err := client.UpdateMe(ctx, mailafrica.UpdateProfileRequest{
    Name: strPtr("Jane Doe"),
})

// Email & phone
client.SetEmail(ctx, "new@example.com")
client.VerifyEmail(ctx, "token")
client.ResendEmailVerification(ctx)
client.SetPhone(ctx, "+255712345678")
client.VerifyPhone(ctx, "123456")
client.ResendPhoneOTP(ctx)
```

### Inbound

```go
// Addresses
addr, err := client.CreateAddress(ctx, mailafrica.CreateAddressRequest{
    LocalPart: "hello",
})
addresses, err := client.ListAddresses(ctx)
client.DeleteAddress(ctx, addr.ID)

// Messages
msgs, pagination, err := client.ListMessages(ctx, mailafrica.MessageListOpts{
    ListOpts: mailafrica.ListOpts{Page: 1, PerPage: 25},
    AddressID: strInt64Ptr(1),
    Unread:    boolPtr(true),
})
msg, err := client.GetMessage(ctx, msgID)
client.MarkMessageRead(ctx, msgID)

// Inbound domains
domain, err := client.CreateInboundDomain(ctx, "example.com")
domains, err := client.ListInboundDomains(ctx)
client.VerifyInboundDomain(ctx, domain.ID)
client.DeleteInboundDomain(ctx, domain.ID)
```

### Outbound

```go
// Send single email
msg, err := client.SendEmail(ctx, mailafrica.SendEmailRequest{
    To:       []string{"recipient@example.com"},
    Subject:  "Hello",
    HTMLBody: "<p>Hello!</p>",
    TextBody: "Hello!",
})

// Batch send
result, err := client.BatchSend(ctx, mailafrica.BatchSendRequest{
    To:       []string{"a@example.com", "b@example.com"},
    Subject:  "Batch",
    HTMLBody: "<p>Hello!</p>",
})

// Sent emails
emails, pagination, err := client.ListSentEmails(ctx, mailafrica.ListOpts{PerPage: 25})
detail, err := client.GetSentEmail(ctx, emailID)

// Templates
tpl, err := client.CreateTemplate(ctx, mailafrica.TemplateRequest{
    Name:     "Welcome",
    Subject:  "Welcome!",
    HTMLBody: "<p>Hi {{name}}</p>",
})
tpls, err := client.ListTemplates(ctx)
tpl, err := client.GetTemplate(ctx, tplID)
tpl, err := client.UpdateTemplate(ctx, tplID, mailafrica.TemplateRequest{...})
client.DeleteTemplate(ctx, tplID)
```

### Domains (Sending)

```go
resp, err := client.AddSendingDomain(ctx, mailafrica.AddSendingDomainRequest{
    Domain:        "example.com",
    FromLocalPart: "hello",
})
// resp.DNSRecords contains DKIM, SPF, DMARC records for DNS setup
fmt.Println("DKIM TXT:", resp.DNSRecords.DKIM.Value)

domains, err := client.ListSendingDomains(ctx)
client.VerifySendingDomain(ctx, domain.ID)
client.DeleteSendingDomain(ctx, domain.ID)

// Sender addresses
addr, err := client.CreateSenderAddress(ctx, domainID, "info")
addrs, err := client.ListSenderAddresses(ctx)
client.DeleteSenderAddress(ctx, addr.ID)
```

### Webhooks

```go
wh, err := client.CreateWebhook(ctx, mailafrica.CreateWebhookRequest{
    AddressID: 1,
    URL:       "https://example.com/hook",
    // Secret auto-generated if omitted
})
webhooks, err := client.ListWebhooks(ctx, addrID)
client.DeleteWebhook(ctx, wh.ID)
deliveries, err := client.ListWebhookDeliveries(ctx, wh.ID)
client.TestWebhook(ctx, wh.ID)
client.TriggerWebhook(ctx, wh.ID)
```

### Sandbox

```go
cred, err := client.CreateSandboxCredential(ctx, mailafrica.CreateCredentialRequest{
    Scopes: strPtr("send,read"),
})
creds, err := client.ListSandboxCredentials(ctx)
client.RevokeSandboxCredential(ctx, cred.ID)

smtpCreds, err := client.GetSMTPSandboxCredentials(ctx)
smtpCreds, err := client.RegenerateSMTPSandboxPassword(ctx)
// smtpCreds.Password is shown only on first generation / regeneration

msgs, pagination, err := client.ListSandboxMessages(ctx, mailafrica.ListOpts{})
msg, err := client.GetSandboxMessage(ctx, msgID)
client.ClearSandboxMessages(ctx)
```

### Billing

```go
balance, err := client.GetBalance(ctx)
topup, err := client.InitiateTopup(ctx, 10000)
topup, err := client.InitiatePhoneTopup(ctx, 10000)
```

### SMS Notifications

```go
notif, err := client.CreateSMSNotification(ctx, mailafrica.CreateSMSNotificationRequest{
    AddressID:   1,
    PhoneNumber: "+255712345678",
    APIKey:      "SENDAFRICA_...",
})
// notif.APIKey is shown only once — store it securely
notifs, err := client.ListSMSNotifications(ctx, addrID)
client.RevokeSMSNotification(ctx, notif.ID)
deliveries, err := client.ListSMSDeliveries(ctx, notif.ID)
```

### Compliance

```go
profile, err := client.GetComplianceProfile(ctx)
profile, err := client.UpdateComplianceProfile(ctx, mailafrica.UpdateComplianceProfileRequest{
    PDPCRegistered:        boolPtr(true),
    PDPCCertificateNumber: strPtr("PDPC/2024/001"),
    DefaultRetentionDays:  intPtr(30),
})
export, err := client.GetAuditExport(ctx)
```

### Agent (AI Auto-Reply)

```go
configs, err := client.ListAgentConfigs(ctx)
config, err := client.GetAgentConfig(ctx, addressID)
config, err := client.UpdateAgentConfig(ctx, addressID, mailafrica.UpdateAgentConfigRequest{
    Mode:     "draft",
    Enabled:  boolPtr(true),
    Persona:  strPtr("You are a helpful assistant."),
})
draft, err := client.GenerateAgentDraft(ctx, addressID, mailafrica.AgentDraftRequest{
    Subject: "Re: Hello",
    Body:    "Thanks for reaching out...",
})
```

### API Keys

```go
key, err := client.CreateAPIKey(ctx, mailafrica.CreateAPIKeyRequest{
    Name:   "CLI Key",
    Scopes: "send,read",
    // ExpiresAt: timePtr(time.Now().Add(365 * 24 * time.Hour)),
})
// key.Key is shown only once — store it securely
keys, err := client.ListAPIKeys(ctx)
client.RevokeAPIKey(ctx, key.APIKey.ID)
```

## Pagination

List endpoints that return collections support pagination:

```go
emails, pagination, err := client.ListSentEmails(ctx, mailafrica.ListOpts{
    Page:    1,
    PerPage: 25, // default; capped at 100
})
fmt.Println("total pages:", pagination.TotalPages)
```

## Hooks

Optional hooks let you observe requests, responses, and errors without depending on a logging framework:

```go
client := mailafrica.New(mailafrica.Config{
    BaseURL: "https://api.mailafrica.online",
    Hooks: &mailafrica.Hooks{
        OnRequest: func(req *http.Request) {
            log.Println("request:", req.Method, req.URL)
        },
        OnResponse: func(resp *http.Response, duration time.Duration) {
            log.Println("response:", resp.StatusCode, duration)
        },
        OnError: func(err error) {
            log.Println("error:", err)
        },
    },
})
```

## Notes

- **Stdlib only**: The SDK uses only the Go standard library. No external HTTP clients, logging frameworks, or OpenAPI generators.
- **Context propagation**: Every method accepts `context.Context` as its first argument and passes it to the HTTP request.
- **No secrets in tests**: Examples and tests use fake keys like `MAIL_test_...`.
- **Admin routes excluded**: The SDK exposes only user-facing endpoints.
- **OAuth flows**: CamelAccounts and Google OAuth require browser redirects and cookies. Use the JSON-friendly endpoints (`Register`, `Login`, `Refresh`, `VerifyEmail`) from the SDK; for OAuth, open the authorization URL in a browser and handle the callback on your frontend.

## Testing

```bash
go test ./...
go vet ./...
```

## Versioning

The SDK follows SemVer. Start at `v0.1.0` for initial development.
