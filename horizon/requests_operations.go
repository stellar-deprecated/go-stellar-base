package horizon

// LoadOperation loads the operation from horizon. err can be either error
// object or horizon.Error object.
func (c *Client) LoadOperation(operationID string) (op Operation, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("operations").AddSegment(operationID).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalOperation := operation{}
	err = decodeResponse(resp, &internalOperation)
	if err != nil {
		return
	}

	op = internalOperation.ToExported()
	return
}

// LoadOperations returns list of operations returned by Horizon `/operations` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadOperations(params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("operations").Params(params...).Build()

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

// LoadOperationsForAccount returns list of operations for a specified account returned by Horizon `/accounts/{accountID}/operations` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadOperationsForAccount(accountID string, params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("accounts").
		AddSegment(accountID).
		AddSegment("operations").
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

// LoadOperationsForLedger returns list of operations for a specified ledger returned by Horizon `/ledgers/{sequence}/operations` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadOperationsForLedger(sequence int32, params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("ledgers").
		AddSegment(sequence).
		AddSegment("operations").
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

// LoadOperationsForTransactions returns list of operations for a specified transaction returned by Horizon `/transactions/{transactionID}/operations` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadOperationsForTransactions(transactionID string, params ...interface{}) (page OperationsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("transactions").
		AddSegment(transactionID).
		AddSegment("operations").
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
