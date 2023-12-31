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

// checks if the NodePortStatusSnapshotDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodePortStatusSnapshotDTO{}

// NodePortStatusSnapshotDTO struct for NodePortStatusSnapshotDTO
type NodePortStatusSnapshotDTO struct {
	// The unique ID that identifies the node
	NodeId *string `json:"nodeId,omitempty"`
	// The API address of the node
	Address *string `json:"address,omitempty"`
	// The API port used to communicate with the node
	ApiPort *int32 `json:"apiPort,omitempty"`
	StatusSnapshot *PortStatusSnapshotDTO `json:"statusSnapshot,omitempty"`
}

// NewNodePortStatusSnapshotDTO instantiates a new NodePortStatusSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodePortStatusSnapshotDTO() *NodePortStatusSnapshotDTO {
	this := NodePortStatusSnapshotDTO{}
	return &this
}

// NewNodePortStatusSnapshotDTOWithDefaults instantiates a new NodePortStatusSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodePortStatusSnapshotDTOWithDefaults() *NodePortStatusSnapshotDTO {
	this := NodePortStatusSnapshotDTO{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *NodePortStatusSnapshotDTO) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePortStatusSnapshotDTO) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *NodePortStatusSnapshotDTO) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *NodePortStatusSnapshotDTO) SetNodeId(v string) {
	o.NodeId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NodePortStatusSnapshotDTO) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePortStatusSnapshotDTO) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NodePortStatusSnapshotDTO) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NodePortStatusSnapshotDTO) SetAddress(v string) {
	o.Address = &v
}

// GetApiPort returns the ApiPort field value if set, zero value otherwise.
func (o *NodePortStatusSnapshotDTO) GetApiPort() int32 {
	if o == nil || IsNil(o.ApiPort) {
		var ret int32
		return ret
	}
	return *o.ApiPort
}

// GetApiPortOk returns a tuple with the ApiPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePortStatusSnapshotDTO) GetApiPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ApiPort) {
		return nil, false
	}
	return o.ApiPort, true
}

// HasApiPort returns a boolean if a field has been set.
func (o *NodePortStatusSnapshotDTO) HasApiPort() bool {
	if o != nil && !IsNil(o.ApiPort) {
		return true
	}

	return false
}

// SetApiPort gets a reference to the given int32 and assigns it to the ApiPort field.
func (o *NodePortStatusSnapshotDTO) SetApiPort(v int32) {
	o.ApiPort = &v
}

// GetStatusSnapshot returns the StatusSnapshot field value if set, zero value otherwise.
func (o *NodePortStatusSnapshotDTO) GetStatusSnapshot() PortStatusSnapshotDTO {
	if o == nil || IsNil(o.StatusSnapshot) {
		var ret PortStatusSnapshotDTO
		return ret
	}
	return *o.StatusSnapshot
}

// GetStatusSnapshotOk returns a tuple with the StatusSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodePortStatusSnapshotDTO) GetStatusSnapshotOk() (*PortStatusSnapshotDTO, bool) {
	if o == nil || IsNil(o.StatusSnapshot) {
		return nil, false
	}
	return o.StatusSnapshot, true
}

// HasStatusSnapshot returns a boolean if a field has been set.
func (o *NodePortStatusSnapshotDTO) HasStatusSnapshot() bool {
	if o != nil && !IsNil(o.StatusSnapshot) {
		return true
	}

	return false
}

// SetStatusSnapshot gets a reference to the given PortStatusSnapshotDTO and assigns it to the StatusSnapshot field.
func (o *NodePortStatusSnapshotDTO) SetStatusSnapshot(v PortStatusSnapshotDTO) {
	o.StatusSnapshot = &v
}

func (o NodePortStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodePortStatusSnapshotDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.StatusSnapshot) {
		toSerialize["statusSnapshot"] = o.StatusSnapshot
	}
	return toSerialize, nil
}

type NullableNodePortStatusSnapshotDTO struct {
	value *NodePortStatusSnapshotDTO
	isSet bool
}

func (v NullableNodePortStatusSnapshotDTO) Get() *NodePortStatusSnapshotDTO {
	return v.value
}

func (v *NullableNodePortStatusSnapshotDTO) Set(val *NodePortStatusSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableNodePortStatusSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableNodePortStatusSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodePortStatusSnapshotDTO(val *NodePortStatusSnapshotDTO) *NullableNodePortStatusSnapshotDTO {
	return &NullableNodePortStatusSnapshotDTO{value: val, isSet: true}
}

func (v NullableNodePortStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodePortStatusSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


