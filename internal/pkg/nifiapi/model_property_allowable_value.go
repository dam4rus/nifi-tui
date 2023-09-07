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

// checks if the PropertyAllowableValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyAllowableValue{}

// PropertyAllowableValue struct for PropertyAllowableValue
type PropertyAllowableValue struct {
	// The internal value
	Value string `json:"value"`
	// The display name of the value, if different from the internal value
	DisplayName *string `json:"displayName,omitempty"`
	// The description of the value, e.g., the behavior it produces.
	Description *string `json:"description,omitempty"`
}

// NewPropertyAllowableValue instantiates a new PropertyAllowableValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyAllowableValue(value string) *PropertyAllowableValue {
	this := PropertyAllowableValue{}
	this.Value = value
	return &this
}

// NewPropertyAllowableValueWithDefaults instantiates a new PropertyAllowableValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyAllowableValueWithDefaults() *PropertyAllowableValue {
	this := PropertyAllowableValue{}
	return &this
}

// GetValue returns the Value field value
func (o *PropertyAllowableValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PropertyAllowableValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PropertyAllowableValue) SetValue(v string) {
	o.Value = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PropertyAllowableValue) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyAllowableValue) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PropertyAllowableValue) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PropertyAllowableValue) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropertyAllowableValue) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyAllowableValue) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyAllowableValue) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropertyAllowableValue) SetDescription(v string) {
	o.Description = &v
}

func (o PropertyAllowableValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyAllowableValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullablePropertyAllowableValue struct {
	value *PropertyAllowableValue
	isSet bool
}

func (v NullablePropertyAllowableValue) Get() *PropertyAllowableValue {
	return v.value
}

func (v *NullablePropertyAllowableValue) Set(val *PropertyAllowableValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyAllowableValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyAllowableValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyAllowableValue(val *PropertyAllowableValue) *NullablePropertyAllowableValue {
	return &NullablePropertyAllowableValue{value: val, isSet: true}
}

func (v NullablePropertyAllowableValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyAllowableValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


