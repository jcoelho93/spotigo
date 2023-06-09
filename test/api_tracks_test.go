/*
Spotify Web API with fixes and improvements from sonallux

Testing TracksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/jcoelho93/spotigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_TracksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TracksApiService AddTracksToPlaylist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.TracksApi.AddTracksToPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService CheckUsersSavedTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TracksApi.CheckUsersSavedTracks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetAnAlbumsTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.TracksApi.GetAnAlbumsTracks(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetAnArtistsTopTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.TracksApi.GetAnArtistsTopTracks(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetAudioAnalysis", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.TracksApi.GetAudioAnalysis(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetAudioFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.TracksApi.GetAudioFeatures(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetPlaylistsTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.TracksApi.GetPlaylistsTracks(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetRecommendations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TracksApi.GetRecommendations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetSeveralAudioFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TracksApi.GetSeveralAudioFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetSeveralTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TracksApi.GetSeveralTracks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetTrack", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.TracksApi.GetTrack(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetUsersSavedTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TracksApi.GetUsersSavedTracks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService GetUsersTopTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TracksApi.GetUsersTopTracks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService RemoveTracksPlaylist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.TracksApi.RemoveTracksPlaylist(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService RemoveTracksUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TracksApi.RemoveTracksUser(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService ReorderOrReplacePlaylistsTracks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var playlistId string

		resp, httpRes, err := apiClient.TracksApi.ReorderOrReplacePlaylistsTracks(context.Background(), playlistId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TracksApiService SaveTracksUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TracksApi.SaveTracksUser(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
