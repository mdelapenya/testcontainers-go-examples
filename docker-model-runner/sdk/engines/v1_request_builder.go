package engines

import (
	sdkabstractions "github.com/microsoft/kiota-abstractions-go"
)

// V1RequestBuilder builds and executes requests for operations under \engines\v1
type V1RequestBuilder struct {
	sdkabstractions.BaseRequestBuilder
}

// NewV1RequestBuilderInternal instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilderInternal(pathParameters map[string]string, requestAdapter sdkabstractions.RequestAdapter) *V1RequestBuilder {
	m := &V1RequestBuilder{
		BaseRequestBuilder: *sdkabstractions.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/engines/v1", pathParameters),
	}
	return m
}

// NewV1RequestBuilder instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilder(rawUrl string, requestAdapter sdkabstractions.RequestAdapter) *V1RequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewV1RequestBuilderInternal(urlParams, requestAdapter)
}

// Models the models property
// returns a *V1ModelsRequestsModelsRequestBuilder when successful
func (m *V1RequestBuilder) Models() *V1ModelsRequestsModelsRequestBuilder {
	return NewV1ModelsRequestsModelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
