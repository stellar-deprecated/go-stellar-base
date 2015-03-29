package stellarcore

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/go-stellar-core/xdr"
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

func TestPayment(t *testing.T) {
	SkipConvey("A payment can be constructed", t, func() {
		seed := "s3Fy8h5LEcYVE8aofthKWHeJpygbntw5HgcekFw93K6XqTW4gEx"
		pub, priv, err := GenerateKeyFromSeed(seed)
		So(err, ShouldBeNil)
		_, _ = pub, priv

		currency := xdr.NewCurrencyNative()
		_ = currency
	})

}
