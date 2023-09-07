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

// checks if the StateMapDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateMapDTO{}

// StateMapDTO struct for StateMapDTO
type StateMapDTO struct {
	// The scope of this StateMap.
	Scope *string `json:"scope,omitempty"`
	// The total number of state entries. When the state map is lengthy, only of portion of the entries are returned.
	TotalEntryCount *int32 `json:"totalEntryCount,omitempty"`
	// The state.
	State []StateEntryDTO `json:"state,omitempty"`
}

// NewStateMapDTO instantiates a new StateMapDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateMapDTO() *StateMapDTO {
	this := StateMapDTO{}
	return &this
}

// NewStateMapDTOWithDefaults instantiates a new StateMapDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateMapDTOWithDefaults() *StateMapDTO {
	this := StateMapDTO{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *StateMapDTO) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateMapDTO) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *StateMapDTO) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *StateMapDTO) SetScope(v string) {
	o.Scope = &v
}

// GetTotalEntryCount returns the TotalEntryCount field value if set, zero value otherwise.
func (o *StateMapDTO) GetTotalEntryCount() int32 {
	if o == nil || IsNil(o.TotalEntryCount) {
		var ret int32
		return ret
	}
	return *o.TotalEntryCount
}

// GetTotalEntryCountOk returns a tuple with the TotalEntryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateMapDTO) GetTotalEntryCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalEntryCount) {
		return nil, false
	}
	return o.TotalEntryCount, true
}

// HasTotalEntryCount returns a boolean if a field has been set.
func (o *StateMapDTO) HasTotalEntryCount() bool {
	if o != nil && !IsNil(o.TotalEntryCount) {
		return true
	}

	return false
}

// SetTotalEntryCount gets a reference to the given int32 and assigns it to the TotalEntryCount field.
func (o *StateMapDTO) SetTotalEntryCount(v int32) {
	o.TotalEntryCount = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StateMapDTO) GetState() []StateEntryDTO {
	if o == nil || IsNil(o.State) {
		var ret []StateEntryDTO
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateMapDTO) GetStateOk() ([]StateEntryDTO, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StateMapDTO) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given []StateEntryDTO and assigns it to the State field.
func (o *StateMapDTO) SetState(v []StateEntryDTO) {
	o.State = v
}

func (o StateMapDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateMapDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.TotalEntryCount) {
		toSerialize["totalEntryCount"] = o.TotalEntryCount
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableStateMapDTO struct {
	value *StateMapDTO
	isSet bool
}

func (v NullableStateMapDTO) Get() *StateMapDTO {
	return v.value
}

func (v *NullableStateMapDTO) Set(val *StateMapDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableStateMapDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableStateMapDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateMapDTO(val *StateMapDTO) *NullableStateMapDTO {
	return &NullableStateMapDTO{value: val, isSet: true}
}

func (v NullableStateMapDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateMapDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

