package build

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuild(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package: github.com/stellar/go-stellar-base/build")
}

// ExampleTransactionBuilder creates and signs a simple transaction, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
//
// It uses the transaction builder system
func ExampleTransactionBuilder() {
	seed := "SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		Payment(
			Destination{"GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA"},
			NativeAmount{"50"},
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAALSRpLtCLv2eboZlEiHDSGR6Hb+zZL92fbSdNpObeE0EAAAAAAAAAAB3NZQAAAAAAAAAAARtDMfAAAABA2oIeQxoJl53RMRWFeLB865zcky39f2gf2PmUubCuJYccEePRSrTC8QQrMOgGwD8a6oe8dgltvezdDsmmXBPyBw==
}

// ExamplePathPayment creates and signs a simple transaction with PathPayment operation, and then
// encodes it into a base64 string capable of being submitted to stellar-core.
func ExamplePathPayment() {
	seed := "SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"
	tx := Transaction(
		SourceAccount{seed},
		Sequence{1},
		PathPayment(
			Destination{"GBDT3K42LOPSHNAEHEJ6AVPADIJ4MAR64QEKKW2LQPBSKLYD22KUEH4P"},
			PathSend{
				Asset: Asset{
					Code: "USD",
					Issuer: "GAWSI2JO2CF36Z43UGMUJCDQ2IMR5B3P5TMS7XM7NUTU3JHG3YJUDQXA",
				},
				MaxAmount: "50",
			},
			PathDestination{
				Asset: Asset{
					Code: "EUR",
					Issuer: "GCPZJ3MJQ3GUGJSBL6R3MLYZS6FKVHG67BPAINMXL3NWNXR5S6XG657P",
				},
				Amount: "100",
			},
			Path{
				Assets: []Asset{
					Asset{Native: true},
					Asset{
						Code: "BTC",
						Issuer: "GAHJZHVKFLATAATJH46C7OK2ZOVRD47GZBGQ7P6OCVF6RJDCEG5JMQBQ",
					},
				},
			},
		),
	)

	txe := tx.Sign(seed)
	txeB64, _ := txe.Base64()

	fmt.Printf("tx base64: %s", txeB64)
	// Output: tx base64: AAAAADZY/nWY0gx6beMpf4S8Ur0qHsjA8fbFtBzBx1cbQzHwAAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAIAAAABVVNEAAAAAAAtJGku0Iu/Z5uhmUSIcNIZHodv7Nkv3Z9tJ02k5t4TQQAAAAAdzWUAAAAAAEc9q5pbnyO0BDkT4FXgGhPGAj7kCKVbS4PDJS8D1pVCAAAAAUVVUgAAAAAAn5TtiYbNQyZBX6O2LxmXiqqc3vheBDWXXttm3j2Xrm8AAAAAO5rKAAAAAAIAAAAAAAAAAUJUQwAAAAAADpyeqirBMAJpPzwvuVrLqxHz5shND7/OFUvopGIhupYAAAAAAAAAARtDMfAAAABAvbv3jz4I/xVsIxamvzyOaa048aJZHmpLpjny5b0pe3ln1geH593O5jOoo4CPBi2AsgqlnAO55cz8yMQ0ibuCAQ==
}
