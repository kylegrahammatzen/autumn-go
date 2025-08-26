package customers

type CreateCustomerOptions struct {
	ID          string  `json:"id"`
	Email       *string `json:"email,omitempty"`
	Name        *string `json:"name,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
}

type UpdateCustomerOptions struct {
	ID          *string `json:"id,omitempty"`
	Email       *string `json:"email,omitempty"`
	Name        *string `json:"name,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
}

type SetBalanceOptions struct {
	Balances []FeatureBalance `json:"balances"`
}

type FeatureBalance struct {
	FeatureID string  `json:"feature_id"`
	Balance   float64 `json:"balance"`
}

type BillingPortalOptions struct {
	ReturnURL *string `json:"return_url,omitempty"`
}

type CustomerProduct struct {
	ID                  string   `json:"id"`
	Name                *string  `json:"name"`
	Group               *string  `json:"group"`
	Status              string   `json:"status"`
	StartedAt           int64    `json:"started_at"`
	CanceledAt          *int64   `json:"canceled_at"`
	SubscriptionIDs     []string `json:"subscription_ids"`
	CurrentPeriodStart  *int64   `json:"current_period_start"`
	CurrentPeriodEnd    *int64   `json:"current_period_end"`
}

type CustomerFeature struct {
	FeatureID      string      `json:"feature_id"`
	Unlimited      bool        `json:"unlimited"`
	Interval       *string     `json:"interval"`
	Balance        interface{} `json:"balance"`
	Usage          interface{} `json:"usage"`
	IncludedUsage  interface{} `json:"included_usage"`
	NextResetAt    *int64      `json:"next_reset_at"`
}

type CustomerResponse struct {
	AutumnID    string            `json:"autumn_id"`
	CreatedAt   int64             `json:"created_at"`
	Env         string            `json:"env"`
	ID          string            `json:"id"`
	Name        *string           `json:"name"`
	Email       *string           `json:"email"`
	Fingerprint *string           `json:"fingerprint"`
	StripeID    *string           `json:"stripe_id"`
	Products    []CustomerProduct `json:"products"`
	Features    interface{}       `json:"features"`
}

type BillingPortalResponse struct {
	URL string `json:"url"`
}