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

// checks if the ConnectionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionDTO{}

// ConnectionDTO struct for ConnectionDTO
type ConnectionDTO struct {
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId *string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId *string `json:"parentGroupId,omitempty"`
	Position *PositionDTO `json:"position,omitempty"`
	Source *ConnectableDTO `json:"source,omitempty"`
	Destination *ConnectableDTO `json:"destination,omitempty"`
	// The name of the connection.
	Name *string `json:"name,omitempty"`
	// The index of the bend point where to place the connection label.
	LabelIndex *int32 `json:"labelIndex,omitempty"`
	// The z index of the connection.
	GetzIndex *int64 `json:"getzIndex,omitempty"`
	// The selected relationship that comprise the connection.
	SelectedRelationships []string `json:"selectedRelationships,omitempty"`
	// The relationships that the source of the connection currently supports.
	AvailableRelationships []string `json:"availableRelationships,omitempty"`
	// The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureObjectThreshold *int64 `json:"backPressureObjectThreshold,omitempty"`
	// The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureDataSizeThreshold *string `json:"backPressureDataSizeThreshold,omitempty"`
	// The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it.
	FlowFileExpiration *string `json:"flowFileExpiration,omitempty"`
	// The comparators used to prioritize the queue.
	Prioritizers []string `json:"prioritizers,omitempty"`
	// The bend points on the connection.
	Bends []PositionDTO `json:"bends,omitempty"`
	// How to load balance the data in this Connection across the nodes in the cluster.
	LoadBalanceStrategy *string `json:"loadBalanceStrategy,omitempty"`
	// The FlowFile Attribute to use for determining which node a FlowFile will go to if the Load Balancing Strategy is set to PARTITION_BY_ATTRIBUTE
	LoadBalancePartitionAttribute *string `json:"loadBalancePartitionAttribute,omitempty"`
	// Whether or not data should be compressed when being transferred between nodes in the cluster.
	LoadBalanceCompression *string `json:"loadBalanceCompression,omitempty"`
	// The current status of the Connection's Load Balancing Activities. Status can indicate that Load Balancing is not configured for the connection, that Load Balancing is configured but inactive (not currently transferring data to another node), or that Load Balancing is configured and actively transferring data to another node.
	LoadBalanceStatus *string `json:"loadBalanceStatus,omitempty"`
}

// NewConnectionDTO instantiates a new ConnectionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionDTO() *ConnectionDTO {
	this := ConnectionDTO{}
	return &this
}

// NewConnectionDTOWithDefaults instantiates a new ConnectionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionDTOWithDefaults() *ConnectionDTO {
	this := ConnectionDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionDTO) SetId(v string) {
	o.Id = &v
}

// GetVersionedComponentId returns the VersionedComponentId field value if set, zero value otherwise.
func (o *ConnectionDTO) GetVersionedComponentId() string {
	if o == nil || IsNil(o.VersionedComponentId) {
		var ret string
		return ret
	}
	return *o.VersionedComponentId
}

// GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetVersionedComponentIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionedComponentId) {
		return nil, false
	}
	return o.VersionedComponentId, true
}

// HasVersionedComponentId returns a boolean if a field has been set.
func (o *ConnectionDTO) HasVersionedComponentId() bool {
	if o != nil && !IsNil(o.VersionedComponentId) {
		return true
	}

	return false
}

// SetVersionedComponentId gets a reference to the given string and assigns it to the VersionedComponentId field.
func (o *ConnectionDTO) SetVersionedComponentId(v string) {
	o.VersionedComponentId = &v
}

// GetParentGroupId returns the ParentGroupId field value if set, zero value otherwise.
func (o *ConnectionDTO) GetParentGroupId() string {
	if o == nil || IsNil(o.ParentGroupId) {
		var ret string
		return ret
	}
	return *o.ParentGroupId
}

// GetParentGroupIdOk returns a tuple with the ParentGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetParentGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentGroupId) {
		return nil, false
	}
	return o.ParentGroupId, true
}

// HasParentGroupId returns a boolean if a field has been set.
func (o *ConnectionDTO) HasParentGroupId() bool {
	if o != nil && !IsNil(o.ParentGroupId) {
		return true
	}

	return false
}

// SetParentGroupId gets a reference to the given string and assigns it to the ParentGroupId field.
func (o *ConnectionDTO) SetParentGroupId(v string) {
	o.ParentGroupId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ConnectionDTO) GetPosition() PositionDTO {
	if o == nil || IsNil(o.Position) {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ConnectionDTO) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *ConnectionDTO) SetPosition(v PositionDTO) {
	o.Position = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConnectionDTO) GetSource() ConnectableDTO {
	if o == nil || IsNil(o.Source) {
		var ret ConnectableDTO
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetSourceOk() (*ConnectableDTO, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConnectionDTO) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ConnectableDTO and assigns it to the Source field.
func (o *ConnectionDTO) SetSource(v ConnectableDTO) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ConnectionDTO) GetDestination() ConnectableDTO {
	if o == nil || IsNil(o.Destination) {
		var ret ConnectableDTO
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetDestinationOk() (*ConnectableDTO, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ConnectionDTO) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given ConnectableDTO and assigns it to the Destination field.
func (o *ConnectionDTO) SetDestination(v ConnectableDTO) {
	o.Destination = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionDTO) SetName(v string) {
	o.Name = &v
}

// GetLabelIndex returns the LabelIndex field value if set, zero value otherwise.
func (o *ConnectionDTO) GetLabelIndex() int32 {
	if o == nil || IsNil(o.LabelIndex) {
		var ret int32
		return ret
	}
	return *o.LabelIndex
}

// GetLabelIndexOk returns a tuple with the LabelIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetLabelIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.LabelIndex) {
		return nil, false
	}
	return o.LabelIndex, true
}

// HasLabelIndex returns a boolean if a field has been set.
func (o *ConnectionDTO) HasLabelIndex() bool {
	if o != nil && !IsNil(o.LabelIndex) {
		return true
	}

	return false
}

// SetLabelIndex gets a reference to the given int32 and assigns it to the LabelIndex field.
func (o *ConnectionDTO) SetLabelIndex(v int32) {
	o.LabelIndex = &v
}

// GetGetzIndex returns the GetzIndex field value if set, zero value otherwise.
func (o *ConnectionDTO) GetGetzIndex() int64 {
	if o == nil || IsNil(o.GetzIndex) {
		var ret int64
		return ret
	}
	return *o.GetzIndex
}

// GetGetzIndexOk returns a tuple with the GetzIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetGetzIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.GetzIndex) {
		return nil, false
	}
	return o.GetzIndex, true
}

// HasGetzIndex returns a boolean if a field has been set.
func (o *ConnectionDTO) HasGetzIndex() bool {
	if o != nil && !IsNil(o.GetzIndex) {
		return true
	}

	return false
}

// SetGetzIndex gets a reference to the given int64 and assigns it to the GetzIndex field.
func (o *ConnectionDTO) SetGetzIndex(v int64) {
	o.GetzIndex = &v
}

// GetSelectedRelationships returns the SelectedRelationships field value if set, zero value otherwise.
func (o *ConnectionDTO) GetSelectedRelationships() []string {
	if o == nil || IsNil(o.SelectedRelationships) {
		var ret []string
		return ret
	}
	return o.SelectedRelationships
}

// GetSelectedRelationshipsOk returns a tuple with the SelectedRelationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetSelectedRelationshipsOk() ([]string, bool) {
	if o == nil || IsNil(o.SelectedRelationships) {
		return nil, false
	}
	return o.SelectedRelationships, true
}

// HasSelectedRelationships returns a boolean if a field has been set.
func (o *ConnectionDTO) HasSelectedRelationships() bool {
	if o != nil && !IsNil(o.SelectedRelationships) {
		return true
	}

	return false
}

// SetSelectedRelationships gets a reference to the given []string and assigns it to the SelectedRelationships field.
func (o *ConnectionDTO) SetSelectedRelationships(v []string) {
	o.SelectedRelationships = v
}

// GetAvailableRelationships returns the AvailableRelationships field value if set, zero value otherwise.
func (o *ConnectionDTO) GetAvailableRelationships() []string {
	if o == nil || IsNil(o.AvailableRelationships) {
		var ret []string
		return ret
	}
	return o.AvailableRelationships
}

// GetAvailableRelationshipsOk returns a tuple with the AvailableRelationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetAvailableRelationshipsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableRelationships) {
		return nil, false
	}
	return o.AvailableRelationships, true
}

// HasAvailableRelationships returns a boolean if a field has been set.
func (o *ConnectionDTO) HasAvailableRelationships() bool {
	if o != nil && !IsNil(o.AvailableRelationships) {
		return true
	}

	return false
}

// SetAvailableRelationships gets a reference to the given []string and assigns it to the AvailableRelationships field.
func (o *ConnectionDTO) SetAvailableRelationships(v []string) {
	o.AvailableRelationships = v
}

// GetBackPressureObjectThreshold returns the BackPressureObjectThreshold field value if set, zero value otherwise.
func (o *ConnectionDTO) GetBackPressureObjectThreshold() int64 {
	if o == nil || IsNil(o.BackPressureObjectThreshold) {
		var ret int64
		return ret
	}
	return *o.BackPressureObjectThreshold
}

// GetBackPressureObjectThresholdOk returns a tuple with the BackPressureObjectThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetBackPressureObjectThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.BackPressureObjectThreshold) {
		return nil, false
	}
	return o.BackPressureObjectThreshold, true
}

// HasBackPressureObjectThreshold returns a boolean if a field has been set.
func (o *ConnectionDTO) HasBackPressureObjectThreshold() bool {
	if o != nil && !IsNil(o.BackPressureObjectThreshold) {
		return true
	}

	return false
}

// SetBackPressureObjectThreshold gets a reference to the given int64 and assigns it to the BackPressureObjectThreshold field.
func (o *ConnectionDTO) SetBackPressureObjectThreshold(v int64) {
	o.BackPressureObjectThreshold = &v
}

// GetBackPressureDataSizeThreshold returns the BackPressureDataSizeThreshold field value if set, zero value otherwise.
func (o *ConnectionDTO) GetBackPressureDataSizeThreshold() string {
	if o == nil || IsNil(o.BackPressureDataSizeThreshold) {
		var ret string
		return ret
	}
	return *o.BackPressureDataSizeThreshold
}

// GetBackPressureDataSizeThresholdOk returns a tuple with the BackPressureDataSizeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetBackPressureDataSizeThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.BackPressureDataSizeThreshold) {
		return nil, false
	}
	return o.BackPressureDataSizeThreshold, true
}

// HasBackPressureDataSizeThreshold returns a boolean if a field has been set.
func (o *ConnectionDTO) HasBackPressureDataSizeThreshold() bool {
	if o != nil && !IsNil(o.BackPressureDataSizeThreshold) {
		return true
	}

	return false
}

// SetBackPressureDataSizeThreshold gets a reference to the given string and assigns it to the BackPressureDataSizeThreshold field.
func (o *ConnectionDTO) SetBackPressureDataSizeThreshold(v string) {
	o.BackPressureDataSizeThreshold = &v
}

// GetFlowFileExpiration returns the FlowFileExpiration field value if set, zero value otherwise.
func (o *ConnectionDTO) GetFlowFileExpiration() string {
	if o == nil || IsNil(o.FlowFileExpiration) {
		var ret string
		return ret
	}
	return *o.FlowFileExpiration
}

// GetFlowFileExpirationOk returns a tuple with the FlowFileExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetFlowFileExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.FlowFileExpiration) {
		return nil, false
	}
	return o.FlowFileExpiration, true
}

// HasFlowFileExpiration returns a boolean if a field has been set.
func (o *ConnectionDTO) HasFlowFileExpiration() bool {
	if o != nil && !IsNil(o.FlowFileExpiration) {
		return true
	}

	return false
}

// SetFlowFileExpiration gets a reference to the given string and assigns it to the FlowFileExpiration field.
func (o *ConnectionDTO) SetFlowFileExpiration(v string) {
	o.FlowFileExpiration = &v
}

// GetPrioritizers returns the Prioritizers field value if set, zero value otherwise.
func (o *ConnectionDTO) GetPrioritizers() []string {
	if o == nil || IsNil(o.Prioritizers) {
		var ret []string
		return ret
	}
	return o.Prioritizers
}

// GetPrioritizersOk returns a tuple with the Prioritizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetPrioritizersOk() ([]string, bool) {
	if o == nil || IsNil(o.Prioritizers) {
		return nil, false
	}
	return o.Prioritizers, true
}

// HasPrioritizers returns a boolean if a field has been set.
func (o *ConnectionDTO) HasPrioritizers() bool {
	if o != nil && !IsNil(o.Prioritizers) {
		return true
	}

	return false
}

// SetPrioritizers gets a reference to the given []string and assigns it to the Prioritizers field.
func (o *ConnectionDTO) SetPrioritizers(v []string) {
	o.Prioritizers = v
}

// GetBends returns the Bends field value if set, zero value otherwise.
func (o *ConnectionDTO) GetBends() []PositionDTO {
	if o == nil || IsNil(o.Bends) {
		var ret []PositionDTO
		return ret
	}
	return o.Bends
}

// GetBendsOk returns a tuple with the Bends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetBendsOk() ([]PositionDTO, bool) {
	if o == nil || IsNil(o.Bends) {
		return nil, false
	}
	return o.Bends, true
}

// HasBends returns a boolean if a field has been set.
func (o *ConnectionDTO) HasBends() bool {
	if o != nil && !IsNil(o.Bends) {
		return true
	}

	return false
}

// SetBends gets a reference to the given []PositionDTO and assigns it to the Bends field.
func (o *ConnectionDTO) SetBends(v []PositionDTO) {
	o.Bends = v
}

// GetLoadBalanceStrategy returns the LoadBalanceStrategy field value if set, zero value otherwise.
func (o *ConnectionDTO) GetLoadBalanceStrategy() string {
	if o == nil || IsNil(o.LoadBalanceStrategy) {
		var ret string
		return ret
	}
	return *o.LoadBalanceStrategy
}

// GetLoadBalanceStrategyOk returns a tuple with the LoadBalanceStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetLoadBalanceStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalanceStrategy) {
		return nil, false
	}
	return o.LoadBalanceStrategy, true
}

// HasLoadBalanceStrategy returns a boolean if a field has been set.
func (o *ConnectionDTO) HasLoadBalanceStrategy() bool {
	if o != nil && !IsNil(o.LoadBalanceStrategy) {
		return true
	}

	return false
}

// SetLoadBalanceStrategy gets a reference to the given string and assigns it to the LoadBalanceStrategy field.
func (o *ConnectionDTO) SetLoadBalanceStrategy(v string) {
	o.LoadBalanceStrategy = &v
}

// GetLoadBalancePartitionAttribute returns the LoadBalancePartitionAttribute field value if set, zero value otherwise.
func (o *ConnectionDTO) GetLoadBalancePartitionAttribute() string {
	if o == nil || IsNil(o.LoadBalancePartitionAttribute) {
		var ret string
		return ret
	}
	return *o.LoadBalancePartitionAttribute
}

// GetLoadBalancePartitionAttributeOk returns a tuple with the LoadBalancePartitionAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetLoadBalancePartitionAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalancePartitionAttribute) {
		return nil, false
	}
	return o.LoadBalancePartitionAttribute, true
}

// HasLoadBalancePartitionAttribute returns a boolean if a field has been set.
func (o *ConnectionDTO) HasLoadBalancePartitionAttribute() bool {
	if o != nil && !IsNil(o.LoadBalancePartitionAttribute) {
		return true
	}

	return false
}

// SetLoadBalancePartitionAttribute gets a reference to the given string and assigns it to the LoadBalancePartitionAttribute field.
func (o *ConnectionDTO) SetLoadBalancePartitionAttribute(v string) {
	o.LoadBalancePartitionAttribute = &v
}

// GetLoadBalanceCompression returns the LoadBalanceCompression field value if set, zero value otherwise.
func (o *ConnectionDTO) GetLoadBalanceCompression() string {
	if o == nil || IsNil(o.LoadBalanceCompression) {
		var ret string
		return ret
	}
	return *o.LoadBalanceCompression
}

// GetLoadBalanceCompressionOk returns a tuple with the LoadBalanceCompression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetLoadBalanceCompressionOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalanceCompression) {
		return nil, false
	}
	return o.LoadBalanceCompression, true
}

// HasLoadBalanceCompression returns a boolean if a field has been set.
func (o *ConnectionDTO) HasLoadBalanceCompression() bool {
	if o != nil && !IsNil(o.LoadBalanceCompression) {
		return true
	}

	return false
}

// SetLoadBalanceCompression gets a reference to the given string and assigns it to the LoadBalanceCompression field.
func (o *ConnectionDTO) SetLoadBalanceCompression(v string) {
	o.LoadBalanceCompression = &v
}

// GetLoadBalanceStatus returns the LoadBalanceStatus field value if set, zero value otherwise.
func (o *ConnectionDTO) GetLoadBalanceStatus() string {
	if o == nil || IsNil(o.LoadBalanceStatus) {
		var ret string
		return ret
	}
	return *o.LoadBalanceStatus
}

// GetLoadBalanceStatusOk returns a tuple with the LoadBalanceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDTO) GetLoadBalanceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalanceStatus) {
		return nil, false
	}
	return o.LoadBalanceStatus, true
}

// HasLoadBalanceStatus returns a boolean if a field has been set.
func (o *ConnectionDTO) HasLoadBalanceStatus() bool {
	if o != nil && !IsNil(o.LoadBalanceStatus) {
		return true
	}

	return false
}

// SetLoadBalanceStatus gets a reference to the given string and assigns it to the LoadBalanceStatus field.
func (o *ConnectionDTO) SetLoadBalanceStatus(v string) {
	o.LoadBalanceStatus = &v
}

func (o ConnectionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.VersionedComponentId) {
		toSerialize["versionedComponentId"] = o.VersionedComponentId
	}
	if !IsNil(o.ParentGroupId) {
		toSerialize["parentGroupId"] = o.ParentGroupId
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LabelIndex) {
		toSerialize["labelIndex"] = o.LabelIndex
	}
	if !IsNil(o.GetzIndex) {
		toSerialize["getzIndex"] = o.GetzIndex
	}
	if !IsNil(o.SelectedRelationships) {
		toSerialize["selectedRelationships"] = o.SelectedRelationships
	}
	if !IsNil(o.AvailableRelationships) {
		toSerialize["availableRelationships"] = o.AvailableRelationships
	}
	if !IsNil(o.BackPressureObjectThreshold) {
		toSerialize["backPressureObjectThreshold"] = o.BackPressureObjectThreshold
	}
	if !IsNil(o.BackPressureDataSizeThreshold) {
		toSerialize["backPressureDataSizeThreshold"] = o.BackPressureDataSizeThreshold
	}
	if !IsNil(o.FlowFileExpiration) {
		toSerialize["flowFileExpiration"] = o.FlowFileExpiration
	}
	if !IsNil(o.Prioritizers) {
		toSerialize["prioritizers"] = o.Prioritizers
	}
	if !IsNil(o.Bends) {
		toSerialize["bends"] = o.Bends
	}
	if !IsNil(o.LoadBalanceStrategy) {
		toSerialize["loadBalanceStrategy"] = o.LoadBalanceStrategy
	}
	if !IsNil(o.LoadBalancePartitionAttribute) {
		toSerialize["loadBalancePartitionAttribute"] = o.LoadBalancePartitionAttribute
	}
	if !IsNil(o.LoadBalanceCompression) {
		toSerialize["loadBalanceCompression"] = o.LoadBalanceCompression
	}
	if !IsNil(o.LoadBalanceStatus) {
		toSerialize["loadBalanceStatus"] = o.LoadBalanceStatus
	}
	return toSerialize, nil
}

type NullableConnectionDTO struct {
	value *ConnectionDTO
	isSet bool
}

func (v NullableConnectionDTO) Get() *ConnectionDTO {
	return v.value
}

func (v *NullableConnectionDTO) Set(val *ConnectionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionDTO(val *ConnectionDTO) *NullableConnectionDTO {
	return &NullableConnectionDTO{value: val, isSet: true}
}

func (v NullableConnectionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


