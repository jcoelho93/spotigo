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

// checks if the SaveShowsUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveShowsUserRequest{}

// SaveShowsUserRequest struct for SaveShowsUserRequest
type SaveShowsUserRequest struct {
	// A JSON array of the [Spotify IDs](https://developer.spotify.com/documentation/web-api/#spotify-uris-and-ids).   A maximum of 50 items can be specified in one request. *Note: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored.*
	Ids []string `json:"ids,omitempty"`
}

// NewSaveShowsUserRequest instantiates a new SaveShowsUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveShowsUserRequest() *SaveShowsUserRequest {
	this := SaveShowsUserRequest{}
	return &this
}

// NewSaveShowsUserRequestWithDefaults instantiates a new SaveShowsUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveShowsUserRequestWithDefaults() *SaveShowsUserRequest {
	this := SaveShowsUserRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *SaveShowsUserRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveShowsUserRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *SaveShowsUserRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *SaveShowsUserRequest) SetIds(v []string) {
	o.Ids = v
}

func (o SaveShowsUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveShowsUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableSaveShowsUserRequest struct {
	value *SaveShowsUserRequest
	isSet bool
}

func (v NullableSaveShowsUserRequest) Get() *SaveShowsUserRequest {
	return v.value
}

func (v *NullableSaveShowsUserRequest) Set(val *SaveShowsUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveShowsUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveShowsUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveShowsUserRequest(val *SaveShowsUserRequest) *NullableSaveShowsUserRequest {
	return &NullableSaveShowsUserRequest{value: val, isSet: true}
}

func (v NullableSaveShowsUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveShowsUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


