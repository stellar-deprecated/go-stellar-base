package xdr_test

import (
	. "github.com/stellar/go-stellar-base/xdr"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("xdr.AccountId#Address()", func() {
	It("returns an empty string when account id is nil", func() {
		addy := (*AccountId)(nil).Address()
		Expect(addy).To(Equal(""))
	})

	It("returns a strkey string when account id is valid", func() {
		var aid AccountId
		aid.SetAddress("GCR22L3WS7TP72S4Z27YTO6JIQYDJK2KLS2TQNHK6Y7XYPA3AGT3X4FH")
		addy := aid.Address()
		Expect(addy).To(Equal("GCR22L3WS7TP72S4Z27YTO6JIQYDJK2KLS2TQNHK6Y7XYPA3AGT3X4FH"))
	})
})
