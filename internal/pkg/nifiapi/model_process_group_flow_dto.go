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

// checks if the ProcessGroupFlowDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessGroupFlowDTO{}

// ProcessGroupFlowDTO struct for ProcessGroupFlowDTO
type ProcessGroupFlowDTO struct {
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri *string `json:"uri,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId *string `json:"parentGroupId,omitempty"`
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	Breadcrumb *FlowBreadcrumbEntity `json:"breadcrumb,omitempty"`
	Flow *FlowDTO `json:"flow,omitempty"`
	// The time the flow for the process group was last refreshed.
	LastRefreshed *string `json:"lastRefreshed,omitempty"`
}

// NewProcessGroupFlowDTO instantiates a new ProcessGroupFlowDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupFlowDTO() *ProcessGroupFlowDTO {
	this := ProcessGroupFlowDTO{}
	return &this
}

// NewProcessGroupFlowDTOWithDefaults instantiates a new ProcessGroupFlowDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupFlowDTOWithDefaults() *ProcessGroupFlowDTO {
	this := ProcessGroupFlowDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProcessGroupFlowDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupFlowDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProcessGroupFlowDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProcessGroupFlowDTO) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ProcessGroupFlowDTO) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupFlowDTO) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ProcessGroupFlowDTO) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ProcessGroupFlowDTO) SetUri(v string) {
	o.Uri = &v
}

// GetParentGroupId returns the ParentGroupId field value if set, zero value otherwise.
func (o *ProcessGroupFlowDTO) GetParentGroupId() string {
	if o == nil || IsNil(o.ParentGroupId) {
		var ret string
		return ret
	}
	return *o.ParentGroupId
}

// GetParentGroupIdOk returns a tuple with the ParentGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupFlowDTO) GetParentGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentGroupId) {
		return nil, false
	}
	return o.ParentGroupId, true
}

// HasParentGroupId returns a boolean if a field has been set.
func (o *ProcessGroupFlowDTO) HasParentGroupId() bool {
	if o != nil && !IsNil(o.ParentGroupId) {
		return true
	}

	return false
}

// SetParentGroupId gets a reference to the given string and assigns it to the ParentGroupId field.
func (o *ProcessGroupFlowDTO) SetParentGroupId(v string) {
	o.ParentGroupId = &v
}

// GetParameterContext returns the ParameterContext field value if set, zero value otherwise.
func (o *ProcessGroupFlowDTO) GetParameterContext() ParameterContextReferenceEntity {
	if o == nil || IsNil(o.ParameterContext) {
		var ret ParameterContextReferenceEntity
		return ret
	}
	return *o.ParameterContext
}

// GetParameterContextOk returns a tuple with the ParameterContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupFlowDTO) GetParameterContextOk() (*ParameterContextReferenceEntity, bool) {
	if o == nil || IsNil(o.ParameterContext) {
		return nil, false
	}
	return o.ParameterContext, true
}

// HasParameterContext returns a boolean if a field has been set.
func (o *ProcessGroupFlowDTO) HasParameterContext() bool {
	if o != nil && !IsNil(o.ParameterContext) {
		return true
	}

	return false
}

// SetParameterContext gets a reference to the given ParameterContextReferenceEntity and assigns it to the ParameterContext field.
func (o *ProcessGroupFlowDTO) SetParameterContext(v ParameterContextReferenceEntity) {
	o.ParameterContext = &v
}

// GetBreadcrumb returns the Breadcrumb field value if set, zero value otherwise.
func (o *ProcessGroupFlowDTO) GetBreadcrumb() FlowBreadcrumbEntity {
	if o == nil || IsNil(o.Breadcrumb) {
		var ret FlowBreadcrumbEntity
		return ret
	}
	return *o.Breadcrumb
}

// GetBreadcrumbOk returns a tuple with the Breadcrumb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupFlowDTO) GetBreadcrumbOk() (*FlowBreadcrumbEntity, bool) {
	if o == nil || IsNil(o.Breadcrumb) {
		return nil, false
	}
	return o.Breadcrumb, true
}

// HasBreadcrumb returns a boolean if a field has been set.
func (o *ProcessGroupFlowDTO) HasBreadcrumb() bool {
	if o != nil && !IsNil(o.Breadcrumb) {
		return true
	}

	return false
}

// SetBreadcrumb gets a reference to the given FlowBreadcrumbEntity and assigns it to the Breadcrumb field.
func (o *ProcessGroupFlowDTO) SetBreadcrumb(v FlowBreadcrumbEntity) {
	o.Breadcrumb = &v
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *ProcessGroupFlowDTO) GetFlow() FlowDTO {
	if o == nil || IsNil(o.Flow) {
		var ret FlowDTO
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupFlowDTO) GetFlowOk() (*FlowDTO, bool) {
	if o == nil || IsNil(o.Flow) {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *ProcessGroupFlowDTO) HasFlow() bool {
	if o != nil && !IsNil(o.Flow) {
		return true
	}

	return false
}

// SetFlow gets a reference to the given FlowDTO and assigns it to the Flow field.
func (o *ProcessGroupFlowDTO) SetFlow(v FlowDTO) {
	o.Flow = &v
}

// GetLastRefreshed returns the LastRefreshed field value if set, zero value otherwise.
func (o *ProcessGroupFlowDTO) GetLastRefreshed() string {
	if o == nil || IsNil(o.LastRefreshed) {
		var ret string
		return ret
	}
	return *o.LastRefreshed
}

// GetLastRefreshedOk returns a tuple with the LastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupFlowDTO) GetLastRefreshedOk() (*string, bool) {
	if o == nil || IsNil(o.LastRefreshed) {
		return nil, false
	}
	return o.LastRefreshed, true
}

// HasLastRefreshed returns a boolean if a field has been set.
func (o *ProcessGroupFlowDTO) HasLastRefreshed() bool {
	if o != nil && !IsNil(o.LastRefreshed) {
		return true
	}

	return false
}

// SetLastRefreshed gets a reference to the given string and assigns it to the LastRefreshed field.
func (o *ProcessGroupFlowDTO) SetLastRefreshed(v string) {
	o.LastRefreshed = &v
}

func (o ProcessGroupFlowDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessGroupFlowDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.ParentGroupId) {
		toSerialize["parentGroupId"] = o.ParentGroupId
	}
	if !IsNil(o.ParameterContext) {
		toSerialize["parameterContext"] = o.ParameterContext
	}
	if !IsNil(o.Breadcrumb) {
		toSerialize["breadcrumb"] = o.Breadcrumb
	}
	if !IsNil(o.Flow) {
		toSerialize["flow"] = o.Flow
	}
	if !IsNil(o.LastRefreshed) {
		toSerialize["lastRefreshed"] = o.LastRefreshed
	}
	return toSerialize, nil
}

type NullableProcessGroupFlowDTO struct {
	value *ProcessGroupFlowDTO
	isSet bool
}

func (v NullableProcessGroupFlowDTO) Get() *ProcessGroupFlowDTO {
	return v.value
}

func (v *NullableProcessGroupFlowDTO) Set(val *ProcessGroupFlowDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupFlowDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupFlowDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupFlowDTO(val *ProcessGroupFlowDTO) *NullableProcessGroupFlowDTO {
	return &NullableProcessGroupFlowDTO{value: val, isSet: true}
}

func (v NullableProcessGroupFlowDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupFlowDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


