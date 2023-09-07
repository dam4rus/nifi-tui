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

// checks if the AboutDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AboutDTO{}

// AboutDTO struct for AboutDTO
type AboutDTO struct {
	// The title to be used on the page and in the about dialog.
	Title *string `json:"title,omitempty"`
	// The version of this NiFi.
	Version *string `json:"version,omitempty"`
	// The URI for the NiFi.
	Uri *string `json:"uri,omitempty"`
	// The URL for the content viewer if configured.
	ContentViewerUrl *string `json:"contentViewerUrl,omitempty"`
	// The timezone of the NiFi instance.
	Timezone *string `json:"timezone,omitempty"`
	// Build tag
	BuildTag *string `json:"buildTag,omitempty"`
	// Build revision or commit hash
	BuildRevision *string `json:"buildRevision,omitempty"`
	// Build branch
	BuildBranch *string `json:"buildBranch,omitempty"`
	// Build timestamp
	BuildTimestamp *string `json:"buildTimestamp,omitempty"`
}

// NewAboutDTO instantiates a new AboutDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAboutDTO() *AboutDTO {
	this := AboutDTO{}
	return &this
}

// NewAboutDTOWithDefaults instantiates a new AboutDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAboutDTOWithDefaults() *AboutDTO {
	this := AboutDTO{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AboutDTO) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AboutDTO) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AboutDTO) SetTitle(v string) {
	o.Title = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AboutDTO) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AboutDTO) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AboutDTO) SetVersion(v string) {
	o.Version = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *AboutDTO) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *AboutDTO) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *AboutDTO) SetUri(v string) {
	o.Uri = &v
}

// GetContentViewerUrl returns the ContentViewerUrl field value if set, zero value otherwise.
func (o *AboutDTO) GetContentViewerUrl() string {
	if o == nil || IsNil(o.ContentViewerUrl) {
		var ret string
		return ret
	}
	return *o.ContentViewerUrl
}

// GetContentViewerUrlOk returns a tuple with the ContentViewerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetContentViewerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContentViewerUrl) {
		return nil, false
	}
	return o.ContentViewerUrl, true
}

// HasContentViewerUrl returns a boolean if a field has been set.
func (o *AboutDTO) HasContentViewerUrl() bool {
	if o != nil && !IsNil(o.ContentViewerUrl) {
		return true
	}

	return false
}

// SetContentViewerUrl gets a reference to the given string and assigns it to the ContentViewerUrl field.
func (o *AboutDTO) SetContentViewerUrl(v string) {
	o.ContentViewerUrl = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *AboutDTO) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *AboutDTO) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *AboutDTO) SetTimezone(v string) {
	o.Timezone = &v
}

// GetBuildTag returns the BuildTag field value if set, zero value otherwise.
func (o *AboutDTO) GetBuildTag() string {
	if o == nil || IsNil(o.BuildTag) {
		var ret string
		return ret
	}
	return *o.BuildTag
}

// GetBuildTagOk returns a tuple with the BuildTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetBuildTagOk() (*string, bool) {
	if o == nil || IsNil(o.BuildTag) {
		return nil, false
	}
	return o.BuildTag, true
}

// HasBuildTag returns a boolean if a field has been set.
func (o *AboutDTO) HasBuildTag() bool {
	if o != nil && !IsNil(o.BuildTag) {
		return true
	}

	return false
}

// SetBuildTag gets a reference to the given string and assigns it to the BuildTag field.
func (o *AboutDTO) SetBuildTag(v string) {
	o.BuildTag = &v
}

// GetBuildRevision returns the BuildRevision field value if set, zero value otherwise.
func (o *AboutDTO) GetBuildRevision() string {
	if o == nil || IsNil(o.BuildRevision) {
		var ret string
		return ret
	}
	return *o.BuildRevision
}

// GetBuildRevisionOk returns a tuple with the BuildRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetBuildRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.BuildRevision) {
		return nil, false
	}
	return o.BuildRevision, true
}

// HasBuildRevision returns a boolean if a field has been set.
func (o *AboutDTO) HasBuildRevision() bool {
	if o != nil && !IsNil(o.BuildRevision) {
		return true
	}

	return false
}

// SetBuildRevision gets a reference to the given string and assigns it to the BuildRevision field.
func (o *AboutDTO) SetBuildRevision(v string) {
	o.BuildRevision = &v
}

// GetBuildBranch returns the BuildBranch field value if set, zero value otherwise.
func (o *AboutDTO) GetBuildBranch() string {
	if o == nil || IsNil(o.BuildBranch) {
		var ret string
		return ret
	}
	return *o.BuildBranch
}

// GetBuildBranchOk returns a tuple with the BuildBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetBuildBranchOk() (*string, bool) {
	if o == nil || IsNil(o.BuildBranch) {
		return nil, false
	}
	return o.BuildBranch, true
}

// HasBuildBranch returns a boolean if a field has been set.
func (o *AboutDTO) HasBuildBranch() bool {
	if o != nil && !IsNil(o.BuildBranch) {
		return true
	}

	return false
}

// SetBuildBranch gets a reference to the given string and assigns it to the BuildBranch field.
func (o *AboutDTO) SetBuildBranch(v string) {
	o.BuildBranch = &v
}

// GetBuildTimestamp returns the BuildTimestamp field value if set, zero value otherwise.
func (o *AboutDTO) GetBuildTimestamp() string {
	if o == nil || IsNil(o.BuildTimestamp) {
		var ret string
		return ret
	}
	return *o.BuildTimestamp
}

// GetBuildTimestampOk returns a tuple with the BuildTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AboutDTO) GetBuildTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.BuildTimestamp) {
		return nil, false
	}
	return o.BuildTimestamp, true
}

// HasBuildTimestamp returns a boolean if a field has been set.
func (o *AboutDTO) HasBuildTimestamp() bool {
	if o != nil && !IsNil(o.BuildTimestamp) {
		return true
	}

	return false
}

// SetBuildTimestamp gets a reference to the given string and assigns it to the BuildTimestamp field.
func (o *AboutDTO) SetBuildTimestamp(v string) {
	o.BuildTimestamp = &v
}

func (o AboutDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AboutDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.ContentViewerUrl) {
		toSerialize["contentViewerUrl"] = o.ContentViewerUrl
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
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

type NullableAboutDTO struct {
	value *AboutDTO
	isSet bool
}

func (v NullableAboutDTO) Get() *AboutDTO {
	return v.value
}

func (v *NullableAboutDTO) Set(val *AboutDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAboutDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAboutDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAboutDTO(val *AboutDTO) *NullableAboutDTO {
	return &NullableAboutDTO{value: val, isSet: true}
}

func (v NullableAboutDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAboutDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


