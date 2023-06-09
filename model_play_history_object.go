/*
Spotify Web API with fixes and improvements from sonallux

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 2023.6.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the PlayHistoryObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayHistoryObject{}

// PlayHistoryObject struct for PlayHistoryObject
type PlayHistoryObject struct {
	Track *PlayHistoryObjectTrack `json:"track,omitempty"`
	// The date and time the track was played.
	PlayedAt *time.Time `json:"played_at,omitempty"`
	Context *PlayHistoryObjectContext `json:"context,omitempty"`
}

// NewPlayHistoryObject instantiates a new PlayHistoryObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayHistoryObject() *PlayHistoryObject {
	this := PlayHistoryObject{}
	return &this
}

// NewPlayHistoryObjectWithDefaults instantiates a new PlayHistoryObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayHistoryObjectWithDefaults() *PlayHistoryObject {
	this := PlayHistoryObject{}
	return &this
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *PlayHistoryObject) GetTrack() PlayHistoryObjectTrack {
	if o == nil || IsNil(o.Track) {
		var ret PlayHistoryObjectTrack
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayHistoryObject) GetTrackOk() (*PlayHistoryObjectTrack, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *PlayHistoryObject) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given PlayHistoryObjectTrack and assigns it to the Track field.
func (o *PlayHistoryObject) SetTrack(v PlayHistoryObjectTrack) {
	o.Track = &v
}

// GetPlayedAt returns the PlayedAt field value if set, zero value otherwise.
func (o *PlayHistoryObject) GetPlayedAt() time.Time {
	if o == nil || IsNil(o.PlayedAt) {
		var ret time.Time
		return ret
	}
	return *o.PlayedAt
}

// GetPlayedAtOk returns a tuple with the PlayedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayHistoryObject) GetPlayedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PlayedAt) {
		return nil, false
	}
	return o.PlayedAt, true
}

// HasPlayedAt returns a boolean if a field has been set.
func (o *PlayHistoryObject) HasPlayedAt() bool {
	if o != nil && !IsNil(o.PlayedAt) {
		return true
	}

	return false
}

// SetPlayedAt gets a reference to the given time.Time and assigns it to the PlayedAt field.
func (o *PlayHistoryObject) SetPlayedAt(v time.Time) {
	o.PlayedAt = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PlayHistoryObject) GetContext() PlayHistoryObjectContext {
	if o == nil || IsNil(o.Context) {
		var ret PlayHistoryObjectContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayHistoryObject) GetContextOk() (*PlayHistoryObjectContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PlayHistoryObject) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given PlayHistoryObjectContext and assigns it to the Context field.
func (o *PlayHistoryObject) SetContext(v PlayHistoryObjectContext) {
	o.Context = &v
}

func (o PlayHistoryObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayHistoryObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	if !IsNil(o.PlayedAt) {
		toSerialize["played_at"] = o.PlayedAt
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return toSerialize, nil
}

type NullablePlayHistoryObject struct {
	value *PlayHistoryObject
	isSet bool
}

func (v NullablePlayHistoryObject) Get() *PlayHistoryObject {
	return v.value
}

func (v *NullablePlayHistoryObject) Set(val *PlayHistoryObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayHistoryObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayHistoryObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayHistoryObject(val *PlayHistoryObject) *NullablePlayHistoryObject {
	return &NullablePlayHistoryObject{value: val, isSet: true}
}

func (v NullablePlayHistoryObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayHistoryObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


