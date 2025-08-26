package products

type PriceTier struct {
	To     interface{} `json:"to"`
	Amount float64     `json:"amount"`
}

type ProductItemConfig struct {
	FeatureID               *string      `json:"feature_id,omitempty"`
	FeatureType             *string      `json:"feature_type,omitempty"`
	IncludedUsage           interface{}  `json:"included_usage,omitempty"`
	Interval                *string      `json:"interval,omitempty"`
	UsageModel              *string      `json:"usage_model,omitempty"`
	Price                   *float64     `json:"price,omitempty"`
	BillingUnits            *float64     `json:"billing_units,omitempty"`
	EntityFeatureID         *string      `json:"entity_feature_id,omitempty"`
	ResetUsageWhenEnabled   *bool        `json:"reset_usage_when_enabled,omitempty"`
	Tiers                   []PriceTier  `json:"tiers,omitempty"`
}

type FreeTrialDuration string

const (
	FreeTrialDurationDay   FreeTrialDuration = "day"
	FreeTrialDurationMonth FreeTrialDuration = "month"
	FreeTrialDurationYear  FreeTrialDuration = "year"
)

type FreeTrialOption struct {
	Duration           FreeTrialDuration `json:"duration"`
	Length             int               `json:"length"`
	UniqueFingerprint  bool              `json:"unique_fingerprint"`
}

type CreateProductOptions struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	IsAddOn     bool                `json:"is_add_on,omitempty"`
	IsDefault   bool                `json:"is_default,omitempty"`
	Items       []ProductItemConfig `json:"items,omitempty"`
	FreeTrial   *FreeTrialOption    `json:"free_trial,omitempty"`
}

type ProductResponse struct {
	AutumnID    string              `json:"autumn_id"`
	CreatedAt   int64               `json:"created_at"`
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Env         string              `json:"env"`
	IsAddOn     bool                `json:"is_add_on"`
	IsDefault   bool                `json:"is_default"`
	Group       *string             `json:"group"`
	Version     int                 `json:"version"`
	Items       []ProductItemConfig `json:"items"`
	FreeTrial   *FreeTrialOption    `json:"free_trial"`
}