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

// checks if the BuildInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildInfo{}

// BuildInfo struct for BuildInfo
type BuildInfo struct {
	// The version number of the built component.
	Version *string `json:"version,omitempty"`
	// The SCM revision id of the source code used for this build.
	Revision *string `json:"revision,omitempty"`
	// The timestamp (milliseconds since Epoch) of the build.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// The target architecture of the built component.
	TargetArch *string `json:"targetArch,omitempty"`
	// The compiler used for the build
	Compiler *string `json:"compiler,omitempty"`
	// The compiler flags used for the build.
	CompilerFlags *string `json:"compilerFlags,omitempty"`
}

// NewBuildInfo instantiates a new BuildInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildInfo() *BuildInfo {
	this := BuildInfo{}
	return &this
}

// NewBuildInfoWithDefaults instantiates a new BuildInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildInfoWithDefaults() *BuildInfo {
	this := BuildInfo{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BuildInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BuildInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *BuildInfo) SetVersion(v string) {
	o.Version = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BuildInfo) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildInfo) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BuildInfo) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BuildInfo) SetRevision(v string) {
	o.Revision = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BuildInfo) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildInfo) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BuildInfo) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *BuildInfo) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTargetArch returns the TargetArch field value if set, zero value otherwise.
func (o *BuildInfo) GetTargetArch() string {
	if o == nil || IsNil(o.TargetArch) {
		var ret string
		return ret
	}
	return *o.TargetArch
}

// GetTargetArchOk returns a tuple with the TargetArch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildInfo) GetTargetArchOk() (*string, bool) {
	if o == nil || IsNil(o.TargetArch) {
		return nil, false
	}
	return o.TargetArch, true
}

// HasTargetArch returns a boolean if a field has been set.
func (o *BuildInfo) HasTargetArch() bool {
	if o != nil && !IsNil(o.TargetArch) {
		return true
	}

	return false
}

// SetTargetArch gets a reference to the given string and assigns it to the TargetArch field.
func (o *BuildInfo) SetTargetArch(v string) {
	o.TargetArch = &v
}

// GetCompiler returns the Compiler field value if set, zero value otherwise.
func (o *BuildInfo) GetCompiler() string {
	if o == nil || IsNil(o.Compiler) {
		var ret string
		return ret
	}
	return *o.Compiler
}

// GetCompilerOk returns a tuple with the Compiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildInfo) GetCompilerOk() (*string, bool) {
	if o == nil || IsNil(o.Compiler) {
		return nil, false
	}
	return o.Compiler, true
}

// HasCompiler returns a boolean if a field has been set.
func (o *BuildInfo) HasCompiler() bool {
	if o != nil && !IsNil(o.Compiler) {
		return true
	}

	return false
}

// SetCompiler gets a reference to the given string and assigns it to the Compiler field.
func (o *BuildInfo) SetCompiler(v string) {
	o.Compiler = &v
}

// GetCompilerFlags returns the CompilerFlags field value if set, zero value otherwise.
func (o *BuildInfo) GetCompilerFlags() string {
	if o == nil || IsNil(o.CompilerFlags) {
		var ret string
		return ret
	}
	return *o.CompilerFlags
}

// GetCompilerFlagsOk returns a tuple with the CompilerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildInfo) GetCompilerFlagsOk() (*string, bool) {
	if o == nil || IsNil(o.CompilerFlags) {
		return nil, false
	}
	return o.CompilerFlags, true
}

// HasCompilerFlags returns a boolean if a field has been set.
func (o *BuildInfo) HasCompilerFlags() bool {
	if o != nil && !IsNil(o.CompilerFlags) {
		return true
	}

	return false
}

// SetCompilerFlags gets a reference to the given string and assigns it to the CompilerFlags field.
func (o *BuildInfo) SetCompilerFlags(v string) {
	o.CompilerFlags = &v
}

func (o BuildInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TargetArch) {
		toSerialize["targetArch"] = o.TargetArch
	}
	if !IsNil(o.Compiler) {
		toSerialize["compiler"] = o.Compiler
	}
	if !IsNil(o.CompilerFlags) {
		toSerialize["compilerFlags"] = o.CompilerFlags
	}
	return toSerialize, nil
}

type NullableBuildInfo struct {
	value *BuildInfo
	isSet bool
}

func (v NullableBuildInfo) Get() *BuildInfo {
	return v.value
}

func (v *NullableBuildInfo) Set(val *BuildInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildInfo(val *BuildInfo) *NullableBuildInfo {
	return &NullableBuildInfo{value: val, isSet: true}
}

func (v NullableBuildInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

