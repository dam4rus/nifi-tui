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

// checks if the DropRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DropRequestDTO{}

// DropRequestDTO struct for DropRequestDTO
type DropRequestDTO struct {
	// The id for this drop request.
	Id *string `json:"id,omitempty"`
	// The URI for future requests to this drop request.
	Uri *string `json:"uri,omitempty"`
	// The timestamp when the query was submitted.
	SubmissionTime *string `json:"submissionTime,omitempty"`
	// The last time this drop request was updated.
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// The current percent complete.
	PercentCompleted *int32 `json:"percentCompleted,omitempty"`
	// Whether the query has finished.
	Finished *bool `json:"finished,omitempty"`
	// The reason, if any, that this drop request failed.
	FailureReason *string `json:"failureReason,omitempty"`
	// The number of flow files currently queued.
	CurrentCount *int32 `json:"currentCount,omitempty"`
	// The size of flow files currently queued in bytes.
	CurrentSize *int64 `json:"currentSize,omitempty"`
	// The count and size of flow files currently queued.
	Current *string `json:"current,omitempty"`
	// The number of flow files to be dropped as a result of this request.
	OriginalCount *int32 `json:"originalCount,omitempty"`
	// The size of flow files to be dropped as a result of this request in bytes.
	OriginalSize *int64 `json:"originalSize,omitempty"`
	// The count and size of flow files to be dropped as a result of this request.
	Original *string `json:"original,omitempty"`
	// The number of flow files that have been dropped thus far.
	DroppedCount *int32 `json:"droppedCount,omitempty"`
	// The size of flow files that have been dropped thus far in bytes.
	DroppedSize *int64 `json:"droppedSize,omitempty"`
	// The count and size of flow files that have been dropped thus far.
	Dropped *string `json:"dropped,omitempty"`
	// The current state of the drop request.
	State *string `json:"state,omitempty"`
}

// NewDropRequestDTO instantiates a new DropRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropRequestDTO() *DropRequestDTO {
	this := DropRequestDTO{}
	return &this
}

// NewDropRequestDTOWithDefaults instantiates a new DropRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropRequestDTOWithDefaults() *DropRequestDTO {
	this := DropRequestDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DropRequestDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DropRequestDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DropRequestDTO) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *DropRequestDTO) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *DropRequestDTO) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *DropRequestDTO) SetUri(v string) {
	o.Uri = &v
}

// GetSubmissionTime returns the SubmissionTime field value if set, zero value otherwise.
func (o *DropRequestDTO) GetSubmissionTime() string {
	if o == nil || IsNil(o.SubmissionTime) {
		var ret string
		return ret
	}
	return *o.SubmissionTime
}

// GetSubmissionTimeOk returns a tuple with the SubmissionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetSubmissionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SubmissionTime) {
		return nil, false
	}
	return o.SubmissionTime, true
}

// HasSubmissionTime returns a boolean if a field has been set.
func (o *DropRequestDTO) HasSubmissionTime() bool {
	if o != nil && !IsNil(o.SubmissionTime) {
		return true
	}

	return false
}

// SetSubmissionTime gets a reference to the given string and assigns it to the SubmissionTime field.
func (o *DropRequestDTO) SetSubmissionTime(v string) {
	o.SubmissionTime = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DropRequestDTO) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DropRequestDTO) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *DropRequestDTO) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetPercentCompleted returns the PercentCompleted field value if set, zero value otherwise.
func (o *DropRequestDTO) GetPercentCompleted() int32 {
	if o == nil || IsNil(o.PercentCompleted) {
		var ret int32
		return ret
	}
	return *o.PercentCompleted
}

// GetPercentCompletedOk returns a tuple with the PercentCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetPercentCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentCompleted) {
		return nil, false
	}
	return o.PercentCompleted, true
}

// HasPercentCompleted returns a boolean if a field has been set.
func (o *DropRequestDTO) HasPercentCompleted() bool {
	if o != nil && !IsNil(o.PercentCompleted) {
		return true
	}

	return false
}

// SetPercentCompleted gets a reference to the given int32 and assigns it to the PercentCompleted field.
func (o *DropRequestDTO) SetPercentCompleted(v int32) {
	o.PercentCompleted = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *DropRequestDTO) GetFinished() bool {
	if o == nil || IsNil(o.Finished) {
		var ret bool
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetFinishedOk() (*bool, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *DropRequestDTO) HasFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given bool and assigns it to the Finished field.
func (o *DropRequestDTO) SetFinished(v bool) {
	o.Finished = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *DropRequestDTO) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *DropRequestDTO) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *DropRequestDTO) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetCurrentCount returns the CurrentCount field value if set, zero value otherwise.
func (o *DropRequestDTO) GetCurrentCount() int32 {
	if o == nil || IsNil(o.CurrentCount) {
		var ret int32
		return ret
	}
	return *o.CurrentCount
}

// GetCurrentCountOk returns a tuple with the CurrentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetCurrentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentCount) {
		return nil, false
	}
	return o.CurrentCount, true
}

// HasCurrentCount returns a boolean if a field has been set.
func (o *DropRequestDTO) HasCurrentCount() bool {
	if o != nil && !IsNil(o.CurrentCount) {
		return true
	}

	return false
}

// SetCurrentCount gets a reference to the given int32 and assigns it to the CurrentCount field.
func (o *DropRequestDTO) SetCurrentCount(v int32) {
	o.CurrentCount = &v
}

// GetCurrentSize returns the CurrentSize field value if set, zero value otherwise.
func (o *DropRequestDTO) GetCurrentSize() int64 {
	if o == nil || IsNil(o.CurrentSize) {
		var ret int64
		return ret
	}
	return *o.CurrentSize
}

// GetCurrentSizeOk returns a tuple with the CurrentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetCurrentSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentSize) {
		return nil, false
	}
	return o.CurrentSize, true
}

// HasCurrentSize returns a boolean if a field has been set.
func (o *DropRequestDTO) HasCurrentSize() bool {
	if o != nil && !IsNil(o.CurrentSize) {
		return true
	}

	return false
}

// SetCurrentSize gets a reference to the given int64 and assigns it to the CurrentSize field.
func (o *DropRequestDTO) SetCurrentSize(v int64) {
	o.CurrentSize = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *DropRequestDTO) GetCurrent() string {
	if o == nil || IsNil(o.Current) {
		var ret string
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetCurrentOk() (*string, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *DropRequestDTO) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given string and assigns it to the Current field.
func (o *DropRequestDTO) SetCurrent(v string) {
	o.Current = &v
}

// GetOriginalCount returns the OriginalCount field value if set, zero value otherwise.
func (o *DropRequestDTO) GetOriginalCount() int32 {
	if o == nil || IsNil(o.OriginalCount) {
		var ret int32
		return ret
	}
	return *o.OriginalCount
}

// GetOriginalCountOk returns a tuple with the OriginalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetOriginalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalCount) {
		return nil, false
	}
	return o.OriginalCount, true
}

// HasOriginalCount returns a boolean if a field has been set.
func (o *DropRequestDTO) HasOriginalCount() bool {
	if o != nil && !IsNil(o.OriginalCount) {
		return true
	}

	return false
}

// SetOriginalCount gets a reference to the given int32 and assigns it to the OriginalCount field.
func (o *DropRequestDTO) SetOriginalCount(v int32) {
	o.OriginalCount = &v
}

// GetOriginalSize returns the OriginalSize field value if set, zero value otherwise.
func (o *DropRequestDTO) GetOriginalSize() int64 {
	if o == nil || IsNil(o.OriginalSize) {
		var ret int64
		return ret
	}
	return *o.OriginalSize
}

// GetOriginalSizeOk returns a tuple with the OriginalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetOriginalSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginalSize) {
		return nil, false
	}
	return o.OriginalSize, true
}

// HasOriginalSize returns a boolean if a field has been set.
func (o *DropRequestDTO) HasOriginalSize() bool {
	if o != nil && !IsNil(o.OriginalSize) {
		return true
	}

	return false
}

// SetOriginalSize gets a reference to the given int64 and assigns it to the OriginalSize field.
func (o *DropRequestDTO) SetOriginalSize(v int64) {
	o.OriginalSize = &v
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *DropRequestDTO) GetOriginal() string {
	if o == nil || IsNil(o.Original) {
		var ret string
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetOriginalOk() (*string, bool) {
	if o == nil || IsNil(o.Original) {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *DropRequestDTO) HasOriginal() bool {
	if o != nil && !IsNil(o.Original) {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given string and assigns it to the Original field.
func (o *DropRequestDTO) SetOriginal(v string) {
	o.Original = &v
}

// GetDroppedCount returns the DroppedCount field value if set, zero value otherwise.
func (o *DropRequestDTO) GetDroppedCount() int32 {
	if o == nil || IsNil(o.DroppedCount) {
		var ret int32
		return ret
	}
	return *o.DroppedCount
}

// GetDroppedCountOk returns a tuple with the DroppedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetDroppedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DroppedCount) {
		return nil, false
	}
	return o.DroppedCount, true
}

// HasDroppedCount returns a boolean if a field has been set.
func (o *DropRequestDTO) HasDroppedCount() bool {
	if o != nil && !IsNil(o.DroppedCount) {
		return true
	}

	return false
}

// SetDroppedCount gets a reference to the given int32 and assigns it to the DroppedCount field.
func (o *DropRequestDTO) SetDroppedCount(v int32) {
	o.DroppedCount = &v
}

// GetDroppedSize returns the DroppedSize field value if set, zero value otherwise.
func (o *DropRequestDTO) GetDroppedSize() int64 {
	if o == nil || IsNil(o.DroppedSize) {
		var ret int64
		return ret
	}
	return *o.DroppedSize
}

// GetDroppedSizeOk returns a tuple with the DroppedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetDroppedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.DroppedSize) {
		return nil, false
	}
	return o.DroppedSize, true
}

// HasDroppedSize returns a boolean if a field has been set.
func (o *DropRequestDTO) HasDroppedSize() bool {
	if o != nil && !IsNil(o.DroppedSize) {
		return true
	}

	return false
}

// SetDroppedSize gets a reference to the given int64 and assigns it to the DroppedSize field.
func (o *DropRequestDTO) SetDroppedSize(v int64) {
	o.DroppedSize = &v
}

// GetDropped returns the Dropped field value if set, zero value otherwise.
func (o *DropRequestDTO) GetDropped() string {
	if o == nil || IsNil(o.Dropped) {
		var ret string
		return ret
	}
	return *o.Dropped
}

// GetDroppedOk returns a tuple with the Dropped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetDroppedOk() (*string, bool) {
	if o == nil || IsNil(o.Dropped) {
		return nil, false
	}
	return o.Dropped, true
}

// HasDropped returns a boolean if a field has been set.
func (o *DropRequestDTO) HasDropped() bool {
	if o != nil && !IsNil(o.Dropped) {
		return true
	}

	return false
}

// SetDropped gets a reference to the given string and assigns it to the Dropped field.
func (o *DropRequestDTO) SetDropped(v string) {
	o.Dropped = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DropRequestDTO) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropRequestDTO) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DropRequestDTO) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DropRequestDTO) SetState(v string) {
	o.State = &v
}

func (o DropRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DropRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.PercentCompleted) {
		toSerialize["percentCompleted"] = o.PercentCompleted
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	if !IsNil(o.CurrentCount) {
		toSerialize["currentCount"] = o.CurrentCount
	}
	if !IsNil(o.CurrentSize) {
		toSerialize["currentSize"] = o.CurrentSize
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.OriginalCount) {
		toSerialize["originalCount"] = o.OriginalCount
	}
	if !IsNil(o.OriginalSize) {
		toSerialize["originalSize"] = o.OriginalSize
	}
	if !IsNil(o.Original) {
		toSerialize["original"] = o.Original
	}
	if !IsNil(o.DroppedCount) {
		toSerialize["droppedCount"] = o.DroppedCount
	}
	if !IsNil(o.DroppedSize) {
		toSerialize["droppedSize"] = o.DroppedSize
	}
	if !IsNil(o.Dropped) {
		toSerialize["dropped"] = o.Dropped
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableDropRequestDTO struct {
	value *DropRequestDTO
	isSet bool
}

func (v NullableDropRequestDTO) Get() *DropRequestDTO {
	return v.value
}

func (v *NullableDropRequestDTO) Set(val *DropRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDropRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDropRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropRequestDTO(val *DropRequestDTO) *NullableDropRequestDTO {
	return &NullableDropRequestDTO{value: val, isSet: true}
}

func (v NullableDropRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


