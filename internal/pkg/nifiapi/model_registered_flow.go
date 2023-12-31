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

// checks if the RegisteredFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisteredFlow{}

// RegisteredFlow struct for RegisteredFlow
type RegisteredFlow struct {
	Identifier *string `json:"identifier,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	BucketIdentifier *string `json:"bucketIdentifier,omitempty"`
	BucketName *string `json:"bucketName,omitempty"`
	CreatedTimestamp *int64 `json:"createdTimestamp,omitempty"`
	LastModifiedTimestamp *int64 `json:"lastModifiedTimestamp,omitempty"`
	Permissions *FlowRegistryPermissions `json:"permissions,omitempty"`
	VersionCount *int64 `json:"versionCount,omitempty"`
	VersionInfo *RegisteredFlowVersionInfo `json:"versionInfo,omitempty"`
}

// NewRegisteredFlow instantiates a new RegisteredFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisteredFlow() *RegisteredFlow {
	this := RegisteredFlow{}
	return &this
}

// NewRegisteredFlowWithDefaults instantiates a new RegisteredFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisteredFlowWithDefaults() *RegisteredFlow {
	this := RegisteredFlow{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *RegisteredFlow) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *RegisteredFlow) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *RegisteredFlow) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegisteredFlow) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegisteredFlow) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegisteredFlow) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegisteredFlow) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegisteredFlow) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegisteredFlow) SetDescription(v string) {
	o.Description = &v
}

// GetBucketIdentifier returns the BucketIdentifier field value if set, zero value otherwise.
func (o *RegisteredFlow) GetBucketIdentifier() string {
	if o == nil || IsNil(o.BucketIdentifier) {
		var ret string
		return ret
	}
	return *o.BucketIdentifier
}

// GetBucketIdentifierOk returns a tuple with the BucketIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetBucketIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.BucketIdentifier) {
		return nil, false
	}
	return o.BucketIdentifier, true
}

// HasBucketIdentifier returns a boolean if a field has been set.
func (o *RegisteredFlow) HasBucketIdentifier() bool {
	if o != nil && !IsNil(o.BucketIdentifier) {
		return true
	}

	return false
}

// SetBucketIdentifier gets a reference to the given string and assigns it to the BucketIdentifier field.
func (o *RegisteredFlow) SetBucketIdentifier(v string) {
	o.BucketIdentifier = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *RegisteredFlow) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *RegisteredFlow) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *RegisteredFlow) SetBucketName(v string) {
	o.BucketName = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *RegisteredFlow) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *RegisteredFlow) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *RegisteredFlow) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

// GetLastModifiedTimestamp returns the LastModifiedTimestamp field value if set, zero value otherwise.
func (o *RegisteredFlow) GetLastModifiedTimestamp() int64 {
	if o == nil || IsNil(o.LastModifiedTimestamp) {
		var ret int64
		return ret
	}
	return *o.LastModifiedTimestamp
}

// GetLastModifiedTimestampOk returns a tuple with the LastModifiedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetLastModifiedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.LastModifiedTimestamp) {
		return nil, false
	}
	return o.LastModifiedTimestamp, true
}

// HasLastModifiedTimestamp returns a boolean if a field has been set.
func (o *RegisteredFlow) HasLastModifiedTimestamp() bool {
	if o != nil && !IsNil(o.LastModifiedTimestamp) {
		return true
	}

	return false
}

// SetLastModifiedTimestamp gets a reference to the given int64 and assigns it to the LastModifiedTimestamp field.
func (o *RegisteredFlow) SetLastModifiedTimestamp(v int64) {
	o.LastModifiedTimestamp = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RegisteredFlow) GetPermissions() FlowRegistryPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret FlowRegistryPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetPermissionsOk() (*FlowRegistryPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RegisteredFlow) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given FlowRegistryPermissions and assigns it to the Permissions field.
func (o *RegisteredFlow) SetPermissions(v FlowRegistryPermissions) {
	o.Permissions = &v
}

// GetVersionCount returns the VersionCount field value if set, zero value otherwise.
func (o *RegisteredFlow) GetVersionCount() int64 {
	if o == nil || IsNil(o.VersionCount) {
		var ret int64
		return ret
	}
	return *o.VersionCount
}

// GetVersionCountOk returns a tuple with the VersionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetVersionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VersionCount) {
		return nil, false
	}
	return o.VersionCount, true
}

// HasVersionCount returns a boolean if a field has been set.
func (o *RegisteredFlow) HasVersionCount() bool {
	if o != nil && !IsNil(o.VersionCount) {
		return true
	}

	return false
}

// SetVersionCount gets a reference to the given int64 and assigns it to the VersionCount field.
func (o *RegisteredFlow) SetVersionCount(v int64) {
	o.VersionCount = &v
}

// GetVersionInfo returns the VersionInfo field value if set, zero value otherwise.
func (o *RegisteredFlow) GetVersionInfo() RegisteredFlowVersionInfo {
	if o == nil || IsNil(o.VersionInfo) {
		var ret RegisteredFlowVersionInfo
		return ret
	}
	return *o.VersionInfo
}

// GetVersionInfoOk returns a tuple with the VersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredFlow) GetVersionInfoOk() (*RegisteredFlowVersionInfo, bool) {
	if o == nil || IsNil(o.VersionInfo) {
		return nil, false
	}
	return o.VersionInfo, true
}

// HasVersionInfo returns a boolean if a field has been set.
func (o *RegisteredFlow) HasVersionInfo() bool {
	if o != nil && !IsNil(o.VersionInfo) {
		return true
	}

	return false
}

// SetVersionInfo gets a reference to the given RegisteredFlowVersionInfo and assigns it to the VersionInfo field.
func (o *RegisteredFlow) SetVersionInfo(v RegisteredFlowVersionInfo) {
	o.VersionInfo = &v
}

func (o RegisteredFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisteredFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BucketIdentifier) {
		toSerialize["bucketIdentifier"] = o.BucketIdentifier
	}
	if !IsNil(o.BucketName) {
		toSerialize["bucketName"] = o.BucketName
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["createdTimestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.LastModifiedTimestamp) {
		toSerialize["lastModifiedTimestamp"] = o.LastModifiedTimestamp
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.VersionCount) {
		toSerialize["versionCount"] = o.VersionCount
	}
	if !IsNil(o.VersionInfo) {
		toSerialize["versionInfo"] = o.VersionInfo
	}
	return toSerialize, nil
}

type NullableRegisteredFlow struct {
	value *RegisteredFlow
	isSet bool
}

func (v NullableRegisteredFlow) Get() *RegisteredFlow {
	return v.value
}

func (v *NullableRegisteredFlow) Set(val *RegisteredFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisteredFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisteredFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisteredFlow(val *RegisteredFlow) *NullableRegisteredFlow {
	return &NullableRegisteredFlow{value: val, isSet: true}
}

func (v NullableRegisteredFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisteredFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


