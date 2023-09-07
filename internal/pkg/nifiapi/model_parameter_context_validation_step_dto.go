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

// checks if the ParameterContextValidationStepDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterContextValidationStepDTO{}

// ParameterContextValidationStepDTO struct for ParameterContextValidationStepDTO
type ParameterContextValidationStepDTO struct {
	// Explanation of what happens in this step
	Description *string `json:"description,omitempty"`
	// Whether or not this step has completed
	Complete *bool `json:"complete,omitempty"`
	// An explanation of why this step failed, or null if this step did not fail
	FailureReason *string `json:"failureReason,omitempty"`
}

// NewParameterContextValidationStepDTO instantiates a new ParameterContextValidationStepDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterContextValidationStepDTO() *ParameterContextValidationStepDTO {
	this := ParameterContextValidationStepDTO{}
	return &this
}

// NewParameterContextValidationStepDTOWithDefaults instantiates a new ParameterContextValidationStepDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterContextValidationStepDTOWithDefaults() *ParameterContextValidationStepDTO {
	this := ParameterContextValidationStepDTO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterContextValidationStepDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextValidationStepDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterContextValidationStepDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterContextValidationStepDTO) SetDescription(v string) {
	o.Description = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *ParameterContextValidationStepDTO) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextValidationStepDTO) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *ParameterContextValidationStepDTO) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *ParameterContextValidationStepDTO) SetComplete(v bool) {
	o.Complete = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *ParameterContextValidationStepDTO) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterContextValidationStepDTO) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *ParameterContextValidationStepDTO) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *ParameterContextValidationStepDTO) SetFailureReason(v string) {
	o.FailureReason = &v
}

func (o ParameterContextValidationStepDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterContextValidationStepDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	return toSerialize, nil
}

type NullableParameterContextValidationStepDTO struct {
	value *ParameterContextValidationStepDTO
	isSet bool
}

func (v NullableParameterContextValidationStepDTO) Get() *ParameterContextValidationStepDTO {
	return v.value
}

func (v *NullableParameterContextValidationStepDTO) Set(val *ParameterContextValidationStepDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterContextValidationStepDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterContextValidationStepDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterContextValidationStepDTO(val *ParameterContextValidationStepDTO) *NullableParameterContextValidationStepDTO {
	return &NullableParameterContextValidationStepDTO{value: val, isSet: true}
}

func (v NullableParameterContextValidationStepDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterContextValidationStepDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


