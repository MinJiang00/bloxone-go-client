/*
IP Address Management API

Testing OptionCodeAPIService

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

var IpamsvcOptionCodePost = openapiclient.IpamsvcOptionCode{
	Id: openapiclient.PtrString("Test Post"),
}
var IpamsvcOptionCodePatch = openapiclient.IpamsvcOptionCode{
	Id: openapiclient.PtrString("Test Patch"),
}

func TestOptionCodeAPIService(t *testing.T) {

	t.Run("Test OptionCodeAPIService OptionCodeCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_code", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.IpamsvcOptionCode
			assert.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			assert.Equal(t, IpamsvcOptionCodePost, reqBody)

			response := openapiclient.IpamsvcCreateOptionCodeResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionCodeAPI.OptionCodeCreate(context.Background()).Body(IpamsvcOptionCodePost).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionCodeAPIService OptionCodeDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodDelete, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_code/"+*IpamsvcOptionCodePost.Id, req.URL.Path)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.OptionCodeAPI.OptionCodeDelete(context.Background(), *IpamsvcOptionCodePost.Id).Execute()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionCodeAPIService OptionCodeList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_code", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.IpamsvcListOptionCodeResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionCodeAPI.OptionCodeList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionCodeAPIService OptionCodeRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_code/"+*IpamsvcOptionCodePost.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.IpamsvcReadOptionCodeResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionCodeAPI.OptionCodeRead(context.Background(), *IpamsvcOptionCodePost.Id).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionCodeAPIService OptionCodeUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_code/"+*IpamsvcOptionCodePatch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.IpamsvcOptionCode
			assert.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			assert.Equal(t, IpamsvcOptionCodePatch, reqBody)

			response := openapiclient.IpamsvcUpdateOptionCodeResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionCodeAPI.OptionCodeUpdate(context.Background(), *IpamsvcOptionCodePatch.Id).Body(IpamsvcOptionCodePatch).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})
}
