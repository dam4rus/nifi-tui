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

// checks if the ParameterProviderConfigurationEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterProviderConfigurationEntity{}

// ParameterProviderConfigurationEntity struct for ParameterProviderConfigurationEntity
type ParameterProviderConfigurationEntity struct {
	// The id of the component.
	Id *string `json:"id,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
	Component *ParameterProviderConfigurationDTO `json:"component,omitempty"`
}

// NewParameterProviderConfigurationEntity instantiates a new ParameterProviderConfigurationEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterProviderConfigurationEntity() *ParameterProviderConfigurationEntity {
	this := ParameterProviderConfigurationEntity{}
	return &this
}

// NewParameterProviderConfigurationEntityWithDefaults instantiates a new ParameterProviderConfigurationEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterProviderConfigurationEntityWithDefaults() *ParameterProviderConfigurationEntity {
	this := ParameterProviderConfigurationEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ParameterProviderConfigurationEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderConfigurationEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ParameterProviderConfigurationEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ParameterProviderConfigurationEntity) SetId(v string) {
	o.Id = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ParameterProviderConfigurationEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderConfigurationEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ParameterProviderConfigurationEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ParameterProviderConfigurationEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ParameterProviderConfigurationEntity) GetComponent() ParameterProviderConfigurationDTO {
	if o == nil || IsNil(o.Component) {
		var ret ParameterProviderConfigurationDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderConfigurationEntity) GetComponentOk() (*ParameterProviderConfigurationDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ParameterProviderConfigurationEntity) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ParameterProviderConfigurationDTO and assigns it to the Component field.
func (o *ParameterProviderConfigurationEntity) SetComponent(v ParameterProviderConfigurationDTO) {
	o.Component = &v
}

func (o ParameterProviderConfigurationEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterProviderConfigurationEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	return toSerialize, nil
}

type NullableParameterProviderConfigurationEntity struct {
	value *ParameterProviderConfigurationEntity
	isSet bool
}

func (v NullableParameterProviderConfigurationEntity) Get() *ParameterProviderConfigurationEntity {
	return v.value
}

func (v *NullableParameterProviderConfigurationEntity) Set(val *ParameterProviderConfigurationEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterProviderConfigurationEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterProviderConfigurationEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterProviderConfigurationEntity(val *ParameterProviderConfigurationEntity) *NullableParameterProviderConfigurationEntity {
	return &NullableParameterProviderConfigurationEntity{value: val, isSet: true}
}

func (v NullableParameterProviderConfigurationEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterProviderConfigurationEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


