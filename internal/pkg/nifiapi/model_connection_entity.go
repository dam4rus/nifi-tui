/*
NiFi Rest API

The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.

API version: 2.0.0
Contact: dev@nifi.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifiapi

import (
	"encoding/json"
)

// checks if the ConnectionEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionEntity{}

// ConnectionEntity struct for ConnectionEntity
type ConnectionEntity struct {
	Revision *RevisionDTO `json:"revision,omitempty"`
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri *string `json:"uri,omitempty"`
	Position *PositionDTO `json:"position,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
	Component *ConnectionDTO `json:"component,omitempty"`
	Status *ConnectionStatusDTO `json:"status,omitempty"`
	// The bend points on the connection.
	Bends []PositionDTO `json:"bends,omitempty"`
	// The index of the bend point where to place the connection label.
	LabelIndex *int32 `json:"labelIndex,omitempty"`
	// The z index of the connection.
	GetzIndex *int64 `json:"getzIndex,omitempty"`
	// The identifier of the source of this connection.
	SourceId *string `json:"sourceId,omitempty"`
	// The identifier of the group of the source of this connection.
	SourceGroupId *string `json:"sourceGroupId,omitempty"`
	// The type of component the source connectable is.
	SourceType string `json:"sourceType"`
	// The identifier of the destination of this connection.
	DestinationId *string `json:"destinationId,omitempty"`
	// The identifier of the group of the destination of this connection.
	DestinationGroupId *string `json:"destinationGroupId,omitempty"`
	// The type of component the destination connectable is.
	DestinationType string `json:"destinationType"`
}

// NewConnectionEntity instantiates a new ConnectionEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionEntity(sourceType string, destinationType string) *ConnectionEntity {
	this := ConnectionEntity{}
	this.SourceType = sourceType
	this.DestinationType = destinationType
	return &this
}

// NewConnectionEntityWithDefaults instantiates a new ConnectionEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionEntityWithDefaults() *ConnectionEntity {
	this := ConnectionEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ConnectionEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ConnectionEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *ConnectionEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ConnectionEntity) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ConnectionEntity) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ConnectionEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ConnectionEntity) GetPosition() PositionDTO {
	if o == nil || IsNil(o.Position) {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ConnectionEntity) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *ConnectionEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ConnectionEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ConnectionEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ConnectionEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *ConnectionEntity) GetBulletins() []BulletinEntity {
	if o == nil || IsNil(o.Bulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.Bulletins) {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *ConnectionEntity) HasBulletins() bool {
	if o != nil && !IsNil(o.Bulletins) {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *ConnectionEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ConnectionEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ConnectionEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ConnectionEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ConnectionEntity) GetComponent() ConnectionDTO {
	if o == nil || IsNil(o.Component) {
		var ret ConnectionDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetComponentOk() (*ConnectionDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ConnectionEntity) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ConnectionDTO and assigns it to the Component field.
func (o *ConnectionEntity) SetComponent(v ConnectionDTO) {
	o.Component = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConnectionEntity) GetStatus() ConnectionStatusDTO {
	if o == nil || IsNil(o.Status) {
		var ret ConnectionStatusDTO
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetStatusOk() (*ConnectionStatusDTO, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectionEntity) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ConnectionStatusDTO and assigns it to the Status field.
func (o *ConnectionEntity) SetStatus(v ConnectionStatusDTO) {
	o.Status = &v
}

// GetBends returns the Bends field value if set, zero value otherwise.
func (o *ConnectionEntity) GetBends() []PositionDTO {
	if o == nil || IsNil(o.Bends) {
		var ret []PositionDTO
		return ret
	}
	return o.Bends
}

// GetBendsOk returns a tuple with the Bends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetBendsOk() ([]PositionDTO, bool) {
	if o == nil || IsNil(o.Bends) {
		return nil, false
	}
	return o.Bends, true
}

// HasBends returns a boolean if a field has been set.
func (o *ConnectionEntity) HasBends() bool {
	if o != nil && !IsNil(o.Bends) {
		return true
	}

	return false
}

// SetBends gets a reference to the given []PositionDTO and assigns it to the Bends field.
func (o *ConnectionEntity) SetBends(v []PositionDTO) {
	o.Bends = v
}

// GetLabelIndex returns the LabelIndex field value if set, zero value otherwise.
func (o *ConnectionEntity) GetLabelIndex() int32 {
	if o == nil || IsNil(o.LabelIndex) {
		var ret int32
		return ret
	}
	return *o.LabelIndex
}

// GetLabelIndexOk returns a tuple with the LabelIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetLabelIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.LabelIndex) {
		return nil, false
	}
	return o.LabelIndex, true
}

// HasLabelIndex returns a boolean if a field has been set.
func (o *ConnectionEntity) HasLabelIndex() bool {
	if o != nil && !IsNil(o.LabelIndex) {
		return true
	}

	return false
}

// SetLabelIndex gets a reference to the given int32 and assigns it to the LabelIndex field.
func (o *ConnectionEntity) SetLabelIndex(v int32) {
	o.LabelIndex = &v
}

// GetGetzIndex returns the GetzIndex field value if set, zero value otherwise.
func (o *ConnectionEntity) GetGetzIndex() int64 {
	if o == nil || IsNil(o.GetzIndex) {
		var ret int64
		return ret
	}
	return *o.GetzIndex
}

// GetGetzIndexOk returns a tuple with the GetzIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetGetzIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.GetzIndex) {
		return nil, false
	}
	return o.GetzIndex, true
}

// HasGetzIndex returns a boolean if a field has been set.
func (o *ConnectionEntity) HasGetzIndex() bool {
	if o != nil && !IsNil(o.GetzIndex) {
		return true
	}

	return false
}

// SetGetzIndex gets a reference to the given int64 and assigns it to the GetzIndex field.
func (o *ConnectionEntity) SetGetzIndex(v int64) {
	o.GetzIndex = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ConnectionEntity) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ConnectionEntity) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ConnectionEntity) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceGroupId returns the SourceGroupId field value if set, zero value otherwise.
func (o *ConnectionEntity) GetSourceGroupId() string {
	if o == nil || IsNil(o.SourceGroupId) {
		var ret string
		return ret
	}
	return *o.SourceGroupId
}

// GetSourceGroupIdOk returns a tuple with the SourceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetSourceGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceGroupId) {
		return nil, false
	}
	return o.SourceGroupId, true
}

// HasSourceGroupId returns a boolean if a field has been set.
func (o *ConnectionEntity) HasSourceGroupId() bool {
	if o != nil && !IsNil(o.SourceGroupId) {
		return true
	}

	return false
}

// SetSourceGroupId gets a reference to the given string and assigns it to the SourceGroupId field.
func (o *ConnectionEntity) SetSourceGroupId(v string) {
	o.SourceGroupId = &v
}

// GetSourceType returns the SourceType field value
func (o *ConnectionEntity) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *ConnectionEntity) SetSourceType(v string) {
	o.SourceType = v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *ConnectionEntity) GetDestinationId() string {
	if o == nil || IsNil(o.DestinationId) {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetDestinationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationId) {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *ConnectionEntity) HasDestinationId() bool {
	if o != nil && !IsNil(o.DestinationId) {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *ConnectionEntity) SetDestinationId(v string) {
	o.DestinationId = &v
}

// GetDestinationGroupId returns the DestinationGroupId field value if set, zero value otherwise.
func (o *ConnectionEntity) GetDestinationGroupId() string {
	if o == nil || IsNil(o.DestinationGroupId) {
		var ret string
		return ret
	}
	return *o.DestinationGroupId
}

// GetDestinationGroupIdOk returns a tuple with the DestinationGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetDestinationGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationGroupId) {
		return nil, false
	}
	return o.DestinationGroupId, true
}

// HasDestinationGroupId returns a boolean if a field has been set.
func (o *ConnectionEntity) HasDestinationGroupId() bool {
	if o != nil && !IsNil(o.DestinationGroupId) {
		return true
	}

	return false
}

// SetDestinationGroupId gets a reference to the given string and assigns it to the DestinationGroupId field.
func (o *ConnectionEntity) SetDestinationGroupId(v string) {
	o.DestinationGroupId = &v
}

// GetDestinationType returns the DestinationType field value
func (o *ConnectionEntity) GetDestinationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value
// and a boolean to check if the value has been set.
func (o *ConnectionEntity) GetDestinationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationType, true
}

// SetDestinationType sets field value
func (o *ConnectionEntity) SetDestinationType(v string) {
	o.DestinationType = v
}

func (o ConnectionEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Bulletins) {
		toSerialize["bulletins"] = o.Bulletins
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Bends) {
		toSerialize["bends"] = o.Bends
	}
	if !IsNil(o.LabelIndex) {
		toSerialize["labelIndex"] = o.LabelIndex
	}
	if !IsNil(o.GetzIndex) {
		toSerialize["getzIndex"] = o.GetzIndex
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.SourceGroupId) {
		toSerialize["sourceGroupId"] = o.SourceGroupId
	}
	toSerialize["sourceType"] = o.SourceType
	if !IsNil(o.DestinationId) {
		toSerialize["destinationId"] = o.DestinationId
	}
	if !IsNil(o.DestinationGroupId) {
		toSerialize["destinationGroupId"] = o.DestinationGroupId
	}
	toSerialize["destinationType"] = o.DestinationType
	return toSerialize, nil
}

type NullableConnectionEntity struct {
	value *ConnectionEntity
	isSet bool
}

func (v NullableConnectionEntity) Get() *ConnectionEntity {
	return v.value
}

func (v *NullableConnectionEntity) Set(val *ConnectionEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionEntity(val *ConnectionEntity) *NullableConnectionEntity {
	return &NullableConnectionEntity{value: val, isSet: true}
}

func (v NullableConnectionEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


