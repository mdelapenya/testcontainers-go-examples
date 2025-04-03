package modelsrequests

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ModelsRequestBuilder builds and executes requests for operations under \models
type ModelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ModelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ModelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByNamespace gets an item from the modelrunner/sdk.modelsRequests.item collection
// returns a *WithNamespaceItemRequestBuilder when successful
func (m *ModelsRequestBuilder) ByNamespace(namespace string)(*WithNamespaceItemRequestBuilder) {
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
func NewModelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModelsRequestBuilder) {
    m := &ModelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/models", pathParameters),
    }
    return m
}
// NewModelsRequestBuilder instantiates a new ModelsRequestBuilder and sets the default values.
func NewModelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ModelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewModelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Create the create property
// returns a *CreateRequestBuilder when successful
func (m *ModelsRequestBuilder) Create()(*CreateRequestBuilder) {
    return NewCreateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list all models
// returns a []Modelsable when successful
func (m *ModelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ModelsRequestBuilderGetRequestConfiguration)([]Modelsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
func (m *ModelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ModelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ModelsRequestBuilder when successful
func (m *ModelsRequestBuilder) WithUrl(rawUrl string)(*ModelsRequestBuilder) {
    return NewModelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
