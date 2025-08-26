package config

type FeatureType string

const (
	FeatureTypeSingleUse     FeatureType = "single_use"
	FeatureTypeContinuousUse FeatureType = "continuous_use"
	FeatureTypeCreditSystem  FeatureType = "credit_system"
	FeatureTypeBoolean       FeatureType = "boolean"
)

type UsageModel string

const (
	UsageModelPrepaid    UsageModel = "prepaid"
	UsageModelPayPerUse  UsageModel = "pay_per_use"
)

type Interval string

const (
	IntervalHour        Interval = "hour"
	IntervalDay         Interval = "day"
	IntervalWeek        Interval = "week"
	IntervalMonth       Interval = "month"
	IntervalQuarter     Interval = "quarter"
	IntervalSemiAnnual  Interval = "semi_annual"
	IntervalYear        Interval = "year"
)

type CreditSchema struct {
	MeteredFeatureID string  `json:"metered_feature_id"`
	CreditCost       float64 `json:"credit_cost"`
}

type Feature struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Type         FeatureType    `json:"type"`
	CreditSchema []CreditSchema `json:"credit_schema,omitempty"`
}

type FeatureItem struct {
	FeatureID       string    `json:"feature_id"`
	IncludedUsage   *float64  `json:"included_usage,omitempty"`
	Interval        *Interval `json:"interval,omitempty"`
	EntityFeatureID *string   `json:"entity_feature_id,omitempty"`
}

type PriceItem struct {
	Price    float64   `json:"price"`
	Interval *Interval `json:"interval,omitempty"`
}

type PricedFeatureItem struct {
	FeatureID       string      `json:"feature_id"`
	Price           float64     `json:"price"`
	BillingUnits    *float64    `json:"billing_units,omitempty"`
	UsageModel      *UsageModel `json:"usage_model,omitempty"`
	IncludedUsage   *float64    `json:"included_usage,omitempty"`
	Interval        *Interval   `json:"interval,omitempty"`
	EntityFeatureID *string     `json:"entity_feature_id,omitempty"`
}

type ProductItem interface {
	isProductItem()
}

func (f FeatureItem) isProductItem()       {}
func (p PriceItem) isProductItem()         {}
func (pf PricedFeatureItem) isProductItem() {}

type Product struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	IsDefault bool          `json:"is_default,omitempty"`
	IsAddOn   bool          `json:"is_add_on,omitempty"`
	Items     []ProductItem `json:"items"`
}

type Config struct {
	Features []Feature `json:"features"`
	Products []Product `json:"products"`
}