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

// checks if the ControllerServiceRunStatusEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllerServiceRunStatusEntity{}

// ControllerServiceRunStatusEntity struct for ControllerServiceRunStatusEntity
type ControllerServiceRunStatusEntity struct {
	Revision *RevisionDTO `json:"revision,omitempty"`
	// The run status of the ControllerService.
	State *string `json:"state,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
	// Indicates whether or not responses should only include fields necessary for rendering the NiFi User Interface. As such, when this value is set to true, some fields may be returned as null values, and the selected fields may change at any time without notice. As a result, this value should not be set to true by any client other than the UI.
	UiOnly *bool `json:"uiOnly,omitempty"`
}

// NewControllerServiceRunStatusEntity instantiates a new ControllerServiceRunStatusEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerServiceRunStatusEntity() *ControllerServiceRunStatusEntity {
	this := ControllerServiceRunStatusEntity{}
	return &this
}

// NewControllerServiceRunStatusEntityWithDefaults instantiates a new ControllerServiceRunStatusEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerServiceRunStatusEntityWithDefaults() *ControllerServiceRunStatusEntity {
	this := ControllerServiceRunStatusEntity{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ControllerServiceRunStatusEntity) GetRevision() RevisionDTO {
	if o == nil || IsNil(o.Revision) {
		var ret RevisionDTO
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceRunStatusEntity) GetRevisionOk() (*RevisionDTO, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ControllerServiceRunStatusEntity) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given RevisionDTO and assigns it to the Revision field.
func (o *ControllerServiceRunStatusEntity) SetRevision(v RevisionDTO) {
	o.Revision = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ControllerServiceRunStatusEntity) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceRunStatusEntity) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ControllerServiceRunStatusEntity) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ControllerServiceRunStatusEntity) SetState(v string) {
	o.State = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *ControllerServiceRunStatusEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceRunStatusEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *ControllerServiceRunStatusEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *ControllerServiceRunStatusEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

// GetUiOnly returns the UiOnly field value if set, zero value otherwise.
func (o *ControllerServiceRunStatusEntity) GetUiOnly() bool {
	if o == nil || IsNil(o.UiOnly) {
		var ret bool
		return ret
	}
	return *o.UiOnly
}

// GetUiOnlyOk returns a tuple with the UiOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceRunStatusEntity) GetUiOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.UiOnly) {
		return nil, false
	}
	return o.UiOnly, true
}

// HasUiOnly returns a boolean if a field has been set.
func (o *ControllerServiceRunStatusEntity) HasUiOnly() bool {
	if o != nil && !IsNil(o.UiOnly) {
		return true
	}

	return false
}

// SetUiOnly gets a reference to the given bool and assigns it to the UiOnly field.
func (o *ControllerServiceRunStatusEntity) SetUiOnly(v bool) {
	o.UiOnly = &v
}

func (o ControllerServiceRunStatusEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllerServiceRunStatusEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	if !IsNil(o.UiOnly) {
		toSerialize["uiOnly"] = o.UiOnly
	}
	return toSerialize, nil
}

type NullableControllerServiceRunStatusEntity struct {
	value *ControllerServiceRunStatusEntity
	isSet bool
}

func (v NullableControllerServiceRunStatusEntity) Get() *ControllerServiceRunStatusEntity {
	return v.value
}

func (v *NullableControllerServiceRunStatusEntity) Set(val *ControllerServiceRunStatusEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerServiceRunStatusEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerServiceRunStatusEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerServiceRunStatusEntity(val *ControllerServiceRunStatusEntity) *NullableControllerServiceRunStatusEntity {
	return &NullableControllerServiceRunStatusEntity{value: val, isSet: true}
}

func (v NullableControllerServiceRunStatusEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerServiceRunStatusEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


