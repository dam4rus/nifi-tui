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

// checks if the LabelsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelsEntity{}

// LabelsEntity struct for LabelsEntity
type LabelsEntity struct {
	Labels []LabelEntity `json:"labels,omitempty"`
}

// NewLabelsEntity instantiates a new LabelsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsEntity() *LabelsEntity {
	this := LabelsEntity{}
	return &this
}

// NewLabelsEntityWithDefaults instantiates a new LabelsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsEntityWithDefaults() *LabelsEntity {
	this := LabelsEntity{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *LabelsEntity) GetLabels() []LabelEntity {
	if o == nil || IsNil(o.Labels) {
		var ret []LabelEntity
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsEntity) GetLabelsOk() ([]LabelEntity, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LabelsEntity) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelEntity and assigns it to the Labels field.
func (o *LabelsEntity) SetLabels(v []LabelEntity) {
	o.Labels = v
}

func (o LabelsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableLabelsEntity struct {
	value *LabelsEntity
	isSet bool
}

func (v NullableLabelsEntity) Get() *LabelsEntity {
	return v.value
}

func (v *NullableLabelsEntity) Set(val *LabelsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsEntity(val *LabelsEntity) *NullableLabelsEntity {
	return &NullableLabelsEntity{value: val, isSet: true}
}

func (v NullableLabelsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

