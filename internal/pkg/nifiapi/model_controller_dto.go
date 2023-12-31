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

// checks if the ControllerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllerDTO{}

// ControllerDTO struct for ControllerDTO
type ControllerDTO struct {
	// The id of the NiFi.
	Id *string `json:"id,omitempty"`
	// The name of the NiFi.
	Name *string `json:"name,omitempty"`
	// The comments for the NiFi.
	Comments *string `json:"comments,omitempty"`
	// The number of running components in the NiFi.
	RunningCount *int32 `json:"runningCount,omitempty"`
	// The number of stopped components in the NiFi.
	StoppedCount *int32 `json:"stoppedCount,omitempty"`
	// The number of invalid components in the NiFi.
	InvalidCount *int32 `json:"invalidCount,omitempty"`
	// The number of disabled components in the NiFi.
	DisabledCount *int32 `json:"disabledCount,omitempty"`
	// The number of active remote ports contained in the NiFi.
	ActiveRemotePortCount *int32 `json:"activeRemotePortCount,omitempty"`
	// The number of inactive remote ports contained in the NiFi.
	InactiveRemotePortCount *int32 `json:"inactiveRemotePortCount,omitempty"`
	// The number of input ports contained in the NiFi.
	InputPortCount *int32 `json:"inputPortCount,omitempty"`
	// The number of output ports in the NiFi.
	OutputPortCount *int32 `json:"outputPortCount,omitempty"`
	// The Socket Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	RemoteSiteListeningPort *int32 `json:"remoteSiteListeningPort,omitempty"`
	// The HTTP(S) Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	RemoteSiteHttpListeningPort *int32 `json:"remoteSiteHttpListeningPort,omitempty"`
	// Indicates whether or not Site-to-Site communications with this instance is secure (2-way authentication).
	SiteToSiteSecure *bool `json:"siteToSiteSecure,omitempty"`
	// If clustered, the id of the Cluster Manager, otherwise the id of the NiFi.
	InstanceId *string `json:"instanceId,omitempty"`
	// The input ports available to send data to for the NiFi.
	InputPorts []PortDTO `json:"inputPorts,omitempty"`
	// The output ports available to received data from the NiFi.
	OutputPorts []PortDTO `json:"outputPorts,omitempty"`
}

// NewControllerDTO instantiates a new ControllerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerDTO() *ControllerDTO {
	this := ControllerDTO{}
	return &this
}

// NewControllerDTOWithDefaults instantiates a new ControllerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerDTOWithDefaults() *ControllerDTO {
	this := ControllerDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ControllerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ControllerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ControllerDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllerDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllerDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllerDTO) SetName(v string) {
	o.Name = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ControllerDTO) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ControllerDTO) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ControllerDTO) SetComments(v string) {
	o.Comments = &v
}

// GetRunningCount returns the RunningCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetRunningCount() int32 {
	if o == nil || IsNil(o.RunningCount) {
		var ret int32
		return ret
	}
	return *o.RunningCount
}

// GetRunningCountOk returns a tuple with the RunningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetRunningCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RunningCount) {
		return nil, false
	}
	return o.RunningCount, true
}

// HasRunningCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasRunningCount() bool {
	if o != nil && !IsNil(o.RunningCount) {
		return true
	}

	return false
}

// SetRunningCount gets a reference to the given int32 and assigns it to the RunningCount field.
func (o *ControllerDTO) SetRunningCount(v int32) {
	o.RunningCount = &v
}

// GetStoppedCount returns the StoppedCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetStoppedCount() int32 {
	if o == nil || IsNil(o.StoppedCount) {
		var ret int32
		return ret
	}
	return *o.StoppedCount
}

// GetStoppedCountOk returns a tuple with the StoppedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetStoppedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.StoppedCount) {
		return nil, false
	}
	return o.StoppedCount, true
}

// HasStoppedCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasStoppedCount() bool {
	if o != nil && !IsNil(o.StoppedCount) {
		return true
	}

	return false
}

// SetStoppedCount gets a reference to the given int32 and assigns it to the StoppedCount field.
func (o *ControllerDTO) SetStoppedCount(v int32) {
	o.StoppedCount = &v
}

// GetInvalidCount returns the InvalidCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetInvalidCount() int32 {
	if o == nil || IsNil(o.InvalidCount) {
		var ret int32
		return ret
	}
	return *o.InvalidCount
}

// GetInvalidCountOk returns a tuple with the InvalidCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetInvalidCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InvalidCount) {
		return nil, false
	}
	return o.InvalidCount, true
}

// HasInvalidCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasInvalidCount() bool {
	if o != nil && !IsNil(o.InvalidCount) {
		return true
	}

	return false
}

// SetInvalidCount gets a reference to the given int32 and assigns it to the InvalidCount field.
func (o *ControllerDTO) SetInvalidCount(v int32) {
	o.InvalidCount = &v
}

// GetDisabledCount returns the DisabledCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetDisabledCount() int32 {
	if o == nil || IsNil(o.DisabledCount) {
		var ret int32
		return ret
	}
	return *o.DisabledCount
}

// GetDisabledCountOk returns a tuple with the DisabledCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetDisabledCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DisabledCount) {
		return nil, false
	}
	return o.DisabledCount, true
}

// HasDisabledCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasDisabledCount() bool {
	if o != nil && !IsNil(o.DisabledCount) {
		return true
	}

	return false
}

// SetDisabledCount gets a reference to the given int32 and assigns it to the DisabledCount field.
func (o *ControllerDTO) SetDisabledCount(v int32) {
	o.DisabledCount = &v
}

// GetActiveRemotePortCount returns the ActiveRemotePortCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetActiveRemotePortCount() int32 {
	if o == nil || IsNil(o.ActiveRemotePortCount) {
		var ret int32
		return ret
	}
	return *o.ActiveRemotePortCount
}

// GetActiveRemotePortCountOk returns a tuple with the ActiveRemotePortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetActiveRemotePortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveRemotePortCount) {
		return nil, false
	}
	return o.ActiveRemotePortCount, true
}

// HasActiveRemotePortCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasActiveRemotePortCount() bool {
	if o != nil && !IsNil(o.ActiveRemotePortCount) {
		return true
	}

	return false
}

// SetActiveRemotePortCount gets a reference to the given int32 and assigns it to the ActiveRemotePortCount field.
func (o *ControllerDTO) SetActiveRemotePortCount(v int32) {
	o.ActiveRemotePortCount = &v
}

// GetInactiveRemotePortCount returns the InactiveRemotePortCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetInactiveRemotePortCount() int32 {
	if o == nil || IsNil(o.InactiveRemotePortCount) {
		var ret int32
		return ret
	}
	return *o.InactiveRemotePortCount
}

// GetInactiveRemotePortCountOk returns a tuple with the InactiveRemotePortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetInactiveRemotePortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InactiveRemotePortCount) {
		return nil, false
	}
	return o.InactiveRemotePortCount, true
}

// HasInactiveRemotePortCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasInactiveRemotePortCount() bool {
	if o != nil && !IsNil(o.InactiveRemotePortCount) {
		return true
	}

	return false
}

// SetInactiveRemotePortCount gets a reference to the given int32 and assigns it to the InactiveRemotePortCount field.
func (o *ControllerDTO) SetInactiveRemotePortCount(v int32) {
	o.InactiveRemotePortCount = &v
}

// GetInputPortCount returns the InputPortCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetInputPortCount() int32 {
	if o == nil || IsNil(o.InputPortCount) {
		var ret int32
		return ret
	}
	return *o.InputPortCount
}

// GetInputPortCountOk returns a tuple with the InputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetInputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InputPortCount) {
		return nil, false
	}
	return o.InputPortCount, true
}

// HasInputPortCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasInputPortCount() bool {
	if o != nil && !IsNil(o.InputPortCount) {
		return true
	}

	return false
}

// SetInputPortCount gets a reference to the given int32 and assigns it to the InputPortCount field.
func (o *ControllerDTO) SetInputPortCount(v int32) {
	o.InputPortCount = &v
}

// GetOutputPortCount returns the OutputPortCount field value if set, zero value otherwise.
func (o *ControllerDTO) GetOutputPortCount() int32 {
	if o == nil || IsNil(o.OutputPortCount) {
		var ret int32
		return ret
	}
	return *o.OutputPortCount
}

// GetOutputPortCountOk returns a tuple with the OutputPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetOutputPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputPortCount) {
		return nil, false
	}
	return o.OutputPortCount, true
}

// HasOutputPortCount returns a boolean if a field has been set.
func (o *ControllerDTO) HasOutputPortCount() bool {
	if o != nil && !IsNil(o.OutputPortCount) {
		return true
	}

	return false
}

// SetOutputPortCount gets a reference to the given int32 and assigns it to the OutputPortCount field.
func (o *ControllerDTO) SetOutputPortCount(v int32) {
	o.OutputPortCount = &v
}

// GetRemoteSiteListeningPort returns the RemoteSiteListeningPort field value if set, zero value otherwise.
func (o *ControllerDTO) GetRemoteSiteListeningPort() int32 {
	if o == nil || IsNil(o.RemoteSiteListeningPort) {
		var ret int32
		return ret
	}
	return *o.RemoteSiteListeningPort
}

// GetRemoteSiteListeningPortOk returns a tuple with the RemoteSiteListeningPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetRemoteSiteListeningPortOk() (*int32, bool) {
	if o == nil || IsNil(o.RemoteSiteListeningPort) {
		return nil, false
	}
	return o.RemoteSiteListeningPort, true
}

// HasRemoteSiteListeningPort returns a boolean if a field has been set.
func (o *ControllerDTO) HasRemoteSiteListeningPort() bool {
	if o != nil && !IsNil(o.RemoteSiteListeningPort) {
		return true
	}

	return false
}

// SetRemoteSiteListeningPort gets a reference to the given int32 and assigns it to the RemoteSiteListeningPort field.
func (o *ControllerDTO) SetRemoteSiteListeningPort(v int32) {
	o.RemoteSiteListeningPort = &v
}

// GetRemoteSiteHttpListeningPort returns the RemoteSiteHttpListeningPort field value if set, zero value otherwise.
func (o *ControllerDTO) GetRemoteSiteHttpListeningPort() int32 {
	if o == nil || IsNil(o.RemoteSiteHttpListeningPort) {
		var ret int32
		return ret
	}
	return *o.RemoteSiteHttpListeningPort
}

// GetRemoteSiteHttpListeningPortOk returns a tuple with the RemoteSiteHttpListeningPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetRemoteSiteHttpListeningPortOk() (*int32, bool) {
	if o == nil || IsNil(o.RemoteSiteHttpListeningPort) {
		return nil, false
	}
	return o.RemoteSiteHttpListeningPort, true
}

// HasRemoteSiteHttpListeningPort returns a boolean if a field has been set.
func (o *ControllerDTO) HasRemoteSiteHttpListeningPort() bool {
	if o != nil && !IsNil(o.RemoteSiteHttpListeningPort) {
		return true
	}

	return false
}

// SetRemoteSiteHttpListeningPort gets a reference to the given int32 and assigns it to the RemoteSiteHttpListeningPort field.
func (o *ControllerDTO) SetRemoteSiteHttpListeningPort(v int32) {
	o.RemoteSiteHttpListeningPort = &v
}

// GetSiteToSiteSecure returns the SiteToSiteSecure field value if set, zero value otherwise.
func (o *ControllerDTO) GetSiteToSiteSecure() bool {
	if o == nil || IsNil(o.SiteToSiteSecure) {
		var ret bool
		return ret
	}
	return *o.SiteToSiteSecure
}

// GetSiteToSiteSecureOk returns a tuple with the SiteToSiteSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetSiteToSiteSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.SiteToSiteSecure) {
		return nil, false
	}
	return o.SiteToSiteSecure, true
}

// HasSiteToSiteSecure returns a boolean if a field has been set.
func (o *ControllerDTO) HasSiteToSiteSecure() bool {
	if o != nil && !IsNil(o.SiteToSiteSecure) {
		return true
	}

	return false
}

// SetSiteToSiteSecure gets a reference to the given bool and assigns it to the SiteToSiteSecure field.
func (o *ControllerDTO) SetSiteToSiteSecure(v bool) {
	o.SiteToSiteSecure = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *ControllerDTO) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ControllerDTO) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *ControllerDTO) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetInputPorts returns the InputPorts field value if set, zero value otherwise.
func (o *ControllerDTO) GetInputPorts() []PortDTO {
	if o == nil || IsNil(o.InputPorts) {
		var ret []PortDTO
		return ret
	}
	return o.InputPorts
}

// GetInputPortsOk returns a tuple with the InputPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetInputPortsOk() ([]PortDTO, bool) {
	if o == nil || IsNil(o.InputPorts) {
		return nil, false
	}
	return o.InputPorts, true
}

// HasInputPorts returns a boolean if a field has been set.
func (o *ControllerDTO) HasInputPorts() bool {
	if o != nil && !IsNil(o.InputPorts) {
		return true
	}

	return false
}

// SetInputPorts gets a reference to the given []PortDTO and assigns it to the InputPorts field.
func (o *ControllerDTO) SetInputPorts(v []PortDTO) {
	o.InputPorts = v
}

// GetOutputPorts returns the OutputPorts field value if set, zero value otherwise.
func (o *ControllerDTO) GetOutputPorts() []PortDTO {
	if o == nil || IsNil(o.OutputPorts) {
		var ret []PortDTO
		return ret
	}
	return o.OutputPorts
}

// GetOutputPortsOk returns a tuple with the OutputPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerDTO) GetOutputPortsOk() ([]PortDTO, bool) {
	if o == nil || IsNil(o.OutputPorts) {
		return nil, false
	}
	return o.OutputPorts, true
}

// HasOutputPorts returns a boolean if a field has been set.
func (o *ControllerDTO) HasOutputPorts() bool {
	if o != nil && !IsNil(o.OutputPorts) {
		return true
	}

	return false
}

// SetOutputPorts gets a reference to the given []PortDTO and assigns it to the OutputPorts field.
func (o *ControllerDTO) SetOutputPorts(v []PortDTO) {
	o.OutputPorts = v
}

func (o ControllerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
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
	if !IsNil(o.InputPortCount) {
		toSerialize["inputPortCount"] = o.InputPortCount
	}
	if !IsNil(o.OutputPortCount) {
		toSerialize["outputPortCount"] = o.OutputPortCount
	}
	if !IsNil(o.RemoteSiteListeningPort) {
		toSerialize["remoteSiteListeningPort"] = o.RemoteSiteListeningPort
	}
	if !IsNil(o.RemoteSiteHttpListeningPort) {
		toSerialize["remoteSiteHttpListeningPort"] = o.RemoteSiteHttpListeningPort
	}
	if !IsNil(o.SiteToSiteSecure) {
		toSerialize["siteToSiteSecure"] = o.SiteToSiteSecure
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.InputPorts) {
		toSerialize["inputPorts"] = o.InputPorts
	}
	if !IsNil(o.OutputPorts) {
		toSerialize["outputPorts"] = o.OutputPorts
	}
	return toSerialize, nil
}

type NullableControllerDTO struct {
	value *ControllerDTO
	isSet bool
}

func (v NullableControllerDTO) Get() *ControllerDTO {
	return v.value
}

func (v *NullableControllerDTO) Set(val *ControllerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerDTO(val *ControllerDTO) *NullableControllerDTO {
	return &NullableControllerDTO{value: val, isSet: true}
}

func (v NullableControllerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


