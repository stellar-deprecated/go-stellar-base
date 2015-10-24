package keypair

import (
	"bytes"
	"github.com/agl/ed25519"
	"github.com/stellar/go-stellar-base/strkey"
)

type Full struct {
	seed string
}

func (kp *Full) Address() string {
	return strkey.MustEncode(strkey.VersionByteAccountID, kp.publicKey()[:])
}

func (kp *Full) Seed() string {
	return kp.seed
}

func (kp *Full) Verify(input []byte, sig []byte) error {
	if len(sig) != 64 {
		return ErrInvalidSignature
	}

	var asig [64]byte
	copy(asig[:], sig[:])

	if !ed25519.Verify(kp.publicKey(), input, &asig) {
		return ErrInvalidSignature
	}
	return nil
}

func (kp *Full) Sign(input []byte) ([]byte, error) {
	_, priv := kp.keys()
	return ed25519.Sign(priv, input)[:], nil
}

func (kp *Full) publicKey() *[32]byte {
	pub, _ := kp.keys()
	return pub
}

func (kp *Full) keys() (*[32]byte, *[64]byte) {
	reader := bytes.NewReader(kp.rawSeed())
	pub, priv, err := ed25519.GenerateKey(reader)
	if err != nil {
		panic(err)
	}
	return pub, priv
}

func (kp *Full) rawSeed() []byte {
	return strkey.MustDecode(strkey.VersionByteSeed, kp.seed)
}
