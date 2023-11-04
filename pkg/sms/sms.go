package sms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	sendSMSPath = "/send"
)

func (c *Client) SendSMS(ctx context.Context, smsData SendSMSRequest) (*SendSMSResponse, error) {
	if err := c.verifySendSMSRequestFields(smsData); err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(smsData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling sms data. err: %v", err)
	}

	URI := fmt.Sprintf("%s%s", c.baseURL, sendSMSPath)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, URI, bytes.NewReader(marshal))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.APIKey)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request. err: %v", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad request, code: %d", res.StatusCode)
	}

	var response *SendSMSResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response. err: %v", err)
	}

	return response, nil
}

func (c *Client) verifySendSMSRequestFields(req SendSMSRequest) error {
	if c.APIKey == "" {
		return errAPIKeyRequired
	}
	if c.Username == "" {
		return errUsernameRequired
	}
	if req.Message == "" || len(req.Message) == 0 {
		return fmt.Errorf("message is required")
	}
	if len(req.Message) > 160 {
		return fmt.Errorf("message is too long")
	}
	if len(req.Recipient) == 0 {
		return fmt.Errorf("recipient is required")
	}
	return nil
}
