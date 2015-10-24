package keypair

import (
	"errors"
	"github.com/stellar/go-stellar-base/strkey"
)

var (
	// ErrInvalidKey will be returned by operations when the keypair being used
	// could not be decoded.
	ErrInvalidKey = errors.New("invalid key")

	// ErrInvalidSignature is returned when the signature is invalid, either
	// through malformation or if it does not verify the message against the
	// provided public key
	ErrInvalidSignature = errors.New("signature verification failed")

	// ErrCannotSign is returned when attempting to sign a message when
	// the keypair does not have the secret key available
	ErrCannotSign = errors.New("cannot sign")
)

// KP is the main interface for this package
type KP interface {
	Address() string
	Verify(input []byte, signature []byte) error
	Sign(input []byte) ([]byte, error)
}

func Parse(addressOrSeed string) (KP, error) {
	_, err := strkey.Decode(strkey.VersionByteAccountID, addressOrSeed)
	if err == nil {
		return &FromAddress{addressOrSeed}, nil
	}

	if err != strkey.ErrInvalidVersionByte {
		return nil, err
	}

	_, err = strkey.Decode(strkey.VersionByteSeed, addressOrSeed)
	if err == nil {
		return &Full{addressOrSeed}, nil
	}

	return nil, err
}
