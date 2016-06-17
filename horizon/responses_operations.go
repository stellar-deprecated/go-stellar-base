package horizon

type Operation struct {
	// Links struct {
	// 	Self        hal.Link `json:"self"`
	// 	Transaction hal.Link `json:"transaction"`
	// 	Effects     hal.Link `json:"effects"`
	// 	Succeeds    hal.Link `json:"succeeds"`
	// 	Precedes    hal.Link `json:"precedes"`
	// } `json:"_links"`

	ID            string `json:"id"`
	PT            string `json:"paging_token"`
	SourceAccount string `json:"source_account"`
	Type          string `json:"type"`
	TypeI         int32  `json:"type_i"`

	*CreateAccount
	*Payment
	*PathPayment
	*ManageData
	*ManageOffer
	*CreatePassiveOffer
	*SetOptions
	*ChangeTrust
	*AllowTrust
	*AccountMerge
	*Inflation
}

// operation is an unexported struct used for unmarshaling only.
// It seems to be the most convinient method for end user.
type operation struct {
	ID            string `json:"id"`
	PT            string `json:"paging_token"`
	SourceAccount string `json:"source_account"`
	Type          string `json:"type"`
	TypeI         int32  `json:"type_i"`

	Asset
	StartingBalance    string   `json:"starting_balance"`
	Funder             string   `json:"funder"`
	Account            string   `json:"account"`
	From               string   `json:"from"`
	To                 string   `json:"to"`
	Amount             string   `json:"amount"`
	Path               []Asset  `json:"path"`
	SourceMax          string   `json:"source_max"`
	SourceAssetType    string   `json:"source_asset_type"`
	SourceAssetCode    string   `json:"source_asset_code,omitempty"`
	SourceAssetIssuer  string   `json:"source_asset_issuer,omitempty"`
	Name               string   `json:"name"`
	Value              string   `json:"value"`
	OfferID            int64    `json:"offer_id"`
	Price              string   `json:"price"`
	PriceR             Price    `json:"price"`
	BuyingAssetType    string   `json:"buying_asset_type"`
	BuyingAssetCode    string   `json:"buying_asset_code,omitempty"`
	BuyingAssetIssuer  string   `json:"buying_asset_issuer,omitempty"`
	SellingAssetType   string   `json:"selling_asset_type"`
	SellingAssetCode   string   `json:"selling_asset_code,omitempty"`
	SellingAssetIssuer string   `json:"selling_asset_issuer,omitempty"`
	HomeDomain         string   `json:"home_domain,omitempty"`
	InflationDest      string   `json:"inflation_dest,omitempty"`
	MasterKeyWeight    *int     `json:"master_key_weight,omitempty"`
	SignerKey          string   `json:"signer_key,omitempty"`
	SignerWeight       *int     `json:"signer_weight,omitempty"`
	SetFlags           []int    `json:"set_flags,omitempty"`
	SetFlagsS          []string `json:"set_flags_s,omitempty"`
	ClearFlags         []int    `json:"clear_flags,omitempty"`
	ClearFlagsS        []string `json:"clear_flags_s,omitempty"`
	LowThreshold       *int     `json:"low_threshold,omitempty"`
	MedThreshold       *int     `json:"med_threshold,omitempty"`
	HighThreshold      *int     `json:"high_threshold,omitempty"`
	Limit              string   `json:"limit"`
	Trustee            string   `json:"trustee"`
	Trustor            string   `json:"trustor"`
	Authorize          bool     `json:"authorize"`
	Into               string   `json:"into"`
}

// ToExported creates Operation object from unexported operation.
func (o operation) ToExported() (operation Operation) {
	operation.ID = o.ID
	operation.PT = o.PT
	operation.SourceAccount = o.SourceAccount
	operation.Type = o.Type
	operation.TypeI = o.TypeI

	switch o.Type {
	case "create_account":
		operation.CreateAccount = &CreateAccount{}
		operation.CreateAccount.StartingBalance = o.StartingBalance
		operation.CreateAccount.Funder = o.Funder
		operation.CreateAccount.Account = o.Account
	case "payment":
		operation.Payment = &Payment{}
		operation.Payment.Asset = o.Asset
		operation.Payment.From = o.From
		operation.Payment.To = o.To
		operation.Payment.Amount = o.Amount
	case "path_payment":
		operation.PathPayment = &PathPayment{}
		operation.PathPayment.Asset = o.Asset
		operation.PathPayment.From = o.From
		operation.PathPayment.To = o.To
		operation.PathPayment.Amount = o.Amount
		operation.PathPayment.Path = o.Path
		operation.PathPayment.SourceMax = o.SourceMax
		operation.PathPayment.SourceAssetType = o.SourceAssetType
		operation.PathPayment.SourceAssetCode = o.SourceAssetCode
		operation.PathPayment.SourceAssetIssuer = o.SourceAssetIssuer
	case "manage_data":
		operation.ManageData = &ManageData{}
		operation.ManageData.Name = o.Name
		operation.ManageData.Value = o.Value
	case "manage_offer":
		operation.ManageOffer = &ManageOffer{}
		operation.ManageOffer.OfferID = o.OfferID
		operation.ManageOffer.Amount = o.Amount
		operation.ManageOffer.Price = o.Price
		operation.ManageOffer.PriceR = o.PriceR
		operation.ManageOffer.BuyingAssetType = o.BuyingAssetType
		operation.ManageOffer.BuyingAssetCode = o.BuyingAssetCode
		operation.ManageOffer.BuyingAssetIssuer = o.BuyingAssetIssuer
		operation.ManageOffer.SellingAssetType = o.SellingAssetType
		operation.ManageOffer.SellingAssetCode = o.SellingAssetCode
		operation.ManageOffer.SellingAssetIssuer = o.SellingAssetIssuer
	case "create_passive_offer":
		operation.CreatePassiveOffer = &CreatePassiveOffer{}
		operation.CreatePassiveOffer.OfferID = o.OfferID
		operation.CreatePassiveOffer.Amount = o.Amount
		operation.CreatePassiveOffer.Price = o.Price
		operation.CreatePassiveOffer.PriceR = o.PriceR
		operation.CreatePassiveOffer.BuyingAssetType = o.BuyingAssetType
		operation.CreatePassiveOffer.BuyingAssetCode = o.BuyingAssetCode
		operation.CreatePassiveOffer.BuyingAssetIssuer = o.BuyingAssetIssuer
		operation.CreatePassiveOffer.SellingAssetType = o.SellingAssetType
		operation.CreatePassiveOffer.SellingAssetCode = o.SellingAssetCode
		operation.CreatePassiveOffer.SellingAssetIssuer = o.SellingAssetIssuer
	case "set_options":
		operation.SetOptions = &SetOptions{}
		operation.SetOptions.HomeDomain = o.HomeDomain
		operation.SetOptions.InflationDest = o.InflationDest
		operation.SetOptions.MasterKeyWeight = o.MasterKeyWeight
		operation.SetOptions.SignerKey = o.SignerKey
		operation.SetOptions.SignerWeight = o.SignerWeight
		operation.SetOptions.SetFlags = o.SetFlags
		operation.SetOptions.SetFlagsS = o.SetFlagsS
		operation.SetOptions.ClearFlags = o.ClearFlags
		operation.SetOptions.ClearFlagsS = o.ClearFlagsS
		operation.SetOptions.LowThreshold = o.LowThreshold
		operation.SetOptions.MedThreshold = o.MedThreshold
		operation.SetOptions.HighThreshold = o.HighThreshold
	case "change_trust":
		operation.ChangeTrust = &ChangeTrust{}
		operation.ChangeTrust.Asset = o.Asset
		operation.ChangeTrust.Limit = o.Limit
		operation.ChangeTrust.Trustee = o.Trustee
		operation.ChangeTrust.Trustor = o.Trustor
	case "allow_trust":
		operation.AllowTrust = &AllowTrust{}
		operation.AllowTrust.Asset = o.Asset
		operation.AllowTrust.Authorize = o.Authorize
		operation.AllowTrust.Trustee = o.Trustee
		operation.AllowTrust.Trustor = o.Trustor
	case "account_merge":
		operation.AccountMerge = &AccountMerge{}
		operation.AccountMerge.Account = o.Account
		operation.AccountMerge.Into = o.Into
	case "inflation":
		operation.Inflation = &Inflation{}
	}

	return
}

type OperationsPage struct {
	Links    `json:"_links"`
	Embedded struct {
		Records []Operation `json:"records"`
	} `json:"_embedded"`
}

// operationsPage is an unexported struct used for unmarshaling only.
// It seems to be the most convinient method for end user.
type operationsPage struct {
	Links    `json:"_links"`
	Embedded struct {
		Records []operation `json:"records"`
	} `json:"_embedded"`
}

// ToExported creates Effect objects from unexported effect.
func (e operationsPage) ToExported() (page OperationsPage) {
	page.Links = e.Links
	for _, operation := range e.Embedded.Records {
		page.Embedded.Records = append(page.Embedded.Records, operation.ToExported())
	}
	return
}

type CreateAccount struct {
	StartingBalance string `json:"starting_balance"`
	Funder          string `json:"funder"`
	Account         string `json:"account"`
}

type Payment struct {
	Asset
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

type PathPayment struct {
	Payment
	Path              []Asset `json:"path"`
	SourceMax         string  `json:"source_max"`
	SourceAssetType   string  `json:"source_asset_type"`
	SourceAssetCode   string  `json:"source_asset_code,omitempty"`
	SourceAssetIssuer string  `json:"source_asset_issuer,omitempty"`
}

// ManageData represents a ManageData operation as it is serialized into json
// for the horizon API.
type ManageData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ManageOffer struct {
	OfferID            int64  `json:"offer_id"`
	Amount             string `json:"amount"`
	Price              string `json:"price"`
	PriceR             Price  `json:"price"`
	BuyingAssetType    string `json:"buying_asset_type"`
	BuyingAssetCode    string `json:"buying_asset_code,omitempty"`
	BuyingAssetIssuer  string `json:"buying_asset_issuer,omitempty"`
	SellingAssetType   string `json:"selling_asset_type"`
	SellingAssetCode   string `json:"selling_asset_code,omitempty"`
	SellingAssetIssuer string `json:"selling_asset_issuer,omitempty"`
}

type CreatePassiveOffer struct {
	ManageOffer
}

type SetOptions struct {
	HomeDomain    string `json:"home_domain,omitempty"`
	InflationDest string `json:"inflation_dest,omitempty"`

	MasterKeyWeight *int   `json:"master_key_weight,omitempty"`
	SignerKey       string `json:"signer_key,omitempty"`
	SignerWeight    *int   `json:"signer_weight,omitempty"`

	SetFlags    []int    `json:"set_flags,omitempty"`
	SetFlagsS   []string `json:"set_flags_s,omitempty"`
	ClearFlags  []int    `json:"clear_flags,omitempty"`
	ClearFlagsS []string `json:"clear_flags_s,omitempty"`

	LowThreshold  *int `json:"low_threshold,omitempty"`
	MedThreshold  *int `json:"med_threshold,omitempty"`
	HighThreshold *int `json:"high_threshold,omitempty"`
}

type ChangeTrust struct {
	Asset
	Limit   string `json:"limit"`
	Trustee string `json:"trustee"`
	Trustor string `json:"trustor"`
}

type AllowTrust struct {
	Asset
	Trustee   string `json:"trustee"`
	Trustor   string `json:"trustor"`
	Authorize bool   `json:"authorize"`
}

type AccountMerge struct {
	Account string `json:"account"`
	Into    string `json:"into"`
}

type Inflation struct {
}
