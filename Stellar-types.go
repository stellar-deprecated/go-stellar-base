// Automatically generated from xdr/Stellar-types.x
// DO NOT EDIT or your changes may be overwritten

package stellarcore

type Uint512 [64]byte

type Uint256 [32]byte

type Uint64 uint64

type Int64 int64

type Uint32 uint32

type Int32 int32

type AccountId [32]byte

type Signature [64]byte

type Hash [32]byte

type Thresholds [4]byte

type SequenceNumber Uint64

type CurrencyType uint32

const (
	Native CurrencyType = iota
	Iso4217
)

type IsoCurrencyIssuer struct {
	CurrencyCode [4]byte
	Issuer       AccountId
}

type Currency struct {
	Type  CurrencyType
	isoCi *IsoCurrencyIssuer
}

type Price struct {
	N Int32
	D Int32
}
