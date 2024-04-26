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

// checks if the ACL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACL{}

// ACL Named ACL (Access Control List).
type ACL struct {
	// Optional. Comment for ACL.
	Comment *string `json:"comment,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// Optional. Ordered list of access control elements.  Elements are evaluated in order to determine access. If evaluation reaches the end of the list then access is denied.
	List []ACLItem `json:"list,omitempty"`
	// ACL object name.
	Name string `json:"name"`
	// Tagging specifics.
	Tags                 map[string]interface{} `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ACL ACL

// NewACL instantiates a new ACL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACL(name string) *ACL {
	this := ACL{}
	this.Name = name
	return &this
}

// NewACLWithDefaults instantiates a new ACL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLWithDefaults() *ACL {
	this := ACL{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ACL) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACL) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ACL) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ACL) SetComment(v string) {
	o.Comment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ACL) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACL) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ACL) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ACL) SetId(v string) {
	o.Id = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ACL) GetList() []ACLItem {
	if o == nil || IsNil(o.List) {
		var ret []ACLItem
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACL) GetListOk() ([]ACLItem, bool) {
	if o == nil || IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ACL) HasList() bool {
	if o != nil && !IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []ACLItem and assigns it to the List field.
func (o *ACL) SetList(v []ACLItem) {
	o.List = v
}

// GetName returns the Name field value
func (o *ACL) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ACL) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ACL) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ACL) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACL) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ACL) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ACL) SetTags(v map[string]interface{}) {
	o.Tags = v
}

func (o ACL) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.List) {
		toSerialize["list"] = o.List
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ACL) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varACL := _ACL{}

	err = json.Unmarshal(data, &varACL)

	if err != nil {
		return err
	}

	*o = ACL(varACL)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "id")
		delete(additionalProperties, "list")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableACL struct {
	value *ACL
	isSet bool
}

func (v NullableACL) Get() *ACL {
	return v.value
}

func (v *NullableACL) Set(val *ACL) {
	v.value = val
	v.isSet = true
}

func (v NullableACL) IsSet() bool {
	return v.isSet
}

func (v *NullableACL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACL(val *ACL) *NullableACL {
	return &NullableACL{value: val, isSet: true}
}

func (v NullableACL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}