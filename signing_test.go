package stellarcore

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSigning(t *testing.T) {
	Convey("Given the master passphrase", t, func() {
		rawSeed := "masterpassphrasemasterpassphrase"

		Convey("and a key pair from the seed", func() {
			pub, priv, err := GenerateKeyFromRawSeed(rawSeed)

			So(err, ShouldBeNil)
			So(len(pub), ShouldEqual, 32)
			So(len(priv), ShouldEqual, 64)

			Convey("When signing a message", func() {
				message := []byte("hello world")
				signature := Sign(priv, message)

				Convey("The signature should be verifiable", func() {
					valid := Verify(pub, message, signature)
					So(valid, ShouldBeTrue)
				})
			})
		})
	})
}
