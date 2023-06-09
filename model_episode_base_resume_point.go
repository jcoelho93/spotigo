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

// checks if the EpisodeBaseResumePoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpisodeBaseResumePoint{}

// EpisodeBaseResumePoint The user's most recent position in the episode. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'. 
type EpisodeBaseResumePoint struct {
	// Whether or not the episode has been fully played by the user. 
	FullyPlayed *bool `json:"fully_played,omitempty"`
	// The user's most recent position in the episode in milliseconds. 
	ResumePositionMs *int32 `json:"resume_position_ms,omitempty"`
}

// NewEpisodeBaseResumePoint instantiates a new EpisodeBaseResumePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpisodeBaseResumePoint() *EpisodeBaseResumePoint {
	this := EpisodeBaseResumePoint{}
	return &this
}

// NewEpisodeBaseResumePointWithDefaults instantiates a new EpisodeBaseResumePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpisodeBaseResumePointWithDefaults() *EpisodeBaseResumePoint {
	this := EpisodeBaseResumePoint{}
	return &this
}

// GetFullyPlayed returns the FullyPlayed field value if set, zero value otherwise.
func (o *EpisodeBaseResumePoint) GetFullyPlayed() bool {
	if o == nil || IsNil(o.FullyPlayed) {
		var ret bool
		return ret
	}
	return *o.FullyPlayed
}

// GetFullyPlayedOk returns a tuple with the FullyPlayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeBaseResumePoint) GetFullyPlayedOk() (*bool, bool) {
	if o == nil || IsNil(o.FullyPlayed) {
		return nil, false
	}
	return o.FullyPlayed, true
}

// HasFullyPlayed returns a boolean if a field has been set.
func (o *EpisodeBaseResumePoint) HasFullyPlayed() bool {
	if o != nil && !IsNil(o.FullyPlayed) {
		return true
	}

	return false
}

// SetFullyPlayed gets a reference to the given bool and assigns it to the FullyPlayed field.
func (o *EpisodeBaseResumePoint) SetFullyPlayed(v bool) {
	o.FullyPlayed = &v
}

// GetResumePositionMs returns the ResumePositionMs field value if set, zero value otherwise.
func (o *EpisodeBaseResumePoint) GetResumePositionMs() int32 {
	if o == nil || IsNil(o.ResumePositionMs) {
		var ret int32
		return ret
	}
	return *o.ResumePositionMs
}

// GetResumePositionMsOk returns a tuple with the ResumePositionMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeBaseResumePoint) GetResumePositionMsOk() (*int32, bool) {
	if o == nil || IsNil(o.ResumePositionMs) {
		return nil, false
	}
	return o.ResumePositionMs, true
}

// HasResumePositionMs returns a boolean if a field has been set.
func (o *EpisodeBaseResumePoint) HasResumePositionMs() bool {
	if o != nil && !IsNil(o.ResumePositionMs) {
		return true
	}

	return false
}

// SetResumePositionMs gets a reference to the given int32 and assigns it to the ResumePositionMs field.
func (o *EpisodeBaseResumePoint) SetResumePositionMs(v int32) {
	o.ResumePositionMs = &v
}

func (o EpisodeBaseResumePoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpisodeBaseResumePoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FullyPlayed) {
		toSerialize["fully_played"] = o.FullyPlayed
	}
	if !IsNil(o.ResumePositionMs) {
		toSerialize["resume_position_ms"] = o.ResumePositionMs
	}
	return toSerialize, nil
}

type NullableEpisodeBaseResumePoint struct {
	value *EpisodeBaseResumePoint
	isSet bool
}

func (v NullableEpisodeBaseResumePoint) Get() *EpisodeBaseResumePoint {
	return v.value
}

func (v *NullableEpisodeBaseResumePoint) Set(val *EpisodeBaseResumePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisodeBaseResumePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisodeBaseResumePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisodeBaseResumePoint(val *EpisodeBaseResumePoint) *NullableEpisodeBaseResumePoint {
	return &NullableEpisodeBaseResumePoint{value: val, isSet: true}
}

func (v NullableEpisodeBaseResumePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisodeBaseResumePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


