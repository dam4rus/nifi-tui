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

// checks if the ActionEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionEntity{}

// ActionEntity struct for ActionEntity
type ActionEntity struct {
	Id *int32 `json:"id,omitempty"`
	// The timestamp of the action.
	Timestamp *string `json:"timestamp,omitempty"`
	SourceId *string `json:"sourceId,omitempty"`
	// Indicates whether the user can read a given resource.
	CanRead *bool `json:"canRead,omitempty"`
	Action *ActionDTO `json:"action,omitempty"`
}

// NewActionEntity instantiates a new ActionEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionEntity() *ActionEntity {
	this := ActionEntity{}
	return &this
}

// NewActionEntityWithDefaults instantiates a new ActionEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionEntityWithDefaults() *ActionEntity {
	this := ActionEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActionEntity) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionEntity) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActionEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ActionEntity) SetId(v int32) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ActionEntity) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionEntity) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ActionEntity) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ActionEntity) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ActionEntity) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionEntity) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ActionEntity) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ActionEntity) SetSourceId(v string) {
	o.SourceId = &v
}

// GetCanRead returns the CanRead field value if set, zero value otherwise.
func (o *ActionEntity) GetCanRead() bool {
	if o == nil || IsNil(o.CanRead) {
		var ret bool
		return ret
	}
	return *o.CanRead
}

// GetCanReadOk returns a tuple with the CanRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionEntity) GetCanReadOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRead) {
		return nil, false
	}
	return o.CanRead, true
}

// HasCanRead returns a boolean if a field has been set.
func (o *ActionEntity) HasCanRead() bool {
	if o != nil && !IsNil(o.CanRead) {
		return true
	}

	return false
}

// SetCanRead gets a reference to the given bool and assigns it to the CanRead field.
func (o *ActionEntity) SetCanRead(v bool) {
	o.CanRead = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ActionEntity) GetAction() ActionDTO {
	if o == nil || IsNil(o.Action) {
		var ret ActionDTO
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionEntity) GetActionOk() (*ActionDTO, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ActionEntity) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given ActionDTO and assigns it to the Action field.
func (o *ActionEntity) SetAction(v ActionDTO) {
	o.Action = &v
}

func (o ActionEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.CanRead) {
		toSerialize["canRead"] = o.CanRead
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableActionEntity struct {
	value *ActionEntity
	isSet bool
}

func (v NullableActionEntity) Get() *ActionEntity {
	return v.value
}

func (v *NullableActionEntity) Set(val *ActionEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableActionEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableActionEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionEntity(val *ActionEntity) *NullableActionEntity {
	return &NullableActionEntity{value: val, isSet: true}
}

func (v NullableActionEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


