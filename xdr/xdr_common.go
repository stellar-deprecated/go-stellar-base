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

func EncodeInt(encoder *xdr.Encoder, value *int32) (int, error) {
	return encoder.EncodeInt(*value)
}

func EncodeOptionalInt(encoder *xdr.Encoder, value *int32) (int, error) {
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

	bytesWritten, err = EncodeInt(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func EncodeIntFixedArray(encoder *xdr.Encoder, value []int32, size int) (int, error) {
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
		bytesWritten, err = EncodeInt(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeIntArray(encoder *xdr.Encoder, value []int32, maxSize int32) (int, error) {
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
		bytesWritten, err = EncodeInt(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeUint(encoder *xdr.Encoder, value *uint32) (int, error) {
	return encoder.EncodeUint(*value)
}

func EncodeOptionalUint(encoder *xdr.Encoder, value *uint32) (int, error) {
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

	bytesWritten, err = EncodeUint(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func EncodeUintFixedArray(encoder *xdr.Encoder, value []uint32, size int) (int, error) {
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
		bytesWritten, err = EncodeUint(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeUintArray(encoder *xdr.Encoder, value []uint32, maxSize int32) (int, error) {
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
		bytesWritten, err = EncodeUint(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeHyper(encoder *xdr.Encoder, value *int64) (int, error) {
	return encoder.EncodeHyper(*value)
}

func EncodeOptionalHyper(encoder *xdr.Encoder, value *int64) (int, error) {
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

	bytesWritten, err = EncodeHyper(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func EncodeHyperFixedArray(encoder *xdr.Encoder, value []int64, size int) (int, error) {
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
		bytesWritten, err = EncodeHyper(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeHyperArray(encoder *xdr.Encoder, value []int64, maxSize int32) (int, error) {
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
		bytesWritten, err = EncodeHyper(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeUhyper(encoder *xdr.Encoder, value *uint64) (int, error) {
	return encoder.EncodeUhyper(*value)
}

func EncodeOptionalUhyper(encoder *xdr.Encoder, value *uint64) (int, error) {
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

	bytesWritten, err = EncodeUhyper(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func EncodeUhyperFixedArray(encoder *xdr.Encoder, value []uint64, size int) (int, error) {
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
		bytesWritten, err = EncodeUhyper(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeUhyperArray(encoder *xdr.Encoder, value []uint64, maxSize int32) (int, error) {
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
		bytesWritten, err = EncodeUhyper(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeFloat(encoder *xdr.Encoder, value *float32) (int, error) {
	return encoder.EncodeFloat(*value)
}

func EncodeOptionalFloat(encoder *xdr.Encoder, value *float32) (int, error) {
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

	bytesWritten, err = EncodeFloat(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func EncodeFloatFixedArray(encoder *xdr.Encoder, value []float32, size int) (int, error) {
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
		bytesWritten, err = EncodeFloat(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeFloatArray(encoder *xdr.Encoder, value []float32, maxSize int32) (int, error) {
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
		bytesWritten, err = EncodeFloat(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeDouble(encoder *xdr.Encoder, value *float64) (int, error) {
	return encoder.EncodeDouble(*value)
}

func EncodeOptionalDouble(encoder *xdr.Encoder, value *float64) (int, error) {
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

	bytesWritten, err = EncodeDouble(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func EncodeDoubleFixedArray(encoder *xdr.Encoder, value []float64, size int) (int, error) {
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
		bytesWritten, err = EncodeDouble(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeDoubleArray(encoder *xdr.Encoder, value []float64, maxSize int32) (int, error) {
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
		bytesWritten, err = EncodeDouble(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeBool(encoder *xdr.Encoder, value *bool) (int, error) {
	return encoder.EncodeBool(*value)
}

func EncodeOptionalBool(encoder *xdr.Encoder, value *bool) (int, error) {
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

	bytesWritten, err = EncodeBool(encoder, value)
	totalWritten += bytesWritten
	if err != nil {
		return totalWritten, err
	}

	return totalWritten, nil
}

func EncodeBoolFixedArray(encoder *xdr.Encoder, value []bool, size int) (int, error) {
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
		bytesWritten, err = EncodeBool(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
}

func EncodeBoolArray(encoder *xdr.Encoder, value []bool, maxSize int32) (int, error) {
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
		bytesWritten, err = EncodeBool(encoder, &element)
		totalWritten += bytesWritten
		if err != nil {
			return totalWritten, err
		}

	}

	return totalWritten, nil
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

func EncodeFixedOpaque(encoder *xdr.Encoder, val []byte, size int32) (int, error) {

	if len(val) != int(size) {
		errMsg := fmt.Sprintf("xdr: value wrong size:%d, expected:%d", len(val), size)
		return 0, errors.New(errMsg)
	}

	return encoder.EncodeFixedOpaque(val)
}

func EncodeOpaque(encoder *xdr.Encoder, val []byte, maxSize int32) (int, error) {

	if len(val) > int(maxSize) {
		errMsg := fmt.Sprintf("xdr: value to large:%d, max:%d", len(val), maxSize)
		return 0, errors.New(errMsg)
	}

	return encoder.EncodeOpaque(val)
}

func EncodeString(encoder *xdr.Encoder, val *string, maxSize int32) (int, error) {

	if len(*val) > int(maxSize) {
		errMsg := fmt.Sprintf("xdr: value to large:%d, max:%d", len(*val), maxSize)
		return 0, errors.New(errMsg)
	}

	return encoder.EncodeString(*val)
}
