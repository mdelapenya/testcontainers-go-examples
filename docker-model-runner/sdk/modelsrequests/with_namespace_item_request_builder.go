package modelsrequests

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithNamespaceItemRequestBuilder builds and executes requests for operations under \models\{namespace}
type WithNamespaceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByName gets an item from the modelrunner/sdk.modelsRequests.item.item collection
// returns a *ItemWithNameItemRequestBuilder when successful
func (m *WithNamespaceItemRequestBuilder) ByName(name string)(*ItemWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewItemWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithNamespaceItemRequestBuilderInternal instantiates a new WithNamespaceItemRequestBuilder and sets the default values.
func NewWithNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithNamespaceItemRequestBuilder) {
    m := &WithNamespaceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/models/{namespace}", pathParameters),
    }
    return m
}
// NewWithNamespaceItemRequestBuilder instantiates a new WithNamespaceItemRequestBuilder and sets the default values.
func NewWithNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
