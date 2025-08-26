# Autumn Go

Unofficial Go SDK for Autumn's subscription and usage tracking API.

[![Go Reference](https://pkg.go.dev/badge/github.com/kylegrahammatzen/autumn-go.svg)](https://pkg.go.dev/github.com/kylegrahammatzen/autumn-go)
[![Discord](https://img.shields.io/badge/Join%20Community-5865F2?logo=discord&logoColor=white)](https://discord.gg/53emPtY9tA)

## Getting Started

Installation (Go 1.19+ required):

```bash
go get github.com/kylegrahammatzen/autumn-go
```

## Usage

1. Define products and features using configuration files (similar to TypeScript CLI)
2. Attach customers to products/subscription plans
3. Check if customers have access to features before granting access
4. Track usage when customers consume features
5. Let Autumn handle all the complex billing logic and Stripe integration

### Config

```go
import "github.com/kylegrahammatzen/autumn-go"

// Define features
var messages = autumn.NewFeature("messages", "Messages", autumn.FeatureTypeSingleUse)

// Create products 
var pro = autumn.NewProduct("pro", "Pro",
    autumn.PriceItemWithInterval(20, autumn.IntervalMonth),
    autumn.FeatureItemWithUsage(messages.ID, 1000, autumn.IntervalPtr(autumn.IntervalMonth)),
)

// Convert to API calls
autumnProducts := autumn.Config{
    Features: []autumn.FeatureConfig{messages},
    Products: []autumn.ProductConfig{pro},
}.ToAutumnProducts()

client := autumn.NewClient("your-api-key")
product, err := client.Products.Create(autumnProducts[0])
```


## Examples

See [examples/](examples/) for complete working examples including error handling and different use cases.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup and guidelines.

## License

`autumn-go` is distributed under the terms of the [MIT](https://spdx.org/licenses/MIT.html) license.