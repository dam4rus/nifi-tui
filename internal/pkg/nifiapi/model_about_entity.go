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

// checks if the AboutEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AboutEntity{}

// AboutEntity struct for AboutEntity
type AboutEntity struct {
	About *AboutDTO `json:"about,omitempty"`
}

// NewAboutEntity instantiates a new AboutEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAboutEntity() *AboutEntity {
	this := AboutEntity{}
	return &this
}

// NewAboutEntityWithDefaults instantiates a new AboutEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAboutEntityWithDefaults() *AboutEntity {
	this := AboutEntity{}
	return &this
}

// GetAbout returns the About field value if set, zero value otherwise.
func (o *AboutEntity) GetAbout() AboutDTO {
	if o == nil || IsNil(o.About) {
		var ret AboutDTO
		return ret
	}
	return *o.About
}

// GetAboutOk returns a tuple with the About field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutEntity) GetAboutOk() (*AboutDTO, bool) {
	if o == nil || IsNil(o.About) {
		return nil, false
	}
	return o.About, true
}

// HasAbout returns a boolean if a field has been set.
func (o *AboutEntity) HasAbout() bool {
	if o != nil && !IsNil(o.About) {
		return true
	}

	return false
}

// SetAbout gets a reference to the given AboutDTO and assigns it to the About field.
func (o *AboutEntity) SetAbout(v AboutDTO) {
	o.About = &v
}

func (o AboutEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AboutEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.About) {
		toSerialize["about"] = o.About
	}
	return toSerialize, nil
}

type NullableAboutEntity struct {
	value *AboutEntity
	isSet bool
}

func (v NullableAboutEntity) Get() *AboutEntity {
	return v.value
}

func (v *NullableAboutEntity) Set(val *AboutEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAboutEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAboutEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAboutEntity(val *AboutEntity) *NullableAboutEntity {
	return &NullableAboutEntity{value: val, isSet: true}
}

func (v NullableAboutEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAboutEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

