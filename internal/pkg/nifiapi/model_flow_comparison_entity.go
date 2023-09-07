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

// checks if the FlowComparisonEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowComparisonEntity{}

// FlowComparisonEntity struct for FlowComparisonEntity
type FlowComparisonEntity struct {
	// The list of differences for each component in the flow that is not the same between the two flows
	ComponentDifferences []ComponentDifferenceDTO `json:"componentDifferences,omitempty"`
}

// NewFlowComparisonEntity instantiates a new FlowComparisonEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowComparisonEntity() *FlowComparisonEntity {
	this := FlowComparisonEntity{}
	return &this
}

// NewFlowComparisonEntityWithDefaults instantiates a new FlowComparisonEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowComparisonEntityWithDefaults() *FlowComparisonEntity {
	this := FlowComparisonEntity{}
	return &this
}

// GetComponentDifferences returns the ComponentDifferences field value if set, zero value otherwise.
func (o *FlowComparisonEntity) GetComponentDifferences() []ComponentDifferenceDTO {
	if o == nil || IsNil(o.ComponentDifferences) {
		var ret []ComponentDifferenceDTO
		return ret
	}
	return o.ComponentDifferences
}

// GetComponentDifferencesOk returns a tuple with the ComponentDifferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowComparisonEntity) GetComponentDifferencesOk() ([]ComponentDifferenceDTO, bool) {
	if o == nil || IsNil(o.ComponentDifferences) {
		return nil, false
	}
	return o.ComponentDifferences, true
}

// HasComponentDifferences returns a boolean if a field has been set.
func (o *FlowComparisonEntity) HasComponentDifferences() bool {
	if o != nil && !IsNil(o.ComponentDifferences) {
		return true
	}

	return false
}

// SetComponentDifferences gets a reference to the given []ComponentDifferenceDTO and assigns it to the ComponentDifferences field.
func (o *FlowComparisonEntity) SetComponentDifferences(v []ComponentDifferenceDTO) {
	o.ComponentDifferences = v
}

func (o FlowComparisonEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowComparisonEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentDifferences) {
		toSerialize["componentDifferences"] = o.ComponentDifferences
	}
	return toSerialize, nil
}

type NullableFlowComparisonEntity struct {
	value *FlowComparisonEntity
	isSet bool
}

func (v NullableFlowComparisonEntity) Get() *FlowComparisonEntity {
	return v.value
}

func (v *NullableFlowComparisonEntity) Set(val *FlowComparisonEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowComparisonEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowComparisonEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowComparisonEntity(val *FlowComparisonEntity) *NullableFlowComparisonEntity {
	return &NullableFlowComparisonEntity{value: val, isSet: true}
}

func (v NullableFlowComparisonEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowComparisonEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

