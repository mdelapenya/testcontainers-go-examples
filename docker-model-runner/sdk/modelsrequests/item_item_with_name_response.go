package modelsrequests

import (
	sdkserialization "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemWithNameGetResponseable instead.
type ItemItemWithNameResponse struct {
	ItemItemWithNameGetResponse
}

// NewItemItemWithNameResponse instantiates a new ItemItemWithNameResponse and sets the default values.
func NewItemItemWithNameResponse() *ItemItemWithNameResponse {
	m := &ItemItemWithNameResponse{
		ItemItemWithNameGetResponse: *NewItemItemWithNameGetResponse(),
	}
	return m
}

// CreateItemItemWithNameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemWithNameResponseFromDiscriminatorValue(parseNode sdkserialization.ParseNode) (sdkserialization.Parsable, error) {
	return NewItemItemWithNameResponse(), nil
}

// Deprecated: This class is obsolete. Use ItemItemWithNameGetResponseable instead.
type ItemItemWithNameResponseable interface {
	ItemItemWithNameGetResponseable
	sdkserialization.Parsable
}
