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

// checks if the LinkedTrackObjectExternalUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedTrackObjectExternalUrls{}

// LinkedTrackObjectExternalUrls Known external URLs for this track. 
type LinkedTrackObjectExternalUrls struct {
	// The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object. 
	Spotify *string `json:"spotify,omitempty"`
}

// NewLinkedTrackObjectExternalUrls instantiates a new LinkedTrackObjectExternalUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedTrackObjectExternalUrls() *LinkedTrackObjectExternalUrls {
	this := LinkedTrackObjectExternalUrls{}
	return &this
}

// NewLinkedTrackObjectExternalUrlsWithDefaults instantiates a new LinkedTrackObjectExternalUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedTrackObjectExternalUrlsWithDefaults() *LinkedTrackObjectExternalUrls {
	this := LinkedTrackObjectExternalUrls{}
	return &this
}

// GetSpotify returns the Spotify field value if set, zero value otherwise.
func (o *LinkedTrackObjectExternalUrls) GetSpotify() string {
	if o == nil || IsNil(o.Spotify) {
		var ret string
		return ret
	}
	return *o.Spotify
}

// GetSpotifyOk returns a tuple with the Spotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTrackObjectExternalUrls) GetSpotifyOk() (*string, bool) {
	if o == nil || IsNil(o.Spotify) {
		return nil, false
	}
	return o.Spotify, true
}

// HasSpotify returns a boolean if a field has been set.
func (o *LinkedTrackObjectExternalUrls) HasSpotify() bool {
	if o != nil && !IsNil(o.Spotify) {
		return true
	}

	return false
}

// SetSpotify gets a reference to the given string and assigns it to the Spotify field.
func (o *LinkedTrackObjectExternalUrls) SetSpotify(v string) {
	o.Spotify = &v
}

func (o LinkedTrackObjectExternalUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedTrackObjectExternalUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Spotify) {
		toSerialize["spotify"] = o.Spotify
	}
	return toSerialize, nil
}

type NullableLinkedTrackObjectExternalUrls struct {
	value *LinkedTrackObjectExternalUrls
	isSet bool
}

func (v NullableLinkedTrackObjectExternalUrls) Get() *LinkedTrackObjectExternalUrls {
	return v.value
}

func (v *NullableLinkedTrackObjectExternalUrls) Set(val *LinkedTrackObjectExternalUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedTrackObjectExternalUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedTrackObjectExternalUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedTrackObjectExternalUrls(val *LinkedTrackObjectExternalUrls) *NullableLinkedTrackObjectExternalUrls {
	return &NullableLinkedTrackObjectExternalUrls{value: val, isSet: true}
}

func (v NullableLinkedTrackObjectExternalUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedTrackObjectExternalUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


