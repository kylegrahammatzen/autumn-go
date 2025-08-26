package main

import (
	"fmt"
	"log"

	autumn "github.com/kylegrahammatzen/autumn-go"
)

func main() {
	client := autumn.NewClient("am_sk_test_XESp2wyPE...")

	attachReq := autumn.AttachRequest{
		CustomerID: "john_doe",
		ProductID:  autumn.StringPtr("chat_messages"),
	}

	attachResp, err := client.Attach(attachReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Attached customer: %+v\n", attachResp)

	checkReq := autumn.CheckRequest{
		CustomerID: "john_doe",
		ProductID:  autumn.StringPtr("chat_messages"),
	}

	checkResp, err := client.Check(checkReq)
	if err != nil {
		log.Fatal(err)
	}

	if checkResp.Allowed {
		fmt.Println("Sending chat message...")
	}

	trackReq := autumn.TrackRequest{
		CustomerID: "john_doe",
		FeatureID:  autumn.StringPtr("chat_messages"),
		Value:      1,
	}

	trackResp, err := client.Track(trackReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tracked usage: %+v\n", trackResp)

	checkResp2, err := client.Check(checkReq)
	if err != nil {
		log.Fatal(err)
	}

	if !checkResp2.Allowed {
		fmt.Println("Customer has run out of chat messages.")
	}
}