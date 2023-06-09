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

// checks if the RemoveTracksPlaylistRequestTracksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveTracksPlaylistRequestTracksInner{}

// RemoveTracksPlaylistRequestTracksInner struct for RemoveTracksPlaylistRequestTracksInner
type RemoveTracksPlaylistRequestTracksInner struct {
	// Spotify URI
	Uri *string `json:"uri,omitempty"`
}

// NewRemoveTracksPlaylistRequestTracksInner instantiates a new RemoveTracksPlaylistRequestTracksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveTracksPlaylistRequestTracksInner() *RemoveTracksPlaylistRequestTracksInner {
	this := RemoveTracksPlaylistRequestTracksInner{}
	return &this
}

// NewRemoveTracksPlaylistRequestTracksInnerWithDefaults instantiates a new RemoveTracksPlaylistRequestTracksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveTracksPlaylistRequestTracksInnerWithDefaults() *RemoveTracksPlaylistRequestTracksInner {
	this := RemoveTracksPlaylistRequestTracksInner{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *RemoveTracksPlaylistRequestTracksInner) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveTracksPlaylistRequestTracksInner) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *RemoveTracksPlaylistRequestTracksInner) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *RemoveTracksPlaylistRequestTracksInner) SetUri(v string) {
	o.Uri = &v
}

func (o RemoveTracksPlaylistRequestTracksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveTracksPlaylistRequestTracksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableRemoveTracksPlaylistRequestTracksInner struct {
	value *RemoveTracksPlaylistRequestTracksInner
	isSet bool
}

func (v NullableRemoveTracksPlaylistRequestTracksInner) Get() *RemoveTracksPlaylistRequestTracksInner {
	return v.value
}

func (v *NullableRemoveTracksPlaylistRequestTracksInner) Set(val *RemoveTracksPlaylistRequestTracksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveTracksPlaylistRequestTracksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveTracksPlaylistRequestTracksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveTracksPlaylistRequestTracksInner(val *RemoveTracksPlaylistRequestTracksInner) *NullableRemoveTracksPlaylistRequestTracksInner {
	return &NullableRemoveTracksPlaylistRequestTracksInner{value: val, isSet: true}
}

func (v NullableRemoveTracksPlaylistRequestTracksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveTracksPlaylistRequestTracksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


