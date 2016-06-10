package horizon

import (
	"net/url"
)

// LoadTransaction loads the transaction from horizon. err can be either error
// object or horizon.Error object.
func (c *Client) LoadTransaction(transactionID string) (transaction Transaction, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("transactions").AddSegment(transactionID).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &transaction)
	return
}

// LoadTransactions returns list of transactions returned by Horizon `/transactions` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadTransactions(params ...interface{}) (page TransactionsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("transactions").Params(params...).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &page)
	return
}

func (c *Client) LoadTransactionsForAccount(accountID string, params ...interface{}) (page TransactionsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("accounts").
		AddSegment(accountID).
		AddSegment("transactions").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &page)
	return
}

func (c *Client) LoadTransactionsForLedger(sequence int32, params ...interface{}) (page TransactionsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.
		AddSegment("ledgers").
		AddSegment(sequence).
		AddSegment("transactions").
		Params(params...).
		Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &page)
	return
}

// SubmitTransaction submits a transaction to the network. err can be either error object or horizon.Error object.
func (c *Client) SubmitTransaction(transactionEnvelopeXdr string) (response TransactionSuccess, err error) {
	v := url.Values{}
	v.Set("tx", transactionEnvelopeXdr)

	c.initHTTPClient()
	resp, err := c.Client.PostForm(c.URL+"/transactions", v)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &response)
	return
}
