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

// checks if the ProvenanceSearchValueDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvenanceSearchValueDTO{}

// ProvenanceSearchValueDTO struct for ProvenanceSearchValueDTO
type ProvenanceSearchValueDTO struct {
	// The search value.
	Value *string `json:"value,omitempty"`
	// Query for all except for search value.
	Inverse *bool `json:"inverse,omitempty"`
}

// NewProvenanceSearchValueDTO instantiates a new ProvenanceSearchValueDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvenanceSearchValueDTO() *ProvenanceSearchValueDTO {
	this := ProvenanceSearchValueDTO{}
	return &this
}

// NewProvenanceSearchValueDTOWithDefaults instantiates a new ProvenanceSearchValueDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvenanceSearchValueDTOWithDefaults() *ProvenanceSearchValueDTO {
	this := ProvenanceSearchValueDTO{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ProvenanceSearchValueDTO) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvenanceSearchValueDTO) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ProvenanceSearchValueDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ProvenanceSearchValueDTO) SetValue(v string) {
	o.Value = &v
}

// GetInverse returns the Inverse field value if set, zero value otherwise.
func (o *ProvenanceSearchValueDTO) GetInverse() bool {
	if o == nil || IsNil(o.Inverse) {
		var ret bool
		return ret
	}
	return *o.Inverse
}

// GetInverseOk returns a tuple with the Inverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvenanceSearchValueDTO) GetInverseOk() (*bool, bool) {
	if o == nil || IsNil(o.Inverse) {
		return nil, false
	}
	return o.Inverse, true
}

// HasInverse returns a boolean if a field has been set.
func (o *ProvenanceSearchValueDTO) HasInverse() bool {
	if o != nil && !IsNil(o.Inverse) {
		return true
	}

	return false
}

// SetInverse gets a reference to the given bool and assigns it to the Inverse field.
func (o *ProvenanceSearchValueDTO) SetInverse(v bool) {
	o.Inverse = &v
}

func (o ProvenanceSearchValueDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvenanceSearchValueDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Inverse) {
		toSerialize["inverse"] = o.Inverse
	}
	return toSerialize, nil
}

type NullableProvenanceSearchValueDTO struct {
	value *ProvenanceSearchValueDTO
	isSet bool
}

func (v NullableProvenanceSearchValueDTO) Get() *ProvenanceSearchValueDTO {
	return v.value
}

func (v *NullableProvenanceSearchValueDTO) Set(val *ProvenanceSearchValueDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProvenanceSearchValueDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProvenanceSearchValueDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvenanceSearchValueDTO(val *ProvenanceSearchValueDTO) *NullableProvenanceSearchValueDTO {
	return &NullableProvenanceSearchValueDTO{value: val, isSet: true}
}

func (v NullableProvenanceSearchValueDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvenanceSearchValueDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


