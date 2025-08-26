package autumn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.useautumn.com/v1"
	defaultTimeout = 30 * time.Second
)

type Client struct {
	token      string
	httpClient *http.Client
	Customers  *CustomersClient
	Products   *ProductsClient
	Features   *FeaturesClient
}

func NewClient(token string) *Client {
	client := &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
	client.Customers = &CustomersClient{client: client}
	client.Products = &ProductsClient{client: client}
	client.Features = &FeaturesClient{client: client}
	return client
}


func (c *Client) request(method, endpoint string, payload interface{}, result interface{}) error {
	url := baseURL + endpoint

	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal payload: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "autumn-go")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return handleErrorResponse(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) Attach(customerID string, options AttachOptions) (*AttachResponse, error) {
	if options.ProductID == nil && len(options.ProductIDs) == 0 {
		return nil, &Error{Code: "INVALID_REQUEST", Message: "Either ProductID or ProductIDs must be provided"}
	}
	if options.ProductID != nil && len(options.ProductIDs) > 0 {
		return nil, &Error{Code: "INVALID_REQUEST", Message: "Only one of ProductID or ProductIDs must be provided"}
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
	err := c.request("POST", "/attach", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Check(customerID string, options CheckOptions) (*CheckResponse, error) {
	if options.ProductID == nil && options.FeatureID == nil {
		return nil, &Error{Code: "INVALID_REQUEST", Message: "Either ProductID or FeatureID must be provided"}
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
	err := c.request("POST", "/check", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Track(customerID string, options TrackOptions) (*TrackResponse, error) {
	if options.FeatureID == nil && options.EventName == nil {
		return nil, &Error{Code: "INVALID_REQUEST", Message: "Either FeatureID or EventName must be provided"}
	}

	if options.Value == 0 {
		options.Value = 1
	}

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
	err := c.request("POST", "/track", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Checkout(customerID string, options CheckoutOptions) (*CheckoutResponse, error) {
	if options.ProductID == nil && len(options.ProductIDs) == 0 {
		return nil, &Error{Code: "INVALID_REQUEST", Message: "Either ProductID or ProductIDs must be provided"}
	}

	req := CheckoutRequest{
		CustomerID:            customerID,
		ProductID:             "",
		ProductIDs:            options.ProductIDs,
		SuccessURL:            options.SuccessURL,
		Options:               options.Options,
		Reward:                options.Reward,
		EntityID:              options.EntityID,
		CustomerData:          options.CustomerData,
		CheckoutSessionParams: options.CheckoutSessionParams,
	}

	if options.ProductID != nil {
		req.ProductID = *options.ProductID
	}

	var response CheckoutResponse
	err := c.request("POST", "/checkout", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// CustomersClient handles all customer management operations
type CustomersClient struct {
	client *Client
}

func (c *CustomersClient) Create(options CreateCustomerOptions) (*CustomerResponse, error) {
	var response CustomerResponse
	err := c.client.request("POST", "/customers", options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *CustomersClient) Get(customerID string) (*CustomerResponse, error) {
	var response CustomerResponse
	err := c.client.request("GET", "/customers/"+customerID, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *CustomersClient) Update(customerID string, options UpdateCustomerOptions) (*CustomerResponse, error) {
	var response CustomerResponse
	err := c.client.request("POST", "/customers/"+customerID, options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *CustomersClient) SetBalance(customerID string, options SetBalanceOptions) error {
	err := c.client.request("POST", "/customers/"+customerID+"/balances", options, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomersClient) BillingPortal(customerID string, options BillingPortalOptions) (*BillingPortalResponse, error) {
	var response BillingPortalResponse
	err := c.client.request("POST", "/customers/"+customerID+"/billing_portal", options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ProductsClient handles all product management operations
type ProductsClient struct {
	client *Client
}

func (p *ProductsClient) Create(options CreateProductOptions) (*ProductResponse, error) {
	var response ProductResponse
	err := p.client.request("POST", "/products", options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *ProductsClient) Get(productID string) (*ProductResponse, error) {
	var response ProductResponse
	err := p.client.request("GET", "/products/"+productID, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// FeaturesClient handles feature usage and entity management
type FeaturesClient struct {
	client *Client
}

func (f *FeaturesClient) SetUsage(options SetUsageOptions) error {
	err := f.client.request("POST", "/usage", options, nil)
	if err != nil {
		return err
	}
	return nil
}

func (f *FeaturesClient) CreateEntity(customerID string, entity EntityOptions) error {
	err := f.client.request("POST", "/customers/"+customerID+"/entities", []EntityOptions{entity}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (f *FeaturesClient) CreateEntities(customerID string, entities []EntityOptions) error {
	err := f.client.request("POST", "/customers/"+customerID+"/entities", entities, nil)
	if err != nil {
		return err
	}
	return nil
}

func (f *FeaturesClient) DeleteEntity(customerID string, entityID string) error {
	err := f.client.request("DELETE", "/customers/"+customerID+"/entities/"+entityID, nil, nil)
	if err != nil {
		return err
	}
	return nil
}