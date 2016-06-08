package horizon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/stellar/go-stellar-base/xdr"
)

// OrderDirection is a type that represents `order` param values
type OrderDirection string

const (
	// OrderAsc represents ascending order
	OrderAsc OrderDirection = "asc"
	// OrderDesc represents descending order
	OrderDesc OrderDirection = "desc"
)

// Cursor is a type that represents `cursor` param values
type Cursor string

// Limit is a type that represents `limit` param values
type Limit int

// PageParams contains all query params that allow fetching
// a specific results page. Helpful for loading next/previous page of results.
type PageParams struct {
	Order OrderDirection
	Cursor
	Limit
}

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{URL: "https://horizon-testnet.stellar.org"}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{URL: "https://horizon.stellar.org"}

// Error struct contains the problem returned by Horizon
type Error struct {
	Response *http.Response
	Problem  Problem
}

func (herror *Error) Error() string {
	return fmt.Sprintf("Horizon error: [%s] %s\nExtras: %s", herror.Problem.Title, herror.Problem.Detail, herror.Problem.Extras)
}

// HTTPClient contains http.Client methods used by horizon.Client.
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// Client struct contains data required to connect to Horizon instance
type Client struct {
	// URL of Horizon server to connect
	URL string
	// Will be populated with &http.Client when nil. If you want to configure your http.Client make sure Timeout is at least 10 seconds.
	Client HTTPClient
	// clientInit initializes http client once
	clientInit sync.Once
}

// SequenceForAccount implements build.SequenceProvider
func (c *Client) SequenceForAccount(
	accountID string,
) (xdr.SequenceNumber, error) {

	a, err := c.LoadAccount(accountID)
	if err != nil {
		return 0, err
	}

	seq, err := strconv.ParseUint(a.Sequence, 10, 64)
	if err != nil {
		return 0, err
	}

	return xdr.SequenceNumber(seq), nil
}

func (c *Client) initHTTPClient() {
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
