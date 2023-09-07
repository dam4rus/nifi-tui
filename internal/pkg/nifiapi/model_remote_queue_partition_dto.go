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

// checks if the RemoteQueuePartitionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteQueuePartitionDTO{}

// RemoteQueuePartitionDTO struct for RemoteQueuePartitionDTO
type RemoteQueuePartitionDTO struct {
	// Total number of FlowFiles owned by the Connection
	TotalFlowFileCount *int32 `json:"totalFlowFileCount,omitempty"`
	// Total number of bytes that make up the content for the FlowFiles owned by this Connection
	TotalByteCount *int64 `json:"totalByteCount,omitempty"`
	// Total number of FlowFiles that exist in the Connection's Active Queue, immediately available to be offered up to a component
	ActiveQueueFlowFileCount *int32 `json:"activeQueueFlowFileCount,omitempty"`
	// Total number of bytes that make up the content for the FlowFiles that are present in the Connection's Active Queue
	ActiveQueueByteCount *int64 `json:"activeQueueByteCount,omitempty"`
	// The total number of FlowFiles that are swapped out for this Connection
	SwapFlowFileCount *int32 `json:"swapFlowFileCount,omitempty"`
	// Total number of bytes that make up the content for the FlowFiles that are swapped out to disk for the Connection
	SwapByteCount *int64 `json:"swapByteCount,omitempty"`
	// The number of Swap Files that exist for this Connection
	SwapFiles *int32 `json:"swapFiles,omitempty"`
	// The number of In-Flight FlowFiles for this Connection. These are FlowFiles that belong to the connection but are currently being operated on by a Processor, Port, etc.
	InFlightFlowFileCount *int32 `json:"inFlightFlowFileCount,omitempty"`
	// The number bytes that make up the content of the FlowFiles that are In-Flight
	InFlightByteCount *int64 `json:"inFlightByteCount,omitempty"`
	// The Node Identifier that this queue partition is sending to
	NodeIdentifier *string `json:"nodeIdentifier,omitempty"`
}

// NewRemoteQueuePartitionDTO instantiates a new RemoteQueuePartitionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteQueuePartitionDTO() *RemoteQueuePartitionDTO {
	this := RemoteQueuePartitionDTO{}
	return &this
}

// NewRemoteQueuePartitionDTOWithDefaults instantiates a new RemoteQueuePartitionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteQueuePartitionDTOWithDefaults() *RemoteQueuePartitionDTO {
	this := RemoteQueuePartitionDTO{}
	return &this
}

// GetTotalFlowFileCount returns the TotalFlowFileCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetTotalFlowFileCount() int32 {
	if o == nil || IsNil(o.TotalFlowFileCount) {
		var ret int32
		return ret
	}
	return *o.TotalFlowFileCount
}

// GetTotalFlowFileCountOk returns a tuple with the TotalFlowFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetTotalFlowFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalFlowFileCount) {
		return nil, false
	}
	return o.TotalFlowFileCount, true
}

// HasTotalFlowFileCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasTotalFlowFileCount() bool {
	if o != nil && !IsNil(o.TotalFlowFileCount) {
		return true
	}

	return false
}

// SetTotalFlowFileCount gets a reference to the given int32 and assigns it to the TotalFlowFileCount field.
func (o *RemoteQueuePartitionDTO) SetTotalFlowFileCount(v int32) {
	o.TotalFlowFileCount = &v
}

// GetTotalByteCount returns the TotalByteCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetTotalByteCount() int64 {
	if o == nil || IsNil(o.TotalByteCount) {
		var ret int64
		return ret
	}
	return *o.TotalByteCount
}

// GetTotalByteCountOk returns a tuple with the TotalByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetTotalByteCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalByteCount) {
		return nil, false
	}
	return o.TotalByteCount, true
}

// HasTotalByteCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasTotalByteCount() bool {
	if o != nil && !IsNil(o.TotalByteCount) {
		return true
	}

	return false
}

// SetTotalByteCount gets a reference to the given int64 and assigns it to the TotalByteCount field.
func (o *RemoteQueuePartitionDTO) SetTotalByteCount(v int64) {
	o.TotalByteCount = &v
}

// GetActiveQueueFlowFileCount returns the ActiveQueueFlowFileCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetActiveQueueFlowFileCount() int32 {
	if o == nil || IsNil(o.ActiveQueueFlowFileCount) {
		var ret int32
		return ret
	}
	return *o.ActiveQueueFlowFileCount
}

// GetActiveQueueFlowFileCountOk returns a tuple with the ActiveQueueFlowFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetActiveQueueFlowFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveQueueFlowFileCount) {
		return nil, false
	}
	return o.ActiveQueueFlowFileCount, true
}

// HasActiveQueueFlowFileCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasActiveQueueFlowFileCount() bool {
	if o != nil && !IsNil(o.ActiveQueueFlowFileCount) {
		return true
	}

	return false
}

// SetActiveQueueFlowFileCount gets a reference to the given int32 and assigns it to the ActiveQueueFlowFileCount field.
func (o *RemoteQueuePartitionDTO) SetActiveQueueFlowFileCount(v int32) {
	o.ActiveQueueFlowFileCount = &v
}

// GetActiveQueueByteCount returns the ActiveQueueByteCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetActiveQueueByteCount() int64 {
	if o == nil || IsNil(o.ActiveQueueByteCount) {
		var ret int64
		return ret
	}
	return *o.ActiveQueueByteCount
}

// GetActiveQueueByteCountOk returns a tuple with the ActiveQueueByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetActiveQueueByteCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ActiveQueueByteCount) {
		return nil, false
	}
	return o.ActiveQueueByteCount, true
}

// HasActiveQueueByteCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasActiveQueueByteCount() bool {
	if o != nil && !IsNil(o.ActiveQueueByteCount) {
		return true
	}

	return false
}

// SetActiveQueueByteCount gets a reference to the given int64 and assigns it to the ActiveQueueByteCount field.
func (o *RemoteQueuePartitionDTO) SetActiveQueueByteCount(v int64) {
	o.ActiveQueueByteCount = &v
}

// GetSwapFlowFileCount returns the SwapFlowFileCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetSwapFlowFileCount() int32 {
	if o == nil || IsNil(o.SwapFlowFileCount) {
		var ret int32
		return ret
	}
	return *o.SwapFlowFileCount
}

// GetSwapFlowFileCountOk returns a tuple with the SwapFlowFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetSwapFlowFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SwapFlowFileCount) {
		return nil, false
	}
	return o.SwapFlowFileCount, true
}

// HasSwapFlowFileCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasSwapFlowFileCount() bool {
	if o != nil && !IsNil(o.SwapFlowFileCount) {
		return true
	}

	return false
}

// SetSwapFlowFileCount gets a reference to the given int32 and assigns it to the SwapFlowFileCount field.
func (o *RemoteQueuePartitionDTO) SetSwapFlowFileCount(v int32) {
	o.SwapFlowFileCount = &v
}

// GetSwapByteCount returns the SwapByteCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetSwapByteCount() int64 {
	if o == nil || IsNil(o.SwapByteCount) {
		var ret int64
		return ret
	}
	return *o.SwapByteCount
}

// GetSwapByteCountOk returns a tuple with the SwapByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetSwapByteCountOk() (*int64, bool) {
	if o == nil || IsNil(o.SwapByteCount) {
		return nil, false
	}
	return o.SwapByteCount, true
}

// HasSwapByteCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasSwapByteCount() bool {
	if o != nil && !IsNil(o.SwapByteCount) {
		return true
	}

	return false
}

// SetSwapByteCount gets a reference to the given int64 and assigns it to the SwapByteCount field.
func (o *RemoteQueuePartitionDTO) SetSwapByteCount(v int64) {
	o.SwapByteCount = &v
}

// GetSwapFiles returns the SwapFiles field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetSwapFiles() int32 {
	if o == nil || IsNil(o.SwapFiles) {
		var ret int32
		return ret
	}
	return *o.SwapFiles
}

// GetSwapFilesOk returns a tuple with the SwapFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetSwapFilesOk() (*int32, bool) {
	if o == nil || IsNil(o.SwapFiles) {
		return nil, false
	}
	return o.SwapFiles, true
}

// HasSwapFiles returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasSwapFiles() bool {
	if o != nil && !IsNil(o.SwapFiles) {
		return true
	}

	return false
}

// SetSwapFiles gets a reference to the given int32 and assigns it to the SwapFiles field.
func (o *RemoteQueuePartitionDTO) SetSwapFiles(v int32) {
	o.SwapFiles = &v
}

// GetInFlightFlowFileCount returns the InFlightFlowFileCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetInFlightFlowFileCount() int32 {
	if o == nil || IsNil(o.InFlightFlowFileCount) {
		var ret int32
		return ret
	}
	return *o.InFlightFlowFileCount
}

// GetInFlightFlowFileCountOk returns a tuple with the InFlightFlowFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetInFlightFlowFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InFlightFlowFileCount) {
		return nil, false
	}
	return o.InFlightFlowFileCount, true
}

// HasInFlightFlowFileCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasInFlightFlowFileCount() bool {
	if o != nil && !IsNil(o.InFlightFlowFileCount) {
		return true
	}

	return false
}

// SetInFlightFlowFileCount gets a reference to the given int32 and assigns it to the InFlightFlowFileCount field.
func (o *RemoteQueuePartitionDTO) SetInFlightFlowFileCount(v int32) {
	o.InFlightFlowFileCount = &v
}

// GetInFlightByteCount returns the InFlightByteCount field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetInFlightByteCount() int64 {
	if o == nil || IsNil(o.InFlightByteCount) {
		var ret int64
		return ret
	}
	return *o.InFlightByteCount
}

// GetInFlightByteCountOk returns a tuple with the InFlightByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetInFlightByteCountOk() (*int64, bool) {
	if o == nil || IsNil(o.InFlightByteCount) {
		return nil, false
	}
	return o.InFlightByteCount, true
}

// HasInFlightByteCount returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasInFlightByteCount() bool {
	if o != nil && !IsNil(o.InFlightByteCount) {
		return true
	}

	return false
}

// SetInFlightByteCount gets a reference to the given int64 and assigns it to the InFlightByteCount field.
func (o *RemoteQueuePartitionDTO) SetInFlightByteCount(v int64) {
	o.InFlightByteCount = &v
}

// GetNodeIdentifier returns the NodeIdentifier field value if set, zero value otherwise.
func (o *RemoteQueuePartitionDTO) GetNodeIdentifier() string {
	if o == nil || IsNil(o.NodeIdentifier) {
		var ret string
		return ret
	}
	return *o.NodeIdentifier
}

// GetNodeIdentifierOk returns a tuple with the NodeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteQueuePartitionDTO) GetNodeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.NodeIdentifier) {
		return nil, false
	}
	return o.NodeIdentifier, true
}

// HasNodeIdentifier returns a boolean if a field has been set.
func (o *RemoteQueuePartitionDTO) HasNodeIdentifier() bool {
	if o != nil && !IsNil(o.NodeIdentifier) {
		return true
	}

	return false
}

// SetNodeIdentifier gets a reference to the given string and assigns it to the NodeIdentifier field.
func (o *RemoteQueuePartitionDTO) SetNodeIdentifier(v string) {
	o.NodeIdentifier = &v
}

func (o RemoteQueuePartitionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteQueuePartitionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalFlowFileCount) {
		toSerialize["totalFlowFileCount"] = o.TotalFlowFileCount
	}
	if !IsNil(o.TotalByteCount) {
		toSerialize["totalByteCount"] = o.TotalByteCount
	}
	if !IsNil(o.ActiveQueueFlowFileCount) {
		toSerialize["activeQueueFlowFileCount"] = o.ActiveQueueFlowFileCount
	}
	if !IsNil(o.ActiveQueueByteCount) {
		toSerialize["activeQueueByteCount"] = o.ActiveQueueByteCount
	}
	if !IsNil(o.SwapFlowFileCount) {
		toSerialize["swapFlowFileCount"] = o.SwapFlowFileCount
	}
	if !IsNil(o.SwapByteCount) {
		toSerialize["swapByteCount"] = o.SwapByteCount
	}
	if !IsNil(o.SwapFiles) {
		toSerialize["swapFiles"] = o.SwapFiles
	}
	if !IsNil(o.InFlightFlowFileCount) {
		toSerialize["inFlightFlowFileCount"] = o.InFlightFlowFileCount
	}
	if !IsNil(o.InFlightByteCount) {
		toSerialize["inFlightByteCount"] = o.InFlightByteCount
	}
	if !IsNil(o.NodeIdentifier) {
		toSerialize["nodeIdentifier"] = o.NodeIdentifier
	}
	return toSerialize, nil
}

type NullableRemoteQueuePartitionDTO struct {
	value *RemoteQueuePartitionDTO
	isSet bool
}

func (v NullableRemoteQueuePartitionDTO) Get() *RemoteQueuePartitionDTO {
	return v.value
}

func (v *NullableRemoteQueuePartitionDTO) Set(val *RemoteQueuePartitionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteQueuePartitionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteQueuePartitionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteQueuePartitionDTO(val *RemoteQueuePartitionDTO) *NullableRemoteQueuePartitionDTO {
	return &NullableRemoteQueuePartitionDTO{value: val, isSet: true}
}

func (v NullableRemoteQueuePartitionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteQueuePartitionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


