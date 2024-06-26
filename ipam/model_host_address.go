/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the HostAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostAddress{}

// HostAddress The __HostAddress__ object represents addresses associated with a Host object.
type HostAddress struct {
	// Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.
	Address *string `json:"address,omitempty"`
	// The resource identifier.
	Ref *string `json:"ref,omitempty"`
	// The resource identifier.
	Space                *string `json:"space,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HostAddress HostAddress

// NewHostAddress instantiates a new HostAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostAddress() *HostAddress {
	this := HostAddress{}
	return &this
}

// NewHostAddressWithDefaults instantiates a new HostAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostAddressWithDefaults() *HostAddress {
	this := HostAddress{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *HostAddress) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostAddress) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *HostAddress) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *HostAddress) SetAddress(v string) {
	o.Address = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *HostAddress) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostAddress) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *HostAddress) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *HostAddress) SetRef(v string) {
	o.Ref = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *HostAddress) GetSpace() string {
	if o == nil || IsNil(o.Space) {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostAddress) GetSpaceOk() (*string, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *HostAddress) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *HostAddress) SetSpace(v string) {
	o.Space = &v
}

func (o HostAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HostAddress) UnmarshalJSON(data []byte) (err error) {
	varHostAddress := _HostAddress{}

	err = json.Unmarshal(data, &varHostAddress)

	if err != nil {
		return err
	}

	*o = HostAddress(varHostAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "ref")
		delete(additionalProperties, "space")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHostAddress struct {
	value *HostAddress
	isSet bool
}

func (v NullableHostAddress) Get() *HostAddress {
	return v.value
}

func (v *NullableHostAddress) Set(val *HostAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableHostAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableHostAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostAddress(val *HostAddress) *NullableHostAddress {
	return &NullableHostAddress{value: val, isSet: true}
}

func (v NullableHostAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
