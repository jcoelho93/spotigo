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

// checks if the PagingSavedAlbumObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingSavedAlbumObjectAllOf{}

// PagingSavedAlbumObjectAllOf struct for PagingSavedAlbumObjectAllOf
type PagingSavedAlbumObjectAllOf struct {
	Items []SavedAlbumObject `json:"items,omitempty"`
}

// NewPagingSavedAlbumObjectAllOf instantiates a new PagingSavedAlbumObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingSavedAlbumObjectAllOf() *PagingSavedAlbumObjectAllOf {
	this := PagingSavedAlbumObjectAllOf{}
	return &this
}

// NewPagingSavedAlbumObjectAllOfWithDefaults instantiates a new PagingSavedAlbumObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingSavedAlbumObjectAllOfWithDefaults() *PagingSavedAlbumObjectAllOf {
	this := PagingSavedAlbumObjectAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PagingSavedAlbumObjectAllOf) GetItems() []SavedAlbumObject {
	if o == nil || IsNil(o.Items) {
		var ret []SavedAlbumObject
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingSavedAlbumObjectAllOf) GetItemsOk() ([]SavedAlbumObject, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PagingSavedAlbumObjectAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SavedAlbumObject and assigns it to the Items field.
func (o *PagingSavedAlbumObjectAllOf) SetItems(v []SavedAlbumObject) {
	o.Items = v
}

func (o PagingSavedAlbumObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingSavedAlbumObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullablePagingSavedAlbumObjectAllOf struct {
	value *PagingSavedAlbumObjectAllOf
	isSet bool
}

func (v NullablePagingSavedAlbumObjectAllOf) Get() *PagingSavedAlbumObjectAllOf {
	return v.value
}

func (v *NullablePagingSavedAlbumObjectAllOf) Set(val *PagingSavedAlbumObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingSavedAlbumObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingSavedAlbumObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingSavedAlbumObjectAllOf(val *PagingSavedAlbumObjectAllOf) *NullablePagingSavedAlbumObjectAllOf {
	return &NullablePagingSavedAlbumObjectAllOf{value: val, isSet: true}
}

func (v NullablePagingSavedAlbumObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingSavedAlbumObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


