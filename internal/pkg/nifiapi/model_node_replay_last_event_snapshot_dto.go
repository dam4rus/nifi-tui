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

// checks if the NodeReplayLastEventSnapshotDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeReplayLastEventSnapshotDTO{}

// NodeReplayLastEventSnapshotDTO struct for NodeReplayLastEventSnapshotDTO
type NodeReplayLastEventSnapshotDTO struct {
	// The unique ID that identifies the node
	NodeId *string `json:"nodeId,omitempty"`
	// The API address of the node
	Address *string `json:"address,omitempty"`
	// The API port used to communicate with the node
	ApiPort *int32 `json:"apiPort,omitempty"`
	Snapshot *ReplayLastEventSnapshotDTO `json:"snapshot,omitempty"`
}

// NewNodeReplayLastEventSnapshotDTO instantiates a new NodeReplayLastEventSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeReplayLastEventSnapshotDTO() *NodeReplayLastEventSnapshotDTO {
	this := NodeReplayLastEventSnapshotDTO{}
	return &this
}

// NewNodeReplayLastEventSnapshotDTOWithDefaults instantiates a new NodeReplayLastEventSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeReplayLastEventSnapshotDTOWithDefaults() *NodeReplayLastEventSnapshotDTO {
	this := NodeReplayLastEventSnapshotDTO{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *NodeReplayLastEventSnapshotDTO) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReplayLastEventSnapshotDTO) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *NodeReplayLastEventSnapshotDTO) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *NodeReplayLastEventSnapshotDTO) SetNodeId(v string) {
	o.NodeId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NodeReplayLastEventSnapshotDTO) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReplayLastEventSnapshotDTO) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NodeReplayLastEventSnapshotDTO) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NodeReplayLastEventSnapshotDTO) SetAddress(v string) {
	o.Address = &v
}

// GetApiPort returns the ApiPort field value if set, zero value otherwise.
func (o *NodeReplayLastEventSnapshotDTO) GetApiPort() int32 {
	if o == nil || IsNil(o.ApiPort) {
		var ret int32
		return ret
	}
	return *o.ApiPort
}

// GetApiPortOk returns a tuple with the ApiPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReplayLastEventSnapshotDTO) GetApiPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ApiPort) {
		return nil, false
	}
	return o.ApiPort, true
}

// HasApiPort returns a boolean if a field has been set.
func (o *NodeReplayLastEventSnapshotDTO) HasApiPort() bool {
	if o != nil && !IsNil(o.ApiPort) {
		return true
	}

	return false
}

// SetApiPort gets a reference to the given int32 and assigns it to the ApiPort field.
func (o *NodeReplayLastEventSnapshotDTO) SetApiPort(v int32) {
	o.ApiPort = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *NodeReplayLastEventSnapshotDTO) GetSnapshot() ReplayLastEventSnapshotDTO {
	if o == nil || IsNil(o.Snapshot) {
		var ret ReplayLastEventSnapshotDTO
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReplayLastEventSnapshotDTO) GetSnapshotOk() (*ReplayLastEventSnapshotDTO, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *NodeReplayLastEventSnapshotDTO) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given ReplayLastEventSnapshotDTO and assigns it to the Snapshot field.
func (o *NodeReplayLastEventSnapshotDTO) SetSnapshot(v ReplayLastEventSnapshotDTO) {
	o.Snapshot = &v
}

func (o NodeReplayLastEventSnapshotDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeReplayLastEventSnapshotDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodeId) {
		toSerialize["nodeId"] = o.NodeId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ApiPort) {
		toSerialize["apiPort"] = o.ApiPort
	}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	return toSerialize, nil
}

type NullableNodeReplayLastEventSnapshotDTO struct {
	value *NodeReplayLastEventSnapshotDTO
	isSet bool
}

func (v NullableNodeReplayLastEventSnapshotDTO) Get() *NodeReplayLastEventSnapshotDTO {
	return v.value
}

func (v *NullableNodeReplayLastEventSnapshotDTO) Set(val *NodeReplayLastEventSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeReplayLastEventSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeReplayLastEventSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeReplayLastEventSnapshotDTO(val *NodeReplayLastEventSnapshotDTO) *NullableNodeReplayLastEventSnapshotDTO {
	return &NullableNodeReplayLastEventSnapshotDTO{value: val, isSet: true}
}

func (v NullableNodeReplayLastEventSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeReplayLastEventSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

