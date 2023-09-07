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

// checks if the FlowFileEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowFileEntity{}

// FlowFileEntity struct for FlowFileEntity
type FlowFileEntity struct {
	FlowFile *FlowFileDTO `json:"flowFile,omitempty"`
}

// NewFlowFileEntity instantiates a new FlowFileEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowFileEntity() *FlowFileEntity {
	this := FlowFileEntity{}
	return &this
}

// NewFlowFileEntityWithDefaults instantiates a new FlowFileEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowFileEntityWithDefaults() *FlowFileEntity {
	this := FlowFileEntity{}
	return &this
}

// GetFlowFile returns the FlowFile field value if set, zero value otherwise.
func (o *FlowFileEntity) GetFlowFile() FlowFileDTO {
	if o == nil || IsNil(o.FlowFile) {
		var ret FlowFileDTO
		return ret
	}
	return *o.FlowFile
}

// GetFlowFileOk returns a tuple with the FlowFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowFileEntity) GetFlowFileOk() (*FlowFileDTO, bool) {
	if o == nil || IsNil(o.FlowFile) {
		return nil, false
	}
	return o.FlowFile, true
}

// HasFlowFile returns a boolean if a field has been set.
func (o *FlowFileEntity) HasFlowFile() bool {
	if o != nil && !IsNil(o.FlowFile) {
		return true
	}

	return false
}

// SetFlowFile gets a reference to the given FlowFileDTO and assigns it to the FlowFile field.
func (o *FlowFileEntity) SetFlowFile(v FlowFileDTO) {
	o.FlowFile = &v
}

func (o FlowFileEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowFileEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowFile) {
		toSerialize["flowFile"] = o.FlowFile
	}
	return toSerialize, nil
}

type NullableFlowFileEntity struct {
	value *FlowFileEntity
	isSet bool
}

func (v NullableFlowFileEntity) Get() *FlowFileEntity {
	return v.value
}

func (v *NullableFlowFileEntity) Set(val *FlowFileEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowFileEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowFileEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowFileEntity(val *FlowFileEntity) *NullableFlowFileEntity {
	return &NullableFlowFileEntity{value: val, isSet: true}
}

func (v NullableFlowFileEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowFileEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

