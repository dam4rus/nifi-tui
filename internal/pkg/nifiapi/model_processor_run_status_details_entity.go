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

// checks if the ProcessorRunStatusDetailsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessorRunStatusDetailsEntity{}

// ProcessorRunStatusDetailsEntity struct for ProcessorRunStatusDetailsEntity
type ProcessorRunStatusDetailsEntity struct {
	Revision *RevisionDTO `json:"revision,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
	RunStatusDetails *ProcessorRunStatusDetailsDTO `json:"runStatusDetails,omitempty"`
}

// NewProcessorRunStatusDetailsEntity instantiates a new ProcessorRunStatusDetailsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorRunStatusDetailsEntity() *ProcessorRunStatusDetailsEntity {
	this := ProcessorRunStatusDetailsEntity{}
	return &this
}

// NewProcessorRunStatusDetailsEntityWithDefaults instantiates a new ProcessorRunStatusDetailsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorRunStatusDetailsEntityWithDefaults() *ProcessorRunStatusDetailsEntity {
	this := ProcessorRunStatusDetailsEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *ProcessorRunStatusDetailsEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ProcessorRunStatusDetailsEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetRunStatusDetails returns the RunStatusDetails field value if set, zero value otherwise.
func (o *ProcessorRunStatusDetailsEntity) GetRunStatusDetails() ProcessorRunStatusDetailsDTO {
	if o == nil || IsNil(o.RunStatusDetails) {
		var ret ProcessorRunStatusDetailsDTO
		return ret
	}
	return *o.RunStatusDetails
}

// GetRunStatusDetailsOk returns a tuple with the RunStatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorRunStatusDetailsEntity) GetRunStatusDetailsOk() (*ProcessorRunStatusDetailsDTO, bool) {
	if o == nil || IsNil(o.RunStatusDetails) {
		return nil, false
	}
	return o.RunStatusDetails, true
}

// HasRunStatusDetails returns a boolean if a field has been set.
func (o *ProcessorRunStatusDetailsEntity) HasRunStatusDetails() bool {
	if o != nil && !IsNil(o.RunStatusDetails) {
		return true
	}

	return false
}

// SetRunStatusDetails gets a reference to the given ProcessorRunStatusDetailsDTO and assigns it to the RunStatusDetails field.
func (o *ProcessorRunStatusDetailsEntity) SetRunStatusDetails(v ProcessorRunStatusDetailsDTO) {
	o.RunStatusDetails = &v
}

func (o ProcessorRunStatusDetailsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessorRunStatusDetailsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.RunStatusDetails) {
		toSerialize["runStatusDetails"] = o.RunStatusDetails
	}
	return toSerialize, nil
}

type NullableProcessorRunStatusDetailsEntity struct {
	value *ProcessorRunStatusDetailsEntity
	isSet bool
}

func (v NullableProcessorRunStatusDetailsEntity) Get() *ProcessorRunStatusDetailsEntity {
	return v.value
}

func (v *NullableProcessorRunStatusDetailsEntity) Set(val *ProcessorRunStatusDetailsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorRunStatusDetailsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorRunStatusDetailsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorRunStatusDetailsEntity(val *ProcessorRunStatusDetailsEntity) *NullableProcessorRunStatusDetailsEntity {
	return &NullableProcessorRunStatusDetailsEntity{value: val, isSet: true}
}

func (v NullableProcessorRunStatusDetailsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorRunStatusDetailsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

