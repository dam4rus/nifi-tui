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

// checks if the ConfigVerificationResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigVerificationResultDTO{}

// ConfigVerificationResultDTO struct for ConfigVerificationResultDTO
type ConfigVerificationResultDTO struct {
	// The outcome of the verification
	Outcome *string `json:"outcome,omitempty"`
	// The name of the verification step
	VerificationStepName *string `json:"verificationStepName,omitempty"`
	// An explanation of why the step was or was not successful
	Explanation *string `json:"explanation,omitempty"`
}

// NewConfigVerificationResultDTO instantiates a new ConfigVerificationResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigVerificationResultDTO() *ConfigVerificationResultDTO {
	this := ConfigVerificationResultDTO{}
	return &this
}

// NewConfigVerificationResultDTOWithDefaults instantiates a new ConfigVerificationResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigVerificationResultDTOWithDefaults() *ConfigVerificationResultDTO {
	this := ConfigVerificationResultDTO{}
	return &this
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *ConfigVerificationResultDTO) GetOutcome() string {
	if o == nil || IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigVerificationResultDTO) GetOutcomeOk() (*string, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *ConfigVerificationResultDTO) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *ConfigVerificationResultDTO) SetOutcome(v string) {
	o.Outcome = &v
}

// GetVerificationStepName returns the VerificationStepName field value if set, zero value otherwise.
func (o *ConfigVerificationResultDTO) GetVerificationStepName() string {
	if o == nil || IsNil(o.VerificationStepName) {
		var ret string
		return ret
	}
	return *o.VerificationStepName
}

// GetVerificationStepNameOk returns a tuple with the VerificationStepName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigVerificationResultDTO) GetVerificationStepNameOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationStepName) {
		return nil, false
	}
	return o.VerificationStepName, true
}

// HasVerificationStepName returns a boolean if a field has been set.
func (o *ConfigVerificationResultDTO) HasVerificationStepName() bool {
	if o != nil && !IsNil(o.VerificationStepName) {
		return true
	}

	return false
}

// SetVerificationStepName gets a reference to the given string and assigns it to the VerificationStepName field.
func (o *ConfigVerificationResultDTO) SetVerificationStepName(v string) {
	o.VerificationStepName = &v
}

// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *ConfigVerificationResultDTO) GetExplanation() string {
	if o == nil || IsNil(o.Explanation) {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigVerificationResultDTO) GetExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.Explanation) {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *ConfigVerificationResultDTO) HasExplanation() bool {
	if o != nil && !IsNil(o.Explanation) {
		return true
	}

	return false
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *ConfigVerificationResultDTO) SetExplanation(v string) {
	o.Explanation = &v
}

func (o ConfigVerificationResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigVerificationResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.VerificationStepName) {
		toSerialize["verificationStepName"] = o.VerificationStepName
	}
	if !IsNil(o.Explanation) {
		toSerialize["explanation"] = o.Explanation
	}
	return toSerialize, nil
}

type NullableConfigVerificationResultDTO struct {
	value *ConfigVerificationResultDTO
	isSet bool
}

func (v NullableConfigVerificationResultDTO) Get() *ConfigVerificationResultDTO {
	return v.value
}

func (v *NullableConfigVerificationResultDTO) Set(val *ConfigVerificationResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigVerificationResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigVerificationResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigVerificationResultDTO(val *ConfigVerificationResultDTO) *NullableConfigVerificationResultDTO {
	return &NullableConfigVerificationResultDTO{value: val, isSet: true}
}

func (v NullableConfigVerificationResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigVerificationResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

