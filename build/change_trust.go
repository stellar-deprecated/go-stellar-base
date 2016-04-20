package build

import (
	"errors"

	"github.com/stellar/go-stellar-base/amount"
	"github.com/stellar/go-stellar-base/xdr"
)

// ChangeTrust groups the creation of a new ChangeTrustBuilder with a call to Mutate.
func ChangeTrust(muts ...interface{}) (result ChangeTrustBuilder) {
	result.Mutate(muts...)
	return
}

// ChangeTrustMutator is a interface that wraps the
// MutatePayment operation.  types may implement this interface to
// specify how they modify an xdr.ChangeTrustOp object
type ChangeTrustMutator interface {
	MutateChangeTrust(*xdr.ChangeTrustOp) error
}

// ChangeTrustBuilder represents a transaction that is being built.
type ChangeTrustBuilder struct {
	O   xdr.Operation
	CT  xdr.ChangeTrustOp
	Err error
}

// Mutate applies the provided mutators to this builder's payment or operation.
func (b *ChangeTrustBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case ChangeTrustMutator:
			err = mut.MutateChangeTrust(&b.CT)
		case OperationMutator:
			err = mut.MutateOperation(&b.O)
		default:
			err = errors.New("Mutator type not allowed")
		}

		if err != nil {
			b.Err = err
			return
		}
	}
}

// MutateChangeTrust for Asset sets the ChangeTrustOp's Line field
func (m Asset) MutateChangeTrust(o *xdr.ChangeTrustOp) (err error) {
	if m.Native {
		return errors.New("Native asset not allowed")
	}

	o.Line, err = createAlphaNumAsset(m.Code, m.Issuer)
	return
}

// MutateChangeTrust for Limit sets the ChangeTrustOp's Limit field
func (m Limit) MutateChangeTrust(o *xdr.ChangeTrustOp) (err error) {
	o.Limit, err = amount.Parse(m.Amount)
	return
}
