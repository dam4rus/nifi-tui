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

// checks if the VersionedFlowUpdateRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionedFlowUpdateRequestDTO{}

// VersionedFlowUpdateRequestDTO struct for VersionedFlowUpdateRequestDTO
type VersionedFlowUpdateRequestDTO struct {
	// The unique ID of this request.
	RequestId *string `json:"requestId,omitempty"`
	// The unique ID of the Process Group being updated
	ProcessGroupId *string `json:"processGroupId,omitempty"`
	// The URI for future requests to this drop request.
	Uri *string `json:"uri,omitempty"`
	// The last time this request was updated.
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// Whether or not this request has completed
	Complete *bool `json:"complete,omitempty"`
	// An explanation of why this request failed, or null if this request has not failed
	FailureReason *string `json:"failureReason,omitempty"`
	// The percentage complete for the request, between 0 and 100
	PercentCompleted *int32 `json:"percentCompleted,omitempty"`
	// The state of the request
	State *string `json:"state,omitempty"`
	VersionControlInformation *VersionControlInformationDTO `json:"versionControlInformation,omitempty"`
}

// NewVersionedFlowUpdateRequestDTO instantiates a new VersionedFlowUpdateRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionedFlowUpdateRequestDTO() *VersionedFlowUpdateRequestDTO {
	this := VersionedFlowUpdateRequestDTO{}
	return &this
}

// NewVersionedFlowUpdateRequestDTOWithDefaults instantiates a new VersionedFlowUpdateRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionedFlowUpdateRequestDTOWithDefaults() *VersionedFlowUpdateRequestDTO {
	this := VersionedFlowUpdateRequestDTO{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *VersionedFlowUpdateRequestDTO) SetRequestId(v string) {
	o.RequestId = &v
}

// GetProcessGroupId returns the ProcessGroupId field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetProcessGroupId() string {
	if o == nil || IsNil(o.ProcessGroupId) {
		var ret string
		return ret
	}
	return *o.ProcessGroupId
}

// GetProcessGroupIdOk returns a tuple with the ProcessGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetProcessGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessGroupId) {
		return nil, false
	}
	return o.ProcessGroupId, true
}

// HasProcessGroupId returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasProcessGroupId() bool {
	if o != nil && !IsNil(o.ProcessGroupId) {
		return true
	}

	return false
}

// SetProcessGroupId gets a reference to the given string and assigns it to the ProcessGroupId field.
func (o *VersionedFlowUpdateRequestDTO) SetProcessGroupId(v string) {
	o.ProcessGroupId = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *VersionedFlowUpdateRequestDTO) SetUri(v string) {
	o.Uri = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *VersionedFlowUpdateRequestDTO) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *VersionedFlowUpdateRequestDTO) SetComplete(v bool) {
	o.Complete = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *VersionedFlowUpdateRequestDTO) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetPercentCompleted returns the PercentCompleted field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetPercentCompleted() int32 {
	if o == nil || IsNil(o.PercentCompleted) {
		var ret int32
		return ret
	}
	return *o.PercentCompleted
}

// GetPercentCompletedOk returns a tuple with the PercentCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetPercentCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentCompleted) {
		return nil, false
	}
	return o.PercentCompleted, true
}

// HasPercentCompleted returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasPercentCompleted() bool {
	if o != nil && !IsNil(o.PercentCompleted) {
		return true
	}

	return false
}

// SetPercentCompleted gets a reference to the given int32 and assigns it to the PercentCompleted field.
func (o *VersionedFlowUpdateRequestDTO) SetPercentCompleted(v int32) {
	o.PercentCompleted = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VersionedFlowUpdateRequestDTO) SetState(v string) {
	o.State = &v
}

// GetVersionControlInformation returns the VersionControlInformation field value if set, zero value otherwise.
func (o *VersionedFlowUpdateRequestDTO) GetVersionControlInformation() VersionControlInformationDTO {
	if o == nil || IsNil(o.VersionControlInformation) {
		var ret VersionControlInformationDTO
		return ret
	}
	return *o.VersionControlInformation
}

// GetVersionControlInformationOk returns a tuple with the VersionControlInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionedFlowUpdateRequestDTO) GetVersionControlInformationOk() (*VersionControlInformationDTO, bool) {
	if o == nil || IsNil(o.VersionControlInformation) {
		return nil, false
	}
	return o.VersionControlInformation, true
}

// HasVersionControlInformation returns a boolean if a field has been set.
func (o *VersionedFlowUpdateRequestDTO) HasVersionControlInformation() bool {
	if o != nil && !IsNil(o.VersionControlInformation) {
		return true
	}

	return false
}

// SetVersionControlInformation gets a reference to the given VersionControlInformationDTO and assigns it to the VersionControlInformation field.
func (o *VersionedFlowUpdateRequestDTO) SetVersionControlInformation(v VersionControlInformationDTO) {
	o.VersionControlInformation = &v
}

func (o VersionedFlowUpdateRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionedFlowUpdateRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.ProcessGroupId) {
		toSerialize["processGroupId"] = o.ProcessGroupId
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	if !IsNil(o.PercentCompleted) {
		toSerialize["percentCompleted"] = o.PercentCompleted
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.VersionControlInformation) {
		toSerialize["versionControlInformation"] = o.VersionControlInformation
	}
	return toSerialize, nil
}

type NullableVersionedFlowUpdateRequestDTO struct {
	value *VersionedFlowUpdateRequestDTO
	isSet bool
}

func (v NullableVersionedFlowUpdateRequestDTO) Get() *VersionedFlowUpdateRequestDTO {
	return v.value
}

func (v *NullableVersionedFlowUpdateRequestDTO) Set(val *VersionedFlowUpdateRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionedFlowUpdateRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionedFlowUpdateRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionedFlowUpdateRequestDTO(val *VersionedFlowUpdateRequestDTO) *NullableVersionedFlowUpdateRequestDTO {
	return &NullableVersionedFlowUpdateRequestDTO{value: val, isSet: true}
}

func (v NullableVersionedFlowUpdateRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionedFlowUpdateRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


