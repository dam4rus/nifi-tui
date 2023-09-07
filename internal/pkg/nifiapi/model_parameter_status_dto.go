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

// checks if the ParameterStatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterStatusDTO{}

// ParameterStatusDTO struct for ParameterStatusDTO
type ParameterStatusDTO struct {
	Parameter *ParameterEntity `json:"parameter,omitempty"`
	// Indicates the status of the parameter, compared to the existing parameter context
	Status *string `json:"status,omitempty"`
}

// NewParameterStatusDTO instantiates a new ParameterStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterStatusDTO() *ParameterStatusDTO {
	this := ParameterStatusDTO{}
	return &this
}

// NewParameterStatusDTOWithDefaults instantiates a new ParameterStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterStatusDTOWithDefaults() *ParameterStatusDTO {
	this := ParameterStatusDTO{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ParameterStatusDTO) GetParameter() ParameterEntity {
	if o == nil || IsNil(o.Parameter) {
		var ret ParameterEntity
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterStatusDTO) GetParameterOk() (*ParameterEntity, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ParameterStatusDTO) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given ParameterEntity and assigns it to the Parameter field.
func (o *ParameterStatusDTO) SetParameter(v ParameterEntity) {
	o.Parameter = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ParameterStatusDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterStatusDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ParameterStatusDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ParameterStatusDTO) SetStatus(v string) {
	o.Status = &v
}

func (o ParameterStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterStatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableParameterStatusDTO struct {
	value *ParameterStatusDTO
	isSet bool
}

func (v NullableParameterStatusDTO) Get() *ParameterStatusDTO {
	return v.value
}

func (v *NullableParameterStatusDTO) Set(val *ParameterStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterStatusDTO(val *ParameterStatusDTO) *NullableParameterStatusDTO {
	return &NullableParameterStatusDTO{value: val, isSet: true}
}

func (v NullableParameterStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


