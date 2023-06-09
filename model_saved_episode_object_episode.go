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

// checks if the SavedEpisodeObjectEpisode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedEpisodeObjectEpisode{}

// SavedEpisodeObjectEpisode Information about the episode.
type SavedEpisodeObjectEpisode struct {
	// A URL to a 30 second preview (MP3 format) of the episode. `null` if not available. 
	AudioPreviewUrl string `json:"audio_preview_url"`
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
	// True if the episode is hosted outside of Spotify's CDN. 
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// True if the episode is playable in the given market. Otherwise false. 
	IsPlayable bool `json:"is_playable"`
	// The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the `languages` field instead. 
	// Deprecated
	Language *string `json:"language,omitempty"`
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
	Restrictions *EpisodeBaseRestrictions `json:"restrictions,omitempty"`
	Show EpisodeObjectAllOfShow `json:"show"`
}

// NewSavedEpisodeObjectEpisode instantiates a new SavedEpisodeObjectEpisode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedEpisodeObjectEpisode(audioPreviewUrl string, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls EpisodeBaseExternalUrls, href string, id string, images []ImageObject, isExternallyHosted bool, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, resumePoint EpisodeBaseResumePoint, type_ string, uri string, show EpisodeObjectAllOfShow) *SavedEpisodeObjectEpisode {
	this := SavedEpisodeObjectEpisode{}
	this.AudioPreviewUrl = audioPreviewUrl
	this.Description = description
	this.HtmlDescription = htmlDescription
	this.DurationMs = durationMs
	this.Explicit = explicit
	this.ExternalUrls = externalUrls
	this.Href = href
	this.Id = id
	this.Images = images
	this.IsExternallyHosted = isExternallyHosted
	this.IsPlayable = isPlayable
	this.Languages = languages
	this.Name = name
	this.ReleaseDate = releaseDate
	this.ReleaseDatePrecision = releaseDatePrecision
	this.ResumePoint = resumePoint
	this.Type = type_
	this.Uri = uri
	this.Show = show
	return &this
}

// NewSavedEpisodeObjectEpisodeWithDefaults instantiates a new SavedEpisodeObjectEpisode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedEpisodeObjectEpisodeWithDefaults() *SavedEpisodeObjectEpisode {
	this := SavedEpisodeObjectEpisode{}
	return &this
}

// GetAudioPreviewUrl returns the AudioPreviewUrl field value
func (o *SavedEpisodeObjectEpisode) GetAudioPreviewUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudioPreviewUrl
}

// GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetAudioPreviewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudioPreviewUrl, true
}

// SetAudioPreviewUrl sets field value
func (o *SavedEpisodeObjectEpisode) SetAudioPreviewUrl(v string) {
	o.AudioPreviewUrl = v
}

// GetDescription returns the Description field value
func (o *SavedEpisodeObjectEpisode) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SavedEpisodeObjectEpisode) SetDescription(v string) {
	o.Description = v
}

// GetHtmlDescription returns the HtmlDescription field value
func (o *SavedEpisodeObjectEpisode) GetHtmlDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlDescription
}

// GetHtmlDescriptionOk returns a tuple with the HtmlDescription field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetHtmlDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlDescription, true
}

// SetHtmlDescription sets field value
func (o *SavedEpisodeObjectEpisode) SetHtmlDescription(v string) {
	o.HtmlDescription = v
}

// GetDurationMs returns the DurationMs field value
func (o *SavedEpisodeObjectEpisode) GetDurationMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetDurationMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMs, true
}

// SetDurationMs sets field value
func (o *SavedEpisodeObjectEpisode) SetDurationMs(v int32) {
	o.DurationMs = v
}

// GetExplicit returns the Explicit field value
func (o *SavedEpisodeObjectEpisode) GetExplicit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetExplicitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explicit, true
}

// SetExplicit sets field value
func (o *SavedEpisodeObjectEpisode) SetExplicit(v bool) {
	o.Explicit = v
}

// GetExternalUrls returns the ExternalUrls field value
func (o *SavedEpisodeObjectEpisode) GetExternalUrls() EpisodeBaseExternalUrls {
	if o == nil {
		var ret EpisodeBaseExternalUrls
		return ret
	}

	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetExternalUrlsOk() (*EpisodeBaseExternalUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrls, true
}

// SetExternalUrls sets field value
func (o *SavedEpisodeObjectEpisode) SetExternalUrls(v EpisodeBaseExternalUrls) {
	o.ExternalUrls = v
}

// GetHref returns the Href field value
func (o *SavedEpisodeObjectEpisode) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *SavedEpisodeObjectEpisode) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *SavedEpisodeObjectEpisode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SavedEpisodeObjectEpisode) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *SavedEpisodeObjectEpisode) GetImages() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetImagesOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *SavedEpisodeObjectEpisode) SetImages(v []ImageObject) {
	o.Images = v
}

// GetIsExternallyHosted returns the IsExternallyHosted field value
func (o *SavedEpisodeObjectEpisode) GetIsExternallyHosted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExternallyHosted
}

// GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetIsExternallyHostedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExternallyHosted, true
}

// SetIsExternallyHosted sets field value
func (o *SavedEpisodeObjectEpisode) SetIsExternallyHosted(v bool) {
	o.IsExternallyHosted = v
}

// GetIsPlayable returns the IsPlayable field value
func (o *SavedEpisodeObjectEpisode) GetIsPlayable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPlayable
}

// GetIsPlayableOk returns a tuple with the IsPlayable field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetIsPlayableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPlayable, true
}

// SetIsPlayable sets field value
func (o *SavedEpisodeObjectEpisode) SetIsPlayable(v bool) {
	o.IsPlayable = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
// Deprecated
func (o *SavedEpisodeObjectEpisode) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SavedEpisodeObjectEpisode) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *SavedEpisodeObjectEpisode) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
// Deprecated
func (o *SavedEpisodeObjectEpisode) SetLanguage(v string) {
	o.Language = &v
}

// GetLanguages returns the Languages field value
func (o *SavedEpisodeObjectEpisode) GetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Languages, true
}

// SetLanguages sets field value
func (o *SavedEpisodeObjectEpisode) SetLanguages(v []string) {
	o.Languages = v
}

// GetName returns the Name field value
func (o *SavedEpisodeObjectEpisode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SavedEpisodeObjectEpisode) SetName(v string) {
	o.Name = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *SavedEpisodeObjectEpisode) GetReleaseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *SavedEpisodeObjectEpisode) SetReleaseDate(v string) {
	o.ReleaseDate = v
}

// GetReleaseDatePrecision returns the ReleaseDatePrecision field value
func (o *SavedEpisodeObjectEpisode) GetReleaseDatePrecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDatePrecision
}

// GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetReleaseDatePrecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDatePrecision, true
}

// SetReleaseDatePrecision sets field value
func (o *SavedEpisodeObjectEpisode) SetReleaseDatePrecision(v string) {
	o.ReleaseDatePrecision = v
}

// GetResumePoint returns the ResumePoint field value
func (o *SavedEpisodeObjectEpisode) GetResumePoint() EpisodeBaseResumePoint {
	if o == nil {
		var ret EpisodeBaseResumePoint
		return ret
	}

	return o.ResumePoint
}

// GetResumePointOk returns a tuple with the ResumePoint field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetResumePointOk() (*EpisodeBaseResumePoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResumePoint, true
}

// SetResumePoint sets field value
func (o *SavedEpisodeObjectEpisode) SetResumePoint(v EpisodeBaseResumePoint) {
	o.ResumePoint = v
}

// GetType returns the Type field value
func (o *SavedEpisodeObjectEpisode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SavedEpisodeObjectEpisode) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *SavedEpisodeObjectEpisode) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *SavedEpisodeObjectEpisode) SetUri(v string) {
	o.Uri = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *SavedEpisodeObjectEpisode) GetRestrictions() EpisodeBaseRestrictions {
	if o == nil || IsNil(o.Restrictions) {
		var ret EpisodeBaseRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetRestrictionsOk() (*EpisodeBaseRestrictions, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *SavedEpisodeObjectEpisode) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given EpisodeBaseRestrictions and assigns it to the Restrictions field.
func (o *SavedEpisodeObjectEpisode) SetRestrictions(v EpisodeBaseRestrictions) {
	o.Restrictions = &v
}

// GetShow returns the Show field value
func (o *SavedEpisodeObjectEpisode) GetShow() EpisodeObjectAllOfShow {
	if o == nil {
		var ret EpisodeObjectAllOfShow
		return ret
	}

	return o.Show
}

// GetShowOk returns a tuple with the Show field value
// and a boolean to check if the value has been set.
func (o *SavedEpisodeObjectEpisode) GetShowOk() (*EpisodeObjectAllOfShow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Show, true
}

// SetShow sets field value
func (o *SavedEpisodeObjectEpisode) SetShow(v EpisodeObjectAllOfShow) {
	o.Show = v
}

func (o SavedEpisodeObjectEpisode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedEpisodeObjectEpisode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audio_preview_url"] = o.AudioPreviewUrl
	toSerialize["description"] = o.Description
	toSerialize["html_description"] = o.HtmlDescription
	toSerialize["duration_ms"] = o.DurationMs
	toSerialize["explicit"] = o.Explicit
	toSerialize["external_urls"] = o.ExternalUrls
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	toSerialize["is_externally_hosted"] = o.IsExternallyHosted
	toSerialize["is_playable"] = o.IsPlayable
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
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
	toSerialize["show"] = o.Show
	return toSerialize, nil
}

type NullableSavedEpisodeObjectEpisode struct {
	value *SavedEpisodeObjectEpisode
	isSet bool
}

func (v NullableSavedEpisodeObjectEpisode) Get() *SavedEpisodeObjectEpisode {
	return v.value
}

func (v *NullableSavedEpisodeObjectEpisode) Set(val *SavedEpisodeObjectEpisode) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedEpisodeObjectEpisode) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedEpisodeObjectEpisode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedEpisodeObjectEpisode(val *SavedEpisodeObjectEpisode) *NullableSavedEpisodeObjectEpisode {
	return &NullableSavedEpisodeObjectEpisode{value: val, isSet: true}
}

func (v NullableSavedEpisodeObjectEpisode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedEpisodeObjectEpisode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


