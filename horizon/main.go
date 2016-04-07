package horizon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{URL: "https://horizon-testnet.stellar.org"}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{URL: "https://horizon.stellar.org"}

// HorizonError struct contains the problem returned by Horizon
type Error struct {
	Response *http.Response
	Problem  Problem
}

func (herror *Error) Error() string {
	return "Horizon error"
}

type HorizonHttpClient interface {
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// Client struct contains data required to connect to Horizon instance
type Client struct {
	// URL of Horizon server to connect
	URL string
	// Will be populated with &http.Client when nil. If you want to configure your http.Client make sure Timeout is at least 10 seconds.
	Client HorizonHttpClient
	// clientInit initializes http client once
	clientInit sync.Once
}

// LoadAccount loads the account state from horizon. err can be either error object or horizon.Error object.
func (c *Client) LoadAccount(accountId string) (account Account, err error) {
	c.initHttpClient()
	resp, err := c.Client.Get(c.URL + "/accounts/" + accountId)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &account)
	return
}

// LoadLedger loads the state of the given ledger from Horizon.
// set ledgerNum to 0 to get the current ledger
func (c *Client) LoadLedger(ledgerNum uint64) (ledger Ledger, err error) {
	c.initHttpClient()
	var resp *http.Response
	if ledgerNum == 0 {
		resp, err = c.Client.Get(c.URL + "/ledgers?cursor=now&limit=1&order=desc")
		if err != nil {
			return
		}
		var ledgerList LedgerList
		err = decodeResponse(resp, &ledgerList)
		if err != nil || len(ledgerList.Embedded.Records) < 1 {
			return
		}
		ledger = ledgerList.Embedded.Records[0]
	} else {
		resp, err = c.Client.Get(fmt.Sprintf("%s/ledgers/%d", c.URL, ledgerNum))
		if err != nil {
			return
		}

		err = decodeResponse(resp, &ledger)
	}
	return
}

// SubmitTransaction submits a transaction to the network. err can be either error object or horizon.Error object.
func (c *Client) SubmitTransaction(transactionEnvelopeXdr string) (response TransactionSuccess, err error) {
	v := url.Values{}
	v.Set("tx", transactionEnvelopeXdr)

	c.initHttpClient()
	resp, err := c.Client.PostForm(c.URL+"/transactions", v)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &response)
	return
}

func (c *Client) initHttpClient() {
	c.clientInit.Do(func() {
		if c.Client == nil {
			c.Client = &http.Client{}
		}
	})
}

func decodeResponse(resp *http.Response, object interface{}) (err error) {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		horizonError := &Error{
			Response: resp,
		}
		decodeError := decoder.Decode(&horizonError.Problem)
		if decodeError != nil {
			return decodeError
		}
		return horizonError
	}

	err = decoder.Decode(&object)
	if err != nil {
		return
	}
	return
}
