package engines

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// V1RequestBuilder builds and executes requests for operations under \engines\v1
type V1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewV1RequestBuilderInternal instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1RequestBuilder) {
    m := &V1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/engines/v1", pathParameters),
    }
    return m
}
// NewV1RequestBuilder instantiates a new V1RequestBuilder and sets the default values.
func NewV1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*V1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewV1RequestBuilderInternal(urlParams, requestAdapter)
}
// Models the models property
// returns a *V1ModelsRequestsModelsRequestBuilder when successful
func (m *V1RequestBuilder) Models()(*V1ModelsRequestsModelsRequestBuilder) {
    return NewV1ModelsRequestsModelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
