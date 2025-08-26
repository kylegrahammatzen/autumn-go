# Customer Management Example

Complete example demonstrating customer lifecycle management using the Autumn Go SDK.

## Setup

1. Get your API key from https://app.useautumn.com/sandbox/dev
2. Replace the example key in `main.go` with your actual key
3. Run the example:
   ```bash
   cd examples/customer-management
   go run main.go
   ```

## Output

```
Creating new customer...
Created customer:
  ID: jane_doe
  Name: Jane Doe
  Email: jane@example.com
  Created At: 1677649423000

Fetching customer details...
Customer details:
  ID: jane_doe
  Autumn ID: cus_2w5dzidzFD1cESxOGnn9frVuVcm
  Products: 1
  Features: 2

Setting bonus feature balance...
Bonus credits applied successfully!

Updating customer information...
Updated customer name to: Jane Smith

Generating billing portal URL...
Billing portal URL: https://billing.stripe.com/session/<hash>

Final customer state:
Customer: Jane Smith (jane_doe)
  Products: 1 active
  Features: 2 available
  - chat_messages: 50 remaining
  - api_calls: unlimited
```

## Features Demonstrated

- Create new customers with contact information
- Retrieve detailed customer data including products and features
- Manually set feature balances (useful for bonuses or support)
- Update customer information
- Generate Stripe billing portal URLs for self-service
- Display comprehensive customer state