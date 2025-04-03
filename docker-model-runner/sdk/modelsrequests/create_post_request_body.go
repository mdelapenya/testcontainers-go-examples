package modelsrequests

import (
	sdkserialization "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreatePostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The from property
	from *string
}

// NewCreatePostRequestBody instantiates a new CreatePostRequestBody and sets the default values.
func NewCreatePostRequestBody() *CreatePostRequestBody {
	m := &CreatePostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCreatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreatePostRequestBodyFromDiscriminatorValue(parseNode sdkserialization.ParseNode) (sdkserialization.Parsable, error) {
	return NewCreatePostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreatePostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(sdkserialization.ParseNode)(error) when successful
func (m *CreatePostRequestBody) GetFieldDeserializers() map[string]func(sdkserialization.ParseNode) error {
	res := make(map[string]func(sdkserialization.ParseNode) error)
	res["from"] = func(n sdkserialization.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFrom(val)
		}
		return nil
	}
	return res
}

// GetFrom gets the from property value. The from property
// returns a *string when successful
func (m *CreatePostRequestBody) GetFrom() *string {
	return m.from
}

// Serialize serializes information the current object
func (m *CreatePostRequestBody) Serialize(writer sdkserialization.SerializationWriter) error {
	{
		err := writer.WriteStringValue("from", m.GetFrom())
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
func (m *CreatePostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFrom sets the from property value. The from property
func (m *CreatePostRequestBody) SetFrom(value *string) {
	m.from = value
}

type CreatePostRequestBodyable interface {
	sdkserialization.AdditionalDataHolder
	sdkserialization.Parsable
	GetFrom() *string
	SetFrom(value *string)
}
