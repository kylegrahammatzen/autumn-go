package autumn

import (
	"github.com/kylegrahammatzen/autumn-go/autumn"
	"github.com/kylegrahammatzen/autumn-go/autumn/config"
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

type Config = config.Config
type FeatureConfig = config.Feature
type FeatureType = config.FeatureType
type UsageModel = config.UsageModel
type Interval = config.Interval
type CreditSchema = config.CreditSchema
type FeatureItem = config.FeatureItem
type PriceItem = config.PriceItem
type PricedFeatureItem = config.PricedFeatureItem
type ProductConfig = config.Product

const (
	FeatureTypeSingleUse     = config.FeatureTypeSingleUse
	FeatureTypeContinuousUse = config.FeatureTypeContinuousUse
	FeatureTypeCreditSystem  = config.FeatureTypeCreditSystem
	FeatureTypeBoolean       = config.FeatureTypeBoolean
	UsageModelPrepaid        = config.UsageModelPrepaid
	UsageModelPayPerUse      = config.UsageModelPayPerUse
	IntervalHour             = config.IntervalHour
	IntervalDay              = config.IntervalDay
	IntervalWeek             = config.IntervalWeek
	IntervalMonth            = config.IntervalMonth
	IntervalQuarter          = config.IntervalQuarter
	IntervalSemiAnnual       = config.IntervalSemiAnnual
	IntervalYear             = config.IntervalYear
)

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

func NewFeature(id, name string, featureType FeatureType) FeatureConfig {
	return config.FeatureConfig(id, name, featureType)
}

func CreditSystemFeature(id, name string, creditSchema []CreditSchema) FeatureConfig {
	return config.CreditSystemFeature(id, name, creditSchema)
}

func FeatureItemConfig(featureID string) FeatureItem {
	return config.FeatureItemConfig(featureID)
}

func FeatureItemWithUsage(featureID string, includedUsage float64, interval *Interval) FeatureItem {
	return config.FeatureItemWithUsage(featureID, includedUsage, interval)
}

func PriceItemConfig(price float64) PriceItem {
	return config.PriceItemConfig(price)
}

func PriceItemWithInterval(price float64, interval Interval) PriceItem {
	return config.PriceItemWithInterval(price, interval)
}

func PricedFeatureItemConfig(featureID string, price float64) PricedFeatureItem {
	return config.PricedFeatureItemConfig(featureID, price)
}

func PricedFeatureItemWithBilling(featureID string, price float64, billingUnits float64, usageModel UsageModel) PricedFeatureItem {
	return config.PricedFeatureItemWithBilling(featureID, price, billingUnits, usageModel)
}

func NewProduct(id, name string, items ...config.ProductItem) ProductConfig {
	return config.ProductConfig(id, name, items...)
}

func DefaultProduct(id, name string, items ...config.ProductItem) ProductConfig {
	return config.DefaultProduct(id, name, items...)
}

func AddOnProduct(id, name string, items ...config.ProductItem) ProductConfig {
	return config.AddOnProduct(id, name, items...)
}

func IntervalPtr(i Interval) *Interval {
	return config.IntervalPtr(i)
}
