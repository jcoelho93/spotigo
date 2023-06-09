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

// checks if the TrackObjectAlbum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackObjectAlbum{}

// TrackObjectAlbum The album on which the track appears. The album object includes a link in `href` to full information about the album. 
type TrackObjectAlbum struct {
	// The type of the album. 
	AlbumType string `json:"album_type"`
	// The number of tracks in the album.
	TotalTracks int32 `json:"total_tracks"`
	// The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._ 
	AvailableMarkets []string `json:"available_markets"`
	ExternalUrls AlbumBaseExternalUrls `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the album. 
	Href string `json:"href"`
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the album. 
	Id string `json:"id"`
	// The cover art for the album in various sizes, widest first. 
	Images []ImageObject `json:"images"`
	// The name of the album. In case of an album takedown, the value may be an empty string. 
	Name string `json:"name"`
	// The date the album was first released. 
	ReleaseDate string `json:"release_date"`
	// The precision with which `release_date` value is known. 
	ReleaseDatePrecision string `json:"release_date_precision"`
	Restrictions *AlbumBaseRestrictions `json:"restrictions,omitempty"`
	// The object type. 
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the album. 
	Uri string `json:"uri"`
	// The field is present when getting an artist's albums. Compare to album_type this field represents relationship between the artist and the album. 
	AlbumGroup *string `json:"album_group,omitempty"`
	// The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist. 
	Artists []SimplifiedArtistObject `json:"artists"`
}

// NewTrackObjectAlbum instantiates a new TrackObjectAlbum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackObjectAlbum(albumType string, totalTracks int32, availableMarkets []string, externalUrls AlbumBaseExternalUrls, href string, id string, images []ImageObject, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, artists []SimplifiedArtistObject) *TrackObjectAlbum {
	this := TrackObjectAlbum{}
	this.AlbumType = albumType
	this.TotalTracks = totalTracks
	this.AvailableMarkets = availableMarkets
	this.ExternalUrls = externalUrls
	this.Href = href
	this.Id = id
	this.Images = images
	this.Name = name
	this.ReleaseDate = releaseDate
	this.ReleaseDatePrecision = releaseDatePrecision
	this.Type = type_
	this.Uri = uri
	this.Artists = artists
	return &this
}

// NewTrackObjectAlbumWithDefaults instantiates a new TrackObjectAlbum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackObjectAlbumWithDefaults() *TrackObjectAlbum {
	this := TrackObjectAlbum{}
	return &this
}

// GetAlbumType returns the AlbumType field value
func (o *TrackObjectAlbum) GetAlbumType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlbumType
}

// GetAlbumTypeOk returns a tuple with the AlbumType field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetAlbumTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlbumType, true
}

// SetAlbumType sets field value
func (o *TrackObjectAlbum) SetAlbumType(v string) {
	o.AlbumType = v
}

// GetTotalTracks returns the TotalTracks field value
func (o *TrackObjectAlbum) GetTotalTracks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTracks
}

// GetTotalTracksOk returns a tuple with the TotalTracks field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetTotalTracksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTracks, true
}

// SetTotalTracks sets field value
func (o *TrackObjectAlbum) SetTotalTracks(v int32) {
	o.TotalTracks = v
}

// GetAvailableMarkets returns the AvailableMarkets field value
func (o *TrackObjectAlbum) GetAvailableMarkets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableMarkets
}

// GetAvailableMarketsOk returns a tuple with the AvailableMarkets field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetAvailableMarketsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableMarkets, true
}

// SetAvailableMarkets sets field value
func (o *TrackObjectAlbum) SetAvailableMarkets(v []string) {
	o.AvailableMarkets = v
}

// GetExternalUrls returns the ExternalUrls field value
func (o *TrackObjectAlbum) GetExternalUrls() AlbumBaseExternalUrls {
	if o == nil {
		var ret AlbumBaseExternalUrls
		return ret
	}

	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetExternalUrlsOk() (*AlbumBaseExternalUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrls, true
}

// SetExternalUrls sets field value
func (o *TrackObjectAlbum) SetExternalUrls(v AlbumBaseExternalUrls) {
	o.ExternalUrls = v
}

// GetHref returns the Href field value
func (o *TrackObjectAlbum) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *TrackObjectAlbum) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *TrackObjectAlbum) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TrackObjectAlbum) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *TrackObjectAlbum) GetImages() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetImagesOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *TrackObjectAlbum) SetImages(v []ImageObject) {
	o.Images = v
}

// GetName returns the Name field value
func (o *TrackObjectAlbum) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TrackObjectAlbum) SetName(v string) {
	o.Name = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *TrackObjectAlbum) GetReleaseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *TrackObjectAlbum) SetReleaseDate(v string) {
	o.ReleaseDate = v
}

// GetReleaseDatePrecision returns the ReleaseDatePrecision field value
func (o *TrackObjectAlbum) GetReleaseDatePrecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDatePrecision
}

// GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetReleaseDatePrecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDatePrecision, true
}

// SetReleaseDatePrecision sets field value
func (o *TrackObjectAlbum) SetReleaseDatePrecision(v string) {
	o.ReleaseDatePrecision = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *TrackObjectAlbum) GetRestrictions() AlbumBaseRestrictions {
	if o == nil || IsNil(o.Restrictions) {
		var ret AlbumBaseRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetRestrictionsOk() (*AlbumBaseRestrictions, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *TrackObjectAlbum) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given AlbumBaseRestrictions and assigns it to the Restrictions field.
func (o *TrackObjectAlbum) SetRestrictions(v AlbumBaseRestrictions) {
	o.Restrictions = &v
}

// GetType returns the Type field value
func (o *TrackObjectAlbum) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TrackObjectAlbum) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *TrackObjectAlbum) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *TrackObjectAlbum) SetUri(v string) {
	o.Uri = v
}

// GetAlbumGroup returns the AlbumGroup field value if set, zero value otherwise.
func (o *TrackObjectAlbum) GetAlbumGroup() string {
	if o == nil || IsNil(o.AlbumGroup) {
		var ret string
		return ret
	}
	return *o.AlbumGroup
}

// GetAlbumGroupOk returns a tuple with the AlbumGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetAlbumGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AlbumGroup) {
		return nil, false
	}
	return o.AlbumGroup, true
}

// HasAlbumGroup returns a boolean if a field has been set.
func (o *TrackObjectAlbum) HasAlbumGroup() bool {
	if o != nil && !IsNil(o.AlbumGroup) {
		return true
	}

	return false
}

// SetAlbumGroup gets a reference to the given string and assigns it to the AlbumGroup field.
func (o *TrackObjectAlbum) SetAlbumGroup(v string) {
	o.AlbumGroup = &v
}

// GetArtists returns the Artists field value
func (o *TrackObjectAlbum) GetArtists() []SimplifiedArtistObject {
	if o == nil {
		var ret []SimplifiedArtistObject
		return ret
	}

	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value
// and a boolean to check if the value has been set.
func (o *TrackObjectAlbum) GetArtistsOk() ([]SimplifiedArtistObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artists, true
}

// SetArtists sets field value
func (o *TrackObjectAlbum) SetArtists(v []SimplifiedArtistObject) {
	o.Artists = v
}

func (o TrackObjectAlbum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackObjectAlbum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["album_type"] = o.AlbumType
	toSerialize["total_tracks"] = o.TotalTracks
	toSerialize["available_markets"] = o.AvailableMarkets
	toSerialize["external_urls"] = o.ExternalUrls
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	toSerialize["name"] = o.Name
	toSerialize["release_date"] = o.ReleaseDate
	toSerialize["release_date_precision"] = o.ReleaseDatePrecision
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	if !IsNil(o.AlbumGroup) {
		toSerialize["album_group"] = o.AlbumGroup
	}
	toSerialize["artists"] = o.Artists
	return toSerialize, nil
}

type NullableTrackObjectAlbum struct {
	value *TrackObjectAlbum
	isSet bool
}

func (v NullableTrackObjectAlbum) Get() *TrackObjectAlbum {
	return v.value
}

func (v *NullableTrackObjectAlbum) Set(val *TrackObjectAlbum) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackObjectAlbum) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackObjectAlbum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackObjectAlbum(val *TrackObjectAlbum) *NullableTrackObjectAlbum {
	return &NullableTrackObjectAlbum{value: val, isSet: true}
}

func (v NullableTrackObjectAlbum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackObjectAlbum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


