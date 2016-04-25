package build

import (
	"errors"
	"math"

	"github.com/shopspring/decimal"
	"github.com/stellar/go-stellar-base/amount"
	"github.com/stellar/go-stellar-base/xdr"
)

// CreateOffer creates a new offer
func CreateOffer(rate Rate, amount Amount) (result ManageOfferBuilder) {
	return ManageOffer(false, rate, amount)
}

// CreatePassiveOffer creates a new passive offer
func CreatePassiveOffer(rate Rate, amount Amount) (result ManageOfferBuilder) {
	return ManageOffer(true, rate, amount)
}

// UpdateOffer updates an existing offer
func UpdateOffer(rate Rate, amount Amount, offerID OfferID) (result ManageOfferBuilder) {
	return ManageOffer(false, rate, amount, offerID)
}

// DeleteOffer deletes an existing offer
func DeleteOffer(rate Rate, offerID OfferID) (result ManageOfferBuilder) {
	return ManageOffer(false, rate, Amount("0"), offerID)
}

// ManageOffer groups the creation of a new ManageOfferBuilder with a call to Mutate.
func ManageOffer(passiveOffer bool, muts ...interface{}) (result ManageOfferBuilder) {
	result.PassiveOffer = passiveOffer
	result.Mutate(muts...)
	return
}

// ManageOfferMutator is a interface that wraps the
// MutateManageOffer operation.  types may implement this interface to
// specify how they modify an xdr.ManageOfferOp object
type ManageOfferMutator interface {
	MutateManageOffer(interface{}) error
}

// ManageOfferBuilder represents a transaction that is being built.
type ManageOfferBuilder struct {
	PassiveOffer bool
	O            xdr.Operation
	MO           xdr.ManageOfferOp
	PO           xdr.CreatePassiveOfferOp
	Err          error
}

// Mutate applies the provided mutators to this builder's offer or operation.
func (b *ManageOfferBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case ManageOfferMutator:
			if b.PassiveOffer {
				err = mut.MutateManageOffer(&b.PO)
			} else {
				err = mut.MutateManageOffer(&b.MO)
			}
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

// MutateManageOffer for Amount sets the ManageOfferOp's Amount field
func (m Amount) MutateManageOffer(o interface{}) (err error) {
	switch o := o.(type) {
	default:
		err = errors.New("Unexpected operation type")
	case *xdr.ManageOfferOp:
		o.Amount, err = amount.Parse(string(m))
	case *xdr.CreatePassiveOfferOp:
		o.Amount, err = amount.Parse(string(m))
	}
	return
}

// MutateManageOffer for OfferID sets the ManageOfferOp's OfferID field
func (m OfferID) MutateManageOffer(o interface{}) (err error) {
	switch o := o.(type) {
	default:
		err = errors.New("Unexpected operation type")
	case *xdr.ManageOfferOp:
		o.OfferId = xdr.Uint64(m)
	}
	return
}

// MutateManageOffer for Rate sets the ManageOfferOp's selling, buying and price fields
func (m Rate) MutateManageOffer(o interface{}) (err error) {
	switch o := o.(type) {
	default:
		err = errors.New("Unexpected operation type")
	case *xdr.ManageOfferOp:
		o.Selling, err = m.Selling.ToXdrObject()
		if err != nil {
			return
		}

		o.Buying, err = m.Buying.ToXdrObject()
		if err != nil {
			return
		}

		o.Price, err = continuedFraction(m.Price)
	case *xdr.CreatePassiveOfferOp:
		o.Selling, err = m.Selling.ToXdrObject()
		if err != nil {
			return
		}

		o.Buying, err = m.Buying.ToXdrObject()
		if err != nil {
			return
		}

		o.Price, err = continuedFraction(m.Price)
	}
	return
}

// continuedFraction calculates and returns the best rational approximation of the given real number.
func continuedFraction(price Price) (xdrPrice xdr.Price, err error) {
	number, err := decimal.NewFromString(string(price))
	if err != nil {
		return
	}

	maxInt32 := decimal.New(int64(math.MaxInt32), 0)

	fractions := [][2]decimal.Decimal{
		{decimal.Decimal{}, decimal.New(1, 0)},
		{decimal.New(1, 0), decimal.Decimal{}},
	}

	i := 2
	for {
		if number.Cmp(maxInt32) == 1 {
			break
		}

		a := number.Floor()
		f := number.Sub(a)
		h := a.Mul(fractions[i-1][0]).Add(fractions[i-2][0])
		k := a.Mul(fractions[i-1][1]).Add(fractions[i-2][1])
		if h.Cmp(maxInt32) == 1 || k.Cmp(maxInt32) == 1 {
			break
		}

		fractions = append(fractions, [2]decimal.Decimal{h, k})
		if f.Cmp(decimal.Decimal{}) == 0 {
			break
		}
		number = decimal.New(1, 0).Div(f)
		i++
	}

	n, d := fractions[len(fractions)-1][0], fractions[len(fractions)-1][1]

	if n.Cmp(decimal.Decimal{}) == 0 || d.Cmp(decimal.Decimal{}) == 0 {
		return xdrPrice, errors.New("Couldn't find approximation")
	}

	return xdr.Price{
		N: xdr.Int32(n.IntPart()),
		D: xdr.Int32(d.IntPart()),
	}, nil
}
