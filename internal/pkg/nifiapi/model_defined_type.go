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

// checks if the DefinedType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefinedType{}

// DefinedType struct for DefinedType
type DefinedType struct {
	// The group name of the bundle that provides the referenced type.
	Group *string `json:"group,omitempty"`
	// The artifact name of the bundle that provides the referenced type.
	Artifact *string `json:"artifact,omitempty"`
	// The version of the bundle that provides the referenced type.
	Version *string `json:"version,omitempty"`
	// The fully-qualified class type
	Type string `json:"type"`
	// The description of the type.
	TypeDescription *string `json:"typeDescription,omitempty"`
}

// NewDefinedType instantiates a new DefinedType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefinedType(type_ string) *DefinedType {
	this := DefinedType{}
	this.Type = type_
	return &this
}

// NewDefinedTypeWithDefaults instantiates a new DefinedType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefinedTypeWithDefaults() *DefinedType {
	this := DefinedType{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *DefinedType) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefinedType) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *DefinedType) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *DefinedType) SetGroup(v string) {
	o.Group = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *DefinedType) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefinedType) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *DefinedType) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *DefinedType) SetArtifact(v string) {
	o.Artifact = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DefinedType) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefinedType) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DefinedType) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DefinedType) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value
func (o *DefinedType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DefinedType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DefinedType) SetType(v string) {
	o.Type = v
}

// GetTypeDescription returns the TypeDescription field value if set, zero value otherwise.
func (o *DefinedType) GetTypeDescription() string {
	if o == nil || IsNil(o.TypeDescription) {
		var ret string
		return ret
	}
	return *o.TypeDescription
}

// GetTypeDescriptionOk returns a tuple with the TypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefinedType) GetTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TypeDescription) {
		return nil, false
	}
	return o.TypeDescription, true
}

// HasTypeDescription returns a boolean if a field has been set.
func (o *DefinedType) HasTypeDescription() bool {
	if o != nil && !IsNil(o.TypeDescription) {
		return true
	}

	return false
}

// SetTypeDescription gets a reference to the given string and assigns it to the TypeDescription field.
func (o *DefinedType) SetTypeDescription(v string) {
	o.TypeDescription = &v
}

func (o DefinedType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefinedType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.TypeDescription) {
		toSerialize["typeDescription"] = o.TypeDescription
	}
	return toSerialize, nil
}

type NullableDefinedType struct {
	value *DefinedType
	isSet bool
}

func (v NullableDefinedType) Get() *DefinedType {
	return v.value
}

func (v *NullableDefinedType) Set(val *DefinedType) {
	v.value = val
	v.isSet = true
}

func (v NullableDefinedType) IsSet() bool {
	return v.isSet
}

func (v *NullableDefinedType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefinedType(val *DefinedType) *NullableDefinedType {
	return &NullableDefinedType{value: val, isSet: true}
}

func (v NullableDefinedType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefinedType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


