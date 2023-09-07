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

// checks if the AttributeDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeDTO{}

// AttributeDTO struct for AttributeDTO
type AttributeDTO struct {
	// The attribute name.
	Name *string `json:"name,omitempty"`
	// The attribute value.
	Value *string `json:"value,omitempty"`
	// The value of the attribute before the event took place.
	PreviousValue *string `json:"previousValue,omitempty"`
}

// NewAttributeDTO instantiates a new AttributeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeDTO() *AttributeDTO {
	this := AttributeDTO{}
	return &this
}

// NewAttributeDTOWithDefaults instantiates a new AttributeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeDTOWithDefaults() *AttributeDTO {
	this := AttributeDTO{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttributeDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttributeDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttributeDTO) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AttributeDTO) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDTO) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AttributeDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AttributeDTO) SetValue(v string) {
	o.Value = &v
}

// GetPreviousValue returns the PreviousValue field value if set, zero value otherwise.
func (o *AttributeDTO) GetPreviousValue() string {
	if o == nil || IsNil(o.PreviousValue) {
		var ret string
		return ret
	}
	return *o.PreviousValue
}

// GetPreviousValueOk returns a tuple with the PreviousValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDTO) GetPreviousValueOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousValue) {
		return nil, false
	}
	return o.PreviousValue, true
}

// HasPreviousValue returns a boolean if a field has been set.
func (o *AttributeDTO) HasPreviousValue() bool {
	if o != nil && !IsNil(o.PreviousValue) {
		return true
	}

	return false
}

// SetPreviousValue gets a reference to the given string and assigns it to the PreviousValue field.
func (o *AttributeDTO) SetPreviousValue(v string) {
	o.PreviousValue = &v
}

func (o AttributeDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.PreviousValue) {
		toSerialize["previousValue"] = o.PreviousValue
	}
	return toSerialize, nil
}

type NullableAttributeDTO struct {
	value *AttributeDTO
	isSet bool
}

func (v NullableAttributeDTO) Get() *AttributeDTO {
	return v.value
}

func (v *NullableAttributeDTO) Set(val *AttributeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeDTO(val *AttributeDTO) *NullableAttributeDTO {
	return &NullableAttributeDTO{value: val, isSet: true}
}

func (v NullableAttributeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

