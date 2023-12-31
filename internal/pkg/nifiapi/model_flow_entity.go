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

// checks if the FlowEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowEntity{}

// FlowEntity struct for FlowEntity
type FlowEntity struct {
	Flow *FlowDTO `json:"flow,omitempty"`
}

// NewFlowEntity instantiates a new FlowEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowEntity() *FlowEntity {
	this := FlowEntity{}
	return &this
}

// NewFlowEntityWithDefaults instantiates a new FlowEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowEntityWithDefaults() *FlowEntity {
	this := FlowEntity{}
	return &this
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *FlowEntity) GetFlow() FlowDTO {
	if o == nil || IsNil(o.Flow) {
		var ret FlowDTO
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowEntity) GetFlowOk() (*FlowDTO, bool) {
	if o == nil || IsNil(o.Flow) {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *FlowEntity) HasFlow() bool {
	if o != nil && !IsNil(o.Flow) {
		return true
	}

	return false
}

// SetFlow gets a reference to the given FlowDTO and assigns it to the Flow field.
func (o *FlowEntity) SetFlow(v FlowDTO) {
	o.Flow = &v
}

func (o FlowEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flow) {
		toSerialize["flow"] = o.Flow
	}
	return toSerialize, nil
}

type NullableFlowEntity struct {
	value *FlowEntity
	isSet bool
}

func (v NullableFlowEntity) Get() *FlowEntity {
	return v.value
}

func (v *NullableFlowEntity) Set(val *FlowEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowEntity(val *FlowEntity) *NullableFlowEntity {
	return &NullableFlowEntity{value: val, isSet: true}
}

func (v NullableFlowEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


