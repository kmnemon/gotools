package shodan

const BaseURL = "http://127.0.0.1:1234"

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
