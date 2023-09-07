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

// checks if the ReplayLastEventSnapshotDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplayLastEventSnapshotDTO{}

// ReplayLastEventSnapshotDTO struct for ReplayLastEventSnapshotDTO
type ReplayLastEventSnapshotDTO struct {
	// The IDs of the events that were successfully replayed
	EventsReplayed []int64 `json:"eventsReplayed,omitempty"`
	// If unable to replay an event, specifies why the event could not be replayed
	FailureExplanation *string `json:"failureExplanation,omitempty"`
	// Whether or not an event was available. This may not be populated if there was a failure.
	EventAvailable *bool `json:"eventAvailable,omitempty"`
}

// NewReplayLastEventSnapshotDTO instantiates a new ReplayLastEventSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplayLastEventSnapshotDTO() *ReplayLastEventSnapshotDTO {
	this := ReplayLastEventSnapshotDTO{}
	return &this
}

// NewReplayLastEventSnapshotDTOWithDefaults instantiates a new ReplayLastEventSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplayLastEventSnapshotDTOWithDefaults() *ReplayLastEventSnapshotDTO {
	this := ReplayLastEventSnapshotDTO{}
	return &this
}

// GetEventsReplayed returns the EventsReplayed field value if set, zero value otherwise.
func (o *ReplayLastEventSnapshotDTO) GetEventsReplayed() []int64 {
	if o == nil || IsNil(o.EventsReplayed) {
		var ret []int64
		return ret
	}
	return o.EventsReplayed
}

// GetEventsReplayedOk returns a tuple with the EventsReplayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayLastEventSnapshotDTO) GetEventsReplayedOk() ([]int64, bool) {
	if o == nil || IsNil(o.EventsReplayed) {
		return nil, false
	}
	return o.EventsReplayed, true
}

// HasEventsReplayed returns a boolean if a field has been set.
func (o *ReplayLastEventSnapshotDTO) HasEventsReplayed() bool {
	if o != nil && !IsNil(o.EventsReplayed) {
		return true
	}

	return false
}

// SetEventsReplayed gets a reference to the given []int64 and assigns it to the EventsReplayed field.
func (o *ReplayLastEventSnapshotDTO) SetEventsReplayed(v []int64) {
	o.EventsReplayed = v
}

// GetFailureExplanation returns the FailureExplanation field value if set, zero value otherwise.
func (o *ReplayLastEventSnapshotDTO) GetFailureExplanation() string {
	if o == nil || IsNil(o.FailureExplanation) {
		var ret string
		return ret
	}
	return *o.FailureExplanation
}

// GetFailureExplanationOk returns a tuple with the FailureExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayLastEventSnapshotDTO) GetFailureExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.FailureExplanation) {
		return nil, false
	}
	return o.FailureExplanation, true
}

// HasFailureExplanation returns a boolean if a field has been set.
func (o *ReplayLastEventSnapshotDTO) HasFailureExplanation() bool {
	if o != nil && !IsNil(o.FailureExplanation) {
		return true
	}

	return false
}

// SetFailureExplanation gets a reference to the given string and assigns it to the FailureExplanation field.
func (o *ReplayLastEventSnapshotDTO) SetFailureExplanation(v string) {
	o.FailureExplanation = &v
}

// GetEventAvailable returns the EventAvailable field value if set, zero value otherwise.
func (o *ReplayLastEventSnapshotDTO) GetEventAvailable() bool {
	if o == nil || IsNil(o.EventAvailable) {
		var ret bool
		return ret
	}
	return *o.EventAvailable
}

// GetEventAvailableOk returns a tuple with the EventAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayLastEventSnapshotDTO) GetEventAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.EventAvailable) {
		return nil, false
	}
	return o.EventAvailable, true
}

// HasEventAvailable returns a boolean if a field has been set.
func (o *ReplayLastEventSnapshotDTO) HasEventAvailable() bool {
	if o != nil && !IsNil(o.EventAvailable) {
		return true
	}

	return false
}

// SetEventAvailable gets a reference to the given bool and assigns it to the EventAvailable field.
func (o *ReplayLastEventSnapshotDTO) SetEventAvailable(v bool) {
	o.EventAvailable = &v
}

func (o ReplayLastEventSnapshotDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplayLastEventSnapshotDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventsReplayed) {
		toSerialize["eventsReplayed"] = o.EventsReplayed
	}
	if !IsNil(o.FailureExplanation) {
		toSerialize["failureExplanation"] = o.FailureExplanation
	}
	if !IsNil(o.EventAvailable) {
		toSerialize["eventAvailable"] = o.EventAvailable
	}
	return toSerialize, nil
}

type NullableReplayLastEventSnapshotDTO struct {
	value *ReplayLastEventSnapshotDTO
	isSet bool
}

func (v NullableReplayLastEventSnapshotDTO) Get() *ReplayLastEventSnapshotDTO {
	return v.value
}

func (v *NullableReplayLastEventSnapshotDTO) Set(val *ReplayLastEventSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableReplayLastEventSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableReplayLastEventSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplayLastEventSnapshotDTO(val *ReplayLastEventSnapshotDTO) *NullableReplayLastEventSnapshotDTO {
	return &NullableReplayLastEventSnapshotDTO{value: val, isSet: true}
}

func (v NullableReplayLastEventSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplayLastEventSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


