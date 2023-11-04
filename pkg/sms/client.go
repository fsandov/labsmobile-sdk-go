package sms

import (
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.labsmobile.com/json"
	defaultTimeout = 2000
)

var (
	errAPIKeyRequired   = fmt.Errorf("api key is required")
	errUsernameRequired = fmt.Errorf("username is required")
)

type (
	clientOptions struct {
		Timeout int64
	}

	// ClientOption is the type for the options of the Client.
	ClientOption func(*clientOptions)
)

// WithTimeOut sets the timeout for the http client in milliseconds. Default value is 2000. This is optional.
func WithTimeOut(timeout int64) ClientOption {
	return func(o *clientOptions) {
		if timeout > 0 {
			o.Timeout = timeout
		} else {
			o.Timeout = defaultTimeout
		}
	}
}

func (o *clientOptions) apply(opts ...ClientOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// Client is the client for the labosmobile.com API. It contains the http client, the baseURL, the APIKey and the SecretKey.
type Client struct {
	httpClient *http.Client
	baseURL    string
	APIKey     string
	Username   string
}

func NewClient(apiKey string, username string, opts ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, errAPIKeyRequired
	}

	if username == "" {
		return nil, errUsernameRequired
	}

	options := &clientOptions{}
	options.apply(opts...)

	return &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(options.Timeout) * time.Millisecond,
		},
		baseURL:  baseURL,
		APIKey:   apiKey,
		Username: username,
	}, nil
}
