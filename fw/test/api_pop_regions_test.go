/*
BloxOne FW API

Testing PopRegionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fw

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/bloxone-go-client/internal"
)

func Test_fw_PopRegionsAPIService(t *testing.T) {

	configuration := internal.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PopRegionsAPIService PopRegionsListPoPRegions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PopRegionsAPI.PopRegionsListPoPRegions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PopRegionsAPIService PopRegionsReadPoPRegion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.PopRegionsAPI.PopRegionsReadPoPRegion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}