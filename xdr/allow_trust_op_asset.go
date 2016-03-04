package xdr

import (
	"fmt"
)

// ToAsset converts `a` to a proper xdr.Asset
func (a AllowTrustOpAsset) ToAsset() (ret Asset) {
	var err error

	switch a.Type {
	case AssetTypeAssetTypeCreditAlphanum4:
		ret, err = NewAsset(AssetTypeAssetTypeCreditAlphanum4, a.MustAssetCode4())
	case AssetTypeAssetTypeCreditAlphanum12:
		ret, err = NewAsset(AssetTypeAssetTypeCreditAlphanum12, a.MustAssetCode12())
	default:
		err = fmt.Errorf("Unexpected type for AllowTrustOpAsset: %d", a.Type)
	}

	if err != nil {
		panic(err)
	}
	return
}
