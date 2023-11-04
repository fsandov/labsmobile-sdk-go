package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	getBalance = "/balance"
)

func (c *Client) GetBalance(ctx context.Context) (*GetBalanceResponse, error) {
	if c.APIKey == "" {
		return nil, errAPIKeyRequired
	}
	if c.Username == "" {
		return nil, errUsernameRequired
	}
	URI := fmt.Sprintf("%s%s", c.baseURL, getBalance)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, URI, nil)
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

	var response *GetBalanceResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response. err: %v", err)
	}

	return response, nil
}
