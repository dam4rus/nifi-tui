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

// checks if the ProcessorStatusSnapshotEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessorStatusSnapshotEntity{}

// ProcessorStatusSnapshotEntity struct for ProcessorStatusSnapshotEntity
type ProcessorStatusSnapshotEntity struct {
	// The id of the processor.
	Id *string `json:"id,omitempty"`
	ProcessorStatusSnapshot *ProcessorStatusSnapshotDTO `json:"processorStatusSnapshot,omitempty"`
	// Indicates whether the user can read a given resource.
	CanRead *bool `json:"canRead,omitempty"`
}

// NewProcessorStatusSnapshotEntity instantiates a new ProcessorStatusSnapshotEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorStatusSnapshotEntity() *ProcessorStatusSnapshotEntity {
	this := ProcessorStatusSnapshotEntity{}
	return &this
}

// NewProcessorStatusSnapshotEntityWithDefaults instantiates a new ProcessorStatusSnapshotEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorStatusSnapshotEntityWithDefaults() *ProcessorStatusSnapshotEntity {
	this := ProcessorStatusSnapshotEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessorStatusSnapshotEntity) SetId(v string) {
	o.Id = &v
}

// GetProcessorStatusSnapshot returns the ProcessorStatusSnapshot field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotEntity) GetProcessorStatusSnapshot() ProcessorStatusSnapshotDTO {
	if o == nil || IsNil(o.ProcessorStatusSnapshot) {
		var ret ProcessorStatusSnapshotDTO
		return ret
	}
	return *o.ProcessorStatusSnapshot
}

// GetProcessorStatusSnapshotOk returns a tuple with the ProcessorStatusSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotEntity) GetProcessorStatusSnapshotOk() (*ProcessorStatusSnapshotDTO, bool) {
	if o == nil || IsNil(o.ProcessorStatusSnapshot) {
		return nil, false
	}
	return o.ProcessorStatusSnapshot, true
}

// HasProcessorStatusSnapshot returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotEntity) HasProcessorStatusSnapshot() bool {
	if o != nil && !IsNil(o.ProcessorStatusSnapshot) {
		return true
	}

	return false
}

// SetProcessorStatusSnapshot gets a reference to the given ProcessorStatusSnapshotDTO and assigns it to the ProcessorStatusSnapshot field.
func (o *ProcessorStatusSnapshotEntity) SetProcessorStatusSnapshot(v ProcessorStatusSnapshotDTO) {
	o.ProcessorStatusSnapshot = &v
}

// GetCanRead returns the CanRead field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotEntity) GetCanRead() bool {
	if o == nil || IsNil(o.CanRead) {
		var ret bool
		return ret
	}
	return *o.CanRead
}

// GetCanReadOk returns a tuple with the CanRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotEntity) GetCanReadOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRead) {
		return nil, false
	}
	return o.CanRead, true
}

// HasCanRead returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotEntity) HasCanRead() bool {
	if o != nil && !IsNil(o.CanRead) {
		return true
	}

	return false
}

// SetCanRead gets a reference to the given bool and assigns it to the CanRead field.
func (o *ProcessorStatusSnapshotEntity) SetCanRead(v bool) {
	o.CanRead = &v
}

func (o ProcessorStatusSnapshotEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessorStatusSnapshotEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProcessorStatusSnapshot) {
		toSerialize["processorStatusSnapshot"] = o.ProcessorStatusSnapshot
	}
	if !IsNil(o.CanRead) {
		toSerialize["canRead"] = o.CanRead
	}
	return toSerialize, nil
}

type NullableProcessorStatusSnapshotEntity struct {
	value *ProcessorStatusSnapshotEntity
	isSet bool
}

func (v NullableProcessorStatusSnapshotEntity) Get() *ProcessorStatusSnapshotEntity {
	return v.value
}

func (v *NullableProcessorStatusSnapshotEntity) Set(val *ProcessorStatusSnapshotEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorStatusSnapshotEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorStatusSnapshotEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorStatusSnapshotEntity(val *ProcessorStatusSnapshotEntity) *NullableProcessorStatusSnapshotEntity {
	return &NullableProcessorStatusSnapshotEntity{value: val, isSet: true}
}

func (v NullableProcessorStatusSnapshotEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorStatusSnapshotEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


