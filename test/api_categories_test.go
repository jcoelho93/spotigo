/*
Spotify Web API with fixes and improvements from sonallux

Testing CategoriesApiService

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

func Test_openapi_CategoriesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CategoriesApiService GetACategoriesPlaylists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.CategoriesApi.GetACategoriesPlaylists(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesApiService GetACategory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var categoryId string

		resp, httpRes, err := apiClient.CategoriesApi.GetACategory(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesApiService GetCategories", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CategoriesApi.GetCategories(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
