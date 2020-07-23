package swimlane

import (
	"os"
	"testing"
)


func NewTestClient() (*Client, error) {
	BaseURL := os.Getenv("SWIMLANE_BASE_URL")
	Token := os.Getenv("SWIMLANE_TOKEN")
	client, err := NewClient(BaseURL, Token)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestNewClientInitialises(t *testing.T) {
	_, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
}
