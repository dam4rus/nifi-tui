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

// checks if the ConnectableComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectableComponent{}

// ConnectableComponent struct for ConnectableComponent
type ConnectableComponent struct {
	// The id of the connectable component.
	Id string `json:"id"`
	// The type of component the connectable is.
	Type string `json:"type"`
	// The id of the group that the connectable component resides in
	GroupId string `json:"groupId"`
	// The name of the connectable component
	Name *string `json:"name,omitempty"`
	// The comments for the connectable component.
	Comments *string `json:"comments,omitempty"`
	// The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component
	InstanceIdentifier *string `json:"instanceIdentifier,omitempty"`
}

// NewConnectableComponent instantiates a new ConnectableComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectableComponent(id string, type_ string, groupId string) *ConnectableComponent {
	this := ConnectableComponent{}
	this.Id = id
	this.Type = type_
	this.GroupId = groupId
	return &this
}

// NewConnectableComponentWithDefaults instantiates a new ConnectableComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectableComponentWithDefaults() *ConnectableComponent {
	this := ConnectableComponent{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectableComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectableComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectableComponent) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ConnectableComponent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectableComponent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectableComponent) SetType(v string) {
	o.Type = v
}

// GetGroupId returns the GroupId field value
func (o *ConnectableComponent) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ConnectableComponent) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ConnectableComponent) SetGroupId(v string) {
	o.GroupId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectableComponent) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectableComponent) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectableComponent) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectableComponent) SetName(v string) {
	o.Name = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ConnectableComponent) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectableComponent) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ConnectableComponent) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ConnectableComponent) SetComments(v string) {
	o.Comments = &v
}

// GetInstanceIdentifier returns the InstanceIdentifier field value if set, zero value otherwise.
func (o *ConnectableComponent) GetInstanceIdentifier() string {
	if o == nil || IsNil(o.InstanceIdentifier) {
		var ret string
		return ret
	}
	return *o.InstanceIdentifier
}

// GetInstanceIdentifierOk returns a tuple with the InstanceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectableComponent) GetInstanceIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceIdentifier) {
		return nil, false
	}
	return o.InstanceIdentifier, true
}

// HasInstanceIdentifier returns a boolean if a field has been set.
func (o *ConnectableComponent) HasInstanceIdentifier() bool {
	if o != nil && !IsNil(o.InstanceIdentifier) {
		return true
	}

	return false
}

// SetInstanceIdentifier gets a reference to the given string and assigns it to the InstanceIdentifier field.
func (o *ConnectableComponent) SetInstanceIdentifier(v string) {
	o.InstanceIdentifier = &v
}

func (o ConnectableComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectableComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["groupId"] = o.GroupId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.InstanceIdentifier) {
		toSerialize["instanceIdentifier"] = o.InstanceIdentifier
	}
	return toSerialize, nil
}

type NullableConnectableComponent struct {
	value *ConnectableComponent
	isSet bool
}

func (v NullableConnectableComponent) Get() *ConnectableComponent {
	return v.value
}

func (v *NullableConnectableComponent) Set(val *ConnectableComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectableComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectableComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectableComponent(val *ConnectableComponent) *NullableConnectableComponent {
	return &NullableConnectableComponent{value: val, isSet: true}
}

func (v NullableConnectableComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectableComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


