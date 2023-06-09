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

// checks if the SavedAudiobookObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedAudiobookObject{}

// SavedAudiobookObject struct for SavedAudiobookObject
type SavedAudiobookObject struct {
	// The date and time the audiobook was saved Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object. 
	AddedAt *time.Time `json:"added_at,omitempty"`
	Audiobook *SavedAudiobookObjectAudiobook `json:"audiobook,omitempty"`
}

// NewSavedAudiobookObject instantiates a new SavedAudiobookObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedAudiobookObject() *SavedAudiobookObject {
	this := SavedAudiobookObject{}
	return &this
}

// NewSavedAudiobookObjectWithDefaults instantiates a new SavedAudiobookObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedAudiobookObjectWithDefaults() *SavedAudiobookObject {
	this := SavedAudiobookObject{}
	return &this
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *SavedAudiobookObject) GetAddedAt() time.Time {
	if o == nil || IsNil(o.AddedAt) {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAudiobookObject) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddedAt) {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *SavedAudiobookObject) HasAddedAt() bool {
	if o != nil && !IsNil(o.AddedAt) {
		return true
	}

	return false
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *SavedAudiobookObject) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetAudiobook returns the Audiobook field value if set, zero value otherwise.
func (o *SavedAudiobookObject) GetAudiobook() SavedAudiobookObjectAudiobook {
	if o == nil || IsNil(o.Audiobook) {
		var ret SavedAudiobookObjectAudiobook
		return ret
	}
	return *o.Audiobook
}

// GetAudiobookOk returns a tuple with the Audiobook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAudiobookObject) GetAudiobookOk() (*SavedAudiobookObjectAudiobook, bool) {
	if o == nil || IsNil(o.Audiobook) {
		return nil, false
	}
	return o.Audiobook, true
}

// HasAudiobook returns a boolean if a field has been set.
func (o *SavedAudiobookObject) HasAudiobook() bool {
	if o != nil && !IsNil(o.Audiobook) {
		return true
	}

	return false
}

// SetAudiobook gets a reference to the given SavedAudiobookObjectAudiobook and assigns it to the Audiobook field.
func (o *SavedAudiobookObject) SetAudiobook(v SavedAudiobookObjectAudiobook) {
	o.Audiobook = &v
}

func (o SavedAudiobookObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedAudiobookObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddedAt) {
		toSerialize["added_at"] = o.AddedAt
	}
	if !IsNil(o.Audiobook) {
		toSerialize["audiobook"] = o.Audiobook
	}
	return toSerialize, nil
}

type NullableSavedAudiobookObject struct {
	value *SavedAudiobookObject
	isSet bool
}

func (v NullableSavedAudiobookObject) Get() *SavedAudiobookObject {
	return v.value
}

func (v *NullableSavedAudiobookObject) Set(val *SavedAudiobookObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedAudiobookObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedAudiobookObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedAudiobookObject(val *SavedAudiobookObject) *NullableSavedAudiobookObject {
	return &NullableSavedAudiobookObject{value: val, isSet: true}
}

func (v NullableSavedAudiobookObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedAudiobookObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

