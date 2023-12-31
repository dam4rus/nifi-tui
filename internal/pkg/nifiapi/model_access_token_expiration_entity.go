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

// checks if the AccessTokenExpirationEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenExpirationEntity{}

// AccessTokenExpirationEntity struct for AccessTokenExpirationEntity
type AccessTokenExpirationEntity struct {
	AccessTokenExpiration *AccessTokenExpirationDTO `json:"accessTokenExpiration,omitempty"`
}

// NewAccessTokenExpirationEntity instantiates a new AccessTokenExpirationEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenExpirationEntity() *AccessTokenExpirationEntity {
	this := AccessTokenExpirationEntity{}
	return &this
}

// NewAccessTokenExpirationEntityWithDefaults instantiates a new AccessTokenExpirationEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenExpirationEntityWithDefaults() *AccessTokenExpirationEntity {
	this := AccessTokenExpirationEntity{}
	return &this
}

// GetAccessTokenExpiration returns the AccessTokenExpiration field value if set, zero value otherwise.
func (o *AccessTokenExpirationEntity) GetAccessTokenExpiration() AccessTokenExpirationDTO {
	if o == nil || IsNil(o.AccessTokenExpiration) {
		var ret AccessTokenExpirationDTO
		return ret
	}
	return *o.AccessTokenExpiration
}

// GetAccessTokenExpirationOk returns a tuple with the AccessTokenExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenExpirationEntity) GetAccessTokenExpirationOk() (*AccessTokenExpirationDTO, bool) {
	if o == nil || IsNil(o.AccessTokenExpiration) {
		return nil, false
	}
	return o.AccessTokenExpiration, true
}

// HasAccessTokenExpiration returns a boolean if a field has been set.
func (o *AccessTokenExpirationEntity) HasAccessTokenExpiration() bool {
	if o != nil && !IsNil(o.AccessTokenExpiration) {
		return true
	}

	return false
}

// SetAccessTokenExpiration gets a reference to the given AccessTokenExpirationDTO and assigns it to the AccessTokenExpiration field.
func (o *AccessTokenExpirationEntity) SetAccessTokenExpiration(v AccessTokenExpirationDTO) {
	o.AccessTokenExpiration = &v
}

func (o AccessTokenExpirationEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenExpirationEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessTokenExpiration) {
		toSerialize["accessTokenExpiration"] = o.AccessTokenExpiration
	}
	return toSerialize, nil
}

type NullableAccessTokenExpirationEntity struct {
	value *AccessTokenExpirationEntity
	isSet bool
}

func (v NullableAccessTokenExpirationEntity) Get() *AccessTokenExpirationEntity {
	return v.value
}

func (v *NullableAccessTokenExpirationEntity) Set(val *AccessTokenExpirationEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenExpirationEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenExpirationEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenExpirationEntity(val *AccessTokenExpirationEntity) *NullableAccessTokenExpirationEntity {
	return &NullableAccessTokenExpirationEntity{value: val, isSet: true}
}

func (v NullableAccessTokenExpirationEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenExpirationEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


