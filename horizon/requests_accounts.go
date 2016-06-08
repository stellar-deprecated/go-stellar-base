package horizon

// LoadAccount loads the account state from horizon. err can be either error
// object or horizon.Error object.
func (c *Client) LoadAccount(accountID string) (account Account, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("accounts").AddSegment(accountID).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &account)
	return
}

// LoadAccounts returns list of accounts returned by Horizon `/accounts` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadAccounts(params ...interface{}) (page AccountsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("accounts").Params(params...).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &page)
	return
}
