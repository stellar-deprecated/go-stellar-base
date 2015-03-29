package stellarbase

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/go-stellar-base/xdr"
	"testing"
)

func TestPrice(t *testing.T) {

	Convey("A price can be constructed", t, func() {
		price := xdr.Price{1, 10}

		So(price.N, ShouldEqual, 1)
		So(price.D, ShouldEqual, 10)
	})
}

func TestCurrency(t *testing.T) {

	Convey("A native currency can be constructed", t, func() {
		currency := xdr.NewCurrencyNative()
		So(currency.Type(), ShouldEqual, xdr.CurrencyTypeNative)
	})

	Convey("Given an IsoCurrencyIssuer", t, func() {
		currencyCode := [4]byte{'u', 's', 'd', 0x00}
		accountID := [32]byte{}
		isoci := xdr.IsoCurrencyIssuer{currencyCode, accountID}

		Convey("An iso currency can be constructed", func() {
			currency := xdr.NewCurrencyIso4217(isoci)
			So(currency.Type(), ShouldEqual, xdr.CurrencyTypeIso4217)
			So(currency.IsoCi(), ShouldResemble, isoci)
		})
	})

}

func TestLowLevelPayment(t *testing.T) {
	SkipConvey("A payment can be constructed", t, func() {
		seed := "s3Fy8h5LEcYVE8aofthKWHeJpygbntw5HgcekFw93K6XqTW4gEx"
		pub, priv, err := GenerateKeyFromSeed(seed)
		So(err, ShouldBeNil)
		_, _ = pub, priv

		currency := xdr.NewCurrencyNative()
		accountId := xdr.AccountId(pub.KeyData())

		pOp := xdr.PaymentOp{
			Destination: accountId,
			Currency:    currency,
			Amount:      100000000,
			SendMax:     100000000,
		}

		opBody := xdr.NewOperationBodyPayment(pOp)

		p := xdr.Operation{Body: opBody}

		tx := xdr.Transaction{
			Account:    accountId,
			MaxFee:     10,
			SeqNum:     1,
			MinLedger:  1,
			MaxLedger:  1000,
			Operations: []xdr.Operation{p},
		}

		_ = tx
	})

}
