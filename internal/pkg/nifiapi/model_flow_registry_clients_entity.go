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

// checks if the FlowRegistryClientsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowRegistryClientsEntity{}

// FlowRegistryClientsEntity struct for FlowRegistryClientsEntity
type FlowRegistryClientsEntity struct {
	Registries []FlowRegistryClientEntity `json:"registries,omitempty"`
}

// NewFlowRegistryClientsEntity instantiates a new FlowRegistryClientsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowRegistryClientsEntity() *FlowRegistryClientsEntity {
	this := FlowRegistryClientsEntity{}
	return &this
}

// NewFlowRegistryClientsEntityWithDefaults instantiates a new FlowRegistryClientsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowRegistryClientsEntityWithDefaults() *FlowRegistryClientsEntity {
	this := FlowRegistryClientsEntity{}
	return &this
}

// GetRegistries returns the Registries field value if set, zero value otherwise.
func (o *FlowRegistryClientsEntity) GetRegistries() []FlowRegistryClientEntity {
	if o == nil || IsNil(o.Registries) {
		var ret []FlowRegistryClientEntity
		return ret
	}
	return o.Registries
}

// GetRegistriesOk returns a tuple with the Registries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowRegistryClientsEntity) GetRegistriesOk() ([]FlowRegistryClientEntity, bool) {
	if o == nil || IsNil(o.Registries) {
		return nil, false
	}
	return o.Registries, true
}

// HasRegistries returns a boolean if a field has been set.
func (o *FlowRegistryClientsEntity) HasRegistries() bool {
	if o != nil && !IsNil(o.Registries) {
		return true
	}

	return false
}

// SetRegistries gets a reference to the given []FlowRegistryClientEntity and assigns it to the Registries field.
func (o *FlowRegistryClientsEntity) SetRegistries(v []FlowRegistryClientEntity) {
	o.Registries = v
}

func (o FlowRegistryClientsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowRegistryClientsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Registries) {
		toSerialize["registries"] = o.Registries
	}
	return toSerialize, nil
}

type NullableFlowRegistryClientsEntity struct {
	value *FlowRegistryClientsEntity
	isSet bool
}

func (v NullableFlowRegistryClientsEntity) Get() *FlowRegistryClientsEntity {
	return v.value
}

func (v *NullableFlowRegistryClientsEntity) Set(val *FlowRegistryClientsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowRegistryClientsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowRegistryClientsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowRegistryClientsEntity(val *FlowRegistryClientsEntity) *NullableFlowRegistryClientsEntity {
	return &NullableFlowRegistryClientsEntity{value: val, isSet: true}
}

func (v NullableFlowRegistryClientsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowRegistryClientsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


