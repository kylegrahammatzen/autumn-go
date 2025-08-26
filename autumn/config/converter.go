package config

import (
	"github.com/kylegrahammatzen/autumn-go/autumn/products"
)

func (c *Config) ToAutumnProducts() []products.CreateProductOptions {
	var autumnProducts []products.CreateProductOptions

	for _, product := range c.Products {
		var items []products.ProductItemConfig

		for _, item := range product.Items {
			switch v := item.(type) {
			case FeatureItem:
				var includedUsage interface{}
				if v.IncludedUsage != nil {
					includedUsage = *v.IncludedUsage
				}
				items = append(items, products.ProductItemConfig{
					FeatureID:       StringPtr(v.FeatureID),
					FeatureType:     inferFeatureType(v.FeatureID, c.Features),
					IncludedUsage:   includedUsage,
					Interval:        intervalToString(v.Interval),
					UsageModel:      StringPtr("prepaid"),
					EntityFeatureID: v.EntityFeatureID,
				})

			case PriceItem:
				items = append(items, products.ProductItemConfig{
					Price:    &v.Price,
					Interval: intervalToString(v.Interval),
				})

			case PricedFeatureItem:
				var includedUsage interface{}
				if v.IncludedUsage != nil {
					includedUsage = *v.IncludedUsage
				}
				items = append(items, products.ProductItemConfig{
					FeatureID:       StringPtr(v.FeatureID),
					FeatureType:     inferFeatureType(v.FeatureID, c.Features),
					Price:           &v.Price,
					BillingUnits:    v.BillingUnits,
					UsageModel:      usageModelToString(v.UsageModel),
					IncludedUsage:   includedUsage,
					Interval:        intervalToString(v.Interval),
					EntityFeatureID: v.EntityFeatureID,
				})
			}
		}

		autumnProducts = append(autumnProducts, products.CreateProductOptions{
			ID:        product.ID,
			Name:      product.Name,
			IsAddOn:   product.IsAddOn,
			IsDefault: product.IsDefault,
			Items:     items,
		})
	}

	return autumnProducts
}

func inferFeatureType(featureID string, features []Feature) *string {
	for _, feature := range features {
		if feature.ID == featureID {
			featureType := string(feature.Type)
			return &featureType
		}
	}
	return nil
}

func intervalToString(interval *Interval) *string {
	if interval == nil {
		return nil
	}
	s := string(*interval)
	return &s
}

func usageModelToString(usageModel *UsageModel) *string {
	if usageModel == nil {
		return nil
	}
	s := string(*usageModel)
	return &s
}