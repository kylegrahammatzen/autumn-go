package features

type SetUsageOptions struct {
	FeatureID  string  `json:"feature_id"`
	CustomerID string  `json:"customer_id"`
	Value      float64 `json:"value"`
}