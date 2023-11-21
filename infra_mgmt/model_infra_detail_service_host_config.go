/*
Infrastructure Management API

The **Infrastructure Management API** provides a RESTful interface to manage Infrastructure Hosts and Services objects.  The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):  | Service name | Service type |   | ------ | ------ |   | Access Authentication | authn |   | Anycast | anycast |   | Data Connector | cdc |   | DHCP | dhcp |   | DNS | dns |   | DNS Forwarding Proxy | dfp |   | NIOS Grid Connector | orpheus |   | MS AD Sync | msad |   | NTP | ntp |   | BGP | bgp |   | RIP | rip |   | OSPF | ospf |    ---   ### Hosts API  The Hosts API is used to manage the Infrastructure Host resources. These include various operations related to hosts such as viewing, creating, updating, replacing, disconnecting, and deleting Hosts. Management of Hosts is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Hosts tab.  ---   ### Services API  The Services API is used to manage the Infrastructure Service resources (a.k.a. BloxOne applications). These include various operations related to hosts such as viewing, creating, updating, starting/stopping, configuring, and deleting Services. Management of Services is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Services tab.  ---   ### Detail APIs  The Detail APIs are read-only APIs used to list all the Infrastructure resources (Hosts and Services). Each resource record returned also contains information about its other associated resources and the status information for itself and the associated resource(s) (i.e., Host/Service status).  ---

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_mgmt

import (
	"encoding/json"
	"time"
)

// checks if the InfraDetailServiceHostConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraDetailServiceHostConfig{}

// InfraDetailServiceHostConfig struct for InfraDetailServiceHostConfig
type InfraDetailServiceHostConfig struct {
	// The current version of the Service deployed on the Host.
	CurrentVersion *string                  `json:"current_version,omitempty"`
	Status         *InfraShortServiceStatus `json:"status,omitempty"`
	UpgradedAt     *time.Time               `json:"upgraded_at,omitempty"`
}

// NewInfraDetailServiceHostConfig instantiates a new InfraDetailServiceHostConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraDetailServiceHostConfig() *InfraDetailServiceHostConfig {
	this := InfraDetailServiceHostConfig{}
	return &this
}

// NewInfraDetailServiceHostConfigWithDefaults instantiates a new InfraDetailServiceHostConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraDetailServiceHostConfigWithDefaults() *InfraDetailServiceHostConfig {
	this := InfraDetailServiceHostConfig{}
	return &this
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *InfraDetailServiceHostConfig) GetCurrentVersion() string {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailServiceHostConfig) GetCurrentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *InfraDetailServiceHostConfig) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *InfraDetailServiceHostConfig) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InfraDetailServiceHostConfig) GetStatus() InfraShortServiceStatus {
	if o == nil || IsNil(o.Status) {
		var ret InfraShortServiceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailServiceHostConfig) GetStatusOk() (*InfraShortServiceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InfraDetailServiceHostConfig) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given InfraShortServiceStatus and assigns it to the Status field.
func (o *InfraDetailServiceHostConfig) SetStatus(v InfraShortServiceStatus) {
	o.Status = &v
}

// GetUpgradedAt returns the UpgradedAt field value if set, zero value otherwise.
func (o *InfraDetailServiceHostConfig) GetUpgradedAt() time.Time {
	if o == nil || IsNil(o.UpgradedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpgradedAt
}

// GetUpgradedAtOk returns a tuple with the UpgradedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraDetailServiceHostConfig) GetUpgradedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpgradedAt) {
		return nil, false
	}
	return o.UpgradedAt, true
}

// HasUpgradedAt returns a boolean if a field has been set.
func (o *InfraDetailServiceHostConfig) HasUpgradedAt() bool {
	if o != nil && !IsNil(o.UpgradedAt) {
		return true
	}

	return false
}

// SetUpgradedAt gets a reference to the given time.Time and assigns it to the UpgradedAt field.
func (o *InfraDetailServiceHostConfig) SetUpgradedAt(v time.Time) {
	o.UpgradedAt = &v
}

func (o InfraDetailServiceHostConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraDetailServiceHostConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentVersion) {
		toSerialize["current_version"] = o.CurrentVersion
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpgradedAt) {
		toSerialize["upgraded_at"] = o.UpgradedAt
	}
	return toSerialize, nil
}

type NullableInfraDetailServiceHostConfig struct {
	value *InfraDetailServiceHostConfig
	isSet bool
}

func (v NullableInfraDetailServiceHostConfig) Get() *InfraDetailServiceHostConfig {
	return v.value
}

func (v *NullableInfraDetailServiceHostConfig) Set(val *InfraDetailServiceHostConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraDetailServiceHostConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraDetailServiceHostConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraDetailServiceHostConfig(val *InfraDetailServiceHostConfig) *NullableInfraDetailServiceHostConfig {
	return &NullableInfraDetailServiceHostConfig{value: val, isSet: true}
}

func (v NullableInfraDetailServiceHostConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraDetailServiceHostConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}