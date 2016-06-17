package horizon

import (
	"bytes"
	"errors"
	// "fmt"
	"io/ioutil"
	"log"
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

// ExampleAccountsPage shows how to load consecutive pages of accounts.
func ExampleAccountsPage() {
	accounts, err := DefaultTestNetClient.LoadAccounts(Limit(100), OrderDesc)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("First Page:")
	for i, account := range accounts.Embedded.Records {
		log.Println(i, account.ID)
	}

	accounts, err = DefaultTestNetClient.LoadAccounts(accounts.Links.GetNextPageParams())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Second Page:")
	for i, account := range accounts.Embedded.Records {
		log.Println(i, account.ID)
	}
}

var _ build.SequenceProvider = TestHorizonClient

var _ = Describe("Horizon", func() {
	Describe("initHTTPClient", func() {
		It("does not run into race condition", func() {
			// Race condition should be detected by race-detector:
			// http://blog.golang.org/race-detector
			init := func() {
				DefaultTestNetClient.initHTTPClient()
			}
			go init()
			go init()
		})
	})

	Describe("LoadAccount", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
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
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 404,
					Body:       ioutil.NopCloser(bytes.NewBufferString(notFoundResponse)),
				},
			}

			_, err := TestHorizonClient.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).NotTo(BeNil())
			horizonError, ok := err.(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Resource Missing"))
		})

		It("connection error", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Error: errors.New("http.Client error"),
			}

			_, err := TestHorizonClient.LoadAccount("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("http.Client error"))
			_, ok := err.(*Error)
			Expect(ok).To(BeFalse())
		})
	})

	Describe("LoadAccounts", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(accountsResponse)),
				},
			}

			accounts, err := TestHorizonClient.LoadAccounts()
			Expect(err).To(BeNil())

			nextParams, err := accounts.Links.GetNextPageParams()
			Expect(err).To(BeNil())
			Expect(nextParams.Order).To(Equal(OrderDirection("asc")))
			Expect(nextParams.Limit).To(Equal(Limit(10)))
			Expect(nextParams.Cursor).To(Equal(Cursor("8388071133185")))

			prevParams, err := accounts.Links.GetPrevPageParams()
			Expect(err).To(BeNil())
			Expect(prevParams.Order).To(Equal(OrderDirection("desc")))
			Expect(prevParams.Limit).To(Equal(Limit(10)))
			Expect(prevParams.Cursor).To(Equal(Cursor("1")))

			Expect(accounts.Embedded.Records[0].ID).To(Equal("GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"))
			Expect(accounts.Embedded.Records[0].PT).To(Equal("1"))
			Expect(accounts.Embedded.Records[9].ID).To(Equal("GBLZWJINHGQ4YLBCQVVI6EGUH3OH63KPA2RMBF6RNCPS2IF5GCGF4AZO"))
			Expect(accounts.Embedded.Records[9].PT).To(Equal("8388071133185"))
		})
	})

	Describe("LoadEffects", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(effectsResponse)),
				},
			}

			effects, err := TestHorizonClient.LoadEffects()
			Expect(err).To(BeNil())

			Expect(effects.Embedded.Records[0].ID).To(Equal("0005987652562063361-0000000004"))
			Expect(effects.Embedded.Records[0].Type).To(Equal("trade"))
			Expect(effects.Embedded.Records[0].Trade.Seller).To(Equal("GC3AMRZPOWP2VA4JA3CNV23G3KW7FU6GRBYIGJC5HE5ZRET6ILBPP7TY"))
			Expect(effects.Embedded.Records[0].Trade.SoldAssetCode).To(Equal("ZAR"))

			Expect(effects.Embedded.Records[1].ID).To(Equal("0005987652562063361-0000000003"))
			Expect(effects.Embedded.Records[1].Type).To(Equal("trade"))
			Expect(effects.Embedded.Records[1].Trade.Seller).To(Equal("GBLZWJINHGQ4YLBCQVVI6EGUH3OH63KPA2RMBF6RNCPS2IF5GCGF4AZO"))

			Expect(effects.Embedded.Records[2].ID).To(Equal("0005987652562063361-0000000002"))
			Expect(effects.Embedded.Records[2].Type).To(Equal("account_debited"))
			Expect(effects.Embedded.Records[2].AccountDebited.Amount).To(Equal("10.0196334"))
			Expect(effects.Embedded.Records[2].AccountDebited.Asset.Type).To(Equal("credit_alphanum4"))
			Expect(effects.Embedded.Records[2].AccountDebited.Asset.Code).To(Equal("USD"))
			Expect(effects.Embedded.Records[2].AccountDebited.Asset.Issuer).To(Equal("GBWLLK6C5HF6PZQ7HKSWHCXH355AFF55SPEYYM7BMDHKO2PIT66QP7GJ"))
		})
	})

	Describe("LoadLedger", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(ledgerResponse)),
				},
			}

			ledger, err := TestHorizonClient.LoadLedger(1)
			Expect(err).To(BeNil())
			Expect(ledger.ID).To(Equal("63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99"))
			Expect(ledger.PT).To(Equal("4294967296"))
			Expect(ledger.TotalCoins).To(Equal("100000000000.0000000"))
		})
	})

	Describe("LoadLedgers", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(ledgersResponse)),
				},
			}

			ledgers, err := TestHorizonClient.LoadLedgers()
			Expect(err).To(BeNil())
			Expect(ledgers.Embedded.Records[0].ID).To(Equal("63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99"))
			Expect(ledgers.Embedded.Records[0].PT).To(Equal("4294967296"))
			Expect(ledgers.Embedded.Records[0].TotalCoins).To(Equal("100000000000.0000000"))
			Expect(ledgers.Embedded.Records[9].ID).To(Equal("651a974af2c39bed94975269aa3f762d4e311728b8075a7a394fadf75db4a2e7"))
			Expect(ledgers.Embedded.Records[9].PT).To(Equal("42949672960"))
			Expect(ledgers.Embedded.Records[9].Sequence).To(Equal(int32(10)))
		})
	})

	Describe("LoadOperation", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(operationResponse)),
				},
			}

			operation, err := TestHorizonClient.LoadOperation("6299132180303873")
			Expect(err).To(BeNil())
			Expect(operation.ID).To(Equal("6299132180303873"))
			Expect(operation.PT).To(Equal("6299132180303873"))
			Expect(operation.Type).To(Equal("path_payment"))

			Expect(operation.Payment).To(BeNil())
			Expect(operation.PathPayment).NotTo(BeNil())
			Expect(operation.PathPayment.From).To(Equal("GA4UBSTU6JCBV6HFXPHMTES5V2SH4RHGT7P2WAZ3KEJK5XNNLZFHQ56O"))
			Expect(operation.PathPayment.To).To(Equal("GAMAXXPWISK3WSPDKB2QPOKAVB7DOPDCDKY7IJPVFJOGMTEOV7F57DFX"))
			Expect(operation.PathPayment.Amount).To(Equal("1.0000000"))
			Expect(operation.PathPayment.Asset.Type).To(Equal("credit_alphanum4"))
			Expect(operation.PathPayment.Asset.Code).To(Equal("ZAR"))
			Expect(operation.PathPayment.Asset.Issuer).To(Equal("GCU3KWOEHGFJL2FTEHFPY6DIIBX5HKVZABV35YI6JNVNYDHJSGPHR22G"))
			Expect(operation.PathPayment.SourceMax).To(Equal("20.0000000"))
			Expect(operation.PathPayment.SourceAssetType).To(Equal("credit_alphanum4"))
			Expect(operation.PathPayment.SourceAssetCode).To(Equal("USD"))
			Expect(operation.PathPayment.SourceAssetIssuer).To(Equal("GDGCTUHHOY7CX6LLY3MP3YAYJPSP3R7HXK4TFCX3JWT33K2AUG6AYME5"))
		})
	})

	Describe("LoadOperations", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(operationsResponse)),
				},
			}

			operations, err := TestHorizonClient.LoadOperations()
			Expect(err).To(BeNil())

			nextParams, err := operations.Links.GetNextPageParams()
			Expect(err).To(BeNil())
			Expect(nextParams.Order).To(Equal(OrderDirection("desc")))
			Expect(nextParams.Limit).To(Equal(Limit(3)))
			Expect(nextParams.Cursor).To(Equal(Cursor("6663890867851265")))

			Expect(operations.Embedded.Records[0].ID).To(Equal("6664032601772033"))
			Expect(operations.Embedded.Records[0].PT).To(Equal("6664032601772033"))
			Expect(operations.Embedded.Records[0].Type).To(Equal("change_trust"))
			Expect(operations.Embedded.Records[0].ChangeTrust).NotTo(BeNil())

			Expect(operations.Embedded.Records[2].ID).To(Equal("6663890867851265"))
			Expect(operations.Embedded.Records[2].PT).To(Equal("6663890867851265"))
			Expect(operations.Embedded.Records[2].Type).To(Equal("change_trust"))
			Expect(operations.Embedded.Records[2].ChangeTrust).NotTo(BeNil())
		})
	})

	Describe("LoadTransaction", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(transactionResponse)),
				},
			}

			ledger, err := TestHorizonClient.LoadTransaction("e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1")
			Expect(err).To(BeNil())
			Expect(ledger.ID).To(Equal("e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1"))
			Expect(ledger.PT).To(Equal("4672924422144"))
			Expect(ledger.Ledger).To(Equal(int32(1088)))
		})
	})

	Describe("LoadTransactions", func() {
		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(transactionsResponse)),
				},
			}

			transactions, err := TestHorizonClient.LoadTransactions()
			Expect(err).To(BeNil())
			Expect(transactions.Embedded.Records[0].ID).To(Equal("e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1"))
			Expect(transactions.Embedded.Records[0].PT).To(Equal("4672924422144"))
			Expect(transactions.Embedded.Records[0].Ledger).To(Equal(int32(1088)))
			Expect(transactions.Embedded.Records[2].ID).To(Equal("02a142294acde3e8920fc1e3879b34763df56b9c387a6dcf16d613b21cb9b87a"))
			Expect(transactions.Embedded.Records[2].PT).To(Equal("6859062775808"))
			Expect(transactions.Embedded.Records[2].Ledger).To(Equal(int32(1597)))
		})
	})

	Describe("SubmitTransaction", func() {
		var tx = "AAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAZAAT3TUAAAAwAAAAAAAAAAAAAAABAAAAAAAAAAMAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAG/dhGXAAAAQLuStfImg0OeeGAQmvLkJSZ1MPSkCzCYNbGqX5oYNuuOqZ5SmWhEsC7uOD9ha4V7KengiwNlc0oMNqBVo22S7gk="

		It("success response", func() {
			TestHorizonClient.Client = &TestHTTPClient{
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
			TestHorizonClient.Client = &TestHTTPClient{
				Response: http.Response{
					StatusCode: 400,
					Body:       ioutil.NopCloser(bytes.NewBufferString(transactionFailure)),
				},
			}

			_, err := TestHorizonClient.SubmitTransaction(tx)
			Expect(err).NotTo(BeNil())
			horizonError, ok := err.(*Error)
			Expect(ok).To(BeTrue())
			Expect(horizonError.Problem.Title).To(Equal("Transaction Failed"))
		})

		It("connection error", func() {
			TestHorizonClient.Client = &TestHTTPClient{
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
	Client: &TestHTTPClient{},
}

type TestHTTPClient struct {
	Response http.Response
	Error    error
}

func (tc *TestHTTPClient) Get(url string) (*http.Response, error) {
	return &tc.Response, tc.Error
}

func (tc *TestHTTPClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return &tc.Response, tc.Error
}
