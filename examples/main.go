package main

import (
	"fmt"
	"log"

	"github.com/kylegrahammatzen/autumn-go"
)

func main() {
	client := autumn.NewClient("am_sk_test_XESp2wyPE...")

	attachOptions := autumn.AttachOptions{
		ProductID: autumn.StringPtr("chat_messages"),
	}

	attachResp, err := client.Attach("john_doe", attachOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Attached customer: %+v\n", attachResp)

	checkOptions := autumn.CheckOptions{
		ProductID: autumn.StringPtr("chat_messages"),
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
	fmt.Printf("Tracked usage: %+v\n", trackResp)

	checkResp2, err := client.Check("john_doe", checkOptions)
	if err != nil {
		log.Fatal(err)
	}

	if !checkResp2.Allowed {
		fmt.Println("Customer has run out of chat messages.")
	}
}