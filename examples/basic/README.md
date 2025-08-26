# Basic Example

Complete example showing how to use the Autumn Go SDK for subscription management and usage tracking.

## Setup

1. Get your API key from https://app.useautumn.com/sandbox/dev
2. Replace the example key in `main.go` with your actual key
3. Run the example:
   ```bash
   cd examples/basic
   go run main.go
   ```

## Output

```
Attached customer:
  Customer ID: john_doe
  Status: free_product_attached
  Message: Successfully attached example to john_doe
  Product IDs: [example]
Sending chat message...
Tracked usage:
  ID: evt_31pWSliPGtn56OwfbRZax17Ob2X
  Code: event_received
  Customer ID: john_doe
  Feature ID: chat_messages
  Event Name: N/A
Checking access again after usage...
Second check result:
  Allowed: false
  Customer ID: john_doe
  Code: feature_found
Customer has run out of chat messages.
Creating checkout session...
Checkout completed:
  Customer ID: john_doe
  Total: $0.00 USD
  Has Prorations: false
```

A $0.00 checkout means no payment is required.