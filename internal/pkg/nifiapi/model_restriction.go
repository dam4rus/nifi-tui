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

// checks if the Restriction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Restriction{}

// Restriction struct for Restriction
type Restriction struct {
	// The permission required for this restriction
	RequiredPermission *string `json:"requiredPermission,omitempty"`
	// The explanation of this restriction
	Explanation *string `json:"explanation,omitempty"`
}

// NewRestriction instantiates a new Restriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestriction() *Restriction {
	this := Restriction{}
	return &this
}

// NewRestrictionWithDefaults instantiates a new Restriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictionWithDefaults() *Restriction {
	this := Restriction{}
	return &this
}

// GetRequiredPermission returns the RequiredPermission field value if set, zero value otherwise.
func (o *Restriction) GetRequiredPermission() string {
	if o == nil || IsNil(o.RequiredPermission) {
		var ret string
		return ret
	}
	return *o.RequiredPermission
}

// GetRequiredPermissionOk returns a tuple with the RequiredPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restriction) GetRequiredPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.RequiredPermission) {
		return nil, false
	}
	return o.RequiredPermission, true
}

// HasRequiredPermission returns a boolean if a field has been set.
func (o *Restriction) HasRequiredPermission() bool {
	if o != nil && !IsNil(o.RequiredPermission) {
		return true
	}

	return false
}

// SetRequiredPermission gets a reference to the given string and assigns it to the RequiredPermission field.
func (o *Restriction) SetRequiredPermission(v string) {
	o.RequiredPermission = &v
}

// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *Restriction) GetExplanation() string {
	if o == nil || IsNil(o.Explanation) {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restriction) GetExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.Explanation) {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *Restriction) HasExplanation() bool {
	if o != nil && !IsNil(o.Explanation) {
		return true
	}

	return false
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *Restriction) SetExplanation(v string) {
	o.Explanation = &v
}

func (o Restriction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Restriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequiredPermission) {
		toSerialize["requiredPermission"] = o.RequiredPermission
	}
	if !IsNil(o.Explanation) {
		toSerialize["explanation"] = o.Explanation
	}
	return toSerialize, nil
}

type NullableRestriction struct {
	value *Restriction
	isSet bool
}

func (v NullableRestriction) Get() *Restriction {
	return v.value
}

func (v *NullableRestriction) Set(val *Restriction) {
	v.value = val
	v.isSet = true
}

func (v NullableRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestriction(val *Restriction) *NullableRestriction {
	return &NullableRestriction{value: val, isSet: true}
}

func (v NullableRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


