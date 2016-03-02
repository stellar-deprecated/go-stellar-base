package xdr

import (
	"errors"
	"fmt"

	"github.com/stellar/go-stellar-base/strkey"
)

// SetAddress modifies the receiver, setting it's value to the AccountId form
// of the provided address.
func (aid *AccountId) SetAddress(address string) error {
	if aid == nil {
		return nil
	}

	raw, err := strkey.Decode(strkey.VersionByteAccountID, address)
	if err != nil {
		return err
	}

	if len(raw) != 32 {
		return errors.New("invalid address")
	}

	var ui Uint256
	copy(ui[:], raw)

	*aid, err = NewAccountId(CryptoKeyTypeKeyTypeEd25519, ui)

	return err
}

// Address returns the strkey encoded form of this AccountId.  This method will
// panic if the accountid is backed by a public key of an unknown type.
func (aid *AccountId) Address() string {
	if aid == nil {
		return ""
	}

	switch aid.Type {
	case CryptoKeyTypeKeyTypeEd25519:
		ed := aid.MustEd25519()
		raw := make([]byte, 32)
		copy(raw, ed[:])
		return strkey.MustEncode(strkey.VersionByteAccountID, raw)
	default:
		panic(fmt.Errorf("Unknown account id type: %v", aid.Type))
	}
}
