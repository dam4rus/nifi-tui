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

// checks if the ExplicitRestrictionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExplicitRestrictionDTO{}

// ExplicitRestrictionDTO struct for ExplicitRestrictionDTO
type ExplicitRestrictionDTO struct {
	RequiredPermission *RequiredPermissionDTO `json:"requiredPermission,omitempty"`
	// The description of why the usage of this component is restricted for this required permission.
	Explanation *string `json:"explanation,omitempty"`
}

// NewExplicitRestrictionDTO instantiates a new ExplicitRestrictionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplicitRestrictionDTO() *ExplicitRestrictionDTO {
	this := ExplicitRestrictionDTO{}
	return &this
}

// NewExplicitRestrictionDTOWithDefaults instantiates a new ExplicitRestrictionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplicitRestrictionDTOWithDefaults() *ExplicitRestrictionDTO {
	this := ExplicitRestrictionDTO{}
	return &this
}

// GetRequiredPermission returns the RequiredPermission field value if set, zero value otherwise.
func (o *ExplicitRestrictionDTO) GetRequiredPermission() RequiredPermissionDTO {
	if o == nil || IsNil(o.RequiredPermission) {
		var ret RequiredPermissionDTO
		return ret
	}
	return *o.RequiredPermission
}

// GetRequiredPermissionOk returns a tuple with the RequiredPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplicitRestrictionDTO) GetRequiredPermissionOk() (*RequiredPermissionDTO, bool) {
	if o == nil || IsNil(o.RequiredPermission) {
		return nil, false
	}
	return o.RequiredPermission, true
}

// HasRequiredPermission returns a boolean if a field has been set.
func (o *ExplicitRestrictionDTO) HasRequiredPermission() bool {
	if o != nil && !IsNil(o.RequiredPermission) {
		return true
	}

	return false
}

// SetRequiredPermission gets a reference to the given RequiredPermissionDTO and assigns it to the RequiredPermission field.
func (o *ExplicitRestrictionDTO) SetRequiredPermission(v RequiredPermissionDTO) {
	o.RequiredPermission = &v
}

// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *ExplicitRestrictionDTO) GetExplanation() string {
	if o == nil || IsNil(o.Explanation) {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExplicitRestrictionDTO) GetExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.Explanation) {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *ExplicitRestrictionDTO) HasExplanation() bool {
	if o != nil && !IsNil(o.Explanation) {
		return true
	}

	return false
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *ExplicitRestrictionDTO) SetExplanation(v string) {
	o.Explanation = &v
}

func (o ExplicitRestrictionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplicitRestrictionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequiredPermission) {
		toSerialize["requiredPermission"] = o.RequiredPermission
	}
	if !IsNil(o.Explanation) {
		toSerialize["explanation"] = o.Explanation
	}
	return toSerialize, nil
}

type NullableExplicitRestrictionDTO struct {
	value *ExplicitRestrictionDTO
	isSet bool
}

func (v NullableExplicitRestrictionDTO) Get() *ExplicitRestrictionDTO {
	return v.value
}

func (v *NullableExplicitRestrictionDTO) Set(val *ExplicitRestrictionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableExplicitRestrictionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableExplicitRestrictionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplicitRestrictionDTO(val *ExplicitRestrictionDTO) *NullableExplicitRestrictionDTO {
	return &NullableExplicitRestrictionDTO{value: val, isSet: true}
}

func (v NullableExplicitRestrictionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplicitRestrictionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


