/*
BloxOne FW API

Testing ApplicationFiltersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fw

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/fw"
)

func TestApplicationFiltersAPIService(t *testing.T) {

	apiClient := fw.NewAPIClient()

	t.Run("Test ApplicationFiltersAPIService CreateApplicationFilter", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationFiltersAPI.CreateApplicationFilter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFiltersAPIService DeleteApplicationFilters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ApplicationFiltersAPI.DeleteApplicationFilters(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFiltersAPIService DeleteSingleApplicationFilters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.ApplicationFiltersAPI.DeleteSingleApplicationFilters(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFiltersAPIService ListApplicationFilters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ApplicationFiltersAPI.ListApplicationFilters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFiltersAPIService ReadApplicationFilter", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ApplicationFiltersAPI.ReadApplicationFilter(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationFiltersAPIService UpdateApplicationFilter", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ApplicationFiltersAPI.UpdateApplicationFilter(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
