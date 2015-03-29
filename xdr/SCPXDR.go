// Automatically generated from xdr/SCPXDR.x
// DO NOT EDIT or your changes may be overwritten

package xdr

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-xdr/xdr2"
)

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

	*result = ScpStatementPledges{}
	return totalRead, nil
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
