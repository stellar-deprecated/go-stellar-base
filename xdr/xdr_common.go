// Automatically generated from #{@output.source_paths.join(",")}
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-xdr/xdr2"
	"math"
)

const MaxXdrElements = math.MaxInt32

// Primitive decoders

func DecodeInt(decoder *xdr.Decoder, result *int32) (int, error) {
	val, bytesRead, err := decoder.DecodeInt()

	if err == nil {
		*result = val
	}

	return bytesRead, err
}

func DecodeOptionalInt(decoder *xdr.Decoder, result **int32) (int, error) {
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

	var val int32

	bytesRead, err = DecodeInt(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}

func DecodeIntFixedArray(decoder *xdr.Decoder, result []int32, size int) (int, error) {
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
		bytesRead, err = DecodeInt(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeIntArray(decoder *xdr.Decoder, result *[]int32, maxSize int32) (int, error) {
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

	var theResult = make([]int32, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeInt(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUint(decoder *xdr.Decoder, result *uint32) (int, error) {
	val, bytesRead, err := decoder.DecodeUint()

	if err == nil {
		*result = val
	}

	return bytesRead, err
}

func DecodeOptionalUint(decoder *xdr.Decoder, result **uint32) (int, error) {
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

	var val uint32

	bytesRead, err = DecodeUint(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}

func DecodeUintFixedArray(decoder *xdr.Decoder, result []uint32, size int) (int, error) {
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
		bytesRead, err = DecodeUint(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUintArray(decoder *xdr.Decoder, result *[]uint32, maxSize int32) (int, error) {
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

	var theResult = make([]uint32, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeUint(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeHyper(decoder *xdr.Decoder, result *int64) (int, error) {
	val, bytesRead, err := decoder.DecodeHyper()

	if err == nil {
		*result = val
	}

	return bytesRead, err
}

func DecodeOptionalHyper(decoder *xdr.Decoder, result **int64) (int, error) {
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

	var val int64

	bytesRead, err = DecodeHyper(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}

func DecodeHyperFixedArray(decoder *xdr.Decoder, result []int64, size int) (int, error) {
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
		bytesRead, err = DecodeHyper(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeHyperArray(decoder *xdr.Decoder, result *[]int64, maxSize int32) (int, error) {
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

	var theResult = make([]int64, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeHyper(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUhyper(decoder *xdr.Decoder, result *uint64) (int, error) {
	val, bytesRead, err := decoder.DecodeUhyper()

	if err == nil {
		*result = val
	}

	return bytesRead, err
}

func DecodeOptionalUhyper(decoder *xdr.Decoder, result **uint64) (int, error) {
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

	var val uint64

	bytesRead, err = DecodeUhyper(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}

func DecodeUhyperFixedArray(decoder *xdr.Decoder, result []uint64, size int) (int, error) {
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
		bytesRead, err = DecodeUhyper(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeUhyperArray(decoder *xdr.Decoder, result *[]uint64, maxSize int32) (int, error) {
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

	var theResult = make([]uint64, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeUhyper(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeFloat(decoder *xdr.Decoder, result *float32) (int, error) {
	val, bytesRead, err := decoder.DecodeFloat()

	if err == nil {
		*result = val
	}

	return bytesRead, err
}

func DecodeOptionalFloat(decoder *xdr.Decoder, result **float32) (int, error) {
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

	var val float32

	bytesRead, err = DecodeFloat(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}

func DecodeFloatFixedArray(decoder *xdr.Decoder, result []float32, size int) (int, error) {
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
		bytesRead, err = DecodeFloat(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeFloatArray(decoder *xdr.Decoder, result *[]float32, maxSize int32) (int, error) {
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

	var theResult = make([]float32, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeFloat(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeDouble(decoder *xdr.Decoder, result *float64) (int, error) {
	val, bytesRead, err := decoder.DecodeDouble()

	if err == nil {
		*result = val
	}

	return bytesRead, err
}

func DecodeOptionalDouble(decoder *xdr.Decoder, result **float64) (int, error) {
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

	var val float64

	bytesRead, err = DecodeDouble(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}

func DecodeDoubleFixedArray(decoder *xdr.Decoder, result []float64, size int) (int, error) {
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
		bytesRead, err = DecodeDouble(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeDoubleArray(decoder *xdr.Decoder, result *[]float64, maxSize int32) (int, error) {
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

	var theResult = make([]float64, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeDouble(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeBool(decoder *xdr.Decoder, result *bool) (int, error) {
	val, bytesRead, err := decoder.DecodeBool()

	if err == nil {
		*result = val
	}

	return bytesRead, err
}

func DecodeOptionalBool(decoder *xdr.Decoder, result **bool) (int, error) {
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

	var val bool

	bytesRead, err = DecodeBool(decoder, &val)
	totalRead += bytesRead
	if err != nil {
		return totalRead, err
	}

	*result = &val

	return totalRead, nil
}

func DecodeBoolFixedArray(decoder *xdr.Decoder, result []bool, size int) (int, error) {
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
		bytesRead, err = DecodeBool(decoder, &result[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeBoolArray(decoder *xdr.Decoder, result *[]bool, maxSize int32) (int, error) {
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

	var theResult = make([]bool, size)
	*result = theResult

	for i := int32(0); i < size; i++ {
		bytesRead, err = DecodeBool(decoder, &theResult[i])
		totalRead += bytesRead
		if err != nil {
			return totalRead, err
		}

	}

	return totalRead, nil
}

func DecodeFixedOpaque(decoder *xdr.Decoder, result []byte, size int32) (int, error) {

	if len(result) != int(size) {
		errMsg := fmt.Sprintf("xdr: bad opaque len:%d, expected %d", len(result), size)
		return 0, errors.New(errMsg)
	}

	val, bytesRead, err := decoder.DecodeFixedOpaque(size)

	if err == nil {
		copy(result, val)
	}

	return bytesRead, err
}

func DecodeOpaque(decoder *xdr.Decoder, result *[]byte, maxSize int32) (int, error) {
	val, bytesRead, err := decoder.DecodeOpaque()

	if err != nil {
		return bytesRead, err
	}

	if len(val) > int(maxSize) {
		errMsg := fmt.Sprintf("xdr: encoded opaque size too large:%d, max:%d", len(val), maxSize)
		return bytesRead, errors.New(errMsg)
	}

	*result = val

	return bytesRead, err
}

func DecodeString(decoder *xdr.Decoder, result *string, maxSize int32) (int, error) {
	val, bytesRead, err := decoder.DecodeString()

	if err != nil {
		return bytesRead, err
	}

	if len(val) > int(maxSize) {
		errMsg := fmt.Sprintf("xdr: encoded string size too large:%d, max:%d", len(val), maxSize)
		return bytesRead, errors.New(errMsg)
	}

	*result = val

	return bytesRead, err
}
