// This file contains response structs from horizon
package horizon

import (
	u "net/url"
	"strconv"
	"time"
)

type Problem struct {
	Type     string                 `json:"type"`
	Title    string                 `json:"title"`
	Status   int                    `json:"status"`
	Detail   string                 `json:"detail,omitempty"`
	Instance string                 `json:"instance,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

type Account struct {
	Links struct {
		Self         Link `json:"self"`
		Transactions Link `json:"transactions"`
		Operations   Link `json:"operations"`
		Payments     Link `json:"payments"`
		Effects      Link `json:"effects"`
		Offers       Link `json:"offers"`
	} `json:"_links"`

	HistoryAccount
	Sequence             string            `json:"sequence"`
	SubentryCount        int32             `json:"subentry_count"`
	InflationDestination string            `json:"inflation_destination,omitempty"`
	HomeDomain           string            `json:"home_domain,omitempty"`
	Thresholds           AccountThresholds `json:"thresholds"`
	Flags                AccountFlags      `json:"flags"`
	Balances             []Balance         `json:"balances"`
	Signers              []Signer          `json:"signers"`
	Data                 map[string]string `json:"data"`
}

func (a Account) GetNativeBalance() string {
	for _, balance := range a.Balances {
		if balance.Asset.Type == "native" {
			return balance.Balance
		}
	}

	return "0"
}

type AccountFlags struct {
	AuthRequired  bool `json:"auth_required"`
	AuthRevocable bool `json:"auth_revocable"`
}

type AccountThresholds struct {
	LowThreshold  byte `json:"low_threshold"`
	MedThreshold  byte `json:"med_threshold"`
	HighThreshold byte `json:"high_threshold"`
}

type AccountsPage struct {
	Links    `json:"_links"`
	Embedded struct {
		Records []HistoryAccount `json:"records"`
	} `json:"_embedded"`
}

type Asset struct {
	Type   string `json:"asset_type"`
	Code   string `json:"asset_code,omitempty"`
	Issuer string `json:"asset_issuer,omitempty"`
}

type Balance struct {
	Balance string `json:"balance"`
	Limit   string `json:"limit,omitempty"`
	Asset
}

type HistoryAccount struct {
	ID        string `json:"id"`
	PT        string `json:"paging_token"`
	AccountID string `json:"account_id"`
}

// Ledger represents a single closed ledger
type Ledger struct {
	ID               string    `json:"id"`
	PT               string    `json:"paging_token"`
	Hash             string    `json:"hash"`
	PrevHash         string    `json:"prev_hash,omitempty"`
	Sequence         int32     `json:"sequence"`
	TransactionCount int32     `json:"transaction_count"`
	OperationCount   int32     `json:"operation_count"`
	ClosedAt         time.Time `json:"closed_at"`
	TotalCoins       string    `json:"total_coins"`
	FeePool          string    `json:"fee_pool"`
	BaseFee          int32     `json:"base_fee"`
	BaseReserve      string    `json:"base_reserve"`
	MaxTxSetSize     int32     `json:"max_tx_set_size"`
}

type LedgersPage struct {
	Links    `json:"_links"`
	Embedded struct {
		Records []Ledger `json:"records"`
	} `json:"_embedded"`
}

type Link struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated,omitempty"`
}

type Links struct {
	Self Link `json:"self"`
	Next Link `json:"next"`
	Prev Link `json:"prev"`
}

func urlToPageParams(link string) (params PageParams, err error) {
	url, err := u.Parse(link)
	if err != nil {
		return
	}

	values, err := u.ParseQuery(url.RawQuery)
	if err != nil {
		return
	}

	limit, err := strconv.Atoi(values.Get("limit"))
	if err != nil {
		return
	}

	params.Order = OrderDirection(values.Get("order"))
	params.Cursor = Cursor(values.Get("cursor"))
	params.Limit = Limit(limit)
	return
}

func (links Links) GetNextPageParams() (PageParams, error) {
	return urlToPageParams(links.Next.Href)
}

func (links Links) GetPrevPageParams() (PageParams, error) {
	return urlToPageParams(links.Prev.Href)
}

// Transaction represents a single, successful transaction
type Transaction struct {
	ID              string    `json:"id"`
	PT              string    `json:"paging_token"`
	Hash            string    `json:"hash"`
	Ledger          int32     `json:"ledger"`
	LedgerCloseTime time.Time `json:"created_at"`
	Account         string    `json:"source_account"`
	AccountSequence string    `json:"source_account_sequence"`
	FeePaid         int32     `json:"fee_paid"`
	OperationCount  int32     `json:"operation_count"`
	EnvelopeXdr     string    `json:"envelope_xdr"`
	ResultXdr       string    `json:"result_xdr"`
	ResultMetaXdr   string    `json:"result_meta_xdr"`
	FeeMetaXdr      string    `json:"fee_meta_xdr"`
	MemoType        string    `json:"memo_type"`
	Memo            string    `json:"memo,omitempty"`
	Signatures      []string  `json:"signatures"`
	ValidAfter      string    `json:"valid_after,omitempty"`
	ValidBefore     string    `json:"valid_before,omitempty"`
}

type TransactionsPage struct {
	Links    `json:"_links"`
	Embedded struct {
		Records []Transaction `json:"records"`
	} `json:"_embedded"`
}

type TransactionSuccess struct {
	Links struct {
		Transaction Link `json:"transaction"`
	} `json:"_links"`
	Hash   string `json:"hash"`
	Ledger int32  `json:"ledger"`
	Env    string `json:"envelope_xdr"`
	Result string `json:"result_xdr"`
	Meta   string `json:"result_meta_xdr"`
}

type Signer struct {
	PublicKey string `json:"public_key"`
	Weight    int32  `json:"weight"`
}
