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

// checks if the AccessTokenExpirationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenExpirationDTO{}

// AccessTokenExpirationDTO struct for AccessTokenExpirationDTO
type AccessTokenExpirationDTO struct {
	// Token Expiration
	Expiration *string `json:"expiration,omitempty"`
}

// NewAccessTokenExpirationDTO instantiates a new AccessTokenExpirationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenExpirationDTO() *AccessTokenExpirationDTO {
	this := AccessTokenExpirationDTO{}
	return &this
}

// NewAccessTokenExpirationDTOWithDefaults instantiates a new AccessTokenExpirationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenExpirationDTOWithDefaults() *AccessTokenExpirationDTO {
	this := AccessTokenExpirationDTO{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *AccessTokenExpirationDTO) GetExpiration() string {
	if o == nil || IsNil(o.Expiration) {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenExpirationDTO) GetExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *AccessTokenExpirationDTO) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *AccessTokenExpirationDTO) SetExpiration(v string) {
	o.Expiration = &v
}

func (o AccessTokenExpirationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenExpirationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	return toSerialize, nil
}

type NullableAccessTokenExpirationDTO struct {
	value *AccessTokenExpirationDTO
	isSet bool
}

func (v NullableAccessTokenExpirationDTO) Get() *AccessTokenExpirationDTO {
	return v.value
}

func (v *NullableAccessTokenExpirationDTO) Set(val *AccessTokenExpirationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenExpirationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenExpirationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenExpirationDTO(val *AccessTokenExpirationDTO) *NullableAccessTokenExpirationDTO {
	return &NullableAccessTokenExpirationDTO{value: val, isSet: true}
}

func (v NullableAccessTokenExpirationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenExpirationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


