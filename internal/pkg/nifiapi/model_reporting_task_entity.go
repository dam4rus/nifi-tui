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

// checks if the ReportingTaskEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingTaskEntity{}

// ReportingTaskEntity struct for ReportingTaskEntity
type ReportingTaskEntity struct {
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
	Component *ReportingTaskDTO `json:"component,omitempty"`
	OperatePermissions *PermissionsDTO `json:"operatePermissions,omitempty"`
	Status *ReportingTaskStatusDTO `json:"status,omitempty"`
}

// NewReportingTaskEntity instantiates a new ReportingTaskEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingTaskEntity() *ReportingTaskEntity {
	this := ReportingTaskEntity{}
	return &this
}

// NewReportingTaskEntityWithDefaults instantiates a new ReportingTaskEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingTaskEntityWithDefaults() *ReportingTaskEntity {
	this := ReportingTaskEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *ReportingTaskEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportingTaskEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ReportingTaskEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetPosition() PositionDTO {
	if o == nil || IsNil(o.Position) {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *ReportingTaskEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ReportingTaskEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetBulletins() []BulletinEntity {
	if o == nil || IsNil(o.Bulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.Bulletins) {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasBulletins() bool {
	if o != nil && !IsNil(o.Bulletins) {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *ReportingTaskEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ReportingTaskEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetComponent() ReportingTaskDTO {
	if o == nil || IsNil(o.Component) {
		var ret ReportingTaskDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetComponentOk() (*ReportingTaskDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ReportingTaskDTO and assigns it to the Component field.
func (o *ReportingTaskEntity) SetComponent(v ReportingTaskDTO) {
	o.Component = &v
}

// GetOperatePermissions returns the OperatePermissions field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetOperatePermissions() PermissionsDTO {
	if o == nil || IsNil(o.OperatePermissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.OperatePermissions
}

// GetOperatePermissionsOk returns a tuple with the OperatePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetOperatePermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.OperatePermissions) {
		return nil, false
	}
	return o.OperatePermissions, true
}

// HasOperatePermissions returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasOperatePermissions() bool {
	if o != nil && !IsNil(o.OperatePermissions) {
		return true
	}

	return false
}

// SetOperatePermissions gets a reference to the given PermissionsDTO and assigns it to the OperatePermissions field.
func (o *ReportingTaskEntity) SetOperatePermissions(v PermissionsDTO) {
	o.OperatePermissions = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportingTaskEntity) GetStatus() ReportingTaskStatusDTO {
	if o == nil || IsNil(o.Status) {
		var ret ReportingTaskStatusDTO
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingTaskEntity) GetStatusOk() (*ReportingTaskStatusDTO, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportingTaskEntity) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ReportingTaskStatusDTO and assigns it to the Status field.
func (o *ReportingTaskEntity) SetStatus(v ReportingTaskStatusDTO) {
	o.Status = &v
}

func (o ReportingTaskEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingTaskEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OperatePermissions) {
		toSerialize["operatePermissions"] = o.OperatePermissions
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableReportingTaskEntity struct {
	value *ReportingTaskEntity
	isSet bool
}

func (v NullableReportingTaskEntity) Get() *ReportingTaskEntity {
	return v.value
}

func (v *NullableReportingTaskEntity) Set(val *ReportingTaskEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingTaskEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingTaskEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingTaskEntity(val *ReportingTaskEntity) *NullableReportingTaskEntity {
	return &NullableReportingTaskEntity{value: val, isSet: true}
}

func (v NullableReportingTaskEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingTaskEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


