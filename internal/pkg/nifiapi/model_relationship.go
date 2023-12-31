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

// checks if the Relationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Relationship{}

// Relationship struct for Relationship
type Relationship struct {
	// The name of the relationship
	Name *string `json:"name,omitempty"`
	// The description of the relationship
	Description *string `json:"description,omitempty"`
}

// NewRelationship instantiates a new Relationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationship() *Relationship {
	this := Relationship{}
	return &this
}

// NewRelationshipWithDefaults instantiates a new Relationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWithDefaults() *Relationship {
	this := Relationship{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Relationship) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Relationship) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Relationship) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Relationship) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Relationship) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Relationship) SetDescription(v string) {
	o.Description = &v
}

func (o Relationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Relationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableRelationship struct {
	value *Relationship
	isSet bool
}

func (v NullableRelationship) Get() *Relationship {
	return v.value
}

func (v *NullableRelationship) Set(val *Relationship) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationship(val *Relationship) *NullableRelationship {
	return &NullableRelationship{value: val, isSet: true}
}

func (v NullableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


