package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/xdr"
)

var _ = Describe("ManageOffer", func() {

	Describe("ManageOfferBuilder", func() {
		var (
			subject ManageOfferBuilder
			mut     interface{}

			address = "GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ"
			bad     = "foo"

			rate = Rate{
				Selling: CreditAsset("EUR", "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"),
				Buying:  NativeAsset(),
				Price:   Price("41.265"),
			}
		)

		JustBeforeEach(func() {
			subject = ManageOfferBuilder{}
			subject.Mutate(mut)
		})

		Describe("CreateOffer", func() {
			Context("creates offer properly", func() {
				It("sets values properly", func() {
					builder := CreateOffer(rate, "20")

					Expect(builder.MO.Amount).To(Equal(xdr.Int64(200000000)))

					Expect(builder.MO.Selling.Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum4))
					Expect(builder.MO.Selling.AlphaNum4.AssetCode).To(Equal([4]byte{'E', 'U', 'R', 0}))
					aid, _ := stellarbase.AddressToAccountId(rate.Selling.Issuer)
					Expect(builder.MO.Selling.AlphaNum4.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
					Expect(builder.MO.Selling.AlphaNum12).To(BeNil())

					Expect(builder.MO.Buying.Type).To(Equal(xdr.AssetTypeAssetTypeNative))
					Expect(builder.MO.Buying.AlphaNum4).To(BeNil())
					Expect(builder.MO.Buying.AlphaNum12).To(BeNil())

					Expect(builder.MO.Price.N).To(Equal(xdr.Int32(8253)))
					Expect(builder.MO.Price.D).To(Equal(xdr.Int32(200)))

					Expect(builder.MO.OfferId).To(Equal(xdr.Uint64(0)))
				})
			})
		})

		Describe("UpdateOffer", func() {
			Context("updates the offer properly", func() {
				It("sets values properly", func() {
					builder := UpdateOffer(rate, "100", 5)

					Expect(builder.MO.Amount).To(Equal(xdr.Int64(1000000000)))

					Expect(builder.MO.Selling.Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum4))
					Expect(builder.MO.Selling.AlphaNum4.AssetCode).To(Equal([4]byte{'E', 'U', 'R', 0}))
					aid, _ := stellarbase.AddressToAccountId(rate.Selling.Issuer)
					Expect(builder.MO.Selling.AlphaNum4.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
					Expect(builder.MO.Selling.AlphaNum12).To(BeNil())

					Expect(builder.MO.Buying.Type).To(Equal(xdr.AssetTypeAssetTypeNative))
					Expect(builder.MO.Buying.AlphaNum4).To(BeNil())
					Expect(builder.MO.Buying.AlphaNum12).To(BeNil())

					Expect(builder.MO.Price.N).To(Equal(xdr.Int32(8253)))
					Expect(builder.MO.Price.D).To(Equal(xdr.Int32(200)))

					Expect(builder.MO.OfferId).To(Equal(xdr.Uint64(5)))
				})
			})
		})

		Describe("DeleteOffer", func() {
			Context("deletes the offer properly", func() {
				It("sets values properly", func() {
					builder := DeleteOffer(rate, 10)

					Expect(builder.MO.Amount).To(Equal(xdr.Int64(0)))

					Expect(builder.MO.Selling.Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum4))
					Expect(builder.MO.Selling.AlphaNum4.AssetCode).To(Equal([4]byte{'E', 'U', 'R', 0}))
					aid, _ := stellarbase.AddressToAccountId(rate.Selling.Issuer)
					Expect(builder.MO.Selling.AlphaNum4.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
					Expect(builder.MO.Selling.AlphaNum12).To(BeNil())

					Expect(builder.MO.Buying.Type).To(Equal(xdr.AssetTypeAssetTypeNative))
					Expect(builder.MO.Buying.AlphaNum4).To(BeNil())
					Expect(builder.MO.Buying.AlphaNum12).To(BeNil())

					Expect(builder.MO.Price.N).To(Equal(xdr.Int32(8253)))
					Expect(builder.MO.Price.D).To(Equal(xdr.Int32(200)))

					Expect(builder.MO.OfferId).To(Equal(xdr.Uint64(10)))
				})
			})
		})

		Describe("SourceAccount", func() {
			Context("using a valid stellar address", func() {
				BeforeEach(func() { mut = SourceAccount{address} })

				It("succeeds", func() {
					Expect(subject.Err).NotTo(HaveOccurred())
				})

				It("sets the destination to the correct xdr.AccountId", func() {
					aid, _ := stellarbase.AddressToAccountId(address)
					Expect(subject.O.SourceAccount.MustEd25519()).To(Equal(aid.MustEd25519()))
				})
			})

			Context("using an invalid value", func() {
				BeforeEach(func() { mut = SourceAccount{bad} })
				It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
			})
		})
	})

	Describe("continuedFraction", func() {
		It("succeeds", func() {
			Expect(continuedFraction("0.1")).To(Equal(xdr.Price{1, 10}))
			Expect(continuedFraction("0.01")).To(Equal(xdr.Price{1, 100}))
			Expect(continuedFraction("0.001")).To(Equal(xdr.Price{1, 1000}))
			Expect(continuedFraction("543.017930")).To(Equal(xdr.Price{54301793, 100000}))
			Expect(continuedFraction("319.69983")).To(Equal(xdr.Price{31969983, 100000}))
			Expect(continuedFraction("0.93")).To(Equal(xdr.Price{93, 100}))
			Expect(continuedFraction("0.5")).To(Equal(xdr.Price{1, 2}))
			Expect(continuedFraction("1.730")).To(Equal(xdr.Price{173, 100}))
			Expect(continuedFraction("0.85334384")).To(Equal(xdr.Price{5333399, 6250000}))
			Expect(continuedFraction("5.5")).To(Equal(xdr.Price{11, 2}))
			Expect(continuedFraction("2.72783")).To(Equal(xdr.Price{272783, 100000}))
			Expect(continuedFraction("638082.0")).To(Equal(xdr.Price{638082, 1}))
			Expect(continuedFraction("2.93850088")).To(Equal(xdr.Price{36731261, 12500000}))
			Expect(continuedFraction("58.04")).To(Equal(xdr.Price{1451, 25}))
			Expect(continuedFraction("41.265")).To(Equal(xdr.Price{8253, 200}))
			Expect(continuedFraction("5.1476")).To(Equal(xdr.Price{12869, 2500}))
			Expect(continuedFraction("95.14")).To(Equal(xdr.Price{4757, 50}))
			Expect(continuedFraction("0.74580")).To(Equal(xdr.Price{3729, 5000}))
			Expect(continuedFraction("4119.0")).To(Equal(xdr.Price{4119, 1}))
		})

		It("fails for invalid values", func() {
			var err error
			_, err = continuedFraction("0.0000000003")
			Expect(err).Should(HaveOccurred())
			_, err = continuedFraction("2147483648")
			Expect(err).Should(HaveOccurred())
		})
	})
})
