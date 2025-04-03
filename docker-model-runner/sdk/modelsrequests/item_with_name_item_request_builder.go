package modelsrequests

import (
	"context"

	sdkabstractions "github.com/microsoft/kiota-abstractions-go"
)

// ItemWithNameItemRequestBuilder builds and executes requests for operations under \models\{namespace}\{name}
type ItemWithNameItemRequestBuilder struct {
	sdkabstractions.BaseRequestBuilder
}

// ItemWithNameItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithNameItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *sdkabstractions.RequestHeaders
	// Request options
	Options []sdkabstractions.RequestOption
}

// ItemWithNameItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWithNameItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *sdkabstractions.RequestHeaders
	// Request options
	Options []sdkabstractions.RequestOption
}

// NewItemWithNameItemRequestBuilderInternal instantiates a new ItemWithNameItemRequestBuilder and sets the default values.
func NewItemWithNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter sdkabstractions.RequestAdapter) *ItemWithNameItemRequestBuilder {
	m := &ItemWithNameItemRequestBuilder{
		BaseRequestBuilder: *sdkabstractions.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/models/{namespace}/{name}", pathParameters),
	}
	return m
}

// NewItemWithNameItemRequestBuilder instantiates a new ItemWithNameItemRequestBuilder and sets the default values.
func NewItemWithNameItemRequestBuilder(rawUrl string, requestAdapter sdkabstractions.RequestAdapter) *ItemWithNameItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemWithNameItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete a model
// returns a []byte when successful
func (m *ItemWithNameItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemWithNameItemRequestBuilderDeleteRequestConfiguration) ([]byte, error) {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.([]byte), nil
}

// Get get model details
// Deprecated: This method is obsolete. Use GetAsWithNameGetResponse instead.
// returns a ItemItemWithNameResponseable when successful
func (m *ItemWithNameItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemWithNameItemRequestBuilderGetRequestConfiguration) (ItemItemWithNameResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemWithNameResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemItemWithNameResponseable), nil
}

// GetAsWithNameGetResponse get model details
// returns a ItemItemWithNameGetResponseable when successful
func (m *ItemWithNameItemRequestBuilder) GetAsWithNameGetResponse(ctx context.Context, requestConfiguration *ItemWithNameItemRequestBuilderGetRequestConfiguration) (ItemItemWithNameGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemWithNameGetResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemItemWithNameGetResponseable), nil
}

// ToDeleteRequestInformation delete a model
// returns a *RequestInformation when successful
func (m *ItemWithNameItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemWithNameItemRequestBuilderDeleteRequestConfiguration) (*sdkabstractions.RequestInformation, error) {
	requestInfo := sdkabstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(sdkabstractions.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToGetRequestInformation get model details
// returns a *RequestInformation when successful
func (m *ItemWithNameItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemWithNameItemRequestBuilderGetRequestConfiguration) (*sdkabstractions.RequestInformation, error) {
	requestInfo := sdkabstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(sdkabstractions.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemWithNameItemRequestBuilder when successful
func (m *ItemWithNameItemRequestBuilder) WithUrl(rawUrl string) *ItemWithNameItemRequestBuilder {
	return NewItemWithNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
