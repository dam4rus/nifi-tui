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

// checks if the PropertyDependencyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyDependencyDTO{}

// PropertyDependencyDTO struct for PropertyDependencyDTO
type PropertyDependencyDTO struct {
	// The name of the property that is being depended upon
	PropertyName *string `json:"propertyName,omitempty"`
	// The values for the property that satisfies the dependency, or null if the dependency is satisfied by the presence of any value for the associated property name
	DependentValues []string `json:"dependentValues,omitempty"`
}

// NewPropertyDependencyDTO instantiates a new PropertyDependencyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDependencyDTO() *PropertyDependencyDTO {
	this := PropertyDependencyDTO{}
	return &this
}

// NewPropertyDependencyDTOWithDefaults instantiates a new PropertyDependencyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDependencyDTOWithDefaults() *PropertyDependencyDTO {
	this := PropertyDependencyDTO{}
	return &this
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *PropertyDependencyDTO) GetPropertyName() string {
	if o == nil || IsNil(o.PropertyName) {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDependencyDTO) GetPropertyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyName) {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *PropertyDependencyDTO) HasPropertyName() bool {
	if o != nil && !IsNil(o.PropertyName) {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *PropertyDependencyDTO) SetPropertyName(v string) {
	o.PropertyName = &v
}

// GetDependentValues returns the DependentValues field value if set, zero value otherwise.
func (o *PropertyDependencyDTO) GetDependentValues() []string {
	if o == nil || IsNil(o.DependentValues) {
		var ret []string
		return ret
	}
	return o.DependentValues
}

// GetDependentValuesOk returns a tuple with the DependentValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDependencyDTO) GetDependentValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.DependentValues) {
		return nil, false
	}
	return o.DependentValues, true
}

// HasDependentValues returns a boolean if a field has been set.
func (o *PropertyDependencyDTO) HasDependentValues() bool {
	if o != nil && !IsNil(o.DependentValues) {
		return true
	}

	return false
}

// SetDependentValues gets a reference to the given []string and assigns it to the DependentValues field.
func (o *PropertyDependencyDTO) SetDependentValues(v []string) {
	o.DependentValues = v
}

func (o PropertyDependencyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyDependencyDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyName) {
		toSerialize["propertyName"] = o.PropertyName
	}
	if !IsNil(o.DependentValues) {
		toSerialize["dependentValues"] = o.DependentValues
	}
	return toSerialize, nil
}

type NullablePropertyDependencyDTO struct {
	value *PropertyDependencyDTO
	isSet bool
}

func (v NullablePropertyDependencyDTO) Get() *PropertyDependencyDTO {
	return v.value
}

func (v *NullablePropertyDependencyDTO) Set(val *PropertyDependencyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDependencyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDependencyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDependencyDTO(val *PropertyDependencyDTO) *NullablePropertyDependencyDTO {
	return &NullablePropertyDependencyDTO{value: val, isSet: true}
}

func (v NullablePropertyDependencyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDependencyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


