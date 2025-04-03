package engines

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EnginesRequestBuilder builds and executes requests for operations under \engines
type EnginesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewEnginesRequestBuilderInternal instantiates a new EnginesRequestBuilder and sets the default values.
func NewEnginesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnginesRequestBuilder) {
    m := &EnginesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/engines", pathParameters),
    }
    return m
}
// NewEnginesRequestBuilder instantiates a new EnginesRequestBuilder and sets the default values.
func NewEnginesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnginesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnginesRequestBuilderInternal(urlParams, requestAdapter)
}
// V1 the v1 property
// returns a *V1RequestBuilder when successful
func (m *EnginesRequestBuilder) V1()(*V1RequestBuilder) {
    return NewV1RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
