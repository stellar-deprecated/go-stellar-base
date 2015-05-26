package stellarbase

import "github.com/stellar/go-stellar-base/xdr"

//go:generate rake xdr:update
//go:generate go fmt ./xdr

// AddressToAccountId converts the provided address into a xdr.AccountId
func AddressToAccountId(address string) (result xdr.AccountId, err error) {
	bytes, err := DecodeBase58Check(VersionByteAccountID, address)

	if err != nil {
		return
	}

	copy(result[:], bytes)
	return
}
