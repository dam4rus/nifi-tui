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

// checks if the RunStatusDetailsRequestEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunStatusDetailsRequestEntity{}

// RunStatusDetailsRequestEntity struct for RunStatusDetailsRequestEntity
type RunStatusDetailsRequestEntity struct {
	// The IDs of all processors whose run status details should be provided
	ProcessorIds []string `json:"processorIds,omitempty"`
}

// NewRunStatusDetailsRequestEntity instantiates a new RunStatusDetailsRequestEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunStatusDetailsRequestEntity() *RunStatusDetailsRequestEntity {
	this := RunStatusDetailsRequestEntity{}
	return &this
}

// NewRunStatusDetailsRequestEntityWithDefaults instantiates a new RunStatusDetailsRequestEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunStatusDetailsRequestEntityWithDefaults() *RunStatusDetailsRequestEntity {
	this := RunStatusDetailsRequestEntity{}
	return &this
}

// GetProcessorIds returns the ProcessorIds field value if set, zero value otherwise.
func (o *RunStatusDetailsRequestEntity) GetProcessorIds() []string {
	if o == nil || IsNil(o.ProcessorIds) {
		var ret []string
		return ret
	}
	return o.ProcessorIds
}

// GetProcessorIdsOk returns a tuple with the ProcessorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunStatusDetailsRequestEntity) GetProcessorIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProcessorIds) {
		return nil, false
	}
	return o.ProcessorIds, true
}

// HasProcessorIds returns a boolean if a field has been set.
func (o *RunStatusDetailsRequestEntity) HasProcessorIds() bool {
	if o != nil && !IsNil(o.ProcessorIds) {
		return true
	}

	return false
}

// SetProcessorIds gets a reference to the given []string and assigns it to the ProcessorIds field.
func (o *RunStatusDetailsRequestEntity) SetProcessorIds(v []string) {
	o.ProcessorIds = v
}

func (o RunStatusDetailsRequestEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunStatusDetailsRequestEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProcessorIds) {
		toSerialize["processorIds"] = o.ProcessorIds
	}
	return toSerialize, nil
}

type NullableRunStatusDetailsRequestEntity struct {
	value *RunStatusDetailsRequestEntity
	isSet bool
}

func (v NullableRunStatusDetailsRequestEntity) Get() *RunStatusDetailsRequestEntity {
	return v.value
}

func (v *NullableRunStatusDetailsRequestEntity) Set(val *RunStatusDetailsRequestEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRunStatusDetailsRequestEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRunStatusDetailsRequestEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunStatusDetailsRequestEntity(val *RunStatusDetailsRequestEntity) *NullableRunStatusDetailsRequestEntity {
	return &NullableRunStatusDetailsRequestEntity{value: val, isSet: true}
}

func (v NullableRunStatusDetailsRequestEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunStatusDetailsRequestEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


