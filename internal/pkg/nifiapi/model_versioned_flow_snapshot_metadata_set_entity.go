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

// checks if the VersionedFlowSnapshotMetadataSetEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionedFlowSnapshotMetadataSetEntity{}

// VersionedFlowSnapshotMetadataSetEntity struct for VersionedFlowSnapshotMetadataSetEntity
type VersionedFlowSnapshotMetadataSetEntity struct {
	VersionedFlowSnapshotMetadataSet []VersionedFlowSnapshotMetadataEntity `json:"versionedFlowSnapshotMetadataSet,omitempty"`
}

// NewVersionedFlowSnapshotMetadataSetEntity instantiates a new VersionedFlowSnapshotMetadataSetEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedFlowSnapshotMetadataSetEntity() *VersionedFlowSnapshotMetadataSetEntity {
	this := VersionedFlowSnapshotMetadataSetEntity{}
	return &this
}

// NewVersionedFlowSnapshotMetadataSetEntityWithDefaults instantiates a new VersionedFlowSnapshotMetadataSetEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedFlowSnapshotMetadataSetEntityWithDefaults() *VersionedFlowSnapshotMetadataSetEntity {
	this := VersionedFlowSnapshotMetadataSetEntity{}
	return &this
}

// GetVersionedFlowSnapshotMetadataSet returns the VersionedFlowSnapshotMetadataSet field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotMetadataSetEntity) GetVersionedFlowSnapshotMetadataSet() []VersionedFlowSnapshotMetadataEntity {
	if o == nil || IsNil(o.VersionedFlowSnapshotMetadataSet) {
		var ret []VersionedFlowSnapshotMetadataEntity
		return ret
	}
	return o.VersionedFlowSnapshotMetadataSet
}

// GetVersionedFlowSnapshotMetadataSetOk returns a tuple with the VersionedFlowSnapshotMetadataSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotMetadataSetEntity) GetVersionedFlowSnapshotMetadataSetOk() ([]VersionedFlowSnapshotMetadataEntity, bool) {
	if o == nil || IsNil(o.VersionedFlowSnapshotMetadataSet) {
		return nil, false
	}
	return o.VersionedFlowSnapshotMetadataSet, true
}

// HasVersionedFlowSnapshotMetadataSet returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotMetadataSetEntity) HasVersionedFlowSnapshotMetadataSet() bool {
	if o != nil && !IsNil(o.VersionedFlowSnapshotMetadataSet) {
		return true
	}

	return false
}

// SetVersionedFlowSnapshotMetadataSet gets a reference to the given []VersionedFlowSnapshotMetadataEntity and assigns it to the VersionedFlowSnapshotMetadataSet field.
func (o *VersionedFlowSnapshotMetadataSetEntity) SetVersionedFlowSnapshotMetadataSet(v []VersionedFlowSnapshotMetadataEntity) {
	o.VersionedFlowSnapshotMetadataSet = v
}

func (o VersionedFlowSnapshotMetadataSetEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionedFlowSnapshotMetadataSetEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionedFlowSnapshotMetadataSet) {
		toSerialize["versionedFlowSnapshotMetadataSet"] = o.VersionedFlowSnapshotMetadataSet
	}
	return toSerialize, nil
}

type NullableVersionedFlowSnapshotMetadataSetEntity struct {
	value *VersionedFlowSnapshotMetadataSetEntity
	isSet bool
}

func (v NullableVersionedFlowSnapshotMetadataSetEntity) Get() *VersionedFlowSnapshotMetadataSetEntity {
	return v.value
}

func (v *NullableVersionedFlowSnapshotMetadataSetEntity) Set(val *VersionedFlowSnapshotMetadataSetEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedFlowSnapshotMetadataSetEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedFlowSnapshotMetadataSetEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedFlowSnapshotMetadataSetEntity(val *VersionedFlowSnapshotMetadataSetEntity) *NullableVersionedFlowSnapshotMetadataSetEntity {
	return &NullableVersionedFlowSnapshotMetadataSetEntity{value: val, isSet: true}
}

func (v NullableVersionedFlowSnapshotMetadataSetEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedFlowSnapshotMetadataSetEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


