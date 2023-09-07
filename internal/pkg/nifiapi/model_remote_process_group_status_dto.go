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

// checks if the RemoteProcessGroupStatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteProcessGroupStatusDTO{}

// RemoteProcessGroupStatusDTO struct for RemoteProcessGroupStatusDTO
type RemoteProcessGroupStatusDTO struct {
	// The unique ID of the process group that the Processor belongs to
	GroupId *string `json:"groupId,omitempty"`
	// The unique ID of the Processor
	Id *string `json:"id,omitempty"`
	// The name of the remote process group.
	Name *string `json:"name,omitempty"`
	// The URI of the target system.
	TargetUri *string `json:"targetUri,omitempty"`
	// The transmission status of the remote process group.
	TransmissionStatus *string `json:"transmissionStatus,omitempty"`
	// The time the status for the process group was last refreshed.
	StatsLastRefreshed *string `json:"statsLastRefreshed,omitempty"`
	// Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid)
	ValidationStatus *string `json:"validationStatus,omitempty"`
	AggregateSnapshot *RemoteProcessGroupStatusSnapshotDTO `json:"aggregateSnapshot,omitempty"`
	// A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots []NodeRemoteProcessGroupStatusSnapshotDTO `json:"nodeSnapshots,omitempty"`
}

// NewRemoteProcessGroupStatusDTO instantiates a new RemoteProcessGroupStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteProcessGroupStatusDTO() *RemoteProcessGroupStatusDTO {
	this := RemoteProcessGroupStatusDTO{}
	return &this
}

// NewRemoteProcessGroupStatusDTOWithDefaults instantiates a new RemoteProcessGroupStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteProcessGroupStatusDTOWithDefaults() *RemoteProcessGroupStatusDTO {
	this := RemoteProcessGroupStatusDTO{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *RemoteProcessGroupStatusDTO) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteProcessGroupStatusDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemoteProcessGroupStatusDTO) SetName(v string) {
	o.Name = &v
}

// GetTargetUri returns the TargetUri field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetTargetUri() string {
	if o == nil || IsNil(o.TargetUri) {
		var ret string
		return ret
	}
	return *o.TargetUri
}

// GetTargetUriOk returns a tuple with the TargetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetTargetUriOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUri) {
		return nil, false
	}
	return o.TargetUri, true
}

// HasTargetUri returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasTargetUri() bool {
	if o != nil && !IsNil(o.TargetUri) {
		return true
	}

	return false
}

// SetTargetUri gets a reference to the given string and assigns it to the TargetUri field.
func (o *RemoteProcessGroupStatusDTO) SetTargetUri(v string) {
	o.TargetUri = &v
}

// GetTransmissionStatus returns the TransmissionStatus field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetTransmissionStatus() string {
	if o == nil || IsNil(o.TransmissionStatus) {
		var ret string
		return ret
	}
	return *o.TransmissionStatus
}

// GetTransmissionStatusOk returns a tuple with the TransmissionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetTransmissionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TransmissionStatus) {
		return nil, false
	}
	return o.TransmissionStatus, true
}

// HasTransmissionStatus returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasTransmissionStatus() bool {
	if o != nil && !IsNil(o.TransmissionStatus) {
		return true
	}

	return false
}

// SetTransmissionStatus gets a reference to the given string and assigns it to the TransmissionStatus field.
func (o *RemoteProcessGroupStatusDTO) SetTransmissionStatus(v string) {
	o.TransmissionStatus = &v
}

// GetStatsLastRefreshed returns the StatsLastRefreshed field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetStatsLastRefreshed() string {
	if o == nil || IsNil(o.StatsLastRefreshed) {
		var ret string
		return ret
	}
	return *o.StatsLastRefreshed
}

// GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetStatsLastRefreshedOk() (*string, bool) {
	if o == nil || IsNil(o.StatsLastRefreshed) {
		return nil, false
	}
	return o.StatsLastRefreshed, true
}

// HasStatsLastRefreshed returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasStatsLastRefreshed() bool {
	if o != nil && !IsNil(o.StatsLastRefreshed) {
		return true
	}

	return false
}

// SetStatsLastRefreshed gets a reference to the given string and assigns it to the StatsLastRefreshed field.
func (o *RemoteProcessGroupStatusDTO) SetStatsLastRefreshed(v string) {
	o.StatsLastRefreshed = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetValidationStatus() string {
	if o == nil || IsNil(o.ValidationStatus) {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetValidationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationStatus) {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasValidationStatus() bool {
	if o != nil && !IsNil(o.ValidationStatus) {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *RemoteProcessGroupStatusDTO) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

// GetAggregateSnapshot returns the AggregateSnapshot field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetAggregateSnapshot() RemoteProcessGroupStatusSnapshotDTO {
	if o == nil || IsNil(o.AggregateSnapshot) {
		var ret RemoteProcessGroupStatusSnapshotDTO
		return ret
	}
	return *o.AggregateSnapshot
}

// GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetAggregateSnapshotOk() (*RemoteProcessGroupStatusSnapshotDTO, bool) {
	if o == nil || IsNil(o.AggregateSnapshot) {
		return nil, false
	}
	return o.AggregateSnapshot, true
}

// HasAggregateSnapshot returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasAggregateSnapshot() bool {
	if o != nil && !IsNil(o.AggregateSnapshot) {
		return true
	}

	return false
}

// SetAggregateSnapshot gets a reference to the given RemoteProcessGroupStatusSnapshotDTO and assigns it to the AggregateSnapshot field.
func (o *RemoteProcessGroupStatusDTO) SetAggregateSnapshot(v RemoteProcessGroupStatusSnapshotDTO) {
	o.AggregateSnapshot = &v
}

// GetNodeSnapshots returns the NodeSnapshots field value if set, zero value otherwise.
func (o *RemoteProcessGroupStatusDTO) GetNodeSnapshots() []NodeRemoteProcessGroupStatusSnapshotDTO {
	if o == nil || IsNil(o.NodeSnapshots) {
		var ret []NodeRemoteProcessGroupStatusSnapshotDTO
		return ret
	}
	return o.NodeSnapshots
}

// GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteProcessGroupStatusDTO) GetNodeSnapshotsOk() ([]NodeRemoteProcessGroupStatusSnapshotDTO, bool) {
	if o == nil || IsNil(o.NodeSnapshots) {
		return nil, false
	}
	return o.NodeSnapshots, true
}

// HasNodeSnapshots returns a boolean if a field has been set.
func (o *RemoteProcessGroupStatusDTO) HasNodeSnapshots() bool {
	if o != nil && !IsNil(o.NodeSnapshots) {
		return true
	}

	return false
}

// SetNodeSnapshots gets a reference to the given []NodeRemoteProcessGroupStatusSnapshotDTO and assigns it to the NodeSnapshots field.
func (o *RemoteProcessGroupStatusDTO) SetNodeSnapshots(v []NodeRemoteProcessGroupStatusSnapshotDTO) {
	o.NodeSnapshots = v
}

func (o RemoteProcessGroupStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteProcessGroupStatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TargetUri) {
		toSerialize["targetUri"] = o.TargetUri
	}
	if !IsNil(o.TransmissionStatus) {
		toSerialize["transmissionStatus"] = o.TransmissionStatus
	}
	if !IsNil(o.StatsLastRefreshed) {
		toSerialize["statsLastRefreshed"] = o.StatsLastRefreshed
	}
	if !IsNil(o.ValidationStatus) {
		toSerialize["validationStatus"] = o.ValidationStatus
	}
	if !IsNil(o.AggregateSnapshot) {
		toSerialize["aggregateSnapshot"] = o.AggregateSnapshot
	}
	if !IsNil(o.NodeSnapshots) {
		toSerialize["nodeSnapshots"] = o.NodeSnapshots
	}
	return toSerialize, nil
}

type NullableRemoteProcessGroupStatusDTO struct {
	value *RemoteProcessGroupStatusDTO
	isSet bool
}

func (v NullableRemoteProcessGroupStatusDTO) Get() *RemoteProcessGroupStatusDTO {
	return v.value
}

func (v *NullableRemoteProcessGroupStatusDTO) Set(val *RemoteProcessGroupStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteProcessGroupStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteProcessGroupStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteProcessGroupStatusDTO(val *RemoteProcessGroupStatusDTO) *NullableRemoteProcessGroupStatusDTO {
	return &NullableRemoteProcessGroupStatusDTO{value: val, isSet: true}
}

func (v NullableRemoteProcessGroupStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteProcessGroupStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

