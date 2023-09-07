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

// checks if the ConnectionDiagnosticsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionDiagnosticsDTO{}

// ConnectionDiagnosticsDTO struct for ConnectionDiagnosticsDTO
type ConnectionDiagnosticsDTO struct {
	Connection *ConnectionDTO `json:"connection,omitempty"`
	AggregateSnapshot *ConnectionDiagnosticsSnapshotDTO `json:"aggregateSnapshot,omitempty"`
	// A list of values for each node in the cluster, if clustered.
	NodeSnapshots []ConnectionDiagnosticsSnapshotDTO `json:"nodeSnapshots,omitempty"`
}

// NewConnectionDiagnosticsDTO instantiates a new ConnectionDiagnosticsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionDiagnosticsDTO() *ConnectionDiagnosticsDTO {
	this := ConnectionDiagnosticsDTO{}
	return &this
}

// NewConnectionDiagnosticsDTOWithDefaults instantiates a new ConnectionDiagnosticsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionDiagnosticsDTOWithDefaults() *ConnectionDiagnosticsDTO {
	this := ConnectionDiagnosticsDTO{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ConnectionDiagnosticsDTO) GetConnection() ConnectionDTO {
	if o == nil || IsNil(o.Connection) {
		var ret ConnectionDTO
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDiagnosticsDTO) GetConnectionOk() (*ConnectionDTO, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ConnectionDiagnosticsDTO) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given ConnectionDTO and assigns it to the Connection field.
func (o *ConnectionDiagnosticsDTO) SetConnection(v ConnectionDTO) {
	o.Connection = &v
}

// GetAggregateSnapshot returns the AggregateSnapshot field value if set, zero value otherwise.
func (o *ConnectionDiagnosticsDTO) GetAggregateSnapshot() ConnectionDiagnosticsSnapshotDTO {
	if o == nil || IsNil(o.AggregateSnapshot) {
		var ret ConnectionDiagnosticsSnapshotDTO
		return ret
	}
	return *o.AggregateSnapshot
}

// GetAggregateSnapshotOk returns a tuple with the AggregateSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDiagnosticsDTO) GetAggregateSnapshotOk() (*ConnectionDiagnosticsSnapshotDTO, bool) {
	if o == nil || IsNil(o.AggregateSnapshot) {
		return nil, false
	}
	return o.AggregateSnapshot, true
}

// HasAggregateSnapshot returns a boolean if a field has been set.
func (o *ConnectionDiagnosticsDTO) HasAggregateSnapshot() bool {
	if o != nil && !IsNil(o.AggregateSnapshot) {
		return true
	}

	return false
}

// SetAggregateSnapshot gets a reference to the given ConnectionDiagnosticsSnapshotDTO and assigns it to the AggregateSnapshot field.
func (o *ConnectionDiagnosticsDTO) SetAggregateSnapshot(v ConnectionDiagnosticsSnapshotDTO) {
	o.AggregateSnapshot = &v
}

// GetNodeSnapshots returns the NodeSnapshots field value if set, zero value otherwise.
func (o *ConnectionDiagnosticsDTO) GetNodeSnapshots() []ConnectionDiagnosticsSnapshotDTO {
	if o == nil || IsNil(o.NodeSnapshots) {
		var ret []ConnectionDiagnosticsSnapshotDTO
		return ret
	}
	return o.NodeSnapshots
}

// GetNodeSnapshotsOk returns a tuple with the NodeSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDiagnosticsDTO) GetNodeSnapshotsOk() ([]ConnectionDiagnosticsSnapshotDTO, bool) {
	if o == nil || IsNil(o.NodeSnapshots) {
		return nil, false
	}
	return o.NodeSnapshots, true
}

// HasNodeSnapshots returns a boolean if a field has been set.
func (o *ConnectionDiagnosticsDTO) HasNodeSnapshots() bool {
	if o != nil && !IsNil(o.NodeSnapshots) {
		return true
	}

	return false
}

// SetNodeSnapshots gets a reference to the given []ConnectionDiagnosticsSnapshotDTO and assigns it to the NodeSnapshots field.
func (o *ConnectionDiagnosticsDTO) SetNodeSnapshots(v []ConnectionDiagnosticsSnapshotDTO) {
	o.NodeSnapshots = v
}

func (o ConnectionDiagnosticsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionDiagnosticsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !IsNil(o.AggregateSnapshot) {
		toSerialize["aggregateSnapshot"] = o.AggregateSnapshot
	}
	if !IsNil(o.NodeSnapshots) {
		toSerialize["nodeSnapshots"] = o.NodeSnapshots
	}
	return toSerialize, nil
}

type NullableConnectionDiagnosticsDTO struct {
	value *ConnectionDiagnosticsDTO
	isSet bool
}

func (v NullableConnectionDiagnosticsDTO) Get() *ConnectionDiagnosticsDTO {
	return v.value
}

func (v *NullableConnectionDiagnosticsDTO) Set(val *ConnectionDiagnosticsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionDiagnosticsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionDiagnosticsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionDiagnosticsDTO(val *ConnectionDiagnosticsDTO) *NullableConnectionDiagnosticsDTO {
	return &NullableConnectionDiagnosticsDTO{value: val, isSet: true}
}

func (v NullableConnectionDiagnosticsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionDiagnosticsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


