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

// checks if the DeviceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceObject{}

// DeviceObject struct for DeviceObject
type DeviceObject struct {
	// The device ID.
	Id NullableString `json:"id,omitempty"`
	// If this device is the currently active device.
	IsActive *bool `json:"is_active,omitempty"`
	// If this device is currently in a private session.
	IsPrivateSession *bool `json:"is_private_session,omitempty"`
	// Whether controlling this device is restricted. At present if this is \"true\" then no Web API commands will be accepted by this device.
	IsRestricted *bool `json:"is_restricted,omitempty"`
	// A human-readable name for the device. Some devices have a name that the user can configure (e.g. \\\"Loudest speaker\\\") and some devices have a generic name associated with the manufacturer or device model.
	Name *string `json:"name,omitempty"`
	// Device type, such as \"computer\", \"smartphone\" or \"speaker\".
	Type *string `json:"type,omitempty"`
	// The current volume in percent.
	VolumePercent NullableInt32 `json:"volume_percent,omitempty"`
}

// NewDeviceObject instantiates a new DeviceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceObject() *DeviceObject {
	this := DeviceObject{}
	return &this
}

// NewDeviceObjectWithDefaults instantiates a new DeviceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceObjectWithDefaults() *DeviceObject {
	this := DeviceObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceObject) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DeviceObject) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DeviceObject) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DeviceObject) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DeviceObject) UnsetId() {
	o.Id.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *DeviceObject) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceObject) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *DeviceObject) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *DeviceObject) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsPrivateSession returns the IsPrivateSession field value if set, zero value otherwise.
func (o *DeviceObject) GetIsPrivateSession() bool {
	if o == nil || IsNil(o.IsPrivateSession) {
		var ret bool
		return ret
	}
	return *o.IsPrivateSession
}

// GetIsPrivateSessionOk returns a tuple with the IsPrivateSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceObject) GetIsPrivateSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivateSession) {
		return nil, false
	}
	return o.IsPrivateSession, true
}

// HasIsPrivateSession returns a boolean if a field has been set.
func (o *DeviceObject) HasIsPrivateSession() bool {
	if o != nil && !IsNil(o.IsPrivateSession) {
		return true
	}

	return false
}

// SetIsPrivateSession gets a reference to the given bool and assigns it to the IsPrivateSession field.
func (o *DeviceObject) SetIsPrivateSession(v bool) {
	o.IsPrivateSession = &v
}

// GetIsRestricted returns the IsRestricted field value if set, zero value otherwise.
func (o *DeviceObject) GetIsRestricted() bool {
	if o == nil || IsNil(o.IsRestricted) {
		var ret bool
		return ret
	}
	return *o.IsRestricted
}

// GetIsRestrictedOk returns a tuple with the IsRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceObject) GetIsRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRestricted) {
		return nil, false
	}
	return o.IsRestricted, true
}

// HasIsRestricted returns a boolean if a field has been set.
func (o *DeviceObject) HasIsRestricted() bool {
	if o != nil && !IsNil(o.IsRestricted) {
		return true
	}

	return false
}

// SetIsRestricted gets a reference to the given bool and assigns it to the IsRestricted field.
func (o *DeviceObject) SetIsRestricted(v bool) {
	o.IsRestricted = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceObject) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeviceObject) SetType(v string) {
	o.Type = &v
}

// GetVolumePercent returns the VolumePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceObject) GetVolumePercent() int32 {
	if o == nil || IsNil(o.VolumePercent.Get()) {
		var ret int32
		return ret
	}
	return *o.VolumePercent.Get()
}

// GetVolumePercentOk returns a tuple with the VolumePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceObject) GetVolumePercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumePercent.Get(), o.VolumePercent.IsSet()
}

// HasVolumePercent returns a boolean if a field has been set.
func (o *DeviceObject) HasVolumePercent() bool {
	if o != nil && o.VolumePercent.IsSet() {
		return true
	}

	return false
}

// SetVolumePercent gets a reference to the given NullableInt32 and assigns it to the VolumePercent field.
func (o *DeviceObject) SetVolumePercent(v int32) {
	o.VolumePercent.Set(&v)
}
// SetVolumePercentNil sets the value for VolumePercent to be an explicit nil
func (o *DeviceObject) SetVolumePercentNil() {
	o.VolumePercent.Set(nil)
}

// UnsetVolumePercent ensures that no value is present for VolumePercent, not even an explicit nil
func (o *DeviceObject) UnsetVolumePercent() {
	o.VolumePercent.Unset()
}

func (o DeviceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsPrivateSession) {
		toSerialize["is_private_session"] = o.IsPrivateSession
	}
	if !IsNil(o.IsRestricted) {
		toSerialize["is_restricted"] = o.IsRestricted
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.VolumePercent.IsSet() {
		toSerialize["volume_percent"] = o.VolumePercent.Get()
	}
	return toSerialize, nil
}

type NullableDeviceObject struct {
	value *DeviceObject
	isSet bool
}

func (v NullableDeviceObject) Get() *DeviceObject {
	return v.value
}

func (v *NullableDeviceObject) Set(val *DeviceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceObject(val *DeviceObject) *NullableDeviceObject {
	return &NullableDeviceObject{value: val, isSet: true}
}

func (v NullableDeviceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


