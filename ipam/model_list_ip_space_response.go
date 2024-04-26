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

// checks if the ListIPSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIPSpaceResponse{}

// ListIPSpaceResponse The response format to retrieve __IPSpace__ objects.
type ListIPSpaceResponse struct {
	// The list of IPSpace objects.
	Results              []IPSpace `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListIPSpaceResponse ListIPSpaceResponse

// NewListIPSpaceResponse instantiates a new ListIPSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIPSpaceResponse() *ListIPSpaceResponse {
	this := ListIPSpaceResponse{}
	return &this
}

// NewListIPSpaceResponseWithDefaults instantiates a new ListIPSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIPSpaceResponseWithDefaults() *ListIPSpaceResponse {
	this := ListIPSpaceResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListIPSpaceResponse) GetResults() []IPSpace {
	if o == nil || IsNil(o.Results) {
		var ret []IPSpace
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIPSpaceResponse) GetResultsOk() ([]IPSpace, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListIPSpaceResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []IPSpace and assigns it to the Results field.
func (o *ListIPSpaceResponse) SetResults(v []IPSpace) {
	o.Results = v
}

func (o ListIPSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIPSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListIPSpaceResponse) UnmarshalJSON(data []byte) (err error) {
	varListIPSpaceResponse := _ListIPSpaceResponse{}

	err = json.Unmarshal(data, &varListIPSpaceResponse)

	if err != nil {
		return err
	}

	*o = ListIPSpaceResponse(varListIPSpaceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListIPSpaceResponse struct {
	value *ListIPSpaceResponse
	isSet bool
}

func (v NullableListIPSpaceResponse) Get() *ListIPSpaceResponse {
	return v.value
}

func (v *NullableListIPSpaceResponse) Set(val *ListIPSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListIPSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListIPSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIPSpaceResponse(val *ListIPSpaceResponse) *NullableListIPSpaceResponse {
	return &NullableListIPSpaceResponse{value: val, isSet: true}
}

func (v NullableListIPSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIPSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}