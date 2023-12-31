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

// checks if the StateEntryDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateEntryDTO{}

// StateEntryDTO struct for StateEntryDTO
type StateEntryDTO struct {
	// The key for this state.
	Key *string `json:"key,omitempty"`
	// The value for this state.
	Value *string `json:"value,omitempty"`
	// The identifier for the node where the state originated.
	ClusterNodeId *string `json:"clusterNodeId,omitempty"`
	// The label for the node where the state originated.
	ClusterNodeAddress *string `json:"clusterNodeAddress,omitempty"`
}

// NewStateEntryDTO instantiates a new StateEntryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateEntryDTO() *StateEntryDTO {
	this := StateEntryDTO{}
	return &this
}

// NewStateEntryDTOWithDefaults instantiates a new StateEntryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateEntryDTOWithDefaults() *StateEntryDTO {
	this := StateEntryDTO{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StateEntryDTO) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateEntryDTO) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StateEntryDTO) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StateEntryDTO) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StateEntryDTO) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateEntryDTO) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StateEntryDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *StateEntryDTO) SetValue(v string) {
	o.Value = &v
}

// GetClusterNodeId returns the ClusterNodeId field value if set, zero value otherwise.
func (o *StateEntryDTO) GetClusterNodeId() string {
	if o == nil || IsNil(o.ClusterNodeId) {
		var ret string
		return ret
	}
	return *o.ClusterNodeId
}

// GetClusterNodeIdOk returns a tuple with the ClusterNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateEntryDTO) GetClusterNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterNodeId) {
		return nil, false
	}
	return o.ClusterNodeId, true
}

// HasClusterNodeId returns a boolean if a field has been set.
func (o *StateEntryDTO) HasClusterNodeId() bool {
	if o != nil && !IsNil(o.ClusterNodeId) {
		return true
	}

	return false
}

// SetClusterNodeId gets a reference to the given string and assigns it to the ClusterNodeId field.
func (o *StateEntryDTO) SetClusterNodeId(v string) {
	o.ClusterNodeId = &v
}

// GetClusterNodeAddress returns the ClusterNodeAddress field value if set, zero value otherwise.
func (o *StateEntryDTO) GetClusterNodeAddress() string {
	if o == nil || IsNil(o.ClusterNodeAddress) {
		var ret string
		return ret
	}
	return *o.ClusterNodeAddress
}

// GetClusterNodeAddressOk returns a tuple with the ClusterNodeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateEntryDTO) GetClusterNodeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterNodeAddress) {
		return nil, false
	}
	return o.ClusterNodeAddress, true
}

// HasClusterNodeAddress returns a boolean if a field has been set.
func (o *StateEntryDTO) HasClusterNodeAddress() bool {
	if o != nil && !IsNil(o.ClusterNodeAddress) {
		return true
	}

	return false
}

// SetClusterNodeAddress gets a reference to the given string and assigns it to the ClusterNodeAddress field.
func (o *StateEntryDTO) SetClusterNodeAddress(v string) {
	o.ClusterNodeAddress = &v
}

func (o StateEntryDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateEntryDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ClusterNodeId) {
		toSerialize["clusterNodeId"] = o.ClusterNodeId
	}
	if !IsNil(o.ClusterNodeAddress) {
		toSerialize["clusterNodeAddress"] = o.ClusterNodeAddress
	}
	return toSerialize, nil
}

type NullableStateEntryDTO struct {
	value *StateEntryDTO
	isSet bool
}

func (v NullableStateEntryDTO) Get() *StateEntryDTO {
	return v.value
}

func (v *NullableStateEntryDTO) Set(val *StateEntryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableStateEntryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableStateEntryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateEntryDTO(val *StateEntryDTO) *NullableStateEntryDTO {
	return &NullableStateEntryDTO{value: val, isSet: true}
}

func (v NullableStateEntryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateEntryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


