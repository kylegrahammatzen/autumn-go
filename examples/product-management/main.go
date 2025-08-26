package main

import (
	"fmt"
	"log"

	"github.com/kylegrahammatzen/autumn-go"
)

func main() {
	client := autumn.NewClient("am_sk_test_9Kj2LmN8qR...")

	// 1. Create a new product
	fmt.Println("Creating new product...")
	product, err := client.Products.Create(autumn.CreateProductOptions{
		ID:   "my-premium-plan",
		Name: "Premium Plan",
		Items: []autumn.ProductItemConfig{
			{
				FeatureID:     autumn.StringPtr("api_calls"),
				FeatureType:   autumn.StringPtr("single_use"),
				IncludedUsage: 1000,
				Interval:      autumn.StringPtr("month"),
				UsageModel:    autumn.StringPtr("prepaid"),
				Price:         autumn.FloatPtr(29.99),
			},
		},
		FreeTrial: &autumn.FreeTrialOption{
			Duration:          autumn.FreeTrialDurationDay,
			Length:            14,
			UniqueFingerprint: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`Created product:
  ID: %s
  Name: %s
  Autumn ID: %s
  Environment: %s
  Version: %d
  Items: %d
`, product.ID, product.Name, product.AutumnID, product.Env, product.Version, len(product.Items))

	// 2. Get the product details
	fmt.Println("\nFetching product details...")
	retrievedProduct, err := client.Products.Get("my-premium-plan")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`Product details:
  ID: %s
  Name: %s
  Add-on: %t
  Default: %t
  Items: %d
`, retrievedProduct.ID, retrievedProduct.Name, retrievedProduct.IsAddOn, retrievedProduct.IsDefault, len(retrievedProduct.Items))

	// 3. Show product items
	fmt.Println("\nProduct items:")
	for i, item := range retrievedProduct.Items {
		fmt.Printf("  Item %d:\n", i+1)
		if item.FeatureID == nil {
			fmt.Printf("    Type: Price\n")
			if item.Price != nil {
				fmt.Printf("    Price: $%.2f\n", *item.Price)
			}
		} else {
			fmt.Printf("    Feature: %s\n", *item.FeatureID)
			if item.FeatureType != nil {
				fmt.Printf("    Type: %s\n", *item.FeatureType)
			}
			fmt.Printf("    Included Usage: %v\n", item.IncludedUsage)
			if item.UsageModel != nil {
				fmt.Printf("    Usage Model: %s\n", *item.UsageModel)
			}
		}
		if item.Interval != nil {
			fmt.Printf("    Interval: %s\n", *item.Interval)
		}
		fmt.Println()
	}

	// 4. Show free trial info
	if retrievedProduct.FreeTrial != nil {
		fmt.Printf(`Free trial:
  Duration: %d %s
  Unique fingerprint required: %t
`, retrievedProduct.FreeTrial.Length, retrievedProduct.FreeTrial.Duration, retrievedProduct.FreeTrial.UniqueFingerprint)
	}
}
