// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package provisioning_test

import (
	"context"
	"strings"
	"testing"

	"github.com/azure/azure-dev/cli/azd/pkg/environment"
	"github.com/azure/azure-dev/cli/azd/pkg/infra"
	. "github.com/azure/azure-dev/cli/azd/pkg/infra/provisioning"
	"github.com/azure/azure-dev/cli/azd/pkg/infra/provisioning/test"
	"github.com/azure/azure-dev/cli/azd/pkg/input"
	"github.com/azure/azure-dev/cli/azd/test/mocks"
	"github.com/stretchr/testify/require"
)

func TestManagerPlan(t *testing.T) {
	env := environment.EphemeralWithValues("test-env", map[string]string{
		"AZURE_LOCATION": "eastus2",
	})
	options := Options{Provider: "test"}
	interactive := false

	mockContext := mocks.NewMockContext(context.Background())
	test.RegisterTestProvider()
	mgr, _ := NewManager(*mockContext.Context, env, "", options, interactive)

	deploymentPlan, err := mgr.Plan(*mockContext.Context)

	require.NotNil(t, deploymentPlan)
	require.Nil(t, err)
	require.Equal(t, deploymentPlan.Deployment.Parameters["location"].Value, env.Values["AZURE_LOCATION"])
}

func TestManagerGetDeployment(t *testing.T) {
	env := environment.EphemeralWithValues("test-env", map[string]string{
		"AZURE_LOCATION": "eastus2",
	})
	options := Options{Provider: "test"}
	interactive := false

	mockContext := mocks.NewMockContext(context.Background())
	test.RegisterTestProvider()
	mgr, _ := NewManager(*mockContext.Context, env, "", options, interactive)

	provisioningScope := infra.NewSubscriptionScope(*mockContext.Context, "eastus2", env.GetSubscriptionId(), env.GetEnvName())
	getResult, err := mgr.GetDeployment(*mockContext.Context, provisioningScope)

	require.NotNil(t, getResult)
	require.Nil(t, err)
}

func TestManagerDeploy(t *testing.T) {
	env := environment.EphemeralWithValues("test-env", map[string]string{
		"AZURE_LOCATION": "eastus2",
	})
	options := Options{Provider: "test"}
	interactive := false

	mockContext := mocks.NewMockContext(context.Background())
	test.RegisterTestProvider()
	mgr, _ := NewManager(*mockContext.Context, env, "", options, interactive)

	deploymentPlan, _ := mgr.Plan(*mockContext.Context)
	provisioningScope := infra.NewSubscriptionScope(*mockContext.Context, "eastus2", env.GetSubscriptionId(), env.GetEnvName())
	deployResult, err := mgr.Deploy(*mockContext.Context, deploymentPlan, provisioningScope)

	require.NotNil(t, deployResult)
	require.Nil(t, err)
}

func TestManagerDestroyWithPositiveConfirmation(t *testing.T) {
	env := environment.EphemeralWithValues("test-env", map[string]string{
		"AZURE_LOCATION": "eastus2",
	})
	options := Options{Provider: "test"}
	interactive := false

	mockContext := mocks.NewMockContext(context.Background())
	mockContext.Console.WhenConfirm(func(options input.ConsoleOptions) bool {
		return strings.Contains(options.Message, "Are you sure you want to destroy?")
	}).Respond(true)

	test.RegisterTestProvider()
	mgr, _ := NewManager(*mockContext.Context, env, "", options, interactive)

	deploymentPlan, _ := mgr.Plan(*mockContext.Context)
	destroyOptions := NewDestroyOptions(false, false)
	destroyResult, err := mgr.Destroy(*mockContext.Context, &deploymentPlan.Deployment, destroyOptions)

	require.NotNil(t, destroyResult)
	require.Nil(t, err)
	require.Contains(t, mockContext.Console.Output(), "Are you sure you want to destroy?")
}

func TestManagerDestroyWithNegativeConfirmation(t *testing.T) {
	env := environment.EphemeralWithValues("test-env", map[string]string{
		"AZURE_LOCATION": "eastus2",
	})
	options := Options{Provider: "test"}
	interactive := false

	mockContext := mocks.NewMockContext(context.Background())
	mockContext.Console.WhenConfirm(func(options input.ConsoleOptions) bool {
		return strings.Contains(options.Message, "Are you sure you want to destroy?")
	}).Respond(false)

	test.RegisterTestProvider()
	mgr, _ := NewManager(*mockContext.Context, env, "", options, interactive)

	deploymentPlan, _ := mgr.Plan(*mockContext.Context)
	destroyOptions := NewDestroyOptions(false, false)
	destroyResult, err := mgr.Destroy(*mockContext.Context, &deploymentPlan.Deployment, destroyOptions)

	require.Nil(t, destroyResult)
	require.NotNil(t, err)
	require.Contains(t, mockContext.Console.Output(), "Are you sure you want to destroy?")
}