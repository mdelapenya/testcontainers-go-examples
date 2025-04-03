package engines

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1ModelsRequestsModelsRequestBuilder builds and executes requests for operations under \engines\v1\models
type V1ModelsRequestsModelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// V1ModelsRequestsModelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type V1ModelsRequestsModelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewV1ModelsRequestsModelsRequestBuilderInternal instantiates a new V1ModelsRequestsModelsRequestBuilder and sets the default values.
func NewV1ModelsRequestsModelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ModelsRequestsModelsRequestBuilder) {
    m := &V1ModelsRequestsModelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/engines/v1/models", pathParameters),
    }
    return m
}
// NewV1ModelsRequestsModelsRequestBuilder instantiates a new V1ModelsRequestsModelsRequestBuilder and sets the default values.
func NewV1ModelsRequestsModelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1ModelsRequestsModelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1ModelsRequestsModelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list available models
// Deprecated: This method is obsolete. Use GetAsModelsGetResponse instead.
// returns a V1ModelsRequestsModelsResponseable when successful
func (m *V1ModelsRequestsModelsRequestBuilder) Get(ctx context.Context, requestConfiguration *V1ModelsRequestsModelsRequestBuilderGetRequestConfiguration)(V1ModelsRequestsModelsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1ModelsRequestsModelsResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1ModelsRequestsModelsResponseable), nil
}
// GetAsModelsGetResponse list available models
// returns a V1ModelsRequestsModelsGetResponseable when successful
func (m *V1ModelsRequestsModelsRequestBuilder) GetAsModelsGetResponse(ctx context.Context, requestConfiguration *V1ModelsRequestsModelsRequestBuilderGetRequestConfiguration)(V1ModelsRequestsModelsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateV1ModelsRequestsModelsGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(V1ModelsRequestsModelsGetResponseable), nil
}
// ToGetRequestInformation list available models
// returns a *RequestInformation when successful
func (m *V1ModelsRequestsModelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *V1ModelsRequestsModelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *V1ModelsRequestsModelsRequestBuilder when successful
func (m *V1ModelsRequestsModelsRequestBuilder) WithUrl(rawUrl string)(*V1ModelsRequestsModelsRequestBuilder) {
    return NewV1ModelsRequestsModelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
