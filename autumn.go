package autumn

import (
	"github.com/kylegrahammatzen/autumn-go/autumn"
	"github.com/kylegrahammatzen/autumn-go/autumn/customers"
	"github.com/kylegrahammatzen/autumn-go/autumn/features"
	"github.com/kylegrahammatzen/autumn-go/autumn/features/entities"
	"github.com/kylegrahammatzen/autumn-go/autumn/products"
)

type Client = autumn.Client

type AttachOptions = autumn.AttachOptions
type CheckOptions = autumn.CheckOptions
type TrackOptions = autumn.TrackOptions
type CheckoutOptions = autumn.CheckoutOptions

type AttachResponse = autumn.AttachResponse
type CheckResponse = autumn.CheckResponse
type TrackResponse = autumn.TrackResponse
type CheckoutResponse = autumn.CheckoutResponse

type CustomerData = autumn.CustomerData
type Feature = autumn.Feature
type AttachOption = autumn.AttachOption

type CreateCustomerOptions = customers.CreateCustomerOptions
type UpdateCustomerOptions = customers.UpdateCustomerOptions
type SetBalanceOptions = customers.SetBalanceOptions
type FeatureBalance = customers.FeatureBalance
type BillingPortalOptions = customers.BillingPortalOptions
type CustomerResponse = customers.CustomerResponse
type BillingPortalResponse = customers.BillingPortalResponse
type CustomerProduct = customers.CustomerProduct
type CustomerFeature = customers.CustomerFeature

type SetUsageOptions = features.SetUsageOptions
type EntityOptions = entities.EntityOptions

type PriceTier = products.PriceTier
type ProductItemConfig = products.ProductItemConfig
type FreeTrialDuration = products.FreeTrialDuration
type FreeTrialOption = products.FreeTrialOption
type CreateProductOptions = products.CreateProductOptions
type ProductResponse = products.ProductResponse

const (
	FreeTrialDurationDay   = products.FreeTrialDurationDay
	FreeTrialDurationMonth = products.FreeTrialDurationMonth
	FreeTrialDurationYear  = products.FreeTrialDurationYear
)

func NewClient(token string) *Client {
	return autumn.NewClient(token)
}

func StringPtr(s string) *string {
	return autumn.StringPtr(s)
}

func FloatPtr(f float64) *float64 {
	return autumn.FloatPtr(f)
}
