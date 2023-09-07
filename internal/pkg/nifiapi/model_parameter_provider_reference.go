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

// checks if the ParameterProviderReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterProviderReference{}

// ParameterProviderReference struct for ParameterProviderReference
type ParameterProviderReference struct {
	// The identifier of the parameter provider
	Identifier *string `json:"identifier,omitempty"`
	// The name of the parameter provider
	Name *string `json:"name,omitempty"`
	// The fully qualified name of the parameter provider class.
	Type *string `json:"type,omitempty"`
	Bundle *Bundle `json:"bundle,omitempty"`
}

// NewParameterProviderReference instantiates a new ParameterProviderReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterProviderReference() *ParameterProviderReference {
	this := ParameterProviderReference{}
	return &this
}

// NewParameterProviderReferenceWithDefaults instantiates a new ParameterProviderReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterProviderReferenceWithDefaults() *ParameterProviderReference {
	this := ParameterProviderReference{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ParameterProviderReference) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderReference) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ParameterProviderReference) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *ParameterProviderReference) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParameterProviderReference) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderReference) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParameterProviderReference) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParameterProviderReference) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ParameterProviderReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ParameterProviderReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ParameterProviderReference) SetType(v string) {
	o.Type = &v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *ParameterProviderReference) GetBundle() Bundle {
	if o == nil || IsNil(o.Bundle) {
		var ret Bundle
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterProviderReference) GetBundleOk() (*Bundle, bool) {
	if o == nil || IsNil(o.Bundle) {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *ParameterProviderReference) HasBundle() bool {
	if o != nil && !IsNil(o.Bundle) {
		return true
	}

	return false
}

// SetBundle gets a reference to the given Bundle and assigns it to the Bundle field.
func (o *ParameterProviderReference) SetBundle(v Bundle) {
	o.Bundle = &v
}

func (o ParameterProviderReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterProviderReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Bundle) {
		toSerialize["bundle"] = o.Bundle
	}
	return toSerialize, nil
}

type NullableParameterProviderReference struct {
	value *ParameterProviderReference
	isSet bool
}

func (v NullableParameterProviderReference) Get() *ParameterProviderReference {
	return v.value
}

func (v *NullableParameterProviderReference) Set(val *ParameterProviderReference) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterProviderReference) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterProviderReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterProviderReference(val *ParameterProviderReference) *NullableParameterProviderReference {
	return &NullableParameterProviderReference{value: val, isSet: true}
}

func (v NullableParameterProviderReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterProviderReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

