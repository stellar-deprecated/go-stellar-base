// Automatically generated from xdr/Stellar-ledger-entries.x
// DO NOT EDIT or your changes may be overwritten

package stellarcore

type LedgerEntryType uint32

const (
	Account LedgerEntryType = iota
	Trustline
	Offer
)

type Signer struct {
	PubKey Uint256
	Weight Uint32
}

type KeyValue struct {
	Key   Uint32
	Value []byte
}

type AccountFlags uint32

const (
	AuthRequiredFlag AccountFlags = iota
)

type AccountEntry struct {
	AccountId     Uint256
	Balance       Int64
	SeqNum        SequenceNumber
	NumSubEntries Uint32
	InflationDest *Uint256
	Thresholds    [4]byte
	Signers       []Signer
	Data          []KeyValue
	Flags         Uint32
}

type TrustLineEntry struct {
	AccountId  Uint256
	Currency   Currency
	Limit      Int64
	Balance    Int64
	Authorized bool
}

type OfferEntry struct {
	AccountId Uint256
	OfferId   Uint64
	TakerGets Currency
	TakerPays Currency
	Amount    Int64
	Price     Price
	Flags     Int32
}

type LedgerEntry struct {
	Type      LedgerEntryType
	account   *AccountEntry
	trustLine *TrustLineEntry
	offer     *OfferEntry
}
