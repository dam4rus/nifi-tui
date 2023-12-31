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

// checks if the BatchSettingsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchSettingsDTO{}

// BatchSettingsDTO struct for BatchSettingsDTO
type BatchSettingsDTO struct {
	// Preferred number of flow files to include in a transaction.
	Count *int32 `json:"count,omitempty"`
	// Preferred number of bytes to include in a transaction.
	Size *string `json:"size,omitempty"`
	// Preferred amount of time that a transaction should span.
	Duration *string `json:"duration,omitempty"`
}

// NewBatchSettingsDTO instantiates a new BatchSettingsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchSettingsDTO() *BatchSettingsDTO {
	this := BatchSettingsDTO{}
	return &this
}

// NewBatchSettingsDTOWithDefaults instantiates a new BatchSettingsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchSettingsDTOWithDefaults() *BatchSettingsDTO {
	this := BatchSettingsDTO{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *BatchSettingsDTO) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettingsDTO) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *BatchSettingsDTO) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *BatchSettingsDTO) SetCount(v int32) {
	o.Count = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BatchSettingsDTO) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettingsDTO) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BatchSettingsDTO) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *BatchSettingsDTO) SetSize(v string) {
	o.Size = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *BatchSettingsDTO) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchSettingsDTO) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *BatchSettingsDTO) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *BatchSettingsDTO) SetDuration(v string) {
	o.Duration = &v
}

func (o BatchSettingsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchSettingsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

type NullableBatchSettingsDTO struct {
	value *BatchSettingsDTO
	isSet bool
}

func (v NullableBatchSettingsDTO) Get() *BatchSettingsDTO {
	return v.value
}

func (v *NullableBatchSettingsDTO) Set(val *BatchSettingsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchSettingsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchSettingsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchSettingsDTO(val *BatchSettingsDTO) *NullableBatchSettingsDTO {
	return &NullableBatchSettingsDTO{value: val, isSet: true}
}

func (v NullableBatchSettingsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchSettingsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


