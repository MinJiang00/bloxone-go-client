/*
DNS Configuration API

Testing AclAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dns_config_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/internal"
	openapiclient "github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func Test_dns_config_ConvertRnameAPIService(t *testing.T) {

	t.Run("Test ConvertRnameAPIService ConvertRnameConvertRName", func(t *testing.T) {

		// Define a dummy email address
		emailAddress := "test@example.com"

		// Create a mock HTTP client
		testClient := NewTestClient(func(req *http.Request) *http.Response {
			// Check the request parameters
			require.Equal(t, "GET", req.Method)
			require.Equal(t, "/api/ddi/v1/dns/convert_rname/"+emailAddress, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			// Create a dummy response
			response := dns_config.ConfigConvertRNameResponse{
				Rname: openapiclient.PtrString("test.example.com."),
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			// Return the dummy response
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		configuration := internal.NewConfiguration()
		configuration.HTTPClient = testClient

		convertRename := dns_config.NewAPIClient(configuration)
		ctx := context.Background()

		resp, httpRes, err := convertRename.ConvertRnameAPI.ConvertRnameConvertRName(ctx, emailAddress).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "test.example.com.", *resp.Rname)

	})

}
