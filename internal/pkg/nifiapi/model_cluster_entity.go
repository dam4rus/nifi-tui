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

// checks if the ClusterEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterEntity{}

// ClusterEntity struct for ClusterEntity
type ClusterEntity struct {
	Cluster *ClusterDTO `json:"cluster,omitempty"`
}

// NewClusterEntity instantiates a new ClusterEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterEntity() *ClusterEntity {
	this := ClusterEntity{}
	return &this
}

// NewClusterEntityWithDefaults instantiates a new ClusterEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterEntityWithDefaults() *ClusterEntity {
	this := ClusterEntity{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ClusterEntity) GetCluster() ClusterDTO {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterDTO
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEntity) GetClusterOk() (*ClusterDTO, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ClusterEntity) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterDTO and assigns it to the Cluster field.
func (o *ClusterEntity) SetCluster(v ClusterDTO) {
	o.Cluster = &v
}

func (o ClusterEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	return toSerialize, nil
}

type NullableClusterEntity struct {
	value *ClusterEntity
	isSet bool
}

func (v NullableClusterEntity) Get() *ClusterEntity {
	return v.value
}

func (v *NullableClusterEntity) Set(val *ClusterEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterEntity(val *ClusterEntity) *NullableClusterEntity {
	return &NullableClusterEntity{value: val, isSet: true}
}

func (v NullableClusterEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


