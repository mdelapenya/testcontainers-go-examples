package engines

import (
	sdkserialization "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use V1ModelsRequestsModelsGetResponseable instead.
type V1ModelsRequestsModelsResponse struct {
	V1ModelsRequestsModelsGetResponse
}

// NewV1ModelsRequestsModelsResponse instantiates a new V1ModelsRequestsModelsResponse and sets the default values.
func NewV1ModelsRequestsModelsResponse() *V1ModelsRequestsModelsResponse {
	m := &V1ModelsRequestsModelsResponse{
		V1ModelsRequestsModelsGetResponse: *NewV1ModelsRequestsModelsGetResponse(),
	}
	return m
}

// CreateV1ModelsRequestsModelsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ModelsRequestsModelsResponseFromDiscriminatorValue(parseNode sdkserialization.ParseNode) (sdkserialization.Parsable, error) {
	return NewV1ModelsRequestsModelsResponse(), nil
}

// Deprecated: This class is obsolete. Use V1ModelsRequestsModelsGetResponseable instead.
type V1ModelsRequestsModelsResponseable interface {
	sdkserialization.Parsable
	V1ModelsRequestsModelsGetResponseable
}
