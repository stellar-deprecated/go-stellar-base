package horizon

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sync"

	res "github.com/stellar/go-stellar-base/horizon/response"
)

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{URL: "https://horizon-testnet.stellar.org"}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{URL: "https://horizon.stellar.org"}

// HorizonError struct contains error and problem returned by Horizon
type HorizonError struct {
	Err     error
	Problem *res.Problem
}

func (herror *HorizonError) Error() string {
	return herror.Err.Error()
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

// LoadAccount loads the account state from horizon. horizonError is HorizonError struct that implements error interface.
func (c *Client) LoadAccount(accountId string) (account res.Account, horizonError error) {
	c.clientInit.Do(c.initHttpClient)
	resp, err := c.Client.Get(c.URL + "/accounts/" + accountId)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	horizonError = decodeResponse(resp, &account)
	return
}

// SubmitTransaction submits a transaction to the network. horizonError is HorizonError struct that implements error interface.
func (c *Client) SubmitTransaction(transactionEnvelopeXdr string) (response res.TransactionSuccess, horizonError error) {
	v := url.Values{}
	v.Set("tx", transactionEnvelopeXdr)

	c.clientInit.Do(c.initHttpClient)
	resp, err := c.Client.PostForm(c.URL+"/transactions", v)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	horizonError = decodeResponse(resp, &response)
	return
}

func (c *Client) initHttpClient() {
	if c.Client == nil {
		c.Client = &http.Client{}
	}
}

func decodeResponse(resp *http.Response, object interface{}) (horizonError error) {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		horizonError = &HorizonError{
			Err:     errors.New("Error response"),
			Problem: &res.Problem{},
		}
		err := decoder.Decode(&horizonError.(*HorizonError).Problem)
		if err != nil {
			horizonError = &HorizonError{Err: err}
		}
		return
	}

	err := decoder.Decode(&object)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}
	return
}
