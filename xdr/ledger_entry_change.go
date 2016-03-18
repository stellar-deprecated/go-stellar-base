package xdr

import "fmt"

// EntryType is a helper to get at the entry type for a change.
func (change *LedgerEntryChange) EntryType() LedgerEntryType {
	switch change.Type {
	case LedgerEntryChangeTypeLedgerEntryCreated:
		return change.MustCreated().Data.Type
	case LedgerEntryChangeTypeLedgerEntryRemoved:
		return change.MustRemoved().Type
	case LedgerEntryChangeTypeLedgerEntryUpdated:
		return change.MustUpdated().Data.Type
	case LedgerEntryChangeTypeLedgerEntryState:
		return change.MustState().Data.Type
	default:
		panic(fmt.Errorf("Unknown change type: %v", change.Type))
	}
}
