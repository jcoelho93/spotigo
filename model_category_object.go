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

// checks if the CategoryObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryObject{}

// CategoryObject struct for CategoryObject
type CategoryObject struct {
	// A link to the Web API endpoint returning full details of the category. 
	Href string `json:"href"`
	// The category icon, in various sizes. 
	Icons []ImageObject `json:"icons"`
	// The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) of the category. 
	Id string `json:"id"`
	// The name of the category. 
	Name string `json:"name"`
}

// NewCategoryObject instantiates a new CategoryObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryObject(href string, icons []ImageObject, id string, name string) *CategoryObject {
	this := CategoryObject{}
	this.Href = href
	this.Icons = icons
	this.Id = id
	this.Name = name
	return &this
}

// NewCategoryObjectWithDefaults instantiates a new CategoryObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryObjectWithDefaults() *CategoryObject {
	this := CategoryObject{}
	return &this
}

// GetHref returns the Href field value
func (o *CategoryObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *CategoryObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *CategoryObject) SetHref(v string) {
	o.Href = v
}

// GetIcons returns the Icons field value
func (o *CategoryObject) GetIcons() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Icons
}

// GetIconsOk returns a tuple with the Icons field value
// and a boolean to check if the value has been set.
func (o *CategoryObject) GetIconsOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icons, true
}

// SetIcons sets field value
func (o *CategoryObject) SetIcons(v []ImageObject) {
	o.Icons = v
}

// GetId returns the Id field value
func (o *CategoryObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CategoryObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CategoryObject) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CategoryObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CategoryObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CategoryObject) SetName(v string) {
	o.Name = v
}

func (o CategoryObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["icons"] = o.Icons
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCategoryObject struct {
	value *CategoryObject
	isSet bool
}

func (v NullableCategoryObject) Get() *CategoryObject {
	return v.value
}

func (v *NullableCategoryObject) Set(val *CategoryObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryObject(val *CategoryObject) *NullableCategoryObject {
	return &NullableCategoryObject{value: val, isSet: true}
}

func (v NullableCategoryObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


