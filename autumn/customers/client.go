package customers

import "github.com/kylegrahammatzen/autumn-go/autumn/core"

type Client struct {
	core *core.Client
}

func NewClient(coreClient *core.Client) *Client {
	return &Client{core: coreClient}
}

func (c *Client) Create(options CreateCustomerOptions) (*CustomerResponse, error) {
	var response CustomerResponse
	err := c.core.Request("POST", "/customers", options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Get(customerID string) (*CustomerResponse, error) {
	var response CustomerResponse
	err := c.core.Request("GET", "/customers/"+customerID, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Update(customerID string, options UpdateCustomerOptions) (*CustomerResponse, error) {
	var response CustomerResponse
	err := c.core.Request("POST", "/customers/"+customerID, options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SetBalance(customerID string, options SetBalanceOptions) error {
	err := c.core.Request("POST", "/customers/"+customerID+"/balances", options, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) BillingPortal(customerID string, options BillingPortalOptions) (*BillingPortalResponse, error) {
	var response BillingPortalResponse
	err := c.core.Request("POST", "/customers/"+customerID+"/billing_portal", options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}