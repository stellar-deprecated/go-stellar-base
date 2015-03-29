package stellarcore

import (
	"crypto/sha256"
)

func Hash(message []byte) [32]byte {
	return sha256.Sum256(message)
}
