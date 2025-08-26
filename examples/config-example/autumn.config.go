package main

import "github.com/kylegrahammatzen/autumn-go"

var messages = autumn.NewFeature("messages", "Messages", autumn.FeatureTypeSingleUse)
var seats = autumn.NewFeature("seats", "Seats", autumn.FeatureTypeContinuousUse)
var sso = autumn.NewFeature("sso", "SSO Auth", autumn.FeatureTypeBoolean)
var storage = autumn.NewFeature("storage", "Storage", autumn.FeatureTypeContinuousUse)

var credits = autumn.CreditSystemFeature("ai_credits", "AI Credits", []autumn.CreditSchema{
	{
		MeteredFeatureID: "basic_message",
		CreditCost:       1,
	},
	{
		MeteredFeatureID: "premium_message", 
		CreditCost:       5,
	},
})

var free = autumn.DefaultProduct("free", "Free",
	autumn.FeatureItemWithUsage(messages.ID, 5, autumn.IntervalPtr(autumn.IntervalMonth)),
	autumn.FeatureItemWithUsage(seats.ID, 3, nil),
	autumn.FeatureItemConfig(sso.ID),
)

var pro = autumn.NewProduct("pro", "Pro",
	autumn.PriceItemWithInterval(20, autumn.IntervalMonth),
	autumn.PriceItemConfig(5),
	autumn.FeatureItemWithUsage(messages.ID, 100, autumn.IntervalPtr(autumn.IntervalMonth)),
	autumn.PricedFeatureItemWithBilling(messages.ID, 0.01, 10, autumn.UsageModelPayPerUse),
)

var team = autumn.NewProduct("team", "Team",
	autumn.PricedFeatureItemConfig(seats.ID, 20),
)

var topUp = autumn.AddOnProduct("top_up", "Credit Top-up",
	autumn.PricedFeatureItemWithBilling(credits.ID, 5, 100, autumn.UsageModelPrepaid),
)

var Config = autumn.Config{
	Features: []autumn.FeatureConfig{messages, seats, sso, storage, credits},
	Products: []autumn.ProductConfig{free, pro, team, topUp},
}