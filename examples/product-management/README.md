# Product Management Example

Complete example demonstrating product creation and management using the Autumn Go SDK.

## Setup

1. Get your API key from https://app.useautumn.com/sandbox/dev
2. Replace the example key in `main.go` with your actual key
3. Run the example:
   ```bash
   cd examples/product-management
   go run main.go
   ```

## Output

```
Creating new product...
Created product:
  ID: my-premium-plan
  Name: Premium Plan
  Autumn ID: prod_abc123
  Environment: sandbox
  Version: 1
  Items: 3

Fetching product details...
Product details:
  ID: my-premium-plan
  Name: Premium Plan
  Add-on: false
  Default: false
  Items: 3

Product items:
  Item 1:
    Type: Price
    Price: $29.99
    Interval: month

  Item 2:
    Feature: api_calls
    Type: single_use
    Included Usage: 1000
    Usage Model: prepaid
    Interval: month

  Item 3:
    Feature: storage
    Type: continuous_use
    Included Usage: Infinity
    Usage Model: prepaid
    Interval: month

Free trial:
  Duration: 14 days
  Unique fingerprint required: true
```

## Features Demonstrated

- Create products with pricing and feature configuration
- Retrieve detailed product information
- Configure different product item types (price, features)
- Set up free trials with fingerprint protection
- Handle both single-use and continuous-use features
- Display comprehensive product structure