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

// checks if the SnippetEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnippetEntity{}

// SnippetEntity struct for SnippetEntity
type SnippetEntity struct {
	Snippet *SnippetDTO `json:"snippet,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewSnippetEntity instantiates a new SnippetEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippetEntity() *SnippetEntity {
	this := SnippetEntity{}
	return &this
}

// NewSnippetEntityWithDefaults instantiates a new SnippetEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetEntityWithDefaults() *SnippetEntity {
	this := SnippetEntity{}
	return &this
}

// GetSnippet returns the Snippet field value if set, zero value otherwise.
func (o *SnippetEntity) GetSnippet() SnippetDTO {
	if o == nil || IsNil(o.Snippet) {
		var ret SnippetDTO
		return ret
	}
	return *o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetEntity) GetSnippetOk() (*SnippetDTO, bool) {
	if o == nil || IsNil(o.Snippet) {
		return nil, false
	}
	return o.Snippet, true
}

// HasSnippet returns a boolean if a field has been set.
func (o *SnippetEntity) HasSnippet() bool {
	if o != nil && !IsNil(o.Snippet) {
		return true
	}

	return false
}

// SetSnippet gets a reference to the given SnippetDTO and assigns it to the Snippet field.
func (o *SnippetEntity) SetSnippet(v SnippetDTO) {
	o.Snippet = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *SnippetEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *SnippetEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *SnippetEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o SnippetEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnippetEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Snippet) {
		toSerialize["snippet"] = o.Snippet
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return toSerialize, nil
}

type NullableSnippetEntity struct {
	value *SnippetEntity
	isSet bool
}

func (v NullableSnippetEntity) Get() *SnippetEntity {
	return v.value
}

func (v *NullableSnippetEntity) Set(val *SnippetEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableSnippetEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableSnippetEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnippetEntity(val *SnippetEntity) *NullableSnippetEntity {
	return &NullableSnippetEntity{value: val, isSet: true}
}

func (v NullableSnippetEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnippetEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


