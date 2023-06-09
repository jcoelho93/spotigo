/*
Spotify Web API with fixes and improvements from sonallux

Testing PlaylistsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_PlaylistsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PlaylistsApiService AddTracksToPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsApi.AddTracksToPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService ChangePlaylistDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		httpRes, err := apiClient.PlaylistsApi.ChangePlaylistDetails(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService CheckIfUserFollowsPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsApi.CheckIfUserFollowsPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService CreatePlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.PlaylistsApi.CreatePlaylist(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService FollowPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		httpRes, err := apiClient.PlaylistsApi.FollowPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService GetACategoriesPlaylists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.PlaylistsApi.GetACategoriesPlaylists(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService GetAListOfCurrentUsersPlaylists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PlaylistsApi.GetAListOfCurrentUsersPlaylists(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService GetFeaturedPlaylists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PlaylistsApi.GetFeaturedPlaylists(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService GetListUsersPlaylists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.PlaylistsApi.GetListUsersPlaylists(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService GetPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsApi.GetPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService GetPlaylistCover", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsApi.GetPlaylistCover(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService GetPlaylistsTracks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsApi.GetPlaylistsTracks(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService RemoveTracksPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsApi.RemoveTracksPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService ReorderOrReplacePlaylistsTracks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.PlaylistsApi.ReorderOrReplacePlaylistsTracks(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService UnfollowPlaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		httpRes, err := apiClient.PlaylistsApi.UnfollowPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlaylistsApiService UploadCustomPlaylistCover", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var playlistId string

		httpRes, err := apiClient.PlaylistsApi.UploadCustomPlaylistCover(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}