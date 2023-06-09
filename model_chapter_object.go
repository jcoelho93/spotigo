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

// checks if the ChapterObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChapterObject{}

// ChapterObject struct for ChapterObject
type ChapterObject struct {
	// A URL to a 30 second preview (MP3 format) of the episode. `null` if not available. 
	AudioPreviewUrl string `json:"audio_preview_url"`
	// A list of the countries in which the chapter can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. 
	AvailableMarkets []string `json:"available_markets,omitempty"`
	// The number of the chapter 
	ChapterNumber int32 `json:"chapter_number"`
	// A description of the episode. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed. 
	Description string `json:"description"`
	// A description of the episode. This field may contain HTML tags. 
	HtmlDescription string `json:"html_description"`
	// The episode length in milliseconds. 
	DurationMs int32 `json:"duration_ms"`
	// Whether or not the episode has explicit content (true = yes it does; false = no it does not OR unknown). 
	Explicit bool `json:"explicit"`
	ExternalUrls EpisodeBaseExternalUrls `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the episode. 
	Href string `json:"href"`
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the episode. 
	Id string `json:"id"`
	// The cover art for the episode in various sizes, widest first. 
	Images []ImageObject `json:"images"`
	// True if the episode is playable in the given market. Otherwise false. 
	IsPlayable bool `json:"is_playable"`
	// A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code. 
	Languages []string `json:"languages"`
	// The name of the episode. 
	Name string `json:"name"`
	// The date the episode was first released, for example `\"1981-12-15\"`. Depending on the precision, it might be shown as `\"1981\"` or `\"1981-12\"`. 
	ReleaseDate string `json:"release_date"`
	// The precision with which `release_date` value is known. 
	ReleaseDatePrecision string `json:"release_date_precision"`
	ResumePoint EpisodeBaseResumePoint `json:"resume_point"`
	// The object type. 
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the episode. 
	Uri string `json:"uri"`
	Restrictions *ChapterBaseRestrictions `json:"restrictions,omitempty"`
	Audiobook ChapterObjectAllOfAudiobook `json:"audiobook"`
}

// NewChapterObject instantiates a new ChapterObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChapterObject(audioPreviewUrl string, chapterNumber int32, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls EpisodeBaseExternalUrls, href string, id string, images []ImageObject, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, resumePoint EpisodeBaseResumePoint, type_ string, uri string, audiobook ChapterObjectAllOfAudiobook) *ChapterObject {
	this := ChapterObject{}
	this.AudioPreviewUrl = audioPreviewUrl
	this.ChapterNumber = chapterNumber
	this.Description = description
	this.HtmlDescription = htmlDescription
	this.DurationMs = durationMs
	this.Explicit = explicit
	this.ExternalUrls = externalUrls
	this.Href = href
	this.Id = id
	this.Images = images
	this.IsPlayable = isPlayable
	this.Languages = languages
	this.Name = name
	this.ReleaseDate = releaseDate
	this.ReleaseDatePrecision = releaseDatePrecision
	this.ResumePoint = resumePoint
	this.Type = type_
	this.Uri = uri
	this.Audiobook = audiobook
	return &this
}

// NewChapterObjectWithDefaults instantiates a new ChapterObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChapterObjectWithDefaults() *ChapterObject {
	this := ChapterObject{}
	return &this
}

// GetAudioPreviewUrl returns the AudioPreviewUrl field value
func (o *ChapterObject) GetAudioPreviewUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudioPreviewUrl
}

// GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetAudioPreviewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudioPreviewUrl, true
}

// SetAudioPreviewUrl sets field value
func (o *ChapterObject) SetAudioPreviewUrl(v string) {
	o.AudioPreviewUrl = v
}

// GetAvailableMarkets returns the AvailableMarkets field value if set, zero value otherwise.
func (o *ChapterObject) GetAvailableMarkets() []string {
	if o == nil || IsNil(o.AvailableMarkets) {
		var ret []string
		return ret
	}
	return o.AvailableMarkets
}

// GetAvailableMarketsOk returns a tuple with the AvailableMarkets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetAvailableMarketsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableMarkets) {
		return nil, false
	}
	return o.AvailableMarkets, true
}

// HasAvailableMarkets returns a boolean if a field has been set.
func (o *ChapterObject) HasAvailableMarkets() bool {
	if o != nil && !IsNil(o.AvailableMarkets) {
		return true
	}

	return false
}

// SetAvailableMarkets gets a reference to the given []string and assigns it to the AvailableMarkets field.
func (o *ChapterObject) SetAvailableMarkets(v []string) {
	o.AvailableMarkets = v
}

// GetChapterNumber returns the ChapterNumber field value
func (o *ChapterObject) GetChapterNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChapterNumber
}

// GetChapterNumberOk returns a tuple with the ChapterNumber field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetChapterNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChapterNumber, true
}

// SetChapterNumber sets field value
func (o *ChapterObject) SetChapterNumber(v int32) {
	o.ChapterNumber = v
}

// GetDescription returns the Description field value
func (o *ChapterObject) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ChapterObject) SetDescription(v string) {
	o.Description = v
}

// GetHtmlDescription returns the HtmlDescription field value
func (o *ChapterObject) GetHtmlDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlDescription
}

// GetHtmlDescriptionOk returns a tuple with the HtmlDescription field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetHtmlDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlDescription, true
}

// SetHtmlDescription sets field value
func (o *ChapterObject) SetHtmlDescription(v string) {
	o.HtmlDescription = v
}

// GetDurationMs returns the DurationMs field value
func (o *ChapterObject) GetDurationMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetDurationMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMs, true
}

// SetDurationMs sets field value
func (o *ChapterObject) SetDurationMs(v int32) {
	o.DurationMs = v
}

// GetExplicit returns the Explicit field value
func (o *ChapterObject) GetExplicit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetExplicitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explicit, true
}

// SetExplicit sets field value
func (o *ChapterObject) SetExplicit(v bool) {
	o.Explicit = v
}

// GetExternalUrls returns the ExternalUrls field value
func (o *ChapterObject) GetExternalUrls() EpisodeBaseExternalUrls {
	if o == nil {
		var ret EpisodeBaseExternalUrls
		return ret
	}

	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetExternalUrlsOk() (*EpisodeBaseExternalUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrls, true
}

// SetExternalUrls sets field value
func (o *ChapterObject) SetExternalUrls(v EpisodeBaseExternalUrls) {
	o.ExternalUrls = v
}

// GetHref returns the Href field value
func (o *ChapterObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ChapterObject) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *ChapterObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChapterObject) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *ChapterObject) GetImages() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetImagesOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *ChapterObject) SetImages(v []ImageObject) {
	o.Images = v
}

// GetIsPlayable returns the IsPlayable field value
func (o *ChapterObject) GetIsPlayable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPlayable
}

// GetIsPlayableOk returns a tuple with the IsPlayable field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetIsPlayableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPlayable, true
}

// SetIsPlayable sets field value
func (o *ChapterObject) SetIsPlayable(v bool) {
	o.IsPlayable = v
}

// GetLanguages returns the Languages field value
func (o *ChapterObject) GetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Languages, true
}

// SetLanguages sets field value
func (o *ChapterObject) SetLanguages(v []string) {
	o.Languages = v
}

// GetName returns the Name field value
func (o *ChapterObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChapterObject) SetName(v string) {
	o.Name = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *ChapterObject) GetReleaseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *ChapterObject) SetReleaseDate(v string) {
	o.ReleaseDate = v
}

// GetReleaseDatePrecision returns the ReleaseDatePrecision field value
func (o *ChapterObject) GetReleaseDatePrecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDatePrecision
}

// GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetReleaseDatePrecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDatePrecision, true
}

// SetReleaseDatePrecision sets field value
func (o *ChapterObject) SetReleaseDatePrecision(v string) {
	o.ReleaseDatePrecision = v
}

// GetResumePoint returns the ResumePoint field value
func (o *ChapterObject) GetResumePoint() EpisodeBaseResumePoint {
	if o == nil {
		var ret EpisodeBaseResumePoint
		return ret
	}

	return o.ResumePoint
}

// GetResumePointOk returns a tuple with the ResumePoint field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetResumePointOk() (*EpisodeBaseResumePoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResumePoint, true
}

// SetResumePoint sets field value
func (o *ChapterObject) SetResumePoint(v EpisodeBaseResumePoint) {
	o.ResumePoint = v
}

// GetType returns the Type field value
func (o *ChapterObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChapterObject) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *ChapterObject) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *ChapterObject) SetUri(v string) {
	o.Uri = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *ChapterObject) GetRestrictions() ChapterBaseRestrictions {
	if o == nil || IsNil(o.Restrictions) {
		var ret ChapterBaseRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetRestrictionsOk() (*ChapterBaseRestrictions, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *ChapterObject) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given ChapterBaseRestrictions and assigns it to the Restrictions field.
func (o *ChapterObject) SetRestrictions(v ChapterBaseRestrictions) {
	o.Restrictions = &v
}

// GetAudiobook returns the Audiobook field value
func (o *ChapterObject) GetAudiobook() ChapterObjectAllOfAudiobook {
	if o == nil {
		var ret ChapterObjectAllOfAudiobook
		return ret
	}

	return o.Audiobook
}

// GetAudiobookOk returns a tuple with the Audiobook field value
// and a boolean to check if the value has been set.
func (o *ChapterObject) GetAudiobookOk() (*ChapterObjectAllOfAudiobook, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audiobook, true
}

// SetAudiobook sets field value
func (o *ChapterObject) SetAudiobook(v ChapterObjectAllOfAudiobook) {
	o.Audiobook = v
}

func (o ChapterObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChapterObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audio_preview_url"] = o.AudioPreviewUrl
	if !IsNil(o.AvailableMarkets) {
		toSerialize["available_markets"] = o.AvailableMarkets
	}
	toSerialize["chapter_number"] = o.ChapterNumber
	toSerialize["description"] = o.Description
	toSerialize["html_description"] = o.HtmlDescription
	toSerialize["duration_ms"] = o.DurationMs
	toSerialize["explicit"] = o.Explicit
	toSerialize["external_urls"] = o.ExternalUrls
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	toSerialize["is_playable"] = o.IsPlayable
	toSerialize["languages"] = o.Languages
	toSerialize["name"] = o.Name
	toSerialize["release_date"] = o.ReleaseDate
	toSerialize["release_date_precision"] = o.ReleaseDatePrecision
	toSerialize["resume_point"] = o.ResumePoint
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	toSerialize["audiobook"] = o.Audiobook
	return toSerialize, nil
}

type NullableChapterObject struct {
	value *ChapterObject
	isSet bool
}

func (v NullableChapterObject) Get() *ChapterObject {
	return v.value
}

func (v *NullableChapterObject) Set(val *ChapterObject) {
	v.value = val
	v.isSet = true
}

func (v NullableChapterObject) IsSet() bool {
	return v.isSet
}

func (v *NullableChapterObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChapterObject(val *ChapterObject) *NullableChapterObject {
	return &NullableChapterObject{value: val, isSet: true}
}

func (v NullableChapterObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChapterObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


