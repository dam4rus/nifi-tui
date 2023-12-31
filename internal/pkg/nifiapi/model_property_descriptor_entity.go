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

// checks if the PropertyDescriptorEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyDescriptorEntity{}

// PropertyDescriptorEntity struct for PropertyDescriptorEntity
type PropertyDescriptorEntity struct {
	PropertyDescriptor *PropertyDescriptorDTO `json:"propertyDescriptor,omitempty"`
}

// NewPropertyDescriptorEntity instantiates a new PropertyDescriptorEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDescriptorEntity() *PropertyDescriptorEntity {
	this := PropertyDescriptorEntity{}
	return &this
}

// NewPropertyDescriptorEntityWithDefaults instantiates a new PropertyDescriptorEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDescriptorEntityWithDefaults() *PropertyDescriptorEntity {
	this := PropertyDescriptorEntity{}
	return &this
}

// GetPropertyDescriptor returns the PropertyDescriptor field value if set, zero value otherwise.
func (o *PropertyDescriptorEntity) GetPropertyDescriptor() PropertyDescriptorDTO {
	if o == nil || IsNil(o.PropertyDescriptor) {
		var ret PropertyDescriptorDTO
		return ret
	}
	return *o.PropertyDescriptor
}

// GetPropertyDescriptorOk returns a tuple with the PropertyDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorEntity) GetPropertyDescriptorOk() (*PropertyDescriptorDTO, bool) {
	if o == nil || IsNil(o.PropertyDescriptor) {
		return nil, false
	}
	return o.PropertyDescriptor, true
}

// HasPropertyDescriptor returns a boolean if a field has been set.
func (o *PropertyDescriptorEntity) HasPropertyDescriptor() bool {
	if o != nil && !IsNil(o.PropertyDescriptor) {
		return true
	}

	return false
}

// SetPropertyDescriptor gets a reference to the given PropertyDescriptorDTO and assigns it to the PropertyDescriptor field.
func (o *PropertyDescriptorEntity) SetPropertyDescriptor(v PropertyDescriptorDTO) {
	o.PropertyDescriptor = &v
}

func (o PropertyDescriptorEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyDescriptorEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyDescriptor) {
		toSerialize["propertyDescriptor"] = o.PropertyDescriptor
	}
	return toSerialize, nil
}

type NullablePropertyDescriptorEntity struct {
	value *PropertyDescriptorEntity
	isSet bool
}

func (v NullablePropertyDescriptorEntity) Get() *PropertyDescriptorEntity {
	return v.value
}

func (v *NullablePropertyDescriptorEntity) Set(val *PropertyDescriptorEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDescriptorEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDescriptorEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDescriptorEntity(val *PropertyDescriptorEntity) *NullablePropertyDescriptorEntity {
	return &NullablePropertyDescriptorEntity{value: val, isSet: true}
}

func (v NullablePropertyDescriptorEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDescriptorEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


