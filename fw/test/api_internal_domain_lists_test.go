/*
BloxOne FW API

Testing InternalDomainListsAPIService

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

func TestInternalDomainListsAPIService(t *testing.T) {

	apiClient := fw.NewAPIClient()

	t.Run("Test InternalDomainListsAPIService CreateInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InternalDomainListsAPI.CreateInternalDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService DeleteInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.InternalDomainListsAPI.DeleteInternalDomains(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService DeleteSingleInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.InternalDomainListsAPI.DeleteSingleInternalDomains(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService InternalDomainsItemsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InternalDomainListsAPI.InternalDomainsItemsPartialUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService ListInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InternalDomainListsAPI.ListInternalDomains(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService ReadInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InternalDomainListsAPI.ReadInternalDomains(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InternalDomainListsAPIService UpdateInternalDomains", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.InternalDomainListsAPI.UpdateInternalDomains(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
