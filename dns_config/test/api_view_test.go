/*
DNS Configuration API

Testing ViewAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dns_config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/internal"
)

func Test_dns_config_ViewAPIService(t *testing.T) {

	configuration := internal.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ViewAPIService ViewBulkCopy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ViewAPI.ViewBulkCopy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ViewAPIService ViewCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ViewAPI.ViewCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ViewAPIService ViewDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.ViewAPI.ViewDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ViewAPIService ViewList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ViewAPI.ViewList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ViewAPIService ViewRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ViewAPI.ViewRead(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ViewAPIService ViewUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.ViewAPI.ViewUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}