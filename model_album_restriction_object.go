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

// checks if the AlbumRestrictionObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumRestrictionObject{}

// AlbumRestrictionObject struct for AlbumRestrictionObject
type AlbumRestrictionObject struct {
	// The reason for the restriction. Albums may be restricted if the content is not available in a given market, to the user's subscription type, or when the user's account is set to not play explicit content. Additional reasons may be added in the future. 
	Reason *string `json:"reason,omitempty"`
}

// NewAlbumRestrictionObject instantiates a new AlbumRestrictionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumRestrictionObject() *AlbumRestrictionObject {
	this := AlbumRestrictionObject{}
	return &this
}

// NewAlbumRestrictionObjectWithDefaults instantiates a new AlbumRestrictionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumRestrictionObjectWithDefaults() *AlbumRestrictionObject {
	this := AlbumRestrictionObject{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AlbumRestrictionObject) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumRestrictionObject) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AlbumRestrictionObject) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AlbumRestrictionObject) SetReason(v string) {
	o.Reason = &v
}

func (o AlbumRestrictionObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumRestrictionObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableAlbumRestrictionObject struct {
	value *AlbumRestrictionObject
	isSet bool
}

func (v NullableAlbumRestrictionObject) Get() *AlbumRestrictionObject {
	return v.value
}

func (v *NullableAlbumRestrictionObject) Set(val *AlbumRestrictionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumRestrictionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumRestrictionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumRestrictionObject(val *AlbumRestrictionObject) *NullableAlbumRestrictionObject {
	return &NullableAlbumRestrictionObject{value: val, isSet: true}
}

func (v NullableAlbumRestrictionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumRestrictionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


