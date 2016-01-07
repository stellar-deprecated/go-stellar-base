package build

import (
	"errors"
	"github.com/stellar/go-stellar-base/xdr"
)

// AllowTrust groups the creation of a new AllowTrustBuilder with a call to Mutate.
func AllowTrust(muts ...interface{}) (result AllowTrustBuilder) {
	result.Mutate(muts...)
	return
}

// AllowTrustMutator is a interface that wraps the
// MutatePayment operation.  types may implement this interface to
// specify how they modify an xdr.AllowTrustOp object
type AllowTrustMutator interface {
	MutateAllowTrust(*xdr.AllowTrustOp) error
}

// AllowTrustBuilder represents a transaction that is being built.
type AllowTrustBuilder struct {
	O   xdr.Operation
	AT  xdr.AllowTrustOp
	Err error
}

// Mutate applies the provided mutators to this builder's payment or operation.
func (b *AllowTrustBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case AllowTrustMutator:
			err = mut.MutateAllowTrust(&b.AT)
		case OperationMutator:
			err = mut.MutateOperation(&b.O)
		}

		if err != nil {
			b.Err = err
			return
		}
	}
}

// MutateAllowTrust for Authorize sets the AllowTrustOp's Authorize field
func (m Authorize) MutateAllowTrust(o *xdr.AllowTrustOp) {
	o.Authorize = m.Value
}

// MutateAllowTrust for Asset sets the AllowTrustOp's Asset field
func (m AllowTrustAsset) MutateAllowTrust(o *xdr.AllowTrustOp) (err error) {
	var assetType xdr.AssetType
	var code []byte
	length := len(m.Code)

	switch {
	case length >= 1 && length <= 4:
		assetType = xdr.AssetTypeAssetTypeCreditAlphanum4
		code = make([]byte, 4)
	case length >= 5 && length <= 12:
		assetType = xdr.AssetTypeAssetTypeCreditAlphanum12
		code = make([]byte, 12)
	default:
		return errors.New("Asset code length is invalid")
	}

	byteArray := []byte(m.Code)
	copy(code, byteArray[:length])

	o.Asset, err = xdr.NewAllowTrustOpAsset(assetType, code)
	return
}

// MutateAllowTrust for Trustor sets the AllowTrustOp's Trustor field
func (m Trustor) MutateAllowTrust(o *xdr.AllowTrustOp) error {
	return setAccountId(m.Address, &o.Trustor)
}
