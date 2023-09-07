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

// checks if the VersionedFlowCoordinates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionedFlowCoordinates{}

// VersionedFlowCoordinates struct for VersionedFlowCoordinates
type VersionedFlowCoordinates struct {
	// The identifier of the Flow Registry that contains the flow
	RegistryId *string `json:"registryId,omitempty"`
	// The location of the Flow Registry that stores the flow
	StorageLocation *string `json:"storageLocation,omitempty"`
	// The URL of the Flow Registry that contains the flow
	RegistryUrl *string `json:"registryUrl,omitempty"`
	// The UUID of the bucket that the flow resides in
	BucketId *string `json:"bucketId,omitempty"`
	// The UUID of the flow
	FlowId *string `json:"flowId,omitempty"`
	// The version of the flow
	Version *int32 `json:"version,omitempty"`
	// Whether or not these coordinates point to the latest version of the flow
	Latest *bool `json:"latest,omitempty"`
}

// NewVersionedFlowCoordinates instantiates a new VersionedFlowCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedFlowCoordinates() *VersionedFlowCoordinates {
	this := VersionedFlowCoordinates{}
	return &this
}

// NewVersionedFlowCoordinatesWithDefaults instantiates a new VersionedFlowCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedFlowCoordinatesWithDefaults() *VersionedFlowCoordinates {
	this := VersionedFlowCoordinates{}
	return &this
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *VersionedFlowCoordinates) GetRegistryId() string {
	if o == nil || IsNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowCoordinates) GetRegistryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryId) {
		return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *VersionedFlowCoordinates) HasRegistryId() bool {
	if o != nil && !IsNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *VersionedFlowCoordinates) SetRegistryId(v string) {
	o.RegistryId = &v
}

// GetStorageLocation returns the StorageLocation field value if set, zero value otherwise.
func (o *VersionedFlowCoordinates) GetStorageLocation() string {
	if o == nil || IsNil(o.StorageLocation) {
		var ret string
		return ret
	}
	return *o.StorageLocation
}

// GetStorageLocationOk returns a tuple with the StorageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowCoordinates) GetStorageLocationOk() (*string, bool) {
	if o == nil || IsNil(o.StorageLocation) {
		return nil, false
	}
	return o.StorageLocation, true
}

// HasStorageLocation returns a boolean if a field has been set.
func (o *VersionedFlowCoordinates) HasStorageLocation() bool {
	if o != nil && !IsNil(o.StorageLocation) {
		return true
	}

	return false
}

// SetStorageLocation gets a reference to the given string and assigns it to the StorageLocation field.
func (o *VersionedFlowCoordinates) SetStorageLocation(v string) {
	o.StorageLocation = &v
}

// GetRegistryUrl returns the RegistryUrl field value if set, zero value otherwise.
func (o *VersionedFlowCoordinates) GetRegistryUrl() string {
	if o == nil || IsNil(o.RegistryUrl) {
		var ret string
		return ret
	}
	return *o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowCoordinates) GetRegistryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryUrl) {
		return nil, false
	}
	return o.RegistryUrl, true
}

// HasRegistryUrl returns a boolean if a field has been set.
func (o *VersionedFlowCoordinates) HasRegistryUrl() bool {
	if o != nil && !IsNil(o.RegistryUrl) {
		return true
	}

	return false
}

// SetRegistryUrl gets a reference to the given string and assigns it to the RegistryUrl field.
func (o *VersionedFlowCoordinates) SetRegistryUrl(v string) {
	o.RegistryUrl = &v
}

// GetBucketId returns the BucketId field value if set, zero value otherwise.
func (o *VersionedFlowCoordinates) GetBucketId() string {
	if o == nil || IsNil(o.BucketId) {
		var ret string
		return ret
	}
	return *o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowCoordinates) GetBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.BucketId) {
		return nil, false
	}
	return o.BucketId, true
}

// HasBucketId returns a boolean if a field has been set.
func (o *VersionedFlowCoordinates) HasBucketId() bool {
	if o != nil && !IsNil(o.BucketId) {
		return true
	}

	return false
}

// SetBucketId gets a reference to the given string and assigns it to the BucketId field.
func (o *VersionedFlowCoordinates) SetBucketId(v string) {
	o.BucketId = &v
}

// GetFlowId returns the FlowId field value if set, zero value otherwise.
func (o *VersionedFlowCoordinates) GetFlowId() string {
	if o == nil || IsNil(o.FlowId) {
		var ret string
		return ret
	}
	return *o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowCoordinates) GetFlowIdOk() (*string, bool) {
	if o == nil || IsNil(o.FlowId) {
		return nil, false
	}
	return o.FlowId, true
}

// HasFlowId returns a boolean if a field has been set.
func (o *VersionedFlowCoordinates) HasFlowId() bool {
	if o != nil && !IsNil(o.FlowId) {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given string and assigns it to the FlowId field.
func (o *VersionedFlowCoordinates) SetFlowId(v string) {
	o.FlowId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VersionedFlowCoordinates) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowCoordinates) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VersionedFlowCoordinates) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *VersionedFlowCoordinates) SetVersion(v int32) {
	o.Version = &v
}

// GetLatest returns the Latest field value if set, zero value otherwise.
func (o *VersionedFlowCoordinates) GetLatest() bool {
	if o == nil || IsNil(o.Latest) {
		var ret bool
		return ret
	}
	return *o.Latest
}

// GetLatestOk returns a tuple with the Latest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowCoordinates) GetLatestOk() (*bool, bool) {
	if o == nil || IsNil(o.Latest) {
		return nil, false
	}
	return o.Latest, true
}

// HasLatest returns a boolean if a field has been set.
func (o *VersionedFlowCoordinates) HasLatest() bool {
	if o != nil && !IsNil(o.Latest) {
		return true
	}

	return false
}

// SetLatest gets a reference to the given bool and assigns it to the Latest field.
func (o *VersionedFlowCoordinates) SetLatest(v bool) {
	o.Latest = &v
}

func (o VersionedFlowCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionedFlowCoordinates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegistryId) {
		toSerialize["registryId"] = o.RegistryId
	}
	if !IsNil(o.StorageLocation) {
		toSerialize["storageLocation"] = o.StorageLocation
	}
	if !IsNil(o.RegistryUrl) {
		toSerialize["registryUrl"] = o.RegistryUrl
	}
	if !IsNil(o.BucketId) {
		toSerialize["bucketId"] = o.BucketId
	}
	if !IsNil(o.FlowId) {
		toSerialize["flowId"] = o.FlowId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Latest) {
		toSerialize["latest"] = o.Latest
	}
	return toSerialize, nil
}

type NullableVersionedFlowCoordinates struct {
	value *VersionedFlowCoordinates
	isSet bool
}

func (v NullableVersionedFlowCoordinates) Get() *VersionedFlowCoordinates {
	return v.value
}

func (v *NullableVersionedFlowCoordinates) Set(val *VersionedFlowCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedFlowCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedFlowCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedFlowCoordinates(val *VersionedFlowCoordinates) *NullableVersionedFlowCoordinates {
	return &NullableVersionedFlowCoordinates{value: val, isSet: true}
}

func (v NullableVersionedFlowCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedFlowCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


