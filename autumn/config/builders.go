package config

func FeatureConfig(id, name string, featureType FeatureType) Feature {
	return Feature{
		ID:   id,
		Name: name,
		Type: featureType,
	}
}

func CreditSystemFeature(id, name string, creditSchema []CreditSchema) Feature {
	return Feature{
		ID:           id,
		Name:         name,
		Type:         FeatureTypeCreditSystem,
		CreditSchema: creditSchema,
	}
}

func FeatureItemConfig(featureID string) FeatureItem {
	return FeatureItem{
		FeatureID: featureID,
	}
}

func FeatureItemWithUsage(featureID string, includedUsage float64, interval *Interval) FeatureItem {
	return FeatureItem{
		FeatureID:     featureID,
		IncludedUsage: &includedUsage,
		Interval:      interval,
	}
}

func FeatureItemWithEntity(featureID string, includedUsage float64, interval *Interval, entityFeatureID string) FeatureItem {
	return FeatureItem{
		FeatureID:       featureID,
		IncludedUsage:   &includedUsage,
		Interval:        interval,
		EntityFeatureID: &entityFeatureID,
	}
}

func PriceItemConfig(price float64) PriceItem {
	return PriceItem{
		Price: price,
	}
}

func PriceItemWithInterval(price float64, interval Interval) PriceItem {
	return PriceItem{
		Price:    price,
		Interval: &interval,
	}
}

func PricedFeatureItemConfig(featureID string, price float64) PricedFeatureItem {
	return PricedFeatureItem{
		FeatureID: featureID,
		Price:     price,
	}
}

func PricedFeatureItemWithBilling(featureID string, price float64, billingUnits float64, usageModel UsageModel) PricedFeatureItem {
	return PricedFeatureItem{
		FeatureID:    featureID,
		Price:        price,
		BillingUnits: &billingUnits,
		UsageModel:   &usageModel,
	}
}

func PricedFeatureItemFull(featureID string, price float64, billingUnits *float64, usageModel *UsageModel, includedUsage *float64, interval *Interval, entityFeatureID *string) PricedFeatureItem {
	return PricedFeatureItem{
		FeatureID:       featureID,
		Price:           price,
		BillingUnits:    billingUnits,
		UsageModel:      usageModel,
		IncludedUsage:   includedUsage,
		Interval:        interval,
		EntityFeatureID: entityFeatureID,
	}
}

func ProductConfig(id, name string, items ...ProductItem) Product {
	return Product{
		ID:    id,
		Name:  name,
		Items: items,
	}
}

func DefaultProduct(id, name string, items ...ProductItem) Product {
	return Product{
		ID:        id,
		Name:      name,
		IsDefault: true,
		Items:     items,
	}
}

func AddOnProduct(id, name string, items ...ProductItem) Product {
	return Product{
		ID:      id,
		Name:    name,
		IsAddOn: true,
		Items:   items,
	}
}

func IntervalPtr(i Interval) *Interval {
	return &i
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func StringPtr(s string) *string {
	return &s
}

func UsageModelPtr(u UsageModel) *UsageModel {
	return &u
}