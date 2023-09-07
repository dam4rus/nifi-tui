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

// checks if the VerifyConfigRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyConfigRequestDTO{}

// VerifyConfigRequestDTO struct for VerifyConfigRequestDTO
type VerifyConfigRequestDTO struct {
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
	UpdateSteps []VerifyConfigUpdateStepDTO `json:"updateSteps,omitempty"`
	// The ID of the component whose configuration was verified
	ComponentId *string `json:"componentId,omitempty"`
	// The configured component properties
	Properties *map[string]string `json:"properties,omitempty"`
	// FlowFile Attributes that should be used to evaluate Expression Language for resolving property values
	Attributes *map[string]string `json:"attributes,omitempty"`
	// The Results of the verification
	Results []ConfigVerificationResultDTO `json:"results,omitempty"`
}

// NewVerifyConfigRequestDTO instantiates a new VerifyConfigRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyConfigRequestDTO() *VerifyConfigRequestDTO {
	this := VerifyConfigRequestDTO{}
	return &this
}

// NewVerifyConfigRequestDTOWithDefaults instantiates a new VerifyConfigRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyConfigRequestDTOWithDefaults() *VerifyConfigRequestDTO {
	this := VerifyConfigRequestDTO{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *VerifyConfigRequestDTO) SetRequestId(v string) {
	o.RequestId = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *VerifyConfigRequestDTO) SetUri(v string) {
	o.Uri = &v
}

// GetSubmissionTime returns the SubmissionTime field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetSubmissionTime() time.Time {
	if o == nil || IsNil(o.SubmissionTime) {
		var ret time.Time
		return ret
	}
	return *o.SubmissionTime
}

// GetSubmissionTimeOk returns a tuple with the SubmissionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetSubmissionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmissionTime) {
		return nil, false
	}
	return o.SubmissionTime, true
}

// HasSubmissionTime returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasSubmissionTime() bool {
	if o != nil && !IsNil(o.SubmissionTime) {
		return true
	}

	return false
}

// SetSubmissionTime gets a reference to the given time.Time and assigns it to the SubmissionTime field.
func (o *VerifyConfigRequestDTO) SetSubmissionTime(v time.Time) {
	o.SubmissionTime = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *VerifyConfigRequestDTO) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *VerifyConfigRequestDTO) SetComplete(v bool) {
	o.Complete = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *VerifyConfigRequestDTO) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetPercentCompleted returns the PercentCompleted field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetPercentCompleted() int32 {
	if o == nil || IsNil(o.PercentCompleted) {
		var ret int32
		return ret
	}
	return *o.PercentCompleted
}

// GetPercentCompletedOk returns a tuple with the PercentCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetPercentCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentCompleted) {
		return nil, false
	}
	return o.PercentCompleted, true
}

// HasPercentCompleted returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasPercentCompleted() bool {
	if o != nil && !IsNil(o.PercentCompleted) {
		return true
	}

	return false
}

// SetPercentCompleted gets a reference to the given int32 and assigns it to the PercentCompleted field.
func (o *VerifyConfigRequestDTO) SetPercentCompleted(v int32) {
	o.PercentCompleted = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VerifyConfigRequestDTO) SetState(v string) {
	o.State = &v
}

// GetUpdateSteps returns the UpdateSteps field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetUpdateSteps() []VerifyConfigUpdateStepDTO {
	if o == nil || IsNil(o.UpdateSteps) {
		var ret []VerifyConfigUpdateStepDTO
		return ret
	}
	return o.UpdateSteps
}

// GetUpdateStepsOk returns a tuple with the UpdateSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetUpdateStepsOk() ([]VerifyConfigUpdateStepDTO, bool) {
	if o == nil || IsNil(o.UpdateSteps) {
		return nil, false
	}
	return o.UpdateSteps, true
}

// HasUpdateSteps returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasUpdateSteps() bool {
	if o != nil && !IsNil(o.UpdateSteps) {
		return true
	}

	return false
}

// SetUpdateSteps gets a reference to the given []VerifyConfigUpdateStepDTO and assigns it to the UpdateSteps field.
func (o *VerifyConfigRequestDTO) SetUpdateSteps(v []VerifyConfigUpdateStepDTO) {
	o.UpdateSteps = v
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetComponentId() string {
	if o == nil || IsNil(o.ComponentId) {
		var ret string
		return ret
	}
	return *o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetComponentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentId) {
		return nil, false
	}
	return o.ComponentId, true
}

// HasComponentId returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasComponentId() bool {
	if o != nil && !IsNil(o.ComponentId) {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given string and assigns it to the ComponentId field.
func (o *VerifyConfigRequestDTO) SetComponentId(v string) {
	o.ComponentId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *VerifyConfigRequestDTO) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *VerifyConfigRequestDTO) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *VerifyConfigRequestDTO) GetResults() []ConfigVerificationResultDTO {
	if o == nil || IsNil(o.Results) {
		var ret []ConfigVerificationResultDTO
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyConfigRequestDTO) GetResultsOk() ([]ConfigVerificationResultDTO, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *VerifyConfigRequestDTO) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ConfigVerificationResultDTO and assigns it to the Results field.
func (o *VerifyConfigRequestDTO) SetResults(v []ConfigVerificationResultDTO) {
	o.Results = v
}

func (o VerifyConfigRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyConfigRequestDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ComponentId) {
		toSerialize["componentId"] = o.ComponentId
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableVerifyConfigRequestDTO struct {
	value *VerifyConfigRequestDTO
	isSet bool
}

func (v NullableVerifyConfigRequestDTO) Get() *VerifyConfigRequestDTO {
	return v.value
}

func (v *NullableVerifyConfigRequestDTO) Set(val *VerifyConfigRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyConfigRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyConfigRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyConfigRequestDTO(val *VerifyConfigRequestDTO) *NullableVerifyConfigRequestDTO {
	return &NullableVerifyConfigRequestDTO{value: val, isSet: true}
}

func (v NullableVerifyConfigRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyConfigRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


