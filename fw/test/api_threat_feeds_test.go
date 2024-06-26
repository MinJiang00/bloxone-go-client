/*
BloxOne FW API

Testing ThreatFeedsAPIService

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

func TestThreatFeedsAPIService(t *testing.T) {

	apiClient := fw.NewAPIClient()

	t.Run("Test ThreatFeedsAPIService ListThreatFeeds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ThreatFeedsAPI.ListThreatFeeds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
