# CursorPagingObjectCursors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The cursor to use as key to find the next page of items. | [optional] 
**Before** | Pointer to **string** | The cursor to use as key to find the previous page of items. | [optional] 

## Methods

### NewCursorPagingObjectCursors

`func NewCursorPagingObjectCursors() *CursorPagingObjectCursors`

NewCursorPagingObjectCursors instantiates a new CursorPagingObjectCursors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPagingObjectCursorsWithDefaults

`func NewCursorPagingObjectCursorsWithDefaults() *CursorPagingObjectCursors`

NewCursorPagingObjectCursorsWithDefaults instantiates a new CursorPagingObjectCursors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *CursorPagingObjectCursors) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CursorPagingObjectCursors) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CursorPagingObjectCursors) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *CursorPagingObjectCursors) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *CursorPagingObjectCursors) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *CursorPagingObjectCursors) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *CursorPagingObjectCursors) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *CursorPagingObjectCursors) HasBefore() bool`

HasBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


