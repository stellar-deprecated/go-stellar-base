// Automatically generated from xdr/Stellar-overlay.x
// DO NOT EDIT or your changes may be overwritten

package xdr

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-xdr/xdr2"
)

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

	bytesRead, err = DecodeString(decoder, result.Msg[:], 100)
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

	bytesRead, err = DecodeString(decoder, result.VersionStr[:], 100)
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
	peers       *PeerAddress
	txSetHash   *Uint256
	txSet       *TransactionSet
	ledgerHash  *Uint256
	validations *ScpEnvelope
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
func NewStellarMessagePeer(val PeerAddress) StellarMessage {
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
func NewStellarMessageValidation(val ScpEnvelope) StellarMessage {
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
func (u *StellarMessage) Peers() PeerAddress {
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
func (u *StellarMessage) Validations() ScpEnvelope {
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

	*result = StellarMessage{}
	return totalRead, nil
}
