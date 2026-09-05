package mailafrica

import (
	"context"
	"fmt"
	"time"
)

// ExampleClient_Register demonstrates user registration.
func ExampleClient_Register() {
	ctx := context.Background()
	client := New(Config{
		BaseURL: "https://api.mailafrica.online",
	})

	resp, err := client.Register(ctx, RegisterRequest{
		Email:       strPtr("user@example.com"),
		Phone:       strPtr("+255712345678"),
		Password:    "secure-password",
		Name:        "John Doe",
		CompanyName: "Acme Corp",
	})
	if err != nil {
		fmt.Println("register error:", err)
		return
	}
	fmt.Println("registered user ID:", resp.User.ID)
}

// ExampleClient_SendEmail demonstrates sending a single email.
func ExampleClient_SendEmail() {
	ctx := context.Background()
	client := New(Config{
		BaseURL: "https://api.mailafrica.online",
		APIKey:  "MAIL_test_...",
	})

	msg, err := client.SendEmail(ctx, SendEmailRequest{
		To:       []string{"recipient@example.com"},
		Subject:  "Hello from MailAfrica",
		HTMLBody: "<p>Hello!</p>",
		TextBody: "Hello!",
	})
	if err != nil {
		fmt.Println("send error:", err)
		return
	}
	fmt.Println("sent message ID:", msg.ID)
}

// ExampleClient_ListAddresses demonstrates listing inbound addresses.
func ExampleClient_ListAddresses() {
	ctx := context.Background()
	client := New(Config{
		BaseURL: "https://api.mailafrica.online",
		JWT:     "eyJ...",
	})

	addresses, err := client.ListAddresses(ctx)
	if err != nil {
		fmt.Println("list error:", err)
		return
	}
	for _, addr := range addresses {
		fmt.Println(addr.LocalPart)
	}
}

// ExampleClient_AddSendingDomain demonstrates adding a sending domain.
func ExampleClient_AddSendingDomain() {
	ctx := context.Background()
	client := New(Config{
		BaseURL: "https://api.mailafrica.online",
		APIKey:  "MAIL_test_...",
	})

	resp, err := client.AddSendingDomain(ctx, AddSendingDomainRequest{
		Domain:        "example.com",
		FromLocalPart: "hello",
	})
	if err != nil {
		fmt.Println("add domain error:", err)
		return
	}
	fmt.Println("add DKIM TXT:", resp.DNSRecords.DKIM.Value)
}

// ExampleClient_GetBalance demonstrates checking wallet balance.
func ExampleClient_GetBalance() {
	ctx := context.Background()
	client := New(Config{
		BaseURL: "https://api.mailafrica.online",
		APIKey:  "MAIL_test_...",
	})

	balance, err := client.GetBalance(ctx)
	if err != nil {
		fmt.Println("balance error:", err)
		return
	}
	fmt.Println("balance (TZS):", balance.BalanceTZS)
}

// ExampleIsInsufficientBalance demonstrates error inspection.
func ExampleIsInsufficientBalance() {
	ctx := context.Background()
	client := New(Config{BaseURL: "https://api.mailafrica.online"})

	_, err := client.SendEmail(ctx, SendEmailRequest{
		To:      []string{"recipient@example.com"},
		Subject: "Hello",
	})
	if err != nil {
		if IsInsufficientBalance(err) {
			fmt.Println("wallet is empty — please top up")
			return
		}
		fmt.Println("send failed:", err)
	}
}

func strPtr(s string) *string {
	return &s
}

// ExampleClient_BatchSend demonstrates batch sending.
func ExampleClient_BatchSend() {
	ctx := context.Background()
	client := New(Config{
		BaseURL: "https://api.mailafrica.online",
		APIKey:  "MAIL_test_...",
	})

	result, err := client.BatchSend(ctx, BatchSendRequest{
		To:       []string{"a@example.com", "b@example.com"},
		Subject:  "Batch hello",
		HTMLBody: "<p>Hello!</p>",
		TextBody: "Hello!",
	})
	if err != nil {
		fmt.Println("batch error:", err)
		return
	}
	fmt.Printf("sent %d of %d\n", result.Sent, result.Total)
}

// ExampleClient_CreateWebhook demonstrates webhook creation.
func ExampleClient_CreateWebhook() {
	ctx := context.Background()
	client := New(Config{
		BaseURL: "https://api.mailafrica.online",
		APIKey:  "MAIL_test_...",
	})

	wh, err := client.CreateWebhook(ctx, CreateWebhookRequest{
		AddressID: 1,
		URL:       "https://example.com/webhook",
		Secret:    "whsec_mysecret",
	})
	if err != nil {
		fmt.Println("webhook error:", err)
		return
	}
	fmt.Println("webhook ID:", wh.ID)
	// Store wh.Secret securely; it is shown only once.
	_ = time.Now()
}
