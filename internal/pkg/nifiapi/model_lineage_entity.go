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

// checks if the LineageEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineageEntity{}

// LineageEntity struct for LineageEntity
type LineageEntity struct {
	Lineage *LineageDTO `json:"lineage,omitempty"`
}

// NewLineageEntity instantiates a new LineageEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineageEntity() *LineageEntity {
	this := LineageEntity{}
	return &this
}

// NewLineageEntityWithDefaults instantiates a new LineageEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineageEntityWithDefaults() *LineageEntity {
	this := LineageEntity{}
	return &this
}

// GetLineage returns the Lineage field value if set, zero value otherwise.
func (o *LineageEntity) GetLineage() LineageDTO {
	if o == nil || IsNil(o.Lineage) {
		var ret LineageDTO
		return ret
	}
	return *o.Lineage
}

// GetLineageOk returns a tuple with the Lineage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineageEntity) GetLineageOk() (*LineageDTO, bool) {
	if o == nil || IsNil(o.Lineage) {
		return nil, false
	}
	return o.Lineage, true
}

// HasLineage returns a boolean if a field has been set.
func (o *LineageEntity) HasLineage() bool {
	if o != nil && !IsNil(o.Lineage) {
		return true
	}

	return false
}

// SetLineage gets a reference to the given LineageDTO and assigns it to the Lineage field.
func (o *LineageEntity) SetLineage(v LineageDTO) {
	o.Lineage = &v
}

func (o LineageEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineageEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lineage) {
		toSerialize["lineage"] = o.Lineage
	}
	return toSerialize, nil
}

type NullableLineageEntity struct {
	value *LineageEntity
	isSet bool
}

func (v NullableLineageEntity) Get() *LineageEntity {
	return v.value
}

func (v *NullableLineageEntity) Set(val *LineageEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableLineageEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableLineageEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineageEntity(val *LineageEntity) *NullableLineageEntity {
	return &NullableLineageEntity{value: val, isSet: true}
}

func (v NullableLineageEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineageEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


