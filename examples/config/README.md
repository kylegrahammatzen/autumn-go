# Configuration-Based Product Management

Define your Autumn products and features using Go configuration files, similar to the TypeScript CLI approach.

## Setup

1. Get your API key from https://app.useautumn.com/sandbox/dev
2. Replace the example key in `main.go` with your actual key
3. Run the example:
   ```bash
   cd examples/config
   go run .
   ```

## Features Demonstrated

- Define features with different types (single-use, continuous-use, boolean, credit system)
- Create products with pricing and feature configurations
- Convert configuration to Autumn API calls
- Validate business rules with real API responses

## Anatomy

Import the package and assemble your configuration:

```go
import "github.com/kylegrahammatzen/autumn-go"

// Define features
var messages = autumn.NewFeature("messages", "Messages", autumn.FeatureTypeSingleUse)
var seats = autumn.NewFeature("seats", "Seats", autumn.FeatureTypeContinuousUse)
var sso = autumn.NewFeature("sso", "SSO Auth", autumn.FeatureTypeBoolean)

// Create products with pricing and features
var free = autumn.DefaultProduct("free", "Free",
    autumn.FeatureItemWithUsage(messages.ID, 5, autumn.IntervalPtr(autumn.IntervalMonth)),
    autumn.FeatureItemConfig(sso.ID),
)

var pro = autumn.NewProduct("pro", "Pro",
    autumn.PriceItemWithInterval(20, autumn.IntervalMonth),
    autumn.FeatureItemWithUsage(messages.ID, 100, autumn.IntervalPtr(autumn.IntervalMonth)),
)

// Assemble configuration
var Config = autumn.Config{
    Features: []autumn.FeatureConfig{messages, seats, sso},
    Products: []autumn.ProductConfig{free, pro},
}

// Convert to Autumn API calls
autumnProducts := Config.ToAutumnProducts()
product, err := client.Products.Create(autumnProducts[0])
```

## Output

```
Generated 4 products from config:

Product 1: Free (free)
  Default: true, Add-on: false
  Items: 3
    Item 1:
      Feature: messages
      Type: single_use
      Included: 5
      Usage Model: prepaid
      Interval: month
    Item 2:
      Feature: seats
      Type: continuous_use
      Included: 3
      Usage Model: prepaid
    Item 3:
      Feature: sso
      Type: boolean
      Usage Model: prepaid

Product 2: Pro (pro)
  Default: false, Add-on: false
  Items: 4
    Item 1:
      Price: $20.00
      Interval: month
    Item 2:
      Price: $5.00
    Item 3:
      Feature: messages
      Type: single_use
      Included: 100
      Usage Model: prepaid
      Interval: month
    Item 4:
      Feature: messages
      Type: single_use
      Usage Model: pay_per_use
      Price: $0.01
      Billing Units: 10

Product 3: Team (team)
  Default: false, Add-on: false
  Items: 1
    Item 1:
      Feature: seats
      Type: continuous_use
      Price: $20.00

Product 4: Credit Top-up (top_up)
  Default: false, Add-on: true
  Items: 1
    Item 1:
      Feature: ai_credits
      Type: credit_system
      Usage Model: prepaid
      Price: $5.00
      Billing Units: 100

Example: Creating the 'pro' product via API...
Note: API call would create product (demo mode): autumn: Usage prices cannot be one-off if not set to prepaid (feature: messages) (code: invalid_price)

Config JSON representation:
{
  "features": [
    {
      "id": "messages",
      "name": "Messages", 
      "type": "single_use"
    },
    {
      "id": "seats",
      "name": "Seats",
      "type": "continuous_use"
    },
    {
      "id": "sso",
      "name": "SSO Auth",
      "type": "boolean"
    },
    {
      "id": "storage", 
      "name": "Storage",
      "type": "continuous_use"
    },
    {
      "id": "ai_credits",
      "name": "AI Credits",
      "type": "credit_system",
      "credit_schema": [
        {
          "metered_feature_id": "basic_message",
          "credit_cost": 1
        },
        {
          "metered_feature_id": "premium_message", 
          "credit_cost": 5
        }
      ]
    }
  ],
  "products": [
    {
      "id": "free",
      "name": "Free",
      "is_default": true,
      "items": [
        {
          "feature_id": "messages",
          "included_usage": 5,
          "interval": "month"
        },
        {
          "feature_id": "seats", 
          "included_usage": 3
        },
        {
          "feature_id": "sso"
        }
      ]
    },
    {
      "id": "pro",
      "name": "Pro", 
      "items": [
        {
          "price": 20,
          "interval": "month"
        },
        {
          "price": 5
        },
        {
          "feature_id": "messages",
          "included_usage": 100,
          "interval": "month"
        },
        {
          "feature_id": "messages",
          "price": 0.01,
          "billing_units": 10,
          "usage_model": "pay_per_use"
        }
      ]
    },
    {
      "id": "team",
      "name": "Team",
      "items": [
        {
          "feature_id": "seats",
          "price": 20
        }
      ]
    },
    {
      "id": "top_up",
      "name": "Credit Top-up",
      "is_add_on": true,
      "items": [
        {
          "feature_id": "ai_credits",
          "price": 5,
          "billing_units": 100,
          "usage_model": "prepaid"
        }
      ]
    }
  ]
}
```