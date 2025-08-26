package autumn

import "time"

type AttachOptions struct {
	ProductID     *string           `json:"product_id,omitempty"`
	ProductIDs    []string          `json:"product_ids,omitempty"`
	SuccessURL    *string           `json:"success_url,omitempty"`
	ForceCheckout bool              `json:"force_checkout,omitempty"`
	Features      []Feature         `json:"features,omitempty"`
	EntityID      *string           `json:"entity_id,omitempty"`
	CustomerData  *CustomerData     `json:"customer_data,omitempty"`
	FreeTrial     *bool             `json:"free_trial,omitempty"`
	Options       []AttachOption    `json:"options,omitempty"`
}

type CheckOptions struct {
	ProductID       *string       `json:"product_id,omitempty"`
	FeatureID       *string       `json:"feature_id,omitempty"`
	RequiredBalance *int          `json:"required_balance,omitempty"`
	SendEvent       bool          `json:"send_event,omitempty"`
	WithPreview     bool          `json:"with_preview,omitempty"`
	EntityID        *string       `json:"entity_id,omitempty"`
	CustomerData    *CustomerData `json:"customer_data,omitempty"`
}

type TrackOptions struct {
	FeatureID      *string                `json:"feature_id,omitempty"`
	Value          int                    `json:"value,omitempty"`
	EntityID       *string                `json:"entity_id,omitempty"`
	EventName      *string                `json:"event_name,omitempty"`
	IdempotencyKey *string                `json:"idempotency_key,omitempty"`
	Properties     map[string]interface{} `json:"properties,omitempty"`
	CustomerData   *CustomerData          `json:"customer_data,omitempty"`
}

type CheckoutOptions struct {
	SuccessURL *string `json:"success_url,omitempty"`
}

type AttachRequest struct {
	CustomerID    string            `json:"customer_id"`
	ProductID     *string           `json:"product_id,omitempty"`
	ProductIDs    []string          `json:"product_ids,omitempty"`
	SuccessURL    *string           `json:"success_url,omitempty"`
	ForceCheckout bool              `json:"force_checkout,omitempty"`
	Features      []Feature         `json:"features,omitempty"`
	EntityID      *string           `json:"entity_id,omitempty"`
	CustomerData  *CustomerData     `json:"customer_data,omitempty"`
	FreeTrial     *bool             `json:"free_trial,omitempty"`
	Options       []AttachOption    `json:"options,omitempty"`
}

type CheckRequest struct {
	CustomerID      string        `json:"customer_id"`
	ProductID       *string       `json:"product_id,omitempty"`
	FeatureID       *string       `json:"feature_id,omitempty"`
	RequiredBalance *int          `json:"required_balance,omitempty"`
	SendEvent       bool          `json:"send_event,omitempty"`
	WithPreview     bool          `json:"with_preview,omitempty"`
	EntityID        *string       `json:"entity_id,omitempty"`
	CustomerData    *CustomerData `json:"customer_data,omitempty"`
}

type TrackRequest struct {
	CustomerID     string                 `json:"customer_id"`
	FeatureID      *string                `json:"feature_id,omitempty"`
	Value          int                    `json:"value,omitempty"`
	EntityID       *string                `json:"entity_id,omitempty"`
	EventName      *string                `json:"event_name,omitempty"`
	IdempotencyKey *string                `json:"idempotency_key,omitempty"`
	Properties     map[string]interface{} `json:"properties,omitempty"`
	CustomerData   *CustomerData          `json:"customer_data,omitempty"`
}

type CheckoutRequest struct {
	CustomerID string  `json:"customer_id"`
	ProductID  string  `json:"product_id"`
	SuccessURL *string `json:"success_url,omitempty"`
}

type AttachResponse struct {
	CustomerID  string   `json:"customer_id"`
	Code        string   `json:"code"`
	Message     string   `json:"message"`
	CheckoutURL *string  `json:"checkout_url"`
	ProductIDs  []string `json:"product_ids"`
	Success     *bool    `json:"success"`
}

type CheckResponse struct {
	Allowed        bool             `json:"allowed"`
	CustomerID     string           `json:"customer_id"`
	Code           string           `json:"code"`
	Balance        *float64         `json:"balance"`
	FeatureID      *string          `json:"feature_id"`
	ProductID      *string          `json:"product_id"`
	Status         *string          `json:"status"`
	FeaturePreview *FeaturePreview  `json:"feature_preview"`
	ProductPreview *ProductPreview  `json:"product_preview"`
}

type TrackResponse struct {
	ID         string  `json:"id"`
	Code       string  `json:"code"`
	CustomerID string  `json:"customer_id"`
	FeatureID  *string `json:"feature_id"`
	EventName  *string `json:"event_name"`
}

type CheckoutResponse struct {
	URL            *string        `json:"url"`
	CustomerID     string         `json:"customer_id"`
	HasProrations  bool           `json:"has_prorations"`
	Total          float64        `json:"total"`
	Currency       string         `json:"currency"`
	Lines          []CheckoutLine `json:"lines"`
	Options        []AttachOption `json:"options"`
	Product        Product        `json:"product"`
	CurrentProduct *Product       `json:"current_product"`
	NextCycle      *Cycle         `json:"next_cycle"`
}

type CustomerData struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type Feature struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Limit    *int   `json:"limit"`
	Reset    string `json:"reset"`
	IsUnique bool   `json:"is_unique"`
}

type AttachOption struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type CheckoutLine struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Amount   float64 `json:"amount"`
}

type Product struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	CreatedAt  int64         `json:"created_at"`
	Env        string        `json:"env"`
	Version    int           `json:"version"`
	IsAddOn    bool          `json:"is_add_on"`
	IsDefault  bool          `json:"is_default"`
	Items      []ProductItem `json:"items"`
	FreeTrial  *FreeTrial    `json:"free_trial"`
}

type ProductItem struct {
	Type        string  `json:"type"`
	PriceID     string  `json:"price_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Interval    *string `json:"interval"`
	Description *string `json:"description"`
}

type FreeTrial struct {
	Days int `json:"days"`
}

type Cycle struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type FeaturePreview struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductPreview struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}