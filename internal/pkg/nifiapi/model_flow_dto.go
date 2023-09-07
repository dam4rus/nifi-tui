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

// checks if the FlowDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowDTO{}

// FlowDTO struct for FlowDTO
type FlowDTO struct {
	// The process groups in this flow.
	ProcessGroups []ProcessGroupEntity `json:"processGroups,omitempty"`
	// The remote process groups in this flow.
	RemoteProcessGroups []RemoteProcessGroupEntity `json:"remoteProcessGroups,omitempty"`
	// The processors in this flow.
	Processors []ProcessorEntity `json:"processors,omitempty"`
	// The input ports in this flow.
	InputPorts []PortEntity `json:"inputPorts,omitempty"`
	// The output ports in this flow.
	OutputPorts []PortEntity `json:"outputPorts,omitempty"`
	// The connections in this flow.
	Connections []ConnectionEntity `json:"connections,omitempty"`
	// The labels in this flow.
	Labels []LabelEntity `json:"labels,omitempty"`
	// The funnels in this flow.
	Funnels []FunnelEntity `json:"funnels,omitempty"`
}

// NewFlowDTO instantiates a new FlowDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowDTO() *FlowDTO {
	this := FlowDTO{}
	return &this
}

// NewFlowDTOWithDefaults instantiates a new FlowDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowDTOWithDefaults() *FlowDTO {
	this := FlowDTO{}
	return &this
}

// GetProcessGroups returns the ProcessGroups field value if set, zero value otherwise.
func (o *FlowDTO) GetProcessGroups() []ProcessGroupEntity {
	if o == nil || IsNil(o.ProcessGroups) {
		var ret []ProcessGroupEntity
		return ret
	}
	return o.ProcessGroups
}

// GetProcessGroupsOk returns a tuple with the ProcessGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetProcessGroupsOk() ([]ProcessGroupEntity, bool) {
	if o == nil || IsNil(o.ProcessGroups) {
		return nil, false
	}
	return o.ProcessGroups, true
}

// HasProcessGroups returns a boolean if a field has been set.
func (o *FlowDTO) HasProcessGroups() bool {
	if o != nil && !IsNil(o.ProcessGroups) {
		return true
	}

	return false
}

// SetProcessGroups gets a reference to the given []ProcessGroupEntity and assigns it to the ProcessGroups field.
func (o *FlowDTO) SetProcessGroups(v []ProcessGroupEntity) {
	o.ProcessGroups = v
}

// GetRemoteProcessGroups returns the RemoteProcessGroups field value if set, zero value otherwise.
func (o *FlowDTO) GetRemoteProcessGroups() []RemoteProcessGroupEntity {
	if o == nil || IsNil(o.RemoteProcessGroups) {
		var ret []RemoteProcessGroupEntity
		return ret
	}
	return o.RemoteProcessGroups
}

// GetRemoteProcessGroupsOk returns a tuple with the RemoteProcessGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetRemoteProcessGroupsOk() ([]RemoteProcessGroupEntity, bool) {
	if o == nil || IsNil(o.RemoteProcessGroups) {
		return nil, false
	}
	return o.RemoteProcessGroups, true
}

// HasRemoteProcessGroups returns a boolean if a field has been set.
func (o *FlowDTO) HasRemoteProcessGroups() bool {
	if o != nil && !IsNil(o.RemoteProcessGroups) {
		return true
	}

	return false
}

// SetRemoteProcessGroups gets a reference to the given []RemoteProcessGroupEntity and assigns it to the RemoteProcessGroups field.
func (o *FlowDTO) SetRemoteProcessGroups(v []RemoteProcessGroupEntity) {
	o.RemoteProcessGroups = v
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *FlowDTO) GetProcessors() []ProcessorEntity {
	if o == nil || IsNil(o.Processors) {
		var ret []ProcessorEntity
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetProcessorsOk() ([]ProcessorEntity, bool) {
	if o == nil || IsNil(o.Processors) {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *FlowDTO) HasProcessors() bool {
	if o != nil && !IsNil(o.Processors) {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []ProcessorEntity and assigns it to the Processors field.
func (o *FlowDTO) SetProcessors(v []ProcessorEntity) {
	o.Processors = v
}

// GetInputPorts returns the InputPorts field value if set, zero value otherwise.
func (o *FlowDTO) GetInputPorts() []PortEntity {
	if o == nil || IsNil(o.InputPorts) {
		var ret []PortEntity
		return ret
	}
	return o.InputPorts
}

// GetInputPortsOk returns a tuple with the InputPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetInputPortsOk() ([]PortEntity, bool) {
	if o == nil || IsNil(o.InputPorts) {
		return nil, false
	}
	return o.InputPorts, true
}

// HasInputPorts returns a boolean if a field has been set.
func (o *FlowDTO) HasInputPorts() bool {
	if o != nil && !IsNil(o.InputPorts) {
		return true
	}

	return false
}

// SetInputPorts gets a reference to the given []PortEntity and assigns it to the InputPorts field.
func (o *FlowDTO) SetInputPorts(v []PortEntity) {
	o.InputPorts = v
}

// GetOutputPorts returns the OutputPorts field value if set, zero value otherwise.
func (o *FlowDTO) GetOutputPorts() []PortEntity {
	if o == nil || IsNil(o.OutputPorts) {
		var ret []PortEntity
		return ret
	}
	return o.OutputPorts
}

// GetOutputPortsOk returns a tuple with the OutputPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetOutputPortsOk() ([]PortEntity, bool) {
	if o == nil || IsNil(o.OutputPorts) {
		return nil, false
	}
	return o.OutputPorts, true
}

// HasOutputPorts returns a boolean if a field has been set.
func (o *FlowDTO) HasOutputPorts() bool {
	if o != nil && !IsNil(o.OutputPorts) {
		return true
	}

	return false
}

// SetOutputPorts gets a reference to the given []PortEntity and assigns it to the OutputPorts field.
func (o *FlowDTO) SetOutputPorts(v []PortEntity) {
	o.OutputPorts = v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *FlowDTO) GetConnections() []ConnectionEntity {
	if o == nil || IsNil(o.Connections) {
		var ret []ConnectionEntity
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetConnectionsOk() ([]ConnectionEntity, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *FlowDTO) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []ConnectionEntity and assigns it to the Connections field.
func (o *FlowDTO) SetConnections(v []ConnectionEntity) {
	o.Connections = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *FlowDTO) GetLabels() []LabelEntity {
	if o == nil || IsNil(o.Labels) {
		var ret []LabelEntity
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetLabelsOk() ([]LabelEntity, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *FlowDTO) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelEntity and assigns it to the Labels field.
func (o *FlowDTO) SetLabels(v []LabelEntity) {
	o.Labels = v
}

// GetFunnels returns the Funnels field value if set, zero value otherwise.
func (o *FlowDTO) GetFunnels() []FunnelEntity {
	if o == nil || IsNil(o.Funnels) {
		var ret []FunnelEntity
		return ret
	}
	return o.Funnels
}

// GetFunnelsOk returns a tuple with the Funnels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowDTO) GetFunnelsOk() ([]FunnelEntity, bool) {
	if o == nil || IsNil(o.Funnels) {
		return nil, false
	}
	return o.Funnels, true
}

// HasFunnels returns a boolean if a field has been set.
func (o *FlowDTO) HasFunnels() bool {
	if o != nil && !IsNil(o.Funnels) {
		return true
	}

	return false
}

// SetFunnels gets a reference to the given []FunnelEntity and assigns it to the Funnels field.
func (o *FlowDTO) SetFunnels(v []FunnelEntity) {
	o.Funnels = v
}

func (o FlowDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProcessGroups) {
		toSerialize["processGroups"] = o.ProcessGroups
	}
	if !IsNil(o.RemoteProcessGroups) {
		toSerialize["remoteProcessGroups"] = o.RemoteProcessGroups
	}
	if !IsNil(o.Processors) {
		toSerialize["processors"] = o.Processors
	}
	if !IsNil(o.InputPorts) {
		toSerialize["inputPorts"] = o.InputPorts
	}
	if !IsNil(o.OutputPorts) {
		toSerialize["outputPorts"] = o.OutputPorts
	}
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Funnels) {
		toSerialize["funnels"] = o.Funnels
	}
	return toSerialize, nil
}

type NullableFlowDTO struct {
	value *FlowDTO
	isSet bool
}

func (v NullableFlowDTO) Get() *FlowDTO {
	return v.value
}

func (v *NullableFlowDTO) Set(val *FlowDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDTO(val *FlowDTO) *NullableFlowDTO {
	return &NullableFlowDTO{value: val, isSet: true}
}

func (v NullableFlowDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

