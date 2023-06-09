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

// checks if the PlaylistObjectOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaylistObjectOwner{}

// PlaylistObjectOwner The user who owns the playlist 
type PlaylistObjectOwner struct {
	ExternalUrls *PublicUserObjectExternalUrls `json:"external_urls,omitempty"`
	Followers *PublicUserObjectFollowers `json:"followers,omitempty"`
	// A link to the Web API endpoint for this user. 
	Href *string `json:"href,omitempty"`
	// The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user. 
	Id *string `json:"id,omitempty"`
	// The object type. 
	Type *string `json:"type,omitempty"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this user. 
	Uri *string `json:"uri,omitempty"`
	// The name displayed on the user's profile. `null` if not available. 
	DisplayName NullableString `json:"display_name,omitempty"`
}

// NewPlaylistObjectOwner instantiates a new PlaylistObjectOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistObjectOwner() *PlaylistObjectOwner {
	this := PlaylistObjectOwner{}
	return &this
}

// NewPlaylistObjectOwnerWithDefaults instantiates a new PlaylistObjectOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistObjectOwnerWithDefaults() *PlaylistObjectOwner {
	this := PlaylistObjectOwner{}
	return &this
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *PlaylistObjectOwner) GetExternalUrls() PublicUserObjectExternalUrls {
	if o == nil || IsNil(o.ExternalUrls) {
		var ret PublicUserObjectExternalUrls
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistObjectOwner) GetExternalUrlsOk() (*PublicUserObjectExternalUrls, bool) {
	if o == nil || IsNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *PlaylistObjectOwner) HasExternalUrls() bool {
	if o != nil && !IsNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given PublicUserObjectExternalUrls and assigns it to the ExternalUrls field.
func (o *PlaylistObjectOwner) SetExternalUrls(v PublicUserObjectExternalUrls) {
	o.ExternalUrls = &v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *PlaylistObjectOwner) GetFollowers() PublicUserObjectFollowers {
	if o == nil || IsNil(o.Followers) {
		var ret PublicUserObjectFollowers
		return ret
	}
	return *o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistObjectOwner) GetFollowersOk() (*PublicUserObjectFollowers, bool) {
	if o == nil || IsNil(o.Followers) {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *PlaylistObjectOwner) HasFollowers() bool {
	if o != nil && !IsNil(o.Followers) {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given PublicUserObjectFollowers and assigns it to the Followers field.
func (o *PlaylistObjectOwner) SetFollowers(v PublicUserObjectFollowers) {
	o.Followers = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PlaylistObjectOwner) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistObjectOwner) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PlaylistObjectOwner) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PlaylistObjectOwner) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlaylistObjectOwner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistObjectOwner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlaylistObjectOwner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlaylistObjectOwner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlaylistObjectOwner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistObjectOwner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlaylistObjectOwner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlaylistObjectOwner) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *PlaylistObjectOwner) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistObjectOwner) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PlaylistObjectOwner) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *PlaylistObjectOwner) SetUri(v string) {
	o.Uri = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaylistObjectOwner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaylistObjectOwner) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PlaylistObjectOwner) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *PlaylistObjectOwner) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *PlaylistObjectOwner) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *PlaylistObjectOwner) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o PlaylistObjectOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaylistObjectOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !IsNil(o.Followers) {
		toSerialize["followers"] = o.Followers
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	return toSerialize, nil
}

type NullablePlaylistObjectOwner struct {
	value *PlaylistObjectOwner
	isSet bool
}

func (v NullablePlaylistObjectOwner) Get() *PlaylistObjectOwner {
	return v.value
}

func (v *NullablePlaylistObjectOwner) Set(val *PlaylistObjectOwner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistObjectOwner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistObjectOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistObjectOwner(val *PlaylistObjectOwner) *NullablePlaylistObjectOwner {
	return &NullablePlaylistObjectOwner{value: val, isSet: true}
}

func (v NullablePlaylistObjectOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistObjectOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


