package engines

import (
	sdkabstractions "github.com/microsoft/kiota-abstractions-go"
)

// EnginesRequestBuilder builds and executes requests for operations under \engines
type EnginesRequestBuilder struct {
	sdkabstractions.BaseRequestBuilder
}

// NewEnginesRequestBuilderInternal instantiates a new EnginesRequestBuilder and sets the default values.
func NewEnginesRequestBuilderInternal(pathParameters map[string]string, requestAdapter sdkabstractions.RequestAdapter) *EnginesRequestBuilder {
	m := &EnginesRequestBuilder{
		BaseRequestBuilder: *sdkabstractions.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/engines", pathParameters),
	}
	return m
}

// NewEnginesRequestBuilder instantiates a new EnginesRequestBuilder and sets the default values.
func NewEnginesRequestBuilder(rawUrl string, requestAdapter sdkabstractions.RequestAdapter) *EnginesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEnginesRequestBuilderInternal(urlParams, requestAdapter)
}

// V1 the v1 property
// returns a *V1RequestBuilder when successful
func (m *EnginesRequestBuilder) V1() *V1RequestBuilder {
	return NewV1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
