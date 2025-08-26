# Basic Example

---

This example demonstrates the core functionality of the Autumn Go SDK for subscription and usage tracking.

## What This Example Does

---

1. Attaches a customer to the "example" product
2. Checks access to ensure the customer can use features
3. Tracks usage when the customer consumes a chat message
4. Checks access again to see remaining balance after usage
5. Handles checkout if the customer runs out of credits

## Running the Example

---

```bash
cd examples/basic
go run main.go
```

## Expected Output

---

The example will show:
- Customer attachment confirmation
- Usage tracking response
- Feature access check with balance information
- Checkout flow (if needed) or confirmation of remaining access

## Understanding the Output

---

### Balance Information

When checking feature access, the SDK returns balance details:
- Current Balance: Remaining usage credits
- Required Balance: Credits needed for the operation
- Allowed: Whether the customer can proceed

### Checkout Behavior

If a customer runs out of credits, the example initiates checkout:
- $0.00 total: Customer already has access or no payment needed
- Checkout URL: Stripe payment link for upgrades (when payment required)

## Configuration

---

The example uses:
- Product ID: `example` 
- Feature ID: `chat_messages`
- Customer ID: `john_doe`
- API Key: Set in the client initialization

Replace these values with your own Autumn dashboard configuration for testing.