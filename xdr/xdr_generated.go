// Automatically generated from xdr/Stellar-types.x,xdr/Stellar-ledger-entries.x,xdr/Stellar-ledger.x,xdr/Stellar-transaction.x,xdr/Stellar-overlay.x,xdr/SCPXDR.x
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

func EncodeUint512(encoder *xdr.Encoder, value *Uint512) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[64]byte)(value)[:], 64)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalUint512(encoder *xdr.Encoder, value *Uint512) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeUint512(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeUint512FixedArray(encoder *xdr.Encoder, value []Uint512, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint512(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeUint512Array(encoder *xdr.Encoder, value []Uint512, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint512(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeUint256(encoder *xdr.Encoder, value *Uint256) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[32]byte)(value)[:], 32)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalUint256(encoder *xdr.Encoder, value *Uint256) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeUint256(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeUint256FixedArray(encoder *xdr.Encoder, value []Uint256, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint256(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeUint256Array(encoder *xdr.Encoder, value []Uint256, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint256(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeUint64(encoder *xdr.Encoder, value *Uint64) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeUhyper(encoder, (*uint64)(value))
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalUint64(encoder *xdr.Encoder, value *Uint64) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeUint64(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeUint64FixedArray(encoder *xdr.Encoder, value []Uint64, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint64(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeUint64Array(encoder *xdr.Encoder, value []Uint64, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint64(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeInt64(encoder *xdr.Encoder, value *Int64) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeHyper(encoder, (*int64)(value))
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalInt64(encoder *xdr.Encoder, value *Int64) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeInt64(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeInt64FixedArray(encoder *xdr.Encoder, value []Int64, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeInt64(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeInt64Array(encoder *xdr.Encoder, value []Int64, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeInt64(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeUint32(encoder *xdr.Encoder, value *Uint32) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeUint(encoder, (*uint32)(value))
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalUint32(encoder *xdr.Encoder, value *Uint32) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeUint32(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeUint32FixedArray(encoder *xdr.Encoder, value []Uint32, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint32(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeUint32Array(encoder *xdr.Encoder, value []Uint32, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeUint32(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeInt32(encoder *xdr.Encoder, value *Int32) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeInt(encoder, (*int32)(value))
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalInt32(encoder *xdr.Encoder, value *Int32) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeInt32(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeInt32FixedArray(encoder *xdr.Encoder, value []Int32, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeInt32(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeInt32Array(encoder *xdr.Encoder, value []Int32, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeInt32(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeAccountId(encoder *xdr.Encoder, value *AccountId) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[32]byte)(value)[:], 32)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalAccountId(encoder *xdr.Encoder, value *AccountId) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAccountId(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAccountIdFixedArray(encoder *xdr.Encoder, value []AccountId, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountId(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAccountIdArray(encoder *xdr.Encoder, value []AccountId, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountId(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeSignature(encoder *xdr.Encoder, value *Signature) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[64]byte)(value)[:], 64)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalSignature(encoder *xdr.Encoder, value *Signature) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSignature(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSignatureFixedArray(encoder *xdr.Encoder, value []Signature, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSignature(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSignatureArray(encoder *xdr.Encoder, value []Signature, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSignature(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeHash(encoder *xdr.Encoder, value *Hash) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[32]byte)(value)[:], 32)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalHash(encoder *xdr.Encoder, value *Hash) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeHash(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeHashFixedArray(encoder *xdr.Encoder, value []Hash, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeHash(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeHashArray(encoder *xdr.Encoder, value []Hash, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeHash(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeThresholds(encoder *xdr.Encoder, value *Thresholds) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[4]byte)(value)[:], 4)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalThresholds(encoder *xdr.Encoder, value *Thresholds) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeThresholds(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeThresholdsFixedArray(encoder *xdr.Encoder, value []Thresholds, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeThresholds(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeThresholdsArray(encoder *xdr.Encoder, value []Thresholds, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeThresholds(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeSequenceNumber(encoder *xdr.Encoder, value *SequenceNumber) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeUint64(encoder, (*Uint64)(value))
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalSequenceNumber(encoder *xdr.Encoder, value *SequenceNumber) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSequenceNumber(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSequenceNumberFixedArray(encoder *xdr.Encoder, value []SequenceNumber, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSequenceNumber(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSequenceNumberArray(encoder *xdr.Encoder, value []SequenceNumber, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSequenceNumber(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeCurrencyType(encoder *xdr.Encoder, value *CurrencyType) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), CurrencyTypeMap)
	return bytesWritten, err
}
func EncodeOptionalCurrencyType(encoder *xdr.Encoder, value *CurrencyType) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCurrencyType(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCurrencyTypeFixedArray(encoder *xdr.Encoder, value []CurrencyType, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCurrencyType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCurrencyTypeArray(encoder *xdr.Encoder, value []CurrencyType, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCurrencyType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeIsoCurrencyIssuer(encoder *xdr.Encoder, value *IsoCurrencyIssuer) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeFixedOpaque(encoder, value.CurrencyCode[:][:], 4)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeAccountId(encoder, &value.Issuer)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalIsoCurrencyIssuer(encoder *xdr.Encoder, value *IsoCurrencyIssuer) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeIsoCurrencyIssuer(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeIsoCurrencyIssuerFixedArray(encoder *xdr.Encoder, value []IsoCurrencyIssuer, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeIsoCurrencyIssuer(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeIsoCurrencyIssuerArray(encoder *xdr.Encoder, value []IsoCurrencyIssuer, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeIsoCurrencyIssuer(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

	if discriminant == CurrencyTypeIso4217 {
		var iso4217 IsoCurrencyIssuer
		bytesRead, err = DecodeIsoCurrencyIssuer(decoder, &iso4217)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewCurrencyIso4217(iso4217)
	}

	return totalRead, nil
}
func DecodeOptionalCurrency(decoder *xdr.Decoder, result **Currency) (int, error) {
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

	var val Currency

	bytesRead, err = DecodeCurrency(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCurrencyFixedArray(decoder *xdr.Decoder, result []Currency, size int) (int, error) {
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
		bytesRead, err = DecodeCurrency(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCurrencyArray(decoder *xdr.Decoder, result *[]Currency, maxSize int32) (int, error) {
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

	var theResult = make([]Currency, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCurrency(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCurrency(encoder *xdr.Encoder, value *Currency) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeCurrencyType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == CurrencyTypeIso4217 {
		bytesWritten, err = EncodeIsoCurrencyIssuer(encoder, value.isoCi)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalCurrency(encoder *xdr.Encoder, value *Currency) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCurrency(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCurrencyFixedArray(encoder *xdr.Encoder, value []Currency, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCurrency(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCurrencyArray(encoder *xdr.Encoder, value []Currency, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCurrency(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodePrice(encoder *xdr.Encoder, value *Price) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeInt32(encoder, &value.N)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt32(encoder, &value.D)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalPrice(encoder *xdr.Encoder, value *Price) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodePrice(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodePriceFixedArray(encoder *xdr.Encoder, value []Price, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodePrice(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodePriceArray(encoder *xdr.Encoder, value []Price, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodePrice(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum LedgerEntryType {
//       ACCOUNT,
//       TRUSTLINE,
//       OFFER
//   };
//
// ===========================================================================

type LedgerEntryType int32

const (
	LedgerEntryTypeAccount   LedgerEntryType = 0
	LedgerEntryTypeTrustline                 = 1
	LedgerEntryTypeOffer                     = 2
)

var LedgerEntryTypeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
}

func DecodeLedgerEntryType(decoder *xdr.Decoder, result *LedgerEntryType) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(LedgerEntryTypeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = LedgerEntryType(val)
	return bytesRead, err
}
func DecodeOptionalLedgerEntryType(decoder *xdr.Decoder, result **LedgerEntryType) (int, error) {
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

	var val LedgerEntryType

	bytesRead, err = DecodeLedgerEntryType(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerEntryTypeFixedArray(decoder *xdr.Decoder, result []LedgerEntryType, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerEntryType(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerEntryTypeArray(decoder *xdr.Decoder, result *[]LedgerEntryType, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerEntryType, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerEntryType(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerEntryType(encoder *xdr.Encoder, value *LedgerEntryType) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), LedgerEntryTypeMap)
	return bytesWritten, err
}
func EncodeOptionalLedgerEntryType(encoder *xdr.Encoder, value *LedgerEntryType) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerEntryType(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerEntryTypeFixedArray(encoder *xdr.Encoder, value []LedgerEntryType, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerEntryType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerEntryTypeArray(encoder *xdr.Encoder, value []LedgerEntryType, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerEntryType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct Signer
//   {
//       uint256 pubKey;
//       uint32 weight;  // really only need 1byte
//   };
//
// ===========================================================================

type Signer struct {
	PubKey Uint256
	Weight Uint32
}

func DecodeSigner(decoder *xdr.Decoder, result *Signer) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.PubKey)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.Weight)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalSigner(decoder *xdr.Decoder, result **Signer) (int, error) {
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

	var val Signer

	bytesRead, err = DecodeSigner(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSignerFixedArray(decoder *xdr.Decoder, result []Signer, size int) (int, error) {
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
		bytesRead, err = DecodeSigner(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSignerArray(decoder *xdr.Decoder, result *[]Signer, maxSize int32) (int, error) {
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

	var theResult = make([]Signer, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSigner(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeSigner(encoder *xdr.Encoder, value *Signer) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.PubKey)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.Weight)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalSigner(encoder *xdr.Encoder, value *Signer) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSigner(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSignerFixedArray(encoder *xdr.Encoder, value []Signer, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSigner(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSignerArray(encoder *xdr.Encoder, value []Signer, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSigner(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct KeyValue
//   {
//       uint32 key;
//       opaque value<64>;
//   };
//
// ===========================================================================

type KeyValue struct {
	Key   Uint32
	Value []byte
}

func DecodeKeyValue(decoder *xdr.Decoder, result *KeyValue) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint32(decoder, &result.Key)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeFixedOpaque(decoder, result.Value[:], 64)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalKeyValue(decoder *xdr.Decoder, result **KeyValue) (int, error) {
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

	var val KeyValue

	bytesRead, err = DecodeKeyValue(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeKeyValueFixedArray(decoder *xdr.Decoder, result []KeyValue, size int) (int, error) {
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
		bytesRead, err = DecodeKeyValue(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeKeyValueArray(decoder *xdr.Decoder, result *[]KeyValue, maxSize int32) (int, error) {
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

	var theResult = make([]KeyValue, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeKeyValue(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeKeyValue(encoder *xdr.Encoder, value *KeyValue) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint32(encoder, &value.Key)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeFixedOpaque(encoder, value.Value[:][:], 64)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalKeyValue(encoder *xdr.Encoder, value *KeyValue) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeKeyValue(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeKeyValueFixedArray(encoder *xdr.Encoder, value []KeyValue, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeKeyValue(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeKeyValueArray(encoder *xdr.Encoder, value []KeyValue, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeKeyValue(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum AccountFlags
//   { // masks for each flag
//       AUTH_REQUIRED_FLAG = 0x1
//   };
//
// ===========================================================================

type AccountFlags int32

const (
	AccountFlagsAuthRequiredFlag AccountFlags = 1
)

var AccountFlagsMap = map[int32]bool{
	1: true,
}

func DecodeAccountFlags(decoder *xdr.Decoder, result *AccountFlags) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(AccountFlagsMap)

	if err != nil {
		return bytesRead, err
	}
	*result = AccountFlags(val)
	return bytesRead, err
}
func DecodeOptionalAccountFlags(decoder *xdr.Decoder, result **AccountFlags) (int, error) {
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

	var val AccountFlags

	bytesRead, err = DecodeAccountFlags(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAccountFlagsFixedArray(decoder *xdr.Decoder, result []AccountFlags, size int) (int, error) {
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
		bytesRead, err = DecodeAccountFlags(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAccountFlagsArray(decoder *xdr.Decoder, result *[]AccountFlags, maxSize int32) (int, error) {
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

	var theResult = make([]AccountFlags, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAccountFlags(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAccountFlags(encoder *xdr.Encoder, value *AccountFlags) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), AccountFlagsMap)
	return bytesWritten, err
}
func EncodeOptionalAccountFlags(encoder *xdr.Encoder, value *AccountFlags) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAccountFlags(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAccountFlagsFixedArray(encoder *xdr.Encoder, value []AccountFlags, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountFlags(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAccountFlagsArray(encoder *xdr.Encoder, value []AccountFlags, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountFlags(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct AccountEntry
//   {
//       uint256 accountID;
//       int64 balance;
//       SequenceNumber seqNum;  // last sequence number used for this account
//       uint32 numSubEntries;
//       uint256 *inflationDest;
//       opaque thresholds[4]; // [weight of master|threshold1|threshold2|threshold3]
//       Signer signers<20>;
//       KeyValue data<>;
//
//       uint32 flags; // see AccountFlags
//   };
//
// ===========================================================================

type AccountEntry struct {
	AccountId     Uint256
	Balance       Int64
	SeqNum        SequenceNumber
	NumSubEntries Uint32
	InflationDest *Uint256
	Thresholds    [4]byte
	Signers       []Signer
	Data          []KeyValue
	Flags         Uint32
}

func DecodeAccountEntry(decoder *xdr.Decoder, result *AccountEntry) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.AccountId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Balance)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeSequenceNumber(decoder, &result.SeqNum)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.NumSubEntries)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	var inflationDest *Uint256
	bytesRead, err = DecodeOptionalUint256(decoder, &inflationDest)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.InflationDest = inflationDest

	bytesRead, err = DecodeFixedOpaque(decoder, result.Thresholds[:], 4)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeSignerArray(decoder, &result.Signers, 20)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeKeyValueArray(decoder, &result.Data, MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.Flags)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalAccountEntry(decoder *xdr.Decoder, result **AccountEntry) (int, error) {
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

	var val AccountEntry

	bytesRead, err = DecodeAccountEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAccountEntryFixedArray(decoder *xdr.Decoder, result []AccountEntry, size int) (int, error) {
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
		bytesRead, err = DecodeAccountEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAccountEntryArray(decoder *xdr.Decoder, result *[]AccountEntry, maxSize int32) (int, error) {
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

	var theResult = make([]AccountEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAccountEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAccountEntry(encoder *xdr.Encoder, value *AccountEntry) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.AccountId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Balance)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeSequenceNumber(encoder, &value.SeqNum)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.NumSubEntries)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOptionalUint256(encoder, value.InflationDest)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeFixedOpaque(encoder, value.Thresholds[:][:], 4)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeSignerArray(encoder, value.Signers[:], 20)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeKeyValueArray(encoder, value.Data[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.Flags)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalAccountEntry(encoder *xdr.Encoder, value *AccountEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAccountEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAccountEntryFixedArray(encoder *xdr.Encoder, value []AccountEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAccountEntryArray(encoder *xdr.Encoder, value []AccountEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TrustLineEntry
//   {
//       uint256 accountID;
//       Currency currency;
//       int64 limit;
//       int64 balance;
//       bool authorized;  // if the issuer has authorized this guy to hold its credit
//   };
//
// ===========================================================================

type TrustLineEntry struct {
	AccountId  Uint256
	Currency   Currency
	Limit      Int64
	Balance    Int64
	Authorized bool
}

func DecodeTrustLineEntry(decoder *xdr.Decoder, result *TrustLineEntry) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.AccountId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.Currency)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Limit)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Balance)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeBool(decoder, &result.Authorized)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTrustLineEntry(decoder *xdr.Decoder, result **TrustLineEntry) (int, error) {
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

	var val TrustLineEntry

	bytesRead, err = DecodeTrustLineEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTrustLineEntryFixedArray(decoder *xdr.Decoder, result []TrustLineEntry, size int) (int, error) {
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
		bytesRead, err = DecodeTrustLineEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTrustLineEntryArray(decoder *xdr.Decoder, result *[]TrustLineEntry, maxSize int32) (int, error) {
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

	var theResult = make([]TrustLineEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTrustLineEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTrustLineEntry(encoder *xdr.Encoder, value *TrustLineEntry) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.AccountId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.Currency)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Limit)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Balance)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeBool(encoder, &value.Authorized)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTrustLineEntry(encoder *xdr.Encoder, value *TrustLineEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTrustLineEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTrustLineEntryFixedArray(encoder *xdr.Encoder, value []TrustLineEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTrustLineEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTrustLineEntryArray(encoder *xdr.Encoder, value []TrustLineEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTrustLineEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct OfferEntry
//   {
//       uint256 accountID;
//       uint64 offerID;
//       Currency takerGets;  // A
//       Currency takerPays;  // B
//       int64 amount;    // amount of A
//
//       Price price;
//                       // price of A in terms of B
//                       // price*10,000,000
//                       // price=AmountB/AmountA=priceNumerator/priceDenominator
//                       // price is after fees
//       int32 flags;
//   };
//
// ===========================================================================

type OfferEntry struct {
	AccountId Uint256
	OfferId   Uint64
	TakerGets Currency
	TakerPays Currency
	Amount    Int64
	Price     Price
	Flags     Int32
}

func DecodeOfferEntry(decoder *xdr.Decoder, result *OfferEntry) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.AccountId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint64(decoder, &result.OfferId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.TakerGets)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.TakerPays)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Amount)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodePrice(decoder, &result.Price)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt32(decoder, &result.Flags)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalOfferEntry(decoder *xdr.Decoder, result **OfferEntry) (int, error) {
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

	var val OfferEntry

	bytesRead, err = DecodeOfferEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeOfferEntryFixedArray(decoder *xdr.Decoder, result []OfferEntry, size int) (int, error) {
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
		bytesRead, err = DecodeOfferEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeOfferEntryArray(decoder *xdr.Decoder, result *[]OfferEntry, maxSize int32) (int, error) {
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

	var theResult = make([]OfferEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeOfferEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeOfferEntry(encoder *xdr.Encoder, value *OfferEntry) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.AccountId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint64(encoder, &value.OfferId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.TakerGets)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.TakerPays)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Amount)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodePrice(encoder, &value.Price)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt32(encoder, &value.Flags)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalOfferEntry(encoder *xdr.Encoder, value *OfferEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeOfferEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOfferEntryFixedArray(encoder *xdr.Encoder, value []OfferEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeOfferEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeOfferEntryArray(encoder *xdr.Encoder, value []OfferEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeOfferEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union LedgerEntry switch (LedgerEntryType type)
//   {
//       case ACCOUNT:
//           AccountEntry account;
//
//       case TRUSTLINE:
//           TrustLineEntry trustLine;
//
//       case OFFER:
//           OfferEntry offer;
//   };
//
// ===========================================================================

type LedgerEntry struct {
	aType     LedgerEntryType
	account   *AccountEntry
	trustLine *TrustLineEntry
	offer     *OfferEntry
}

func NewLedgerEntryAccount(val AccountEntry) LedgerEntry {
	return LedgerEntry{
		aType:   LedgerEntryTypeAccount,
		account: &val,
	}
}
func NewLedgerEntryTrustline(val TrustLineEntry) LedgerEntry {
	return LedgerEntry{
		aType:     LedgerEntryTypeTrustline,
		trustLine: &val,
	}
}
func NewLedgerEntryOffer(val OfferEntry) LedgerEntry {
	return LedgerEntry{
		aType: LedgerEntryTypeOffer,
		offer: &val,
	}
}
func (u *LedgerEntry) Type() LedgerEntryType {
	return u.aType
}
func (u *LedgerEntry) Account() AccountEntry {
	return *u.account
}
func (u *LedgerEntry) TrustLine() TrustLineEntry {
	return *u.trustLine
}
func (u *LedgerEntry) Offer() OfferEntry {
	return *u.offer
}
func DecodeLedgerEntry(decoder *xdr.Decoder, result *LedgerEntry) (int, error) {
	var (
		discriminant LedgerEntryType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeLedgerEntryType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == LedgerEntryTypeAccount {
		var account AccountEntry
		bytesRead, err = DecodeAccountEntry(decoder, &account)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewLedgerEntryAccount(account)
	}

	if discriminant == LedgerEntryTypeTrustline {
		var trustline TrustLineEntry
		bytesRead, err = DecodeTrustLineEntry(decoder, &trustline)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewLedgerEntryTrustline(trustline)
	}

	if discriminant == LedgerEntryTypeOffer {
		var offer OfferEntry
		bytesRead, err = DecodeOfferEntry(decoder, &offer)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewLedgerEntryOffer(offer)
	}

	return totalRead, nil
}
func DecodeOptionalLedgerEntry(decoder *xdr.Decoder, result **LedgerEntry) (int, error) {
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

	var val LedgerEntry

	bytesRead, err = DecodeLedgerEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerEntryFixedArray(decoder *xdr.Decoder, result []LedgerEntry, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerEntryArray(decoder *xdr.Decoder, result *[]LedgerEntry, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerEntry(encoder *xdr.Encoder, value *LedgerEntry) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeLedgerEntryType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == LedgerEntryTypeAccount {
		bytesWritten, err = EncodeAccountEntry(encoder, value.account)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == LedgerEntryTypeTrustline {
		bytesWritten, err = EncodeTrustLineEntry(encoder, value.trustLine)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == LedgerEntryTypeOffer {
		bytesWritten, err = EncodeOfferEntry(encoder, value.offer)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalLedgerEntry(encoder *xdr.Encoder, value *LedgerEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerEntryFixedArray(encoder *xdr.Encoder, value []LedgerEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerEntryArray(encoder *xdr.Encoder, value []LedgerEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct LedgerHeader
//   {
//       Hash previousLedgerHash;// hash of the previous ledger header
//       Hash txSetHash;         // the tx set that was SCP confirmed
//       Hash txSetResultHash;   // the TransactionResultSet that led to this ledger
//       Hash clfHash;           // hash of the ledger state
//
//       uint32 ledgerSeq;       // sequence number of this ledger
//       uint64 closeTime;       // network close time
//
//       int64 totalCoins;       // total number of stroops in existence
//
//       int64 feePool;          // fees burned since last inflation run
//       uint32 inflationSeq;    // inflation sequence number
//
//       uint64 idPool;          // last used global ID, used for generating objects
//
//       int32 baseFee;          // base fee per operation in stroops
//       int32 baseReserve;      // account base reserve in stroops
//
//   };
//
// ===========================================================================

type LedgerHeader struct {
	PreviousLedgerHash Hash
	TxSetHash          Hash
	TxSetResultHash    Hash
	ClfHash            Hash
	LedgerSeq          Uint32
	CloseTime          Uint64
	TotalCoins         Int64
	FeePool            Int64
	InflationSeq       Uint32
	IdPool             Uint64
	BaseFee            Int32
	BaseReserve        Int32
}

func DecodeLedgerHeader(decoder *xdr.Decoder, result *LedgerHeader) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeHash(decoder, &result.PreviousLedgerHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeHash(decoder, &result.TxSetHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeHash(decoder, &result.TxSetResultHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeHash(decoder, &result.ClfHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.LedgerSeq)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint64(decoder, &result.CloseTime)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.TotalCoins)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.FeePool)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.InflationSeq)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint64(decoder, &result.IdPool)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt32(decoder, &result.BaseFee)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt32(decoder, &result.BaseReserve)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalLedgerHeader(decoder *xdr.Decoder, result **LedgerHeader) (int, error) {
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

	var val LedgerHeader

	bytesRead, err = DecodeLedgerHeader(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerHeaderFixedArray(decoder *xdr.Decoder, result []LedgerHeader, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerHeader(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerHeaderArray(decoder *xdr.Decoder, result *[]LedgerHeader, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerHeader, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerHeader(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerHeader(encoder *xdr.Encoder, value *LedgerHeader) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeHash(encoder, &value.PreviousLedgerHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeHash(encoder, &value.TxSetHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeHash(encoder, &value.TxSetResultHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeHash(encoder, &value.ClfHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.LedgerSeq)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint64(encoder, &value.CloseTime)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.TotalCoins)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.FeePool)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.InflationSeq)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint64(encoder, &value.IdPool)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt32(encoder, &value.BaseFee)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt32(encoder, &value.BaseReserve)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalLedgerHeader(encoder *xdr.Encoder, value *LedgerHeader) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerHeader(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerHeaderFixedArray(encoder *xdr.Encoder, value []LedgerHeader, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerHeader(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerHeaderArray(encoder *xdr.Encoder, value []LedgerHeader, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerHeader(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct {
//               uint256 accountID;
//           }
//
// ===========================================================================

type LedgerKeyAccount struct {
	AccountId Uint256
}

func DecodeLedgerKeyAccount(decoder *xdr.Decoder, result *LedgerKeyAccount) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.AccountId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalLedgerKeyAccount(decoder *xdr.Decoder, result **LedgerKeyAccount) (int, error) {
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

	var val LedgerKeyAccount

	bytesRead, err = DecodeLedgerKeyAccount(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerKeyAccountFixedArray(decoder *xdr.Decoder, result []LedgerKeyAccount, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerKeyAccount(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerKeyAccountArray(decoder *xdr.Decoder, result *[]LedgerKeyAccount, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerKeyAccount, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerKeyAccount(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerKeyAccount(encoder *xdr.Encoder, value *LedgerKeyAccount) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.AccountId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalLedgerKeyAccount(encoder *xdr.Encoder, value *LedgerKeyAccount) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerKeyAccount(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerKeyAccountFixedArray(encoder *xdr.Encoder, value []LedgerKeyAccount, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKeyAccount(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerKeyAccountArray(encoder *xdr.Encoder, value []LedgerKeyAccount, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKeyAccount(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct {
//               uint256 accountID;
//               Currency currency;
//           }
//
// ===========================================================================

type LedgerKeyTrustLine struct {
	AccountId Uint256
	Currency  Currency
}

func DecodeLedgerKeyTrustLine(decoder *xdr.Decoder, result *LedgerKeyTrustLine) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.AccountId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.Currency)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalLedgerKeyTrustLine(decoder *xdr.Decoder, result **LedgerKeyTrustLine) (int, error) {
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

	var val LedgerKeyTrustLine

	bytesRead, err = DecodeLedgerKeyTrustLine(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerKeyTrustLineFixedArray(decoder *xdr.Decoder, result []LedgerKeyTrustLine, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerKeyTrustLine(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerKeyTrustLineArray(decoder *xdr.Decoder, result *[]LedgerKeyTrustLine, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerKeyTrustLine, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerKeyTrustLine(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerKeyTrustLine(encoder *xdr.Encoder, value *LedgerKeyTrustLine) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.AccountId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.Currency)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalLedgerKeyTrustLine(encoder *xdr.Encoder, value *LedgerKeyTrustLine) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerKeyTrustLine(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerKeyTrustLineFixedArray(encoder *xdr.Encoder, value []LedgerKeyTrustLine, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKeyTrustLine(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerKeyTrustLineArray(encoder *xdr.Encoder, value []LedgerKeyTrustLine, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKeyTrustLine(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct {
//               uint256 accountID; // GRAYDON: ok to drop this since offerID is unique now?
//               uint64 offerID;
//           }
//
// ===========================================================================

type LedgerKeyOffer struct {
	AccountId Uint256
	OfferId   Uint64
}

func DecodeLedgerKeyOffer(decoder *xdr.Decoder, result *LedgerKeyOffer) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.AccountId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint64(decoder, &result.OfferId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalLedgerKeyOffer(decoder *xdr.Decoder, result **LedgerKeyOffer) (int, error) {
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

	var val LedgerKeyOffer

	bytesRead, err = DecodeLedgerKeyOffer(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerKeyOfferFixedArray(decoder *xdr.Decoder, result []LedgerKeyOffer, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerKeyOffer(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerKeyOfferArray(decoder *xdr.Decoder, result *[]LedgerKeyOffer, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerKeyOffer, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerKeyOffer(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerKeyOffer(encoder *xdr.Encoder, value *LedgerKeyOffer) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.AccountId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint64(encoder, &value.OfferId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalLedgerKeyOffer(encoder *xdr.Encoder, value *LedgerKeyOffer) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerKeyOffer(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerKeyOfferFixedArray(encoder *xdr.Encoder, value []LedgerKeyOffer, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKeyOffer(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerKeyOfferArray(encoder *xdr.Encoder, value []LedgerKeyOffer, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKeyOffer(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union LedgerKey switch (LedgerEntryType type)
//   {
//       case ACCOUNT:
//           struct {
//               uint256 accountID;
//           } account;
//
//       case TRUSTLINE:
//           struct {
//               uint256 accountID;
//               Currency currency;
//           } trustLine;
//
//       case OFFER:
//           struct {
//               uint256 accountID; // GRAYDON: ok to drop this since offerID is unique now?
//               uint64 offerID;
//           } offer;
//   };
//
// ===========================================================================

type LedgerKey struct {
	aType     LedgerEntryType
	account   *LedgerKeyAccount
	trustLine *LedgerKeyTrustLine
	offer     *LedgerKeyOffer
}

func NewLedgerKeyAccount(val LedgerKeyAccount) LedgerKey {
	return LedgerKey{
		aType:   LedgerEntryTypeAccount,
		account: &val,
	}
}
func NewLedgerKeyTrustline(val LedgerKeyTrustLine) LedgerKey {
	return LedgerKey{
		aType:     LedgerEntryTypeTrustline,
		trustLine: &val,
	}
}
func NewLedgerKeyOffer(val LedgerKeyOffer) LedgerKey {
	return LedgerKey{
		aType: LedgerEntryTypeOffer,
		offer: &val,
	}
}
func (u *LedgerKey) Type() LedgerEntryType {
	return u.aType
}
func (u *LedgerKey) Account() LedgerKeyAccount {
	return *u.account
}
func (u *LedgerKey) TrustLine() LedgerKeyTrustLine {
	return *u.trustLine
}
func (u *LedgerKey) Offer() LedgerKeyOffer {
	return *u.offer
}
func DecodeLedgerKey(decoder *xdr.Decoder, result *LedgerKey) (int, error) {
	var (
		discriminant LedgerEntryType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeLedgerEntryType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == LedgerEntryTypeAccount {
		var account LedgerKeyAccount
		bytesRead, err = DecodeLedgerKeyAccount(decoder, &account)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewLedgerKeyAccount(account)
	}

	if discriminant == LedgerEntryTypeTrustline {
		var trustline LedgerKeyTrustLine
		bytesRead, err = DecodeLedgerKeyTrustLine(decoder, &trustline)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewLedgerKeyTrustline(trustline)
	}

	if discriminant == LedgerEntryTypeOffer {
		var offer LedgerKeyOffer
		bytesRead, err = DecodeLedgerKeyOffer(decoder, &offer)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewLedgerKeyOffer(offer)
	}

	return totalRead, nil
}
func DecodeOptionalLedgerKey(decoder *xdr.Decoder, result **LedgerKey) (int, error) {
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

	var val LedgerKey

	bytesRead, err = DecodeLedgerKey(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerKeyFixedArray(decoder *xdr.Decoder, result []LedgerKey, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerKey(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerKeyArray(decoder *xdr.Decoder, result *[]LedgerKey, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerKey, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerKey(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerKey(encoder *xdr.Encoder, value *LedgerKey) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeLedgerEntryType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == LedgerEntryTypeAccount {
		bytesWritten, err = EncodeLedgerKeyAccount(encoder, value.account)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == LedgerEntryTypeTrustline {
		bytesWritten, err = EncodeLedgerKeyTrustLine(encoder, value.trustLine)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == LedgerEntryTypeOffer {
		bytesWritten, err = EncodeLedgerKeyOffer(encoder, value.offer)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalLedgerKey(encoder *xdr.Encoder, value *LedgerKey) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerKey(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerKeyFixedArray(encoder *xdr.Encoder, value []LedgerKey, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKey(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerKeyArray(encoder *xdr.Encoder, value []LedgerKey, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerKey(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum CLFType {
//       LIVEENTRY,
//       DEADENTRY
//   };
//
// ===========================================================================

type ClfType int32

const (
	ClfTypeLiveentry ClfType = 0
	ClfTypeDeadentry         = 1
)

var ClfTypeMap = map[int32]bool{
	0: true,
	1: true,
}

func DecodeClfType(decoder *xdr.Decoder, result *ClfType) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(ClfTypeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = ClfType(val)
	return bytesRead, err
}
func DecodeOptionalClfType(decoder *xdr.Decoder, result **ClfType) (int, error) {
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

	var val ClfType

	bytesRead, err = DecodeClfType(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeClfTypeFixedArray(decoder *xdr.Decoder, result []ClfType, size int) (int, error) {
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
		bytesRead, err = DecodeClfType(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeClfTypeArray(decoder *xdr.Decoder, result *[]ClfType, maxSize int32) (int, error) {
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

	var theResult = make([]ClfType, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeClfType(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeClfType(encoder *xdr.Encoder, value *ClfType) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), ClfTypeMap)
	return bytesWritten, err
}
func EncodeOptionalClfType(encoder *xdr.Encoder, value *ClfType) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeClfType(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeClfTypeFixedArray(encoder *xdr.Encoder, value []ClfType, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeClfType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeClfTypeArray(encoder *xdr.Encoder, value []ClfType, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeClfType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union CLFEntry switch (CLFType type)
//   {
//       case LIVEENTRY:
//           LedgerEntry liveEntry;
//
//       case DEADENTRY:
//           LedgerKey deadEntry;
//   };
//
// ===========================================================================

type ClfEntry struct {
	aType     ClfType
	liveEntry *LedgerEntry
	deadEntry *LedgerKey
}

func NewClfEntryLiveentry(val LedgerEntry) ClfEntry {
	return ClfEntry{
		aType:     ClfTypeLiveentry,
		liveEntry: &val,
	}
}
func NewClfEntryDeadentry(val LedgerKey) ClfEntry {
	return ClfEntry{
		aType:     ClfTypeDeadentry,
		deadEntry: &val,
	}
}
func (u *ClfEntry) Type() ClfType {
	return u.aType
}
func (u *ClfEntry) LiveEntry() LedgerEntry {
	return *u.liveEntry
}
func (u *ClfEntry) DeadEntry() LedgerKey {
	return *u.deadEntry
}
func DecodeClfEntry(decoder *xdr.Decoder, result *ClfEntry) (int, error) {
	var (
		discriminant ClfType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeClfType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == ClfTypeLiveentry {
		var liveentry LedgerEntry
		bytesRead, err = DecodeLedgerEntry(decoder, &liveentry)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewClfEntryLiveentry(liveentry)
	}

	if discriminant == ClfTypeDeadentry {
		var deadentry LedgerKey
		bytesRead, err = DecodeLedgerKey(decoder, &deadentry)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewClfEntryDeadentry(deadentry)
	}

	return totalRead, nil
}
func DecodeOptionalClfEntry(decoder *xdr.Decoder, result **ClfEntry) (int, error) {
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

	var val ClfEntry

	bytesRead, err = DecodeClfEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeClfEntryFixedArray(decoder *xdr.Decoder, result []ClfEntry, size int) (int, error) {
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
		bytesRead, err = DecodeClfEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeClfEntryArray(decoder *xdr.Decoder, result *[]ClfEntry, maxSize int32) (int, error) {
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

	var theResult = make([]ClfEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeClfEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeClfEntry(encoder *xdr.Encoder, value *ClfEntry) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeClfType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == ClfTypeLiveentry {
		bytesWritten, err = EncodeLedgerEntry(encoder, value.liveEntry)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == ClfTypeDeadentry {
		bytesWritten, err = EncodeLedgerKey(encoder, value.deadEntry)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalClfEntry(encoder *xdr.Encoder, value *ClfEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeClfEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeClfEntryFixedArray(encoder *xdr.Encoder, value []ClfEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeClfEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeClfEntryArray(encoder *xdr.Encoder, value []ClfEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeClfEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionSet
//   {
//       Hash previousLedgerHash;
//       TransactionEnvelope txs<5000>;
//   };
//
// ===========================================================================

type TransactionSet struct {
	PreviousLedgerHash Hash
	Txes               []TransactionEnvelope
}

func DecodeTransactionSet(decoder *xdr.Decoder, result *TransactionSet) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeHash(decoder, &result.PreviousLedgerHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeTransactionEnvelopeArray(decoder, &result.Txes, 5000)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionSet(decoder *xdr.Decoder, result **TransactionSet) (int, error) {
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

	var val TransactionSet

	bytesRead, err = DecodeTransactionSet(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionSetFixedArray(decoder *xdr.Decoder, result []TransactionSet, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionSet(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionSetArray(decoder *xdr.Decoder, result *[]TransactionSet, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionSet, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionSet(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionSet(encoder *xdr.Encoder, value *TransactionSet) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeHash(encoder, &value.PreviousLedgerHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeTransactionEnvelopeArray(encoder, value.Txes[:], 5000)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionSet(encoder *xdr.Encoder, value *TransactionSet) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionSet(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionSetFixedArray(encoder *xdr.Encoder, value []TransactionSet, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionSet(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionSetArray(encoder *xdr.Encoder, value []TransactionSet, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionSet(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionResultPair
//   {
//       Hash transactionHash;
//       TransactionResult result;   // result for the transaction
//   };
//
// ===========================================================================

type TransactionResultPair struct {
	TransactionHash Hash
	Result          TransactionResult
}

func DecodeTransactionResultPair(decoder *xdr.Decoder, result *TransactionResultPair) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeHash(decoder, &result.TransactionHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeTransactionResult(decoder, &result.Result)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionResultPair(decoder *xdr.Decoder, result **TransactionResultPair) (int, error) {
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

	var val TransactionResultPair

	bytesRead, err = DecodeTransactionResultPair(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionResultPairFixedArray(decoder *xdr.Decoder, result []TransactionResultPair, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionResultPair(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionResultPairArray(decoder *xdr.Decoder, result *[]TransactionResultPair, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionResultPair, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionResultPair(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionResultPair(encoder *xdr.Encoder, value *TransactionResultPair) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeHash(encoder, &value.TransactionHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeTransactionResult(encoder, &value.Result)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionResultPair(encoder *xdr.Encoder, value *TransactionResultPair) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionResultPair(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionResultPairFixedArray(encoder *xdr.Encoder, value []TransactionResultPair, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultPair(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionResultPairArray(encoder *xdr.Encoder, value []TransactionResultPair, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultPair(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionResultSet
//   {
//       TransactionResultPair results<5000>;
//   };
//
// ===========================================================================

type TransactionResultSet struct {
	Results []TransactionResultPair
}

func DecodeTransactionResultSet(decoder *xdr.Decoder, result *TransactionResultSet) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeTransactionResultPairArray(decoder, &result.Results, 5000)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionResultSet(decoder *xdr.Decoder, result **TransactionResultSet) (int, error) {
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

	var val TransactionResultSet

	bytesRead, err = DecodeTransactionResultSet(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionResultSetFixedArray(decoder *xdr.Decoder, result []TransactionResultSet, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionResultSet(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionResultSetArray(decoder *xdr.Decoder, result *[]TransactionResultSet, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionResultSet, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionResultSet(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionResultSet(encoder *xdr.Encoder, value *TransactionResultSet) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeTransactionResultPairArray(encoder, value.Results[:], 5000)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionResultSet(encoder *xdr.Encoder, value *TransactionResultSet) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionResultSet(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionResultSetFixedArray(encoder *xdr.Encoder, value []TransactionResultSet, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultSet(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionResultSetArray(encoder *xdr.Encoder, value []TransactionResultSet, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultSet(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionMeta
//   {
//       CLFEntry entries<>;
//   };
//
// ===========================================================================

type TransactionMeta struct {
	Entries []ClfEntry
}

func DecodeTransactionMeta(decoder *xdr.Decoder, result *TransactionMeta) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeClfEntryArray(decoder, &result.Entries, MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionMeta(decoder *xdr.Decoder, result **TransactionMeta) (int, error) {
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

	var val TransactionMeta

	bytesRead, err = DecodeTransactionMeta(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionMetaFixedArray(decoder *xdr.Decoder, result []TransactionMeta, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionMeta(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionMetaArray(decoder *xdr.Decoder, result *[]TransactionMeta, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionMeta, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionMeta(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionMeta(encoder *xdr.Encoder, value *TransactionMeta) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeClfEntryArray(encoder, value.Entries[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionMeta(encoder *xdr.Encoder, value *TransactionMeta) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionMeta(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionMetaFixedArray(encoder *xdr.Encoder, value []TransactionMeta, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionMeta(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionMetaArray(encoder *xdr.Encoder, value []TransactionMeta, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionMeta(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionHistoryEntry
//   {
//       uint32 ledgerSeq;
//       TransactionSet txSet;
//   };
//
// ===========================================================================

type TransactionHistoryEntry struct {
	LedgerSeq Uint32
	TxSet     TransactionSet
}

func DecodeTransactionHistoryEntry(decoder *xdr.Decoder, result *TransactionHistoryEntry) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint32(decoder, &result.LedgerSeq)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeTransactionSet(decoder, &result.TxSet)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionHistoryEntry(decoder *xdr.Decoder, result **TransactionHistoryEntry) (int, error) {
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

	var val TransactionHistoryEntry

	bytesRead, err = DecodeTransactionHistoryEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionHistoryEntryFixedArray(decoder *xdr.Decoder, result []TransactionHistoryEntry, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionHistoryEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionHistoryEntryArray(decoder *xdr.Decoder, result *[]TransactionHistoryEntry, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionHistoryEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionHistoryEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionHistoryEntry(encoder *xdr.Encoder, value *TransactionHistoryEntry) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint32(encoder, &value.LedgerSeq)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeTransactionSet(encoder, &value.TxSet)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionHistoryEntry(encoder *xdr.Encoder, value *TransactionHistoryEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionHistoryEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionHistoryEntryFixedArray(encoder *xdr.Encoder, value []TransactionHistoryEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionHistoryEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionHistoryEntryArray(encoder *xdr.Encoder, value []TransactionHistoryEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionHistoryEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionHistoryResultEntry
//   {
//       uint32 ledgerSeq;
//       TransactionResultSet txResultSet;
//   };
//
// ===========================================================================

type TransactionHistoryResultEntry struct {
	LedgerSeq   Uint32
	TxResultSet TransactionResultSet
}

func DecodeTransactionHistoryResultEntry(decoder *xdr.Decoder, result *TransactionHistoryResultEntry) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint32(decoder, &result.LedgerSeq)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeTransactionResultSet(decoder, &result.TxResultSet)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionHistoryResultEntry(decoder *xdr.Decoder, result **TransactionHistoryResultEntry) (int, error) {
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

	var val TransactionHistoryResultEntry

	bytesRead, err = DecodeTransactionHistoryResultEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionHistoryResultEntryFixedArray(decoder *xdr.Decoder, result []TransactionHistoryResultEntry, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionHistoryResultEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionHistoryResultEntryArray(decoder *xdr.Decoder, result *[]TransactionHistoryResultEntry, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionHistoryResultEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionHistoryResultEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionHistoryResultEntry(encoder *xdr.Encoder, value *TransactionHistoryResultEntry) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint32(encoder, &value.LedgerSeq)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeTransactionResultSet(encoder, &value.TxResultSet)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionHistoryResultEntry(encoder *xdr.Encoder, value *TransactionHistoryResultEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionHistoryResultEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionHistoryResultEntryFixedArray(encoder *xdr.Encoder, value []TransactionHistoryResultEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionHistoryResultEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionHistoryResultEntryArray(encoder *xdr.Encoder, value []TransactionHistoryResultEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionHistoryResultEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct LedgerHeaderHistoryEntry
//   {
//       Hash hash;
//       LedgerHeader header;
//   };
//
// ===========================================================================

type LedgerHeaderHistoryEntry struct {
	Hash   Hash
	Header LedgerHeader
}

func DecodeLedgerHeaderHistoryEntry(decoder *xdr.Decoder, result *LedgerHeaderHistoryEntry) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeHash(decoder, &result.Hash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeLedgerHeader(decoder, &result.Header)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalLedgerHeaderHistoryEntry(decoder *xdr.Decoder, result **LedgerHeaderHistoryEntry) (int, error) {
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

	var val LedgerHeaderHistoryEntry

	bytesRead, err = DecodeLedgerHeaderHistoryEntry(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeLedgerHeaderHistoryEntryFixedArray(decoder *xdr.Decoder, result []LedgerHeaderHistoryEntry, size int) (int, error) {
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
		bytesRead, err = DecodeLedgerHeaderHistoryEntry(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeLedgerHeaderHistoryEntryArray(decoder *xdr.Decoder, result *[]LedgerHeaderHistoryEntry, maxSize int32) (int, error) {
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

	var theResult = make([]LedgerHeaderHistoryEntry, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeLedgerHeaderHistoryEntry(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeLedgerHeaderHistoryEntry(encoder *xdr.Encoder, value *LedgerHeaderHistoryEntry) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeHash(encoder, &value.Hash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeLedgerHeader(encoder, &value.Header)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalLedgerHeaderHistoryEntry(encoder *xdr.Encoder, value *LedgerHeaderHistoryEntry) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeLedgerHeaderHistoryEntry(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeLedgerHeaderHistoryEntryFixedArray(encoder *xdr.Encoder, value []LedgerHeaderHistoryEntry, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerHeaderHistoryEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeLedgerHeaderHistoryEntryArray(encoder *xdr.Encoder, value []LedgerHeaderHistoryEntry, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeLedgerHeaderHistoryEntry(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct DecoratedSignature
//   {
//       opaque hint[4];     // first 4 bytes of the public key, used as a hint
//       uint512 signature;  // actual signature
//   };
//
// ===========================================================================

type DecoratedSignature struct {
	Hint      [4]byte
	Signature Uint512
}

func DecodeDecoratedSignature(decoder *xdr.Decoder, result *DecoratedSignature) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeFixedOpaque(decoder, result.Hint[:], 4)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint512(decoder, &result.Signature)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalDecoratedSignature(decoder *xdr.Decoder, result **DecoratedSignature) (int, error) {
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

	var val DecoratedSignature

	bytesRead, err = DecodeDecoratedSignature(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeDecoratedSignatureFixedArray(decoder *xdr.Decoder, result []DecoratedSignature, size int) (int, error) {
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
		bytesRead, err = DecodeDecoratedSignature(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeDecoratedSignatureArray(decoder *xdr.Decoder, result *[]DecoratedSignature, maxSize int32) (int, error) {
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

	var theResult = make([]DecoratedSignature, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeDecoratedSignature(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeDecoratedSignature(encoder *xdr.Encoder, value *DecoratedSignature) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeFixedOpaque(encoder, value.Hint[:][:], 4)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint512(encoder, &value.Signature)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalDecoratedSignature(encoder *xdr.Encoder, value *DecoratedSignature) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeDecoratedSignature(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeDecoratedSignatureFixedArray(encoder *xdr.Encoder, value []DecoratedSignature, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeDecoratedSignature(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeDecoratedSignatureArray(encoder *xdr.Encoder, value []DecoratedSignature, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeDecoratedSignature(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum OperationType
//   {
//       PAYMENT,
//       CREATE_OFFER,
//       CANCEL_OFFER,
//       SET_OPTIONS,
//       CHANGE_TRUST,
//       ALLOW_TRUST,
//       ACCOUNT_MERGE,
//       INFLATION
//   };
//
// ===========================================================================

type OperationType int32

const (
	OperationTypePayment      OperationType = 0
	OperationTypeCreateOffer                = 1
	OperationTypeCancelOffer                = 2
	OperationTypeSetOption                  = 3
	OperationTypeChangeTrust                = 4
	OperationTypeAllowTrust                 = 5
	OperationTypeAccountMerge               = 6
	OperationTypeInflation                  = 7
)

var OperationTypeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
	7: true,
}

func DecodeOperationType(decoder *xdr.Decoder, result *OperationType) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(OperationTypeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = OperationType(val)
	return bytesRead, err
}
func DecodeOptionalOperationType(decoder *xdr.Decoder, result **OperationType) (int, error) {
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

	var val OperationType

	bytesRead, err = DecodeOperationType(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeOperationTypeFixedArray(decoder *xdr.Decoder, result []OperationType, size int) (int, error) {
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
		bytesRead, err = DecodeOperationType(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeOperationTypeArray(decoder *xdr.Decoder, result *[]OperationType, maxSize int32) (int, error) {
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

	var theResult = make([]OperationType, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeOperationType(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeOperationType(encoder *xdr.Encoder, value *OperationType) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), OperationTypeMap)
	return bytesWritten, err
}
func EncodeOptionalOperationType(encoder *xdr.Encoder, value *OperationType) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeOperationType(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOperationTypeFixedArray(encoder *xdr.Encoder, value []OperationType, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeOperationTypeArray(encoder *xdr.Encoder, value []OperationType, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct PaymentOp
//   {
//       AccountID destination;
//       Currency currency;      // what they end up with
//       int64 amount;           // amount they end up with
//       Currency path<5>;        // what hops it must go through to get there
//       int64 sendMax;          // the maximum amount of the source currency to
//                               // send (excluding fees).
//                               // The operation will fail if can't be met
//       opaque memo<32>;
//       opaque sourceMemo<32>;  // used to return a payment
//   };
//
// ===========================================================================

type PaymentOp struct {
	Destination AccountId
	Currency    Currency
	Amount      Int64
	Path        []Currency
	SendMax     Int64
	Memo        []byte
	SourceMemo  []byte
}

func DecodePaymentOp(decoder *xdr.Decoder, result *PaymentOp) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeAccountId(decoder, &result.Destination)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.Currency)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Amount)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrencyArray(decoder, &result.Path, 5)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.SendMax)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeFixedOpaque(decoder, result.Memo[:], 32)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeFixedOpaque(decoder, result.SourceMemo[:], 32)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalPaymentOp(decoder *xdr.Decoder, result **PaymentOp) (int, error) {
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

	var val PaymentOp

	bytesRead, err = DecodePaymentOp(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodePaymentOpFixedArray(decoder *xdr.Decoder, result []PaymentOp, size int) (int, error) {
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
		bytesRead, err = DecodePaymentOp(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodePaymentOpArray(decoder *xdr.Decoder, result *[]PaymentOp, maxSize int32) (int, error) {
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

	var theResult = make([]PaymentOp, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodePaymentOp(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodePaymentOp(encoder *xdr.Encoder, value *PaymentOp) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeAccountId(encoder, &value.Destination)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.Currency)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Amount)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrencyArray(encoder, value.Path[:], 5)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.SendMax)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeFixedOpaque(encoder, value.Memo[:][:], 32)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeFixedOpaque(encoder, value.SourceMemo[:][:], 32)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalPaymentOp(encoder *xdr.Encoder, value *PaymentOp) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodePaymentOp(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodePaymentOpFixedArray(encoder *xdr.Encoder, value []PaymentOp, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodePaymentOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodePaymentOpArray(encoder *xdr.Encoder, value []PaymentOp, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodePaymentOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct CreateOfferOp
//   {
//       Currency takerGets;
//       Currency takerPays;
//       int64 amount;        // amount taker gets
//       Price price;         // =takerPaysAmount/takerGetsAmount
//
//       uint64 offerID;      // set if you want to change an existing offer
//       uint32 flags;        // passive: only take offers that cross this. not offers that match it
//   };
//
// ===========================================================================

type CreateOfferOp struct {
	TakerGets Currency
	TakerPays Currency
	Amount    Int64
	Price     Price
	OfferId   Uint64
	Flags     Uint32
}

func DecodeCreateOfferOp(decoder *xdr.Decoder, result *CreateOfferOp) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeCurrency(decoder, &result.TakerGets)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.TakerPays)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Amount)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodePrice(decoder, &result.Price)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint64(decoder, &result.OfferId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.Flags)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalCreateOfferOp(decoder *xdr.Decoder, result **CreateOfferOp) (int, error) {
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

	var val CreateOfferOp

	bytesRead, err = DecodeCreateOfferOp(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCreateOfferOpFixedArray(decoder *xdr.Decoder, result []CreateOfferOp, size int) (int, error) {
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
		bytesRead, err = DecodeCreateOfferOp(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCreateOfferOpArray(decoder *xdr.Decoder, result *[]CreateOfferOp, maxSize int32) (int, error) {
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

	var theResult = make([]CreateOfferOp, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCreateOfferOp(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCreateOfferOp(encoder *xdr.Encoder, value *CreateOfferOp) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeCurrency(encoder, &value.TakerGets)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.TakerPays)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Amount)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodePrice(encoder, &value.Price)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint64(encoder, &value.OfferId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.Flags)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalCreateOfferOp(encoder *xdr.Encoder, value *CreateOfferOp) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCreateOfferOp(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCreateOfferOpFixedArray(encoder *xdr.Encoder, value []CreateOfferOp, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCreateOfferOpArray(encoder *xdr.Encoder, value []CreateOfferOp, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct SetOptionsOp
//   {
//       AccountID* inflationDest;
//       uint32*    clearFlags;
//       uint32*    setFlags;
//       KeyValue* data;
//       Thresholds* thresholds;
//       Signer* signer;
//   };
//
// ===========================================================================

type SetOptionsOp struct {
	InflationDest *AccountId
	ClearFlags    *Uint32
	SetFlags      *Uint32
	Data          *KeyValue
	Thresholds    *Thresholds
	Signer        *Signer
}

func DecodeSetOptionsOp(decoder *xdr.Decoder, result *SetOptionsOp) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	var inflationDest *AccountId
	bytesRead, err = DecodeOptionalAccountId(decoder, &inflationDest)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.InflationDest = inflationDest

	var clearFlags *Uint32
	bytesRead, err = DecodeOptionalUint32(decoder, &clearFlags)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.ClearFlags = clearFlags

	var setFlags *Uint32
	bytesRead, err = DecodeOptionalUint32(decoder, &setFlags)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.SetFlags = setFlags

	var data *KeyValue
	bytesRead, err = DecodeOptionalKeyValue(decoder, &data)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.Data = data

	var thresholds *Thresholds
	bytesRead, err = DecodeOptionalThresholds(decoder, &thresholds)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.Thresholds = thresholds

	var signer *Signer
	bytesRead, err = DecodeOptionalSigner(decoder, &signer)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.Signer = signer

	return totalRead, nil
}
func DecodeOptionalSetOptionsOp(decoder *xdr.Decoder, result **SetOptionsOp) (int, error) {
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

	var val SetOptionsOp

	bytesRead, err = DecodeSetOptionsOp(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSetOptionsOpFixedArray(decoder *xdr.Decoder, result []SetOptionsOp, size int) (int, error) {
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
		bytesRead, err = DecodeSetOptionsOp(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSetOptionsOpArray(decoder *xdr.Decoder, result *[]SetOptionsOp, maxSize int32) (int, error) {
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

	var theResult = make([]SetOptionsOp, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSetOptionsOp(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeSetOptionsOp(encoder *xdr.Encoder, value *SetOptionsOp) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeOptionalAccountId(encoder, value.InflationDest)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOptionalUint32(encoder, value.ClearFlags)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOptionalUint32(encoder, value.SetFlags)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOptionalKeyValue(encoder, value.Data)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOptionalThresholds(encoder, value.Thresholds)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOptionalSigner(encoder, value.Signer)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalSetOptionsOp(encoder *xdr.Encoder, value *SetOptionsOp) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSetOptionsOp(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSetOptionsOpFixedArray(encoder *xdr.Encoder, value []SetOptionsOp, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSetOptionsOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSetOptionsOpArray(encoder *xdr.Encoder, value []SetOptionsOp, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSetOptionsOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct ChangeTrustOp
//   {
//       Currency line;
//       int64 limit;
//   };
//
// ===========================================================================

type ChangeTrustOp struct {
	Line  Currency
	Limit Int64
}

func DecodeChangeTrustOp(decoder *xdr.Decoder, result *ChangeTrustOp) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeCurrency(decoder, &result.Line)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Limit)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalChangeTrustOp(decoder *xdr.Decoder, result **ChangeTrustOp) (int, error) {
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

	var val ChangeTrustOp

	bytesRead, err = DecodeChangeTrustOp(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeChangeTrustOpFixedArray(decoder *xdr.Decoder, result []ChangeTrustOp, size int) (int, error) {
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
		bytesRead, err = DecodeChangeTrustOp(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeChangeTrustOpArray(decoder *xdr.Decoder, result *[]ChangeTrustOp, maxSize int32) (int, error) {
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

	var theResult = make([]ChangeTrustOp, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeChangeTrustOp(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeChangeTrustOp(encoder *xdr.Encoder, value *ChangeTrustOp) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeCurrency(encoder, &value.Line)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Limit)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalChangeTrustOp(encoder *xdr.Encoder, value *ChangeTrustOp) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeChangeTrustOp(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeChangeTrustOpFixedArray(encoder *xdr.Encoder, value []ChangeTrustOp, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeChangeTrustOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeChangeTrustOpArray(encoder *xdr.Encoder, value []ChangeTrustOp, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeChangeTrustOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union switch(CurrencyType type)
//       {
//           // NATIVE is not allowed
//           case ISO4217:
//               opaque currencyCode[4];
//
//           // add other currency types here in the future
//       }
//
// ===========================================================================

type AllowTrustOpCurrency struct {
	aType        CurrencyType
	currencyCode *[4]byte
}

func NewAllowTrustOpCurrencyIso4217(val [4]byte) AllowTrustOpCurrency {
	return AllowTrustOpCurrency{
		aType:        CurrencyTypeIso4217,
		currencyCode: &val,
	}
}
func (u *AllowTrustOpCurrency) Type() CurrencyType {
	return u.aType
}
func (u *AllowTrustOpCurrency) CurrencyCode() [4]byte {
	return *u.currencyCode
}
func DecodeAllowTrustOpCurrency(decoder *xdr.Decoder, result *AllowTrustOpCurrency) (int, error) {
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

	if discriminant == CurrencyTypeIso4217 {
		var iso4217 [4]byte
		bytesRead, err = DecodeFixedOpaque(decoder, iso4217[:], 4)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewAllowTrustOpCurrencyIso4217(iso4217)
	}

	return totalRead, nil
}
func DecodeOptionalAllowTrustOpCurrency(decoder *xdr.Decoder, result **AllowTrustOpCurrency) (int, error) {
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

	var val AllowTrustOpCurrency

	bytesRead, err = DecodeAllowTrustOpCurrency(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAllowTrustOpCurrencyFixedArray(decoder *xdr.Decoder, result []AllowTrustOpCurrency, size int) (int, error) {
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
		bytesRead, err = DecodeAllowTrustOpCurrency(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAllowTrustOpCurrencyArray(decoder *xdr.Decoder, result *[]AllowTrustOpCurrency, maxSize int32) (int, error) {
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

	var theResult = make([]AllowTrustOpCurrency, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAllowTrustOpCurrency(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAllowTrustOpCurrency(encoder *xdr.Encoder, value *AllowTrustOpCurrency) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeCurrencyType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == CurrencyTypeIso4217 {
		bytesWritten, err = EncodeFixedOpaque(encoder, value.currencyCode[:], 4)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalAllowTrustOpCurrency(encoder *xdr.Encoder, value *AllowTrustOpCurrency) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAllowTrustOpCurrency(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAllowTrustOpCurrencyFixedArray(encoder *xdr.Encoder, value []AllowTrustOpCurrency, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustOpCurrency(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAllowTrustOpCurrencyArray(encoder *xdr.Encoder, value []AllowTrustOpCurrency, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustOpCurrency(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct AllowTrustOp
//   {
//       AccountID trustor;
//       union switch(CurrencyType type)
//       {
//           // NATIVE is not allowed
//           case ISO4217:
//               opaque currencyCode[4];
//
//           // add other currency types here in the future
//       } currency;
//
//       bool authorize;
//   };
//
// ===========================================================================

type AllowTrustOp struct {
	Trustor   AccountId
	Currency  AllowTrustOpCurrency
	Authorize bool
}

func DecodeAllowTrustOp(decoder *xdr.Decoder, result *AllowTrustOp) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeAccountId(decoder, &result.Trustor)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeAllowTrustOpCurrency(decoder, &result.Currency)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeBool(decoder, &result.Authorize)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalAllowTrustOp(decoder *xdr.Decoder, result **AllowTrustOp) (int, error) {
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

	var val AllowTrustOp

	bytesRead, err = DecodeAllowTrustOp(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAllowTrustOpFixedArray(decoder *xdr.Decoder, result []AllowTrustOp, size int) (int, error) {
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
		bytesRead, err = DecodeAllowTrustOp(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAllowTrustOpArray(decoder *xdr.Decoder, result *[]AllowTrustOp, maxSize int32) (int, error) {
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

	var theResult = make([]AllowTrustOp, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAllowTrustOp(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAllowTrustOp(encoder *xdr.Encoder, value *AllowTrustOp) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeAccountId(encoder, &value.Trustor)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeAllowTrustOpCurrency(encoder, &value.Currency)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeBool(encoder, &value.Authorize)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalAllowTrustOp(encoder *xdr.Encoder, value *AllowTrustOp) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAllowTrustOp(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAllowTrustOpFixedArray(encoder *xdr.Encoder, value []AllowTrustOp, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAllowTrustOpArray(encoder *xdr.Encoder, value []AllowTrustOp, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustOp(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union switch (OperationType type)
//       {
//           case PAYMENT:
//               PaymentOp paymentOp;
//           case CREATE_OFFER:
//               CreateOfferOp createOfferOp;
//           case CANCEL_OFFER:
//               uint64 offerID;
//           case SET_OPTIONS:
//               SetOptionsOp setOptionsOp;
//           case CHANGE_TRUST:
//               ChangeTrustOp changeTrustOp;
//           case ALLOW_TRUST:
//               AllowTrustOp allowTrustOp;
//           case ACCOUNT_MERGE:
//               uint256 destination;
//           case INFLATION:
//               uint32 inflationSeq;
//       }
//
// ===========================================================================

type OperationBody struct {
	aType         OperationType
	paymentOp     *PaymentOp
	createOfferOp *CreateOfferOp
	offerId       *Uint64
	setOptionsOp  *SetOptionsOp
	changeTrustOp *ChangeTrustOp
	allowTrustOp  *AllowTrustOp
	destination   *Uint256
	inflationSeq  *Uint32
}

func NewOperationBodyPayment(val PaymentOp) OperationBody {
	return OperationBody{
		aType:     OperationTypePayment,
		paymentOp: &val,
	}
}
func NewOperationBodyCreateOffer(val CreateOfferOp) OperationBody {
	return OperationBody{
		aType:         OperationTypeCreateOffer,
		createOfferOp: &val,
	}
}
func NewOperationBodyCancelOffer(val Uint64) OperationBody {
	return OperationBody{
		aType:   OperationTypeCancelOffer,
		offerId: &val,
	}
}
func NewOperationBodySetOption(val SetOptionsOp) OperationBody {
	return OperationBody{
		aType:        OperationTypeSetOption,
		setOptionsOp: &val,
	}
}
func NewOperationBodyChangeTrust(val ChangeTrustOp) OperationBody {
	return OperationBody{
		aType:         OperationTypeChangeTrust,
		changeTrustOp: &val,
	}
}
func NewOperationBodyAllowTrust(val AllowTrustOp) OperationBody {
	return OperationBody{
		aType:        OperationTypeAllowTrust,
		allowTrustOp: &val,
	}
}
func NewOperationBodyAccountMerge(val Uint256) OperationBody {
	return OperationBody{
		aType:       OperationTypeAccountMerge,
		destination: &val,
	}
}
func NewOperationBodyInflation(val Uint32) OperationBody {
	return OperationBody{
		aType:        OperationTypeInflation,
		inflationSeq: &val,
	}
}
func (u *OperationBody) Type() OperationType {
	return u.aType
}
func (u *OperationBody) PaymentOp() PaymentOp {
	return *u.paymentOp
}
func (u *OperationBody) CreateOfferOp() CreateOfferOp {
	return *u.createOfferOp
}
func (u *OperationBody) OfferId() Uint64 {
	return *u.offerId
}
func (u *OperationBody) SetOptionsOp() SetOptionsOp {
	return *u.setOptionsOp
}
func (u *OperationBody) ChangeTrustOp() ChangeTrustOp {
	return *u.changeTrustOp
}
func (u *OperationBody) AllowTrustOp() AllowTrustOp {
	return *u.allowTrustOp
}
func (u *OperationBody) Destination() Uint256 {
	return *u.destination
}
func (u *OperationBody) InflationSeq() Uint32 {
	return *u.inflationSeq
}
func DecodeOperationBody(decoder *xdr.Decoder, result *OperationBody) (int, error) {
	var (
		discriminant OperationType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeOperationType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == OperationTypePayment {
		var payment PaymentOp
		bytesRead, err = DecodePaymentOp(decoder, &payment)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodyPayment(payment)
	}

	if discriminant == OperationTypeCreateOffer {
		var createOffer CreateOfferOp
		bytesRead, err = DecodeCreateOfferOp(decoder, &createOffer)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodyCreateOffer(createOffer)
	}

	if discriminant == OperationTypeCancelOffer {
		var cancelOffer Uint64
		bytesRead, err = DecodeUint64(decoder, &cancelOffer)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodyCancelOffer(cancelOffer)
	}

	if discriminant == OperationTypeSetOption {
		var setOptions SetOptionsOp
		bytesRead, err = DecodeSetOptionsOp(decoder, &setOptions)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodySetOption(setOptions)
	}

	if discriminant == OperationTypeChangeTrust {
		var changeTrust ChangeTrustOp
		bytesRead, err = DecodeChangeTrustOp(decoder, &changeTrust)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodyChangeTrust(changeTrust)
	}

	if discriminant == OperationTypeAllowTrust {
		var allowTrust AllowTrustOp
		bytesRead, err = DecodeAllowTrustOp(decoder, &allowTrust)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodyAllowTrust(allowTrust)
	}

	if discriminant == OperationTypeAccountMerge {
		var accountMerge Uint256
		bytesRead, err = DecodeUint256(decoder, &accountMerge)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodyAccountMerge(accountMerge)
	}

	if discriminant == OperationTypeInflation {
		var inflation Uint32
		bytesRead, err = DecodeUint32(decoder, &inflation)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationBodyInflation(inflation)
	}

	return totalRead, nil
}
func DecodeOptionalOperationBody(decoder *xdr.Decoder, result **OperationBody) (int, error) {
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

	var val OperationBody

	bytesRead, err = DecodeOperationBody(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeOperationBodyFixedArray(decoder *xdr.Decoder, result []OperationBody, size int) (int, error) {
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
		bytesRead, err = DecodeOperationBody(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeOperationBodyArray(decoder *xdr.Decoder, result *[]OperationBody, maxSize int32) (int, error) {
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

	var theResult = make([]OperationBody, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeOperationBody(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeOperationBody(encoder *xdr.Encoder, value *OperationBody) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeOperationType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == OperationTypePayment {
		bytesWritten, err = EncodePaymentOp(encoder, value.paymentOp)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeCreateOffer {
		bytesWritten, err = EncodeCreateOfferOp(encoder, value.createOfferOp)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeCancelOffer {
		bytesWritten, err = EncodeUint64(encoder, value.offerId)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeSetOption {
		bytesWritten, err = EncodeSetOptionsOp(encoder, value.setOptionsOp)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeChangeTrust {
		bytesWritten, err = EncodeChangeTrustOp(encoder, value.changeTrustOp)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeAllowTrust {
		bytesWritten, err = EncodeAllowTrustOp(encoder, value.allowTrustOp)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeAccountMerge {
		bytesWritten, err = EncodeUint256(encoder, value.destination)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeInflation {
		bytesWritten, err = EncodeUint32(encoder, value.inflationSeq)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalOperationBody(encoder *xdr.Encoder, value *OperationBody) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeOperationBody(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOperationBodyFixedArray(encoder *xdr.Encoder, value []OperationBody, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationBody(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeOperationBodyArray(encoder *xdr.Encoder, value []OperationBody, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationBody(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct Operation
//   {
//       AccountID* sourceAccount; // defaults to the account from the transaction
//       union switch (OperationType type)
//       {
//           case PAYMENT:
//               PaymentOp paymentOp;
//           case CREATE_OFFER:
//               CreateOfferOp createOfferOp;
//           case CANCEL_OFFER:
//               uint64 offerID;
//           case SET_OPTIONS:
//               SetOptionsOp setOptionsOp;
//           case CHANGE_TRUST:
//               ChangeTrustOp changeTrustOp;
//           case ALLOW_TRUST:
//               AllowTrustOp allowTrustOp;
//           case ACCOUNT_MERGE:
//               uint256 destination;
//           case INFLATION:
//               uint32 inflationSeq;
//       } body;
//   };
//
// ===========================================================================

type Operation struct {
	SourceAccount *AccountId
	Body          OperationBody
}

func DecodeOperation(decoder *xdr.Decoder, result *Operation) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	var sourceAccount *AccountId
	bytesRead, err = DecodeOptionalAccountId(decoder, &sourceAccount)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.SourceAccount = sourceAccount

	bytesRead, err = DecodeOperationBody(decoder, &result.Body)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalOperation(decoder *xdr.Decoder, result **Operation) (int, error) {
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

	var val Operation

	bytesRead, err = DecodeOperation(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeOperationFixedArray(decoder *xdr.Decoder, result []Operation, size int) (int, error) {
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
		bytesRead, err = DecodeOperation(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeOperationArray(decoder *xdr.Decoder, result *[]Operation, maxSize int32) (int, error) {
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

	var theResult = make([]Operation, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeOperation(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeOperation(encoder *xdr.Encoder, value *Operation) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeOptionalAccountId(encoder, value.SourceAccount)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOperationBody(encoder, &value.Body)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalOperation(encoder *xdr.Encoder, value *Operation) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeOperation(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOperationFixedArray(encoder *xdr.Encoder, value []Operation, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperation(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeOperationArray(encoder *xdr.Encoder, value []Operation, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperation(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct Transaction
//   {
//       AccountID account;
//       int32 maxFee;
//       SequenceNumber seqNum;
//       uint32 minLedger;
//       uint32 maxLedger;
//
//       Operation operations<100>;
//   };
//
// ===========================================================================

type Transaction struct {
	Account    AccountId
	MaxFee     Int32
	SeqNum     SequenceNumber
	MinLedger  Uint32
	MaxLedger  Uint32
	Operations []Operation
}

func DecodeTransaction(decoder *xdr.Decoder, result *Transaction) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeAccountId(decoder, &result.Account)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt32(decoder, &result.MaxFee)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeSequenceNumber(decoder, &result.SeqNum)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.MinLedger)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.MaxLedger)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeOperationArray(decoder, &result.Operations, 100)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransaction(decoder *xdr.Decoder, result **Transaction) (int, error) {
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

	var val Transaction

	bytesRead, err = DecodeTransaction(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionFixedArray(decoder *xdr.Decoder, result []Transaction, size int) (int, error) {
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
		bytesRead, err = DecodeTransaction(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionArray(decoder *xdr.Decoder, result *[]Transaction, maxSize int32) (int, error) {
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

	var theResult = make([]Transaction, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransaction(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransaction(encoder *xdr.Encoder, value *Transaction) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeAccountId(encoder, &value.Account)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt32(encoder, &value.MaxFee)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeSequenceNumber(encoder, &value.SeqNum)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.MinLedger)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.MaxLedger)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOperationArray(encoder, value.Operations[:], 100)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransaction(encoder *xdr.Encoder, value *Transaction) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransaction(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionFixedArray(encoder *xdr.Encoder, value []Transaction, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransaction(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionArray(encoder *xdr.Encoder, value []Transaction, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransaction(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionEnvelope
//   {
//       Transaction tx;
//       DecoratedSignature signatures<20>;
//   };
//
// ===========================================================================

type TransactionEnvelope struct {
	Tx         Transaction
	Signatures []DecoratedSignature
}

func DecodeTransactionEnvelope(decoder *xdr.Decoder, result *TransactionEnvelope) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeTransaction(decoder, &result.Tx)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeDecoratedSignatureArray(decoder, &result.Signatures, 20)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionEnvelope(decoder *xdr.Decoder, result **TransactionEnvelope) (int, error) {
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

	var val TransactionEnvelope

	bytesRead, err = DecodeTransactionEnvelope(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionEnvelopeFixedArray(decoder *xdr.Decoder, result []TransactionEnvelope, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionEnvelope(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionEnvelopeArray(decoder *xdr.Decoder, result *[]TransactionEnvelope, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionEnvelope, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionEnvelope(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionEnvelope(encoder *xdr.Encoder, value *TransactionEnvelope) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeTransaction(encoder, &value.Tx)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeDecoratedSignatureArray(encoder, value.Signatures[:], 20)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionEnvelope(encoder *xdr.Encoder, value *TransactionEnvelope) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionEnvelope(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionEnvelopeFixedArray(encoder *xdr.Encoder, value []TransactionEnvelope, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionEnvelope(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionEnvelopeArray(encoder *xdr.Encoder, value []TransactionEnvelope, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionEnvelope(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct ClaimOfferAtom
//   {
//       AccountID offerOwner;
//       uint64 offerID;
//       Currency currencyClaimed; // redundant but emited for clarity
//       int64 amountClaimed;
//       // should we also include the amount that the owner gets in return?
//   };
//
// ===========================================================================

type ClaimOfferAtom struct {
	OfferOwner      AccountId
	OfferId         Uint64
	CurrencyClaimed Currency
	AmountClaimed   Int64
}

func DecodeClaimOfferAtom(decoder *xdr.Decoder, result *ClaimOfferAtom) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeAccountId(decoder, &result.OfferOwner)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint64(decoder, &result.OfferId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.CurrencyClaimed)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.AmountClaimed)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalClaimOfferAtom(decoder *xdr.Decoder, result **ClaimOfferAtom) (int, error) {
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

	var val ClaimOfferAtom

	bytesRead, err = DecodeClaimOfferAtom(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeClaimOfferAtomFixedArray(decoder *xdr.Decoder, result []ClaimOfferAtom, size int) (int, error) {
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
		bytesRead, err = DecodeClaimOfferAtom(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeClaimOfferAtomArray(decoder *xdr.Decoder, result *[]ClaimOfferAtom, maxSize int32) (int, error) {
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

	var theResult = make([]ClaimOfferAtom, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeClaimOfferAtom(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeClaimOfferAtom(encoder *xdr.Encoder, value *ClaimOfferAtom) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeAccountId(encoder, &value.OfferOwner)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint64(encoder, &value.OfferId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.CurrencyClaimed)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.AmountClaimed)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalClaimOfferAtom(encoder *xdr.Encoder, value *ClaimOfferAtom) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeClaimOfferAtom(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeClaimOfferAtomFixedArray(encoder *xdr.Encoder, value []ClaimOfferAtom, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeClaimOfferAtom(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeClaimOfferAtomArray(encoder *xdr.Encoder, value []ClaimOfferAtom, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeClaimOfferAtom(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum OperationResultCode
//   {
//       opSKIP,
//       opINNER,
//
//       opBAD_AUTH, // not enough signatures to perform operation
//       opNO_ACCOUNT
//   };
//
// ===========================================================================

type OperationResultCode int32

const (
	OperationResultCodeOpSkip      OperationResultCode = 0
	OperationResultCodeOpInner                         = 1
	OperationResultCodeOpBadAuth                       = 2
	OperationResultCodeOpNoAccount                     = 3
)

var OperationResultCodeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
}

func DecodeOperationResultCode(decoder *xdr.Decoder, result *OperationResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(OperationResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = OperationResultCode(val)
	return bytesRead, err
}
func DecodeOptionalOperationResultCode(decoder *xdr.Decoder, result **OperationResultCode) (int, error) {
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

	var val OperationResultCode

	bytesRead, err = DecodeOperationResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeOperationResultCodeFixedArray(decoder *xdr.Decoder, result []OperationResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeOperationResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeOperationResultCodeArray(decoder *xdr.Decoder, result *[]OperationResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]OperationResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeOperationResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeOperationResultCode(encoder *xdr.Encoder, value *OperationResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), OperationResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalOperationResultCode(encoder *xdr.Encoder, value *OperationResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeOperationResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOperationResultCodeFixedArray(encoder *xdr.Encoder, value []OperationResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeOperationResultCodeArray(encoder *xdr.Encoder, value []OperationResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union switch(OperationType type)
//           {
//               case PAYMENT:
//                   PaymentResult paymentResult;
//               case CREATE_OFFER:
//                   CreateOfferResult createOfferResult;
//               case CANCEL_OFFER:
//                   CancelOfferResult cancelOfferResult;
//               case SET_OPTIONS:
//                   SetOptionsResult setOptionsResult;
//               case CHANGE_TRUST:
//                   ChangeTrustResult changeTrustResult;
//               case ALLOW_TRUST:
//                   AllowTrustResult allowTrustResult;
//               case ACCOUNT_MERGE:
//                   AccountMergeResult accountMergeResult;
//               case INFLATION:
//                   InflationResult inflationResult;
//           }
//
// ===========================================================================

type OperationResultTr struct {
	aType              OperationType
	paymentResult      *PaymentResult
	createOfferResult  *CreateOfferResult
	cancelOfferResult  *CancelOfferResult
	setOptionsResult   *SetOptionsResult
	changeTrustResult  *ChangeTrustResult
	allowTrustResult   *AllowTrustResult
	accountMergeResult *AccountMergeResult
	inflationResult    *InflationResult
}

func NewOperationResultTrPayment(val PaymentResult) OperationResultTr {
	return OperationResultTr{
		aType:         OperationTypePayment,
		paymentResult: &val,
	}
}
func NewOperationResultTrCreateOffer(val CreateOfferResult) OperationResultTr {
	return OperationResultTr{
		aType:             OperationTypeCreateOffer,
		createOfferResult: &val,
	}
}
func NewOperationResultTrCancelOffer(val CancelOfferResult) OperationResultTr {
	return OperationResultTr{
		aType:             OperationTypeCancelOffer,
		cancelOfferResult: &val,
	}
}
func NewOperationResultTrSetOption(val SetOptionsResult) OperationResultTr {
	return OperationResultTr{
		aType:            OperationTypeSetOption,
		setOptionsResult: &val,
	}
}
func NewOperationResultTrChangeTrust(val ChangeTrustResult) OperationResultTr {
	return OperationResultTr{
		aType:             OperationTypeChangeTrust,
		changeTrustResult: &val,
	}
}
func NewOperationResultTrAllowTrust(val AllowTrustResult) OperationResultTr {
	return OperationResultTr{
		aType:            OperationTypeAllowTrust,
		allowTrustResult: &val,
	}
}
func NewOperationResultTrAccountMerge(val AccountMergeResult) OperationResultTr {
	return OperationResultTr{
		aType:              OperationTypeAccountMerge,
		accountMergeResult: &val,
	}
}
func NewOperationResultTrInflation(val InflationResult) OperationResultTr {
	return OperationResultTr{
		aType:           OperationTypeInflation,
		inflationResult: &val,
	}
}
func (u *OperationResultTr) Type() OperationType {
	return u.aType
}
func (u *OperationResultTr) PaymentResult() PaymentResult {
	return *u.paymentResult
}
func (u *OperationResultTr) CreateOfferResult() CreateOfferResult {
	return *u.createOfferResult
}
func (u *OperationResultTr) CancelOfferResult() CancelOfferResult {
	return *u.cancelOfferResult
}
func (u *OperationResultTr) SetOptionsResult() SetOptionsResult {
	return *u.setOptionsResult
}
func (u *OperationResultTr) ChangeTrustResult() ChangeTrustResult {
	return *u.changeTrustResult
}
func (u *OperationResultTr) AllowTrustResult() AllowTrustResult {
	return *u.allowTrustResult
}
func (u *OperationResultTr) AccountMergeResult() AccountMergeResult {
	return *u.accountMergeResult
}
func (u *OperationResultTr) InflationResult() InflationResult {
	return *u.inflationResult
}
func DecodeOperationResultTr(decoder *xdr.Decoder, result *OperationResultTr) (int, error) {
	var (
		discriminant OperationType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeOperationType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == OperationTypePayment {
		var payment PaymentResult
		bytesRead, err = DecodePaymentResult(decoder, &payment)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrPayment(payment)
	}

	if discriminant == OperationTypeCreateOffer {
		var createOffer CreateOfferResult
		bytesRead, err = DecodeCreateOfferResult(decoder, &createOffer)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrCreateOffer(createOffer)
	}

	if discriminant == OperationTypeCancelOffer {
		var cancelOffer CancelOfferResult
		bytesRead, err = DecodeCancelOfferResult(decoder, &cancelOffer)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrCancelOffer(cancelOffer)
	}

	if discriminant == OperationTypeSetOption {
		var setOptions SetOptionsResult
		bytesRead, err = DecodeSetOptionsResult(decoder, &setOptions)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrSetOption(setOptions)
	}

	if discriminant == OperationTypeChangeTrust {
		var changeTrust ChangeTrustResult
		bytesRead, err = DecodeChangeTrustResult(decoder, &changeTrust)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrChangeTrust(changeTrust)
	}

	if discriminant == OperationTypeAllowTrust {
		var allowTrust AllowTrustResult
		bytesRead, err = DecodeAllowTrustResult(decoder, &allowTrust)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrAllowTrust(allowTrust)
	}

	if discriminant == OperationTypeAccountMerge {
		var accountMerge AccountMergeResult
		bytesRead, err = DecodeAccountMergeResult(decoder, &accountMerge)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrAccountMerge(accountMerge)
	}

	if discriminant == OperationTypeInflation {
		var inflation InflationResult
		bytesRead, err = DecodeInflationResult(decoder, &inflation)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultTrInflation(inflation)
	}

	return totalRead, nil
}
func DecodeOptionalOperationResultTr(decoder *xdr.Decoder, result **OperationResultTr) (int, error) {
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

	var val OperationResultTr

	bytesRead, err = DecodeOperationResultTr(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeOperationResultTrFixedArray(decoder *xdr.Decoder, result []OperationResultTr, size int) (int, error) {
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
		bytesRead, err = DecodeOperationResultTr(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeOperationResultTrArray(decoder *xdr.Decoder, result *[]OperationResultTr, maxSize int32) (int, error) {
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

	var theResult = make([]OperationResultTr, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeOperationResultTr(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeOperationResultTr(encoder *xdr.Encoder, value *OperationResultTr) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeOperationType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == OperationTypePayment {
		bytesWritten, err = EncodePaymentResult(encoder, value.paymentResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeCreateOffer {
		bytesWritten, err = EncodeCreateOfferResult(encoder, value.createOfferResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeCancelOffer {
		bytesWritten, err = EncodeCancelOfferResult(encoder, value.cancelOfferResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeSetOption {
		bytesWritten, err = EncodeSetOptionsResult(encoder, value.setOptionsResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeChangeTrust {
		bytesWritten, err = EncodeChangeTrustResult(encoder, value.changeTrustResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeAllowTrust {
		bytesWritten, err = EncodeAllowTrustResult(encoder, value.allowTrustResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeAccountMerge {
		bytesWritten, err = EncodeAccountMergeResult(encoder, value.accountMergeResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == OperationTypeInflation {
		bytesWritten, err = EncodeInflationResult(encoder, value.inflationResult)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalOperationResultTr(encoder *xdr.Encoder, value *OperationResultTr) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeOperationResultTr(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOperationResultTrFixedArray(encoder *xdr.Encoder, value []OperationResultTr, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationResultTr(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeOperationResultTrArray(encoder *xdr.Encoder, value []OperationResultTr, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationResultTr(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union OperationResult switch(OperationResultCode code)
//   {
//       case opINNER:
//           union switch(OperationType type)
//           {
//               case PAYMENT:
//                   PaymentResult paymentResult;
//               case CREATE_OFFER:
//                   CreateOfferResult createOfferResult;
//               case CANCEL_OFFER:
//                   CancelOfferResult cancelOfferResult;
//               case SET_OPTIONS:
//                   SetOptionsResult setOptionsResult;
//               case CHANGE_TRUST:
//                   ChangeTrustResult changeTrustResult;
//               case ALLOW_TRUST:
//                   AllowTrustResult allowTrustResult;
//               case ACCOUNT_MERGE:
//                   AccountMergeResult accountMergeResult;
//               case INFLATION:
//                   InflationResult inflationResult;
//           } tr;
//       default:
//           void;
//   };
//
// ===========================================================================

type OperationResult struct {
	code OperationResultCode
	tr   *OperationResultTr
}

func NewOperationResultOpSkip() OperationResult {
	return OperationResult{
		code: OperationResultCodeOpSkip,
	}
}
func NewOperationResultOpInner(val OperationResultTr) OperationResult {
	return OperationResult{
		code: OperationResultCodeOpInner,
		tr:   &val,
	}
}
func NewOperationResultOpBadAuth() OperationResult {
	return OperationResult{
		code: OperationResultCodeOpBadAuth,
	}
}
func NewOperationResultOpNoAccount() OperationResult {
	return OperationResult{
		code: OperationResultCodeOpNoAccount,
	}
}
func (u *OperationResult) Code() OperationResultCode {
	return u.code
}
func (u *OperationResult) Tr() OperationResultTr {
	return *u.tr
}
func DecodeOperationResult(decoder *xdr.Decoder, result *OperationResult) (int, error) {
	var (
		discriminant OperationResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeOperationResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == OperationResultCodeOpInner {
		var opInner OperationResultTr
		bytesRead, err = DecodeOperationResultTr(decoder, &opInner)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewOperationResultOpInner(opInner)
	}

	return totalRead, nil
}
func DecodeOptionalOperationResult(decoder *xdr.Decoder, result **OperationResult) (int, error) {
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

	var val OperationResult

	bytesRead, err = DecodeOperationResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeOperationResultFixedArray(decoder *xdr.Decoder, result []OperationResult, size int) (int, error) {
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
		bytesRead, err = DecodeOperationResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeOperationResultArray(decoder *xdr.Decoder, result *[]OperationResult, maxSize int32) (int, error) {
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

	var theResult = make([]OperationResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeOperationResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeOperationResult(encoder *xdr.Encoder, value *OperationResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeOperationResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.code == OperationResultCodeOpInner {
		bytesWritten, err = EncodeOperationResultTr(encoder, value.tr)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalOperationResult(encoder *xdr.Encoder, value *OperationResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeOperationResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOperationResultFixedArray(encoder *xdr.Encoder, value []OperationResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeOperationResultArray(encoder *xdr.Encoder, value []OperationResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeOperationResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum TransactionResultCode
//   {
//       txSUCCESS,
//       txFAILED,
//       txBAD_LEDGER,
//       txDUPLICATE,
//       txMALFORMED,
//       txBAD_SEQ,
//
//       txBAD_AUTH, // not enough signatures to perform transaction
//       txINSUFFICIENT_BALANCE,
//       txNO_ACCOUNT,
//       txINSUFFICIENT_FEE // max fee is too small
//   };
//
// ===========================================================================

type TransactionResultCode int32

const (
	TransactionResultCodeTxSuccess             TransactionResultCode = 0
	TransactionResultCodeTxFailed                                    = 1
	TransactionResultCodeTxBadLedger                                 = 2
	TransactionResultCodeTxDuplicate                                 = 3
	TransactionResultCodeTxMalformed                                 = 4
	TransactionResultCodeTxBadSeq                                    = 5
	TransactionResultCodeTxBadAuth                                   = 6
	TransactionResultCodeTxInsufficientBalance                       = 7
	TransactionResultCodeTxNoAccount                                 = 8
	TransactionResultCodeTxInsufficientFee                           = 9
)

var TransactionResultCodeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
	7: true,
	8: true,
	9: true,
}

func DecodeTransactionResultCode(decoder *xdr.Decoder, result *TransactionResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(TransactionResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = TransactionResultCode(val)
	return bytesRead, err
}
func DecodeOptionalTransactionResultCode(decoder *xdr.Decoder, result **TransactionResultCode) (int, error) {
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

	var val TransactionResultCode

	bytesRead, err = DecodeTransactionResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionResultCodeFixedArray(decoder *xdr.Decoder, result []TransactionResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionResultCodeArray(decoder *xdr.Decoder, result *[]TransactionResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionResultCode(encoder *xdr.Encoder, value *TransactionResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), TransactionResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalTransactionResultCode(encoder *xdr.Encoder, value *TransactionResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionResultCodeFixedArray(encoder *xdr.Encoder, value []TransactionResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionResultCodeArray(encoder *xdr.Encoder, value []TransactionResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union switch(TransactionResultCode code)
//       {
//           case txSUCCESS:
//           case txFAILED:
//               OperationResult results<>;
//           default:
//               void;
//       }
//
// ===========================================================================

type TransactionResultResult struct {
	code    TransactionResultCode
	results *[]OperationResult
}

func NewTransactionResultResultTxSuccess(val []OperationResult) TransactionResultResult {
	return TransactionResultResult{
		code:    TransactionResultCodeTxSuccess,
		results: &val,
	}
}
func NewTransactionResultResultTxFailed(val []OperationResult) TransactionResultResult {
	return TransactionResultResult{
		code:    TransactionResultCodeTxFailed,
		results: &val,
	}
}
func NewTransactionResultResultTxBadLedger() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxBadLedger,
	}
}
func NewTransactionResultResultTxDuplicate() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxDuplicate,
	}
}
func NewTransactionResultResultTxMalformed() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxMalformed,
	}
}
func NewTransactionResultResultTxBadSeq() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxBadSeq,
	}
}
func NewTransactionResultResultTxBadAuth() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxBadAuth,
	}
}
func NewTransactionResultResultTxInsufficientBalance() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxInsufficientBalance,
	}
}
func NewTransactionResultResultTxNoAccount() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxNoAccount,
	}
}
func NewTransactionResultResultTxInsufficientFee() TransactionResultResult {
	return TransactionResultResult{
		code: TransactionResultCodeTxInsufficientFee,
	}
}
func (u *TransactionResultResult) Code() TransactionResultCode {
	return u.code
}
func (u *TransactionResultResult) Results() []OperationResult {
	return *u.results
}
func DecodeTransactionResultResult(decoder *xdr.Decoder, result *TransactionResultResult) (int, error) {
	var (
		discriminant TransactionResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeTransactionResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == TransactionResultCodeTxSuccess {
		var txSuccess []OperationResult
		bytesRead, err = DecodeOperationResultArray(decoder, &txSuccess, MaxXdrElements)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewTransactionResultResultTxSuccess(txSuccess)
	}

	if discriminant == TransactionResultCodeTxFailed {
		var txFailed []OperationResult
		bytesRead, err = DecodeOperationResultArray(decoder, &txFailed, MaxXdrElements)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewTransactionResultResultTxFailed(txFailed)
	}

	return totalRead, nil
}
func DecodeOptionalTransactionResultResult(decoder *xdr.Decoder, result **TransactionResultResult) (int, error) {
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

	var val TransactionResultResult

	bytesRead, err = DecodeTransactionResultResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionResultResultFixedArray(decoder *xdr.Decoder, result []TransactionResultResult, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionResultResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionResultResultArray(decoder *xdr.Decoder, result *[]TransactionResultResult, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionResultResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionResultResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionResultResult(encoder *xdr.Encoder, value *TransactionResultResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeTransactionResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.code == TransactionResultCodeTxSuccess {
		bytesWritten, err = EncodeOperationResultArray(encoder, value.results, MaxXdrElements)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.code == TransactionResultCodeTxFailed {
		bytesWritten, err = EncodeOperationResultArray(encoder, value.results, MaxXdrElements)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalTransactionResultResult(encoder *xdr.Encoder, value *TransactionResultResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionResultResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionResultResultFixedArray(encoder *xdr.Encoder, value []TransactionResultResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionResultResultArray(encoder *xdr.Encoder, value []TransactionResultResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResultResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct TransactionResult
//   {
//       int64 feeCharged;
//
//       union switch(TransactionResultCode code)
//       {
//           case txSUCCESS:
//           case txFAILED:
//               OperationResult results<>;
//           default:
//               void;
//       } result;
//   };
//
// ===========================================================================

type TransactionResult struct {
	FeeCharged Int64
	Result     TransactionResultResult
}

func DecodeTransactionResult(decoder *xdr.Decoder, result *TransactionResult) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeInt64(decoder, &result.FeeCharged)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeTransactionResultResult(decoder, &result.Result)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalTransactionResult(decoder *xdr.Decoder, result **TransactionResult) (int, error) {
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

	var val TransactionResult

	bytesRead, err = DecodeTransactionResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeTransactionResultFixedArray(decoder *xdr.Decoder, result []TransactionResult, size int) (int, error) {
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
		bytesRead, err = DecodeTransactionResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeTransactionResultArray(decoder *xdr.Decoder, result *[]TransactionResult, maxSize int32) (int, error) {
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

	var theResult = make([]TransactionResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeTransactionResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeTransactionResult(encoder *xdr.Encoder, value *TransactionResult) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeInt64(encoder, &value.FeeCharged)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeTransactionResultResult(encoder, &value.Result)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalTransactionResult(encoder *xdr.Encoder, value *TransactionResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeTransactionResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeTransactionResultFixedArray(encoder *xdr.Encoder, value []TransactionResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeTransactionResultArray(encoder *xdr.Encoder, value []TransactionResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeTransactionResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum PaymentResultCode
//   {
//       SUCCESS,
//       SUCCESS_MULTI,
//       UNDERFUNDED,
//       NO_DESTINATION,
//       MALFORMED,
//       NO_TRUST,
//       NOT_AUTHORIZED,
//       LINE_FULL,
//       OVERSENDMAX
//   };
//
// ===========================================================================

type PaymentResultCode int32

const (
	PaymentResultCodeSuccess       PaymentResultCode = 0
	PaymentResultCodeSuccessMulti                    = 1
	PaymentResultCodeUnderfunded                     = 2
	PaymentResultCodeNoDestination                   = 3
	PaymentResultCodeMalformed                       = 4
	PaymentResultCodeNoTrust                         = 5
	PaymentResultCodeNotAuthorized                   = 6
	PaymentResultCodeLineFull                        = 7
	PaymentResultCodeOversendmax                     = 8
)

var PaymentResultCodeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
	7: true,
	8: true,
}

func DecodePaymentResultCode(decoder *xdr.Decoder, result *PaymentResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(PaymentResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = PaymentResultCode(val)
	return bytesRead, err
}
func DecodeOptionalPaymentResultCode(decoder *xdr.Decoder, result **PaymentResultCode) (int, error) {
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

	var val PaymentResultCode

	bytesRead, err = DecodePaymentResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodePaymentResultCodeFixedArray(decoder *xdr.Decoder, result []PaymentResultCode, size int) (int, error) {
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
		bytesRead, err = DecodePaymentResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodePaymentResultCodeArray(decoder *xdr.Decoder, result *[]PaymentResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]PaymentResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodePaymentResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodePaymentResultCode(encoder *xdr.Encoder, value *PaymentResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), PaymentResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalPaymentResultCode(encoder *xdr.Encoder, value *PaymentResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodePaymentResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodePaymentResultCodeFixedArray(encoder *xdr.Encoder, value []PaymentResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodePaymentResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodePaymentResultCodeArray(encoder *xdr.Encoder, value []PaymentResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodePaymentResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct SimplePaymentResult
//   {
//       AccountID destination;
//       Currency currency;
//       int64 amount;
//   };
//
// ===========================================================================

type SimplePaymentResult struct {
	Destination AccountId
	Currency    Currency
	Amount      Int64
}

func DecodeSimplePaymentResult(decoder *xdr.Decoder, result *SimplePaymentResult) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeAccountId(decoder, &result.Destination)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCurrency(decoder, &result.Currency)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Amount)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalSimplePaymentResult(decoder *xdr.Decoder, result **SimplePaymentResult) (int, error) {
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

	var val SimplePaymentResult

	bytesRead, err = DecodeSimplePaymentResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSimplePaymentResultFixedArray(decoder *xdr.Decoder, result []SimplePaymentResult, size int) (int, error) {
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
		bytesRead, err = DecodeSimplePaymentResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSimplePaymentResultArray(decoder *xdr.Decoder, result *[]SimplePaymentResult, maxSize int32) (int, error) {
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

	var theResult = make([]SimplePaymentResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSimplePaymentResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeSimplePaymentResult(encoder *xdr.Encoder, value *SimplePaymentResult) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeAccountId(encoder, &value.Destination)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCurrency(encoder, &value.Currency)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Amount)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalSimplePaymentResult(encoder *xdr.Encoder, value *SimplePaymentResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSimplePaymentResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSimplePaymentResultFixedArray(encoder *xdr.Encoder, value []SimplePaymentResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSimplePaymentResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSimplePaymentResultArray(encoder *xdr.Encoder, value []SimplePaymentResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSimplePaymentResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct SuccessMultiResult
//   {
//       ClaimOfferAtom offers<>;
//       SimplePaymentResult last;
//   };
//
// ===========================================================================

type SuccessMultiResult struct {
	Offers []ClaimOfferAtom
	Last   SimplePaymentResult
}

func DecodeSuccessMultiResult(decoder *xdr.Decoder, result *SuccessMultiResult) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeClaimOfferAtomArray(decoder, &result.Offers, MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeSimplePaymentResult(decoder, &result.Last)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalSuccessMultiResult(decoder *xdr.Decoder, result **SuccessMultiResult) (int, error) {
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

	var val SuccessMultiResult

	bytesRead, err = DecodeSuccessMultiResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSuccessMultiResultFixedArray(decoder *xdr.Decoder, result []SuccessMultiResult, size int) (int, error) {
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
		bytesRead, err = DecodeSuccessMultiResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSuccessMultiResultArray(decoder *xdr.Decoder, result *[]SuccessMultiResult, maxSize int32) (int, error) {
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

	var theResult = make([]SuccessMultiResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSuccessMultiResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeSuccessMultiResult(encoder *xdr.Encoder, value *SuccessMultiResult) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeClaimOfferAtomArray(encoder, value.Offers[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeSimplePaymentResult(encoder, &value.Last)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalSuccessMultiResult(encoder *xdr.Encoder, value *SuccessMultiResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSuccessMultiResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSuccessMultiResultFixedArray(encoder *xdr.Encoder, value []SuccessMultiResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSuccessMultiResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSuccessMultiResultArray(encoder *xdr.Encoder, value []SuccessMultiResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSuccessMultiResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union PaymentResult switch(PaymentResultCode code)
//   {
//       case SUCCESS:
//           void;
//       case SUCCESS_MULTI:
//           SuccessMultiResult multi;
//       default:
//           void;
//   };
//
// ===========================================================================

type PaymentResult struct {
	code  PaymentResultCode
	multi *SuccessMultiResult
}

func NewPaymentResultSuccess() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeSuccess,
	}
}
func NewPaymentResultSuccessMulti(val SuccessMultiResult) PaymentResult {
	return PaymentResult{
		code:  PaymentResultCodeSuccessMulti,
		multi: &val,
	}
}
func NewPaymentResultUnderfunded() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeUnderfunded,
	}
}
func NewPaymentResultNoDestination() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeNoDestination,
	}
}
func NewPaymentResultMalformed() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeMalformed,
	}
}
func NewPaymentResultNoTrust() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeNoTrust,
	}
}
func NewPaymentResultNotAuthorized() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeNotAuthorized,
	}
}
func NewPaymentResultLineFull() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeLineFull,
	}
}
func NewPaymentResultOversendmax() PaymentResult {
	return PaymentResult{
		code: PaymentResultCodeOversendmax,
	}
}
func (u *PaymentResult) Code() PaymentResultCode {
	return u.code
}
func (u *PaymentResult) Multi() SuccessMultiResult {
	return *u.multi
}
func DecodePaymentResult(decoder *xdr.Decoder, result *PaymentResult) (int, error) {
	var (
		discriminant PaymentResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodePaymentResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == PaymentResultCodeSuccessMulti {
		var successMulti SuccessMultiResult
		bytesRead, err = DecodeSuccessMultiResult(decoder, &successMulti)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewPaymentResultSuccessMulti(successMulti)
	}

	return totalRead, nil
}
func DecodeOptionalPaymentResult(decoder *xdr.Decoder, result **PaymentResult) (int, error) {
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

	var val PaymentResult

	bytesRead, err = DecodePaymentResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodePaymentResultFixedArray(decoder *xdr.Decoder, result []PaymentResult, size int) (int, error) {
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
		bytesRead, err = DecodePaymentResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodePaymentResultArray(decoder *xdr.Decoder, result *[]PaymentResult, maxSize int32) (int, error) {
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

	var theResult = make([]PaymentResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodePaymentResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodePaymentResult(encoder *xdr.Encoder, value *PaymentResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodePaymentResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.code == PaymentResultCodeSuccessMulti {
		bytesWritten, err = EncodeSuccessMultiResult(encoder, value.multi)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalPaymentResult(encoder *xdr.Encoder, value *PaymentResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodePaymentResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodePaymentResultFixedArray(encoder *xdr.Encoder, value []PaymentResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodePaymentResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodePaymentResultArray(encoder *xdr.Encoder, value []PaymentResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodePaymentResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum CreateOfferResultCode
//   {
//       SUCCESS,
//       NO_TRUST,
//       NOT_AUTHORIZED,
//       MALFORMED,
//       UNDERFUNDED,
//       CROSS_SELF,
//       NOT_FOUND
//   };
//
// ===========================================================================

type CreateOfferResultCode int32

const (
	CreateOfferResultCodeSuccess       CreateOfferResultCode = 0
	CreateOfferResultCodeNoTrust                             = 1
	CreateOfferResultCodeNotAuthorized                       = 2
	CreateOfferResultCodeMalformed                           = 3
	CreateOfferResultCodeUnderfunded                         = 4
	CreateOfferResultCodeCrossSelf                           = 5
	CreateOfferResultCodeNotFound                            = 6
)

var CreateOfferResultCodeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
}

func DecodeCreateOfferResultCode(decoder *xdr.Decoder, result *CreateOfferResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(CreateOfferResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = CreateOfferResultCode(val)
	return bytesRead, err
}
func DecodeOptionalCreateOfferResultCode(decoder *xdr.Decoder, result **CreateOfferResultCode) (int, error) {
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

	var val CreateOfferResultCode

	bytesRead, err = DecodeCreateOfferResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCreateOfferResultCodeFixedArray(decoder *xdr.Decoder, result []CreateOfferResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeCreateOfferResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCreateOfferResultCodeArray(decoder *xdr.Decoder, result *[]CreateOfferResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]CreateOfferResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCreateOfferResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCreateOfferResultCode(encoder *xdr.Encoder, value *CreateOfferResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), CreateOfferResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalCreateOfferResultCode(encoder *xdr.Encoder, value *CreateOfferResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCreateOfferResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCreateOfferResultCodeFixedArray(encoder *xdr.Encoder, value []CreateOfferResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCreateOfferResultCodeArray(encoder *xdr.Encoder, value []CreateOfferResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum CreateOfferEffect
//   {
//       CREATED,
//       UPDATED,
//       EMPTY
//   };
//
// ===========================================================================

type CreateOfferEffect int32

const (
	CreateOfferEffectCreated CreateOfferEffect = 0
	CreateOfferEffectUpdated                   = 1
	CreateOfferEffectEmpty                     = 2
)

var CreateOfferEffectMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
}

func DecodeCreateOfferEffect(decoder *xdr.Decoder, result *CreateOfferEffect) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(CreateOfferEffectMap)

	if err != nil {
		return bytesRead, err
	}
	*result = CreateOfferEffect(val)
	return bytesRead, err
}
func DecodeOptionalCreateOfferEffect(decoder *xdr.Decoder, result **CreateOfferEffect) (int, error) {
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

	var val CreateOfferEffect

	bytesRead, err = DecodeCreateOfferEffect(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCreateOfferEffectFixedArray(decoder *xdr.Decoder, result []CreateOfferEffect, size int) (int, error) {
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
		bytesRead, err = DecodeCreateOfferEffect(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCreateOfferEffectArray(decoder *xdr.Decoder, result *[]CreateOfferEffect, maxSize int32) (int, error) {
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

	var theResult = make([]CreateOfferEffect, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCreateOfferEffect(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCreateOfferEffect(encoder *xdr.Encoder, value *CreateOfferEffect) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), CreateOfferEffectMap)
	return bytesWritten, err
}
func EncodeOptionalCreateOfferEffect(encoder *xdr.Encoder, value *CreateOfferEffect) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCreateOfferEffect(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCreateOfferEffectFixedArray(encoder *xdr.Encoder, value []CreateOfferEffect, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferEffect(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCreateOfferEffectArray(encoder *xdr.Encoder, value []CreateOfferEffect, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferEffect(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union switch(CreateOfferEffect effect)
//       {
//           case CREATED:
//               OfferEntry offerCreated;
//           default:
//               void;
//       }
//
// ===========================================================================

type CreateOfferSuccessResultOffer struct {
	effect       CreateOfferEffect
	offerCreated *OfferEntry
}

func NewCreateOfferSuccessResultOfferCreated(val OfferEntry) CreateOfferSuccessResultOffer {
	return CreateOfferSuccessResultOffer{
		effect:       CreateOfferEffectCreated,
		offerCreated: &val,
	}
}
func NewCreateOfferSuccessResultOfferUpdated() CreateOfferSuccessResultOffer {
	return CreateOfferSuccessResultOffer{
		effect: CreateOfferEffectUpdated,
	}
}
func NewCreateOfferSuccessResultOfferEmpty() CreateOfferSuccessResultOffer {
	return CreateOfferSuccessResultOffer{
		effect: CreateOfferEffectEmpty,
	}
}
func (u *CreateOfferSuccessResultOffer) Effect() CreateOfferEffect {
	return u.effect
}
func (u *CreateOfferSuccessResultOffer) OfferCreated() OfferEntry {
	return *u.offerCreated
}
func DecodeCreateOfferSuccessResultOffer(decoder *xdr.Decoder, result *CreateOfferSuccessResultOffer) (int, error) {
	var (
		discriminant CreateOfferEffect
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeCreateOfferEffect(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == CreateOfferEffectCreated {
		var created OfferEntry
		bytesRead, err = DecodeOfferEntry(decoder, &created)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewCreateOfferSuccessResultOfferCreated(created)
	}

	return totalRead, nil
}
func DecodeOptionalCreateOfferSuccessResultOffer(decoder *xdr.Decoder, result **CreateOfferSuccessResultOffer) (int, error) {
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

	var val CreateOfferSuccessResultOffer

	bytesRead, err = DecodeCreateOfferSuccessResultOffer(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCreateOfferSuccessResultOfferFixedArray(decoder *xdr.Decoder, result []CreateOfferSuccessResultOffer, size int) (int, error) {
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
		bytesRead, err = DecodeCreateOfferSuccessResultOffer(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCreateOfferSuccessResultOfferArray(decoder *xdr.Decoder, result *[]CreateOfferSuccessResultOffer, maxSize int32) (int, error) {
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

	var theResult = make([]CreateOfferSuccessResultOffer, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCreateOfferSuccessResultOffer(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCreateOfferSuccessResultOffer(encoder *xdr.Encoder, value *CreateOfferSuccessResultOffer) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeCreateOfferEffect(encoder, &value.effect)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.effect == CreateOfferEffectCreated {
		bytesWritten, err = EncodeOfferEntry(encoder, value.offerCreated)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalCreateOfferSuccessResultOffer(encoder *xdr.Encoder, value *CreateOfferSuccessResultOffer) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCreateOfferSuccessResultOffer(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCreateOfferSuccessResultOfferFixedArray(encoder *xdr.Encoder, value []CreateOfferSuccessResultOffer, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferSuccessResultOffer(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCreateOfferSuccessResultOfferArray(encoder *xdr.Encoder, value []CreateOfferSuccessResultOffer, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferSuccessResultOffer(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct CreateOfferSuccessResult
//   {
//       ClaimOfferAtom offersClaimed<>;
//
//       union switch(CreateOfferEffect effect)
//       {
//           case CREATED:
//               OfferEntry offerCreated;
//           default:
//               void;
//       } offer;
//   };
//
// ===========================================================================

type CreateOfferSuccessResult struct {
	OffersClaimed []ClaimOfferAtom
	Offer         CreateOfferSuccessResultOffer
}

func DecodeCreateOfferSuccessResult(decoder *xdr.Decoder, result *CreateOfferSuccessResult) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeClaimOfferAtomArray(decoder, &result.OffersClaimed, MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeCreateOfferSuccessResultOffer(decoder, &result.Offer)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalCreateOfferSuccessResult(decoder *xdr.Decoder, result **CreateOfferSuccessResult) (int, error) {
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

	var val CreateOfferSuccessResult

	bytesRead, err = DecodeCreateOfferSuccessResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCreateOfferSuccessResultFixedArray(decoder *xdr.Decoder, result []CreateOfferSuccessResult, size int) (int, error) {
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
		bytesRead, err = DecodeCreateOfferSuccessResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCreateOfferSuccessResultArray(decoder *xdr.Decoder, result *[]CreateOfferSuccessResult, maxSize int32) (int, error) {
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

	var theResult = make([]CreateOfferSuccessResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCreateOfferSuccessResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCreateOfferSuccessResult(encoder *xdr.Encoder, value *CreateOfferSuccessResult) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeClaimOfferAtomArray(encoder, value.OffersClaimed[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeCreateOfferSuccessResultOffer(encoder, &value.Offer)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalCreateOfferSuccessResult(encoder *xdr.Encoder, value *CreateOfferSuccessResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCreateOfferSuccessResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCreateOfferSuccessResultFixedArray(encoder *xdr.Encoder, value []CreateOfferSuccessResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferSuccessResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCreateOfferSuccessResultArray(encoder *xdr.Encoder, value []CreateOfferSuccessResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferSuccessResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union CreateOfferResult switch(CreateOfferResultCode code)
//   {
//       case SUCCESS:
//           CreateOfferSuccessResult success;
//       default:
//           void;
//   };
//
// ===========================================================================

type CreateOfferResult struct {
	code    CreateOfferResultCode
	success *CreateOfferSuccessResult
}

func NewCreateOfferResultSuccess(val CreateOfferSuccessResult) CreateOfferResult {
	return CreateOfferResult{
		code:    CreateOfferResultCodeSuccess,
		success: &val,
	}
}
func NewCreateOfferResultNoTrust() CreateOfferResult {
	return CreateOfferResult{
		code: CreateOfferResultCodeNoTrust,
	}
}
func NewCreateOfferResultNotAuthorized() CreateOfferResult {
	return CreateOfferResult{
		code: CreateOfferResultCodeNotAuthorized,
	}
}
func NewCreateOfferResultMalformed() CreateOfferResult {
	return CreateOfferResult{
		code: CreateOfferResultCodeMalformed,
	}
}
func NewCreateOfferResultUnderfunded() CreateOfferResult {
	return CreateOfferResult{
		code: CreateOfferResultCodeUnderfunded,
	}
}
func NewCreateOfferResultCrossSelf() CreateOfferResult {
	return CreateOfferResult{
		code: CreateOfferResultCodeCrossSelf,
	}
}
func NewCreateOfferResultNotFound() CreateOfferResult {
	return CreateOfferResult{
		code: CreateOfferResultCodeNotFound,
	}
}
func (u *CreateOfferResult) Code() CreateOfferResultCode {
	return u.code
}
func (u *CreateOfferResult) Success() CreateOfferSuccessResult {
	return *u.success
}
func DecodeCreateOfferResult(decoder *xdr.Decoder, result *CreateOfferResult) (int, error) {
	var (
		discriminant CreateOfferResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeCreateOfferResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == CreateOfferResultCodeSuccess {
		var success CreateOfferSuccessResult
		bytesRead, err = DecodeCreateOfferSuccessResult(decoder, &success)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewCreateOfferResultSuccess(success)
	}

	return totalRead, nil
}
func DecodeOptionalCreateOfferResult(decoder *xdr.Decoder, result **CreateOfferResult) (int, error) {
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

	var val CreateOfferResult

	bytesRead, err = DecodeCreateOfferResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCreateOfferResultFixedArray(decoder *xdr.Decoder, result []CreateOfferResult, size int) (int, error) {
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
		bytesRead, err = DecodeCreateOfferResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCreateOfferResultArray(decoder *xdr.Decoder, result *[]CreateOfferResult, maxSize int32) (int, error) {
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

	var theResult = make([]CreateOfferResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCreateOfferResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCreateOfferResult(encoder *xdr.Encoder, value *CreateOfferResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeCreateOfferResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.code == CreateOfferResultCodeSuccess {
		bytesWritten, err = EncodeCreateOfferSuccessResult(encoder, value.success)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalCreateOfferResult(encoder *xdr.Encoder, value *CreateOfferResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCreateOfferResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCreateOfferResultFixedArray(encoder *xdr.Encoder, value []CreateOfferResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCreateOfferResultArray(encoder *xdr.Encoder, value []CreateOfferResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCreateOfferResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum CancelOfferResultCode
//   {
//       SUCCESS,
//       NOT_FOUND
//   };
//
// ===========================================================================

type CancelOfferResultCode int32

const (
	CancelOfferResultCodeSuccess  CancelOfferResultCode = 0
	CancelOfferResultCodeNotFound                       = 1
)

var CancelOfferResultCodeMap = map[int32]bool{
	0: true,
	1: true,
}

func DecodeCancelOfferResultCode(decoder *xdr.Decoder, result *CancelOfferResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(CancelOfferResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = CancelOfferResultCode(val)
	return bytesRead, err
}
func DecodeOptionalCancelOfferResultCode(decoder *xdr.Decoder, result **CancelOfferResultCode) (int, error) {
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

	var val CancelOfferResultCode

	bytesRead, err = DecodeCancelOfferResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCancelOfferResultCodeFixedArray(decoder *xdr.Decoder, result []CancelOfferResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeCancelOfferResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCancelOfferResultCodeArray(decoder *xdr.Decoder, result *[]CancelOfferResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]CancelOfferResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCancelOfferResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCancelOfferResultCode(encoder *xdr.Encoder, value *CancelOfferResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), CancelOfferResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalCancelOfferResultCode(encoder *xdr.Encoder, value *CancelOfferResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCancelOfferResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCancelOfferResultCodeFixedArray(encoder *xdr.Encoder, value []CancelOfferResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCancelOfferResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCancelOfferResultCodeArray(encoder *xdr.Encoder, value []CancelOfferResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCancelOfferResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union CancelOfferResult switch(CancelOfferResultCode code)
//   {
//       case SUCCESS:
//           void;
//       default:
//           void;
//   };
//
// ===========================================================================

type CancelOfferResult struct {
	code CancelOfferResultCode
}

func NewCancelOfferResultSuccess() CancelOfferResult {
	return CancelOfferResult{
		code: CancelOfferResultCodeSuccess,
	}
}
func NewCancelOfferResultNotFound() CancelOfferResult {
	return CancelOfferResult{
		code: CancelOfferResultCodeNotFound,
	}
}
func (u *CancelOfferResult) Code() CancelOfferResultCode {
	return u.code
}
func DecodeCancelOfferResult(decoder *xdr.Decoder, result *CancelOfferResult) (int, error) {
	var (
		discriminant CancelOfferResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeCancelOfferResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalCancelOfferResult(decoder *xdr.Decoder, result **CancelOfferResult) (int, error) {
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

	var val CancelOfferResult

	bytesRead, err = DecodeCancelOfferResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeCancelOfferResultFixedArray(decoder *xdr.Decoder, result []CancelOfferResult, size int) (int, error) {
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
		bytesRead, err = DecodeCancelOfferResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeCancelOfferResultArray(decoder *xdr.Decoder, result *[]CancelOfferResult, maxSize int32) (int, error) {
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

	var theResult = make([]CancelOfferResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeCancelOfferResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeCancelOfferResult(encoder *xdr.Encoder, value *CancelOfferResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeCancelOfferResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalCancelOfferResult(encoder *xdr.Encoder, value *CancelOfferResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeCancelOfferResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeCancelOfferResultFixedArray(encoder *xdr.Encoder, value []CancelOfferResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeCancelOfferResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeCancelOfferResultArray(encoder *xdr.Encoder, value []CancelOfferResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeCancelOfferResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum SetOptionsResultCode
//   {
//       SUCCESS,
//       RATE_FIXED,
//       RATE_TOO_HIGH,
//       BELOW_MIN_BALANCE,
//       MALFORMED
//   };
//
// ===========================================================================

type SetOptionsResultCode int32

const (
	SetOptionsResultCodeSuccess         SetOptionsResultCode = 0
	SetOptionsResultCodeRateFixed                            = 1
	SetOptionsResultCodeRateTooHigh                          = 2
	SetOptionsResultCodeBelowMinBalance                      = 3
	SetOptionsResultCodeMalformed                            = 4
)

var SetOptionsResultCodeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
	4: true,
}

func DecodeSetOptionsResultCode(decoder *xdr.Decoder, result *SetOptionsResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(SetOptionsResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = SetOptionsResultCode(val)
	return bytesRead, err
}
func DecodeOptionalSetOptionsResultCode(decoder *xdr.Decoder, result **SetOptionsResultCode) (int, error) {
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

	var val SetOptionsResultCode

	bytesRead, err = DecodeSetOptionsResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSetOptionsResultCodeFixedArray(decoder *xdr.Decoder, result []SetOptionsResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeSetOptionsResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSetOptionsResultCodeArray(decoder *xdr.Decoder, result *[]SetOptionsResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]SetOptionsResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSetOptionsResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeSetOptionsResultCode(encoder *xdr.Encoder, value *SetOptionsResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), SetOptionsResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalSetOptionsResultCode(encoder *xdr.Encoder, value *SetOptionsResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSetOptionsResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSetOptionsResultCodeFixedArray(encoder *xdr.Encoder, value []SetOptionsResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSetOptionsResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSetOptionsResultCodeArray(encoder *xdr.Encoder, value []SetOptionsResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSetOptionsResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union SetOptionsResult switch(SetOptionsResultCode code)
//   {
//       case SUCCESS:
//           void;
//       default:
//           void;
//   };
//
// ===========================================================================

type SetOptionsResult struct {
	code SetOptionsResultCode
}

func NewSetOptionsResultSuccess() SetOptionsResult {
	return SetOptionsResult{
		code: SetOptionsResultCodeSuccess,
	}
}
func NewSetOptionsResultRateFixed() SetOptionsResult {
	return SetOptionsResult{
		code: SetOptionsResultCodeRateFixed,
	}
}
func NewSetOptionsResultRateTooHigh() SetOptionsResult {
	return SetOptionsResult{
		code: SetOptionsResultCodeRateTooHigh,
	}
}
func NewSetOptionsResultBelowMinBalance() SetOptionsResult {
	return SetOptionsResult{
		code: SetOptionsResultCodeBelowMinBalance,
	}
}
func NewSetOptionsResultMalformed() SetOptionsResult {
	return SetOptionsResult{
		code: SetOptionsResultCodeMalformed,
	}
}
func (u *SetOptionsResult) Code() SetOptionsResultCode {
	return u.code
}
func DecodeSetOptionsResult(decoder *xdr.Decoder, result *SetOptionsResult) (int, error) {
	var (
		discriminant SetOptionsResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeSetOptionsResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalSetOptionsResult(decoder *xdr.Decoder, result **SetOptionsResult) (int, error) {
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

	var val SetOptionsResult

	bytesRead, err = DecodeSetOptionsResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeSetOptionsResultFixedArray(decoder *xdr.Decoder, result []SetOptionsResult, size int) (int, error) {
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
		bytesRead, err = DecodeSetOptionsResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeSetOptionsResultArray(decoder *xdr.Decoder, result *[]SetOptionsResult, maxSize int32) (int, error) {
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

	var theResult = make([]SetOptionsResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeSetOptionsResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeSetOptionsResult(encoder *xdr.Encoder, value *SetOptionsResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeSetOptionsResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalSetOptionsResult(encoder *xdr.Encoder, value *SetOptionsResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeSetOptionsResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeSetOptionsResultFixedArray(encoder *xdr.Encoder, value []SetOptionsResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeSetOptionsResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeSetOptionsResultArray(encoder *xdr.Encoder, value []SetOptionsResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeSetOptionsResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum ChangeTrustResultCode
//   {
//       SUCCESS,
//       NO_ACCOUNT
//   };
//
// ===========================================================================

type ChangeTrustResultCode int32

const (
	ChangeTrustResultCodeSuccess   ChangeTrustResultCode = 0
	ChangeTrustResultCodeNoAccount                       = 1
)

var ChangeTrustResultCodeMap = map[int32]bool{
	0: true,
	1: true,
}

func DecodeChangeTrustResultCode(decoder *xdr.Decoder, result *ChangeTrustResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(ChangeTrustResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = ChangeTrustResultCode(val)
	return bytesRead, err
}
func DecodeOptionalChangeTrustResultCode(decoder *xdr.Decoder, result **ChangeTrustResultCode) (int, error) {
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

	var val ChangeTrustResultCode

	bytesRead, err = DecodeChangeTrustResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeChangeTrustResultCodeFixedArray(decoder *xdr.Decoder, result []ChangeTrustResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeChangeTrustResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeChangeTrustResultCodeArray(decoder *xdr.Decoder, result *[]ChangeTrustResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]ChangeTrustResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeChangeTrustResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeChangeTrustResultCode(encoder *xdr.Encoder, value *ChangeTrustResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), ChangeTrustResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalChangeTrustResultCode(encoder *xdr.Encoder, value *ChangeTrustResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeChangeTrustResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeChangeTrustResultCodeFixedArray(encoder *xdr.Encoder, value []ChangeTrustResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeChangeTrustResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeChangeTrustResultCodeArray(encoder *xdr.Encoder, value []ChangeTrustResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeChangeTrustResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union ChangeTrustResult switch(ChangeTrustResultCode code)
//   {
//       case SUCCESS:
//           void;
//       default:
//           void;
//   };
//
// ===========================================================================

type ChangeTrustResult struct {
	code ChangeTrustResultCode
}

func NewChangeTrustResultSuccess() ChangeTrustResult {
	return ChangeTrustResult{
		code: ChangeTrustResultCodeSuccess,
	}
}
func NewChangeTrustResultNoAccount() ChangeTrustResult {
	return ChangeTrustResult{
		code: ChangeTrustResultCodeNoAccount,
	}
}
func (u *ChangeTrustResult) Code() ChangeTrustResultCode {
	return u.code
}
func DecodeChangeTrustResult(decoder *xdr.Decoder, result *ChangeTrustResult) (int, error) {
	var (
		discriminant ChangeTrustResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeChangeTrustResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalChangeTrustResult(decoder *xdr.Decoder, result **ChangeTrustResult) (int, error) {
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

	var val ChangeTrustResult

	bytesRead, err = DecodeChangeTrustResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeChangeTrustResultFixedArray(decoder *xdr.Decoder, result []ChangeTrustResult, size int) (int, error) {
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
		bytesRead, err = DecodeChangeTrustResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeChangeTrustResultArray(decoder *xdr.Decoder, result *[]ChangeTrustResult, maxSize int32) (int, error) {
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

	var theResult = make([]ChangeTrustResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeChangeTrustResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeChangeTrustResult(encoder *xdr.Encoder, value *ChangeTrustResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeChangeTrustResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalChangeTrustResult(encoder *xdr.Encoder, value *ChangeTrustResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeChangeTrustResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeChangeTrustResultFixedArray(encoder *xdr.Encoder, value []ChangeTrustResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeChangeTrustResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeChangeTrustResultArray(encoder *xdr.Encoder, value []ChangeTrustResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeChangeTrustResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum AllowTrustResultCode
//   {
//       SUCCESS,
//       MALFORMED,
//       NO_TRUST_LINE
//   };
//
// ===========================================================================

type AllowTrustResultCode int32

const (
	AllowTrustResultCodeSuccess     AllowTrustResultCode = 0
	AllowTrustResultCodeMalformed                        = 1
	AllowTrustResultCodeNoTrustLine                      = 2
)

var AllowTrustResultCodeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
}

func DecodeAllowTrustResultCode(decoder *xdr.Decoder, result *AllowTrustResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(AllowTrustResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = AllowTrustResultCode(val)
	return bytesRead, err
}
func DecodeOptionalAllowTrustResultCode(decoder *xdr.Decoder, result **AllowTrustResultCode) (int, error) {
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

	var val AllowTrustResultCode

	bytesRead, err = DecodeAllowTrustResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAllowTrustResultCodeFixedArray(decoder *xdr.Decoder, result []AllowTrustResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeAllowTrustResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAllowTrustResultCodeArray(decoder *xdr.Decoder, result *[]AllowTrustResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]AllowTrustResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAllowTrustResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAllowTrustResultCode(encoder *xdr.Encoder, value *AllowTrustResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), AllowTrustResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalAllowTrustResultCode(encoder *xdr.Encoder, value *AllowTrustResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAllowTrustResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAllowTrustResultCodeFixedArray(encoder *xdr.Encoder, value []AllowTrustResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAllowTrustResultCodeArray(encoder *xdr.Encoder, value []AllowTrustResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union AllowTrustResult switch(AllowTrustResultCode code)
//   {
//       case SUCCESS:
//           void;
//       default:
//           void;
//   };
//
// ===========================================================================

type AllowTrustResult struct {
	code AllowTrustResultCode
}

func NewAllowTrustResultSuccess() AllowTrustResult {
	return AllowTrustResult{
		code: AllowTrustResultCodeSuccess,
	}
}
func NewAllowTrustResultMalformed() AllowTrustResult {
	return AllowTrustResult{
		code: AllowTrustResultCodeMalformed,
	}
}
func NewAllowTrustResultNoTrustLine() AllowTrustResult {
	return AllowTrustResult{
		code: AllowTrustResultCodeNoTrustLine,
	}
}
func (u *AllowTrustResult) Code() AllowTrustResultCode {
	return u.code
}
func DecodeAllowTrustResult(decoder *xdr.Decoder, result *AllowTrustResult) (int, error) {
	var (
		discriminant AllowTrustResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeAllowTrustResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalAllowTrustResult(decoder *xdr.Decoder, result **AllowTrustResult) (int, error) {
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

	var val AllowTrustResult

	bytesRead, err = DecodeAllowTrustResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAllowTrustResultFixedArray(decoder *xdr.Decoder, result []AllowTrustResult, size int) (int, error) {
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
		bytesRead, err = DecodeAllowTrustResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAllowTrustResultArray(decoder *xdr.Decoder, result *[]AllowTrustResult, maxSize int32) (int, error) {
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

	var theResult = make([]AllowTrustResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAllowTrustResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAllowTrustResult(encoder *xdr.Encoder, value *AllowTrustResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeAllowTrustResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalAllowTrustResult(encoder *xdr.Encoder, value *AllowTrustResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAllowTrustResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAllowTrustResultFixedArray(encoder *xdr.Encoder, value []AllowTrustResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAllowTrustResultArray(encoder *xdr.Encoder, value []AllowTrustResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAllowTrustResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum AccountMergeResultCode
//   {
//       SUCCESS,
//       MALFORMED,
//       NO_ACCOUNT,
//       HAS_CREDIT
//   };
//
// ===========================================================================

type AccountMergeResultCode int32

const (
	AccountMergeResultCodeSuccess   AccountMergeResultCode = 0
	AccountMergeResultCodeMalformed                        = 1
	AccountMergeResultCodeNoAccount                        = 2
	AccountMergeResultCodeHasCredit                        = 3
)

var AccountMergeResultCodeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
}

func DecodeAccountMergeResultCode(decoder *xdr.Decoder, result *AccountMergeResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(AccountMergeResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = AccountMergeResultCode(val)
	return bytesRead, err
}
func DecodeOptionalAccountMergeResultCode(decoder *xdr.Decoder, result **AccountMergeResultCode) (int, error) {
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

	var val AccountMergeResultCode

	bytesRead, err = DecodeAccountMergeResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAccountMergeResultCodeFixedArray(decoder *xdr.Decoder, result []AccountMergeResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeAccountMergeResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAccountMergeResultCodeArray(decoder *xdr.Decoder, result *[]AccountMergeResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]AccountMergeResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAccountMergeResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAccountMergeResultCode(encoder *xdr.Encoder, value *AccountMergeResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), AccountMergeResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalAccountMergeResultCode(encoder *xdr.Encoder, value *AccountMergeResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAccountMergeResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAccountMergeResultCodeFixedArray(encoder *xdr.Encoder, value []AccountMergeResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountMergeResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAccountMergeResultCodeArray(encoder *xdr.Encoder, value []AccountMergeResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountMergeResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union AccountMergeResult switch(AccountMergeResultCode code)
//   {
//       case SUCCESS:
//           void;
//       default:
//           void;
//   };
//
// ===========================================================================

type AccountMergeResult struct {
	code AccountMergeResultCode
}

func NewAccountMergeResultSuccess() AccountMergeResult {
	return AccountMergeResult{
		code: AccountMergeResultCodeSuccess,
	}
}
func NewAccountMergeResultMalformed() AccountMergeResult {
	return AccountMergeResult{
		code: AccountMergeResultCodeMalformed,
	}
}
func NewAccountMergeResultNoAccount() AccountMergeResult {
	return AccountMergeResult{
		code: AccountMergeResultCodeNoAccount,
	}
}
func NewAccountMergeResultHasCredit() AccountMergeResult {
	return AccountMergeResult{
		code: AccountMergeResultCodeHasCredit,
	}
}
func (u *AccountMergeResult) Code() AccountMergeResultCode {
	return u.code
}
func DecodeAccountMergeResult(decoder *xdr.Decoder, result *AccountMergeResult) (int, error) {
	var (
		discriminant AccountMergeResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeAccountMergeResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalAccountMergeResult(decoder *xdr.Decoder, result **AccountMergeResult) (int, error) {
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

	var val AccountMergeResult

	bytesRead, err = DecodeAccountMergeResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeAccountMergeResultFixedArray(decoder *xdr.Decoder, result []AccountMergeResult, size int) (int, error) {
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
		bytesRead, err = DecodeAccountMergeResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeAccountMergeResultArray(decoder *xdr.Decoder, result *[]AccountMergeResult, maxSize int32) (int, error) {
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

	var theResult = make([]AccountMergeResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeAccountMergeResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeAccountMergeResult(encoder *xdr.Encoder, value *AccountMergeResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeAccountMergeResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalAccountMergeResult(encoder *xdr.Encoder, value *AccountMergeResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeAccountMergeResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeAccountMergeResultFixedArray(encoder *xdr.Encoder, value []AccountMergeResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountMergeResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeAccountMergeResultArray(encoder *xdr.Encoder, value []AccountMergeResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeAccountMergeResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum InflationResultCode
//   {
//       SUCCESS,
//       NOT_TIME
//   };
//
// ===========================================================================

type InflationResultCode int32

const (
	InflationResultCodeSuccess InflationResultCode = 0
	InflationResultCodeNotTime                     = 1
)

var InflationResultCodeMap = map[int32]bool{
	0: true,
	1: true,
}

func DecodeInflationResultCode(decoder *xdr.Decoder, result *InflationResultCode) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(InflationResultCodeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = InflationResultCode(val)
	return bytesRead, err
}
func DecodeOptionalInflationResultCode(decoder *xdr.Decoder, result **InflationResultCode) (int, error) {
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

	var val InflationResultCode

	bytesRead, err = DecodeInflationResultCode(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeInflationResultCodeFixedArray(decoder *xdr.Decoder, result []InflationResultCode, size int) (int, error) {
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
		bytesRead, err = DecodeInflationResultCode(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeInflationResultCodeArray(decoder *xdr.Decoder, result *[]InflationResultCode, maxSize int32) (int, error) {
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

	var theResult = make([]InflationResultCode, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeInflationResultCode(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeInflationResultCode(encoder *xdr.Encoder, value *InflationResultCode) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), InflationResultCodeMap)
	return bytesWritten, err
}
func EncodeOptionalInflationResultCode(encoder *xdr.Encoder, value *InflationResultCode) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeInflationResultCode(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeInflationResultCodeFixedArray(encoder *xdr.Encoder, value []InflationResultCode, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeInflationResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeInflationResultCodeArray(encoder *xdr.Encoder, value []InflationResultCode, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeInflationResultCode(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct inflationPayout // or use PaymentResultAtom to limit types?
//   {
//       AccountID destination;
//       int64 amount;
//   };
//
// ===========================================================================

type InflationPayout struct {
	Destination AccountId
	Amount      Int64
}

func DecodeInflationPayout(decoder *xdr.Decoder, result *InflationPayout) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeAccountId(decoder, &result.Destination)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt64(decoder, &result.Amount)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalInflationPayout(decoder *xdr.Decoder, result **InflationPayout) (int, error) {
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

	var val InflationPayout

	bytesRead, err = DecodeInflationPayout(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeInflationPayoutFixedArray(decoder *xdr.Decoder, result []InflationPayout, size int) (int, error) {
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
		bytesRead, err = DecodeInflationPayout(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeInflationPayoutArray(decoder *xdr.Decoder, result *[]InflationPayout, maxSize int32) (int, error) {
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

	var theResult = make([]InflationPayout, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeInflationPayout(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeInflationPayout(encoder *xdr.Encoder, value *InflationPayout) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeAccountId(encoder, &value.Destination)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt64(encoder, &value.Amount)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalInflationPayout(encoder *xdr.Encoder, value *InflationPayout) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeInflationPayout(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeInflationPayoutFixedArray(encoder *xdr.Encoder, value []InflationPayout, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeInflationPayout(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeInflationPayoutArray(encoder *xdr.Encoder, value []InflationPayout, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeInflationPayout(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union InflationResult switch(InflationResultCode code)
//   {
//       case SUCCESS:
//           inflationPayout payouts<>;
//       default:
//           void;
//   };
//
// ===========================================================================

type InflationResult struct {
	code    InflationResultCode
	payouts *[]InflationPayout
}

func NewInflationResultSuccess(val []InflationPayout) InflationResult {
	return InflationResult{
		code:    InflationResultCodeSuccess,
		payouts: &val,
	}
}
func NewInflationResultNotTime() InflationResult {
	return InflationResult{
		code: InflationResultCodeNotTime,
	}
}
func (u *InflationResult) Code() InflationResultCode {
	return u.code
}
func (u *InflationResult) Payouts() []InflationPayout {
	return *u.payouts
}
func DecodeInflationResult(decoder *xdr.Decoder, result *InflationResult) (int, error) {
	var (
		discriminant InflationResultCode
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeInflationResultCode(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == InflationResultCodeSuccess {
		var success []InflationPayout
		bytesRead, err = DecodeInflationPayoutArray(decoder, &success, MaxXdrElements)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewInflationResultSuccess(success)
	}

	return totalRead, nil
}
func DecodeOptionalInflationResult(decoder *xdr.Decoder, result **InflationResult) (int, error) {
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

	var val InflationResult

	bytesRead, err = DecodeInflationResult(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeInflationResultFixedArray(decoder *xdr.Decoder, result []InflationResult, size int) (int, error) {
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
		bytesRead, err = DecodeInflationResult(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeInflationResultArray(decoder *xdr.Decoder, result *[]InflationResult, maxSize int32) (int, error) {
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

	var theResult = make([]InflationResult, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeInflationResult(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeInflationResult(encoder *xdr.Encoder, value *InflationResult) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeInflationResultCode(encoder, &value.code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.code == InflationResultCodeSuccess {
		bytesWritten, err = EncodeInflationPayoutArray(encoder, value.payouts, MaxXdrElements)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalInflationResult(encoder *xdr.Encoder, value *InflationResult) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeInflationResult(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeInflationResultFixedArray(encoder *xdr.Encoder, value []InflationResult, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeInflationResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeInflationResultArray(encoder *xdr.Encoder, value []InflationResult, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeInflationResult(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct StellarBallotValue
//   {
//       Hash txSetHash;
//       uint64 closeTime;
//       uint32 baseFee;
//   };
//
// ===========================================================================

type StellarBallotValue struct {
	TxSetHash Hash
	CloseTime Uint64
	BaseFee   Uint32
}

func DecodeStellarBallotValue(decoder *xdr.Decoder, result *StellarBallotValue) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeHash(decoder, &result.TxSetHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint64(decoder, &result.CloseTime)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.BaseFee)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalStellarBallotValue(decoder *xdr.Decoder, result **StellarBallotValue) (int, error) {
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

	var val StellarBallotValue

	bytesRead, err = DecodeStellarBallotValue(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeStellarBallotValueFixedArray(decoder *xdr.Decoder, result []StellarBallotValue, size int) (int, error) {
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
		bytesRead, err = DecodeStellarBallotValue(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeStellarBallotValueArray(decoder *xdr.Decoder, result *[]StellarBallotValue, maxSize int32) (int, error) {
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

	var theResult = make([]StellarBallotValue, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeStellarBallotValue(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeStellarBallotValue(encoder *xdr.Encoder, value *StellarBallotValue) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeHash(encoder, &value.TxSetHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint64(encoder, &value.CloseTime)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.BaseFee)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalStellarBallotValue(encoder *xdr.Encoder, value *StellarBallotValue) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeStellarBallotValue(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeStellarBallotValueFixedArray(encoder *xdr.Encoder, value []StellarBallotValue, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeStellarBallotValue(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeStellarBallotValueArray(encoder *xdr.Encoder, value []StellarBallotValue, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeStellarBallotValue(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct StellarBallot
//   {
//       uint256 nodeID;
//       Signature signature;
//       StellarBallotValue value;
//   };
//
// ===========================================================================

type StellarBallot struct {
	NodeId    Uint256
	Signature Signature
	Value     StellarBallotValue
}

func DecodeStellarBallot(decoder *xdr.Decoder, result *StellarBallot) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.NodeId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeSignature(decoder, &result.Signature)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeStellarBallotValue(decoder, &result.Value)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalStellarBallot(decoder *xdr.Decoder, result **StellarBallot) (int, error) {
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

	var val StellarBallot

	bytesRead, err = DecodeStellarBallot(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeStellarBallotFixedArray(decoder *xdr.Decoder, result []StellarBallot, size int) (int, error) {
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
		bytesRead, err = DecodeStellarBallot(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeStellarBallotArray(decoder *xdr.Decoder, result *[]StellarBallot, maxSize int32) (int, error) {
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

	var theResult = make([]StellarBallot, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeStellarBallot(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeStellarBallot(encoder *xdr.Encoder, value *StellarBallot) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.NodeId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeSignature(encoder, &value.Signature)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeStellarBallotValue(encoder, &value.Value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalStellarBallot(encoder *xdr.Encoder, value *StellarBallot) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeStellarBallot(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeStellarBallotFixedArray(encoder *xdr.Encoder, value []StellarBallot, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeStellarBallot(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeStellarBallotArray(encoder *xdr.Encoder, value []StellarBallot, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeStellarBallot(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct Error
//   {
//       int code;
//       string msg<100>;
//   };
//
// ===========================================================================

type Error struct {
	Code int32
	Msg  string
}

func DecodeError(decoder *xdr.Decoder, result *Error) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeInt(decoder, &result.Code)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeString(decoder, &result.Msg, 100)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalError(decoder *xdr.Decoder, result **Error) (int, error) {
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

	var val Error

	bytesRead, err = DecodeError(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeErrorFixedArray(decoder *xdr.Decoder, result []Error, size int) (int, error) {
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
		bytesRead, err = DecodeError(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeErrorArray(decoder *xdr.Decoder, result *[]Error, maxSize int32) (int, error) {
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

	var theResult = make([]Error, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeError(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeError(encoder *xdr.Encoder, value *Error) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeInt(encoder, &value.Code)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeString(encoder, &value.Msg, 100)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalError(encoder *xdr.Encoder, value *Error) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeError(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeErrorFixedArray(encoder *xdr.Encoder, value []Error, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeError(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeErrorArray(encoder *xdr.Encoder, value []Error, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeError(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct Hello
//   {
//       int protocolVersion;
//       string versionStr<100>;
//       int listeningPort;
//       opaque peerID[32];
//   };
//
// ===========================================================================

type Hello struct {
	ProtocolVersion int32
	VersionStr      string
	ListeningPort   int32
	PeerId          [32]byte
}

func DecodeHello(decoder *xdr.Decoder, result *Hello) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeInt(decoder, &result.ProtocolVersion)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeString(decoder, &result.VersionStr, 100)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeInt(decoder, &result.ListeningPort)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeFixedOpaque(decoder, result.PeerId[:], 32)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalHello(decoder *xdr.Decoder, result **Hello) (int, error) {
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

	var val Hello

	bytesRead, err = DecodeHello(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeHelloFixedArray(decoder *xdr.Decoder, result []Hello, size int) (int, error) {
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
		bytesRead, err = DecodeHello(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeHelloArray(decoder *xdr.Decoder, result *[]Hello, maxSize int32) (int, error) {
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

	var theResult = make([]Hello, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeHello(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeHello(encoder *xdr.Encoder, value *Hello) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeInt(encoder, &value.ProtocolVersion)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeString(encoder, &value.VersionStr, 100)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeInt(encoder, &value.ListeningPort)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeFixedOpaque(encoder, value.PeerId[:][:], 32)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalHello(encoder *xdr.Encoder, value *Hello) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeHello(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeHelloFixedArray(encoder *xdr.Encoder, value []Hello, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeHello(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeHelloArray(encoder *xdr.Encoder, value []Hello, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeHello(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct PeerAddress
//   {
//       opaque ip[4];
//       uint32 port;
//       uint32 numFailures;
//   };
//
// ===========================================================================

type PeerAddress struct {
	Ip          [4]byte
	Port        Uint32
	NumFailures Uint32
}

func DecodePeerAddress(decoder *xdr.Decoder, result *PeerAddress) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeFixedOpaque(decoder, result.Ip[:], 4)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.Port)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint32(decoder, &result.NumFailures)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalPeerAddress(decoder *xdr.Decoder, result **PeerAddress) (int, error) {
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

	var val PeerAddress

	bytesRead, err = DecodePeerAddress(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodePeerAddressFixedArray(decoder *xdr.Decoder, result []PeerAddress, size int) (int, error) {
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
		bytesRead, err = DecodePeerAddress(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodePeerAddressArray(decoder *xdr.Decoder, result *[]PeerAddress, maxSize int32) (int, error) {
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

	var theResult = make([]PeerAddress, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodePeerAddress(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodePeerAddress(encoder *xdr.Encoder, value *PeerAddress) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeFixedOpaque(encoder, value.Ip[:][:], 4)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.Port)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint32(encoder, &value.NumFailures)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalPeerAddress(encoder *xdr.Encoder, value *PeerAddress) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodePeerAddress(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodePeerAddressFixedArray(encoder *xdr.Encoder, value []PeerAddress, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodePeerAddress(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodePeerAddressArray(encoder *xdr.Encoder, value []PeerAddress, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodePeerAddress(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum MessageType
//   {
//       ERROR_MSG=0,
//       HELLO=1,
//       DONT_HAVE=2,
//
//       GET_PEERS=3,   // gets a list of peers this guy knows about
//       PEERS=4,
//
//       GET_TX_SET=5,  // gets a particular txset by hash
//       TX_SET=6,
//
//       GET_VALIDATIONS=7, // gets validations for a given ledger hash
//       VALIDATIONS=8,
//
//       TRANSACTION=9, //pass on a tx you have heard about
//       JSON_TRANSACTION=10,
//
//       // SCP
//       GET_SCP_QUORUMSET=11,
//       SCP_QUORUMSET=12,
//       SCP_MESSAGE=13
//   };
//
// ===========================================================================

type MessageType int32

const (
	MessageTypeErrorMsg        MessageType = 0
	MessageTypeHello                       = 1
	MessageTypeDontHave                    = 2
	MessageTypeGetPeer                     = 3
	MessageTypePeer                        = 4
	MessageTypeGetTxSet                    = 5
	MessageTypeTxSet                       = 6
	MessageTypeGetValidation               = 7
	MessageTypeValidation                  = 8
	MessageTypeTransaction                 = 9
	MessageTypeJsonTransaction             = 10
	MessageTypeGetScpQuorumset             = 11
	MessageTypeScpQuorumset                = 12
	MessageTypeScpMessage                  = 13
)

var MessageTypeMap = map[int32]bool{
	0:  true,
	1:  true,
	2:  true,
	3:  true,
	4:  true,
	5:  true,
	6:  true,
	7:  true,
	8:  true,
	9:  true,
	10: true,
	11: true,
	12: true,
	13: true,
}

func DecodeMessageType(decoder *xdr.Decoder, result *MessageType) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(MessageTypeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = MessageType(val)
	return bytesRead, err
}
func DecodeOptionalMessageType(decoder *xdr.Decoder, result **MessageType) (int, error) {
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

	var val MessageType

	bytesRead, err = DecodeMessageType(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeMessageTypeFixedArray(decoder *xdr.Decoder, result []MessageType, size int) (int, error) {
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
		bytesRead, err = DecodeMessageType(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeMessageTypeArray(decoder *xdr.Decoder, result *[]MessageType, maxSize int32) (int, error) {
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

	var theResult = make([]MessageType, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeMessageType(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeMessageType(encoder *xdr.Encoder, value *MessageType) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), MessageTypeMap)
	return bytesWritten, err
}
func EncodeOptionalMessageType(encoder *xdr.Encoder, value *MessageType) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeMessageType(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeMessageTypeFixedArray(encoder *xdr.Encoder, value []MessageType, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeMessageType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeMessageTypeArray(encoder *xdr.Encoder, value []MessageType, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeMessageType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct DontHave
//   {
//       MessageType type;
//       uint256 reqHash;
//   };
//
// ===========================================================================

type DontHave struct {
	Type    MessageType
	ReqHash Uint256
}

func DecodeDontHave(decoder *xdr.Decoder, result *DontHave) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeMessageType(decoder, &result.Type)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeUint256(decoder, &result.ReqHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalDontHave(decoder *xdr.Decoder, result **DontHave) (int, error) {
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

	var val DontHave

	bytesRead, err = DecodeDontHave(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeDontHaveFixedArray(decoder *xdr.Decoder, result []DontHave, size int) (int, error) {
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
		bytesRead, err = DecodeDontHave(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeDontHaveArray(decoder *xdr.Decoder, result *[]DontHave, maxSize int32) (int, error) {
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

	var theResult = make([]DontHave, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeDontHave(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeDontHave(encoder *xdr.Encoder, value *DontHave) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeMessageType(encoder, &value.Type)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeUint256(encoder, &value.ReqHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalDontHave(encoder *xdr.Encoder, value *DontHave) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeDontHave(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeDontHaveFixedArray(encoder *xdr.Encoder, value []DontHave, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeDontHave(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeDontHaveArray(encoder *xdr.Encoder, value []DontHave, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeDontHave(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union StellarMessage switch (MessageType type) {
//       case ERROR_MSG:
//           Error error;
//       case HELLO:
//           Hello hello;
//       case DONT_HAVE:
//           DontHave dontHave;
//       case GET_PEERS:
//           void;
//       case PEERS:
//           PeerAddress peers<>;
//
//       case GET_TX_SET:
//           uint256 txSetHash;
//       case TX_SET:
//           TransactionSet txSet;
//
//       case GET_VALIDATIONS:
//           uint256 ledgerHash;
//       case VALIDATIONS:
//           SCPEnvelope validations<>;
//
//       case TRANSACTION:
//           TransactionEnvelope transaction;
//
//       // SCP
//       case GET_SCP_QUORUMSET:
//           uint256 qSetHash;
//       case SCP_QUORUMSET:
//           SCPQuorumSet qSet;
//       case SCP_MESSAGE:
//           SCPEnvelope envelope;
//   };
//
// ===========================================================================

type StellarMessage struct {
	aType       MessageType
	error       *Error
	hello       *Hello
	dontHave    *DontHave
	peers       *[]PeerAddress
	txSetHash   *Uint256
	txSet       *TransactionSet
	ledgerHash  *Uint256
	validations *[]ScpEnvelope
	transaction *TransactionEnvelope
	qSetHash    *Uint256
	qSet        *ScpQuorumSet
	envelope    *ScpEnvelope
}

func NewStellarMessageErrorMsg(val Error) StellarMessage {
	return StellarMessage{
		aType: MessageTypeErrorMsg,
		error: &val,
	}
}
func NewStellarMessageHello(val Hello) StellarMessage {
	return StellarMessage{
		aType: MessageTypeHello,
		hello: &val,
	}
}
func NewStellarMessageDontHave(val DontHave) StellarMessage {
	return StellarMessage{
		aType:    MessageTypeDontHave,
		dontHave: &val,
	}
}
func NewStellarMessageGetPeer() StellarMessage {
	return StellarMessage{
		aType: MessageTypeGetPeer,
	}
}
func NewStellarMessagePeer(val []PeerAddress) StellarMessage {
	return StellarMessage{
		aType: MessageTypePeer,
		peers: &val,
	}
}
func NewStellarMessageGetTxSet(val Uint256) StellarMessage {
	return StellarMessage{
		aType:     MessageTypeGetTxSet,
		txSetHash: &val,
	}
}
func NewStellarMessageTxSet(val TransactionSet) StellarMessage {
	return StellarMessage{
		aType: MessageTypeTxSet,
		txSet: &val,
	}
}
func NewStellarMessageGetValidation(val Uint256) StellarMessage {
	return StellarMessage{
		aType:      MessageTypeGetValidation,
		ledgerHash: &val,
	}
}
func NewStellarMessageValidation(val []ScpEnvelope) StellarMessage {
	return StellarMessage{
		aType:       MessageTypeValidation,
		validations: &val,
	}
}
func NewStellarMessageTransaction(val TransactionEnvelope) StellarMessage {
	return StellarMessage{
		aType:       MessageTypeTransaction,
		transaction: &val,
	}
}

func NewStellarMessageGetScpQuorumset(val Uint256) StellarMessage {
	return StellarMessage{
		aType:    MessageTypeGetScpQuorumset,
		qSetHash: &val,
	}
}
func NewStellarMessageScpQuorumset(val ScpQuorumSet) StellarMessage {
	return StellarMessage{
		aType: MessageTypeScpQuorumset,
		qSet:  &val,
	}
}
func NewStellarMessageScpMessage(val ScpEnvelope) StellarMessage {
	return StellarMessage{
		aType:    MessageTypeScpMessage,
		envelope: &val,
	}
}
func (u *StellarMessage) Type() MessageType {
	return u.aType
}
func (u *StellarMessage) Error() Error {
	return *u.error
}
func (u *StellarMessage) Hello() Hello {
	return *u.hello
}
func (u *StellarMessage) DontHave() DontHave {
	return *u.dontHave
}
func (u *StellarMessage) Peers() []PeerAddress {
	return *u.peers
}
func (u *StellarMessage) TxSetHash() Uint256 {
	return *u.txSetHash
}
func (u *StellarMessage) TxSet() TransactionSet {
	return *u.txSet
}
func (u *StellarMessage) LedgerHash() Uint256 {
	return *u.ledgerHash
}
func (u *StellarMessage) Validations() []ScpEnvelope {
	return *u.validations
}
func (u *StellarMessage) Transaction() TransactionEnvelope {
	return *u.transaction
}
func (u *StellarMessage) QSetHash() Uint256 {
	return *u.qSetHash
}
func (u *StellarMessage) QSet() ScpQuorumSet {
	return *u.qSet
}
func (u *StellarMessage) Envelope() ScpEnvelope {
	return *u.envelope
}
func DecodeStellarMessage(decoder *xdr.Decoder, result *StellarMessage) (int, error) {
	var (
		discriminant MessageType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeMessageType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == MessageTypeErrorMsg {
		var errorMsg Error
		bytesRead, err = DecodeError(decoder, &errorMsg)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageErrorMsg(errorMsg)
	}

	if discriminant == MessageTypeHello {
		var hello Hello
		bytesRead, err = DecodeHello(decoder, &hello)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageHello(hello)
	}

	if discriminant == MessageTypeDontHave {
		var dontHave DontHave
		bytesRead, err = DecodeDontHave(decoder, &dontHave)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageDontHave(dontHave)
	}

	if discriminant == MessageTypePeer {
		var peers []PeerAddress
		bytesRead, err = DecodePeerAddressArray(decoder, &peers, MaxXdrElements)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessagePeer(peers)
	}

	if discriminant == MessageTypeGetTxSet {
		var getTxSet Uint256
		bytesRead, err = DecodeUint256(decoder, &getTxSet)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageGetTxSet(getTxSet)
	}

	if discriminant == MessageTypeTxSet {
		var txSet TransactionSet
		bytesRead, err = DecodeTransactionSet(decoder, &txSet)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageTxSet(txSet)
	}

	if discriminant == MessageTypeGetValidation {
		var getValidations Uint256
		bytesRead, err = DecodeUint256(decoder, &getValidations)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageGetValidation(getValidations)
	}

	if discriminant == MessageTypeValidation {
		var validations []ScpEnvelope
		bytesRead, err = DecodeScpEnvelopeArray(decoder, &validations, MaxXdrElements)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageValidation(validations)
	}

	if discriminant == MessageTypeTransaction {
		var transaction TransactionEnvelope
		bytesRead, err = DecodeTransactionEnvelope(decoder, &transaction)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageTransaction(transaction)
	}

	if discriminant == MessageTypeGetScpQuorumset {
		var getScpQuorumset Uint256
		bytesRead, err = DecodeUint256(decoder, &getScpQuorumset)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageGetScpQuorumset(getScpQuorumset)
	}

	if discriminant == MessageTypeScpQuorumset {
		var scpQuorumset ScpQuorumSet
		bytesRead, err = DecodeScpQuorumSet(decoder, &scpQuorumset)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageScpQuorumset(scpQuorumset)
	}

	if discriminant == MessageTypeScpMessage {
		var scpMessage ScpEnvelope
		bytesRead, err = DecodeScpEnvelope(decoder, &scpMessage)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewStellarMessageScpMessage(scpMessage)
	}

	return totalRead, nil
}
func DecodeOptionalStellarMessage(decoder *xdr.Decoder, result **StellarMessage) (int, error) {
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

	var val StellarMessage

	bytesRead, err = DecodeStellarMessage(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeStellarMessageFixedArray(decoder *xdr.Decoder, result []StellarMessage, size int) (int, error) {
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
		bytesRead, err = DecodeStellarMessage(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeStellarMessageArray(decoder *xdr.Decoder, result *[]StellarMessage, maxSize int32) (int, error) {
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

	var theResult = make([]StellarMessage, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeStellarMessage(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeStellarMessage(encoder *xdr.Encoder, value *StellarMessage) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeMessageType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == MessageTypeErrorMsg {
		bytesWritten, err = EncodeError(encoder, value.error)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeHello {
		bytesWritten, err = EncodeHello(encoder, value.hello)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeDontHave {
		bytesWritten, err = EncodeDontHave(encoder, value.dontHave)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypePeer {
		bytesWritten, err = EncodePeerAddressArray(encoder, value.peers, MaxXdrElements)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeGetTxSet {
		bytesWritten, err = EncodeUint256(encoder, value.txSetHash)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeTxSet {
		bytesWritten, err = EncodeTransactionSet(encoder, value.txSet)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeGetValidation {
		bytesWritten, err = EncodeUint256(encoder, value.ledgerHash)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeValidation {
		bytesWritten, err = EncodeScpEnvelopeArray(encoder, value.validations, MaxXdrElements)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeTransaction {
		bytesWritten, err = EncodeTransactionEnvelope(encoder, value.transaction)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeGetScpQuorumset {
		bytesWritten, err = EncodeUint256(encoder, value.qSetHash)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeScpQuorumset {
		bytesWritten, err = EncodeScpQuorumSet(encoder, value.qSet)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	if value.aType == MessageTypeScpMessage {
		bytesWritten, err = EncodeScpEnvelope(encoder, value.envelope)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalStellarMessage(encoder *xdr.Encoder, value *StellarMessage) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeStellarMessage(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeStellarMessageFixedArray(encoder *xdr.Encoder, value []StellarMessage, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeStellarMessage(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeStellarMessageArray(encoder *xdr.Encoder, value []StellarMessage, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeStellarMessage(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   typedef opaque Value<>;
//
// ===========================================================================

type Value []byte

func DecodeValue(decoder *xdr.Decoder, result *Value) (int, error) {
	var (
		val       []byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Value(val)
	return totalRead, nil
}
func DecodeOptionalValue(decoder *xdr.Decoder, result **Value) (int, error) {
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

	var val Value

	bytesRead, err = DecodeValue(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeValueFixedArray(decoder *xdr.Decoder, result []Value, size int) (int, error) {
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
		bytesRead, err = DecodeValue(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeValueArray(decoder *xdr.Decoder, result *[]Value, maxSize int32) (int, error) {
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

	var theResult = make([]Value, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeValue(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeValue(encoder *xdr.Encoder, value *Value) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[]byte)(value)[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalValue(encoder *xdr.Encoder, value *Value) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeValue(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeValueFixedArray(encoder *xdr.Encoder, value []Value, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeValue(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeValueArray(encoder *xdr.Encoder, value []Value, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeValue(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   typedef opaque Evidence<>;
//
// ===========================================================================

type Evidence []byte

func DecodeEvidence(decoder *xdr.Decoder, result *Evidence) (int, error) {
	var (
		val       []byte
		totalRead int
		bytesRead int
		err       error
	)

	bytesRead, err = DecodeFixedOpaque(decoder, val[:], MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = Evidence(val)
	return totalRead, nil
}
func DecodeOptionalEvidence(decoder *xdr.Decoder, result **Evidence) (int, error) {
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

	var val Evidence

	bytesRead, err = DecodeEvidence(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeEvidenceFixedArray(decoder *xdr.Decoder, result []Evidence, size int) (int, error) {
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
		bytesRead, err = DecodeEvidence(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeEvidenceArray(decoder *xdr.Decoder, result *[]Evidence, maxSize int32) (int, error) {
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

	var theResult = make([]Evidence, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeEvidence(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeEvidence(encoder *xdr.Encoder, value *Evidence) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeFixedOpaque(encoder, (*[]byte)(value)[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, err
}
func EncodeOptionalEvidence(encoder *xdr.Encoder, value *Evidence) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeEvidence(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeEvidenceFixedArray(encoder *xdr.Encoder, value []Evidence, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeEvidence(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeEvidenceArray(encoder *xdr.Encoder, value []Evidence, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeEvidence(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct SCPBallot
//   {
//       uint32 counter;     // n
//       Value value;        // x
//   };
//
// ===========================================================================

type ScpBallot struct {
	Counter Uint32
	Value   Value
}

func DecodeScpBallot(decoder *xdr.Decoder, result *ScpBallot) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint32(decoder, &result.Counter)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeValue(decoder, &result.Value)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalScpBallot(decoder *xdr.Decoder, result **ScpBallot) (int, error) {
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

	var val ScpBallot

	bytesRead, err = DecodeScpBallot(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeScpBallotFixedArray(decoder *xdr.Decoder, result []ScpBallot, size int) (int, error) {
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
		bytesRead, err = DecodeScpBallot(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeScpBallotArray(decoder *xdr.Decoder, result *[]ScpBallot, maxSize int32) (int, error) {
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

	var theResult = make([]ScpBallot, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeScpBallot(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeScpBallot(encoder *xdr.Encoder, value *ScpBallot) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint32(encoder, &value.Counter)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeValue(encoder, &value.Value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalScpBallot(encoder *xdr.Encoder, value *ScpBallot) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeScpBallot(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeScpBallotFixedArray(encoder *xdr.Encoder, value []ScpBallot, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpBallot(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeScpBallotArray(encoder *xdr.Encoder, value []ScpBallot, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpBallot(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   enum SCPStatementType
//   {
//       PREPARE,
//       PREPARED,
//       COMMIT,
//       COMMITTED
//   };
//
// ===========================================================================

type ScpStatementType int32

const (
	ScpStatementTypePrepare   ScpStatementType = 0
	ScpStatementTypePrepared                   = 1
	ScpStatementTypeCommit                     = 2
	ScpStatementTypeCommitted                  = 3
)

var ScpStatementTypeMap = map[int32]bool{
	0: true,
	1: true,
	2: true,
	3: true,
}

func DecodeScpStatementType(decoder *xdr.Decoder, result *ScpStatementType) (int, error) {
	val, bytesRead, err := decoder.DecodeEnum(ScpStatementTypeMap)

	if err != nil {
		return bytesRead, err
	}
	*result = ScpStatementType(val)
	return bytesRead, err
}
func DecodeOptionalScpStatementType(decoder *xdr.Decoder, result **ScpStatementType) (int, error) {
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

	var val ScpStatementType

	bytesRead, err = DecodeScpStatementType(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeScpStatementTypeFixedArray(decoder *xdr.Decoder, result []ScpStatementType, size int) (int, error) {
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
		bytesRead, err = DecodeScpStatementType(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeScpStatementTypeArray(decoder *xdr.Decoder, result *[]ScpStatementType, maxSize int32) (int, error) {
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

	var theResult = make([]ScpStatementType, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeScpStatementType(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeScpStatementType(encoder *xdr.Encoder, value *ScpStatementType) (int, error) {
	bytesWritten, err := encoder.EncodeEnum(int32(*value), ScpStatementTypeMap)
	return bytesWritten, err
}
func EncodeOptionalScpStatementType(encoder *xdr.Encoder, value *ScpStatementType) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeScpStatementType(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeScpStatementTypeFixedArray(encoder *xdr.Encoder, value []ScpStatementType, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatementType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeScpStatementTypeArray(encoder *xdr.Encoder, value []ScpStatementType, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatementType(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct
//               {
//                   SCPBallot excepted<>;  // B_c
//                   SCPBallot* prepared;   // p
//               }
//
// ===========================================================================

type ScpStatementPrepare struct {
	Excepted []ScpBallot
	Prepared *ScpBallot
}

func DecodeScpStatementPrepare(decoder *xdr.Decoder, result *ScpStatementPrepare) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeScpBallotArray(decoder, &result.Excepted, MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	var prepared *ScpBallot
	bytesRead, err = DecodeOptionalScpBallot(decoder, &prepared)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	result.Prepared = prepared

	return totalRead, nil
}
func DecodeOptionalScpStatementPrepare(decoder *xdr.Decoder, result **ScpStatementPrepare) (int, error) {
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

	var val ScpStatementPrepare

	bytesRead, err = DecodeScpStatementPrepare(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeScpStatementPrepareFixedArray(decoder *xdr.Decoder, result []ScpStatementPrepare, size int) (int, error) {
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
		bytesRead, err = DecodeScpStatementPrepare(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeScpStatementPrepareArray(decoder *xdr.Decoder, result *[]ScpStatementPrepare, maxSize int32) (int, error) {
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

	var theResult = make([]ScpStatementPrepare, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeScpStatementPrepare(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeScpStatementPrepare(encoder *xdr.Encoder, value *ScpStatementPrepare) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeScpBallotArray(encoder, value.Excepted[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeOptionalScpBallot(encoder, value.Prepared)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalScpStatementPrepare(encoder *xdr.Encoder, value *ScpStatementPrepare) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeScpStatementPrepare(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeScpStatementPrepareFixedArray(encoder *xdr.Encoder, value []ScpStatementPrepare, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatementPrepare(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeScpStatementPrepareArray(encoder *xdr.Encoder, value []ScpStatementPrepare, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatementPrepare(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   union switch (SCPStatementType type)
//       {
//           case PREPARE:
//               struct
//               {
//                   SCPBallot excepted<>;  // B_c
//                   SCPBallot* prepared;   // p
//               } prepare;
//           case PREPARED:
//           case COMMIT:
//           case COMMITTED:
//               void;
//       }
//
// ===========================================================================

type ScpStatementPledges struct {
	aType   ScpStatementType
	prepare *ScpStatementPrepare
}

func NewScpStatementPledgesPrepare(val ScpStatementPrepare) ScpStatementPledges {
	return ScpStatementPledges{
		aType:   ScpStatementTypePrepare,
		prepare: &val,
	}
}
func NewScpStatementPledgesPrepared() ScpStatementPledges {
	return ScpStatementPledges{
		aType: ScpStatementTypePrepared,
	}
}
func NewScpStatementPledgesCommit() ScpStatementPledges {
	return ScpStatementPledges{
		aType: ScpStatementTypeCommit,
	}
}
func NewScpStatementPledgesCommitted() ScpStatementPledges {
	return ScpStatementPledges{
		aType: ScpStatementTypeCommitted,
	}
}
func (u *ScpStatementPledges) Type() ScpStatementType {
	return u.aType
}
func (u *ScpStatementPledges) Prepare() ScpStatementPrepare {
	return *u.prepare
}
func DecodeScpStatementPledges(decoder *xdr.Decoder, result *ScpStatementPledges) (int, error) {
	var (
		discriminant ScpStatementType
		totalRead    int
		bytesRead    int
		err          error
	)

	bytesRead, err = DecodeScpStatementType(decoder, &discriminant)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	if discriminant == ScpStatementTypePrepare {
		var prepare ScpStatementPrepare
		bytesRead, err = DecodeScpStatementPrepare(decoder, &prepare)
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

		*result = NewScpStatementPledgesPrepare(prepare)
	}

	return totalRead, nil
}
func DecodeOptionalScpStatementPledges(decoder *xdr.Decoder, result **ScpStatementPledges) (int, error) {
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

	var val ScpStatementPledges

	bytesRead, err = DecodeScpStatementPledges(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeScpStatementPledgesFixedArray(decoder *xdr.Decoder, result []ScpStatementPledges, size int) (int, error) {
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
		bytesRead, err = DecodeScpStatementPledges(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeScpStatementPledgesArray(decoder *xdr.Decoder, result *[]ScpStatementPledges, maxSize int32) (int, error) {
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

	var theResult = make([]ScpStatementPledges, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeScpStatementPledges(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeScpStatementPledges(encoder *xdr.Encoder, value *ScpStatementPledges) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	bytesWritten, err = EncodeScpStatementType(encoder, &value.aType)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	if value.aType == ScpStatementTypePrepare {
		bytesWritten, err = EncodeScpStatementPrepare(encoder, value.prepare)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
func EncodeOptionalScpStatementPledges(encoder *xdr.Encoder, value *ScpStatementPledges) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeScpStatementPledges(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeScpStatementPledgesFixedArray(encoder *xdr.Encoder, value []ScpStatementPledges, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatementPledges(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeScpStatementPledgesArray(encoder *xdr.Encoder, value []ScpStatementPledges, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatementPledges(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct SCPStatement
//   {
//       uint64 slotIndex;      // i
//       SCPBallot ballot;      // b
//       Hash quorumSetHash;    // D
//
//       union switch (SCPStatementType type)
//       {
//           case PREPARE:
//               struct
//               {
//                   SCPBallot excepted<>;  // B_c
//                   SCPBallot* prepared;   // p
//               } prepare;
//           case PREPARED:
//           case COMMIT:
//           case COMMITTED:
//               void;
//       } pledges;
//   };
//
// ===========================================================================

type ScpStatement struct {
	SlotIndex     Uint64
	Ballot        ScpBallot
	QuorumSetHash Hash
	Pledges       ScpStatementPledges
}

func DecodeScpStatement(decoder *xdr.Decoder, result *ScpStatement) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint64(decoder, &result.SlotIndex)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeScpBallot(decoder, &result.Ballot)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeHash(decoder, &result.QuorumSetHash)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeScpStatementPledges(decoder, &result.Pledges)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalScpStatement(decoder *xdr.Decoder, result **ScpStatement) (int, error) {
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

	var val ScpStatement

	bytesRead, err = DecodeScpStatement(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeScpStatementFixedArray(decoder *xdr.Decoder, result []ScpStatement, size int) (int, error) {
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
		bytesRead, err = DecodeScpStatement(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeScpStatementArray(decoder *xdr.Decoder, result *[]ScpStatement, maxSize int32) (int, error) {
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

	var theResult = make([]ScpStatement, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeScpStatement(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeScpStatement(encoder *xdr.Encoder, value *ScpStatement) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint64(encoder, &value.SlotIndex)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeScpBallot(encoder, &value.Ballot)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeHash(encoder, &value.QuorumSetHash)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeScpStatementPledges(encoder, &value.Pledges)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalScpStatement(encoder *xdr.Encoder, value *ScpStatement) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeScpStatement(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeScpStatementFixedArray(encoder *xdr.Encoder, value []ScpStatement, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatement(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeScpStatementArray(encoder *xdr.Encoder, value []ScpStatement, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpStatement(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct SCPEnvelope
//   {
//       uint256 nodeID;         // v
//       SCPStatement statement;
//       Signature signature;
//   };
//
// ===========================================================================

type ScpEnvelope struct {
	NodeId    Uint256
	Statement ScpStatement
	Signature Signature
}

func DecodeScpEnvelope(decoder *xdr.Decoder, result *ScpEnvelope) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint256(decoder, &result.NodeId)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeScpStatement(decoder, &result.Statement)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeSignature(decoder, &result.Signature)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalScpEnvelope(decoder *xdr.Decoder, result **ScpEnvelope) (int, error) {
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

	var val ScpEnvelope

	bytesRead, err = DecodeScpEnvelope(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeScpEnvelopeFixedArray(decoder *xdr.Decoder, result []ScpEnvelope, size int) (int, error) {
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
		bytesRead, err = DecodeScpEnvelope(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeScpEnvelopeArray(decoder *xdr.Decoder, result *[]ScpEnvelope, maxSize int32) (int, error) {
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

	var theResult = make([]ScpEnvelope, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeScpEnvelope(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeScpEnvelope(encoder *xdr.Encoder, value *ScpEnvelope) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint256(encoder, &value.NodeId)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeScpStatement(encoder, &value.Statement)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeSignature(encoder, &value.Signature)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalScpEnvelope(encoder *xdr.Encoder, value *ScpEnvelope) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeScpEnvelope(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeScpEnvelopeFixedArray(encoder *xdr.Encoder, value []ScpEnvelope, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpEnvelope(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeScpEnvelopeArray(encoder *xdr.Encoder, value []ScpEnvelope, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpEnvelope(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

// === xdr source ============================================================
//
//   struct SCPQuorumSet
//   {
//       uint32 threshold;
//       Hash validators<>;
//   };
//
// ===========================================================================

type ScpQuorumSet struct {
	Threshold  Uint32
	Validators []Hash
}

func DecodeScpQuorumSet(decoder *xdr.Decoder, result *ScpQuorumSet) (int, error) {
	totalRead := 0
	bytesRead := 0
	var err error

	bytesRead, err = DecodeUint32(decoder, &result.Threshold)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	bytesRead, err = DecodeHashArray(decoder, &result.Validators, MaxXdrElements)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
func DecodeOptionalScpQuorumSet(decoder *xdr.Decoder, result **ScpQuorumSet) (int, error) {
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

	var val ScpQuorumSet

	bytesRead, err = DecodeScpQuorumSet(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}
func DecodeScpQuorumSetFixedArray(decoder *xdr.Decoder, result []ScpQuorumSet, size int) (int, error) {
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
		bytesRead, err = DecodeScpQuorumSet(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeScpQuorumSetArray(decoder *xdr.Decoder, result *[]ScpQuorumSet, maxSize int32) (int, error) {
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

	var theResult = make([]ScpQuorumSet, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeScpQuorumSet(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func EncodeScpQuorumSet(encoder *xdr.Encoder, value *ScpQuorumSet) (int, error) {
	totalWritten := 0
	bytesWritten := 0
	var err error

	bytesWritten, err = EncodeUint32(encoder, &value.Threshold)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	bytesWritten, err = EncodeHashArray(encoder, value.Validators[:], MaxXdrElements)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeOptionalScpQuorumSet(encoder *xdr.Encoder, value *ScpQuorumSet) (int, error) {
	var (
		isPresent    bool
		totalWritten int
		bytesWritten int
		err          error
	)
	isPresent = value != nil
	bytesWritten, err = EncodeBool(encoder, &isPresent)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	if !isPresent {
		return totalWritten, nil
	}

	bytesWritten, err = EncodeScpQuorumSet(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}
func EncodeScpQuorumSetFixedArray(encoder *xdr.Encoder, value []ScpQuorumSet, size int) (int, error) {
	var (
		totalWritten int
		bytesWritten int
		err          error
	)

	if len(value) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(value), size)
		return 0, errors.New(errMsg)
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpQuorumSet(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeScpQuorumSetArray(encoder *xdr.Encoder, value []ScpQuorumSet, maxSize int32) (int, error) {
	var (
		size         int32
		totalWritten int
		bytesWritten int
		err          error
	)

	size = int32(len(value))

	if size > maxSize {
		errMsg := fmt.Sprintf("xdr: value too large:%d, max:%d", size, maxSize)
		return totalWritten, errors.New(errMsg)
	}

	bytesWritten, err = EncodeInt(encoder, &size)
	totalWritten += bytesWritten

	if err != nil {
		return totalWritten, err
	}

	for _, element := range value {
		bytesWritten, err = EncodeScpQuorumSet(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}
