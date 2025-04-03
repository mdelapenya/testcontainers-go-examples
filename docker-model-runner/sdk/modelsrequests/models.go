package modelsrequests

import (
	sdkserialization "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Models struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The created property
	created *int64
	// The files property
	files []string
	// The id property
	id *string
	// The tags property
	tags []string
}

// NewModels instantiates a new Models and sets the default values.
func NewModels() *Models {
	m := &Models{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateModelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateModelsFromDiscriminatorValue(parseNode sdkserialization.ParseNode) (sdkserialization.Parsable, error) {
	return NewModels(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Models) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCreated gets the created property value. The created property
// returns a *int64 when successful
func (m *Models) GetCreated() *int64 {
	return m.created
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(sdkserialization.ParseNode)(error) when successful
func (m *Models) GetFieldDeserializers() map[string]func(sdkserialization.ParseNode) error {
	res := make(map[string]func(sdkserialization.ParseNode) error)
	res["created"] = func(n sdkserialization.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreated(val)
		}
		return nil
	}
	res["files"] = func(n sdkserialization.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetFiles(res)
		}
		return nil
	}
	res["id"] = func(n sdkserialization.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["tags"] = func(n sdkserialization.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetTags(res)
		}
		return nil
	}
	return res
}

// GetFiles gets the files property value. The files property
// returns a []string when successful
func (m *Models) GetFiles() []string {
	return m.files
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *Models) GetId() *string {
	return m.id
}

// GetTags gets the tags property value. The tags property
// returns a []string when successful
func (m *Models) GetTags() []string {
	return m.tags
}

// Serialize serializes information the current object
func (m *Models) Serialize(writer sdkserialization.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("created", m.GetCreated())
		if err != nil {
			return err
		}
	}
	if m.GetFiles() != nil {
		err := writer.WriteCollectionOfStringValues("files", m.GetFiles())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	if m.GetTags() != nil {
		err := writer.WriteCollectionOfStringValues("tags", m.GetTags())
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
func (m *Models) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCreated sets the created property value. The created property
func (m *Models) SetCreated(value *int64) {
	m.created = value
}

// SetFiles sets the files property value. The files property
func (m *Models) SetFiles(value []string) {
	m.files = value
}

// SetId sets the id property value. The id property
func (m *Models) SetId(value *string) {
	m.id = value
}

// SetTags sets the tags property value. The tags property
func (m *Models) SetTags(value []string) {
	m.tags = value
}

type Modelsable interface {
	sdkserialization.AdditionalDataHolder
	sdkserialization.Parsable
	GetCreated() *int64
	GetFiles() []string
	GetId() *string
	GetTags() []string
	SetCreated(value *int64)
	SetFiles(value []string)
	SetId(value *string)
	SetTags(value []string)
}
