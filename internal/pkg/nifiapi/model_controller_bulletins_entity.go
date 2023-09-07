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

// checks if the ControllerBulletinsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllerBulletinsEntity{}

// ControllerBulletinsEntity struct for ControllerBulletinsEntity
type ControllerBulletinsEntity struct {
	// System level bulletins to be reported to the user.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Controller service bulletins to be reported to the user.
	ControllerServiceBulletins []BulletinEntity `json:"controllerServiceBulletins,omitempty"`
	// Reporting task bulletins to be reported to the user.
	ReportingTaskBulletins []BulletinEntity `json:"reportingTaskBulletins,omitempty"`
	// Parameter provider bulletins to be reported to the user.
	ParameterProviderBulletins []BulletinEntity `json:"parameterProviderBulletins,omitempty"`
	// Flow registry client bulletins to be reported to the user.
	FlowRegistryClientBulletins []BulletinEntity `json:"flowRegistryClientBulletins,omitempty"`
}

// NewControllerBulletinsEntity instantiates a new ControllerBulletinsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerBulletinsEntity() *ControllerBulletinsEntity {
	this := ControllerBulletinsEntity{}
	return &this
}

// NewControllerBulletinsEntityWithDefaults instantiates a new ControllerBulletinsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerBulletinsEntityWithDefaults() *ControllerBulletinsEntity {
	this := ControllerBulletinsEntity{}
	return &this
}

// GetBulletins returns the Bulletins field value if set, zero value otherwise.
func (o *ControllerBulletinsEntity) GetBulletins() []BulletinEntity {
	if o == nil || IsNil(o.Bulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.Bulletins
}

// GetBulletinsOk returns a tuple with the Bulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerBulletinsEntity) GetBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.Bulletins) {
		return nil, false
	}
	return o.Bulletins, true
}

// HasBulletins returns a boolean if a field has been set.
func (o *ControllerBulletinsEntity) HasBulletins() bool {
	if o != nil && !IsNil(o.Bulletins) {
		return true
	}

	return false
}

// SetBulletins gets a reference to the given []BulletinEntity and assigns it to the Bulletins field.
func (o *ControllerBulletinsEntity) SetBulletins(v []BulletinEntity) {
	o.Bulletins = v
}

// GetControllerServiceBulletins returns the ControllerServiceBulletins field value if set, zero value otherwise.
func (o *ControllerBulletinsEntity) GetControllerServiceBulletins() []BulletinEntity {
	if o == nil || IsNil(o.ControllerServiceBulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.ControllerServiceBulletins
}

// GetControllerServiceBulletinsOk returns a tuple with the ControllerServiceBulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerBulletinsEntity) GetControllerServiceBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.ControllerServiceBulletins) {
		return nil, false
	}
	return o.ControllerServiceBulletins, true
}

// HasControllerServiceBulletins returns a boolean if a field has been set.
func (o *ControllerBulletinsEntity) HasControllerServiceBulletins() bool {
	if o != nil && !IsNil(o.ControllerServiceBulletins) {
		return true
	}

	return false
}

// SetControllerServiceBulletins gets a reference to the given []BulletinEntity and assigns it to the ControllerServiceBulletins field.
func (o *ControllerBulletinsEntity) SetControllerServiceBulletins(v []BulletinEntity) {
	o.ControllerServiceBulletins = v
}

// GetReportingTaskBulletins returns the ReportingTaskBulletins field value if set, zero value otherwise.
func (o *ControllerBulletinsEntity) GetReportingTaskBulletins() []BulletinEntity {
	if o == nil || IsNil(o.ReportingTaskBulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.ReportingTaskBulletins
}

// GetReportingTaskBulletinsOk returns a tuple with the ReportingTaskBulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerBulletinsEntity) GetReportingTaskBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.ReportingTaskBulletins) {
		return nil, false
	}
	return o.ReportingTaskBulletins, true
}

// HasReportingTaskBulletins returns a boolean if a field has been set.
func (o *ControllerBulletinsEntity) HasReportingTaskBulletins() bool {
	if o != nil && !IsNil(o.ReportingTaskBulletins) {
		return true
	}

	return false
}

// SetReportingTaskBulletins gets a reference to the given []BulletinEntity and assigns it to the ReportingTaskBulletins field.
func (o *ControllerBulletinsEntity) SetReportingTaskBulletins(v []BulletinEntity) {
	o.ReportingTaskBulletins = v
}

// GetParameterProviderBulletins returns the ParameterProviderBulletins field value if set, zero value otherwise.
func (o *ControllerBulletinsEntity) GetParameterProviderBulletins() []BulletinEntity {
	if o == nil || IsNil(o.ParameterProviderBulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.ParameterProviderBulletins
}

// GetParameterProviderBulletinsOk returns a tuple with the ParameterProviderBulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerBulletinsEntity) GetParameterProviderBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.ParameterProviderBulletins) {
		return nil, false
	}
	return o.ParameterProviderBulletins, true
}

// HasParameterProviderBulletins returns a boolean if a field has been set.
func (o *ControllerBulletinsEntity) HasParameterProviderBulletins() bool {
	if o != nil && !IsNil(o.ParameterProviderBulletins) {
		return true
	}

	return false
}

// SetParameterProviderBulletins gets a reference to the given []BulletinEntity and assigns it to the ParameterProviderBulletins field.
func (o *ControllerBulletinsEntity) SetParameterProviderBulletins(v []BulletinEntity) {
	o.ParameterProviderBulletins = v
}

// GetFlowRegistryClientBulletins returns the FlowRegistryClientBulletins field value if set, zero value otherwise.
func (o *ControllerBulletinsEntity) GetFlowRegistryClientBulletins() []BulletinEntity {
	if o == nil || IsNil(o.FlowRegistryClientBulletins) {
		var ret []BulletinEntity
		return ret
	}
	return o.FlowRegistryClientBulletins
}

// GetFlowRegistryClientBulletinsOk returns a tuple with the FlowRegistryClientBulletins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerBulletinsEntity) GetFlowRegistryClientBulletinsOk() ([]BulletinEntity, bool) {
	if o == nil || IsNil(o.FlowRegistryClientBulletins) {
		return nil, false
	}
	return o.FlowRegistryClientBulletins, true
}

// HasFlowRegistryClientBulletins returns a boolean if a field has been set.
func (o *ControllerBulletinsEntity) HasFlowRegistryClientBulletins() bool {
	if o != nil && !IsNil(o.FlowRegistryClientBulletins) {
		return true
	}

	return false
}

// SetFlowRegistryClientBulletins gets a reference to the given []BulletinEntity and assigns it to the FlowRegistryClientBulletins field.
func (o *ControllerBulletinsEntity) SetFlowRegistryClientBulletins(v []BulletinEntity) {
	o.FlowRegistryClientBulletins = v
}

func (o ControllerBulletinsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllerBulletinsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bulletins) {
		toSerialize["bulletins"] = o.Bulletins
	}
	if !IsNil(o.ControllerServiceBulletins) {
		toSerialize["controllerServiceBulletins"] = o.ControllerServiceBulletins
	}
	if !IsNil(o.ReportingTaskBulletins) {
		toSerialize["reportingTaskBulletins"] = o.ReportingTaskBulletins
	}
	if !IsNil(o.ParameterProviderBulletins) {
		toSerialize["parameterProviderBulletins"] = o.ParameterProviderBulletins
	}
	if !IsNil(o.FlowRegistryClientBulletins) {
		toSerialize["flowRegistryClientBulletins"] = o.FlowRegistryClientBulletins
	}
	return toSerialize, nil
}

type NullableControllerBulletinsEntity struct {
	value *ControllerBulletinsEntity
	isSet bool
}

func (v NullableControllerBulletinsEntity) Get() *ControllerBulletinsEntity {
	return v.value
}

func (v *NullableControllerBulletinsEntity) Set(val *ControllerBulletinsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerBulletinsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerBulletinsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerBulletinsEntity(val *ControllerBulletinsEntity) *NullableControllerBulletinsEntity {
	return &NullableControllerBulletinsEntity{value: val, isSet: true}
}

func (v NullableControllerBulletinsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerBulletinsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

