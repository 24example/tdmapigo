package client

import "resty.dev/v3"

const (
	ApiBaseURL string = "https://api-lk.tdme.ru/api/"
)

type Client struct {
	Request *resty.Request
}

func NewClient(token string) *Client {
	client := resty.New()

	client.SetBaseURL(ApiBaseURL)

	client.
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+token)

	return &Client{
		Request: client.R(),
	}
}
