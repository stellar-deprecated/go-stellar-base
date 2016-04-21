// This file contains response structs from horizon
package horizon

import (
	"encoding/json"
	"strconv"
)

type Problem struct {
	Type     string                 `json:"type"`
	Title    string                 `json:"title"`
	Status   int                    `json:"status"`
	Detail   string                 `json:"detail,omitempty"`
	Instance string                 `json:"instance,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

type Amount int64

func (s *Amount) UnmarshalJSON(d []byte) error {
	var raw string

	err := json.Unmarshal(d, &raw)
	if err != nil {
		return err
	}

	parsed, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return err
	}

	*s = Amount(parsed)
	return nil
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
}

func (a Account) GetNativeBalance() int64 {
	for _, balance := range a.Balances {
		if balance.Asset.Type == "native" {
			return int64(balance.Balance)
		}
	}

	return 0
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

type Asset struct {
	Type   string `json:"asset_type"`
	Code   string `json:"asset_code,omitempty"`
	Issuer string `json:"asset_issuer,omitempty"`
}

type Balance struct {
	Balance Amount `json:"balance"`
	Limit   string `json:"limit,omitempty"`
	Asset
}

type HistoryAccount struct {
	ID        string `json:"id"`
	PT        string `json:"paging_token"`
	AccountID string `json:"account_id"`
}

type Link struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated,omitempty"`
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

type Ledger struct {
	Links struct {
		Self         Link `json:"self"`
		Transactions Link `json:"transactions"`
		Operations   Link `json:"operations"`
		Payments     Link `json:"payments"`
		Effects      Link `json:"effects"`
	} `json:"_links"`

	ID               string `json:"id"`
	PagingToken      string `json:"paging_token"`
	Hash             string `json:"hash"`
	PrevHash         string `json:"prev_hash"`
	Sequence         uint32 `json:"sequence"`
	TransactionCount uint64 `json:"transaction_count"`
	OperationCount   uint64 `json:"operation_count"`
	ClosedAt         string `json:"closed_at"`
	TotalCoins       Amount `json:"total_coins"`
	FeePool          Amount `json:"fee_pool"`
	BaseFee          uint32 `json:"base_fee"`
	BaseReserve      Amount `json:"base_reserve"`
	MaxTxSetSize     uint64 `json:"max_tx_set_size"`
}

type LedgerList struct {
	Links struct {
		Self Link `json:"self"`
		Next Link `json:"next"`
		Prev Link `json:"prev"`
	} `json:"_links"`

	Embedded struct {
		Records []Ledger `json:"records"`
	} `json:"_embedded"`
}
