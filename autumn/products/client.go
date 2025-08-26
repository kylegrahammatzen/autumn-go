package products

import "github.com/kylegrahammatzen/autumn-go/autumn/core"

type Client struct {
	core *core.Client
}

func NewClient(coreClient *core.Client) *Client {
	return &Client{core: coreClient}
}

func (p *Client) Create(options CreateProductOptions) (*ProductResponse, error) {
	var response ProductResponse
	err := p.core.Request("POST", "/products", options, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *Client) Get(productID string) (*ProductResponse, error) {
	var response ProductResponse
	err := p.core.Request("GET", "/products/"+productID, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}