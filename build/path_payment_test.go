package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/xdr"
)

var _ = Describe("PathPayment Mutators", func() {

	var (
		subject PathPaymentBuilder
		mut     interface{}

		address = "GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ"
		bad     = "foo"
	)

	JustBeforeEach(func() {
		subject = PathPaymentBuilder{}
		subject.Mutate(mut)
	})

	Describe("Destination", func() {
		Context("using a valid stellar address", func() {
			BeforeEach(func() { mut = Destination{address} })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the destination to the correct xdr.AccountId", func() {
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.P.Destination.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
		})

		Context("using an invalid value", func() {
			BeforeEach(func() { mut = Destination{bad} })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("PathDestination", func() {
		Context("native", func() {
			BeforeEach(func() {
				mut = PathDestination{
					Asset: Asset{Native: true},
					Amount: "50",
				}
			})
			It("sets the fields properly", func() {
				Expect(subject.P.DestAmount).To(Equal(xdr.Int64(500000000)))
				Expect(subject.P.DestAsset.Type).To(Equal(xdr.AssetTypeAssetTypeNative))
				Expect(subject.P.DestAsset.AlphaNum4).To(BeNil())
				Expect(subject.P.DestAsset.AlphaNum12).To(BeNil())
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("AlphaNum4", func() {
			BeforeEach(func() {
				mut = PathDestination{
					Asset: Asset{Code: "USD", Issuer: address},
					Amount: "50",
				}
			})
			It("sets the asset properly", func() {
				Expect(subject.P.DestAmount).To(Equal(xdr.Int64(500000000)))
				Expect(subject.P.DestAsset.Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum4))
				Expect(subject.P.DestAsset.AlphaNum4.AssetCode).To(Equal([4]byte{'U', 'S', 'D', 0}))
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.P.DestAsset.AlphaNum4.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
				Expect(subject.P.DestAsset.AlphaNum12).To(BeNil())
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("AlphaNum12", func() {
			BeforeEach(func() {
				mut = PathDestination{
					Asset: Asset{Code: "ABCDEF", Issuer: address},
					Amount: "50",
				}
			})
			It("sets the asset properly", func() {
				Expect(subject.P.DestAmount).To(Equal(xdr.Int64(500000000)))
				Expect(subject.P.DestAsset.Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum12))
				Expect(subject.P.DestAsset.AlphaNum4).To(BeNil())
				Expect(subject.P.DestAsset.AlphaNum12.AssetCode).To(Equal([12]byte{'A', 'B', 'C', 'D', 'E', 'F', 0, 0, 0, 0, 0, 0}))
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.P.DestAsset.AlphaNum12.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("issuer invalid", func() {
			BeforeEach(func() {
				mut = PathDestination{
					Asset: Asset{Code: "ABCDEF", Issuer: bad},
					Amount: "50",
				}
			})

			It("failed", func() {
				Expect(subject.Err).To(HaveOccurred())
			})
		})

		Context("amount invalid", func() {
			BeforeEach(func() {
				mut = PathDestination{
					Asset: Asset{Code: "ABCDEF", Issuer: address},
					Amount: "test",
				}
			})

			It("failed", func() {
				Expect(subject.Err).To(HaveOccurred())
			})
		})

		Context("asset code length invalid", func() {
			Context("empty", func() {
				BeforeEach(func() {
					mut = PathDestination{
						Asset: Asset{Code: "", Issuer: address},
						Amount: "50.0",
					}
				})

				It("failed", func() {
					Expect(subject.Err).To(MatchError("Asset code length is invalid"))
				})
			});

			Context("too long", func() {
				BeforeEach(func() {
					mut = PathDestination{
						Asset: Asset{Code: "1234567890123", Issuer: address},
						Amount: "50.0",
					}
				})

				It("failed", func() {
					Expect(subject.Err).To(MatchError("Asset code length is invalid"))
				})
			});
		})
	})

	Describe("PathSend", func() {
		Context("native", func() {
			BeforeEach(func() {
				mut = PathSend{
					Asset: Asset{Native: true},
					MaxAmount: "50",
				}
			})
			It("sets the fields properly", func() {
				Expect(subject.P.SendMax).To(Equal(xdr.Int64(500000000)))
				Expect(subject.P.SendAsset.Type).To(Equal(xdr.AssetTypeAssetTypeNative))
				Expect(subject.P.SendAsset.AlphaNum4).To(BeNil())
				Expect(subject.P.SendAsset.AlphaNum12).To(BeNil())
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("AlphaNum4", func() {
			BeforeEach(func() {
				mut = PathSend{
					Asset: Asset{Code: "USD", Issuer: address},
					MaxAmount: "50",
				}
			})
			It("sets the asset properly", func() {
				Expect(subject.P.SendMax).To(Equal(xdr.Int64(500000000)))
				Expect(subject.P.SendAsset.Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum4))
				Expect(subject.P.SendAsset.AlphaNum4.AssetCode).To(Equal([4]byte{'U', 'S', 'D', 0}))
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.P.SendAsset.AlphaNum4.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
				Expect(subject.P.SendAsset.AlphaNum12).To(BeNil())
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("AlphaNum12", func() {
			BeforeEach(func() {
				mut = PathSend{
					Asset: Asset{Code: "ABCDEF", Issuer: address},
					MaxAmount: "50",
				}
			})
			It("sets the asset properly", func() {
				Expect(subject.P.SendMax).To(Equal(xdr.Int64(500000000)))
				Expect(subject.P.SendAsset.Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum12))
				Expect(subject.P.SendAsset.AlphaNum4).To(BeNil())
				Expect(subject.P.SendAsset.AlphaNum12.AssetCode).To(Equal([12]byte{'A', 'B', 'C', 'D', 'E', 'F', 0, 0, 0, 0, 0, 0}))
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.P.SendAsset.AlphaNum12.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("issuer invalid", func() {
			BeforeEach(func() {
				mut = PathSend{
					Asset: Asset{Code: "ABCDEF", Issuer: bad},
					MaxAmount: "50",
				}
			})

			It("failed", func() {
				Expect(subject.Err).To(HaveOccurred())
			})
		})

		Context("amount invalid", func() {
			BeforeEach(func() {
				mut = PathSend{
					Asset: Asset{Code: "ABCDEF", Issuer: address},
					MaxAmount: "test",
				}
			})

			It("failed", func() {
				Expect(subject.Err).To(HaveOccurred())
			})
		})

		Context("asset code length invalid", func() {
			Context("empty", func() {
				BeforeEach(func() {
					mut = PathSend{
						Asset: Asset{Code: "", Issuer: address},
						MaxAmount: "50.0",
					}
				})

				It("failed", func() {
					Expect(subject.Err).To(MatchError("Asset code length is invalid"))
				})
			});

			Context("too long", func() {
				BeforeEach(func() {
					mut = PathSend{
						Asset: Asset{Code: "1234567890123", Issuer: address},
						MaxAmount: "50.0",
					}
				})

				It("failed", func() {
					Expect(subject.Err).To(MatchError("Asset code length is invalid"))
				})
			});
		})
	})

	Describe("Path", func() {
		Context("no intermediate assets", func() {
			BeforeEach(func() {
				mut = Path{}
			})
			It("sets the fields properly", func() {
				Expect(len(subject.P.Path)).To(Equal(0))
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("credit assets", func() {
			BeforeEach(func() {
				mut = Path{
					Assets: []Asset{
						Asset{Code: "USD", Issuer: address},
						Asset{Code: "12345", Issuer: address},
					},
				}
			})
			It("sets the fields properly", func() {
				Expect(len(subject.P.Path)).To(Equal(2))

				Expect(subject.P.Path[0].Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum4))
				Expect(subject.P.Path[0].AlphaNum4.AssetCode).To(Equal([4]byte{'U', 'S', 'D', 0}))
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.P.Path[0].AlphaNum4.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
				Expect(subject.P.Path[0].AlphaNum12).To(BeNil())

				Expect(subject.P.Path[1].Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum12))
				Expect(subject.P.Path[1].AlphaNum4).To(BeNil())
				Expect(subject.P.Path[1].AlphaNum12.AssetCode).To(Equal([12]byte{'1', '2', '3', '4', '5', 0}))
				aid, _ = stellarbase.AddressToAccountId(address)
				Expect(subject.P.Path[1].AlphaNum12.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("mixed assets", func() {
			BeforeEach(func() {
				mut = Path{
					Assets: []Asset{
						Asset{Native: true},
						Asset{Code: "12345", Issuer: address},
					},
				}
			})
			It("sets the fields properly", func() {
				Expect(len(subject.P.Path)).To(Equal(2))

				Expect(subject.P.Path[0].Type).To(Equal(xdr.AssetTypeAssetTypeNative))
				Expect(subject.P.Path[0].AlphaNum4).To(BeNil())
				Expect(subject.P.Path[0].AlphaNum12).To(BeNil())

				Expect(subject.P.Path[1].Type).To(Equal(xdr.AssetTypeAssetTypeCreditAlphanum12))
				Expect(subject.P.Path[1].AlphaNum4).To(BeNil())
				Expect(subject.P.Path[1].AlphaNum12.AssetCode).To(Equal([12]byte{'1', '2', '3', '4', '5', 0}))
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.P.Path[1].AlphaNum12.Issuer.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})
		})

		Context("issuer invalid", func() {
			BeforeEach(func() {
				mut = Path{
					Assets: []Asset{
						Asset{Native: true},
						Asset{Code: "12345", Issuer: bad},
					},
				}
			})

			It("failed", func() {
				Expect(subject.Err).To(HaveOccurred())
			})
		})

		Context("asset code length invalid", func() {
			Context("empty", func() {
				BeforeEach(func() {
					mut = Path{
						Assets: []Asset{
							Asset{Native: true},
							Asset{Code: "", Issuer: address},
						},
					}
				})

				It("failed", func() {
					Expect(subject.Err).To(MatchError("Asset code length is invalid"))
				})
			});

			Context("too long", func() {
				BeforeEach(func() {
					mut = Path{
						Assets: []Asset{
							Asset{Native: true},
							Asset{Code: "1234567890123", Issuer: address},
						},
					}
				})

				It("failed", func() {
					Expect(subject.Err).To(MatchError("Asset code length is invalid"))
				})
			});
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
