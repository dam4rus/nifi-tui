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

// checks if the VersionedProcessor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionedProcessor{}

// VersionedProcessor struct for VersionedProcessor
type VersionedProcessor struct {
	// The component's unique identifier
	Identifier *string `json:"identifier,omitempty"`
	// The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component
	InstanceIdentifier *string `json:"instanceIdentifier,omitempty"`
	// The component's name
	Name *string `json:"name,omitempty"`
	// The user-supplied comments for the component
	Comments *string `json:"comments,omitempty"`
	Position *Position `json:"position,omitempty"`
	// The type of the extension component
	Type *string `json:"type,omitempty"`
	Bundle *Bundle `json:"bundle,omitempty"`
	// The properties for the component. Properties whose value is not set will only contain the property name.
	Properties *map[string]string `json:"properties,omitempty"`
	// The property descriptors for the component.
	PropertyDescriptors *map[string]VersionedPropertyDescriptor `json:"propertyDescriptors,omitempty"`
	// Stylistic data for rendering in a UI
	Style *map[string]string `json:"style,omitempty"`
	// The annotation data for the processor used to relay configuration between a custom UI and the procesosr.
	AnnotationData *string `json:"annotationData,omitempty"`
	// The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy.
	SchedulingPeriod *string `json:"schedulingPeriod,omitempty"`
	// Indicates whether the processor should be scheduled to run in event or timer driven mode.
	SchedulingStrategy *string `json:"schedulingStrategy,omitempty"`
	// Indicates the node where the process will execute.
	ExecutionNode *string `json:"executionNode,omitempty"`
	// The amout of time that is used when the process penalizes a flowfile.
	PenaltyDuration *string `json:"penaltyDuration,omitempty"`
	// The amount of time that must elapse before this processor is scheduled again after yielding.
	YieldDuration *string `json:"yieldDuration,omitempty"`
	// The level at which the processor will report bulletins.
	BulletinLevel *string `json:"bulletinLevel,omitempty"`
	// The run duration for the processor in milliseconds.
	RunDurationMillis *int64 `json:"runDurationMillis,omitempty"`
	// The number of tasks that should be concurrently schedule for the processor. If the processor doesn't allow parallol processing then any positive input will be ignored.
	ConcurrentlySchedulableTaskCount *int32 `json:"concurrentlySchedulableTaskCount,omitempty"`
	// The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the 'isAutoTerminate' property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated.
	AutoTerminatedRelationships []string `json:"autoTerminatedRelationships,omitempty"`
	// The scheduled state of the component
	ScheduledState *string `json:"scheduledState,omitempty"`
	// Overall number of retries.
	RetryCount *int32 `json:"retryCount,omitempty"`
	// All the relationships should be retried.
	RetriedRelationships []string `json:"retriedRelationships,omitempty"`
	// Determines whether the FlowFile should be penalized or the processor should be yielded between retries.
	BackoffMechanism *string `json:"backoffMechanism,omitempty"`
	// Maximum amount of time to be waited during a retry period.
	MaxBackoffPeriod *string `json:"maxBackoffPeriod,omitempty"`
	ComponentType *string `json:"componentType,omitempty"`
	// The ID of the Process Group that this component belongs to
	GroupIdentifier *string `json:"groupIdentifier,omitempty"`
}

// NewVersionedProcessor instantiates a new VersionedProcessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedProcessor() *VersionedProcessor {
	this := VersionedProcessor{}
	return &this
}

// NewVersionedProcessorWithDefaults instantiates a new VersionedProcessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedProcessorWithDefaults() *VersionedProcessor {
	this := VersionedProcessor{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *VersionedProcessor) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *VersionedProcessor) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *VersionedProcessor) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetInstanceIdentifier returns the InstanceIdentifier field value if set, zero value otherwise.
func (o *VersionedProcessor) GetInstanceIdentifier() string {
	if o == nil || IsNil(o.InstanceIdentifier) {
		var ret string
		return ret
	}
	return *o.InstanceIdentifier
}

// GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetInstanceIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceIdentifier) {
		return nil, false
	}
	return o.InstanceIdentifier, true
}

// HasInstanceIdentifier returns a boolean if a field has been set.
func (o *VersionedProcessor) HasInstanceIdentifier() bool {
	if o != nil && !IsNil(o.InstanceIdentifier) {
		return true
	}

	return false
}

// SetInstanceIdentifier gets a reference to the given string and assigns it to the InstanceIdentifier field.
func (o *VersionedProcessor) SetInstanceIdentifier(v string) {
	o.InstanceIdentifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VersionedProcessor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VersionedProcessor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VersionedProcessor) SetName(v string) {
	o.Name = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VersionedProcessor) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VersionedProcessor) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VersionedProcessor) SetComments(v string) {
	o.Comments = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *VersionedProcessor) GetPosition() Position {
	if o == nil || IsNil(o.Position) {
		var ret Position
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetPositionOk() (*Position, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *VersionedProcessor) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given Position and assigns it to the Position field.
func (o *VersionedProcessor) SetPosition(v Position) {
	o.Position = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VersionedProcessor) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VersionedProcessor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VersionedProcessor) SetType(v string) {
	o.Type = &v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *VersionedProcessor) GetBundle() Bundle {
	if o == nil || IsNil(o.Bundle) {
		var ret Bundle
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetBundleOk() (*Bundle, bool) {
	if o == nil || IsNil(o.Bundle) {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *VersionedProcessor) HasBundle() bool {
	if o != nil && !IsNil(o.Bundle) {
		return true
	}

	return false
}

// SetBundle gets a reference to the given Bundle and assigns it to the Bundle field.
func (o *VersionedProcessor) SetBundle(v Bundle) {
	o.Bundle = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *VersionedProcessor) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *VersionedProcessor) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *VersionedProcessor) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetPropertyDescriptors returns the PropertyDescriptors field value if set, zero value otherwise.
func (o *VersionedProcessor) GetPropertyDescriptors() map[string]VersionedPropertyDescriptor {
	if o == nil || IsNil(o.PropertyDescriptors) {
		var ret map[string]VersionedPropertyDescriptor
		return ret
	}
	return *o.PropertyDescriptors
}

// GetPropertyDescriptorsOk returns a tuple with the PropertyDescriptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetPropertyDescriptorsOk() (*map[string]VersionedPropertyDescriptor, bool) {
	if o == nil || IsNil(o.PropertyDescriptors) {
		return nil, false
	}
	return o.PropertyDescriptors, true
}

// HasPropertyDescriptors returns a boolean if a field has been set.
func (o *VersionedProcessor) HasPropertyDescriptors() bool {
	if o != nil && !IsNil(o.PropertyDescriptors) {
		return true
	}

	return false
}

// SetPropertyDescriptors gets a reference to the given map[string]VersionedPropertyDescriptor and assigns it to the PropertyDescriptors field.
func (o *VersionedProcessor) SetPropertyDescriptors(v map[string]VersionedPropertyDescriptor) {
	o.PropertyDescriptors = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *VersionedProcessor) GetStyle() map[string]string {
	if o == nil || IsNil(o.Style) {
		var ret map[string]string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetStyleOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Style) {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *VersionedProcessor) HasStyle() bool {
	if o != nil && !IsNil(o.Style) {
		return true
	}

	return false
}

// SetStyle gets a reference to the given map[string]string and assigns it to the Style field.
func (o *VersionedProcessor) SetStyle(v map[string]string) {
	o.Style = &v
}

// GetAnnotationData returns the AnnotationData field value if set, zero value otherwise.
func (o *VersionedProcessor) GetAnnotationData() string {
	if o == nil || IsNil(o.AnnotationData) {
		var ret string
		return ret
	}
	return *o.AnnotationData
}

// GetAnnotationDataOk returns a tuple with the AnnotationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetAnnotationDataOk() (*string, bool) {
	if o == nil || IsNil(o.AnnotationData) {
		return nil, false
	}
	return o.AnnotationData, true
}

// HasAnnotationData returns a boolean if a field has been set.
func (o *VersionedProcessor) HasAnnotationData() bool {
	if o != nil && !IsNil(o.AnnotationData) {
		return true
	}

	return false
}

// SetAnnotationData gets a reference to the given string and assigns it to the AnnotationData field.
func (o *VersionedProcessor) SetAnnotationData(v string) {
	o.AnnotationData = &v
}

// GetSchedulingPeriod returns the SchedulingPeriod field value if set, zero value otherwise.
func (o *VersionedProcessor) GetSchedulingPeriod() string {
	if o == nil || IsNil(o.SchedulingPeriod) {
		var ret string
		return ret
	}
	return *o.SchedulingPeriod
}

// GetSchedulingPeriodOk returns a tuple with the SchedulingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetSchedulingPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.SchedulingPeriod) {
		return nil, false
	}
	return o.SchedulingPeriod, true
}

// HasSchedulingPeriod returns a boolean if a field has been set.
func (o *VersionedProcessor) HasSchedulingPeriod() bool {
	if o != nil && !IsNil(o.SchedulingPeriod) {
		return true
	}

	return false
}

// SetSchedulingPeriod gets a reference to the given string and assigns it to the SchedulingPeriod field.
func (o *VersionedProcessor) SetSchedulingPeriod(v string) {
	o.SchedulingPeriod = &v
}

// GetSchedulingStrategy returns the SchedulingStrategy field value if set, zero value otherwise.
func (o *VersionedProcessor) GetSchedulingStrategy() string {
	if o == nil || IsNil(o.SchedulingStrategy) {
		var ret string
		return ret
	}
	return *o.SchedulingStrategy
}

// GetSchedulingStrategyOk returns a tuple with the SchedulingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetSchedulingStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.SchedulingStrategy) {
		return nil, false
	}
	return o.SchedulingStrategy, true
}

// HasSchedulingStrategy returns a boolean if a field has been set.
func (o *VersionedProcessor) HasSchedulingStrategy() bool {
	if o != nil && !IsNil(o.SchedulingStrategy) {
		return true
	}

	return false
}

// SetSchedulingStrategy gets a reference to the given string and assigns it to the SchedulingStrategy field.
func (o *VersionedProcessor) SetSchedulingStrategy(v string) {
	o.SchedulingStrategy = &v
}

// GetExecutionNode returns the ExecutionNode field value if set, zero value otherwise.
func (o *VersionedProcessor) GetExecutionNode() string {
	if o == nil || IsNil(o.ExecutionNode) {
		var ret string
		return ret
	}
	return *o.ExecutionNode
}

// GetExecutionNodeOk returns a tuple with the ExecutionNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetExecutionNodeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionNode) {
		return nil, false
	}
	return o.ExecutionNode, true
}

// HasExecutionNode returns a boolean if a field has been set.
func (o *VersionedProcessor) HasExecutionNode() bool {
	if o != nil && !IsNil(o.ExecutionNode) {
		return true
	}

	return false
}

// SetExecutionNode gets a reference to the given string and assigns it to the ExecutionNode field.
func (o *VersionedProcessor) SetExecutionNode(v string) {
	o.ExecutionNode = &v
}

// GetPenaltyDuration returns the PenaltyDuration field value if set, zero value otherwise.
func (o *VersionedProcessor) GetPenaltyDuration() string {
	if o == nil || IsNil(o.PenaltyDuration) {
		var ret string
		return ret
	}
	return *o.PenaltyDuration
}

// GetPenaltyDurationOk returns a tuple with the PenaltyDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetPenaltyDurationOk() (*string, bool) {
	if o == nil || IsNil(o.PenaltyDuration) {
		return nil, false
	}
	return o.PenaltyDuration, true
}

// HasPenaltyDuration returns a boolean if a field has been set.
func (o *VersionedProcessor) HasPenaltyDuration() bool {
	if o != nil && !IsNil(o.PenaltyDuration) {
		return true
	}

	return false
}

// SetPenaltyDuration gets a reference to the given string and assigns it to the PenaltyDuration field.
func (o *VersionedProcessor) SetPenaltyDuration(v string) {
	o.PenaltyDuration = &v
}

// GetYieldDuration returns the YieldDuration field value if set, zero value otherwise.
func (o *VersionedProcessor) GetYieldDuration() string {
	if o == nil || IsNil(o.YieldDuration) {
		var ret string
		return ret
	}
	return *o.YieldDuration
}

// GetYieldDurationOk returns a tuple with the YieldDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetYieldDurationOk() (*string, bool) {
	if o == nil || IsNil(o.YieldDuration) {
		return nil, false
	}
	return o.YieldDuration, true
}

// HasYieldDuration returns a boolean if a field has been set.
func (o *VersionedProcessor) HasYieldDuration() bool {
	if o != nil && !IsNil(o.YieldDuration) {
		return true
	}

	return false
}

// SetYieldDuration gets a reference to the given string and assigns it to the YieldDuration field.
func (o *VersionedProcessor) SetYieldDuration(v string) {
	o.YieldDuration = &v
}

// GetBulletinLevel returns the BulletinLevel field value if set, zero value otherwise.
func (o *VersionedProcessor) GetBulletinLevel() string {
	if o == nil || IsNil(o.BulletinLevel) {
		var ret string
		return ret
	}
	return *o.BulletinLevel
}

// GetBulletinLevelOk returns a tuple with the BulletinLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetBulletinLevelOk() (*string, bool) {
	if o == nil || IsNil(o.BulletinLevel) {
		return nil, false
	}
	return o.BulletinLevel, true
}

// HasBulletinLevel returns a boolean if a field has been set.
func (o *VersionedProcessor) HasBulletinLevel() bool {
	if o != nil && !IsNil(o.BulletinLevel) {
		return true
	}

	return false
}

// SetBulletinLevel gets a reference to the given string and assigns it to the BulletinLevel field.
func (o *VersionedProcessor) SetBulletinLevel(v string) {
	o.BulletinLevel = &v
}

// GetRunDurationMillis returns the RunDurationMillis field value if set, zero value otherwise.
func (o *VersionedProcessor) GetRunDurationMillis() int64 {
	if o == nil || IsNil(o.RunDurationMillis) {
		var ret int64
		return ret
	}
	return *o.RunDurationMillis
}

// GetRunDurationMillisOk returns a tuple with the RunDurationMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetRunDurationMillisOk() (*int64, bool) {
	if o == nil || IsNil(o.RunDurationMillis) {
		return nil, false
	}
	return o.RunDurationMillis, true
}

// HasRunDurationMillis returns a boolean if a field has been set.
func (o *VersionedProcessor) HasRunDurationMillis() bool {
	if o != nil && !IsNil(o.RunDurationMillis) {
		return true
	}

	return false
}

// SetRunDurationMillis gets a reference to the given int64 and assigns it to the RunDurationMillis field.
func (o *VersionedProcessor) SetRunDurationMillis(v int64) {
	o.RunDurationMillis = &v
}

// GetConcurrentlySchedulableTaskCount returns the ConcurrentlySchedulableTaskCount field value if set, zero value otherwise.
func (o *VersionedProcessor) GetConcurrentlySchedulableTaskCount() int32 {
	if o == nil || IsNil(o.ConcurrentlySchedulableTaskCount) {
		var ret int32
		return ret
	}
	return *o.ConcurrentlySchedulableTaskCount
}

// GetConcurrentlySchedulableTaskCountOk returns a tuple with the ConcurrentlySchedulableTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetConcurrentlySchedulableTaskCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ConcurrentlySchedulableTaskCount) {
		return nil, false
	}
	return o.ConcurrentlySchedulableTaskCount, true
}

// HasConcurrentlySchedulableTaskCount returns a boolean if a field has been set.
func (o *VersionedProcessor) HasConcurrentlySchedulableTaskCount() bool {
	if o != nil && !IsNil(o.ConcurrentlySchedulableTaskCount) {
		return true
	}

	return false
}

// SetConcurrentlySchedulableTaskCount gets a reference to the given int32 and assigns it to the ConcurrentlySchedulableTaskCount field.
func (o *VersionedProcessor) SetConcurrentlySchedulableTaskCount(v int32) {
	o.ConcurrentlySchedulableTaskCount = &v
}

// GetAutoTerminatedRelationships returns the AutoTerminatedRelationships field value if set, zero value otherwise.
func (o *VersionedProcessor) GetAutoTerminatedRelationships() []string {
	if o == nil || IsNil(o.AutoTerminatedRelationships) {
		var ret []string
		return ret
	}
	return o.AutoTerminatedRelationships
}

// GetAutoTerminatedRelationshipsOk returns a tuple with the AutoTerminatedRelationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetAutoTerminatedRelationshipsOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoTerminatedRelationships) {
		return nil, false
	}
	return o.AutoTerminatedRelationships, true
}

// HasAutoTerminatedRelationships returns a boolean if a field has been set.
func (o *VersionedProcessor) HasAutoTerminatedRelationships() bool {
	if o != nil && !IsNil(o.AutoTerminatedRelationships) {
		return true
	}

	return false
}

// SetAutoTerminatedRelationships gets a reference to the given []string and assigns it to the AutoTerminatedRelationships field.
func (o *VersionedProcessor) SetAutoTerminatedRelationships(v []string) {
	o.AutoTerminatedRelationships = v
}

// GetScheduledState returns the ScheduledState field value if set, zero value otherwise.
func (o *VersionedProcessor) GetScheduledState() string {
	if o == nil || IsNil(o.ScheduledState) {
		var ret string
		return ret
	}
	return *o.ScheduledState
}

// GetScheduledStateOk returns a tuple with the ScheduledState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetScheduledStateOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduledState) {
		return nil, false
	}
	return o.ScheduledState, true
}

// HasScheduledState returns a boolean if a field has been set.
func (o *VersionedProcessor) HasScheduledState() bool {
	if o != nil && !IsNil(o.ScheduledState) {
		return true
	}

	return false
}

// SetScheduledState gets a reference to the given string and assigns it to the ScheduledState field.
func (o *VersionedProcessor) SetScheduledState(v string) {
	o.ScheduledState = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *VersionedProcessor) GetRetryCount() int32 {
	if o == nil || IsNil(o.RetryCount) {
		var ret int32
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetRetryCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryCount) {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *VersionedProcessor) HasRetryCount() bool {
	if o != nil && !IsNil(o.RetryCount) {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int32 and assigns it to the RetryCount field.
func (o *VersionedProcessor) SetRetryCount(v int32) {
	o.RetryCount = &v
}

// GetRetriedRelationships returns the RetriedRelationships field value if set, zero value otherwise.
func (o *VersionedProcessor) GetRetriedRelationships() []string {
	if o == nil || IsNil(o.RetriedRelationships) {
		var ret []string
		return ret
	}
	return o.RetriedRelationships
}

// GetRetriedRelationshipsOk returns a tuple with the RetriedRelationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetRetriedRelationshipsOk() ([]string, bool) {
	if o == nil || IsNil(o.RetriedRelationships) {
		return nil, false
	}
	return o.RetriedRelationships, true
}

// HasRetriedRelationships returns a boolean if a field has been set.
func (o *VersionedProcessor) HasRetriedRelationships() bool {
	if o != nil && !IsNil(o.RetriedRelationships) {
		return true
	}

	return false
}

// SetRetriedRelationships gets a reference to the given []string and assigns it to the RetriedRelationships field.
func (o *VersionedProcessor) SetRetriedRelationships(v []string) {
	o.RetriedRelationships = v
}

// GetBackoffMechanism returns the BackoffMechanism field value if set, zero value otherwise.
func (o *VersionedProcessor) GetBackoffMechanism() string {
	if o == nil || IsNil(o.BackoffMechanism) {
		var ret string
		return ret
	}
	return *o.BackoffMechanism
}

// GetBackoffMechanismOk returns a tuple with the BackoffMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetBackoffMechanismOk() (*string, bool) {
	if o == nil || IsNil(o.BackoffMechanism) {
		return nil, false
	}
	return o.BackoffMechanism, true
}

// HasBackoffMechanism returns a boolean if a field has been set.
func (o *VersionedProcessor) HasBackoffMechanism() bool {
	if o != nil && !IsNil(o.BackoffMechanism) {
		return true
	}

	return false
}

// SetBackoffMechanism gets a reference to the given string and assigns it to the BackoffMechanism field.
func (o *VersionedProcessor) SetBackoffMechanism(v string) {
	o.BackoffMechanism = &v
}

// GetMaxBackoffPeriod returns the MaxBackoffPeriod field value if set, zero value otherwise.
func (o *VersionedProcessor) GetMaxBackoffPeriod() string {
	if o == nil || IsNil(o.MaxBackoffPeriod) {
		var ret string
		return ret
	}
	return *o.MaxBackoffPeriod
}

// GetMaxBackoffPeriodOk returns a tuple with the MaxBackoffPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetMaxBackoffPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.MaxBackoffPeriod) {
		return nil, false
	}
	return o.MaxBackoffPeriod, true
}

// HasMaxBackoffPeriod returns a boolean if a field has been set.
func (o *VersionedProcessor) HasMaxBackoffPeriod() bool {
	if o != nil && !IsNil(o.MaxBackoffPeriod) {
		return true
	}

	return false
}

// SetMaxBackoffPeriod gets a reference to the given string and assigns it to the MaxBackoffPeriod field.
func (o *VersionedProcessor) SetMaxBackoffPeriod(v string) {
	o.MaxBackoffPeriod = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *VersionedProcessor) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType) {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetComponentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentType) {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *VersionedProcessor) HasComponentType() bool {
	if o != nil && !IsNil(o.ComponentType) {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *VersionedProcessor) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetGroupIdentifier returns the GroupIdentifier field value if set, zero value otherwise.
func (o *VersionedProcessor) GetGroupIdentifier() string {
	if o == nil || IsNil(o.GroupIdentifier) {
		var ret string
		return ret
	}
	return *o.GroupIdentifier
}

// GetGroupIdentifierOk returns a tuple with the GroupIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedProcessor) GetGroupIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.GroupIdentifier) {
		return nil, false
	}
	return o.GroupIdentifier, true
}

// HasGroupIdentifier returns a boolean if a field has been set.
func (o *VersionedProcessor) HasGroupIdentifier() bool {
	if o != nil && !IsNil(o.GroupIdentifier) {
		return true
	}

	return false
}

// SetGroupIdentifier gets a reference to the given string and assigns it to the GroupIdentifier field.
func (o *VersionedProcessor) SetGroupIdentifier(v string) {
	o.GroupIdentifier = &v
}

func (o VersionedProcessor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionedProcessor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.InstanceIdentifier) {
		toSerialize["instanceIdentifier"] = o.InstanceIdentifier
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Bundle) {
		toSerialize["bundle"] = o.Bundle
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.PropertyDescriptors) {
		toSerialize["propertyDescriptors"] = o.PropertyDescriptors
	}
	if !IsNil(o.Style) {
		toSerialize["style"] = o.Style
	}
	if !IsNil(o.AnnotationData) {
		toSerialize["annotationData"] = o.AnnotationData
	}
	if !IsNil(o.SchedulingPeriod) {
		toSerialize["schedulingPeriod"] = o.SchedulingPeriod
	}
	if !IsNil(o.SchedulingStrategy) {
		toSerialize["schedulingStrategy"] = o.SchedulingStrategy
	}
	if !IsNil(o.ExecutionNode) {
		toSerialize["executionNode"] = o.ExecutionNode
	}
	if !IsNil(o.PenaltyDuration) {
		toSerialize["penaltyDuration"] = o.PenaltyDuration
	}
	if !IsNil(o.YieldDuration) {
		toSerialize["yieldDuration"] = o.YieldDuration
	}
	if !IsNil(o.BulletinLevel) {
		toSerialize["bulletinLevel"] = o.BulletinLevel
	}
	if !IsNil(o.RunDurationMillis) {
		toSerialize["runDurationMillis"] = o.RunDurationMillis
	}
	if !IsNil(o.ConcurrentlySchedulableTaskCount) {
		toSerialize["concurrentlySchedulableTaskCount"] = o.ConcurrentlySchedulableTaskCount
	}
	if !IsNil(o.AutoTerminatedRelationships) {
		toSerialize["autoTerminatedRelationships"] = o.AutoTerminatedRelationships
	}
	if !IsNil(o.ScheduledState) {
		toSerialize["scheduledState"] = o.ScheduledState
	}
	if !IsNil(o.RetryCount) {
		toSerialize["retryCount"] = o.RetryCount
	}
	if !IsNil(o.RetriedRelationships) {
		toSerialize["retriedRelationships"] = o.RetriedRelationships
	}
	if !IsNil(o.BackoffMechanism) {
		toSerialize["backoffMechanism"] = o.BackoffMechanism
	}
	if !IsNil(o.MaxBackoffPeriod) {
		toSerialize["maxBackoffPeriod"] = o.MaxBackoffPeriod
	}
	if !IsNil(o.ComponentType) {
		toSerialize["componentType"] = o.ComponentType
	}
	if !IsNil(o.GroupIdentifier) {
		toSerialize["groupIdentifier"] = o.GroupIdentifier
	}
	return toSerialize, nil
}

type NullableVersionedProcessor struct {
	value *VersionedProcessor
	isSet bool
}

func (v NullableVersionedProcessor) Get() *VersionedProcessor {
	return v.value
}

func (v *NullableVersionedProcessor) Set(val *VersionedProcessor) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedProcessor) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedProcessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedProcessor(val *VersionedProcessor) *NullableVersionedProcessor {
	return &NullableVersionedProcessor{value: val, isSet: true}
}

func (v NullableVersionedProcessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedProcessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


