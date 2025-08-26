package features

import (
	"github.com/kylegrahammatzen/autumn-go/autumn/core"
	"github.com/kylegrahammatzen/autumn-go/autumn/features/entities"
)

type Client struct {
	core     *core.Client
	Entities *entities.Client
}

func NewClient(coreClient *core.Client) *Client {
	return &Client{
		core:     coreClient,
		Entities: entities.NewClient(coreClient),
	}
}

func (f *Client) SetUsage(options SetUsageOptions) error {
	err := f.core.Request("POST", "/usage", options, nil)
	if err != nil {
		return err
	}
	return nil
}

func (f *Client) CreateEntity(customerID string, entity entities.EntityOptions) error {
	return f.Entities.Create(customerID, entity)
}

func (f *Client) CreateEntities(customerID string, entityOptions []entities.EntityOptions) error {
	return f.Entities.CreateMultiple(customerID, entityOptions)
}

func (f *Client) DeleteEntity(customerID string, entityID string) error {
	return f.Entities.Delete(customerID, entityID)
}