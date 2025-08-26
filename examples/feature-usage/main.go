package main

import (
	"fmt"
	"log"

	"github.com/kylegrahammatzen/autumn-go"
)

func main() {
	client := autumn.NewClient("am_sk_test_0kQxxSaElZDI6ny407rZqinEllV8k3fTi51bPnuXVJ")

	customerID := "jane_doe"

	fmt.Println("Cleaning up existing entities...")
	entityIDs := []string{"seat_001", "seat_002", "seat_003"}
	for _, entityID := range entityIDs {
		err := client.Features.DeleteEntity(customerID, entityID)
		if err != nil {
			fmt.Printf("Entity %s not found (expected)\n", entityID)
		} else {
			fmt.Printf("Cleaned up entity: %s\n", entityID)
		}
	}

	fmt.Println("\nSetting feature usage...")
	err := client.Features.SetUsage(autumn.SetUsageOptions{
		CustomerID: customerID,
		FeatureID:  "api_calls",
		Value:      150,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Set usage for 'api_calls' to 150 for customer %s\n", customerID)

	fmt.Println("\nCreating entities...")

	err = client.Features.CreateEntity(customerID, autumn.EntityOptions{
		ID:        "seat_001",
		FeatureID: "seats",
		Name:      "John Doe",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created entity: seat_001 (John Doe)")

	entities := []autumn.EntityOptions{
		{
			ID:        "seat_002",
			FeatureID: "seats",
			Name:      "Jane Smith",
		},
		{
			ID:        "seat_003",
			FeatureID: "seats",
			Name:      "Bob Johnson",
		},
	}

	err = client.Features.CreateEntities(customerID, entities)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created %d additional entities\n", len(entities))

	fmt.Println("\nSetting storage usage...")
	err = client.Features.SetUsage(autumn.SetUsageOptions{
		CustomerID: customerID,
		FeatureID:  "storage_gb",
		Value:      25.5,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set storage usage to 25.5 GB")

	fmt.Println("\nRemoving team member...")
	err = client.Features.DeleteEntity(customerID, "seat_003")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted entity: seat_003 (Bob Johnson)")

	fmt.Println("\nFeature usage management completed successfully!")
	fmt.Printf(`
Summary:
- Set API calls usage to 150
- Created 3 team member seats
- Set storage usage to 25.5 GB
- Removed 1 team member (2 seats remaining)
`)
}
