/*
BloxOne Anycast API

Anycast capability enables HA (High Availability) configuration of BloxOne applications that run on equipment located on customer's premises (on-prem hosts). Anycast supports DNS, as well as DNS-forwarding services.  Anycast-enabled application setups use multiple on-premises installations for one particular application type. Multiple application instances are configured to use the same endpoint address. Anycast capability is collocated with such application instance, monitoring the local application instance and advertising to the upstream router (a customer equipment) a per-instance, local route to the common application endpoint address, as long as the local application instance is available. Depending on the type of the upstream router, the customer may configure local route advertisement via either BGP (Boarder Gateway Protocol) or OSPF (Open Shortest Path First) routing protocols. Both protocols may be enabled as well. Multiple routes to the common application service address provide redundancy without the need to reconfigure application clients.  Should an application instance become unavailable, the local route advertisements stop, resulting in withdrawal of the route (in the upstream router) to the application instance that has gone out of service and ensuring that subsequent application requests thus get routed to the remaining available application instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anycast

import (
	"encoding/json"
)

// checks if the OnpremHostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnpremHostResponse{}

// OnpremHostResponse struct for OnpremHostResponse
type OnpremHostResponse struct {
	Results              *OnpremHost `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnpremHostResponse OnpremHostResponse

// NewOnpremHostResponse instantiates a new OnpremHostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnpremHostResponse() *OnpremHostResponse {
	this := OnpremHostResponse{}
	return &this
}

// NewOnpremHostResponseWithDefaults instantiates a new OnpremHostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnpremHostResponseWithDefaults() *OnpremHostResponse {
	this := OnpremHostResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *OnpremHostResponse) GetResults() OnpremHost {
	if o == nil || IsNil(o.Results) {
		var ret OnpremHost
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremHostResponse) GetResultsOk() (*OnpremHost, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *OnpremHostResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given OnpremHost and assigns it to the Results field.
func (o *OnpremHostResponse) SetResults(v OnpremHost) {
	o.Results = &v
}

func (o OnpremHostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnpremHostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnpremHostResponse) UnmarshalJSON(data []byte) (err error) {
	varOnpremHostResponse := _OnpremHostResponse{}

	err = json.Unmarshal(data, &varOnpremHostResponse)

	if err != nil {
		return err
	}

	*o = OnpremHostResponse(varOnpremHostResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOnpremHostResponse struct {
	value *OnpremHostResponse
	isSet bool
}

func (v NullableOnpremHostResponse) Get() *OnpremHostResponse {
	return v.value
}

func (v *NullableOnpremHostResponse) Set(val *OnpremHostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOnpremHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOnpremHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnpremHostResponse(val *OnpremHostResponse) *NullableOnpremHostResponse {
	return &NullableOnpremHostResponse{value: val, isSet: true}
}

func (v NullableOnpremHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnpremHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}