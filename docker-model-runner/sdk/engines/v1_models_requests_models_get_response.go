package engines

import (
	sdkserialization "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1ModelsRequestsModelsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The data property
	data []V1ModelsRequestsModelsGetResponse_dataable
	// The object property
	object *string
}

// NewV1ModelsRequestsModelsGetResponse instantiates a new V1ModelsRequestsModelsGetResponse and sets the default values.
func NewV1ModelsRequestsModelsGetResponse() *V1ModelsRequestsModelsGetResponse {
	m := &V1ModelsRequestsModelsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateV1ModelsRequestsModelsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ModelsRequestsModelsGetResponseFromDiscriminatorValue(parseNode sdkserialization.ParseNode) (sdkserialization.Parsable, error) {
	return NewV1ModelsRequestsModelsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1ModelsRequestsModelsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetData gets the data property value. The data property
// returns a []V1ModelsRequestsModelsGetResponse_dataable when successful
func (m *V1ModelsRequestsModelsGetResponse) GetData() []V1ModelsRequestsModelsGetResponse_dataable {
	return m.data
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(sdkserialization.ParseNode)(error) when successful
func (m *V1ModelsRequestsModelsGetResponse) GetFieldDeserializers() map[string]func(sdkserialization.ParseNode) error {
	res := make(map[string]func(sdkserialization.ParseNode) error)
	res["data"] = func(n sdkserialization.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateV1ModelsRequestsModelsGetResponse_dataFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]V1ModelsRequestsModelsGetResponse_dataable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(V1ModelsRequestsModelsGetResponse_dataable)
				}
			}
			m.SetData(res)
		}
		return nil
	}
	res["object"] = func(n sdkserialization.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetObject(val)
		}
		return nil
	}
	return res
}

// GetObject gets the object property value. The object property
// returns a *string when successful
func (m *V1ModelsRequestsModelsGetResponse) GetObject() *string {
	return m.object
}

// Serialize serializes information the current object
func (m *V1ModelsRequestsModelsGetResponse) Serialize(writer sdkserialization.SerializationWriter) error {
	if m.GetData() != nil {
		cast := make([]sdkserialization.Parsable, len(m.GetData()))
		for i, v := range m.GetData() {
			if v != nil {
				cast[i] = v.(sdkserialization.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("data", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("object", m.GetObject())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1ModelsRequestsModelsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetData sets the data property value. The data property
func (m *V1ModelsRequestsModelsGetResponse) SetData(value []V1ModelsRequestsModelsGetResponse_dataable) {
	m.data = value
}

// SetObject sets the object property value. The object property
func (m *V1ModelsRequestsModelsGetResponse) SetObject(value *string) {
	m.object = value
}

type V1ModelsRequestsModelsGetResponseable interface {
	sdkserialization.AdditionalDataHolder
	sdkserialization.Parsable
	GetData() []V1ModelsRequestsModelsGetResponse_dataable
	GetObject() *string
	SetData(value []V1ModelsRequestsModelsGetResponse_dataable)
	SetObject(value *string)
}
