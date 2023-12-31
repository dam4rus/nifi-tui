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

// checks if the VariableRegistryUpdateRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableRegistryUpdateRequestDTO{}

// VariableRegistryUpdateRequestDTO struct for VariableRegistryUpdateRequestDTO
type VariableRegistryUpdateRequestDTO struct {
	// The ID of the request
	RequestId *string `json:"requestId,omitempty"`
	// The URI for the request
	Uri *string `json:"uri,omitempty"`
	// The timestamp of when the request was submitted
	SubmissionTime *time.Time `json:"submissionTime,omitempty"`
	// The timestamp of when the request was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Whether or not the request is completed
	Complete *bool `json:"complete,omitempty"`
	// The reason for the request failing, or null if the request has not failed
	FailureReason *string `json:"failureReason,omitempty"`
	// A value between 0 and 100 (inclusive) indicating how close the request is to completion
	PercentCompleted *int32 `json:"percentCompleted,omitempty"`
	// A description of the current state of the request
	State *string `json:"state,omitempty"`
	// The steps that are required in order to complete the request, along with the status of each
	UpdateSteps []VariableRegistryUpdateStepDTO `json:"updateSteps,omitempty"`
	// The unique ID of the Process Group that the variable registry belongs to
	ProcessGroupId *string `json:"processGroupId,omitempty"`
	// A set of all components that will be affected if the value of this variable is changed
	AffectedComponents []AffectedComponentEntity `json:"affectedComponents,omitempty"`
}

// NewVariableRegistryUpdateRequestDTO instantiates a new VariableRegistryUpdateRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableRegistryUpdateRequestDTO() *VariableRegistryUpdateRequestDTO {
	this := VariableRegistryUpdateRequestDTO{}
	return &this
}

// NewVariableRegistryUpdateRequestDTOWithDefaults instantiates a new VariableRegistryUpdateRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableRegistryUpdateRequestDTOWithDefaults() *VariableRegistryUpdateRequestDTO {
	this := VariableRegistryUpdateRequestDTO{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *VariableRegistryUpdateRequestDTO) SetRequestId(v string) {
	o.RequestId = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *VariableRegistryUpdateRequestDTO) SetUri(v string) {
	o.Uri = &v
}

// GetSubmissionTime returns the SubmissionTime field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetSubmissionTime() time.Time {
	if o == nil || IsNil(o.SubmissionTime) {
		var ret time.Time
		return ret
	}
	return *o.SubmissionTime
}

// GetSubmissionTimeOk returns a tuple with the SubmissionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetSubmissionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmissionTime) {
		return nil, false
	}
	return o.SubmissionTime, true
}

// HasSubmissionTime returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasSubmissionTime() bool {
	if o != nil && !IsNil(o.SubmissionTime) {
		return true
	}

	return false
}

// SetSubmissionTime gets a reference to the given time.Time and assigns it to the SubmissionTime field.
func (o *VariableRegistryUpdateRequestDTO) SetSubmissionTime(v time.Time) {
	o.SubmissionTime = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *VariableRegistryUpdateRequestDTO) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *VariableRegistryUpdateRequestDTO) SetComplete(v bool) {
	o.Complete = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *VariableRegistryUpdateRequestDTO) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetPercentCompleted returns the PercentCompleted field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetPercentCompleted() int32 {
	if o == nil || IsNil(o.PercentCompleted) {
		var ret int32
		return ret
	}
	return *o.PercentCompleted
}

// GetPercentCompletedOk returns a tuple with the PercentCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetPercentCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentCompleted) {
		return nil, false
	}
	return o.PercentCompleted, true
}

// HasPercentCompleted returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasPercentCompleted() bool {
	if o != nil && !IsNil(o.PercentCompleted) {
		return true
	}

	return false
}

// SetPercentCompleted gets a reference to the given int32 and assigns it to the PercentCompleted field.
func (o *VariableRegistryUpdateRequestDTO) SetPercentCompleted(v int32) {
	o.PercentCompleted = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VariableRegistryUpdateRequestDTO) SetState(v string) {
	o.State = &v
}

// GetUpdateSteps returns the UpdateSteps field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetUpdateSteps() []VariableRegistryUpdateStepDTO {
	if o == nil || IsNil(o.UpdateSteps) {
		var ret []VariableRegistryUpdateStepDTO
		return ret
	}
	return o.UpdateSteps
}

// GetUpdateStepsOk returns a tuple with the UpdateSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetUpdateStepsOk() ([]VariableRegistryUpdateStepDTO, bool) {
	if o == nil || IsNil(o.UpdateSteps) {
		return nil, false
	}
	return o.UpdateSteps, true
}

// HasUpdateSteps returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasUpdateSteps() bool {
	if o != nil && !IsNil(o.UpdateSteps) {
		return true
	}

	return false
}

// SetUpdateSteps gets a reference to the given []VariableRegistryUpdateStepDTO and assigns it to the UpdateSteps field.
func (o *VariableRegistryUpdateRequestDTO) SetUpdateSteps(v []VariableRegistryUpdateStepDTO) {
	o.UpdateSteps = v
}

// GetProcessGroupId returns the ProcessGroupId field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetProcessGroupId() string {
	if o == nil || IsNil(o.ProcessGroupId) {
		var ret string
		return ret
	}
	return *o.ProcessGroupId
}

// GetProcessGroupIdOk returns a tuple with the ProcessGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetProcessGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessGroupId) {
		return nil, false
	}
	return o.ProcessGroupId, true
}

// HasProcessGroupId returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasProcessGroupId() bool {
	if o != nil && !IsNil(o.ProcessGroupId) {
		return true
	}

	return false
}

// SetProcessGroupId gets a reference to the given string and assigns it to the ProcessGroupId field.
func (o *VariableRegistryUpdateRequestDTO) SetProcessGroupId(v string) {
	o.ProcessGroupId = &v
}

// GetAffectedComponents returns the AffectedComponents field value if set, zero value otherwise.
func (o *VariableRegistryUpdateRequestDTO) GetAffectedComponents() []AffectedComponentEntity {
	if o == nil || IsNil(o.AffectedComponents) {
		var ret []AffectedComponentEntity
		return ret
	}
	return o.AffectedComponents
}

// GetAffectedComponentsOk returns a tuple with the AffectedComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableRegistryUpdateRequestDTO) GetAffectedComponentsOk() ([]AffectedComponentEntity, bool) {
	if o == nil || IsNil(o.AffectedComponents) {
		return nil, false
	}
	return o.AffectedComponents, true
}

// HasAffectedComponents returns a boolean if a field has been set.
func (o *VariableRegistryUpdateRequestDTO) HasAffectedComponents() bool {
	if o != nil && !IsNil(o.AffectedComponents) {
		return true
	}

	return false
}

// SetAffectedComponents gets a reference to the given []AffectedComponentEntity and assigns it to the AffectedComponents field.
func (o *VariableRegistryUpdateRequestDTO) SetAffectedComponents(v []AffectedComponentEntity) {
	o.AffectedComponents = v
}

func (o VariableRegistryUpdateRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableRegistryUpdateRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.SubmissionTime) {
		toSerialize["submissionTime"] = o.SubmissionTime
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
	if !IsNil(o.UpdateSteps) {
		toSerialize["updateSteps"] = o.UpdateSteps
	}
	if !IsNil(o.ProcessGroupId) {
		toSerialize["processGroupId"] = o.ProcessGroupId
	}
	if !IsNil(o.AffectedComponents) {
		toSerialize["affectedComponents"] = o.AffectedComponents
	}
	return toSerialize, nil
}

type NullableVariableRegistryUpdateRequestDTO struct {
	value *VariableRegistryUpdateRequestDTO
	isSet bool
}

func (v NullableVariableRegistryUpdateRequestDTO) Get() *VariableRegistryUpdateRequestDTO {
	return v.value
}

func (v *NullableVariableRegistryUpdateRequestDTO) Set(val *VariableRegistryUpdateRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableRegistryUpdateRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableRegistryUpdateRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableRegistryUpdateRequestDTO(val *VariableRegistryUpdateRequestDTO) *NullableVariableRegistryUpdateRequestDTO {
	return &NullableVariableRegistryUpdateRequestDTO{value: val, isSet: true}
}

func (v NullableVariableRegistryUpdateRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableRegistryUpdateRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


