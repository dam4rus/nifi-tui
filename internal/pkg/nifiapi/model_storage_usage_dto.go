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

// checks if the StorageUsageDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageUsageDTO{}

// StorageUsageDTO struct for StorageUsageDTO
type StorageUsageDTO struct {
	// The identifier of this storage location. The identifier will correspond to the identifier keyed in the storage configuration.
	Identifier *string `json:"identifier,omitempty"`
	// Amount of free space.
	FreeSpace *string `json:"freeSpace,omitempty"`
	// Amount of total space.
	TotalSpace *string `json:"totalSpace,omitempty"`
	// Amount of used space.
	UsedSpace *string `json:"usedSpace,omitempty"`
	// The number of bytes of free space.
	FreeSpaceBytes *int64 `json:"freeSpaceBytes,omitempty"`
	// The number of bytes of total space.
	TotalSpaceBytes *int64 `json:"totalSpaceBytes,omitempty"`
	// The number of bytes of used space.
	UsedSpaceBytes *int64 `json:"usedSpaceBytes,omitempty"`
	// Utilization of this storage location.
	Utilization *string `json:"utilization,omitempty"`
}

// NewStorageUsageDTO instantiates a new StorageUsageDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageUsageDTO() *StorageUsageDTO {
	this := StorageUsageDTO{}
	return &this
}

// NewStorageUsageDTOWithDefaults instantiates a new StorageUsageDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageUsageDTOWithDefaults() *StorageUsageDTO {
	this := StorageUsageDTO{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *StorageUsageDTO) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetFreeSpace returns the FreeSpace field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetFreeSpace() string {
	if o == nil || IsNil(o.FreeSpace) {
		var ret string
		return ret
	}
	return *o.FreeSpace
}

// GetFreeSpaceOk returns a tuple with the FreeSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetFreeSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.FreeSpace) {
		return nil, false
	}
	return o.FreeSpace, true
}

// HasFreeSpace returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasFreeSpace() bool {
	if o != nil && !IsNil(o.FreeSpace) {
		return true
	}

	return false
}

// SetFreeSpace gets a reference to the given string and assigns it to the FreeSpace field.
func (o *StorageUsageDTO) SetFreeSpace(v string) {
	o.FreeSpace = &v
}

// GetTotalSpace returns the TotalSpace field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetTotalSpace() string {
	if o == nil || IsNil(o.TotalSpace) {
		var ret string
		return ret
	}
	return *o.TotalSpace
}

// GetTotalSpaceOk returns a tuple with the TotalSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetTotalSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSpace) {
		return nil, false
	}
	return o.TotalSpace, true
}

// HasTotalSpace returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasTotalSpace() bool {
	if o != nil && !IsNil(o.TotalSpace) {
		return true
	}

	return false
}

// SetTotalSpace gets a reference to the given string and assigns it to the TotalSpace field.
func (o *StorageUsageDTO) SetTotalSpace(v string) {
	o.TotalSpace = &v
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetUsedSpace() string {
	if o == nil || IsNil(o.UsedSpace) {
		var ret string
		return ret
	}
	return *o.UsedSpace
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetUsedSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.UsedSpace) {
		return nil, false
	}
	return o.UsedSpace, true
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasUsedSpace() bool {
	if o != nil && !IsNil(o.UsedSpace) {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given string and assigns it to the UsedSpace field.
func (o *StorageUsageDTO) SetUsedSpace(v string) {
	o.UsedSpace = &v
}

// GetFreeSpaceBytes returns the FreeSpaceBytes field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetFreeSpaceBytes() int64 {
	if o == nil || IsNil(o.FreeSpaceBytes) {
		var ret int64
		return ret
	}
	return *o.FreeSpaceBytes
}

// GetFreeSpaceBytesOk returns a tuple with the FreeSpaceBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetFreeSpaceBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.FreeSpaceBytes) {
		return nil, false
	}
	return o.FreeSpaceBytes, true
}

// HasFreeSpaceBytes returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasFreeSpaceBytes() bool {
	if o != nil && !IsNil(o.FreeSpaceBytes) {
		return true
	}

	return false
}

// SetFreeSpaceBytes gets a reference to the given int64 and assigns it to the FreeSpaceBytes field.
func (o *StorageUsageDTO) SetFreeSpaceBytes(v int64) {
	o.FreeSpaceBytes = &v
}

// GetTotalSpaceBytes returns the TotalSpaceBytes field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetTotalSpaceBytes() int64 {
	if o == nil || IsNil(o.TotalSpaceBytes) {
		var ret int64
		return ret
	}
	return *o.TotalSpaceBytes
}

// GetTotalSpaceBytesOk returns a tuple with the TotalSpaceBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetTotalSpaceBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalSpaceBytes) {
		return nil, false
	}
	return o.TotalSpaceBytes, true
}

// HasTotalSpaceBytes returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasTotalSpaceBytes() bool {
	if o != nil && !IsNil(o.TotalSpaceBytes) {
		return true
	}

	return false
}

// SetTotalSpaceBytes gets a reference to the given int64 and assigns it to the TotalSpaceBytes field.
func (o *StorageUsageDTO) SetTotalSpaceBytes(v int64) {
	o.TotalSpaceBytes = &v
}

// GetUsedSpaceBytes returns the UsedSpaceBytes field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetUsedSpaceBytes() int64 {
	if o == nil || IsNil(o.UsedSpaceBytes) {
		var ret int64
		return ret
	}
	return *o.UsedSpaceBytes
}

// GetUsedSpaceBytesOk returns a tuple with the UsedSpaceBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetUsedSpaceBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedSpaceBytes) {
		return nil, false
	}
	return o.UsedSpaceBytes, true
}

// HasUsedSpaceBytes returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasUsedSpaceBytes() bool {
	if o != nil && !IsNil(o.UsedSpaceBytes) {
		return true
	}

	return false
}

// SetUsedSpaceBytes gets a reference to the given int64 and assigns it to the UsedSpaceBytes field.
func (o *StorageUsageDTO) SetUsedSpaceBytes(v int64) {
	o.UsedSpaceBytes = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *StorageUsageDTO) GetUtilization() string {
	if o == nil || IsNil(o.Utilization) {
		var ret string
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUsageDTO) GetUtilizationOk() (*string, bool) {
	if o == nil || IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *StorageUsageDTO) HasUtilization() bool {
	if o != nil && !IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given string and assigns it to the Utilization field.
func (o *StorageUsageDTO) SetUtilization(v string) {
	o.Utilization = &v
}

func (o StorageUsageDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageUsageDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.FreeSpace) {
		toSerialize["freeSpace"] = o.FreeSpace
	}
	if !IsNil(o.TotalSpace) {
		toSerialize["totalSpace"] = o.TotalSpace
	}
	if !IsNil(o.UsedSpace) {
		toSerialize["usedSpace"] = o.UsedSpace
	}
	if !IsNil(o.FreeSpaceBytes) {
		toSerialize["freeSpaceBytes"] = o.FreeSpaceBytes
	}
	if !IsNil(o.TotalSpaceBytes) {
		toSerialize["totalSpaceBytes"] = o.TotalSpaceBytes
	}
	if !IsNil(o.UsedSpaceBytes) {
		toSerialize["usedSpaceBytes"] = o.UsedSpaceBytes
	}
	if !IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	return toSerialize, nil
}

type NullableStorageUsageDTO struct {
	value *StorageUsageDTO
	isSet bool
}

func (v NullableStorageUsageDTO) Get() *StorageUsageDTO {
	return v.value
}

func (v *NullableStorageUsageDTO) Set(val *StorageUsageDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageUsageDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageUsageDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageUsageDTO(val *StorageUsageDTO) *NullableStorageUsageDTO {
	return &NullableStorageUsageDTO{value: val, isSet: true}
}

func (v NullableStorageUsageDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageUsageDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


