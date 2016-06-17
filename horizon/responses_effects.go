package horizon

// Effect represents effect returned by horizon.
type Effect struct {
	ID      string `json:"id"`
	PT      string `json:"paging_token"`
	Account string `json:"account"`
	Type    string `json:"type"`
	TypeI   int32  `json:"type_i"`

	*AccountCreated
	*AccountCredited
	*AccountDebited
	*AccountThresholdsUpdated
	*AccountHomeDomainUpdated
	*AccountFlagsUpdated

	*SignerCreated
	*SignerRemoved
	*SignerUpdated

	*TrustlineCreated
	*TrustlineRemoved
	*TrustlineUpdated
	*TrustlineAuthorized
	*TrustlineDeauthorized

	*Trade
}

// effect is an unexported struct used for unmarshaling only.
// It seems to be the most convinient method for end user.
type effect struct {
	ID      string `json:"id"`
	PT      string `json:"paging_token"`
	Account string `json:"account"`
	Type    string `json:"type"`
	TypeI   int32  `json:"type_i"`

	Amount            string `json:"amount"`
	AssetCode         string `json:"asset_code,omitempty"`
	AssetType         string `json:"asset_type"`
	AssetIssuer       string `json:"asset_issuer"`
	AuthRequired      *bool  `json:"auth_required_flag,omitempty"`
	AuthRevokable     *bool  `json:"auth_revokable_flag,omitempty"`
	BoughtAmount      string `json:"bought_amount"`
	BoughtAssetCode   string `json:"bought_asset_code,omitempty"`
	BoughtAssetIssuer string `json:"bought_asset_issuer,omitempty"`
	BoughtAssetType   string `json:"bought_asset_type"`
	HighThreshold     int32  `json:"high_threshold"`
	HomeDomain        string `json:"home_domain"`
	Limit             string `json:"limit"`
	LowThreshold      int32  `json:"low_threshold"`
	MedThreshold      int32  `json:"med_threshold"`
	OfferID           int64  `json:"offer_id"`
	PublicKey         string `json:"public_key"`
	Seller            string `json:"seller"`
	SoldAmount        string `json:"sold_amount"`
	SoldAssetCode     string `json:"sold_asset_code,omitempty"`
	SoldAssetIssuer   string `json:"sold_asset_issuer,omitempty"`
	SoldAssetType     string `json:"sold_asset_type"`
	StartingBalance   string `json:"starting_balance"`
	Trustor           string `json:"trustor"`
	Weight            int32  `json:"weight"`
}

// ToExported creates Effect object from unexported effect.
func (e effect) ToExported() (effect Effect) {
	effect.ID = e.ID
	effect.PT = e.PT
	effect.Account = e.Account
	effect.Type = e.Type
	effect.TypeI = e.TypeI

	switch e.Type {
	case "account_created":
		effect.AccountCreated = &AccountCreated{}
		effect.AccountCreated.StartingBalance = e.StartingBalance
	case "account_credited":
		effect.AccountCredited = &AccountCredited{}
		effect.AccountCredited.Asset = Asset{e.AssetType, e.AssetCode, e.AssetIssuer}
		effect.AccountCredited.Amount = e.Amount
	case "account_debited":
		effect.AccountDebited = &AccountDebited{}
		effect.AccountDebited.Asset = Asset{e.AssetType, e.AssetCode, e.AssetIssuer}
		effect.AccountDebited.Amount = e.Amount
	case "account_thresholds_updated":
		effect.AccountThresholdsUpdated = &AccountThresholdsUpdated{}
		effect.AccountThresholdsUpdated.LowThreshold = e.LowThreshold
		effect.AccountThresholdsUpdated.MedThreshold = e.MedThreshold
		effect.AccountThresholdsUpdated.HighThreshold = e.HighThreshold
	case "account_home_domain_updated":
		effect.AccountHomeDomainUpdated = &AccountHomeDomainUpdated{}
		effect.AccountHomeDomainUpdated.HomeDomain = e.HomeDomain
	case "account_flags_updated":
		effect.AccountFlagsUpdated = &AccountFlagsUpdated{}
		effect.AccountFlagsUpdated.AuthRequired = e.AuthRequired
		effect.AccountFlagsUpdated.AuthRevokable = e.AuthRevokable
	case "signer_created":
		effect.SignerCreated = &SignerCreated{}
		effect.SignerCreated.Weight = e.Weight
		effect.SignerCreated.PublicKey = e.PublicKey
	case "signer_removed":
		effect.SignerRemoved = &SignerRemoved{}
		effect.SignerRemoved.Weight = e.Weight
		effect.SignerRemoved.PublicKey = e.PublicKey
	case "signer_updated":
		effect.SignerUpdated = &SignerUpdated{}
		effect.SignerUpdated.Weight = e.Weight
		effect.SignerUpdated.PublicKey = e.PublicKey
	case "trustline_created":
		effect.TrustlineCreated = &TrustlineCreated{}
		effect.TrustlineCreated.Asset = Asset{e.AssetType, e.AssetCode, e.AssetIssuer}
		effect.TrustlineCreated.Limit = e.Limit
	case "trustline_removed":
		effect.TrustlineRemoved = &TrustlineRemoved{}
		effect.TrustlineRemoved.Asset = Asset{e.AssetType, e.AssetCode, e.AssetIssuer}
		effect.TrustlineRemoved.Limit = e.Limit
	case "trustline_updated":
		effect.TrustlineUpdated = &TrustlineUpdated{}
		effect.TrustlineUpdated.Asset = Asset{e.AssetType, e.AssetCode, e.AssetIssuer}
		effect.TrustlineUpdated.Limit = e.Limit
	case "trustline_authorized":
		effect.TrustlineAuthorized = &TrustlineAuthorized{}
		effect.TrustlineAuthorized.Trustor = e.Trustor
		effect.TrustlineAuthorized.AssetType = e.AssetType
		effect.TrustlineAuthorized.AssetCode = e.AssetCode
	case "trustline_deauthorized":
		effect.TrustlineDeauthorized = &TrustlineDeauthorized{}
		effect.TrustlineDeauthorized.Trustor = e.Trustor
		effect.TrustlineDeauthorized.AssetType = e.AssetType
		effect.TrustlineDeauthorized.AssetCode = e.AssetCode
	case "trade":
		effect.Trade = &Trade{}
		effect.Trade.Seller = e.Seller
		effect.Trade.OfferID = e.OfferID
		effect.Trade.SoldAmount = e.SoldAmount
		effect.Trade.SoldAssetType = e.SoldAssetType
		effect.Trade.SoldAssetCode = e.SoldAssetCode
		effect.Trade.SoldAssetIssuer = e.SoldAssetIssuer
		effect.Trade.BoughtAmount = e.BoughtAmount
		effect.Trade.BoughtAssetType = e.BoughtAssetType
		effect.Trade.BoughtAssetCode = e.BoughtAssetCode
		effect.Trade.BoughtAssetIssuer = e.BoughtAssetIssuer
	}

	return
}

type EffectsPage struct {
	Links    `json:"_links"`
	Embedded struct {
		Records []Effect `json:"records"`
	} `json:"_embedded"`
}

// effectsPage is an unexported struct used for unmarshaling only.
// It seems to be the most convinient method for end user.
type effectsPage struct {
	Links    `json:"_links"`
	Embedded struct {
		Records []effect `json:"records"`
	} `json:"_embedded"`
}

// ToExported creates Effect objects from unexported effect.
func (e effectsPage) ToExported() (page EffectsPage) {
	page.Links = e.Links
	for _, effect := range e.Embedded.Records {
		page.Embedded.Records = append(page.Embedded.Records, effect.ToExported())
	}
	return
}

type AccountCreated struct {
	StartingBalance string `json:"starting_balance"`
}

type AccountCredited struct {
	Asset
	Amount string `json:"amount"`
}

type AccountDebited struct {
	Asset
	Amount string `json:"amount"`
}

type AccountThresholdsUpdated struct {
	LowThreshold  int32 `json:"low_threshold"`
	MedThreshold  int32 `json:"med_threshold"`
	HighThreshold int32 `json:"high_threshold"`
}

type AccountHomeDomainUpdated struct {
	HomeDomain string `json:"home_domain"`
}

type AccountFlagsUpdated struct {
	AuthRequired  *bool `json:"auth_required_flag,omitempty"`
	AuthRevokable *bool `json:"auth_revokable_flag,omitempty"`
}

type SignerCreated struct {
	Weight    int32  `json:"weight"`
	PublicKey string `json:"public_key"`
}

type SignerRemoved struct {
	Weight    int32  `json:"weight"`
	PublicKey string `json:"public_key"`
}

type SignerUpdated struct {
	Weight    int32  `json:"weight"`
	PublicKey string `json:"public_key"`
}

type TrustlineCreated struct {
	Asset
	Limit string `json:"limit"`
}

type TrustlineRemoved struct {
	Asset
	Limit string `json:"limit"`
}

type TrustlineUpdated struct {
	Asset
	Limit string `json:"limit"`
}

type TrustlineAuthorized struct {
	Trustor   string `json:"trustor"`
	AssetType string `json:"asset_type"`
	AssetCode string `json:"asset_code,omitempty"`
}

type TrustlineDeauthorized struct {
	Trustor   string `json:"trustor"`
	AssetType string `json:"asset_type"`
	AssetCode string `json:"asset_code,omitempty"`
}

type Trade struct {
	Seller            string `json:"seller"`
	OfferID           int64  `json:"offer_id"`
	SoldAmount        string `json:"sold_amount"`
	SoldAssetType     string `json:"sold_asset_type"`
	SoldAssetCode     string `json:"sold_asset_code,omitempty"`
	SoldAssetIssuer   string `json:"sold_asset_issuer,omitempty"`
	BoughtAmount      string `json:"bought_amount"`
	BoughtAssetType   string `json:"bought_asset_type"`
	BoughtAssetCode   string `json:"bought_asset_code,omitempty"`
	BoughtAssetIssuer string `json:"bought_asset_issuer,omitempty"`
}
