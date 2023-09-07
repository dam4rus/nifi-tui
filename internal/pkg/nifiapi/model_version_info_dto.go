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
	"time"
)

// checks if the VersionInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionInfoDTO{}

// VersionInfoDTO struct for VersionInfoDTO
type VersionInfoDTO struct {
	// The version of this NiFi.
	NiFiVersion *string `json:"niFiVersion,omitempty"`
	// Java JVM vendor
	JavaVendor *string `json:"javaVendor,omitempty"`
	// Java version
	JavaVersion *string `json:"javaVersion,omitempty"`
	// Host operating system name
	OsName *string `json:"osName,omitempty"`
	// Host operating system version
	OsVersion *string `json:"osVersion,omitempty"`
	// Host operating system architecture
	OsArchitecture *string `json:"osArchitecture,omitempty"`
	// Build tag
	BuildTag *string `json:"buildTag,omitempty"`
	// Build revision or commit hash
	BuildRevision *string `json:"buildRevision,omitempty"`
	// Build branch
	BuildBranch *string `json:"buildBranch,omitempty"`
	// Build timestamp
	BuildTimestamp *time.Time `json:"buildTimestamp,omitempty"`
}

// NewVersionInfoDTO instantiates a new VersionInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionInfoDTO() *VersionInfoDTO {
	this := VersionInfoDTO{}
	return &this
}

// NewVersionInfoDTOWithDefaults instantiates a new VersionInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionInfoDTOWithDefaults() *VersionInfoDTO {
	this := VersionInfoDTO{}
	return &this
}

// GetNiFiVersion returns the NiFiVersion field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetNiFiVersion() string {
	if o == nil || IsNil(o.NiFiVersion) {
		var ret string
		return ret
	}
	return *o.NiFiVersion
}

// GetNiFiVersionOk returns a tuple with the NiFiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetNiFiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.NiFiVersion) {
		return nil, false
	}
	return o.NiFiVersion, true
}

// HasNiFiVersion returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasNiFiVersion() bool {
	if o != nil && !IsNil(o.NiFiVersion) {
		return true
	}

	return false
}

// SetNiFiVersion gets a reference to the given string and assigns it to the NiFiVersion field.
func (o *VersionInfoDTO) SetNiFiVersion(v string) {
	o.NiFiVersion = &v
}

// GetJavaVendor returns the JavaVendor field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetJavaVendor() string {
	if o == nil || IsNil(o.JavaVendor) {
		var ret string
		return ret
	}
	return *o.JavaVendor
}

// GetJavaVendorOk returns a tuple with the JavaVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetJavaVendorOk() (*string, bool) {
	if o == nil || IsNil(o.JavaVendor) {
		return nil, false
	}
	return o.JavaVendor, true
}

// HasJavaVendor returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasJavaVendor() bool {
	if o != nil && !IsNil(o.JavaVendor) {
		return true
	}

	return false
}

// SetJavaVendor gets a reference to the given string and assigns it to the JavaVendor field.
func (o *VersionInfoDTO) SetJavaVendor(v string) {
	o.JavaVendor = &v
}

// GetJavaVersion returns the JavaVersion field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetJavaVersion() string {
	if o == nil || IsNil(o.JavaVersion) {
		var ret string
		return ret
	}
	return *o.JavaVersion
}

// GetJavaVersionOk returns a tuple with the JavaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetJavaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.JavaVersion) {
		return nil, false
	}
	return o.JavaVersion, true
}

// HasJavaVersion returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasJavaVersion() bool {
	if o != nil && !IsNil(o.JavaVersion) {
		return true
	}

	return false
}

// SetJavaVersion gets a reference to the given string and assigns it to the JavaVersion field.
func (o *VersionInfoDTO) SetJavaVersion(v string) {
	o.JavaVersion = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetOsName() string {
	if o == nil || IsNil(o.OsName) {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetOsNameOk() (*string, bool) {
	if o == nil || IsNil(o.OsName) {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasOsName() bool {
	if o != nil && !IsNil(o.OsName) {
		return true
	}

	return false
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *VersionInfoDTO) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *VersionInfoDTO) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetOsArchitecture returns the OsArchitecture field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetOsArchitecture() string {
	if o == nil || IsNil(o.OsArchitecture) {
		var ret string
		return ret
	}
	return *o.OsArchitecture
}

// GetOsArchitectureOk returns a tuple with the OsArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetOsArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.OsArchitecture) {
		return nil, false
	}
	return o.OsArchitecture, true
}

// HasOsArchitecture returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasOsArchitecture() bool {
	if o != nil && !IsNil(o.OsArchitecture) {
		return true
	}

	return false
}

// SetOsArchitecture gets a reference to the given string and assigns it to the OsArchitecture field.
func (o *VersionInfoDTO) SetOsArchitecture(v string) {
	o.OsArchitecture = &v
}

// GetBuildTag returns the BuildTag field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildTag() string {
	if o == nil || IsNil(o.BuildTag) {
		var ret string
		return ret
	}
	return *o.BuildTag
}

// GetBuildTagOk returns a tuple with the BuildTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildTagOk() (*string, bool) {
	if o == nil || IsNil(o.BuildTag) {
		return nil, false
	}
	return o.BuildTag, true
}

// HasBuildTag returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildTag() bool {
	if o != nil && !IsNil(o.BuildTag) {
		return true
	}

	return false
}

// SetBuildTag gets a reference to the given string and assigns it to the BuildTag field.
func (o *VersionInfoDTO) SetBuildTag(v string) {
	o.BuildTag = &v
}

// GetBuildRevision returns the BuildRevision field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildRevision() string {
	if o == nil || IsNil(o.BuildRevision) {
		var ret string
		return ret
	}
	return *o.BuildRevision
}

// GetBuildRevisionOk returns a tuple with the BuildRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.BuildRevision) {
		return nil, false
	}
	return o.BuildRevision, true
}

// HasBuildRevision returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildRevision() bool {
	if o != nil && !IsNil(o.BuildRevision) {
		return true
	}

	return false
}

// SetBuildRevision gets a reference to the given string and assigns it to the BuildRevision field.
func (o *VersionInfoDTO) SetBuildRevision(v string) {
	o.BuildRevision = &v
}

// GetBuildBranch returns the BuildBranch field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildBranch() string {
	if o == nil || IsNil(o.BuildBranch) {
		var ret string
		return ret
	}
	return *o.BuildBranch
}

// GetBuildBranchOk returns a tuple with the BuildBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildBranchOk() (*string, bool) {
	if o == nil || IsNil(o.BuildBranch) {
		return nil, false
	}
	return o.BuildBranch, true
}

// HasBuildBranch returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildBranch() bool {
	if o != nil && !IsNil(o.BuildBranch) {
		return true
	}

	return false
}

// SetBuildBranch gets a reference to the given string and assigns it to the BuildBranch field.
func (o *VersionInfoDTO) SetBuildBranch(v string) {
	o.BuildBranch = &v
}

// GetBuildTimestamp returns the BuildTimestamp field value if set, zero value otherwise.
func (o *VersionInfoDTO) GetBuildTimestamp() time.Time {
	if o == nil || IsNil(o.BuildTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.BuildTimestamp
}

// GetBuildTimestampOk returns a tuple with the BuildTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfoDTO) GetBuildTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BuildTimestamp) {
		return nil, false
	}
	return o.BuildTimestamp, true
}

// HasBuildTimestamp returns a boolean if a field has been set.
func (o *VersionInfoDTO) HasBuildTimestamp() bool {
	if o != nil && !IsNil(o.BuildTimestamp) {
		return true
	}

	return false
}

// SetBuildTimestamp gets a reference to the given time.Time and assigns it to the BuildTimestamp field.
func (o *VersionInfoDTO) SetBuildTimestamp(v time.Time) {
	o.BuildTimestamp = &v
}

func (o VersionInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NiFiVersion) {
		toSerialize["niFiVersion"] = o.NiFiVersion
	}
	if !IsNil(o.JavaVendor) {
		toSerialize["javaVendor"] = o.JavaVendor
	}
	if !IsNil(o.JavaVersion) {
		toSerialize["javaVersion"] = o.JavaVersion
	}
	if !IsNil(o.OsName) {
		toSerialize["osName"] = o.OsName
	}
	if !IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !IsNil(o.OsArchitecture) {
		toSerialize["osArchitecture"] = o.OsArchitecture
	}
	if !IsNil(o.BuildTag) {
		toSerialize["buildTag"] = o.BuildTag
	}
	if !IsNil(o.BuildRevision) {
		toSerialize["buildRevision"] = o.BuildRevision
	}
	if !IsNil(o.BuildBranch) {
		toSerialize["buildBranch"] = o.BuildBranch
	}
	if !IsNil(o.BuildTimestamp) {
		toSerialize["buildTimestamp"] = o.BuildTimestamp
	}
	return toSerialize, nil
}

type NullableVersionInfoDTO struct {
	value *VersionInfoDTO
	isSet bool
}

func (v NullableVersionInfoDTO) Get() *VersionInfoDTO {
	return v.value
}

func (v *NullableVersionInfoDTO) Set(val *VersionInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionInfoDTO(val *VersionInfoDTO) *NullableVersionInfoDTO {
	return &NullableVersionInfoDTO{value: val, isSet: true}
}

func (v NullableVersionInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


