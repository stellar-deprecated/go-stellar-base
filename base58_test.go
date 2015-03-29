package stellarcore

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBase58Encoding(t *testing.T) {
	Convey("Given a message", t, func() {
		unencoded := []byte("hello world")
		encoded := "StVgDLaUATiyKyV"

		Convey("When encoding the message", func() {
			actual := EncodeBase58(unencoded)

			Convey("The encoded should be correct", func() {
				So(actual, ShouldEqual, encoded)
			})
		})

		Convey("When decoding the message", func() {
			actual, err := DecodeBase58(encoded)
			So(err, ShouldBeNil)

			Convey("The dedoding should be correct", func() {
				So(actual, ShouldResemble, unencoded)
			})
		})

	})
}
