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

// checks if the VersionedFlowSnapshotEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionedFlowSnapshotEntity{}

// VersionedFlowSnapshotEntity struct for VersionedFlowSnapshotEntity
type VersionedFlowSnapshotEntity struct {
	VersionedFlowSnapshot *RegisteredFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
	ProcessGroupRevision *RevisionDTO `json:"processGroupRevision,omitempty"`
	// The ID of the Registry that this flow belongs to
	RegistryId *string `json:"registryId,omitempty"`
	// If the Process Group to be updated has a child or descendant Process Group that is also under Version Control, this specifies whether or not the contents of that child/descendant Process Group should be updated.
	UpdateDescendantVersionedFlows *bool `json:"updateDescendantVersionedFlows,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewVersionedFlowSnapshotEntity instantiates a new VersionedFlowSnapshotEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedFlowSnapshotEntity() *VersionedFlowSnapshotEntity {
	this := VersionedFlowSnapshotEntity{}
	return &this
}

// NewVersionedFlowSnapshotEntityWithDefaults instantiates a new VersionedFlowSnapshotEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedFlowSnapshotEntityWithDefaults() *VersionedFlowSnapshotEntity {
	this := VersionedFlowSnapshotEntity{}
	return &this
}

// GetVersionedFlowSnapshot returns the VersionedFlowSnapshot field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotEntity) GetVersionedFlowSnapshot() RegisteredFlowSnapshot {
	if o == nil || IsNil(o.VersionedFlowSnapshot) {
		var ret RegisteredFlowSnapshot
		return ret
	}
	return *o.VersionedFlowSnapshot
}

// GetVersionedFlowSnapshotOk returns a tuple with the VersionedFlowSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotEntity) GetVersionedFlowSnapshotOk() (*RegisteredFlowSnapshot, bool) {
	if o == nil || IsNil(o.VersionedFlowSnapshot) {
		return nil, false
	}
	return o.VersionedFlowSnapshot, true
}

// HasVersionedFlowSnapshot returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotEntity) HasVersionedFlowSnapshot() bool {
	if o != nil && !IsNil(o.VersionedFlowSnapshot) {
		return true
	}

	return false
}

// SetVersionedFlowSnapshot gets a reference to the given RegisteredFlowSnapshot and assigns it to the VersionedFlowSnapshot field.
func (o *VersionedFlowSnapshotEntity) SetVersionedFlowSnapshot(v RegisteredFlowSnapshot) {
	o.VersionedFlowSnapshot = &v
}

// GetProcessGroupRevision returns the ProcessGroupRevision field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotEntity) GetProcessGroupRevision() RevisionDTO {
	if o == nil || IsNil(o.ProcessGroupRevision) {
		var ret RevisionDTO
		return ret
	}
	return *o.ProcessGroupRevision
}

// GetProcessGroupRevisionOk returns a tuple with the ProcessGroupRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotEntity) GetProcessGroupRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.ProcessGroupRevision) {
		return nil, false
	}
	return o.ProcessGroupRevision, true
}

// HasProcessGroupRevision returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotEntity) HasProcessGroupRevision() bool {
	if o != nil && !IsNil(o.ProcessGroupRevision) {
		return true
	}

	return false
}

// SetProcessGroupRevision gets a reference to the given RevisionDTO and assigns it to the ProcessGroupRevision field.
func (o *VersionedFlowSnapshotEntity) SetProcessGroupRevision(v RevisionDTO) {
	o.ProcessGroupRevision = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotEntity) GetRegistryId() string {
	if o == nil || IsNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotEntity) GetRegistryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryId) {
		return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotEntity) HasRegistryId() bool {
	if o != nil && !IsNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *VersionedFlowSnapshotEntity) SetRegistryId(v string) {
	o.RegistryId = &v
}

// GetUpdateDescendantVersionedFlows returns the UpdateDescendantVersionedFlows field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotEntity) GetUpdateDescendantVersionedFlows() bool {
	if o == nil || IsNil(o.UpdateDescendantVersionedFlows) {
		var ret bool
		return ret
	}
	return *o.UpdateDescendantVersionedFlows
}

// GetUpdateDescendantVersionedFlowsOk returns a tuple with the UpdateDescendantVersionedFlows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotEntity) GetUpdateDescendantVersionedFlowsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateDescendantVersionedFlows) {
		return nil, false
	}
	return o.UpdateDescendantVersionedFlows, true
}

// HasUpdateDescendantVersionedFlows returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotEntity) HasUpdateDescendantVersionedFlows() bool {
	if o != nil && !IsNil(o.UpdateDescendantVersionedFlows) {
		return true
	}

	return false
}

// SetUpdateDescendantVersionedFlows gets a reference to the given bool and assigns it to the UpdateDescendantVersionedFlows field.
func (o *VersionedFlowSnapshotEntity) SetUpdateDescendantVersionedFlows(v bool) {
	o.UpdateDescendantVersionedFlows = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *VersionedFlowSnapshotEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowSnapshotEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *VersionedFlowSnapshotEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *VersionedFlowSnapshotEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o VersionedFlowSnapshotEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionedFlowSnapshotEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionedFlowSnapshot) {
		toSerialize["versionedFlowSnapshot"] = o.VersionedFlowSnapshot
	}
	if !IsNil(o.ProcessGroupRevision) {
		toSerialize["processGroupRevision"] = o.ProcessGroupRevision
	}
	if !IsNil(o.RegistryId) {
		toSerialize["registryId"] = o.RegistryId
	}
	if !IsNil(o.UpdateDescendantVersionedFlows) {
		toSerialize["updateDescendantVersionedFlows"] = o.UpdateDescendantVersionedFlows
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return toSerialize, nil
}

type NullableVersionedFlowSnapshotEntity struct {
	value *VersionedFlowSnapshotEntity
	isSet bool
}

func (v NullableVersionedFlowSnapshotEntity) Get() *VersionedFlowSnapshotEntity {
	return v.value
}

func (v *NullableVersionedFlowSnapshotEntity) Set(val *VersionedFlowSnapshotEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedFlowSnapshotEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedFlowSnapshotEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedFlowSnapshotEntity(val *VersionedFlowSnapshotEntity) *NullableVersionedFlowSnapshotEntity {
	return &NullableVersionedFlowSnapshotEntity{value: val, isSet: true}
}

func (v NullableVersionedFlowSnapshotEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedFlowSnapshotEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


