package modelsrequests

import (
	"context"

	sdkabstractions "github.com/microsoft/kiota-abstractions-go"
)

// ModelsRequestBuilder builds and executes requests for operations under \models
type ModelsRequestBuilder struct {
	sdkabstractions.BaseRequestBuilder
}

// ModelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ModelsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *sdkabstractions.RequestHeaders
	// Request options
	Options []sdkabstractions.RequestOption
}

// ByNamespace gets an item from the modelrunner/sdk.modelsRequests.item collection
// returns a *WithNamespaceItemRequestBuilder when successful
func (m *ModelsRequestBuilder) ByNamespace(namespace string) *WithNamespaceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if namespace != "" {
		urlTplParams["namespace"] = namespace
	}
	return NewWithNamespaceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewModelsRequestBuilderInternal instantiates a new ModelsRequestBuilder and sets the default values.
func NewModelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter sdkabstractions.RequestAdapter) *ModelsRequestBuilder {
	m := &ModelsRequestBuilder{
		BaseRequestBuilder: *sdkabstractions.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/models", pathParameters),
	}
	return m
}

// NewModelsRequestBuilder instantiates a new ModelsRequestBuilder and sets the default values.
func NewModelsRequestBuilder(rawUrl string, requestAdapter sdkabstractions.RequestAdapter) *ModelsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewModelsRequestBuilderInternal(urlParams, requestAdapter)
}

// Create the create property
// returns a *CreateRequestBuilder when successful
func (m *ModelsRequestBuilder) Create() *CreateRequestBuilder {
	return NewCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get list all models
// returns a []Modelsable when successful
func (m *ModelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ModelsRequestBuilderGetRequestConfiguration) ([]Modelsable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateModelsFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	val := make([]Modelsable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(Modelsable)
		}
	}
	return val, nil
}

// ToGetRequestInformation list all models
// returns a *RequestInformation when successful
func (m *ModelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ModelsRequestBuilderGetRequestConfiguration) (*sdkabstractions.RequestInformation, error) {
	requestInfo := sdkabstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(sdkabstractions.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ModelsRequestBuilder when successful
func (m *ModelsRequestBuilder) WithUrl(rawUrl string) *ModelsRequestBuilder {
	return NewModelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
