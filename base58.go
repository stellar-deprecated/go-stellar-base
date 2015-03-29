package stellarbase

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

const alphabet = "gsphnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCr65jkm8oFqi1tuvAxyz"

var decodeMap [256]byte

func init() {
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	for i := 0; i < len(alphabet); i++ {
		decodeMap[alphabet[i]] = byte(i)
	}
}

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base58 data at input byte " + strconv.FormatInt(int64(e), 10)
}

type InvalidVersionByteError struct {
	Expected VersionByte
	Actual   VersionByte
}

func (e InvalidVersionByteError) Error() string {
	return fmt.Sprintf("illegal base58 version byte expected:%d actual:%d", e.Expected, e.Actual)
}

type VersionByte byte

const (
	VersionByteAccountID VersionByte = 0
	VersionByteNone                  = 1
	VersionByteSeed                  = 33
)

func EncodeBase58(src []byte) string {
	bigInt := new(big.Int)
	bigInt.SetBytes(src)
	leadingZeroes := strings.Repeat("g", leadingZeroCount(src))

	var resultSlice []byte = make([]byte, 0, 256)
	var results = []string{
		leadingZeroes,
		string(EncodeBigToBase58(resultSlice, bigInt)),
	}

	return strings.Join(results, "")
}

func EncodeBase58Check(version VersionByte, src []byte) string {
	start := []byte{}
	withVersion := append(start, byte(version))
	withPayload := append(withVersion, src...)
	withChecksum := append(withPayload, base58CheckSum(withPayload)...)
	// _ = withChecksum
	// return fmt.Sprintf("b: %s\n", withChecksum)
	return string(EncodeBase58(withChecksum))
}

func DecodeBase58(src string) ([]byte, error) {
	leadingGs := make([]byte, leadingGCount(src))

	bigInt, err := DecodeBase58ToBig([]byte(src))

	if err != nil {
		return []byte{}, err
	}

	return append(leadingGs, bigInt.Bytes()...), nil
}

func DecodeBase58Check(version VersionByte, src string) ([]byte, error) {
	decoded, err := DecodeBase58(src)

	if err != nil {
		return []byte{}, err
	}

	decodedVersion := VersionByte(decoded[0])
	payload := decoded[1 : len(decoded)-4]
	checksum := decoded[len(decoded)-4:]

	if decodedVersion != version {
		return []byte{}, InvalidVersionByteError{version, decodedVersion}
	}

	_ = checksum

	return payload, nil
}

// Decode a big integer from the bytes. Returns an error on corrupt
// input.
func DecodeBase58ToBig(src []byte) (*big.Int, error) {
	n := new(big.Int)
	radix := big.NewInt(58)
	for i := 0; i < len(src); i++ {
		b := decodeMap[src[i]]
		if b == 0xFF {
			return nil, CorruptInputError(i)
		}
		n.Mul(n, radix)
		n.Add(n, big.NewInt(int64(b)))
	}
	return n, nil
}

// Encode encodes src, appending to dst. Be sure to use the returned
// new value of dst.
func EncodeBigToBase58(dst []byte, src *big.Int) []byte {
	start := len(dst)
	n := new(big.Int)
	n.Set(src)
	radix := big.NewInt(58)
	zero := big.NewInt(0)

	for n.Cmp(zero) > 0 {
		mod := new(big.Int)
		n.DivMod(n, radix, mod)
		dst = append(dst, alphabet[mod.Int64()])
	}

	// reverse string
	for i, j := start, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return dst
}

func base58CheckSum(message []byte) []byte {
	inner := Hash(message)
	outer := Hash(inner[:])
	return outer[0:4]
}

func leadingZeroCount(src []byte) (result int) {
	for _, val := range src {
		if val != 0x00 {
			return
		}
		result++
	}
	return
}

func leadingGCount(src string) (result int) {
	for _, val := range src {
		if val != 'g' {
			return
		}
		result++
	}
	return
}
