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

// checks if the PortEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortEntity{}

// PortEntity struct for PortEntity
type PortEntity struct {
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
	Component *PortDTO `json:"component,omitempty"`
	Status *PortStatusDTO `json:"status,omitempty"`
	PortType *string `json:"portType,omitempty"`
	OperatePermissions *PermissionsDTO `json:"operatePermissions,omitempty"`
	// Whether this port can be accessed remotely via Site-to-Site protocol.
	AllowRemoteAccess *bool `json:"allowRemoteAccess,omitempty"`
}

// NewPortEntity instantiates a new PortEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortEntity() *PortEntity {
	this := PortEntity{}
	return &this
}

// NewPortEntityWithDefaults instantiates a new PortEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortEntityWithDefaults() *PortEntity {
	this := PortEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *PortEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *PortEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *PortEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *PortEntity) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PortEntity) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *PortEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PortEntity) GetPosition() PositionDTO {
	if o == nil || IsNil(o.Position) {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PortEntity) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *PortEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PortEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PortEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *PortEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *PortEntity) GetBulletins() []BulletinEntity {
	if o == nil || IsNil(o.Bulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.Bulletins) {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *PortEntity) HasBulletins() bool {
	if o != nil && !IsNil(o.Bulletins) {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *PortEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *PortEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *PortEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *PortEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *PortEntity) GetComponent() PortDTO {
	if o == nil || IsNil(o.Component) {
		var ret PortDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetComponentOk() (*PortDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *PortEntity) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given PortDTO and assigns it to the Component field.
func (o *PortEntity) SetComponent(v PortDTO) {
	o.Component = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PortEntity) GetStatus() PortStatusDTO {
	if o == nil || IsNil(o.Status) {
		var ret PortStatusDTO
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetStatusOk() (*PortStatusDTO, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PortEntity) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PortStatusDTO and assigns it to the Status field.
func (o *PortEntity) SetStatus(v PortStatusDTO) {
	o.Status = &v
}

// GetPortType returns the PortType field value if set, zero value otherwise.
func (o *PortEntity) GetPortType() string {
	if o == nil || IsNil(o.PortType) {
		var ret string
		return ret
	}
	return *o.PortType
}

// GetPortTypeOk returns a tuple with the PortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetPortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PortType) {
		return nil, false
	}
	return o.PortType, true
}

// HasPortType returns a boolean if a field has been set.
func (o *PortEntity) HasPortType() bool {
	if o != nil && !IsNil(o.PortType) {
		return true
	}

	return false
}

// SetPortType gets a reference to the given string and assigns it to the PortType field.
func (o *PortEntity) SetPortType(v string) {
	o.PortType = &v
}

// GetOperatePermissions returns the OperatePermissions field value if set, zero value otherwise.
func (o *PortEntity) GetOperatePermissions() PermissionsDTO {
	if o == nil || IsNil(o.OperatePermissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.OperatePermissions
}

// GetOperatePermissionsOk returns a tuple with the OperatePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.OperatePermissions) {
		return nil, false
	}
	return o.OperatePermissions, true
}

// HasOperatePermissions returns a boolean if a field has been set.
func (o *PortEntity) HasOperatePermissions() bool {
	if o != nil && !IsNil(o.OperatePermissions) {
		return true
	}

	return false
}

// SetOperatePermissions gets a reference to the given PermissionsDTO and assigns it to the OperatePermissions field.
func (o *PortEntity) SetOperatePermissions(v PermissionsDTO) {
	o.OperatePermissions = &v
}

// GetAllowRemoteAccess returns the AllowRemoteAccess field value if set, zero value otherwise.
func (o *PortEntity) GetAllowRemoteAccess() bool {
	if o == nil || IsNil(o.AllowRemoteAccess) {
		var ret bool
		return ret
	}
	return *o.AllowRemoteAccess
}

// GetAllowRemoteAccessOk returns a tuple with the AllowRemoteAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortEntity) GetAllowRemoteAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRemoteAccess) {
		return nil, false
	}
	return o.AllowRemoteAccess, true
}

// HasAllowRemoteAccess returns a boolean if a field has been set.
func (o *PortEntity) HasAllowRemoteAccess() bool {
	if o != nil && !IsNil(o.AllowRemoteAccess) {
		return true
	}

	return false
}

// SetAllowRemoteAccess gets a reference to the given bool and assigns it to the AllowRemoteAccess field.
func (o *PortEntity) SetAllowRemoteAccess(v bool) {
	o.AllowRemoteAccess = &v
}

func (o PortEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PortType) {
		toSerialize["portType"] = o.PortType
	}
	if !IsNil(o.OperatePermissions) {
		toSerialize["operatePermissions"] = o.OperatePermissions
	}
	if !IsNil(o.AllowRemoteAccess) {
		toSerialize["allowRemoteAccess"] = o.AllowRemoteAccess
	}
	return toSerialize, nil
}

type NullablePortEntity struct {
	value *PortEntity
	isSet bool
}

func (v NullablePortEntity) Get() *PortEntity {
	return v.value
}

func (v *NullablePortEntity) Set(val *PortEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePortEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePortEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortEntity(val *PortEntity) *NullablePortEntity {
	return &NullablePortEntity{value: val, isSet: true}
}

func (v NullablePortEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


