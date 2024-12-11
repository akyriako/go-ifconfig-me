package ifconfigme

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const ifconfigMeUrl = "https://ifconfig.me/all.json"

type Response struct {
	IpAddr     string `json:"ip_addr"`
	RemoteHost string `json:"remote_host"`
	UserAgent  string `json:"user_agent"`
	Port       string `json:"port"`
	Language   string `json:"language"`
	Method     string `json:"method"`
	Encoding   string `json:"encoding"`
	Mime       string `json:"mime"`
	Via        string `json:"via"`
	Forwarded  string `json:"forwarded"`
}

type Client struct {
	httpClient *http.Client
}

type ClientOption func(*Client)

func WithTransport(transport *http.Transport) ClientOption {
	return func(c *Client) {
		c.httpClient.Transport = transport
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func NewClient(opts ...ClientOption) *Client {
	timeout := 500 * time.Millisecond
	transport := &http.Transport{}

	client := &Client{
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   timeout,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (r *Client) Get() (*Response, error) {
	httpResponse, err := r.httpClient.Get(ifconfigMeUrl)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, fmt.Errorf("http status %d", httpResponse.StatusCode)
	}

	httpBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(httpBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
