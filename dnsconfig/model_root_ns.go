/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
	"fmt"
)

// checks if the RootNS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RootNS{}

// RootNS Root nameserver
type RootNS struct {
	// IPv4 address.
	Address string `json:"address"`
	// FQDN.
	Fqdn string `json:"fqdn"`
	// FQDN in punycode.
	ProtocolFqdn         *string `json:"protocol_fqdn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RootNS RootNS

// NewRootNS instantiates a new RootNS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRootNS(address string, fqdn string) *RootNS {
	this := RootNS{}
	this.Address = address
	this.Fqdn = fqdn
	return &this
}

// NewRootNSWithDefaults instantiates a new RootNS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRootNSWithDefaults() *RootNS {
	this := RootNS{}
	return &this
}

// GetAddress returns the Address field value
func (o *RootNS) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *RootNS) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *RootNS) SetAddress(v string) {
	o.Address = v
}

// GetFqdn returns the Fqdn field value
func (o *RootNS) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *RootNS) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *RootNS) SetFqdn(v string) {
	o.Fqdn = v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *RootNS) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RootNS) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *RootNS) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *RootNS) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

func (o RootNS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RootNS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["fqdn"] = o.Fqdn
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RootNS) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"fqdn",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRootNS := _RootNS{}

	err = json.Unmarshal(data, &varRootNS)

	if err != nil {
		return err
	}

	*o = RootNS(varRootNS)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "protocol_fqdn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRootNS struct {
	value *RootNS
	isSet bool
}

func (v NullableRootNS) Get() *RootNS {
	return v.value
}

func (v *NullableRootNS) Set(val *RootNS) {
	v.value = val
	v.isSet = true
}

func (v NullableRootNS) IsSet() bool {
	return v.isSet
}

func (v *NullableRootNS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRootNS(val *RootNS) *NullableRootNS {
	return &NullableRootNS{value: val, isSet: true}
}

func (v NullableRootNS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRootNS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
