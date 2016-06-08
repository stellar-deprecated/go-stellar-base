package horizon

// LoadLedger loads the ledger from horizon. err can be either error
// object or horizon.Error object.
func (c *Client) LoadLedger(sequence int32) (ledger Ledger, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("ledgers").AddSegment(sequence).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &ledger)
	return
}

// LoadLedgers returns list of ledgers returned by Horizon `/ledgers` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadLedgers(params ...interface{}) (page LedgersPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("ledgers").Params(params...).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &page)
	return
}
