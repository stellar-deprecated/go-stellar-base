package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/xdr"
)

var _ = Describe("Transaction Mutators:", func() {

	var (
		subject *TransactionBuilder
		mut     TransactionMutator
	)

	BeforeEach(func() { subject = &TransactionBuilder{} })
	JustBeforeEach(func() { subject.Mutate(mut) })

	Describe("Defaults", func() {
		BeforeEach(func() { mut = Defaults{} })
		It("sets the fee", func() { Expect(subject.TX.Fee).To(BeEquivalentTo(100)) })
		It("sets the memo", func() { Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoNone)) })
	})

	Describe("PaymentBuilder", func() {
		BeforeEach(func() { mut = Payment() })
		It("adds itself to the tx's operations", func() {
			Expect(subject.TX.Operations).To(HaveLen(1))
		})
	})

	Describe("SourceAccount", func() {
		Context("with a valid address", func() {
			address := "GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ"
			BeforeEach(func() { mut = SourceAccount{address} })
			It("sets the AccountId correctly", func() {
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.TX.SourceAccount.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
		})

		Context("with bad address", func() {
			BeforeEach(func() { mut = SourceAccount{"foo"} })
			It("fails", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("Sequence", func() {
		BeforeEach(func() { mut = Sequence{12345} })
		It("succeeds", func() { Expect(subject.Err).NotTo(HaveOccurred()) })
		It("sets the sequence", func() { Expect(subject.TX.SeqNum).To(BeEquivalentTo(12345)) })
	})

})
