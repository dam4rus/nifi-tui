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

// checks if the RemoteProcessGroupEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteProcessGroupEntity{}

// RemoteProcessGroupEntity struct for RemoteProcessGroupEntity
type RemoteProcessGroupEntity struct {
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
	Component *RemoteProcessGroupDTO `json:"component,omitempty"`
	Status *RemoteProcessGroupStatusDTO `json:"status,omitempty"`
	// The number of remote input ports currently available on the target.
	InputPortCount *int32 `json:"inputPortCount,omitempty"`
	// The number of remote output ports currently available on the target.
	OutputPortCount *int32 `json:"outputPortCount,omitempty"`
	OperatePermissions *PermissionsDTO `json:"operatePermissions,omitempty"`
}

// NewRemoteProcessGroupEntity instantiates a new RemoteProcessGroupEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteProcessGroupEntity() *RemoteProcessGroupEntity {
	this := RemoteProcessGroupEntity{}
	return &this
}

// NewRemoteProcessGroupEntityWithDefaults instantiates a new RemoteProcessGroupEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteProcessGroupEntityWithDefaults() *RemoteProcessGroupEntity {
	this := RemoteProcessGroupEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *RemoteProcessGroupEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteProcessGroupEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *RemoteProcessGroupEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetPosition() PositionDTO {
	if o == nil || IsNil(o.Position) {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *RemoteProcessGroupEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *RemoteProcessGroupEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetBulletins() []BulletinEntity {
	if o == nil || IsNil(o.Bulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.Bulletins) {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasBulletins() bool {
	if o != nil && !IsNil(o.Bulletins) {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *RemoteProcessGroupEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *RemoteProcessGroupEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetComponent() RemoteProcessGroupDTO {
	if o == nil || IsNil(o.Component) {
		var ret RemoteProcessGroupDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetComponentOk() (*RemoteProcessGroupDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given RemoteProcessGroupDTO and assigns it to the Component field.
func (o *RemoteProcessGroupEntity) SetComponent(v RemoteProcessGroupDTO) {
	o.Component = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetStatus() RemoteProcessGroupStatusDTO {
	if o == nil || IsNil(o.Status) {
		var ret RemoteProcessGroupStatusDTO
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetStatusOk() (*RemoteProcessGroupStatusDTO, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RemoteProcessGroupStatusDTO and assigns it to the Status field.
func (o *RemoteProcessGroupEntity) SetStatus(v RemoteProcessGroupStatusDTO) {
	o.Status = &v
}

// GetInputPortCount returns the InputPortCount field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetInputPortCount() int32 {
	if o == nil || IsNil(o.InputPortCount) {
		var ret int32
		return ret
	}
	return *o.InputPortCount
}

// GetInputPortCountOk returns a tuple with the InputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetInputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InputPortCount) {
		return nil, false
	}
	return o.InputPortCount, true
}

// HasInputPortCount returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasInputPortCount() bool {
	if o != nil && !IsNil(o.InputPortCount) {
		return true
	}

	return false
}

// SetInputPortCount gets a reference to the given int32 and assigns it to the InputPortCount field.
func (o *RemoteProcessGroupEntity) SetInputPortCount(v int32) {
	o.InputPortCount = &v
}

// GetOutputPortCount returns the OutputPortCount field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetOutputPortCount() int32 {
	if o == nil || IsNil(o.OutputPortCount) {
		var ret int32
		return ret
	}
	return *o.OutputPortCount
}

// GetOutputPortCountOk returns a tuple with the OutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetOutputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputPortCount) {
		return nil, false
	}
	return o.OutputPortCount, true
}

// HasOutputPortCount returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasOutputPortCount() bool {
	if o != nil && !IsNil(o.OutputPortCount) {
		return true
	}

	return false
}

// SetOutputPortCount gets a reference to the given int32 and assigns it to the OutputPortCount field.
func (o *RemoteProcessGroupEntity) SetOutputPortCount(v int32) {
	o.OutputPortCount = &v
}

// GetOperatePermissions returns the OperatePermissions field value if set, zero value otherwise.
func (o *RemoteProcessGroupEntity) GetOperatePermissions() PermissionsDTO {
	if o == nil || IsNil(o.OperatePermissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.OperatePermissions
}

// GetOperatePermissionsOk returns a tuple with the OperatePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.OperatePermissions) {
		return nil, false
	}
	return o.OperatePermissions, true
}

// HasOperatePermissions returns a boolean if a field has been set.
func (o *RemoteProcessGroupEntity) HasOperatePermissions() bool {
	if o != nil && !IsNil(o.OperatePermissions) {
		return true
	}

	return false
}

// SetOperatePermissions gets a reference to the given PermissionsDTO and assigns it to the OperatePermissions field.
func (o *RemoteProcessGroupEntity) SetOperatePermissions(v PermissionsDTO) {
	o.OperatePermissions = &v
}

func (o RemoteProcessGroupEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteProcessGroupEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.InputPortCount) {
		toSerialize["inputPortCount"] = o.InputPortCount
	}
	if !IsNil(o.OutputPortCount) {
		toSerialize["outputPortCount"] = o.OutputPortCount
	}
	if !IsNil(o.OperatePermissions) {
		toSerialize["operatePermissions"] = o.OperatePermissions
	}
	return toSerialize, nil
}

type NullableRemoteProcessGroupEntity struct {
	value *RemoteProcessGroupEntity
	isSet bool
}

func (v NullableRemoteProcessGroupEntity) Get() *RemoteProcessGroupEntity {
	return v.value
}

func (v *NullableRemoteProcessGroupEntity) Set(val *RemoteProcessGroupEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteProcessGroupEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteProcessGroupEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteProcessGroupEntity(val *RemoteProcessGroupEntity) *NullableRemoteProcessGroupEntity {
	return &NullableRemoteProcessGroupEntity{value: val, isSet: true}
}

func (v NullableRemoteProcessGroupEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteProcessGroupEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

