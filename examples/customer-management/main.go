package main

import (
	"fmt"
	"log"

	"github.com/kylegrahammatzen/autumn-go"
)

func main() {
	client := autumn.NewClient("am_sk_test_YuLgxRagPBqHV63DERwj4XiMNyyGbSXXWZS72DUzQm")

	// 1. Create a new customer
	fmt.Println("Creating new customer...")
	createResp, err := client.Customers.Create(autumn.CreateCustomerOptions{
		ID:    "jane_doe",
		Name:  autumn.StringPtr("Jane Doe"),
		Email: autumn.StringPtr("jane@example.com"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(`Created customer:
  ID: %s
  Name: %s
  Email: %s
  Created At: %d
`, createResp.ID, *createResp.Name, *createResp.Email, createResp.CreatedAt)

	// 2. Get customer details
	fmt.Println("\nFetching customer details...")
	customer, err := client.Customers.Get("jane_doe")
	if err != nil {
		log.Fatal(err)
	}
	featuresCount := 0
	if features, ok := customer.Features.([]interface{}); ok {
		featuresCount = len(features)
	}
	fmt.Printf(`Customer details:
  ID: %s
  Autumn ID: %s
  Products: %d
  Features: %d
`, customer.ID, customer.AutumnID, len(customer.Products), featuresCount)

	// 3. Set feature balance (give bonus credits)
	fmt.Println("\nSetting bonus feature balance...")
	err = client.Customers.SetBalance("jane_doe", autumn.SetBalanceOptions{
		Balances: []autumn.FeatureBalance{
			{
				FeatureID: "chat_messages",
				Balance:   50, // Give 50 bonus messages
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bonus credits applied successfully!")

	// 4. Update customer information
	fmt.Println("\nUpdating customer information...")
	updatedCustomer, err := client.Customers.Update("jane_doe", autumn.UpdateCustomerOptions{
		Name: autumn.StringPtr("Jane Smith"), // Name change after marriage
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated customer name to: %s\n", *updatedCustomer.Name)

	// 5. Generate billing portal URL
	fmt.Println("\nGenerating billing portal URL...")
	billingResp, err := client.Customers.BillingPortal("jane_doe", autumn.BillingPortalOptions{
		ReturnURL: autumn.StringPtr("https://example.com/account"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Billing portal URL: %s\n", billingResp.URL)

	// 6. Final customer state
	fmt.Println("\nFinal customer state:")
	finalCustomer, err := client.Customers.Get("jane_doe")
	if err != nil {
		log.Fatal(err)
	}

	finalFeaturesCount := 0
	if features, ok := finalCustomer.Features.([]interface{}); ok {
		finalFeaturesCount = len(features)
		fmt.Printf(`Customer: %s (%s)
  Products: %d active
  Features: %d available
`, *finalCustomer.Name, finalCustomer.ID, len(finalCustomer.Products), finalFeaturesCount)

		// Show feature details
		for _, f := range features {
			if feature, ok := f.(map[string]interface{}); ok {
				featureID := feature["feature_id"]
				fmt.Printf("  - %v: ", featureID)
				if unlimited, ok := feature["unlimited"].(bool); ok && unlimited {
					fmt.Println("unlimited")
				} else if balance := feature["balance"]; balance != nil {
					fmt.Printf("%v remaining\n", balance)
				} else {
					fmt.Println("no balance info")
				}
			}
		}
	} else {
		fmt.Printf(`Customer: %s (%s)
  Products: %d active
  Features: unknown format
`, *finalCustomer.Name, finalCustomer.ID, len(finalCustomer.Products))
	}
}
