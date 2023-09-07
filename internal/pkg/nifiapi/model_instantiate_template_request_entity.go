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

// checks if the InstantiateTemplateRequestEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstantiateTemplateRequestEntity{}

// InstantiateTemplateRequestEntity struct for InstantiateTemplateRequestEntity
type InstantiateTemplateRequestEntity struct {
	// The x coordinate of the origin of the bounding box where the new components will be placed.
	OriginX *float64 `json:"originX,omitempty"`
	// The y coordinate of the origin of the bounding box where the new components will be placed.
	OriginY *float64 `json:"originY,omitempty"`
	// The identifier of the template.
	TemplateId *string `json:"templateId,omitempty"`
	// The encoding version of the flow snippet. If not specified, this is automatically populated by the node receiving the user request. If the snippet is specified, the version will be the latest. If the snippet is not specified, the version will come from the underlying template. These details need to be replicated throughout the cluster to ensure consistency.
	EncodingVersion *string `json:"encodingVersion,omitempty"`
	Snippet *FlowSnippetDTO `json:"snippet,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewInstantiateTemplateRequestEntity instantiates a new InstantiateTemplateRequestEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantiateTemplateRequestEntity() *InstantiateTemplateRequestEntity {
	this := InstantiateTemplateRequestEntity{}
	return &this
}

// NewInstantiateTemplateRequestEntityWithDefaults instantiates a new InstantiateTemplateRequestEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantiateTemplateRequestEntityWithDefaults() *InstantiateTemplateRequestEntity {
	this := InstantiateTemplateRequestEntity{}
	return &this
}

// GetOriginX returns the OriginX field value if set, zero value otherwise.
func (o *InstantiateTemplateRequestEntity) GetOriginX() float64 {
	if o == nil || IsNil(o.OriginX) {
		var ret float64
		return ret
	}
	return *o.OriginX
}

// GetOriginXOk returns a tuple with the OriginX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantiateTemplateRequestEntity) GetOriginXOk() (*float64, bool) {
	if o == nil || IsNil(o.OriginX) {
		return nil, false
	}
	return o.OriginX, true
}

// HasOriginX returns a boolean if a field has been set.
func (o *InstantiateTemplateRequestEntity) HasOriginX() bool {
	if o != nil && !IsNil(o.OriginX) {
		return true
	}

	return false
}

// SetOriginX gets a reference to the given float64 and assigns it to the OriginX field.
func (o *InstantiateTemplateRequestEntity) SetOriginX(v float64) {
	o.OriginX = &v
}

// GetOriginY returns the OriginY field value if set, zero value otherwise.
func (o *InstantiateTemplateRequestEntity) GetOriginY() float64 {
	if o == nil || IsNil(o.OriginY) {
		var ret float64
		return ret
	}
	return *o.OriginY
}

// GetOriginYOk returns a tuple with the OriginY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantiateTemplateRequestEntity) GetOriginYOk() (*float64, bool) {
	if o == nil || IsNil(o.OriginY) {
		return nil, false
	}
	return o.OriginY, true
}

// HasOriginY returns a boolean if a field has been set.
func (o *InstantiateTemplateRequestEntity) HasOriginY() bool {
	if o != nil && !IsNil(o.OriginY) {
		return true
	}

	return false
}

// SetOriginY gets a reference to the given float64 and assigns it to the OriginY field.
func (o *InstantiateTemplateRequestEntity) SetOriginY(v float64) {
	o.OriginY = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *InstantiateTemplateRequestEntity) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantiateTemplateRequestEntity) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *InstantiateTemplateRequestEntity) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *InstantiateTemplateRequestEntity) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetEncodingVersion returns the EncodingVersion field value if set, zero value otherwise.
func (o *InstantiateTemplateRequestEntity) GetEncodingVersion() string {
	if o == nil || IsNil(o.EncodingVersion) {
		var ret string
		return ret
	}
	return *o.EncodingVersion
}

// GetEncodingVersionOk returns a tuple with the EncodingVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantiateTemplateRequestEntity) GetEncodingVersionOk() (*string, bool) {
	if o == nil || IsNil(o.EncodingVersion) {
		return nil, false
	}
	return o.EncodingVersion, true
}

// HasEncodingVersion returns a boolean if a field has been set.
func (o *InstantiateTemplateRequestEntity) HasEncodingVersion() bool {
	if o != nil && !IsNil(o.EncodingVersion) {
		return true
	}

	return false
}

// SetEncodingVersion gets a reference to the given string and assigns it to the EncodingVersion field.
func (o *InstantiateTemplateRequestEntity) SetEncodingVersion(v string) {
	o.EncodingVersion = &v
}

// GetSnippet returns the Snippet field value if set, zero value otherwise.
func (o *InstantiateTemplateRequestEntity) GetSnippet() FlowSnippetDTO {
	if o == nil || IsNil(o.Snippet) {
		var ret FlowSnippetDTO
		return ret
	}
	return *o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantiateTemplateRequestEntity) GetSnippetOk() (*FlowSnippetDTO, bool) {
	if o == nil || IsNil(o.Snippet) {
		return nil, false
	}
	return o.Snippet, true
}

// HasSnippet returns a boolean if a field has been set.
func (o *InstantiateTemplateRequestEntity) HasSnippet() bool {
	if o != nil && !IsNil(o.Snippet) {
		return true
	}

	return false
}

// SetSnippet gets a reference to the given FlowSnippetDTO and assigns it to the Snippet field.
func (o *InstantiateTemplateRequestEntity) SetSnippet(v FlowSnippetDTO) {
	o.Snippet = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *InstantiateTemplateRequestEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantiateTemplateRequestEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.DisconnectedNodeAcknowledged) {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *InstantiateTemplateRequestEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && !IsNil(o.DisconnectedNodeAcknowledged) {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *InstantiateTemplateRequestEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o InstantiateTemplateRequestEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstantiateTemplateRequestEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginX) {
		toSerialize["originX"] = o.OriginX
	}
	if !IsNil(o.OriginY) {
		toSerialize["originY"] = o.OriginY
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.EncodingVersion) {
		toSerialize["encodingVersion"] = o.EncodingVersion
	}
	if !IsNil(o.Snippet) {
		toSerialize["snippet"] = o.Snippet
	}
	if !IsNil(o.DisconnectedNodeAcknowledged) {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return toSerialize, nil
}

type NullableInstantiateTemplateRequestEntity struct {
	value *InstantiateTemplateRequestEntity
	isSet bool
}

func (v NullableInstantiateTemplateRequestEntity) Get() *InstantiateTemplateRequestEntity {
	return v.value
}

func (v *NullableInstantiateTemplateRequestEntity) Set(val *InstantiateTemplateRequestEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantiateTemplateRequestEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantiateTemplateRequestEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantiateTemplateRequestEntity(val *InstantiateTemplateRequestEntity) *NullableInstantiateTemplateRequestEntity {
	return &NullableInstantiateTemplateRequestEntity{value: val, isSet: true}
}

func (v NullableInstantiateTemplateRequestEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantiateTemplateRequestEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

