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

// checks if the ProcessGroupEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessGroupEntity{}

// ProcessGroupEntity struct for ProcessGroupEntity
type ProcessGroupEntity struct {
	Revision *RevisionDTO `json:"revision,omitempty"`
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri *string `json:"uri,omitempty"`
	Position *PositionDTO `json:"position,omitempty"`
	Permissions *PermissionsDTO `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
	Component *ProcessGroupDTO `json:"component,omitempty"`
	Status *ProcessGroupStatusDTO `json:"status,omitempty"`
	VersionedFlowSnapshot *RegisteredFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
	// The number of running components in this process group.
	RunningCount *int32 `json:"runningCount,omitempty"`
	// The number of stopped components in the process group.
	StoppedCount *int32 `json:"stoppedCount,omitempty"`
	// The number of invalid components in the process group.
	InvalidCount *int32 `json:"invalidCount,omitempty"`
	// The number of disabled components in the process group.
	DisabledCount *int32 `json:"disabledCount,omitempty"`
	// The number of active remote ports in the process group.
	ActiveRemotePortCount *int32 `json:"activeRemotePortCount,omitempty"`
	// The number of inactive remote ports in the process group.
	InactiveRemotePortCount *int32 `json:"inactiveRemotePortCount,omitempty"`
	// The current state of the Process Group, as it relates to the Versioned Flow
	VersionedFlowState *string `json:"versionedFlowState,omitempty"`
	// The number of up to date versioned process groups in the process group.
	UpToDateCount *int32 `json:"upToDateCount,omitempty"`
	// The number of locally modified versioned process groups in the process group.
	LocallyModifiedCount *int32 `json:"locallyModifiedCount,omitempty"`
	// The number of stale versioned process groups in the process group.
	StaleCount *int32 `json:"staleCount,omitempty"`
	// The number of locally modified and stale versioned process groups in the process group.
	LocallyModifiedAndStaleCount *int32 `json:"locallyModifiedAndStaleCount,omitempty"`
	// The number of versioned process groups in the process group that are unable to sync to a registry.
	SyncFailureCount *int32 `json:"syncFailureCount,omitempty"`
	// The number of local input ports in the process group.
	LocalInputPortCount *int32 `json:"localInputPortCount,omitempty"`
	// The number of local output ports in the process group.
	LocalOutputPortCount *int32 `json:"localOutputPortCount,omitempty"`
	// The number of public input ports in the process group.
	PublicInputPortCount *int32 `json:"publicInputPortCount,omitempty"`
	// The number of public output ports in the process group.
	PublicOutputPortCount *int32 `json:"publicOutputPortCount,omitempty"`
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	// The number of input ports in the process group.
	InputPortCount *int32 `json:"inputPortCount,omitempty"`
	// The number of output ports in the process group.
	OutputPortCount *int32 `json:"outputPortCount,omitempty"`
}

// NewProcessGroupEntity instantiates a new ProcessGroupEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupEntity() *ProcessGroupEntity {
	this := ProcessGroupEntity{}
	return &this
}

// NewProcessGroupEntityWithDefaults instantiates a new ProcessGroupEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupEntityWithDefaults() *ProcessGroupEntity {
	this := ProcessGroupEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *ProcessGroupEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessGroupEntity) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ProcessGroupEntity) SetUri(v string) {
	o.Uri = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPosition() PositionDTO {
	if o == nil || IsNil(o.Position) {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *ProcessGroupEntity) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPermissions() PermissionsDTO {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsDTO
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPermissionsOk() (*PermissionsDTO, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsDTO and assigns it to the Permissions field.
func (o *ProcessGroupEntity) SetPermissions(v PermissionsDTO) {
	o.Permissions = &v
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetBulletins() []BulletinEntity {
	if o == nil || IsNil(o.Bulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.Bulletins) {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasBulletins() bool {
	if o != nil && !IsNil(o.Bulletins) {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *ProcessGroupEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ProcessGroupEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetComponent() ProcessGroupDTO {
	if o == nil || IsNil(o.Component) {
		var ret ProcessGroupDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetComponentOk() (*ProcessGroupDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ProcessGroupDTO and assigns it to the Component field.
func (o *ProcessGroupEntity) SetComponent(v ProcessGroupDTO) {
	o.Component = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetStatus() ProcessGroupStatusDTO {
	if o == nil || IsNil(o.Status) {
		var ret ProcessGroupStatusDTO
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetStatusOk() (*ProcessGroupStatusDTO, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProcessGroupStatusDTO and assigns it to the Status field.
func (o *ProcessGroupEntity) SetStatus(v ProcessGroupStatusDTO) {
	o.Status = &v
}

// GetVersionedFlowSnapshot returns the VersionedFlowSnapshot field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetVersionedFlowSnapshot() RegisteredFlowSnapshot {
	if o == nil || IsNil(o.VersionedFlowSnapshot) {
		var ret RegisteredFlowSnapshot
		return ret
	}
	return *o.VersionedFlowSnapshot
}

// GetVersionedFlowSnapshotOk returns a tuple with the VersionedFlowSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetVersionedFlowSnapshotOk() (*RegisteredFlowSnapshot, bool) {
	if o == nil || IsNil(o.VersionedFlowSnapshot) {
		return nil, false
	}
	return o.VersionedFlowSnapshot, true
}

// HasVersionedFlowSnapshot returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasVersionedFlowSnapshot() bool {
	if o != nil && !IsNil(o.VersionedFlowSnapshot) {
		return true
	}

	return false
}

// SetVersionedFlowSnapshot gets a reference to the given RegisteredFlowSnapshot and assigns it to the VersionedFlowSnapshot field.
func (o *ProcessGroupEntity) SetVersionedFlowSnapshot(v RegisteredFlowSnapshot) {
	o.VersionedFlowSnapshot = &v
}

// GetRunningCount returns the RunningCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetRunningCount() int32 {
	if o == nil || IsNil(o.RunningCount) {
		var ret int32
		return ret
	}
	return *o.RunningCount
}

// GetRunningCountOk returns a tuple with the RunningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetRunningCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RunningCount) {
		return nil, false
	}
	return o.RunningCount, true
}

// HasRunningCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasRunningCount() bool {
	if o != nil && !IsNil(o.RunningCount) {
		return true
	}

	return false
}

// SetRunningCount gets a reference to the given int32 and assigns it to the RunningCount field.
func (o *ProcessGroupEntity) SetRunningCount(v int32) {
	o.RunningCount = &v
}

// GetStoppedCount returns the StoppedCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetStoppedCount() int32 {
	if o == nil || IsNil(o.StoppedCount) {
		var ret int32
		return ret
	}
	return *o.StoppedCount
}

// GetStoppedCountOk returns a tuple with the StoppedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetStoppedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StoppedCount) {
		return nil, false
	}
	return o.StoppedCount, true
}

// HasStoppedCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasStoppedCount() bool {
	if o != nil && !IsNil(o.StoppedCount) {
		return true
	}

	return false
}

// SetStoppedCount gets a reference to the given int32 and assigns it to the StoppedCount field.
func (o *ProcessGroupEntity) SetStoppedCount(v int32) {
	o.StoppedCount = &v
}

// GetInvalidCount returns the InvalidCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetInvalidCount() int32 {
	if o == nil || IsNil(o.InvalidCount) {
		var ret int32
		return ret
	}
	return *o.InvalidCount
}

// GetInvalidCountOk returns a tuple with the InvalidCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetInvalidCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InvalidCount) {
		return nil, false
	}
	return o.InvalidCount, true
}

// HasInvalidCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasInvalidCount() bool {
	if o != nil && !IsNil(o.InvalidCount) {
		return true
	}

	return false
}

// SetInvalidCount gets a reference to the given int32 and assigns it to the InvalidCount field.
func (o *ProcessGroupEntity) SetInvalidCount(v int32) {
	o.InvalidCount = &v
}

// GetDisabledCount returns the DisabledCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetDisabledCount() int32 {
	if o == nil || IsNil(o.DisabledCount) {
		var ret int32
		return ret
	}
	return *o.DisabledCount
}

// GetDisabledCountOk returns a tuple with the DisabledCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetDisabledCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DisabledCount) {
		return nil, false
	}
	return o.DisabledCount, true
}

// HasDisabledCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasDisabledCount() bool {
	if o != nil && !IsNil(o.DisabledCount) {
		return true
	}

	return false
}

// SetDisabledCount gets a reference to the given int32 and assigns it to the DisabledCount field.
func (o *ProcessGroupEntity) SetDisabledCount(v int32) {
	o.DisabledCount = &v
}

// GetActiveRemotePortCount returns the ActiveRemotePortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetActiveRemotePortCount() int32 {
	if o == nil || IsNil(o.ActiveRemotePortCount) {
		var ret int32
		return ret
	}
	return *o.ActiveRemotePortCount
}

// GetActiveRemotePortCountOk returns a tuple with the ActiveRemotePortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetActiveRemotePortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveRemotePortCount) {
		return nil, false
	}
	return o.ActiveRemotePortCount, true
}

// HasActiveRemotePortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasActiveRemotePortCount() bool {
	if o != nil && !IsNil(o.ActiveRemotePortCount) {
		return true
	}

	return false
}

// SetActiveRemotePortCount gets a reference to the given int32 and assigns it to the ActiveRemotePortCount field.
func (o *ProcessGroupEntity) SetActiveRemotePortCount(v int32) {
	o.ActiveRemotePortCount = &v
}

// GetInactiveRemotePortCount returns the InactiveRemotePortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetInactiveRemotePortCount() int32 {
	if o == nil || IsNil(o.InactiveRemotePortCount) {
		var ret int32
		return ret
	}
	return *o.InactiveRemotePortCount
}

// GetInactiveRemotePortCountOk returns a tuple with the InactiveRemotePortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetInactiveRemotePortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InactiveRemotePortCount) {
		return nil, false
	}
	return o.InactiveRemotePortCount, true
}

// HasInactiveRemotePortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasInactiveRemotePortCount() bool {
	if o != nil && !IsNil(o.InactiveRemotePortCount) {
		return true
	}

	return false
}

// SetInactiveRemotePortCount gets a reference to the given int32 and assigns it to the InactiveRemotePortCount field.
func (o *ProcessGroupEntity) SetInactiveRemotePortCount(v int32) {
	o.InactiveRemotePortCount = &v
}

// GetVersionedFlowState returns the VersionedFlowState field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetVersionedFlowState() string {
	if o == nil || IsNil(o.VersionedFlowState) {
		var ret string
		return ret
	}
	return *o.VersionedFlowState
}

// GetVersionedFlowStateOk returns a tuple with the VersionedFlowState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetVersionedFlowStateOk() (*string, bool) {
	if o == nil || IsNil(o.VersionedFlowState) {
		return nil, false
	}
	return o.VersionedFlowState, true
}

// HasVersionedFlowState returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasVersionedFlowState() bool {
	if o != nil && !IsNil(o.VersionedFlowState) {
		return true
	}

	return false
}

// SetVersionedFlowState gets a reference to the given string and assigns it to the VersionedFlowState field.
func (o *ProcessGroupEntity) SetVersionedFlowState(v string) {
	o.VersionedFlowState = &v
}

// GetUpToDateCount returns the UpToDateCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetUpToDateCount() int32 {
	if o == nil || IsNil(o.UpToDateCount) {
		var ret int32
		return ret
	}
	return *o.UpToDateCount
}

// GetUpToDateCountOk returns a tuple with the UpToDateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetUpToDateCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UpToDateCount) {
		return nil, false
	}
	return o.UpToDateCount, true
}

// HasUpToDateCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasUpToDateCount() bool {
	if o != nil && !IsNil(o.UpToDateCount) {
		return true
	}

	return false
}

// SetUpToDateCount gets a reference to the given int32 and assigns it to the UpToDateCount field.
func (o *ProcessGroupEntity) SetUpToDateCount(v int32) {
	o.UpToDateCount = &v
}

// GetLocallyModifiedCount returns the LocallyModifiedCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocallyModifiedCount() int32 {
	if o == nil || IsNil(o.LocallyModifiedCount) {
		var ret int32
		return ret
	}
	return *o.LocallyModifiedCount
}

// GetLocallyModifiedCountOk returns a tuple with the LocallyModifiedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocallyModifiedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LocallyModifiedCount) {
		return nil, false
	}
	return o.LocallyModifiedCount, true
}

// HasLocallyModifiedCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocallyModifiedCount() bool {
	if o != nil && !IsNil(o.LocallyModifiedCount) {
		return true
	}

	return false
}

// SetLocallyModifiedCount gets a reference to the given int32 and assigns it to the LocallyModifiedCount field.
func (o *ProcessGroupEntity) SetLocallyModifiedCount(v int32) {
	o.LocallyModifiedCount = &v
}

// GetStaleCount returns the StaleCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetStaleCount() int32 {
	if o == nil || IsNil(o.StaleCount) {
		var ret int32
		return ret
	}
	return *o.StaleCount
}

// GetStaleCountOk returns a tuple with the StaleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetStaleCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StaleCount) {
		return nil, false
	}
	return o.StaleCount, true
}

// HasStaleCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasStaleCount() bool {
	if o != nil && !IsNil(o.StaleCount) {
		return true
	}

	return false
}

// SetStaleCount gets a reference to the given int32 and assigns it to the StaleCount field.
func (o *ProcessGroupEntity) SetStaleCount(v int32) {
	o.StaleCount = &v
}

// GetLocallyModifiedAndStaleCount returns the LocallyModifiedAndStaleCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocallyModifiedAndStaleCount() int32 {
	if o == nil || IsNil(o.LocallyModifiedAndStaleCount) {
		var ret int32
		return ret
	}
	return *o.LocallyModifiedAndStaleCount
}

// GetLocallyModifiedAndStaleCountOk returns a tuple with the LocallyModifiedAndStaleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocallyModifiedAndStaleCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LocallyModifiedAndStaleCount) {
		return nil, false
	}
	return o.LocallyModifiedAndStaleCount, true
}

// HasLocallyModifiedAndStaleCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocallyModifiedAndStaleCount() bool {
	if o != nil && !IsNil(o.LocallyModifiedAndStaleCount) {
		return true
	}

	return false
}

// SetLocallyModifiedAndStaleCount gets a reference to the given int32 and assigns it to the LocallyModifiedAndStaleCount field.
func (o *ProcessGroupEntity) SetLocallyModifiedAndStaleCount(v int32) {
	o.LocallyModifiedAndStaleCount = &v
}

// GetSyncFailureCount returns the SyncFailureCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetSyncFailureCount() int32 {
	if o == nil || IsNil(o.SyncFailureCount) {
		var ret int32
		return ret
	}
	return *o.SyncFailureCount
}

// GetSyncFailureCountOk returns a tuple with the SyncFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetSyncFailureCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SyncFailureCount) {
		return nil, false
	}
	return o.SyncFailureCount, true
}

// HasSyncFailureCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasSyncFailureCount() bool {
	if o != nil && !IsNil(o.SyncFailureCount) {
		return true
	}

	return false
}

// SetSyncFailureCount gets a reference to the given int32 and assigns it to the SyncFailureCount field.
func (o *ProcessGroupEntity) SetSyncFailureCount(v int32) {
	o.SyncFailureCount = &v
}

// GetLocalInputPortCount returns the LocalInputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocalInputPortCount() int32 {
	if o == nil || IsNil(o.LocalInputPortCount) {
		var ret int32
		return ret
	}
	return *o.LocalInputPortCount
}

// GetLocalInputPortCountOk returns a tuple with the LocalInputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocalInputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalInputPortCount) {
		return nil, false
	}
	return o.LocalInputPortCount, true
}

// HasLocalInputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocalInputPortCount() bool {
	if o != nil && !IsNil(o.LocalInputPortCount) {
		return true
	}

	return false
}

// SetLocalInputPortCount gets a reference to the given int32 and assigns it to the LocalInputPortCount field.
func (o *ProcessGroupEntity) SetLocalInputPortCount(v int32) {
	o.LocalInputPortCount = &v
}

// GetLocalOutputPortCount returns the LocalOutputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetLocalOutputPortCount() int32 {
	if o == nil || IsNil(o.LocalOutputPortCount) {
		var ret int32
		return ret
	}
	return *o.LocalOutputPortCount
}

// GetLocalOutputPortCountOk returns a tuple with the LocalOutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetLocalOutputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalOutputPortCount) {
		return nil, false
	}
	return o.LocalOutputPortCount, true
}

// HasLocalOutputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasLocalOutputPortCount() bool {
	if o != nil && !IsNil(o.LocalOutputPortCount) {
		return true
	}

	return false
}

// SetLocalOutputPortCount gets a reference to the given int32 and assigns it to the LocalOutputPortCount field.
func (o *ProcessGroupEntity) SetLocalOutputPortCount(v int32) {
	o.LocalOutputPortCount = &v
}

// GetPublicInputPortCount returns the PublicInputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPublicInputPortCount() int32 {
	if o == nil || IsNil(o.PublicInputPortCount) {
		var ret int32
		return ret
	}
	return *o.PublicInputPortCount
}

// GetPublicInputPortCountOk returns a tuple with the PublicInputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPublicInputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicInputPortCount) {
		return nil, false
	}
	return o.PublicInputPortCount, true
}

// HasPublicInputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPublicInputPortCount() bool {
	if o != nil && !IsNil(o.PublicInputPortCount) {
		return true
	}

	return false
}

// SetPublicInputPortCount gets a reference to the given int32 and assigns it to the PublicInputPortCount field.
func (o *ProcessGroupEntity) SetPublicInputPortCount(v int32) {
	o.PublicInputPortCount = &v
}

// GetPublicOutputPortCount returns the PublicOutputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetPublicOutputPortCount() int32 {
	if o == nil || IsNil(o.PublicOutputPortCount) {
		var ret int32
		return ret
	}
	return *o.PublicOutputPortCount
}

// GetPublicOutputPortCountOk returns a tuple with the PublicOutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetPublicOutputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicOutputPortCount) {
		return nil, false
	}
	return o.PublicOutputPortCount, true
}

// HasPublicOutputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasPublicOutputPortCount() bool {
	if o != nil && !IsNil(o.PublicOutputPortCount) {
		return true
	}

	return false
}

// SetPublicOutputPortCount gets a reference to the given int32 and assigns it to the PublicOutputPortCount field.
func (o *ProcessGroupEntity) SetPublicOutputPortCount(v int32) {
	o.PublicOutputPortCount = &v
}

// GetParameterContext returns the ParameterContext field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetParameterContext() ParameterContextReferenceEntity {
	if o == nil || IsNil(o.ParameterContext) {
		var ret ParameterContextReferenceEntity
		return ret
	}
	return *o.ParameterContext
}

// GetParameterContextOk returns a tuple with the ParameterContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetParameterContextOk() (*ParameterContextReferenceEntity, bool) {
	if o == nil || IsNil(o.ParameterContext) {
		return nil, false
	}
	return o.ParameterContext, true
}

// HasParameterContext returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasParameterContext() bool {
	if o != nil && !IsNil(o.ParameterContext) {
		return true
	}

	return false
}

// SetParameterContext gets a reference to the given ParameterContextReferenceEntity and assigns it to the ParameterContext field.
func (o *ProcessGroupEntity) SetParameterContext(v ParameterContextReferenceEntity) {
	o.ParameterContext = &v
}

// GetInputPortCount returns the InputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetInputPortCount() int32 {
	if o == nil || IsNil(o.InputPortCount) {
		var ret int32
		return ret
	}
	return *o.InputPortCount
}

// GetInputPortCountOk returns a tuple with the InputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetInputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InputPortCount) {
		return nil, false
	}
	return o.InputPortCount, true
}

// HasInputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasInputPortCount() bool {
	if o != nil && !IsNil(o.InputPortCount) {
		return true
	}

	return false
}

// SetInputPortCount gets a reference to the given int32 and assigns it to the InputPortCount field.
func (o *ProcessGroupEntity) SetInputPortCount(v int32) {
	o.InputPortCount = &v
}

// GetOutputPortCount returns the OutputPortCount field value if set, zero value otherwise.
func (o *ProcessGroupEntity) GetOutputPortCount() int32 {
	if o == nil || IsNil(o.OutputPortCount) {
		var ret int32
		return ret
	}
	return *o.OutputPortCount
}

// GetOutputPortCountOk returns a tuple with the OutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupEntity) GetOutputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputPortCount) {
		return nil, false
	}
	return o.OutputPortCount, true
}

// HasOutputPortCount returns a boolean if a field has been set.
func (o *ProcessGroupEntity) HasOutputPortCount() bool {
	if o != nil && !IsNil(o.OutputPortCount) {
		return true
	}

	return false
}

// SetOutputPortCount gets a reference to the given int32 and assigns it to the OutputPortCount field.
func (o *ProcessGroupEntity) SetOutputPortCount(v int32) {
	o.OutputPortCount = &v
}

func (o ProcessGroupEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessGroupEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Bulletins) {
		toSerialize["bulletins"] = o.Bulletins
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.VersionedFlowSnapshot) {
		toSerialize["versionedFlowSnapshot"] = o.VersionedFlowSnapshot
	}
	if !IsNil(o.RunningCount) {
		toSerialize["runningCount"] = o.RunningCount
	}
	if !IsNil(o.StoppedCount) {
		toSerialize["stoppedCount"] = o.StoppedCount
	}
	if !IsNil(o.InvalidCount) {
		toSerialize["invalidCount"] = o.InvalidCount
	}
	if !IsNil(o.DisabledCount) {
		toSerialize["disabledCount"] = o.DisabledCount
	}
	if !IsNil(o.ActiveRemotePortCount) {
		toSerialize["activeRemotePortCount"] = o.ActiveRemotePortCount
	}
	if !IsNil(o.InactiveRemotePortCount) {
		toSerialize["inactiveRemotePortCount"] = o.InactiveRemotePortCount
	}
	if !IsNil(o.VersionedFlowState) {
		toSerialize["versionedFlowState"] = o.VersionedFlowState
	}
	if !IsNil(o.UpToDateCount) {
		toSerialize["upToDateCount"] = o.UpToDateCount
	}
	if !IsNil(o.LocallyModifiedCount) {
		toSerialize["locallyModifiedCount"] = o.LocallyModifiedCount
	}
	if !IsNil(o.StaleCount) {
		toSerialize["staleCount"] = o.StaleCount
	}
	if !IsNil(o.LocallyModifiedAndStaleCount) {
		toSerialize["locallyModifiedAndStaleCount"] = o.LocallyModifiedAndStaleCount
	}
	if !IsNil(o.SyncFailureCount) {
		toSerialize["syncFailureCount"] = o.SyncFailureCount
	}
	if !IsNil(o.LocalInputPortCount) {
		toSerialize["localInputPortCount"] = o.LocalInputPortCount
	}
	if !IsNil(o.LocalOutputPortCount) {
		toSerialize["localOutputPortCount"] = o.LocalOutputPortCount
	}
	if !IsNil(o.PublicInputPortCount) {
		toSerialize["publicInputPortCount"] = o.PublicInputPortCount
	}
	if !IsNil(o.PublicOutputPortCount) {
		toSerialize["publicOutputPortCount"] = o.PublicOutputPortCount
	}
	if !IsNil(o.ParameterContext) {
		toSerialize["parameterContext"] = o.ParameterContext
	}
	if !IsNil(o.InputPortCount) {
		toSerialize["inputPortCount"] = o.InputPortCount
	}
	if !IsNil(o.OutputPortCount) {
		toSerialize["outputPortCount"] = o.OutputPortCount
	}
	return toSerialize, nil
}

type NullableProcessGroupEntity struct {
	value *ProcessGroupEntity
	isSet bool
}

func (v NullableProcessGroupEntity) Get() *ProcessGroupEntity {
	return v.value
}

func (v *NullableProcessGroupEntity) Set(val *ProcessGroupEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupEntity(val *ProcessGroupEntity) *NullableProcessGroupEntity {
	return &NullableProcessGroupEntity{value: val, isSet: true}
}

func (v NullableProcessGroupEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


