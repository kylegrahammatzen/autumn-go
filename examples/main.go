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

	checkResp2, err := client.Check("john_doe", checkOptions)
	if err != nil {
		log.Fatal(err)
	}

	if !checkResp2.Allowed {
		fmt.Println("Customer has run out of chat messages.")
	}
}
