/*
BloxOne Anycast API

Testing OnPremAnycastManagerAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package anycast

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

func TestOnPremAnycastManagerAPIService(t *testing.T) {

	apiClient := anycast.NewAPIClient()

	t.Run("Test OnPremAnycastManagerAPIService CreateAnycastConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.CreateAnycastConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService CreateAnycastVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.CreateAnycastVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService DeleteAnycastConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.DeleteAnycastConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService DeleteAnycastVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.DeleteAnycastVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService DeleteOnpremHost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.DeleteOnpremHost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetAnycastConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetAnycastConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetAnycastConfigList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetAnycastConfigList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetAnycastVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetAnycastVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetOnpremConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ophid string
		var version string

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetOnpremConfig(context.Background(), ophid, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetOnpremConfig2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ophid string
		var version string

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetOnpremConfig2(context.Background(), ophid, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetOnpremHost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetOnpremHost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ophid string

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetStatus(context.Background(), ophid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService GetStatus2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ophid string

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.GetStatus2(context.Background(), ophid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService ListAnycastConfigsWithRuntimeStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.ListAnycastConfigsWithRuntimeStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService ReadAnycastConfigWithRuntimeStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.ReadAnycastConfigWithRuntimeStatus(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService UpdateAnycastConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.UpdateAnycastConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OnPremAnycastManagerAPIService UpdateOnpremHost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int64

		resp, httpRes, err := apiClient.OnPremAnycastManagerAPI.UpdateOnpremHost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
