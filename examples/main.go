package main

import (
	"fmt"
	"log"

	"github.com/kylegrahammatzen/autumn-go"
)

func main() {
	client := autumn.NewClient("am_sk_test_lYjPgVXKljm531lqkBPfpsHQ1x0mx9hH5UBsv8v4bc")

	attachOptions := autumn.AttachOptions{
		ProductID: autumn.StringPtr("example"),
	}

	attachResp, err := client.Attach("john_doe", attachOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(`Attached customer:
  Customer ID: %s
  Status: %s
  Message: %s
  Product IDs: %v
`, attachResp.CustomerID, attachResp.Code, attachResp.Message, attachResp.ProductIDs)

	checkOptions := autumn.CheckOptions{
		ProductID: autumn.StringPtr("example"),
	}

	checkResp, err := client.Check("john_doe", checkOptions)
	if err != nil {
		log.Fatal(err)
	}

	if checkResp.Allowed {
		fmt.Println("Sending chat message...")
	}

	trackOptions := autumn.TrackOptions{
		FeatureID: autumn.StringPtr("chat_messages"),
		Value:     1,
	}

	trackResp, err := client.Track("john_doe", trackOptions)
	if err != nil {
		log.Fatal(err)
	}
	featureID := "N/A"
	if trackResp.FeatureID != nil {
		featureID = *trackResp.FeatureID
	}
	eventName := "N/A"
	if trackResp.EventName != nil {
		eventName = *trackResp.EventName
	}
	fmt.Printf(`Tracked usage:
  ID: %s
  Code: %s
  Customer ID: %s
  Feature ID: %s
  Event Name: %s
`, trackResp.ID, trackResp.Code, trackResp.CustomerID, featureID, eventName)

	fmt.Println("Checking access again after usage...")
	checkResp2, err := client.Check("john_doe", checkOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`Second check result:
  Allowed: %t
  Customer ID: %s
  Code: %s`, checkResp2.Allowed, checkResp2.CustomerID, checkResp2.Code)

	if checkResp2.Balance != nil {
		fmt.Printf(`
  Balance Info:
    Feature ID: %s
    Current Balance: %d
    Required Balance: %d`, checkResp2.Balance.FeatureID, checkResp2.Balance.Balance, checkResp2.Balance.RequiredBalance)
	}
	fmt.Println()

	if !checkResp2.Allowed {
		fmt.Println("Customer has run out of chat messages.")
		
		fmt.Println("Creating checkout session...")
		checkoutOptions := autumn.CheckoutOptions{
			ProductID:  autumn.StringPtr("pro"),
			SuccessURL: autumn.StringPtr("https://example.com/success"),
		}
		
		checkoutResp, err := client.Checkout("john_doe", checkoutOptions)
		if err != nil {
			log.Fatal(err)
		}
		
		if checkoutResp.URL != nil {
			fmt.Printf("Checkout URL: %s\n", *checkoutResp.URL)
		} else {
			fmt.Printf(`Checkout completed:
  Customer ID: %s
  Total: $%.2f %s
  Has Prorations: %t
`, checkoutResp.CustomerID, checkoutResp.Total, checkoutResp.Currency, checkoutResp.HasProrations)
		}
	} else {
		fmt.Println("Customer still has access to chat messages.")
	}
}
