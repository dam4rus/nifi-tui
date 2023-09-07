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

// checks if the ConnectionsEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionsEntity{}

// ConnectionsEntity struct for ConnectionsEntity
type ConnectionsEntity struct {
	Connections []ConnectionEntity `json:"connections,omitempty"`
}

// NewConnectionsEntity instantiates a new ConnectionsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionsEntity() *ConnectionsEntity {
	this := ConnectionsEntity{}
	return &this
}

// NewConnectionsEntityWithDefaults instantiates a new ConnectionsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionsEntityWithDefaults() *ConnectionsEntity {
	this := ConnectionsEntity{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *ConnectionsEntity) GetConnections() []ConnectionEntity {
	if o == nil || IsNil(o.Connections) {
		var ret []ConnectionEntity
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionsEntity) GetConnectionsOk() ([]ConnectionEntity, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *ConnectionsEntity) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []ConnectionEntity and assigns it to the Connections field.
func (o *ConnectionsEntity) SetConnections(v []ConnectionEntity) {
	o.Connections = v
}

func (o ConnectionsEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionsEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	return toSerialize, nil
}

type NullableConnectionsEntity struct {
	value *ConnectionsEntity
	isSet bool
}

func (v NullableConnectionsEntity) Get() *ConnectionsEntity {
	return v.value
}

func (v *NullableConnectionsEntity) Set(val *ConnectionsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionsEntity(val *ConnectionsEntity) *NullableConnectionsEntity {
	return &NullableConnectionsEntity{value: val, isSet: true}
}

func (v NullableConnectionsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

