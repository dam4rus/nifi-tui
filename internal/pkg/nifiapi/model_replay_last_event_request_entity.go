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

// checks if the ReplayLastEventRequestEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplayLastEventRequestEntity{}

// ReplayLastEventRequestEntity struct for ReplayLastEventRequestEntity
type ReplayLastEventRequestEntity struct {
	// The UUID of the component whose last event should be replayed.
	ComponentId *string `json:"componentId,omitempty"`
	// Which nodes are to replay their last provenance event.
	Nodes *string `json:"nodes,omitempty"`
}

// NewReplayLastEventRequestEntity instantiates a new ReplayLastEventRequestEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplayLastEventRequestEntity() *ReplayLastEventRequestEntity {
	this := ReplayLastEventRequestEntity{}
	return &this
}

// NewReplayLastEventRequestEntityWithDefaults instantiates a new ReplayLastEventRequestEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplayLastEventRequestEntityWithDefaults() *ReplayLastEventRequestEntity {
	this := ReplayLastEventRequestEntity{}
	return &this
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise.
func (o *ReplayLastEventRequestEntity) GetComponentId() string {
	if o == nil || IsNil(o.ComponentId) {
		var ret string
		return ret
	}
	return *o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayLastEventRequestEntity) GetComponentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentId) {
		return nil, false
	}
	return o.ComponentId, true
}

// HasComponentId returns a boolean if a field has been set.
func (o *ReplayLastEventRequestEntity) HasComponentId() bool {
	if o != nil && !IsNil(o.ComponentId) {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given string and assigns it to the ComponentId field.
func (o *ReplayLastEventRequestEntity) SetComponentId(v string) {
	o.ComponentId = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ReplayLastEventRequestEntity) GetNodes() string {
	if o == nil || IsNil(o.Nodes) {
		var ret string
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplayLastEventRequestEntity) GetNodesOk() (*string, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ReplayLastEventRequestEntity) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given string and assigns it to the Nodes field.
func (o *ReplayLastEventRequestEntity) SetNodes(v string) {
	o.Nodes = &v
}

func (o ReplayLastEventRequestEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplayLastEventRequestEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentId) {
		toSerialize["componentId"] = o.ComponentId
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	return toSerialize, nil
}

type NullableReplayLastEventRequestEntity struct {
	value *ReplayLastEventRequestEntity
	isSet bool
}

func (v NullableReplayLastEventRequestEntity) Get() *ReplayLastEventRequestEntity {
	return v.value
}

func (v *NullableReplayLastEventRequestEntity) Set(val *ReplayLastEventRequestEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableReplayLastEventRequestEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableReplayLastEventRequestEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplayLastEventRequestEntity(val *ReplayLastEventRequestEntity) *NullableReplayLastEventRequestEntity {
	return &NullableReplayLastEventRequestEntity{value: val, isSet: true}
}

func (v NullableReplayLastEventRequestEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplayLastEventRequestEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


