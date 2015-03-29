package stellarcore

import (
	"github.com/agl/ed25519"
	"strings"
)

type PublicKey [ed25519.PublicKeySize]byte
type PrivateKey [ed25519.PrivateKeySize]byte
type Signature [ed25519.SignatureSize]byte

func GenerateKeyFromRawSeed(str string) (publicKey PublicKey, privateKey PrivateKey, err error) {

	reader := strings.NewReader(str)
	pub, priv, err := ed25519.GenerateKey(reader)

	if err != nil {
		return
	}

	publicKey = *pub
	privateKey = *priv
	return
}

func Sign(privateKey PrivateKey, message []byte) Signature {
	priv := [ed25519.PrivateKeySize]byte(privateKey)
	return *ed25519.Sign(&priv, message)
}

func Verify(publicKey PublicKey, message []byte, signature Signature) bool {
	pub := [ed25519.PublicKeySize]byte(publicKey)
	sig := [ed25519.SignatureSize]byte(signature)

	return ed25519.Verify(&pub, message, &sig)
}
