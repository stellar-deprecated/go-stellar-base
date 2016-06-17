package horizon

// LoadPayments returns list of operations returned by Horizon `/payments` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadPayments(params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("payments").Params(params...).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalOperationsPage := operationsPage{}
	err = decodeResponse(resp, &internalOperationsPage)
	if err != nil {
		return
	}

	page = internalOperationsPage.ToExported()
	return
}

// LoadPaymentsForAccount returns list of operations for a specified account returned by Horizon `/accounts/{accountID}/payments` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadPaymentsForAccount(accountID string, params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("accounts").
		AddSegment(accountID).
		AddSegment("payments").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalOperationsPage := operationsPage{}
	err = decodeResponse(resp, &internalOperationsPage)
	if err != nil {
		return
	}

	page = internalOperationsPage.ToExported()
	return
}

// LoadPaymentsForLedger returns list of operations for a specified ledger returned by Horizon `/ledgers/{sequence}/payments` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadPaymentsForLedger(sequence int32, params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("ledgers").
		AddSegment(sequence).
		AddSegment("payments").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalOperationsPage := operationsPage{}
	err = decodeResponse(resp, &internalOperationsPage)
	if err != nil {
		return
	}

	page = internalOperationsPage.ToExported()
	return
}

// LoadPaymentsForTransactions returns list of operations for a specified transaction returned by Horizon `/transactions/{transactionID}/payments` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadPaymentsForTransactions(transactionID string, params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("transactions").
		AddSegment(transactionID).
		AddSegment("payments").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalOperationsPage := operationsPage{}
	err = decodeResponse(resp, &internalOperationsPage)
	if err != nil {
		return
	}

	page = internalOperationsPage.ToExported()
	return
}
