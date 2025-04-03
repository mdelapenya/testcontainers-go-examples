package sdk

import (
	sdkengines "testcontainers-go-examples/docker-model-runner/sdk/engines"
	sdkmodelsrequests "testcontainers-go-examples/docker-model-runner/sdk/modelsrequests"

	sdkabstractions "github.com/microsoft/kiota-abstractions-go"
	sdkserialization "github.com/microsoft/kiota-abstractions-go/serialization"
	sdkserializationform "github.com/microsoft/kiota-serialization-form-go"
	sdkserializationjson "github.com/microsoft/kiota-serialization-json-go"
	sdkserializationmultipart "github.com/microsoft/kiota-serialization-multipart-go"
	sdkserializationtext "github.com/microsoft/kiota-serialization-text-go"
)

// DockerModelRunner the main entry point of the SDK, exposes the configuration and the fluent API.
type DockerModelRunner struct {
	sdkabstractions.BaseRequestBuilder
}

// NewDockerModelRunner instantiates a new DockerModelRunner and sets the default values.
func NewDockerModelRunner(requestAdapter sdkabstractions.RequestAdapter) *DockerModelRunner {
	m := &DockerModelRunner{
		BaseRequestBuilder: *sdkabstractions.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
	}
	sdkabstractions.RegisterDefaultSerializer(func() sdkserialization.SerializationWriterFactory {
		return sdkserializationjson.NewJsonSerializationWriterFactory()
	})
	sdkabstractions.RegisterDefaultSerializer(func() sdkserialization.SerializationWriterFactory {
		return sdkserializationtext.NewTextSerializationWriterFactory()
	})
	sdkabstractions.RegisterDefaultSerializer(func() sdkserialization.SerializationWriterFactory {
		return sdkserializationform.NewFormSerializationWriterFactory()
	})
	sdkabstractions.RegisterDefaultSerializer(func() sdkserialization.SerializationWriterFactory {
		return sdkserializationmultipart.NewMultipartSerializationWriterFactory()
	})
	sdkabstractions.RegisterDefaultDeserializer(func() sdkserialization.ParseNodeFactory {
		return sdkserializationjson.NewJsonParseNodeFactory()
	})
	sdkabstractions.RegisterDefaultDeserializer(func() sdkserialization.ParseNodeFactory {
		return sdkserializationtext.NewTextParseNodeFactory()
	})
	sdkabstractions.RegisterDefaultDeserializer(func() sdkserialization.ParseNodeFactory {
		return sdkserializationform.NewFormParseNodeFactory()
	})
	return m
}

// Engines the engines property
// returns a *EnginesRequestBuilder when successful
func (m *DockerModelRunner) Engines() *sdkengines.EnginesRequestBuilder {
	return sdkengines.NewEnginesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Models the models property
// returns a *ModelsRequestBuilder when successful
func (m *DockerModelRunner) Models() *sdkmodelsrequests.ModelsRequestBuilder {
	return sdkmodelsrequests.NewModelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
