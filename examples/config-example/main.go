package main

import (
	"encoding/json"
	"fmt"

	"github.com/kylegrahammatzen/autumn-go"
)

func main() {
	client := autumn.NewClient("am_sk_test_MJu6VCtTKTUBpPh9XcT4f7GigFufIfKA2qsAvZ2FGF")

	autumnProducts := Config.ToAutumnProducts()

	fmt.Printf("Generated %d products from config:\n\n", len(autumnProducts))

	for i, product := range autumnProducts {
		fmt.Printf("Product %d: %s (%s)\n", i+1, product.Name, product.ID)
		fmt.Printf("  Default: %t, Add-on: %t\n", product.IsDefault, product.IsAddOn)
		fmt.Printf("  Items: %d\n", len(product.Items))

		for j, item := range product.Items {
			fmt.Printf("    Item %d:\n", j+1)
			if item.FeatureID != nil {
				fmt.Printf("      Feature: %s\n", *item.FeatureID)
				if item.FeatureType != nil {
					fmt.Printf("      Type: %s\n", *item.FeatureType)
				}
				if item.IncludedUsage != nil {
					fmt.Printf("      Included: %v\n", item.IncludedUsage)
				}
				if item.UsageModel != nil {
					fmt.Printf("      Usage Model: %s\n", *item.UsageModel)
				}
			}
			if item.Price != nil {
				fmt.Printf("      Price: $%.2f\n", *item.Price)
			}
			if item.BillingUnits != nil {
				fmt.Printf("      Billing Units: %v\n", *item.BillingUnits)
			}
			if item.Interval != nil {
				fmt.Printf("      Interval: %s\n", *item.Interval)
			}
		}
		fmt.Println()
	}

	fmt.Println("Example: Creating the 'pro' product via API...")

	for _, productConfig := range autumnProducts {
		if productConfig.ID == "pro" {
			product, err := client.Products.Create(productConfig)
			if err != nil {
				fmt.Printf("Note: API call would create product (demo mode): %v\n", err)
			} else {
				fmt.Printf("Created product: %s\n", product.Name)
			}
			break
		}
	}

	fmt.Println("\nConfig JSON representation:")
	configJSON, _ := json.MarshalIndent(Config, "", "  ")
	fmt.Println(string(configJSON))
}
