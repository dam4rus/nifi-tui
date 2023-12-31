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

// checks if the RemoteProcessGroupPortEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteProcessGroupPortEntity{}

// RemoteProcessGroupPortEntity struct for RemoteProcessGroupPortEntity
type RemoteProcessGroupPortEntity struct {
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
	RemoteProcessGroupPort *RemoteProcessGroupPortDTO `json:"remoteProcessGroupPort,omitempty"`
	OperatePermissions *PermissionsDTO `json:"operatePermissions,omitempty"`
}

// NewRemoteProcessGroupPortEntity instantiates a new RemoteProcessGroupPortEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteProcessGroupPortEntity() *RemoteProcessGroupPortEntity {
	this := RemoteProcessGroupPortEntity{}
	return &this
}

// NewRemoteProcessGroupPortEntityWithDefaults instantiates a new RemoteProcessGroupPortEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteProcessGroupPortEntityWithDefaults() *RemoteProcessGroupPortEntity {
	this := RemoteProcessGroupPortEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *RemoteProcessGroupPortEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteProcessGroupPortEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *RemoteProcessGroupPortEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetPosition() PositionDTO {
	if o == nil || IsNil(o.Position) {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *RemoteProcessGroupPortEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *RemoteProcessGroupPortEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetBulletins() []BulletinEntity {
	if o == nil || IsNil(o.Bulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.Bulletins) {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasBulletins() bool {
	if o != nil && !IsNil(o.Bulletins) {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *RemoteProcessGroupPortEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *RemoteProcessGroupPortEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetRemoteProcessGroupPort returns the RemoteProcessGroupPort field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetRemoteProcessGroupPort() RemoteProcessGroupPortDTO {
	if o == nil || IsNil(o.RemoteProcessGroupPort) {
		var ret RemoteProcessGroupPortDTO
		return ret
	}
	return *o.RemoteProcessGroupPort
}

// GetRemoteProcessGroupPortOk returns a tuple with the RemoteProcessGroupPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetRemoteProcessGroupPortOk() (*RemoteProcessGroupPortDTO, bool) {
	if o == nil || IsNil(o.RemoteProcessGroupPort) {
		return nil, false
	}
	return o.RemoteProcessGroupPort, true
}

// HasRemoteProcessGroupPort returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasRemoteProcessGroupPort() bool {
	if o != nil && !IsNil(o.RemoteProcessGroupPort) {
		return true
	}

	return false
}

// SetRemoteProcessGroupPort gets a reference to the given RemoteProcessGroupPortDTO and assigns it to the RemoteProcessGroupPort field.
func (o *RemoteProcessGroupPortEntity) SetRemoteProcessGroupPort(v RemoteProcessGroupPortDTO) {
	o.RemoteProcessGroupPort = &v
}

// GetOperatePermissions returns the OperatePermissions field value if set, zero value otherwise.
func (o *RemoteProcessGroupPortEntity) GetOperatePermissions() PermissionsDTO {
	if o == nil || IsNil(o.OperatePermissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.OperatePermissions
}

// GetOperatePermissionsOk returns a tuple with the OperatePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupPortEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.OperatePermissions) {
		return nil, false
	}
	return o.OperatePermissions, true
}

// HasOperatePermissions returns a boolean if a field has been set.
func (o *RemoteProcessGroupPortEntity) HasOperatePermissions() bool {
	if o != nil && !IsNil(o.OperatePermissions) {
		return true
	}

	return false
}

// SetOperatePermissions gets a reference to the given PermissionsDTO and assigns it to the OperatePermissions field.
func (o *RemoteProcessGroupPortEntity) SetOperatePermissions(v PermissionsDTO) {
	o.OperatePermissions = &v
}

func (o RemoteProcessGroupPortEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteProcessGroupPortEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RemoteProcessGroupPort) {
		toSerialize["remoteProcessGroupPort"] = o.RemoteProcessGroupPort
	}
	if !IsNil(o.OperatePermissions) {
		toSerialize["operatePermissions"] = o.OperatePermissions
	}
	return toSerialize, nil
}

type NullableRemoteProcessGroupPortEntity struct {
	value *RemoteProcessGroupPortEntity
	isSet bool
}

func (v NullableRemoteProcessGroupPortEntity) Get() *RemoteProcessGroupPortEntity {
	return v.value
}

func (v *NullableRemoteProcessGroupPortEntity) Set(val *RemoteProcessGroupPortEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteProcessGroupPortEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteProcessGroupPortEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteProcessGroupPortEntity(val *RemoteProcessGroupPortEntity) *NullableRemoteProcessGroupPortEntity {
	return &NullableRemoteProcessGroupPortEntity{value: val, isSet: true}
}

func (v NullableRemoteProcessGroupPortEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteProcessGroupPortEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


