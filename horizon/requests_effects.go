package horizon

// LoadEffects returns list of effects returned by Horizon `/effects` endpoint.
// `params` can be any of `OrderDirection`, `Cursor`, `Limit` or `PageParams`.
// err can be either error object or horizon.Error object.
func (c *Client) LoadEffects(params ...interface{}) (page EffectsPage, err error) {
	c.initHTTPClient()
	rb := &requestBuilder{BaseURL: c.URL}
	url := rb.AddSegment("effects").Params(params...).Build()

	resp, err := c.Client.Get(url)
	if err != nil {
		return
	}

	internalEffectsPage := effectsPage{}
	err = decodeResponse(resp, &internalEffectsPage)
	if err != nil {
		return
	}

	page = internalEffectsPage.ToExported()
	return
}
