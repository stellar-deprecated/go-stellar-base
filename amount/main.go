// Package amount provides utilities for converting numbers to/from
// the format used internally to stellar-core.
//
// stellar-core represents asset "amounts" as 64-bit integers, but to enable
// fractional units of an asset, horizon, the client-libraries and other built
// on top of stellar-core use a convention, encoding amounts as a string of
// decimal digits with up to seven digits of precision in the fractional
// portion. For example, an amount shown as "101.001" in horizon would be
// represented in stellar-core as 1010010000.
package amount

import (
	"math/big"
)

// One is the value of one whole unit of currency.  Stellar uses 7 fixed digits
// for fractional values, thus One is 10 million (10^7)
const One = 10000000

func Parse(v string) (int64, error) {
	var f, o, r big.Float

	_, _, err := f.Parse(v, 10)
	if err != nil {
		return 0, err
	}

	o.SetInt64(One)
	r.Mul(&f, &o)

	i, _ := r.Int64()
	return i, nil
}

func String(v int64) string {
	var f, o, r big.Float

	f.SetInt64(v)
	o.SetInt64(One)
	r.Quo(&f, &o)

	return r.Text('f', 7)
}
