package autumn

import (
	"github.com/kylegrahammatzen/autumn-go/autumn/core"
	"github.com/kylegrahammatzen/autumn-go/autumn/customers"
	"github.com/kylegrahammatzen/autumn-go/autumn/features"
	"github.com/kylegrahammatzen/autumn-go/autumn/products"
)

type Client struct {
	core      *core.Client
	Customers *customers.Client
	Products  *products.Client
	Features  *features.Client
}

func NewClient(token string) *Client {
	coreClient := core.NewClient(token)
	return &Client{
		core:      coreClient,
		Customers: customers.NewClient(coreClient),
		Products:  products.NewClient(coreClient),
		Features:  features.NewClient(coreClient),
	}
}

func (c *Client) Attach(customerID string, options AttachOptions) (*AttachResponse, error) {
	if options.ProductID != nil && len(options.ProductIDs) > 0 {
		return nil, &core.Error{Code: "INVALID_REQUEST", Message: "Only one of ProductID or ProductIDs must be provided"}
	}

	req := AttachRequest{
		CustomerID:    customerID,
		ProductID:     options.ProductID,
		ProductIDs:    options.ProductIDs,
		SuccessURL:    options.SuccessURL,
		ForceCheckout: options.ForceCheckout,
		Features:      options.Features,
		EntityID:      options.EntityID,
		CustomerData:  options.CustomerData,
		FreeTrial:     options.FreeTrial,
		Options:       options.Options,
	}

	var response AttachResponse
	err := c.core.Request("POST", "/attach", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Check(customerID string, options CheckOptions) (*CheckResponse, error) {
	if options.ProductID == nil && options.FeatureID == nil {
		return nil, &core.Error{Code: "INVALID_REQUEST", Message: "Either ProductID or FeatureID must be provided"}
	}

	req := CheckRequest{
		CustomerID:      customerID,
		ProductID:       options.ProductID,
		FeatureID:       options.FeatureID,
		RequiredBalance: options.RequiredBalance,
		SendEvent:       options.SendEvent,
		WithPreview:     options.WithPreview,
		EntityID:        options.EntityID,
		CustomerData:    options.CustomerData,
	}

	var response CheckResponse
	err := c.core.Request("POST", "/check", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Track(customerID string, options TrackOptions) (*TrackResponse, error) {
	req := TrackRequest{
		CustomerID:     customerID,
		FeatureID:      options.FeatureID,
		Value:          options.Value,
		EntityID:       options.EntityID,
		EventName:      options.EventName,
		IdempotencyKey: options.IdempotencyKey,
		Properties:     options.Properties,
		CustomerData:   options.CustomerData,
	}

	var response TrackResponse
	err := c.core.Request("POST", "/track", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Checkout(customerID string, options CheckoutOptions) (*CheckoutResponse, error) {
	if options.ProductID == nil && len(options.ProductIDs) == 0 {
		return nil, &core.Error{Code: "INVALID_REQUEST", Message: "Either ProductID or ProductIDs must be provided"}
	}

	var productID string
	if options.ProductID != nil {
		productID = *options.ProductID
	}

	req := CheckoutRequest{
		CustomerID:            customerID,
		ProductID:             productID,
		ProductIDs:            options.ProductIDs,
		SuccessURL:            options.SuccessURL,
		Options:               options.Options,
		Reward:                options.Reward,
		EntityID:              options.EntityID,
		CustomerData:          options.CustomerData,
		CheckoutSessionParams: options.CheckoutSessionParams,
	}

	var response CheckoutResponse
	err := c.core.Request("POST", "/checkout", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}