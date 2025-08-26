package entities

import "github.com/kylegrahammatzen/autumn-go/autumn/core"

type Client struct {
	core *core.Client
}

func NewClient(coreClient *core.Client) *Client {
	return &Client{core: coreClient}
}

func (e *Client) Create(customerID string, entity EntityOptions) error {
	err := e.core.Request("POST", "/customers/"+customerID+"/entities", []EntityOptions{entity}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (e *Client) CreateMultiple(customerID string, entities []EntityOptions) error {
	err := e.core.Request("POST", "/customers/"+customerID+"/entities", entities, nil)
	if err != nil {
		return err
	}
	return nil
}

func (e *Client) Delete(customerID string, entityID string) error {
	err := e.core.Request("DELETE", "/customers/"+customerID+"/entities/"+entityID, nil, nil)
	if err != nil {
		return err
	}
	return nil
}