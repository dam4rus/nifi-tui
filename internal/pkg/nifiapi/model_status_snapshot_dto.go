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
	"time"
)

// checks if the StatusSnapshotDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusSnapshotDTO{}

// StatusSnapshotDTO struct for StatusSnapshotDTO
type StatusSnapshotDTO struct {
	// The timestamp of the snapshot.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The status metrics.
	StatusMetrics *map[string]int64 `json:"statusMetrics,omitempty"`
}

// NewStatusSnapshotDTO instantiates a new StatusSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusSnapshotDTO() *StatusSnapshotDTO {
	this := StatusSnapshotDTO{}
	return &this
}

// NewStatusSnapshotDTOWithDefaults instantiates a new StatusSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusSnapshotDTOWithDefaults() *StatusSnapshotDTO {
	this := StatusSnapshotDTO{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *StatusSnapshotDTO) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSnapshotDTO) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *StatusSnapshotDTO) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *StatusSnapshotDTO) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetStatusMetrics returns the StatusMetrics field value if set, zero value otherwise.
func (o *StatusSnapshotDTO) GetStatusMetrics() map[string]int64 {
	if o == nil || IsNil(o.StatusMetrics) {
		var ret map[string]int64
		return ret
	}
	return *o.StatusMetrics
}

// GetStatusMetricsOk returns a tuple with the StatusMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSnapshotDTO) GetStatusMetricsOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.StatusMetrics) {
		return nil, false
	}
	return o.StatusMetrics, true
}

// HasStatusMetrics returns a boolean if a field has been set.
func (o *StatusSnapshotDTO) HasStatusMetrics() bool {
	if o != nil && !IsNil(o.StatusMetrics) {
		return true
	}

	return false
}

// SetStatusMetrics gets a reference to the given map[string]int64 and assigns it to the StatusMetrics field.
func (o *StatusSnapshotDTO) SetStatusMetrics(v map[string]int64) {
	o.StatusMetrics = &v
}

func (o StatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusSnapshotDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.StatusMetrics) {
		toSerialize["statusMetrics"] = o.StatusMetrics
	}
	return toSerialize, nil
}

type NullableStatusSnapshotDTO struct {
	value *StatusSnapshotDTO
	isSet bool
}

func (v NullableStatusSnapshotDTO) Get() *StatusSnapshotDTO {
	return v.value
}

func (v *NullableStatusSnapshotDTO) Set(val *StatusSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusSnapshotDTO(val *StatusSnapshotDTO) *NullableStatusSnapshotDTO {
	return &NullableStatusSnapshotDTO{value: val, isSet: true}
}

func (v NullableStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


