// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azsdk

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/azure/azure-dev/cli/azd/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestOverrideUserAgent(t *testing.T) {
	expectedUserAgent := "custom/agent (5.0.0)"

	mockContext := mocks.NewMockContext(context.Background())
	mockContext.HttpClient.When(func(request *http.Request) bool {
		return true
	}).RespondFn(func(request *http.Request) (*http.Response, error) {
		return mocks.CreateEmptyHttpResponse(request, http.StatusOK)
	})

	client, err := armresources.NewClient("SUBSCRIPTION_ID", &mocks.MockCredentials{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			PerCallPolicies: []policy.Policy{NewUserAgentPolicy(expectedUserAgent)},
			Transport:       mockContext.HttpClient,
		},
	})
	require.NoError(t, err)

	var response *http.Response
	ctx := runtime.WithCaptureResponse(*mockContext.Context, &response)

	_, _ = client.GetByID(ctx, "RESOURCE_ID", "", nil)

	// Contains custom user agent
	require.Contains(t, response.Request.Header.Get("User-Agent"), expectedUserAgent)
	// Still contains original user agent
	require.Contains(t, response.Request.Header.Get("User-Agent"), "azsdk-go-armresources")
}
