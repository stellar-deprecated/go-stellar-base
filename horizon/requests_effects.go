package horizon

// LoadEffects returns list of effects returned by Horizon `/effects` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadEffects(params ...interface{}) (page EffectsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("effects").Params(params...).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalEffectsPage := effectsPage{}
	err = decodeResponse(resp, &internalEffectsPage)
	if err != nil {
		return
	}

	page = internalEffectsPage.ToExported()
	return
}

// LoadEffectsForAccount returns list of effects for a specified account returned by Horizon `/accounts/{accountID}/effects` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadEffectsForAccount(accountID string, params ...interface{}) (page EffectsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("accounts").
		AddSegment(accountID).
		AddSegment("effects").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalEffectsPage := effectsPage{}
	err = decodeResponse(resp, &internalEffectsPage)
	if err != nil {
		return
	}

	page = internalEffectsPage.ToExported()
	return
}

// LoadEffectsForLedger returns list of effects for a specified ledger returned by Horizon `/ledgers/{sequence}/effects` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadEffectsForLedger(sequence int32, params ...interface{}) (page EffectsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("ledgers").
		AddSegment(sequence).
		AddSegment("effects").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalEffectsPage := effectsPage{}
	err = decodeResponse(resp, &internalEffectsPage)
	if err != nil {
		return
	}

	page = internalEffectsPage.ToExported()
	return
}

// LoadEffectsForOperation returns list of effects for a specified operation returned by Horizon `/operations/{operationID}/effects` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadEffectsForOperation(operationID string, params ...interface{}) (page EffectsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("operations").
		AddSegment(operationID).
		AddSegment("effects").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalEffectsPage := effectsPage{}
	err = decodeResponse(resp, &internalEffectsPage)
	if err != nil {
		return
	}

	page = internalEffectsPage.ToExported()
	return
}

// LoadEffectsForTransaction returns list of effects for a specified transaction returned by Horizon `/transactions/{transactionID}/effects` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadEffectsForTransaction(transactionID string, params ...interface{}) (page EffectsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("transactions").
		AddSegment(transactionID).
		AddSegment("effects").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalEffectsPage := effectsPage{}
	err = decodeResponse(resp, &internalEffectsPage)
	if err != nil {
		return
	}

	page = internalEffectsPage.ToExported()
	return
}
