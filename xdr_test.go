package stellarcore

import (
	"./xdr"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPrice(t *testing.T) {

	Convey("A price can be constructed", t, func() {
		price := xdr.Price{1, 10}

		So(price.N, ShouldEqual, 1)
		So(price.D, ShouldEqual, 10)
	})
}
