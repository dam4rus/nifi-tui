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

// checks if the PeersEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeersEntity{}

// PeersEntity struct for PeersEntity
type PeersEntity struct {
	Peers []PeerDTO `json:"peers,omitempty"`
}

// NewPeersEntity instantiates a new PeersEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeersEntity() *PeersEntity {
	this := PeersEntity{}
	return &this
}

// NewPeersEntityWithDefaults instantiates a new PeersEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeersEntityWithDefaults() *PeersEntity {
	this := PeersEntity{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *PeersEntity) GetPeers() []PeerDTO {
	if o == nil || IsNil(o.Peers) {
		var ret []PeerDTO
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeersEntity) GetPeersOk() ([]PeerDTO, bool) {
	if o == nil || IsNil(o.Peers) {
		return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *PeersEntity) HasPeers() bool {
	if o != nil && !IsNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []PeerDTO and assigns it to the Peers field.
func (o *PeersEntity) SetPeers(v []PeerDTO) {
	o.Peers = v
}

func (o PeersEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeersEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return toSerialize, nil
}

type NullablePeersEntity struct {
	value *PeersEntity
	isSet bool
}

func (v NullablePeersEntity) Get() *PeersEntity {
	return v.value
}

func (v *NullablePeersEntity) Set(val *PeersEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePeersEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePeersEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeersEntity(val *PeersEntity) *NullablePeersEntity {
	return &NullablePeersEntity{value: val, isSet: true}
}

func (v NullablePeersEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeersEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


