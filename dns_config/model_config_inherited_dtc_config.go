/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigInheritedDtcConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigInheritedDtcConfig{}

// ConfigInheritedDtcConfig Inheritance configuration for a field of type _DTCConfig_.
type ConfigInheritedDtcConfig struct {
	DefaultTtl *Inheritance2InheritedUInt32 `json:"default_ttl,omitempty"`
}

// NewConfigInheritedDtcConfig instantiates a new ConfigInheritedDtcConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigInheritedDtcConfig() *ConfigInheritedDtcConfig {
	this := ConfigInheritedDtcConfig{}
	return &this
}

// NewConfigInheritedDtcConfigWithDefaults instantiates a new ConfigInheritedDtcConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigInheritedDtcConfigWithDefaults() *ConfigInheritedDtcConfig {
	this := ConfigInheritedDtcConfig{}
	return &this
}

// GetDefaultTtl returns the DefaultTtl field value if set, zero value otherwise.
func (o *ConfigInheritedDtcConfig) GetDefaultTtl() Inheritance2InheritedUInt32 {
	if o == nil || IsNil(o.DefaultTtl) {
		var ret Inheritance2InheritedUInt32
		return ret
	}
	return *o.DefaultTtl
}

// GetDefaultTtlOk returns a tuple with the DefaultTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInheritedDtcConfig) GetDefaultTtlOk() (*Inheritance2InheritedUInt32, bool) {
	if o == nil || IsNil(o.DefaultTtl) {
		return nil, false
	}
	return o.DefaultTtl, true
}

// HasDefaultTtl returns a boolean if a field has been set.
func (o *ConfigInheritedDtcConfig) HasDefaultTtl() bool {
	if o != nil && !IsNil(o.DefaultTtl) {
		return true
	}

	return false
}

// SetDefaultTtl gets a reference to the given Inheritance2InheritedUInt32 and assigns it to the DefaultTtl field.
func (o *ConfigInheritedDtcConfig) SetDefaultTtl(v Inheritance2InheritedUInt32) {
	o.DefaultTtl = &v
}

func (o ConfigInheritedDtcConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigInheritedDtcConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultTtl) {
		toSerialize["default_ttl"] = o.DefaultTtl
	}
	return toSerialize, nil
}

type NullableConfigInheritedDtcConfig struct {
	value *ConfigInheritedDtcConfig
	isSet bool
}

func (v NullableConfigInheritedDtcConfig) Get() *ConfigInheritedDtcConfig {
	return v.value
}

func (v *NullableConfigInheritedDtcConfig) Set(val *ConfigInheritedDtcConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigInheritedDtcConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigInheritedDtcConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigInheritedDtcConfig(val *ConfigInheritedDtcConfig) *NullableConfigInheritedDtcConfig {
	return &NullableConfigInheritedDtcConfig{value: val, isSet: true}
}

func (v NullableConfigInheritedDtcConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigInheritedDtcConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}