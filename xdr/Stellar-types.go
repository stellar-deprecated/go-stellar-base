// Automatically generated from xdr/Stellar-types.x
// DO NOT EDIT or your changes may be overwritten

package xdr

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-xdr/xdr2"
)

// === xdr source ============================================================
//
//   typedef opaque uint512[64];
//
// ===========================================================================

type Uint512 [64]byte

func DecodeUint512(decoder *xdr.Decoder, result *Uint512) (int, error) {
	var (
		val       [64]byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], 64)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Uint512(val)
	return totalRead, nil
}
func DecodeOptionalUint512(decoder *xdr.Decoder, result **Uint512) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Uint512

	bytesRead, err = DecodeUint512(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeUint512FixedArray(decoder *xdr.Decoder, result []Uint512, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeUint512(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUint512Array(decoder *xdr.Decoder, result *[]Uint512, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Uint512, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeUint512(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef opaque uint256[32];
//
// ===========================================================================

type Uint256 [32]byte

func DecodeUint256(decoder *xdr.Decoder, result *Uint256) (int, error) {
	var (
		val       [32]byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], 32)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Uint256(val)
	return totalRead, nil
}
func DecodeOptionalUint256(decoder *xdr.Decoder, result **Uint256) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Uint256

	bytesRead, err = DecodeUint256(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeUint256FixedArray(decoder *xdr.Decoder, result []Uint256, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeUint256(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUint256Array(decoder *xdr.Decoder, result *[]Uint256, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Uint256, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeUint256(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef unsigned hyper uint64;
//
// ===========================================================================

type Uint64 uint64

func DecodeUint64(decoder *xdr.Decoder, result *Uint64) (int, error) {
	var (
		val       uint64
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeUhyper(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Uint64(val)
	return totalRead, nil
}
func DecodeOptionalUint64(decoder *xdr.Decoder, result **Uint64) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Uint64

	bytesRead, err = DecodeUint64(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeUint64FixedArray(decoder *xdr.Decoder, result []Uint64, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeUint64(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUint64Array(decoder *xdr.Decoder, result *[]Uint64, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Uint64, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeUint64(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef hyper int64;
//
// ===========================================================================

type Int64 int64

func DecodeInt64(decoder *xdr.Decoder, result *Int64) (int, error) {
	var (
		val       int64
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeHyper(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Int64(val)
	return totalRead, nil
}
func DecodeOptionalInt64(decoder *xdr.Decoder, result **Int64) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Int64

	bytesRead, err = DecodeInt64(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeInt64FixedArray(decoder *xdr.Decoder, result []Int64, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeInt64(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeInt64Array(decoder *xdr.Decoder, result *[]Int64, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Int64, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeInt64(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef unsigned int uint32;
//
// ===========================================================================

type Uint32 uint32

func DecodeUint32(decoder *xdr.Decoder, result *Uint32) (int, error) {
	var (
		val       uint32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeUint(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Uint32(val)
	return totalRead, nil
}
func DecodeOptionalUint32(decoder *xdr.Decoder, result **Uint32) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Uint32

	bytesRead, err = DecodeUint32(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeUint32FixedArray(decoder *xdr.Decoder, result []Uint32, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeUint32(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUint32Array(decoder *xdr.Decoder, result *[]Uint32, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Uint32, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeUint32(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef int int32;
//
// ===========================================================================

type Int32 int32

func DecodeInt32(decoder *xdr.Decoder, result *Int32) (int, error) {
	var (
		val       int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Int32(val)
	return totalRead, nil
}
func DecodeOptionalInt32(decoder *xdr.Decoder, result **Int32) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Int32

	bytesRead, err = DecodeInt32(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeInt32FixedArray(decoder *xdr.Decoder, result []Int32, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeInt32(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeInt32Array(decoder *xdr.Decoder, result *[]Int32, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Int32, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeInt32(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef opaque AccountID[32];
//
// ===========================================================================

type AccountId [32]byte

func DecodeAccountId(decoder *xdr.Decoder, result *AccountId) (int, error) {
	var (
		val       [32]byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], 32)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = AccountId(val)
	return totalRead, nil
}
func DecodeOptionalAccountId(decoder *xdr.Decoder, result **AccountId) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val AccountId

	bytesRead, err = DecodeAccountId(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAccountIdFixedArray(decoder *xdr.Decoder, result []AccountId, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeAccountId(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAccountIdArray(decoder *xdr.Decoder, result *[]AccountId, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]AccountId, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAccountId(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef opaque Signature[64];
//
// ===========================================================================

type Signature [64]byte

func DecodeSignature(decoder *xdr.Decoder, result *Signature) (int, error) {
	var (
		val       [64]byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], 64)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Signature(val)
	return totalRead, nil
}
func DecodeOptionalSignature(decoder *xdr.Decoder, result **Signature) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Signature

	bytesRead, err = DecodeSignature(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSignatureFixedArray(decoder *xdr.Decoder, result []Signature, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeSignature(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSignatureArray(decoder *xdr.Decoder, result *[]Signature, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Signature, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSignature(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef opaque Hash[32];
//
// ===========================================================================

type Hash [32]byte

func DecodeHash(decoder *xdr.Decoder, result *Hash) (int, error) {
	var (
		val       [32]byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], 32)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Hash(val)
	return totalRead, nil
}
func DecodeOptionalHash(decoder *xdr.Decoder, result **Hash) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Hash

	bytesRead, err = DecodeHash(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeHashFixedArray(decoder *xdr.Decoder, result []Hash, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeHash(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeHashArray(decoder *xdr.Decoder, result *[]Hash, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Hash, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeHash(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef opaque Thresholds[4];
//
// ===========================================================================

type Thresholds [4]byte

func DecodeThresholds(decoder *xdr.Decoder, result *Thresholds) (int, error) {
	var (
		val       [4]byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], 4)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Thresholds(val)
	return totalRead, nil
}
func DecodeOptionalThresholds(decoder *xdr.Decoder, result **Thresholds) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Thresholds

	bytesRead, err = DecodeThresholds(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeThresholdsFixedArray(decoder *xdr.Decoder, result []Thresholds, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeThresholds(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeThresholdsArray(decoder *xdr.Decoder, result *[]Thresholds, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Thresholds, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeThresholds(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   typedef uint64 SequenceNumber;
//
// ===========================================================================

type SequenceNumber Uint64

func DecodeSequenceNumber(decoder *xdr.Decoder, result *SequenceNumber) (int, error) {
	var (
		val       Uint64
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeUint64(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = SequenceNumber(val)
	return totalRead, nil
}
func DecodeOptionalSequenceNumber(decoder *xdr.Decoder, result **SequenceNumber) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val SequenceNumber

	bytesRead, err = DecodeSequenceNumber(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSequenceNumberFixedArray(decoder *xdr.Decoder, result []SequenceNumber, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeSequenceNumber(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSequenceNumberArray(decoder *xdr.Decoder, result *[]SequenceNumber, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]SequenceNumber, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSequenceNumber(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   enum CurrencyType
//   {
//       NATIVE,
//       ISO4217
//   };
//
// ===========================================================================

type CurrencyType int32

const (
	CurrencyTypeNative  CurrencyType = 0
	CurrencyTypeIso4217              = 1
)

var CurrencyTypeMap = map[int32]bool{
	0: true,
	1: true,
}

func DecodeCurrencyType(decoder *xdr.Decoder, result *CurrencyType) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(CurrencyTypeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = CurrencyType(val)
	return bytesRead, err
}
func DecodeOptionalCurrencyType(decoder *xdr.Decoder, result **CurrencyType) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val CurrencyType

	bytesRead, err = DecodeCurrencyType(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCurrencyTypeFixedArray(decoder *xdr.Decoder, result []CurrencyType, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeCurrencyType(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCurrencyTypeArray(decoder *xdr.Decoder, result *[]CurrencyType, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]CurrencyType, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCurrencyType(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   struct ISOCurrencyIssuer
//   {
//       opaque currencyCode[4];
//       AccountID issuer;
//   };
//
// ===========================================================================

type IsoCurrencyIssuer struct {
	CurrencyCode [4]byte
	Issuer       AccountId
}

func DecodeIsoCurrencyIssuer(decoder *xdr.Decoder, result *IsoCurrencyIssuer) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeFixedOpaque(decoder, result.CurrencyCode[:], 4)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeAccountId(decoder, &result.Issuer)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalIsoCurrencyIssuer(decoder *xdr.Decoder, result **IsoCurrencyIssuer) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val IsoCurrencyIssuer

	bytesRead, err = DecodeIsoCurrencyIssuer(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeIsoCurrencyIssuerFixedArray(decoder *xdr.Decoder, result []IsoCurrencyIssuer, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodeIsoCurrencyIssuer(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeIsoCurrencyIssuerArray(decoder *xdr.Decoder, result *[]IsoCurrencyIssuer, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]IsoCurrencyIssuer, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeIsoCurrencyIssuer(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

// === xdr source ============================================================
//
//   union Currency switch(CurrencyType type)
//   {
//       case NATIVE:
//           void;
//
//       case ISO4217:
//           ISOCurrencyIssuer isoCI;
//
//       // add other currency types here in the future
//   };
//
// ===========================================================================

type Currency struct {
	aType CurrencyType
	isoCi *IsoCurrencyIssuer
}

func NewCurrencyNative() Currency {
	return Currency{
		aType: CurrencyTypeNative,
	}
}
func NewCurrencyIso4217(val IsoCurrencyIssuer) Currency {
	return Currency{
		aType: CurrencyTypeIso4217,
		isoCi: &val,
	}
}
func (u *Currency) Type() CurrencyType {
	return u.aType
}
func (u *Currency) IsoCi() IsoCurrencyIssuer {
	return *u.isoCi
}
func DecodeCurrency(decoder *xdr.Decoder, result *Currency) (int, error) {
	var (
		discriminant CurrencyType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeCurrencyType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Currency{}
	return totalRead, nil
}

// === xdr source ============================================================
//
//   struct Price
//   {
//       int32 n;
//       int32 d;
//   };
//
// ===========================================================================

type Price struct {
	N Int32
	D Int32
}

func DecodePrice(decoder *xdr.Decoder, result *Price) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeInt32(decoder, &result.N)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt32(decoder, &result.D)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalPrice(decoder *xdr.Decoder, result **Price) (int, error) {
	var (
		isPresent bool
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeBool(decoder, &isPresent)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if !isPresent {
		return totalRead, err
	}

	var val Price

	bytesRead, err = DecodePrice(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodePriceFixedArray(decoder *xdr.Decoder, result []Price, size int) (int, error) {
	var (
		totalRead int
		bytesRead int
		err       error
	)

	if len(result) != size {
		errMsg := fmt.Sprintf("xdr: bad array len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	for i := 0; i < size; i++ {
		bytesRead, err = DecodePrice(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodePriceArray(decoder *xdr.Decoder, result *[]Price, maxSize int32) (int, error) {
	var (
		size      int32
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeInt(decoder, &size)
	totalRead += bytesRead

	if err != nil {
		return totalRead, err
	}

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: encoded array size too large:%d, max:%d", size, maxSize)
		return totalRead, errors.New(errMsg)
	}

	var theResult = make([]Price, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodePrice(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}
