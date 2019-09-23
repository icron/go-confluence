package confluence

func (api *API) CreatePage(title, spaceKey, value string) (*ContentCollection, error) {
	result := &ContentCollection{}
	req := Content{
		Type:  "page",
		Title: title,
		Space: &Space{Key: spaceKey},
		Body: &Body{
			StorageView: &View{Value: value, Representation: "storage"},
		},
	}
	
	statusCode, err := api.doRequest(
		"POST", "/rest/api/content",
		nil, result, req,
	)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result, nil
	case 403:
		return nil, ErrNoPerms
	case 404:
		return nil, ErrNoContent
	default:
		return nil, makeUnknownError(statusCode)
	}
}

func (api *API) UpdatePage(contentID, title, value string) (*ContentCollection, error) {
	result := &ContentCollection{}
	content, err := api.GetContentByID(contentID, ContentIDParameters{})
	if err != nil {
		return nil, err
	}

	if title != "" {
		content.Title = title
	}
	if value != "" {
		content.Body.StorageView.Value = value
	}
	statusCode, err := api.doRequest(
		"PUT", "/rest/api/content",
		nil, result, content,
	)

	if err != nil {
		return nil, err
	}

	switch statusCode {
	case 200:
		return result, nil
	case 403:
		return nil, ErrNoPerms
	case 404:
		return nil, ErrNoContent
	default:
		return nil, makeUnknownError(statusCode)
	}
}
