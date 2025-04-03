package modelsrequests

import (
	sdkabstractions "github.com/microsoft/kiota-abstractions-go"
)

// WithNamespaceItemRequestBuilder builds and executes requests for operations under \models\{namespace}
type WithNamespaceItemRequestBuilder struct {
	sdkabstractions.BaseRequestBuilder
}

// ByName gets an item from the modelrunner/sdk.modelsRequests.item.item collection
// returns a *ItemWithNameItemRequestBuilder when successful
func (m *WithNamespaceItemRequestBuilder) ByName(name string) *ItemWithNameItemRequestBuilder {
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
func NewWithNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter sdkabstractions.RequestAdapter) *WithNamespaceItemRequestBuilder {
	m := &WithNamespaceItemRequestBuilder{
		BaseRequestBuilder: *sdkabstractions.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/models/{namespace}", pathParameters),
	}
	return m
}

// NewWithNamespaceItemRequestBuilder instantiates a new WithNamespaceItemRequestBuilder and sets the default values.
func NewWithNamespaceItemRequestBuilder(rawUrl string, requestAdapter sdkabstractions.RequestAdapter) *WithNamespaceItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewWithNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
