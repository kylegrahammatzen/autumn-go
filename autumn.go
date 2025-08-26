package autumn

import "github.com/kylegrahammatzen/autumn-go/autumn"

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

type CreateCustomerOptions = autumn.CreateCustomerOptions
type UpdateCustomerOptions = autumn.UpdateCustomerOptions
type SetBalanceOptions = autumn.SetBalanceOptions
type FeatureBalance = autumn.FeatureBalance
type BillingPortalOptions = autumn.BillingPortalOptions
type CustomerResponse = autumn.CustomerResponse
type BillingPortalResponse = autumn.BillingPortalResponse
type CustomerProduct = autumn.CustomerProduct
type CustomerFeature = autumn.CustomerFeature

func NewClient(token string) *Client {
	return autumn.NewClient(token)
}

func StringPtr(s string) *string {
	return autumn.StringPtr(s)
}
