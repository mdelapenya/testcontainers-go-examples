package engines

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1ModelsRequestsModelsGetResponse_data struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The created property
    created *int64
    // The id property
    id *string
    // The object property
    object *string
    // The owned_by property
    owned_by *string
}
// NewV1ModelsRequestsModelsGetResponse_data instantiates a new V1ModelsRequestsModelsGetResponse_data and sets the default values.
func NewV1ModelsRequestsModelsGetResponse_data()(*V1ModelsRequestsModelsGetResponse_data) {
    m := &V1ModelsRequestsModelsGetResponse_data{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1ModelsRequestsModelsGetResponse_dataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1ModelsRequestsModelsGetResponse_dataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1ModelsRequestsModelsGetResponse_data(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1ModelsRequestsModelsGetResponse_data) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreated gets the created property value. The created property
// returns a *int64 when successful
func (m *V1ModelsRequestsModelsGetResponse_data) GetCreated()(*int64) {
    return m.created
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1ModelsRequestsModelsGetResponse_data) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreated(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["object"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObject(val)
        }
        return nil
    }
    res["owned_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnedBy(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *string when successful
func (m *V1ModelsRequestsModelsGetResponse_data) GetId()(*string) {
    return m.id
}
// GetObject gets the object property value. The object property
// returns a *string when successful
func (m *V1ModelsRequestsModelsGetResponse_data) GetObject()(*string) {
    return m.object
}
// GetOwnedBy gets the owned_by property value. The owned_by property
// returns a *string when successful
func (m *V1ModelsRequestsModelsGetResponse_data) GetOwnedBy()(*string) {
    return m.owned_by
}
// Serialize serializes information the current object
func (m *V1ModelsRequestsModelsGetResponse_data) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("created", m.GetCreated())
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
    {
        err := writer.WriteStringValue("object", m.GetObject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("owned_by", m.GetOwnedBy())
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
func (m *V1ModelsRequestsModelsGetResponse_data) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreated sets the created property value. The created property
func (m *V1ModelsRequestsModelsGetResponse_data) SetCreated(value *int64)() {
    m.created = value
}
// SetId sets the id property value. The id property
func (m *V1ModelsRequestsModelsGetResponse_data) SetId(value *string)() {
    m.id = value
}
// SetObject sets the object property value. The object property
func (m *V1ModelsRequestsModelsGetResponse_data) SetObject(value *string)() {
    m.object = value
}
// SetOwnedBy sets the owned_by property value. The owned_by property
func (m *V1ModelsRequestsModelsGetResponse_data) SetOwnedBy(value *string)() {
    m.owned_by = value
}
type V1ModelsRequestsModelsGetResponse_dataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreated()(*int64)
    GetId()(*string)
    GetObject()(*string)
    GetOwnedBy()(*string)
    SetCreated(value *int64)()
    SetId(value *string)()
    SetObject(value *string)()
    SetOwnedBy(value *string)()
}
