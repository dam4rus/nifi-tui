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

// checks if the ParameterProviderReferencingComponentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterProviderReferencingComponentDTO{}

// ParameterProviderReferencingComponentDTO struct for ParameterProviderReferencingComponentDTO
type ParameterProviderReferencingComponentDTO struct {
	// The id of the component referencing a parameter provider.
	Id *string `json:"id,omitempty"`
	// The name of the component referencing a parameter provider.
	Name *string `json:"name,omitempty"`
}

// NewParameterProviderReferencingComponentDTO instantiates a new ParameterProviderReferencingComponentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterProviderReferencingComponentDTO() *ParameterProviderReferencingComponentDTO {
	this := ParameterProviderReferencingComponentDTO{}
	return &this
}

// NewParameterProviderReferencingComponentDTOWithDefaults instantiates a new ParameterProviderReferencingComponentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterProviderReferencingComponentDTOWithDefaults() *ParameterProviderReferencingComponentDTO {
	this := ParameterProviderReferencingComponentDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ParameterProviderReferencingComponentDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderReferencingComponentDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ParameterProviderReferencingComponentDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ParameterProviderReferencingComponentDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParameterProviderReferencingComponentDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderReferencingComponentDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParameterProviderReferencingComponentDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParameterProviderReferencingComponentDTO) SetName(v string) {
	o.Name = &v
}

func (o ParameterProviderReferencingComponentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterProviderReferencingComponentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableParameterProviderReferencingComponentDTO struct {
	value *ParameterProviderReferencingComponentDTO
	isSet bool
}

func (v NullableParameterProviderReferencingComponentDTO) Get() *ParameterProviderReferencingComponentDTO {
	return v.value
}

func (v *NullableParameterProviderReferencingComponentDTO) Set(val *ParameterProviderReferencingComponentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterProviderReferencingComponentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterProviderReferencingComponentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterProviderReferencingComponentDTO(val *ParameterProviderReferencingComponentDTO) *NullableParameterProviderReferencingComponentDTO {
	return &NullableParameterProviderReferencingComponentDTO{value: val, isSet: true}
}

func (v NullableParameterProviderReferencingComponentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterProviderReferencingComponentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


