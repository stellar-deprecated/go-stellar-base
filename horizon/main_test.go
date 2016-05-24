package horizon

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stellar/go-stellar-base/build"
)

func TestHorizon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: github.com/stellar/go-stellar-base/horizon")
}

var _ build.SequenceProvider = TestHorizonClient

var _ = Describe("Horizon", func() {
	Describe("initHttpClient", func() {
		It("does not run into race condition", func() {
			// Race condition should be detected by race-detector:
			// http://blog.golang.org/race-detector
			init := func() {
				DefaultTestNetClient.initHttpClient()
			}
			go init()
			go init()
		})
	})

	Describe("LoadAccount", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHttpClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(accountResponse)),
				},
			}

			account, err := TestHorizonClient.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).To(BeNil())
			Expect(account.ID).To(Equal("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"))
			Expect(account.PT).To(Equal("1"))
			Expect(account.GetNativeBalance()).To(Equal("948522307.6146000"))
			Expect(account.Data["foo"]).To(Equal("+xbxLlS9Exgb4n6tWSK6ruUmejywOykOHw1zCrotEgtNHeBzykVmdMhPipUOI839q1tybb9NUkrsteMoJas1/w=="))
		})

		It("failure response", func() {
			TestHorizonClient.Client = &TestHttpClient{
				Response: http.Response{
					StatusCode: 404,
					Body:       ioutil.NopCloser(bytes.NewBufferString(notFoundResponse)),
				},
			}

			_, err := TestHorizonClient.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("Horizon error"))
			horizonError, ok := err.(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Resource Missing"))
		})

		It("connection error", func() {
			TestHorizonClient.Client = &TestHttpClient{
				Error: errors.New("http.Client error"),
			}

			_, err := TestHorizonClient.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("http.Client error"))
			_, ok := err.(*Error)
			Expect(ok).To(BeFalse())
		})
	})

	Describe("SubmitTransaction", func() {
		var tx = "AAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAZAAT3TUAAAAwAAAAAAAAAAAAAAABAAAAAAAAAAMAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAG/dhGXAAAAQLuStfImg0OeeGAQmvLkJSZ1MPSkCzCYNbGqX5oYNuuOqZ5SmWhEsC7uOD9ha4V7KengiwNlc0oMNqBVo22S7gk="

		It("success response", func() {
			TestHorizonClient.Client = &TestHttpClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(submitResponse)),
				},
			}

			account, err := TestHorizonClient.SubmitTransaction(tx)
			Expect(err).To(BeNil())
			Expect(account.Ledger).To(Equal(int32(3128812)))
		})

		It("failure response", func() {
			TestHorizonClient.Client = &TestHttpClient{
				Response: http.Response{
					StatusCode: 400,
					Body:       ioutil.NopCloser(bytes.NewBufferString(transactionFailure)),
				},
			}

			_, err := TestHorizonClient.SubmitTransaction(tx)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("Horizon error"))
			horizonError, ok := err.(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Transaction Failed"))
		})

		It("connection error", func() {
			TestHorizonClient.Client = &TestHttpClient{
				Error: errors.New("http.Client error"),
			}

			_, err := TestHorizonClient.SubmitTransaction(tx)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("http.Client error"))
			_, ok := err.(*Error)
			Expect(ok).To(BeFalse())
		})
	})
})

var TestHorizonClient = &Client{
	Client: &TestHttpClient{},
}

type TestHttpClient struct {
	Response http.Response
	Error    error
}

func (tc *TestHttpClient) Get(url string) (*http.Response, error) {
	return &tc.Response, tc.Error
}

func (tc *TestHttpClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return &tc.Response, tc.Error
}

var accountResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
    },
    "transactions": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/effects{?cursor,limit,order}",
      "templated": true
    },
    "offers": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/Offers{?cursor,limit,order}",
      "templated": true
    }
  },
  "id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
  "paging_token": "1",
  "account_id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
  "sequence": "7384",
  "subentry_count": 0,
  "thresholds": {
    "low_threshold": 0,
    "med_threshold": 0,
    "high_threshold": 0
  },
  "flags": {
    "auth_required": false,
    "auth_revocable": false
  },
  "balances": [
    {
      "balance": "948522307.6146000",
      "asset_type": "native"
    }
  ],
  "signers": [
    {
      "public_key": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
      "weight": 1
    }
  ],
  "data": {
	"foo": "+xbxLlS9Exgb4n6tWSK6ruUmejywOykOHw1zCrotEgtNHeBzykVmdMhPipUOI839q1tybb9NUkrsteMoJas1/w=="
  }
}`

var notFoundResponse = `{
  "type": "https://stellar.org/horizon-errors/not_found",
  "title": "Resource Missing",
  "status": 404,
  "detail": "The resource at the url requested was not found.  This is usually occurs for one of two reasons:  The url requested is not valid, or no data in our database could be found with the parameters provided.",
  "instance": "horizon-live-001/61KdRW8tKi-18408110"
}`

var submitResponse = `{
  "_links": {
    "transaction": {
      "href": "https://horizon-testnet.stellar.org/transactions/ee14b93fcd31d4cfe835b941a0a8744e23a6677097db1fafe0552d8657bed940"
    }
  },
  "hash": "ee14b93fcd31d4cfe835b941a0a8744e23a6677097db1fafe0552d8657bed940",
  "ledger": 3128812,
  "envelope_xdr": "AAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAZAAT3TUAAAAwAAAAAAAAAAAAAAABAAAAAAAAAAMAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAG/dhGXAAAAQLuStfImg0OeeGAQmvLkJSZ1MPSkCzCYNbGqX5oYNuuOqZ5SmWhEsC7uOD9ha4V7KengiwNlc0oMNqBVo22S7gk=",
  "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAADAAAAAAAAAAAAAAAAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAAAAPEAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAA==",
  "result_meta_xdr": "AAAAAAAAAAEAAAACAAAAAAAvoHwAAAACAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAAAAPEAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAEAL6B8AAAAAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAABZ9zvNAABPdNQAAADAAAAAEAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA=="
}`

var transactionFailure = `{
  "type": "https://stellar.org/horizon-errors/transaction_failed",
  "title": "Transaction Failed",
  "status": 400,
  "detail": "The transaction failed when submitted to the stellar network. The extras.result_codes field on this response contains further details.  Descriptions of each code can be found at: https://www.stellar.org/developers/learn/concepts/list-of-operations.html",
  "instance": "horizon-testnet-001.prd.stellar001.internal.stellar-ops.com/4elYz2fHhC-528285",
  "extras": {
    "envelope_xdr": "AAAAAKpmDL6Z4hvZmkTBkYpHftan4ogzTaO4XTB7joLgQnYYAAAAZAAAAAAABeoyAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAABAAAAAD3sEVVGZGi/NoC3ta/8f/YZKMzyi9ZJpOi0H47x7IqYAAAAAAAAAAAF9eEAAAAAAAAAAAA=",
    "result_codes": {
      "transaction": "tx_no_source_account"
    },
    "result_xdr": "AAAAAAAAAAD////4AAAAAA=="
  }
}`
