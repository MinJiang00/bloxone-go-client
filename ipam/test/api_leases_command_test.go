/*
IP Address Management API

Testing LeasesCommandAPIService

*/

package ipam_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/internal"
	openapiclient "github.com/infobloxopen/bloxone-go-client/ipam"
)

var IpamsvcLeasesCommandPost = openapiclient.IpamsvcLeasesCommand{}

func TestLeasesCommandAPIService(t *testing.T) {

	t.Run("Test LeasesCommandAPIService LeasesCommandCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/leases_command", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.IpamsvcLeasesCommand
			assert.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			assert.Equal(t, IpamsvcLeasesCommandPost, reqBody)

			response := openapiclient.IpamsvcCreateLeasesCommandResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.LeasesCommandAPI.LeasesCommandCreate(context.Background()).Body(IpamsvcLeasesCommandPost).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})
}
