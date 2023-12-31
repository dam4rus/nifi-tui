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

// checks if the OutputPortsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputPortsEntity{}

// OutputPortsEntity struct for OutputPortsEntity
type OutputPortsEntity struct {
	OutputPorts []PortEntity `json:"outputPorts,omitempty"`
}

// NewOutputPortsEntity instantiates a new OutputPortsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputPortsEntity() *OutputPortsEntity {
	this := OutputPortsEntity{}
	return &this
}

// NewOutputPortsEntityWithDefaults instantiates a new OutputPortsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputPortsEntityWithDefaults() *OutputPortsEntity {
	this := OutputPortsEntity{}
	return &this
}

// GetOutputPorts returns the OutputPorts field value if set, zero value otherwise.
func (o *OutputPortsEntity) GetOutputPorts() []PortEntity {
	if o == nil || IsNil(o.OutputPorts) {
		var ret []PortEntity
		return ret
	}
	return o.OutputPorts
}

// GetOutputPortsOk returns a tuple with the OutputPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputPortsEntity) GetOutputPortsOk() ([]PortEntity, bool) {
	if o == nil || IsNil(o.OutputPorts) {
		return nil, false
	}
	return o.OutputPorts, true
}

// HasOutputPorts returns a boolean if a field has been set.
func (o *OutputPortsEntity) HasOutputPorts() bool {
	if o != nil && !IsNil(o.OutputPorts) {
		return true
	}

	return false
}

// SetOutputPorts gets a reference to the given []PortEntity and assigns it to the OutputPorts field.
func (o *OutputPortsEntity) SetOutputPorts(v []PortEntity) {
	o.OutputPorts = v
}

func (o OutputPortsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputPortsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OutputPorts) {
		toSerialize["outputPorts"] = o.OutputPorts
	}
	return toSerialize, nil
}

type NullableOutputPortsEntity struct {
	value *OutputPortsEntity
	isSet bool
}

func (v NullableOutputPortsEntity) Get() *OutputPortsEntity {
	return v.value
}

func (v *NullableOutputPortsEntity) Set(val *OutputPortsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputPortsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputPortsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputPortsEntity(val *OutputPortsEntity) *NullableOutputPortsEntity {
	return &NullableOutputPortsEntity{value: val, isSet: true}
}

func (v NullableOutputPortsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputPortsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


