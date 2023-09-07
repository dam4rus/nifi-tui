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

// checks if the RuntimeManifestEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuntimeManifestEntity{}

// RuntimeManifestEntity struct for RuntimeManifestEntity
type RuntimeManifestEntity struct {
	RuntimeManifest *RuntimeManifest `json:"runtimeManifest,omitempty"`
}

// NewRuntimeManifestEntity instantiates a new RuntimeManifestEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeManifestEntity() *RuntimeManifestEntity {
	this := RuntimeManifestEntity{}
	return &this
}

// NewRuntimeManifestEntityWithDefaults instantiates a new RuntimeManifestEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeManifestEntityWithDefaults() *RuntimeManifestEntity {
	this := RuntimeManifestEntity{}
	return &this
}

// GetRuntimeManifest returns the RuntimeManifest field value if set, zero value otherwise.
func (o *RuntimeManifestEntity) GetRuntimeManifest() RuntimeManifest {
	if o == nil || IsNil(o.RuntimeManifest) {
		var ret RuntimeManifest
		return ret
	}
	return *o.RuntimeManifest
}

// GetRuntimeManifestOk returns a tuple with the RuntimeManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuntimeManifestEntity) GetRuntimeManifestOk() (*RuntimeManifest, bool) {
	if o == nil || IsNil(o.RuntimeManifest) {
		return nil, false
	}
	return o.RuntimeManifest, true
}

// HasRuntimeManifest returns a boolean if a field has been set.
func (o *RuntimeManifestEntity) HasRuntimeManifest() bool {
	if o != nil && !IsNil(o.RuntimeManifest) {
		return true
	}

	return false
}

// SetRuntimeManifest gets a reference to the given RuntimeManifest and assigns it to the RuntimeManifest field.
func (o *RuntimeManifestEntity) SetRuntimeManifest(v RuntimeManifest) {
	o.RuntimeManifest = &v
}

func (o RuntimeManifestEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuntimeManifestEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuntimeManifest) {
		toSerialize["runtimeManifest"] = o.RuntimeManifest
	}
	return toSerialize, nil
}

type NullableRuntimeManifestEntity struct {
	value *RuntimeManifestEntity
	isSet bool
}

func (v NullableRuntimeManifestEntity) Get() *RuntimeManifestEntity {
	return v.value
}

func (v *NullableRuntimeManifestEntity) Set(val *RuntimeManifestEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeManifestEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeManifestEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeManifestEntity(val *RuntimeManifestEntity) *NullableRuntimeManifestEntity {
	return &NullableRuntimeManifestEntity{value: val, isSet: true}
}

func (v NullableRuntimeManifestEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeManifestEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


