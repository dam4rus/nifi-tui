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

// checks if the ScheduleComponentsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleComponentsEntity{}

// ScheduleComponentsEntity struct for ScheduleComponentsEntity
type ScheduleComponentsEntity struct {
	// The id of the ProcessGroup
	Id *string `json:"id,omitempty"`
	// The desired state of the descendant components
	State *string `json:"state,omitempty"`
	// Optional components to schedule. If not specified, all authorized descendant components will be used.
	Components *map[string]RevisionDTO `json:"components,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewScheduleComponentsEntity instantiates a new ScheduleComponentsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleComponentsEntity() *ScheduleComponentsEntity {
	this := ScheduleComponentsEntity{}
	return &this
}

// NewScheduleComponentsEntityWithDefaults instantiates a new ScheduleComponentsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleComponentsEntityWithDefaults() *ScheduleComponentsEntity {
	this := ScheduleComponentsEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduleComponentsEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleComponentsEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduleComponentsEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduleComponentsEntity) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ScheduleComponentsEntity) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleComponentsEntity) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ScheduleComponentsEntity) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ScheduleComponentsEntity) SetState(v string) {
	o.State = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ScheduleComponentsEntity) GetComponents() map[string]RevisionDTO {
	if o == nil || IsNil(o.Components) {
		var ret map[string]RevisionDTO
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleComponentsEntity) GetComponentsOk() (*map[string]RevisionDTO, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ScheduleComponentsEntity) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given map[string]RevisionDTO and assigns it to the Components field.
func (o *ScheduleComponentsEntity) SetComponents(v map[string]RevisionDTO) {
	o.Components = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ScheduleComponentsEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleComponentsEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ScheduleComponentsEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ScheduleComponentsEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o ScheduleComponentsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleComponentsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return toSerialize, nil
}

type NullableScheduleComponentsEntity struct {
	value *ScheduleComponentsEntity
	isSet bool
}

func (v NullableScheduleComponentsEntity) Get() *ScheduleComponentsEntity {
	return v.value
}

func (v *NullableScheduleComponentsEntity) Set(val *ScheduleComponentsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleComponentsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleComponentsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleComponentsEntity(val *ScheduleComponentsEntity) *NullableScheduleComponentsEntity {
	return &NullableScheduleComponentsEntity{value: val, isSet: true}
}

func (v NullableScheduleComponentsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleComponentsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


