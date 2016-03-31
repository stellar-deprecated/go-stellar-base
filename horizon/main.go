package horizon

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/stellar/go-stellar-base/horizon/structs"
)

const TIMEOUT = 30 * time.Second

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{URL: "https://horizon-testnet.stellar.org"}
// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{URL: "https://horizon.stellar.org"}

// HorizonError struct contains error and problem returned by Horizon
type HorizonError struct {
	Err     error
	Problem structs.Problem
}

type HorizonHttpClient interface {
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// Client struct contains data required to connect to Horizon instance
type Client struct {
	URL    string
	client HorizonHttpClient
}

// LoadAccount loads the account state from horizon
func (c *Client) LoadAccount(accountId string) (account structs.Account, horizonError *HorizonError) {
	httpClient := c.getHttpClient()
	resp, err := httpClient.Get(c.URL + "/accounts/" + accountId)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New("Error response")
		horizonError = &HorizonError{Err: err}
		json.Unmarshal(body, &horizonError.Problem)
		return
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	return
}

// SubmitTransaction submits a transaction to the network
func (c *Client) SubmitTransaction(transactionEnvelopeXdr string) (response structs.TransactionSuccess, horizonError *HorizonError) {
	v := url.Values{}
	v.Set("tx", transactionEnvelopeXdr)

	httpClient := c.getHttpClient()
	resp, err := httpClient.PostForm(c.URL+"/transactions", v)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New("Error response")
		horizonError = &HorizonError{Err: err}
		json.Unmarshal(body, &horizonError.Problem)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		horizonError = &HorizonError{Err: err}
		return
	}

	return
}

func (c *Client) getHttpClient() HorizonHttpClient {
	if c.client == nil {
		c.client = &http.Client{
			Timeout: TIMEOUT,
		}
	}

	return c.client
}
