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

// checks if the SaveAlbumsUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveAlbumsUserRequest{}

// SaveAlbumsUserRequest struct for SaveAlbumsUserRequest
type SaveAlbumsUserRequest struct {
	// A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: `[\"4iV5W9uYEdYUVa79Axb7Rh\", \"1301WleyT98MSxVHPZCA6M\"]`<br/>A maximum of 50 items can be specified in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._ 
	Ids []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SaveAlbumsUserRequest SaveAlbumsUserRequest

// NewSaveAlbumsUserRequest instantiates a new SaveAlbumsUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveAlbumsUserRequest() *SaveAlbumsUserRequest {
	this := SaveAlbumsUserRequest{}
	return &this
}

// NewSaveAlbumsUserRequestWithDefaults instantiates a new SaveAlbumsUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveAlbumsUserRequestWithDefaults() *SaveAlbumsUserRequest {
	this := SaveAlbumsUserRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *SaveAlbumsUserRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAlbumsUserRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *SaveAlbumsUserRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *SaveAlbumsUserRequest) SetIds(v []string) {
	o.Ids = v
}

func (o SaveAlbumsUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveAlbumsUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SaveAlbumsUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSaveAlbumsUserRequest := _SaveAlbumsUserRequest{}

	if err = json.Unmarshal(bytes, &varSaveAlbumsUserRequest); err == nil {
		*o = SaveAlbumsUserRequest(varSaveAlbumsUserRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSaveAlbumsUserRequest struct {
	value *SaveAlbumsUserRequest
	isSet bool
}

func (v NullableSaveAlbumsUserRequest) Get() *SaveAlbumsUserRequest {
	return v.value
}

func (v *NullableSaveAlbumsUserRequest) Set(val *SaveAlbumsUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveAlbumsUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveAlbumsUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveAlbumsUserRequest(val *SaveAlbumsUserRequest) *NullableSaveAlbumsUserRequest {
	return &NullableSaveAlbumsUserRequest{value: val, isSet: true}
}

func (v NullableSaveAlbumsUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveAlbumsUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


