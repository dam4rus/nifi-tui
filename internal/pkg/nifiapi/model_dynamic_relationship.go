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

// checks if the DynamicRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicRelationship{}

// DynamicRelationship struct for DynamicRelationship
type DynamicRelationship struct {
	// The description of the dynamic relationship name
	Name *string `json:"name,omitempty"`
	// The description of the dynamic relationship
	Description *string `json:"description,omitempty"`
}

// NewDynamicRelationship instantiates a new DynamicRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicRelationship() *DynamicRelationship {
	this := DynamicRelationship{}
	return &this
}

// NewDynamicRelationshipWithDefaults instantiates a new DynamicRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicRelationshipWithDefaults() *DynamicRelationship {
	this := DynamicRelationship{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DynamicRelationship) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicRelationship) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DynamicRelationship) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DynamicRelationship) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicRelationship) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicRelationship) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicRelationship) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicRelationship) SetDescription(v string) {
	o.Description = &v
}

func (o DynamicRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableDynamicRelationship struct {
	value *DynamicRelationship
	isSet bool
}

func (v NullableDynamicRelationship) Get() *DynamicRelationship {
	return v.value
}

func (v *NullableDynamicRelationship) Set(val *DynamicRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicRelationship(val *DynamicRelationship) *NullableDynamicRelationship {
	return &NullableDynamicRelationship{value: val, isSet: true}
}

func (v NullableDynamicRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

