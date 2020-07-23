package swimlane

import (
	"crypto/tls"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// Client .
type Client struct {
	baseURL            *url.URL
	privateAccessToken string
	rClient            *resty.Client
}

// NewClient generates a new resty client for the Swimlane API.
func NewClient(baseURL string, privateAccessToken string) (*Client, error) {
	client := resty.New()
	client.
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// client.SetDebug(true)

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		msg := fmt.Sprintf("Ran into an error whilst parsing the baseURL: %s", err)
		panic(msg)
	}

	return &Client{
		parsedBaseURL,
		privateAccessToken,
		client,
	}, nil
}

// R generates a new resty Request with the required auth headers and content
// type
func (c *Client) R() *resty.Request {
	client := c.rClient
	request := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Private-Token", c.privateAccessToken)
	return request
}

func (c *Client) parseURL(path string) string {
	parsedURL, err := url.Parse(path)
	if err != nil {
		panic(err)
	}

	fullURL := c.baseURL.ResolveReference(parsedURL)
	return fullURL.String()
}
