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
	BaseURL           = "https://api.useautumn.com"
	Version           = "v1"
	LatestAPIVersion  = "1.2"
	DefaultTimeout    = 30 * time.Second
)

type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

func NewClient(token string, options ...ClientOption) *Client {
	client := &Client{
		token:   token,
		baseURL: BaseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

type ClientOption func(*Client)

func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func (c *Client) request(method, endpoint string, payload interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s%s", c.baseURL, Version, endpoint)

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

func (c *Client) Attach(req AttachRequest) (*AttachResponse, error) {
	var response AttachResponse
	err := c.request("POST", "/attach", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Check(req CheckRequest) (*CheckResponse, error) {
	var response CheckResponse
	err := c.request("POST", "/check", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Track(req TrackRequest) (*TrackResponse, error) {
	var response TrackResponse
	err := c.request("POST", "/track", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Checkout(req CheckoutRequest) (*CheckoutResponse, error) {
	var response CheckoutResponse
	err := c.request("POST", "/checkout", req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}