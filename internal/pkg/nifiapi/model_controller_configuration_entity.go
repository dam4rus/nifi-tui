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

// checks if the ControllerConfigurationEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllerConfigurationEntity{}

// ControllerConfigurationEntity struct for ControllerConfigurationEntity
type ControllerConfigurationEntity struct {
	Revision *RevisionDTO `json:"revision,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
	Component *ControllerConfigurationDTO `json:"component,omitempty"`
}

// NewControllerConfigurationEntity instantiates a new ControllerConfigurationEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerConfigurationEntity() *ControllerConfigurationEntity {
	this := ControllerConfigurationEntity{}
	return &this
}

// NewControllerConfigurationEntityWithDefaults instantiates a new ControllerConfigurationEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerConfigurationEntityWithDefaults() *ControllerConfigurationEntity {
	this := ControllerConfigurationEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ControllerConfigurationEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerConfigurationEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ControllerConfigurationEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *ControllerConfigurationEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ControllerConfigurationEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerConfigurationEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ControllerConfigurationEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ControllerConfigurationEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ControllerConfigurationEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerConfigurationEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ControllerConfigurationEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ControllerConfigurationEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ControllerConfigurationEntity) GetComponent() ControllerConfigurationDTO {
	if o == nil || IsNil(o.Component) {
		var ret ControllerConfigurationDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerConfigurationEntity) GetComponentOk() (*ControllerConfigurationDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ControllerConfigurationEntity) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ControllerConfigurationDTO and assigns it to the Component field.
func (o *ControllerConfigurationEntity) SetComponent(v ControllerConfigurationDTO) {
	o.Component = &v
}

func (o ControllerConfigurationEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllerConfigurationEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	return toSerialize, nil
}

type NullableControllerConfigurationEntity struct {
	value *ControllerConfigurationEntity
	isSet bool
}

func (v NullableControllerConfigurationEntity) Get() *ControllerConfigurationEntity {
	return v.value
}

func (v *NullableControllerConfigurationEntity) Set(val *ControllerConfigurationEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerConfigurationEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerConfigurationEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerConfigurationEntity(val *ControllerConfigurationEntity) *NullableControllerConfigurationEntity {
	return &NullableControllerConfigurationEntity{value: val, isSet: true}
}

func (v NullableControllerConfigurationEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerConfigurationEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


