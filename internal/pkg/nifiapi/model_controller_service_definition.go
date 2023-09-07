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

// checks if the ControllerServiceDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControllerServiceDefinition{}

// ControllerServiceDefinition struct for ControllerServiceDefinition
type ControllerServiceDefinition struct {
	// The group name of the bundle that provides the referenced type.
	Group *string `json:"group,omitempty"`
	// The artifact name of the bundle that provides the referenced type.
	Artifact *string `json:"artifact,omitempty"`
	// The version of the bundle that provides the referenced type.
	Version *string `json:"version,omitempty"`
	// The fully-qualified class type
	Type string `json:"type"`
	// The description of the type.
	TypeDescription *string `json:"typeDescription,omitempty"`
	BuildInfo *BuildInfo `json:"buildInfo,omitempty"`
	// If this type represents a provider for an interface, this lists the APIs it implements
	ProvidedApiImplementations []DefinedType `json:"providedApiImplementations,omitempty"`
	// The tags associated with this type
	Tags []string `json:"tags,omitempty"`
	// The names of other component types that may be related
	SeeAlso []string `json:"seeAlso,omitempty"`
	// Whether or not the component has been deprecated
	Deprecated *bool `json:"deprecated,omitempty"`
	// If this component has been deprecated, this optional field can be used to provide an explanation
	DeprecationReason *string `json:"deprecationReason,omitempty"`
	// If this component has been deprecated, this optional field provides alternatives to use
	DeprecationAlternatives []string `json:"deprecationAlternatives,omitempty"`
	// Whether or not the component has a general restriction
	Restricted *bool `json:"restricted,omitempty"`
	// An optional description of the general restriction
	RestrictedExplanation *string `json:"restrictedExplanation,omitempty"`
	// Explicit restrictions that indicate a require permission to use the component
	ExplicitRestrictions []Restriction `json:"explicitRestrictions,omitempty"`
	Stateful *Stateful `json:"stateful,omitempty"`
	// The system resource considerations for the given component
	SystemResourceConsiderations []SystemResourceConsideration `json:"systemResourceConsiderations,omitempty"`
	// Indicates if the component has additional details documentation
	AdditionalDetails *bool `json:"additionalDetails,omitempty"`
	// Descriptions of configuration properties applicable to this component.
	PropertyDescriptors *map[string]PropertyDescriptor `json:"propertyDescriptors,omitempty"`
	// Whether or not this component makes use of dynamic (user-set) properties.
	SupportsDynamicProperties *bool `json:"supportsDynamicProperties,omitempty"`
	// Describes the dynamic properties supported by this component
	DynamicProperties []DynamicProperty `json:"dynamicProperties,omitempty"`
}

// NewControllerServiceDefinition instantiates a new ControllerServiceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerServiceDefinition(type_ string) *ControllerServiceDefinition {
	this := ControllerServiceDefinition{}
	this.Type = type_
	return &this
}

// NewControllerServiceDefinitionWithDefaults instantiates a new ControllerServiceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerServiceDefinitionWithDefaults() *ControllerServiceDefinition {
	this := ControllerServiceDefinition{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ControllerServiceDefinition) SetGroup(v string) {
	o.Group = &v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *ControllerServiceDefinition) SetArtifact(v string) {
	o.Artifact = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ControllerServiceDefinition) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value
func (o *ControllerServiceDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ControllerServiceDefinition) SetType(v string) {
	o.Type = v
}

// GetTypeDescription returns the TypeDescription field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetTypeDescription() string {
	if o == nil || IsNil(o.TypeDescription) {
		var ret string
		return ret
	}
	return *o.TypeDescription
}

// GetTypeDescriptionOk returns a tuple with the TypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetTypeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.TypeDescription) {
		return nil, false
	}
	return o.TypeDescription, true
}

// HasTypeDescription returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasTypeDescription() bool {
	if o != nil && !IsNil(o.TypeDescription) {
		return true
	}

	return false
}

// SetTypeDescription gets a reference to the given string and assigns it to the TypeDescription field.
func (o *ControllerServiceDefinition) SetTypeDescription(v string) {
	o.TypeDescription = &v
}

// GetBuildInfo returns the BuildInfo field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetBuildInfo() BuildInfo {
	if o == nil || IsNil(o.BuildInfo) {
		var ret BuildInfo
		return ret
	}
	return *o.BuildInfo
}

// GetBuildInfoOk returns a tuple with the BuildInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetBuildInfoOk() (*BuildInfo, bool) {
	if o == nil || IsNil(o.BuildInfo) {
		return nil, false
	}
	return o.BuildInfo, true
}

// HasBuildInfo returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasBuildInfo() bool {
	if o != nil && !IsNil(o.BuildInfo) {
		return true
	}

	return false
}

// SetBuildInfo gets a reference to the given BuildInfo and assigns it to the BuildInfo field.
func (o *ControllerServiceDefinition) SetBuildInfo(v BuildInfo) {
	o.BuildInfo = &v
}

// GetProvidedApiImplementations returns the ProvidedApiImplementations field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetProvidedApiImplementations() []DefinedType {
	if o == nil || IsNil(o.ProvidedApiImplementations) {
		var ret []DefinedType
		return ret
	}
	return o.ProvidedApiImplementations
}

// GetProvidedApiImplementationsOk returns a tuple with the ProvidedApiImplementations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetProvidedApiImplementationsOk() ([]DefinedType, bool) {
	if o == nil || IsNil(o.ProvidedApiImplementations) {
		return nil, false
	}
	return o.ProvidedApiImplementations, true
}

// HasProvidedApiImplementations returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasProvidedApiImplementations() bool {
	if o != nil && !IsNil(o.ProvidedApiImplementations) {
		return true
	}

	return false
}

// SetProvidedApiImplementations gets a reference to the given []DefinedType and assigns it to the ProvidedApiImplementations field.
func (o *ControllerServiceDefinition) SetProvidedApiImplementations(v []DefinedType) {
	o.ProvidedApiImplementations = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ControllerServiceDefinition) SetTags(v []string) {
	o.Tags = v
}

// GetSeeAlso returns the SeeAlso field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetSeeAlso() []string {
	if o == nil || IsNil(o.SeeAlso) {
		var ret []string
		return ret
	}
	return o.SeeAlso
}

// GetSeeAlsoOk returns a tuple with the SeeAlso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetSeeAlsoOk() ([]string, bool) {
	if o == nil || IsNil(o.SeeAlso) {
		return nil, false
	}
	return o.SeeAlso, true
}

// HasSeeAlso returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasSeeAlso() bool {
	if o != nil && !IsNil(o.SeeAlso) {
		return true
	}

	return false
}

// SetSeeAlso gets a reference to the given []string and assigns it to the SeeAlso field.
func (o *ControllerServiceDefinition) SetSeeAlso(v []string) {
	o.SeeAlso = v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *ControllerServiceDefinition) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDeprecationReason returns the DeprecationReason field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetDeprecationReason() string {
	if o == nil || IsNil(o.DeprecationReason) {
		var ret string
		return ret
	}
	return *o.DeprecationReason
}

// GetDeprecationReasonOk returns a tuple with the DeprecationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetDeprecationReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DeprecationReason) {
		return nil, false
	}
	return o.DeprecationReason, true
}

// HasDeprecationReason returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasDeprecationReason() bool {
	if o != nil && !IsNil(o.DeprecationReason) {
		return true
	}

	return false
}

// SetDeprecationReason gets a reference to the given string and assigns it to the DeprecationReason field.
func (o *ControllerServiceDefinition) SetDeprecationReason(v string) {
	o.DeprecationReason = &v
}

// GetDeprecationAlternatives returns the DeprecationAlternatives field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetDeprecationAlternatives() []string {
	if o == nil || IsNil(o.DeprecationAlternatives) {
		var ret []string
		return ret
	}
	return o.DeprecationAlternatives
}

// GetDeprecationAlternativesOk returns a tuple with the DeprecationAlternatives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetDeprecationAlternativesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeprecationAlternatives) {
		return nil, false
	}
	return o.DeprecationAlternatives, true
}

// HasDeprecationAlternatives returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasDeprecationAlternatives() bool {
	if o != nil && !IsNil(o.DeprecationAlternatives) {
		return true
	}

	return false
}

// SetDeprecationAlternatives gets a reference to the given []string and assigns it to the DeprecationAlternatives field.
func (o *ControllerServiceDefinition) SetDeprecationAlternatives(v []string) {
	o.DeprecationAlternatives = v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetRestricted() bool {
	if o == nil || IsNil(o.Restricted) {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.Restricted) {
		return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasRestricted() bool {
	if o != nil && !IsNil(o.Restricted) {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *ControllerServiceDefinition) SetRestricted(v bool) {
	o.Restricted = &v
}

// GetRestrictedExplanation returns the RestrictedExplanation field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetRestrictedExplanation() string {
	if o == nil || IsNil(o.RestrictedExplanation) {
		var ret string
		return ret
	}
	return *o.RestrictedExplanation
}

// GetRestrictedExplanationOk returns a tuple with the RestrictedExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetRestrictedExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.RestrictedExplanation) {
		return nil, false
	}
	return o.RestrictedExplanation, true
}

// HasRestrictedExplanation returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasRestrictedExplanation() bool {
	if o != nil && !IsNil(o.RestrictedExplanation) {
		return true
	}

	return false
}

// SetRestrictedExplanation gets a reference to the given string and assigns it to the RestrictedExplanation field.
func (o *ControllerServiceDefinition) SetRestrictedExplanation(v string) {
	o.RestrictedExplanation = &v
}

// GetExplicitRestrictions returns the ExplicitRestrictions field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetExplicitRestrictions() []Restriction {
	if o == nil || IsNil(o.ExplicitRestrictions) {
		var ret []Restriction
		return ret
	}
	return o.ExplicitRestrictions
}

// GetExplicitRestrictionsOk returns a tuple with the ExplicitRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetExplicitRestrictionsOk() ([]Restriction, bool) {
	if o == nil || IsNil(o.ExplicitRestrictions) {
		return nil, false
	}
	return o.ExplicitRestrictions, true
}

// HasExplicitRestrictions returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasExplicitRestrictions() bool {
	if o != nil && !IsNil(o.ExplicitRestrictions) {
		return true
	}

	return false
}

// SetExplicitRestrictions gets a reference to the given []Restriction and assigns it to the ExplicitRestrictions field.
func (o *ControllerServiceDefinition) SetExplicitRestrictions(v []Restriction) {
	o.ExplicitRestrictions = v
}

// GetStateful returns the Stateful field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetStateful() Stateful {
	if o == nil || IsNil(o.Stateful) {
		var ret Stateful
		return ret
	}
	return *o.Stateful
}

// GetStatefulOk returns a tuple with the Stateful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetStatefulOk() (*Stateful, bool) {
	if o == nil || IsNil(o.Stateful) {
		return nil, false
	}
	return o.Stateful, true
}

// HasStateful returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasStateful() bool {
	if o != nil && !IsNil(o.Stateful) {
		return true
	}

	return false
}

// SetStateful gets a reference to the given Stateful and assigns it to the Stateful field.
func (o *ControllerServiceDefinition) SetStateful(v Stateful) {
	o.Stateful = &v
}

// GetSystemResourceConsiderations returns the SystemResourceConsiderations field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetSystemResourceConsiderations() []SystemResourceConsideration {
	if o == nil || IsNil(o.SystemResourceConsiderations) {
		var ret []SystemResourceConsideration
		return ret
	}
	return o.SystemResourceConsiderations
}

// GetSystemResourceConsiderationsOk returns a tuple with the SystemResourceConsiderations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetSystemResourceConsiderationsOk() ([]SystemResourceConsideration, bool) {
	if o == nil || IsNil(o.SystemResourceConsiderations) {
		return nil, false
	}
	return o.SystemResourceConsiderations, true
}

// HasSystemResourceConsiderations returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasSystemResourceConsiderations() bool {
	if o != nil && !IsNil(o.SystemResourceConsiderations) {
		return true
	}

	return false
}

// SetSystemResourceConsiderations gets a reference to the given []SystemResourceConsideration and assigns it to the SystemResourceConsiderations field.
func (o *ControllerServiceDefinition) SetSystemResourceConsiderations(v []SystemResourceConsideration) {
	o.SystemResourceConsiderations = v
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetAdditionalDetails() bool {
	if o == nil || IsNil(o.AdditionalDetails) {
		var ret bool
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetAdditionalDetailsOk() (*bool, bool) {
	if o == nil || IsNil(o.AdditionalDetails) {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasAdditionalDetails() bool {
	if o != nil && !IsNil(o.AdditionalDetails) {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given bool and assigns it to the AdditionalDetails field.
func (o *ControllerServiceDefinition) SetAdditionalDetails(v bool) {
	o.AdditionalDetails = &v
}

// GetPropertyDescriptors returns the PropertyDescriptors field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetPropertyDescriptors() map[string]PropertyDescriptor {
	if o == nil || IsNil(o.PropertyDescriptors) {
		var ret map[string]PropertyDescriptor
		return ret
	}
	return *o.PropertyDescriptors
}

// GetPropertyDescriptorsOk returns a tuple with the PropertyDescriptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetPropertyDescriptorsOk() (*map[string]PropertyDescriptor, bool) {
	if o == nil || IsNil(o.PropertyDescriptors) {
		return nil, false
	}
	return o.PropertyDescriptors, true
}

// HasPropertyDescriptors returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasPropertyDescriptors() bool {
	if o != nil && !IsNil(o.PropertyDescriptors) {
		return true
	}

	return false
}

// SetPropertyDescriptors gets a reference to the given map[string]PropertyDescriptor and assigns it to the PropertyDescriptors field.
func (o *ControllerServiceDefinition) SetPropertyDescriptors(v map[string]PropertyDescriptor) {
	o.PropertyDescriptors = &v
}

// GetSupportsDynamicProperties returns the SupportsDynamicProperties field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetSupportsDynamicProperties() bool {
	if o == nil || IsNil(o.SupportsDynamicProperties) {
		var ret bool
		return ret
	}
	return *o.SupportsDynamicProperties
}

// GetSupportsDynamicPropertiesOk returns a tuple with the SupportsDynamicProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetSupportsDynamicPropertiesOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsDynamicProperties) {
		return nil, false
	}
	return o.SupportsDynamicProperties, true
}

// HasSupportsDynamicProperties returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasSupportsDynamicProperties() bool {
	if o != nil && !IsNil(o.SupportsDynamicProperties) {
		return true
	}

	return false
}

// SetSupportsDynamicProperties gets a reference to the given bool and assigns it to the SupportsDynamicProperties field.
func (o *ControllerServiceDefinition) SetSupportsDynamicProperties(v bool) {
	o.SupportsDynamicProperties = &v
}

// GetDynamicProperties returns the DynamicProperties field value if set, zero value otherwise.
func (o *ControllerServiceDefinition) GetDynamicProperties() []DynamicProperty {
	if o == nil || IsNil(o.DynamicProperties) {
		var ret []DynamicProperty
		return ret
	}
	return o.DynamicProperties
}

// GetDynamicPropertiesOk returns a tuple with the DynamicProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerServiceDefinition) GetDynamicPropertiesOk() ([]DynamicProperty, bool) {
	if o == nil || IsNil(o.DynamicProperties) {
		return nil, false
	}
	return o.DynamicProperties, true
}

// HasDynamicProperties returns a boolean if a field has been set.
func (o *ControllerServiceDefinition) HasDynamicProperties() bool {
	if o != nil && !IsNil(o.DynamicProperties) {
		return true
	}

	return false
}

// SetDynamicProperties gets a reference to the given []DynamicProperty and assigns it to the DynamicProperties field.
func (o *ControllerServiceDefinition) SetDynamicProperties(v []DynamicProperty) {
	o.DynamicProperties = v
}

func (o ControllerServiceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControllerServiceDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.TypeDescription) {
		toSerialize["typeDescription"] = o.TypeDescription
	}
	if !IsNil(o.BuildInfo) {
		toSerialize["buildInfo"] = o.BuildInfo
	}
	if !IsNil(o.ProvidedApiImplementations) {
		toSerialize["providedApiImplementations"] = o.ProvidedApiImplementations
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.SeeAlso) {
		toSerialize["seeAlso"] = o.SeeAlso
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !IsNil(o.DeprecationReason) {
		toSerialize["deprecationReason"] = o.DeprecationReason
	}
	if !IsNil(o.DeprecationAlternatives) {
		toSerialize["deprecationAlternatives"] = o.DeprecationAlternatives
	}
	if !IsNil(o.Restricted) {
		toSerialize["restricted"] = o.Restricted
	}
	if !IsNil(o.RestrictedExplanation) {
		toSerialize["restrictedExplanation"] = o.RestrictedExplanation
	}
	if !IsNil(o.ExplicitRestrictions) {
		toSerialize["explicitRestrictions"] = o.ExplicitRestrictions
	}
	if !IsNil(o.Stateful) {
		toSerialize["stateful"] = o.Stateful
	}
	if !IsNil(o.SystemResourceConsiderations) {
		toSerialize["systemResourceConsiderations"] = o.SystemResourceConsiderations
	}
	if !IsNil(o.AdditionalDetails) {
		toSerialize["additionalDetails"] = o.AdditionalDetails
	}
	if !IsNil(o.PropertyDescriptors) {
		toSerialize["propertyDescriptors"] = o.PropertyDescriptors
	}
	if !IsNil(o.SupportsDynamicProperties) {
		toSerialize["supportsDynamicProperties"] = o.SupportsDynamicProperties
	}
	if !IsNil(o.DynamicProperties) {
		toSerialize["dynamicProperties"] = o.DynamicProperties
	}
	return toSerialize, nil
}

type NullableControllerServiceDefinition struct {
	value *ControllerServiceDefinition
	isSet bool
}

func (v NullableControllerServiceDefinition) Get() *ControllerServiceDefinition {
	return v.value
}

func (v *NullableControllerServiceDefinition) Set(val *ControllerServiceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerServiceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerServiceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerServiceDefinition(val *ControllerServiceDefinition) *NullableControllerServiceDefinition {
	return &NullableControllerServiceDefinition{value: val, isSet: true}
}

func (v NullableControllerServiceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerServiceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

