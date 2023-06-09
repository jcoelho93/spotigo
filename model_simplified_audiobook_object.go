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

// checks if the SimplifiedAudiobookObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedAudiobookObject{}

// SimplifiedAudiobookObject struct for SimplifiedAudiobookObject
type SimplifiedAudiobookObject struct {
	// The author(s) for the audiobook. 
	Authors []AuthorObject `json:"authors"`
	// A list of the countries in which the audiobook can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. 
	AvailableMarkets []string `json:"available_markets"`
	// The copyright statements of the audiobook. 
	Copyrights []CopyrightObject `json:"copyrights"`
	// A description of the audiobook. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed. 
	Description string `json:"description"`
	// A description of the audiobook. This field may contain HTML tags. 
	HtmlDescription string `json:"html_description"`
	// The edition of the audiobook. 
	Edition *string `json:"edition,omitempty"`
	// Whether or not the audiobook has explicit content (true = yes it does; false = no it does not OR unknown). 
	Explicit bool `json:"explicit"`
	ExternalUrls AudiobookBaseExternalUrls `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the audiobook. 
	Href string `json:"href"`
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook. 
	Id string `json:"id"`
	// The cover art for the audiobook in various sizes, widest first. 
	Images []ImageObject `json:"images"`
	// A list of the languages used in the audiobook, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. 
	Languages []string `json:"languages"`
	// The media type of the audiobook. 
	MediaType string `json:"media_type"`
	// The name of the audiobook. 
	Name string `json:"name"`
	// The narrator(s) for the audiobook. 
	Narrators []NarratorObject `json:"narrators"`
	// The publisher of the audiobook. 
	Publisher string `json:"publisher"`
	// The object type. 
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook. 
	Uri string `json:"uri"`
	// The number of chapters in this audiobook. 
	TotalChapters int32 `json:"total_chapters"`
}

// NewSimplifiedAudiobookObject instantiates a new SimplifiedAudiobookObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedAudiobookObject(authors []AuthorObject, availableMarkets []string, copyrights []CopyrightObject, description string, htmlDescription string, explicit bool, externalUrls AudiobookBaseExternalUrls, href string, id string, images []ImageObject, languages []string, mediaType string, name string, narrators []NarratorObject, publisher string, type_ string, uri string, totalChapters int32) *SimplifiedAudiobookObject {
	this := SimplifiedAudiobookObject{}
	this.Authors = authors
	this.AvailableMarkets = availableMarkets
	this.Copyrights = copyrights
	this.Description = description
	this.HtmlDescription = htmlDescription
	this.Explicit = explicit
	this.ExternalUrls = externalUrls
	this.Href = href
	this.Id = id
	this.Images = images
	this.Languages = languages
	this.MediaType = mediaType
	this.Name = name
	this.Narrators = narrators
	this.Publisher = publisher
	this.Type = type_
	this.Uri = uri
	this.TotalChapters = totalChapters
	return &this
}

// NewSimplifiedAudiobookObjectWithDefaults instantiates a new SimplifiedAudiobookObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedAudiobookObjectWithDefaults() *SimplifiedAudiobookObject {
	this := SimplifiedAudiobookObject{}
	return &this
}

// GetAuthors returns the Authors field value
func (o *SimplifiedAudiobookObject) GetAuthors() []AuthorObject {
	if o == nil {
		var ret []AuthorObject
		return ret
	}

	return o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetAuthorsOk() ([]AuthorObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authors, true
}

// SetAuthors sets field value
func (o *SimplifiedAudiobookObject) SetAuthors(v []AuthorObject) {
	o.Authors = v
}

// GetAvailableMarkets returns the AvailableMarkets field value
func (o *SimplifiedAudiobookObject) GetAvailableMarkets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableMarkets
}

// GetAvailableMarketsOk returns a tuple with the AvailableMarkets field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetAvailableMarketsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableMarkets, true
}

// SetAvailableMarkets sets field value
func (o *SimplifiedAudiobookObject) SetAvailableMarkets(v []string) {
	o.AvailableMarkets = v
}

// GetCopyrights returns the Copyrights field value
func (o *SimplifiedAudiobookObject) GetCopyrights() []CopyrightObject {
	if o == nil {
		var ret []CopyrightObject
		return ret
	}

	return o.Copyrights
}

// GetCopyrightsOk returns a tuple with the Copyrights field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetCopyrightsOk() ([]CopyrightObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Copyrights, true
}

// SetCopyrights sets field value
func (o *SimplifiedAudiobookObject) SetCopyrights(v []CopyrightObject) {
	o.Copyrights = v
}

// GetDescription returns the Description field value
func (o *SimplifiedAudiobookObject) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SimplifiedAudiobookObject) SetDescription(v string) {
	o.Description = v
}

// GetHtmlDescription returns the HtmlDescription field value
func (o *SimplifiedAudiobookObject) GetHtmlDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlDescription
}

// GetHtmlDescriptionOk returns a tuple with the HtmlDescription field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetHtmlDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlDescription, true
}

// SetHtmlDescription sets field value
func (o *SimplifiedAudiobookObject) SetHtmlDescription(v string) {
	o.HtmlDescription = v
}

// GetEdition returns the Edition field value if set, zero value otherwise.
func (o *SimplifiedAudiobookObject) GetEdition() string {
	if o == nil || IsNil(o.Edition) {
		var ret string
		return ret
	}
	return *o.Edition
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetEditionOk() (*string, bool) {
	if o == nil || IsNil(o.Edition) {
		return nil, false
	}
	return o.Edition, true
}

// HasEdition returns a boolean if a field has been set.
func (o *SimplifiedAudiobookObject) HasEdition() bool {
	if o != nil && !IsNil(o.Edition) {
		return true
	}

	return false
}

// SetEdition gets a reference to the given string and assigns it to the Edition field.
func (o *SimplifiedAudiobookObject) SetEdition(v string) {
	o.Edition = &v
}

// GetExplicit returns the Explicit field value
func (o *SimplifiedAudiobookObject) GetExplicit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetExplicitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explicit, true
}

// SetExplicit sets field value
func (o *SimplifiedAudiobookObject) SetExplicit(v bool) {
	o.Explicit = v
}

// GetExternalUrls returns the ExternalUrls field value
func (o *SimplifiedAudiobookObject) GetExternalUrls() AudiobookBaseExternalUrls {
	if o == nil {
		var ret AudiobookBaseExternalUrls
		return ret
	}

	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetExternalUrlsOk() (*AudiobookBaseExternalUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrls, true
}

// SetExternalUrls sets field value
func (o *SimplifiedAudiobookObject) SetExternalUrls(v AudiobookBaseExternalUrls) {
	o.ExternalUrls = v
}

// GetHref returns the Href field value
func (o *SimplifiedAudiobookObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *SimplifiedAudiobookObject) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *SimplifiedAudiobookObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplifiedAudiobookObject) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *SimplifiedAudiobookObject) GetImages() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetImagesOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *SimplifiedAudiobookObject) SetImages(v []ImageObject) {
	o.Images = v
}

// GetLanguages returns the Languages field value
func (o *SimplifiedAudiobookObject) GetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Languages, true
}

// SetLanguages sets field value
func (o *SimplifiedAudiobookObject) SetLanguages(v []string) {
	o.Languages = v
}

// GetMediaType returns the MediaType field value
func (o *SimplifiedAudiobookObject) GetMediaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *SimplifiedAudiobookObject) SetMediaType(v string) {
	o.MediaType = v
}

// GetName returns the Name field value
func (o *SimplifiedAudiobookObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SimplifiedAudiobookObject) SetName(v string) {
	o.Name = v
}

// GetNarrators returns the Narrators field value
func (o *SimplifiedAudiobookObject) GetNarrators() []NarratorObject {
	if o == nil {
		var ret []NarratorObject
		return ret
	}

	return o.Narrators
}

// GetNarratorsOk returns a tuple with the Narrators field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetNarratorsOk() ([]NarratorObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Narrators, true
}

// SetNarrators sets field value
func (o *SimplifiedAudiobookObject) SetNarrators(v []NarratorObject) {
	o.Narrators = v
}

// GetPublisher returns the Publisher field value
func (o *SimplifiedAudiobookObject) GetPublisher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetPublisherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Publisher, true
}

// SetPublisher sets field value
func (o *SimplifiedAudiobookObject) SetPublisher(v string) {
	o.Publisher = v
}

// GetType returns the Type field value
func (o *SimplifiedAudiobookObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SimplifiedAudiobookObject) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *SimplifiedAudiobookObject) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *SimplifiedAudiobookObject) SetUri(v string) {
	o.Uri = v
}

// GetTotalChapters returns the TotalChapters field value
func (o *SimplifiedAudiobookObject) GetTotalChapters() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalChapters
}

// GetTotalChaptersOk returns a tuple with the TotalChapters field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAudiobookObject) GetTotalChaptersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalChapters, true
}

// SetTotalChapters sets field value
func (o *SimplifiedAudiobookObject) SetTotalChapters(v int32) {
	o.TotalChapters = v
}

func (o SimplifiedAudiobookObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedAudiobookObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authors"] = o.Authors
	toSerialize["available_markets"] = o.AvailableMarkets
	toSerialize["copyrights"] = o.Copyrights
	toSerialize["description"] = o.Description
	toSerialize["html_description"] = o.HtmlDescription
	if !IsNil(o.Edition) {
		toSerialize["edition"] = o.Edition
	}
	toSerialize["explicit"] = o.Explicit
	toSerialize["external_urls"] = o.ExternalUrls
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	toSerialize["languages"] = o.Languages
	toSerialize["media_type"] = o.MediaType
	toSerialize["name"] = o.Name
	toSerialize["narrators"] = o.Narrators
	toSerialize["publisher"] = o.Publisher
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	toSerialize["total_chapters"] = o.TotalChapters
	return toSerialize, nil
}

type NullableSimplifiedAudiobookObject struct {
	value *SimplifiedAudiobookObject
	isSet bool
}

func (v NullableSimplifiedAudiobookObject) Get() *SimplifiedAudiobookObject {
	return v.value
}

func (v *NullableSimplifiedAudiobookObject) Set(val *SimplifiedAudiobookObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedAudiobookObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedAudiobookObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedAudiobookObject(val *SimplifiedAudiobookObject) *NullableSimplifiedAudiobookObject {
	return &NullableSimplifiedAudiobookObject{value: val, isSet: true}
}

func (v NullableSimplifiedAudiobookObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedAudiobookObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


