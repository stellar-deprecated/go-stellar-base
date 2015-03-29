// Automatically generated from xdr/Stellar-ledger.x
// DO NOT EDIT or your changes may be overwritten

package xdr

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-xdr/xdr2"
)

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
