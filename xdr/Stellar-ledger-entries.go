// Automatically generated from xdr/Stellar-ledger-entries.x
// DO NOT EDIT or your changes may be overwritten

package xdr

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-xdr/xdr2"
)

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

	*result = LedgerEntry{}
	return totalRead, nil
}
