/*
Spotify Web API with fixes and improvements from sonallux

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 2023.6.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FollowersObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowersObject{}

// FollowersObject struct for FollowersObject
type FollowersObject struct {
	// This will always be set to null, as the Web API does not support it at the moment. 
	Href NullableString `json:"href,omitempty"`
	// The total number of followers. 
	Total *int32 `json:"total,omitempty"`
}

// NewFollowersObject instantiates a new FollowersObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowersObject() *FollowersObject {
	this := FollowersObject{}
	return &this
}

// NewFollowersObjectWithDefaults instantiates a new FollowersObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowersObjectWithDefaults() *FollowersObject {
	this := FollowersObject{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FollowersObject) GetHref() string {
	if o == nil || IsNil(o.Href.Get()) {
		var ret string
		return ret
	}
	return *o.Href.Get()
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FollowersObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Href.Get(), o.Href.IsSet()
}

// HasHref returns a boolean if a field has been set.
func (o *FollowersObject) HasHref() bool {
	if o != nil && o.Href.IsSet() {
		return true
	}

	return false
}

// SetHref gets a reference to the given NullableString and assigns it to the Href field.
func (o *FollowersObject) SetHref(v string) {
	o.Href.Set(&v)
}
// SetHrefNil sets the value for Href to be an explicit nil
func (o *FollowersObject) SetHrefNil() {
	o.Href.Set(nil)
}

// UnsetHref ensures that no value is present for Href, not even an explicit nil
func (o *FollowersObject) UnsetHref() {
	o.Href.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *FollowersObject) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowersObject) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *FollowersObject) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *FollowersObject) SetTotal(v int32) {
	o.Total = &v
}

func (o FollowersObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowersObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Href.IsSet() {
		toSerialize["href"] = o.Href.Get()
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableFollowersObject struct {
	value *FollowersObject
	isSet bool
}

func (v NullableFollowersObject) Get() *FollowersObject {
	return v.value
}

func (v *NullableFollowersObject) Set(val *FollowersObject) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowersObject) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowersObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowersObject(val *FollowersObject) *NullableFollowersObject {
	return &NullableFollowersObject{value: val, isSet: true}
}

func (v NullableFollowersObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowersObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


