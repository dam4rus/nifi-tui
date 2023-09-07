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

// checks if the FlowConfigurationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowConfigurationDTO{}

// FlowConfigurationDTO struct for FlowConfigurationDTO
type FlowConfigurationDTO struct {
	// Whether this NiFi supports a managed authorizer. Managed authorizers can visualize users, groups, and policies in the UI.
	SupportsManagedAuthorizer *bool `json:"supportsManagedAuthorizer,omitempty"`
	// Whether this NiFi supports a configurable authorizer.
	SupportsConfigurableAuthorizer *bool `json:"supportsConfigurableAuthorizer,omitempty"`
	// Whether this NiFi supports configurable users and groups.
	SupportsConfigurableUsersAndGroups *bool `json:"supportsConfigurableUsersAndGroups,omitempty"`
	// The interval in seconds between the automatic NiFi refresh requests.
	AutoRefreshIntervalSeconds *int64 `json:"autoRefreshIntervalSeconds,omitempty"`
	// The current time on the system.
	CurrentTime *string `json:"currentTime,omitempty"`
	// The time offset of the system.
	TimeOffset *int32 `json:"timeOffset,omitempty"`
	// The default back pressure object threshold.
	DefaultBackPressureObjectThreshold *int64 `json:"defaultBackPressureObjectThreshold,omitempty"`
	// The default back pressure data size threshold.
	DefaultBackPressureDataSizeThreshold *string `json:"defaultBackPressureDataSizeThreshold,omitempty"`
}

// NewFlowConfigurationDTO instantiates a new FlowConfigurationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowConfigurationDTO() *FlowConfigurationDTO {
	this := FlowConfigurationDTO{}
	return &this
}

// NewFlowConfigurationDTOWithDefaults instantiates a new FlowConfigurationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowConfigurationDTOWithDefaults() *FlowConfigurationDTO {
	this := FlowConfigurationDTO{}
	return &this
}

// GetSupportsManagedAuthorizer returns the SupportsManagedAuthorizer field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetSupportsManagedAuthorizer() bool {
	if o == nil || IsNil(o.SupportsManagedAuthorizer) {
		var ret bool
		return ret
	}
	return *o.SupportsManagedAuthorizer
}

// GetSupportsManagedAuthorizerOk returns a tuple with the SupportsManagedAuthorizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetSupportsManagedAuthorizerOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsManagedAuthorizer) {
		return nil, false
	}
	return o.SupportsManagedAuthorizer, true
}

// HasSupportsManagedAuthorizer returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasSupportsManagedAuthorizer() bool {
	if o != nil && !IsNil(o.SupportsManagedAuthorizer) {
		return true
	}

	return false
}

// SetSupportsManagedAuthorizer gets a reference to the given bool and assigns it to the SupportsManagedAuthorizer field.
func (o *FlowConfigurationDTO) SetSupportsManagedAuthorizer(v bool) {
	o.SupportsManagedAuthorizer = &v
}

// GetSupportsConfigurableAuthorizer returns the SupportsConfigurableAuthorizer field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetSupportsConfigurableAuthorizer() bool {
	if o == nil || IsNil(o.SupportsConfigurableAuthorizer) {
		var ret bool
		return ret
	}
	return *o.SupportsConfigurableAuthorizer
}

// GetSupportsConfigurableAuthorizerOk returns a tuple with the SupportsConfigurableAuthorizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetSupportsConfigurableAuthorizerOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsConfigurableAuthorizer) {
		return nil, false
	}
	return o.SupportsConfigurableAuthorizer, true
}

// HasSupportsConfigurableAuthorizer returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasSupportsConfigurableAuthorizer() bool {
	if o != nil && !IsNil(o.SupportsConfigurableAuthorizer) {
		return true
	}

	return false
}

// SetSupportsConfigurableAuthorizer gets a reference to the given bool and assigns it to the SupportsConfigurableAuthorizer field.
func (o *FlowConfigurationDTO) SetSupportsConfigurableAuthorizer(v bool) {
	o.SupportsConfigurableAuthorizer = &v
}

// GetSupportsConfigurableUsersAndGroups returns the SupportsConfigurableUsersAndGroups field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetSupportsConfigurableUsersAndGroups() bool {
	if o == nil || IsNil(o.SupportsConfigurableUsersAndGroups) {
		var ret bool
		return ret
	}
	return *o.SupportsConfigurableUsersAndGroups
}

// GetSupportsConfigurableUsersAndGroupsOk returns a tuple with the SupportsConfigurableUsersAndGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetSupportsConfigurableUsersAndGroupsOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsConfigurableUsersAndGroups) {
		return nil, false
	}
	return o.SupportsConfigurableUsersAndGroups, true
}

// HasSupportsConfigurableUsersAndGroups returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasSupportsConfigurableUsersAndGroups() bool {
	if o != nil && !IsNil(o.SupportsConfigurableUsersAndGroups) {
		return true
	}

	return false
}

// SetSupportsConfigurableUsersAndGroups gets a reference to the given bool and assigns it to the SupportsConfigurableUsersAndGroups field.
func (o *FlowConfigurationDTO) SetSupportsConfigurableUsersAndGroups(v bool) {
	o.SupportsConfigurableUsersAndGroups = &v
}

// GetAutoRefreshIntervalSeconds returns the AutoRefreshIntervalSeconds field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetAutoRefreshIntervalSeconds() int64 {
	if o == nil || IsNil(o.AutoRefreshIntervalSeconds) {
		var ret int64
		return ret
	}
	return *o.AutoRefreshIntervalSeconds
}

// GetAutoRefreshIntervalSecondsOk returns a tuple with the AutoRefreshIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetAutoRefreshIntervalSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.AutoRefreshIntervalSeconds) {
		return nil, false
	}
	return o.AutoRefreshIntervalSeconds, true
}

// HasAutoRefreshIntervalSeconds returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasAutoRefreshIntervalSeconds() bool {
	if o != nil && !IsNil(o.AutoRefreshIntervalSeconds) {
		return true
	}

	return false
}

// SetAutoRefreshIntervalSeconds gets a reference to the given int64 and assigns it to the AutoRefreshIntervalSeconds field.
func (o *FlowConfigurationDTO) SetAutoRefreshIntervalSeconds(v int64) {
	o.AutoRefreshIntervalSeconds = &v
}

// GetCurrentTime returns the CurrentTime field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetCurrentTime() string {
	if o == nil || IsNil(o.CurrentTime) {
		var ret string
		return ret
	}
	return *o.CurrentTime
}

// GetCurrentTimeOk returns a tuple with the CurrentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetCurrentTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentTime) {
		return nil, false
	}
	return o.CurrentTime, true
}

// HasCurrentTime returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasCurrentTime() bool {
	if o != nil && !IsNil(o.CurrentTime) {
		return true
	}

	return false
}

// SetCurrentTime gets a reference to the given string and assigns it to the CurrentTime field.
func (o *FlowConfigurationDTO) SetCurrentTime(v string) {
	o.CurrentTime = &v
}

// GetTimeOffset returns the TimeOffset field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetTimeOffset() int32 {
	if o == nil || IsNil(o.TimeOffset) {
		var ret int32
		return ret
	}
	return *o.TimeOffset
}

// GetTimeOffsetOk returns a tuple with the TimeOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetTimeOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeOffset) {
		return nil, false
	}
	return o.TimeOffset, true
}

// HasTimeOffset returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasTimeOffset() bool {
	if o != nil && !IsNil(o.TimeOffset) {
		return true
	}

	return false
}

// SetTimeOffset gets a reference to the given int32 and assigns it to the TimeOffset field.
func (o *FlowConfigurationDTO) SetTimeOffset(v int32) {
	o.TimeOffset = &v
}

// GetDefaultBackPressureObjectThreshold returns the DefaultBackPressureObjectThreshold field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetDefaultBackPressureObjectThreshold() int64 {
	if o == nil || IsNil(o.DefaultBackPressureObjectThreshold) {
		var ret int64
		return ret
	}
	return *o.DefaultBackPressureObjectThreshold
}

// GetDefaultBackPressureObjectThresholdOk returns a tuple with the DefaultBackPressureObjectThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetDefaultBackPressureObjectThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultBackPressureObjectThreshold) {
		return nil, false
	}
	return o.DefaultBackPressureObjectThreshold, true
}

// HasDefaultBackPressureObjectThreshold returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasDefaultBackPressureObjectThreshold() bool {
	if o != nil && !IsNil(o.DefaultBackPressureObjectThreshold) {
		return true
	}

	return false
}

// SetDefaultBackPressureObjectThreshold gets a reference to the given int64 and assigns it to the DefaultBackPressureObjectThreshold field.
func (o *FlowConfigurationDTO) SetDefaultBackPressureObjectThreshold(v int64) {
	o.DefaultBackPressureObjectThreshold = &v
}

// GetDefaultBackPressureDataSizeThreshold returns the DefaultBackPressureDataSizeThreshold field value if set, zero value otherwise.
func (o *FlowConfigurationDTO) GetDefaultBackPressureDataSizeThreshold() string {
	if o == nil || IsNil(o.DefaultBackPressureDataSizeThreshold) {
		var ret string
		return ret
	}
	return *o.DefaultBackPressureDataSizeThreshold
}

// GetDefaultBackPressureDataSizeThresholdOk returns a tuple with the DefaultBackPressureDataSizeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowConfigurationDTO) GetDefaultBackPressureDataSizeThresholdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBackPressureDataSizeThreshold) {
		return nil, false
	}
	return o.DefaultBackPressureDataSizeThreshold, true
}

// HasDefaultBackPressureDataSizeThreshold returns a boolean if a field has been set.
func (o *FlowConfigurationDTO) HasDefaultBackPressureDataSizeThreshold() bool {
	if o != nil && !IsNil(o.DefaultBackPressureDataSizeThreshold) {
		return true
	}

	return false
}

// SetDefaultBackPressureDataSizeThreshold gets a reference to the given string and assigns it to the DefaultBackPressureDataSizeThreshold field.
func (o *FlowConfigurationDTO) SetDefaultBackPressureDataSizeThreshold(v string) {
	o.DefaultBackPressureDataSizeThreshold = &v
}

func (o FlowConfigurationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowConfigurationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportsManagedAuthorizer) {
		toSerialize["supportsManagedAuthorizer"] = o.SupportsManagedAuthorizer
	}
	if !IsNil(o.SupportsConfigurableAuthorizer) {
		toSerialize["supportsConfigurableAuthorizer"] = o.SupportsConfigurableAuthorizer
	}
	if !IsNil(o.SupportsConfigurableUsersAndGroups) {
		toSerialize["supportsConfigurableUsersAndGroups"] = o.SupportsConfigurableUsersAndGroups
	}
	if !IsNil(o.AutoRefreshIntervalSeconds) {
		toSerialize["autoRefreshIntervalSeconds"] = o.AutoRefreshIntervalSeconds
	}
	if !IsNil(o.CurrentTime) {
		toSerialize["currentTime"] = o.CurrentTime
	}
	if !IsNil(o.TimeOffset) {
		toSerialize["timeOffset"] = o.TimeOffset
	}
	if !IsNil(o.DefaultBackPressureObjectThreshold) {
		toSerialize["defaultBackPressureObjectThreshold"] = o.DefaultBackPressureObjectThreshold
	}
	if !IsNil(o.DefaultBackPressureDataSizeThreshold) {
		toSerialize["defaultBackPressureDataSizeThreshold"] = o.DefaultBackPressureDataSizeThreshold
	}
	return toSerialize, nil
}

type NullableFlowConfigurationDTO struct {
	value *FlowConfigurationDTO
	isSet bool
}

func (v NullableFlowConfigurationDTO) Get() *FlowConfigurationDTO {
	return v.value
}

func (v *NullableFlowConfigurationDTO) Set(val *FlowConfigurationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowConfigurationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowConfigurationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowConfigurationDTO(val *FlowConfigurationDTO) *NullableFlowConfigurationDTO {
	return &NullableFlowConfigurationDTO{value: val, isSet: true}
}

func (v NullableFlowConfigurationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowConfigurationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


