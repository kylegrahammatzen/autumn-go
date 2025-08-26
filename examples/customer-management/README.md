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
  Name: Jane Smith
  Email: jane@example.com
  Created At: 1756222842287

Fetching customer details...
Customer details:
  ID: jane_doe
  Autumn ID: 
  Products: 1
  Features: 3

Setting bonus feature balance...
Bonus credits applied successfully!

Updating customer information...
Updated customer name to: Jane Smith

Generating billing portal URL...
Billing portal URL: https://billing.stripe.com/p/session/test_YWNjdF8xUmtCWFZSZVBGRVpYR0FELF...

Final customer state:
Customer: Jane Smith (jane_doe)
  Products: 1 active
  Features: 3 available
  - searches: 5 remaining
  - sources: 2 remaining
  - seats: 0 remaining
```

Note: The features and products shown depend on your Autumn dashboard configuration and default products attached to customers.

## Features Demonstrated

- Create new customers with contact information
- Retrieve detailed customer data including products and features
- Manually set feature balances (useful for bonuses or support)
- Update customer information
- Generate Stripe billing portal URLs for self-service
- Display comprehensive customer state