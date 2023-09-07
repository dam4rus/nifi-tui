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

// checks if the ConnectionStatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionStatusDTO{}

// ConnectionStatusDTO struct for ConnectionStatusDTO
type ConnectionStatusDTO struct {
	// The ID of the connection
	Id *string `json:"id,omitempty"`
	// The ID of the Process Group that the connection belongs to
	GroupId *string `json:"groupId,omitempty"`
	// The name of the connection
	Name *string `json:"name,omitempty"`
	// The timestamp of when the stats were last refreshed
	StatsLastRefreshed *string `json:"statsLastRefreshed,omitempty"`
	// The ID of the source component
	SourceId *string `json:"sourceId,omitempty"`
	// The name of the source component
	SourceName *string `json:"sourceName,omitempty"`
	// The ID of the destination component
	DestinationId *string `json:"destinationId,omitempty"`
	// The name of the destination component
	DestinationName *string `json:"destinationName,omitempty"`
	AggregateSnapshot *ConnectionStatusSnapshotDTO `json:"aggregateSnapshot,omitempty"`
	// A list of status snapshots for each node
	NodeSnapshots []NodeConnectionStatusSnapshotDTO `json:"nodeSnapshots,omitempty"`
}

// NewConnectionStatusDTO instantiates a new ConnectionStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionStatusDTO() *ConnectionStatusDTO {
	this := ConnectionStatusDTO{}
	return &this
}

// NewConnectionStatusDTOWithDefaults instantiates a new ConnectionStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionStatusDTOWithDefaults() *ConnectionStatusDTO {
	this := ConnectionStatusDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionStatusDTO) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ConnectionStatusDTO) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionStatusDTO) SetName(v string) {
	o.Name = &v
}

// GetStatsLastRefreshed returns the StatsLastRefreshed field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetStatsLastRefreshed() string {
	if o == nil || IsNil(o.StatsLastRefreshed) {
		var ret string
		return ret
	}
	return *o.StatsLastRefreshed
}

// GetStatsLastRefreshedOk returns a tuple with the StatsLastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetStatsLastRefreshedOk() (*string, bool) {
	if o == nil || IsNil(o.StatsLastRefreshed) {
		return nil, false
	}
	return o.StatsLastRefreshed, true
}

// HasStatsLastRefreshed returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasStatsLastRefreshed() bool {
	if o != nil && !IsNil(o.StatsLastRefreshed) {
		return true
	}

	return false
}

// SetStatsLastRefreshed gets a reference to the given string and assigns it to the StatsLastRefreshed field.
func (o *ConnectionStatusDTO) SetStatsLastRefreshed(v string) {
	o.StatsLastRefreshed = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ConnectionStatusDTO) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetSourceName() string {
	if o == nil || IsNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasSourceName() bool {
	if o != nil && !IsNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *ConnectionStatusDTO) SetSourceName(v string) {
	o.SourceName = &v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetDestinationId() string {
	if o == nil || IsNil(o.DestinationId) {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetDestinationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationId) {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasDestinationId() bool {
	if o != nil && !IsNil(o.DestinationId) {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *ConnectionStatusDTO) SetDestinationId(v string) {
	o.DestinationId = &v
}

// GetDestinationName returns the DestinationName field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetDestinationName() string {
	if o == nil || IsNil(o.DestinationName) {
		var ret string
		return ret
	}
	return *o.DestinationName
}

// GetDestinationNameOk returns a tuple with the DestinationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetDestinationNameOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationName) {
		return nil, false
	}
	return o.DestinationName, true
}

// HasDestinationName returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasDestinationName() bool {
	if o != nil && !IsNil(o.DestinationName) {
		return true
	}

	return false
}

// SetDestinationName gets a reference to the given string and assigns it to the DestinationName field.
func (o *ConnectionStatusDTO) SetDestinationName(v string) {
	o.DestinationName = &v
}

// GetAggregateSnapshot returns the AggregateSnapshot field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetAggregateSnapshot() ConnectionStatusSnapshotDTO {
	if o == nil || IsNil(o.AggregateSnapshot) {
		var ret ConnectionStatusSnapshotDTO
		return ret
	}
	return *o.AggregateSnapshot
}

// GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetAggregateSnapshotOk() (*ConnectionStatusSnapshotDTO, bool) {
	if o == nil || IsNil(o.AggregateSnapshot) {
		return nil, false
	}
	return o.AggregateSnapshot, true
}

// HasAggregateSnapshot returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasAggregateSnapshot() bool {
	if o != nil && !IsNil(o.AggregateSnapshot) {
		return true
	}

	return false
}

// SetAggregateSnapshot gets a reference to the given ConnectionStatusSnapshotDTO and assigns it to the AggregateSnapshot field.
func (o *ConnectionStatusDTO) SetAggregateSnapshot(v ConnectionStatusSnapshotDTO) {
	o.AggregateSnapshot = &v
}

// GetNodeSnapshots returns the NodeSnapshots field value if set, zero value otherwise.
func (o *ConnectionStatusDTO) GetNodeSnapshots() []NodeConnectionStatusSnapshotDTO {
	if o == nil || IsNil(o.NodeSnapshots) {
		var ret []NodeConnectionStatusSnapshotDTO
		return ret
	}
	return o.NodeSnapshots
}

// GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionStatusDTO) GetNodeSnapshotsOk() ([]NodeConnectionStatusSnapshotDTO, bool) {
	if o == nil || IsNil(o.NodeSnapshots) {
		return nil, false
	}
	return o.NodeSnapshots, true
}

// HasNodeSnapshots returns a boolean if a field has been set.
func (o *ConnectionStatusDTO) HasNodeSnapshots() bool {
	if o != nil && !IsNil(o.NodeSnapshots) {
		return true
	}

	return false
}

// SetNodeSnapshots gets a reference to the given []NodeConnectionStatusSnapshotDTO and assigns it to the NodeSnapshots field.
func (o *ConnectionStatusDTO) SetNodeSnapshots(v []NodeConnectionStatusSnapshotDTO) {
	o.NodeSnapshots = v
}

func (o ConnectionStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionStatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StatsLastRefreshed) {
		toSerialize["statsLastRefreshed"] = o.StatsLastRefreshed
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !IsNil(o.DestinationId) {
		toSerialize["destinationId"] = o.DestinationId
	}
	if !IsNil(o.DestinationName) {
		toSerialize["destinationName"] = o.DestinationName
	}
	if !IsNil(o.AggregateSnapshot) {
		toSerialize["aggregateSnapshot"] = o.AggregateSnapshot
	}
	if !IsNil(o.NodeSnapshots) {
		toSerialize["nodeSnapshots"] = o.NodeSnapshots
	}
	return toSerialize, nil
}

type NullableConnectionStatusDTO struct {
	value *ConnectionStatusDTO
	isSet bool
}

func (v NullableConnectionStatusDTO) Get() *ConnectionStatusDTO {
	return v.value
}

func (v *NullableConnectionStatusDTO) Set(val *ConnectionStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionStatusDTO(val *ConnectionStatusDTO) *NullableConnectionStatusDTO {
	return &NullableConnectionStatusDTO{value: val, isSet: true}
}

func (v NullableConnectionStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


