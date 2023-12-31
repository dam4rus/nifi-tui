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

// checks if the ComponentRestrictionPermissionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentRestrictionPermissionDTO{}

// ComponentRestrictionPermissionDTO struct for ComponentRestrictionPermissionDTO
type ComponentRestrictionPermissionDTO struct {
	RequiredPermission *RequiredPermissionDTO `json:"requiredPermission,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
}

// NewComponentRestrictionPermissionDTO instantiates a new ComponentRestrictionPermissionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentRestrictionPermissionDTO() *ComponentRestrictionPermissionDTO {
	this := ComponentRestrictionPermissionDTO{}
	return &this
}

// NewComponentRestrictionPermissionDTOWithDefaults instantiates a new ComponentRestrictionPermissionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentRestrictionPermissionDTOWithDefaults() *ComponentRestrictionPermissionDTO {
	this := ComponentRestrictionPermissionDTO{}
	return &this
}

// GetRequiredPermission returns the RequiredPermission field value if set, zero value otherwise.
func (o *ComponentRestrictionPermissionDTO) GetRequiredPermission() RequiredPermissionDTO {
	if o == nil || IsNil(o.RequiredPermission) {
		var ret RequiredPermissionDTO
		return ret
	}
	return *o.RequiredPermission
}

// GetRequiredPermissionOk returns a tuple with the RequiredPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentRestrictionPermissionDTO) GetRequiredPermissionOk() (*RequiredPermissionDTO, bool) {
	if o == nil || IsNil(o.RequiredPermission) {
		return nil, false
	}
	return o.RequiredPermission, true
}

// HasRequiredPermission returns a boolean if a field has been set.
func (o *ComponentRestrictionPermissionDTO) HasRequiredPermission() bool {
	if o != nil && !IsNil(o.RequiredPermission) {
		return true
	}

	return false
}

// SetRequiredPermission gets a reference to the given RequiredPermissionDTO and assigns it to the RequiredPermission field.
func (o *ComponentRestrictionPermissionDTO) SetRequiredPermission(v RequiredPermissionDTO) {
	o.RequiredPermission = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ComponentRestrictionPermissionDTO) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentRestrictionPermissionDTO) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ComponentRestrictionPermissionDTO) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ComponentRestrictionPermissionDTO) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

func (o ComponentRestrictionPermissionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentRestrictionPermissionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequiredPermission) {
		toSerialize["requiredPermission"] = o.RequiredPermission
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableComponentRestrictionPermissionDTO struct {
	value *ComponentRestrictionPermissionDTO
	isSet bool
}

func (v NullableComponentRestrictionPermissionDTO) Get() *ComponentRestrictionPermissionDTO {
	return v.value
}

func (v *NullableComponentRestrictionPermissionDTO) Set(val *ComponentRestrictionPermissionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentRestrictionPermissionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentRestrictionPermissionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentRestrictionPermissionDTO(val *ComponentRestrictionPermissionDTO) *NullableComponentRestrictionPermissionDTO {
	return &NullableComponentRestrictionPermissionDTO{value: val, isSet: true}
}

func (v NullableComponentRestrictionPermissionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentRestrictionPermissionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


