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

// checks if the ProcessorStatusSnapshotDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessorStatusSnapshotDTO{}

// ProcessorStatusSnapshotDTO struct for ProcessorStatusSnapshotDTO
type ProcessorStatusSnapshotDTO struct {
	// The id of the processor.
	Id *string `json:"id,omitempty"`
	// The id of the parent process group to which the processor belongs.
	GroupId *string `json:"groupId,omitempty"`
	// The name of the prcessor.
	Name *string `json:"name,omitempty"`
	// The type of the processor.
	Type *string `json:"type,omitempty"`
	// The state of the processor.
	RunStatus *string `json:"runStatus,omitempty"`
	// Indicates the node where the process will execute.
	ExecutionNode *string `json:"executionNode,omitempty"`
	// The number of bytes read by this Processor in the last 5 mintues
	BytesRead *int64 `json:"bytesRead,omitempty"`
	// The number of bytes written by this Processor in the last 5 minutes
	BytesWritten *int64 `json:"bytesWritten,omitempty"`
	// The number of bytes read in the last 5 minutes.
	Read *string `json:"read,omitempty"`
	// The number of bytes written in the last 5 minutes.
	Written *string `json:"written,omitempty"`
	// The number of FlowFiles that have been accepted in the last 5 minutes
	FlowFilesIn *int32 `json:"flowFilesIn,omitempty"`
	// The size of the FlowFiles that have been accepted in the last 5 minutes
	BytesIn *int64 `json:"bytesIn,omitempty"`
	// The count/size of flowfiles that have been accepted in the last 5 minutes.
	Input *string `json:"input,omitempty"`
	// The number of FlowFiles transferred to a Connection in the last 5 minutes
	FlowFilesOut *int32 `json:"flowFilesOut,omitempty"`
	// The size of the FlowFiles transferred to a Connection in the last 5 minutes
	BytesOut *int64 `json:"bytesOut,omitempty"`
	// The count/size of flowfiles that have been processed in the last 5 minutes.
	Output *string `json:"output,omitempty"`
	// The number of times this Processor has run in the last 5 minutes
	TaskCount *int32 `json:"taskCount,omitempty"`
	// The number of nanoseconds that this Processor has spent running in the last 5 minutes
	TasksDurationNanos *int64 `json:"tasksDurationNanos,omitempty"`
	// The total number of task this connectable has completed over the last 5 minutes.
	Tasks *string `json:"tasks,omitempty"`
	// The total duration of all tasks for this connectable over the last 5 minutes.
	TasksDuration *string `json:"tasksDuration,omitempty"`
	// The number of threads currently executing in the processor.
	ActiveThreadCount *int32 `json:"activeThreadCount,omitempty"`
	// The number of threads currently terminated for the processor.
	TerminatedThreadCount *int32 `json:"terminatedThreadCount,omitempty"`
}

// NewProcessorStatusSnapshotDTO instantiates a new ProcessorStatusSnapshotDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorStatusSnapshotDTO() *ProcessorStatusSnapshotDTO {
	this := ProcessorStatusSnapshotDTO{}
	return &this
}

// NewProcessorStatusSnapshotDTOWithDefaults instantiates a new ProcessorStatusSnapshotDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorStatusSnapshotDTOWithDefaults() *ProcessorStatusSnapshotDTO {
	this := ProcessorStatusSnapshotDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessorStatusSnapshotDTO) SetId(v string) {
	o.Id = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ProcessorStatusSnapshotDTO) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProcessorStatusSnapshotDTO) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProcessorStatusSnapshotDTO) SetType(v string) {
	o.Type = &v
}

// GetRunStatus returns the RunStatus field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetRunStatus() string {
	if o == nil || IsNil(o.RunStatus) {
		var ret string
		return ret
	}
	return *o.RunStatus
}

// GetRunStatusOk returns a tuple with the RunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetRunStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RunStatus) {
		return nil, false
	}
	return o.RunStatus, true
}

// HasRunStatus returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasRunStatus() bool {
	if o != nil && !IsNil(o.RunStatus) {
		return true
	}

	return false
}

// SetRunStatus gets a reference to the given string and assigns it to the RunStatus field.
func (o *ProcessorStatusSnapshotDTO) SetRunStatus(v string) {
	o.RunStatus = &v
}

// GetExecutionNode returns the ExecutionNode field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetExecutionNode() string {
	if o == nil || IsNil(o.ExecutionNode) {
		var ret string
		return ret
	}
	return *o.ExecutionNode
}

// GetExecutionNodeOk returns a tuple with the ExecutionNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetExecutionNodeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionNode) {
		return nil, false
	}
	return o.ExecutionNode, true
}

// HasExecutionNode returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasExecutionNode() bool {
	if o != nil && !IsNil(o.ExecutionNode) {
		return true
	}

	return false
}

// SetExecutionNode gets a reference to the given string and assigns it to the ExecutionNode field.
func (o *ProcessorStatusSnapshotDTO) SetExecutionNode(v string) {
	o.ExecutionNode = &v
}

// GetBytesRead returns the BytesRead field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetBytesRead() int64 {
	if o == nil || IsNil(o.BytesRead) {
		var ret int64
		return ret
	}
	return *o.BytesRead
}

// GetBytesReadOk returns a tuple with the BytesRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetBytesReadOk() (*int64, bool) {
	if o == nil || IsNil(o.BytesRead) {
		return nil, false
	}
	return o.BytesRead, true
}

// HasBytesRead returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasBytesRead() bool {
	if o != nil && !IsNil(o.BytesRead) {
		return true
	}

	return false
}

// SetBytesRead gets a reference to the given int64 and assigns it to the BytesRead field.
func (o *ProcessorStatusSnapshotDTO) SetBytesRead(v int64) {
	o.BytesRead = &v
}

// GetBytesWritten returns the BytesWritten field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetBytesWritten() int64 {
	if o == nil || IsNil(o.BytesWritten) {
		var ret int64
		return ret
	}
	return *o.BytesWritten
}

// GetBytesWrittenOk returns a tuple with the BytesWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetBytesWrittenOk() (*int64, bool) {
	if o == nil || IsNil(o.BytesWritten) {
		return nil, false
	}
	return o.BytesWritten, true
}

// HasBytesWritten returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasBytesWritten() bool {
	if o != nil && !IsNil(o.BytesWritten) {
		return true
	}

	return false
}

// SetBytesWritten gets a reference to the given int64 and assigns it to the BytesWritten field.
func (o *ProcessorStatusSnapshotDTO) SetBytesWritten(v int64) {
	o.BytesWritten = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetRead() string {
	if o == nil || IsNil(o.Read) {
		var ret string
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetReadOk() (*string, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given string and assigns it to the Read field.
func (o *ProcessorStatusSnapshotDTO) SetRead(v string) {
	o.Read = &v
}

// GetWritten returns the Written field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetWritten() string {
	if o == nil || IsNil(o.Written) {
		var ret string
		return ret
	}
	return *o.Written
}

// GetWrittenOk returns a tuple with the Written field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetWrittenOk() (*string, bool) {
	if o == nil || IsNil(o.Written) {
		return nil, false
	}
	return o.Written, true
}

// HasWritten returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasWritten() bool {
	if o != nil && !IsNil(o.Written) {
		return true
	}

	return false
}

// SetWritten gets a reference to the given string and assigns it to the Written field.
func (o *ProcessorStatusSnapshotDTO) SetWritten(v string) {
	o.Written = &v
}

// GetFlowFilesIn returns the FlowFilesIn field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetFlowFilesIn() int32 {
	if o == nil || IsNil(o.FlowFilesIn) {
		var ret int32
		return ret
	}
	return *o.FlowFilesIn
}

// GetFlowFilesInOk returns a tuple with the FlowFilesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetFlowFilesInOk() (*int32, bool) {
	if o == nil || IsNil(o.FlowFilesIn) {
		return nil, false
	}
	return o.FlowFilesIn, true
}

// HasFlowFilesIn returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasFlowFilesIn() bool {
	if o != nil && !IsNil(o.FlowFilesIn) {
		return true
	}

	return false
}

// SetFlowFilesIn gets a reference to the given int32 and assigns it to the FlowFilesIn field.
func (o *ProcessorStatusSnapshotDTO) SetFlowFilesIn(v int32) {
	o.FlowFilesIn = &v
}

// GetBytesIn returns the BytesIn field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetBytesIn() int64 {
	if o == nil || IsNil(o.BytesIn) {
		var ret int64
		return ret
	}
	return *o.BytesIn
}

// GetBytesInOk returns a tuple with the BytesIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetBytesInOk() (*int64, bool) {
	if o == nil || IsNil(o.BytesIn) {
		return nil, false
	}
	return o.BytesIn, true
}

// HasBytesIn returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasBytesIn() bool {
	if o != nil && !IsNil(o.BytesIn) {
		return true
	}

	return false
}

// SetBytesIn gets a reference to the given int64 and assigns it to the BytesIn field.
func (o *ProcessorStatusSnapshotDTO) SetBytesIn(v int64) {
	o.BytesIn = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetInput() string {
	if o == nil || IsNil(o.Input) {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetInputOk() (*string, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *ProcessorStatusSnapshotDTO) SetInput(v string) {
	o.Input = &v
}

// GetFlowFilesOut returns the FlowFilesOut field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetFlowFilesOut() int32 {
	if o == nil || IsNil(o.FlowFilesOut) {
		var ret int32
		return ret
	}
	return *o.FlowFilesOut
}

// GetFlowFilesOutOk returns a tuple with the FlowFilesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetFlowFilesOutOk() (*int32, bool) {
	if o == nil || IsNil(o.FlowFilesOut) {
		return nil, false
	}
	return o.FlowFilesOut, true
}

// HasFlowFilesOut returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasFlowFilesOut() bool {
	if o != nil && !IsNil(o.FlowFilesOut) {
		return true
	}

	return false
}

// SetFlowFilesOut gets a reference to the given int32 and assigns it to the FlowFilesOut field.
func (o *ProcessorStatusSnapshotDTO) SetFlowFilesOut(v int32) {
	o.FlowFilesOut = &v
}

// GetBytesOut returns the BytesOut field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetBytesOut() int64 {
	if o == nil || IsNil(o.BytesOut) {
		var ret int64
		return ret
	}
	return *o.BytesOut
}

// GetBytesOutOk returns a tuple with the BytesOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetBytesOutOk() (*int64, bool) {
	if o == nil || IsNil(o.BytesOut) {
		return nil, false
	}
	return o.BytesOut, true
}

// HasBytesOut returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasBytesOut() bool {
	if o != nil && !IsNil(o.BytesOut) {
		return true
	}

	return false
}

// SetBytesOut gets a reference to the given int64 and assigns it to the BytesOut field.
func (o *ProcessorStatusSnapshotDTO) SetBytesOut(v int64) {
	o.BytesOut = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetOutput() string {
	if o == nil || IsNil(o.Output) {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetOutputOk() (*string, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *ProcessorStatusSnapshotDTO) SetOutput(v string) {
	o.Output = &v
}

// GetTaskCount returns the TaskCount field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetTaskCount() int32 {
	if o == nil || IsNil(o.TaskCount) {
		var ret int32
		return ret
	}
	return *o.TaskCount
}

// GetTaskCountOk returns a tuple with the TaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetTaskCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskCount) {
		return nil, false
	}
	return o.TaskCount, true
}

// HasTaskCount returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasTaskCount() bool {
	if o != nil && !IsNil(o.TaskCount) {
		return true
	}

	return false
}

// SetTaskCount gets a reference to the given int32 and assigns it to the TaskCount field.
func (o *ProcessorStatusSnapshotDTO) SetTaskCount(v int32) {
	o.TaskCount = &v
}

// GetTasksDurationNanos returns the TasksDurationNanos field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetTasksDurationNanos() int64 {
	if o == nil || IsNil(o.TasksDurationNanos) {
		var ret int64
		return ret
	}
	return *o.TasksDurationNanos
}

// GetTasksDurationNanosOk returns a tuple with the TasksDurationNanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetTasksDurationNanosOk() (*int64, bool) {
	if o == nil || IsNil(o.TasksDurationNanos) {
		return nil, false
	}
	return o.TasksDurationNanos, true
}

// HasTasksDurationNanos returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasTasksDurationNanos() bool {
	if o != nil && !IsNil(o.TasksDurationNanos) {
		return true
	}

	return false
}

// SetTasksDurationNanos gets a reference to the given int64 and assigns it to the TasksDurationNanos field.
func (o *ProcessorStatusSnapshotDTO) SetTasksDurationNanos(v int64) {
	o.TasksDurationNanos = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetTasks() string {
	if o == nil || IsNil(o.Tasks) {
		var ret string
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetTasksOk() (*string, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given string and assigns it to the Tasks field.
func (o *ProcessorStatusSnapshotDTO) SetTasks(v string) {
	o.Tasks = &v
}

// GetTasksDuration returns the TasksDuration field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetTasksDuration() string {
	if o == nil || IsNil(o.TasksDuration) {
		var ret string
		return ret
	}
	return *o.TasksDuration
}

// GetTasksDurationOk returns a tuple with the TasksDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetTasksDurationOk() (*string, bool) {
	if o == nil || IsNil(o.TasksDuration) {
		return nil, false
	}
	return o.TasksDuration, true
}

// HasTasksDuration returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasTasksDuration() bool {
	if o != nil && !IsNil(o.TasksDuration) {
		return true
	}

	return false
}

// SetTasksDuration gets a reference to the given string and assigns it to the TasksDuration field.
func (o *ProcessorStatusSnapshotDTO) SetTasksDuration(v string) {
	o.TasksDuration = &v
}

// GetActiveThreadCount returns the ActiveThreadCount field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetActiveThreadCount() int32 {
	if o == nil || IsNil(o.ActiveThreadCount) {
		var ret int32
		return ret
	}
	return *o.ActiveThreadCount
}

// GetActiveThreadCountOk returns a tuple with the ActiveThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetActiveThreadCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveThreadCount) {
		return nil, false
	}
	return o.ActiveThreadCount, true
}

// HasActiveThreadCount returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasActiveThreadCount() bool {
	if o != nil && !IsNil(o.ActiveThreadCount) {
		return true
	}

	return false
}

// SetActiveThreadCount gets a reference to the given int32 and assigns it to the ActiveThreadCount field.
func (o *ProcessorStatusSnapshotDTO) SetActiveThreadCount(v int32) {
	o.ActiveThreadCount = &v
}

// GetTerminatedThreadCount returns the TerminatedThreadCount field value if set, zero value otherwise.
func (o *ProcessorStatusSnapshotDTO) GetTerminatedThreadCount() int32 {
	if o == nil || IsNil(o.TerminatedThreadCount) {
		var ret int32
		return ret
	}
	return *o.TerminatedThreadCount
}

// GetTerminatedThreadCountOk returns a tuple with the TerminatedThreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorStatusSnapshotDTO) GetTerminatedThreadCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TerminatedThreadCount) {
		return nil, false
	}
	return o.TerminatedThreadCount, true
}

// HasTerminatedThreadCount returns a boolean if a field has been set.
func (o *ProcessorStatusSnapshotDTO) HasTerminatedThreadCount() bool {
	if o != nil && !IsNil(o.TerminatedThreadCount) {
		return true
	}

	return false
}

// SetTerminatedThreadCount gets a reference to the given int32 and assigns it to the TerminatedThreadCount field.
func (o *ProcessorStatusSnapshotDTO) SetTerminatedThreadCount(v int32) {
	o.TerminatedThreadCount = &v
}

func (o ProcessorStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessorStatusSnapshotDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RunStatus) {
		toSerialize["runStatus"] = o.RunStatus
	}
	if !IsNil(o.ExecutionNode) {
		toSerialize["executionNode"] = o.ExecutionNode
	}
	if !IsNil(o.BytesRead) {
		toSerialize["bytesRead"] = o.BytesRead
	}
	if !IsNil(o.BytesWritten) {
		toSerialize["bytesWritten"] = o.BytesWritten
	}
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	if !IsNil(o.Written) {
		toSerialize["written"] = o.Written
	}
	if !IsNil(o.FlowFilesIn) {
		toSerialize["flowFilesIn"] = o.FlowFilesIn
	}
	if !IsNil(o.BytesIn) {
		toSerialize["bytesIn"] = o.BytesIn
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.FlowFilesOut) {
		toSerialize["flowFilesOut"] = o.FlowFilesOut
	}
	if !IsNil(o.BytesOut) {
		toSerialize["bytesOut"] = o.BytesOut
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.TaskCount) {
		toSerialize["taskCount"] = o.TaskCount
	}
	if !IsNil(o.TasksDurationNanos) {
		toSerialize["tasksDurationNanos"] = o.TasksDurationNanos
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.TasksDuration) {
		toSerialize["tasksDuration"] = o.TasksDuration
	}
	if !IsNil(o.ActiveThreadCount) {
		toSerialize["activeThreadCount"] = o.ActiveThreadCount
	}
	if !IsNil(o.TerminatedThreadCount) {
		toSerialize["terminatedThreadCount"] = o.TerminatedThreadCount
	}
	return toSerialize, nil
}

type NullableProcessorStatusSnapshotDTO struct {
	value *ProcessorStatusSnapshotDTO
	isSet bool
}

func (v NullableProcessorStatusSnapshotDTO) Get() *ProcessorStatusSnapshotDTO {
	return v.value
}

func (v *NullableProcessorStatusSnapshotDTO) Set(val *ProcessorStatusSnapshotDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorStatusSnapshotDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorStatusSnapshotDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorStatusSnapshotDTO(val *ProcessorStatusSnapshotDTO) *NullableProcessorStatusSnapshotDTO {
	return &NullableProcessorStatusSnapshotDTO{value: val, isSet: true}
}

func (v NullableProcessorStatusSnapshotDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorStatusSnapshotDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


