// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/azure/azure-dev/cli/azd/cmd/actions"
	"github.com/azure/azure-dev/cli/azd/internal"
	"github.com/azure/azure-dev/cli/azd/pkg/account"
	"github.com/azure/azure-dev/cli/azd/pkg/alpha"
	"github.com/azure/azure-dev/cli/azd/pkg/apphost"
	"github.com/azure/azure-dev/cli/azd/pkg/azapi"
	"github.com/azure/azure-dev/cli/azd/pkg/azure"
	"github.com/azure/azure-dev/cli/azd/pkg/cloud"
	"github.com/azure/azure-dev/cli/azd/pkg/environment"
	"github.com/azure/azure-dev/cli/azd/pkg/environment/azdcontext"
	"github.com/azure/azure-dev/cli/azd/pkg/infra"
	"github.com/azure/azure-dev/cli/azd/pkg/input"
	"github.com/azure/azure-dev/cli/azd/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type monitorFlags struct {
	monitorLive     bool
	monitorLogs     bool
	monitorOverview bool
	global          *internal.GlobalCommandOptions
	internal.EnvFlag
}

func (m *monitorFlags) Bind(local *pflag.FlagSet, global *internal.GlobalCommandOptions) {
	local.BoolVar(
		&m.monitorLive,
		"live",
		false,
		"Open a browser to Application Insights Live Metrics. Live Metrics is currently not supported for Python apps.",
	)
	local.BoolVar(&m.monitorLogs, "logs", false, "Open a browser to Application Insights Logs.")
	local.BoolVar(&m.monitorOverview, "overview", false, "Open a browser to Application Insights Overview Dashboard.")
	m.EnvFlag.Bind(local, global)
	m.global = global
}

func newMonitorFlags(cmd *cobra.Command, global *internal.GlobalCommandOptions) *monitorFlags {
	flags := &monitorFlags{}
	flags.Bind(cmd.Flags(), global)

	return flags
}

func newMonitorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "monitor",
		Short: "Monitor a deployed project.",
	}
}

type monitorAction struct {
	azdCtx               *azdcontext.AzdContext
	env                  *environment.Environment
	subResolver          account.SubscriptionTenantResolver
	resourceManager      infra.ResourceManager
	resourceService      *azapi.ResourceService
	console              input.Console
	flags                *monitorFlags
	portalUrlBase        string
	alphaFeaturesManager *alpha.FeatureManager
}

func newMonitorAction(
	azdCtx *azdcontext.AzdContext,
	env *environment.Environment,
	subResolver account.SubscriptionTenantResolver,
	resourceManager infra.ResourceManager,
	resourceService *azapi.ResourceService,
	console input.Console,
	flags *monitorFlags,
	cloud *cloud.Cloud,
	alphaFeatureManager *alpha.FeatureManager,
) actions.Action {
	return &monitorAction{
		azdCtx:               azdCtx,
		env:                  env,
		resourceManager:      resourceManager,
		resourceService:      resourceService,
		console:              console,
		flags:                flags,
		subResolver:          subResolver,
		portalUrlBase:        cloud.PortalUrlBase,
		alphaFeaturesManager: alphaFeatureManager,
	}
}

func (m *monitorAction) Run(ctx context.Context) (*actions.ActionResult, error) {
	if !m.flags.monitorLive && !m.flags.monitorLogs && !m.flags.monitorOverview {
		m.flags.monitorOverview = true
	}

	if m.env.GetSubscriptionId() == "" {
		return nil, errors.New(
			"infrastructure has not been provisioned. Run `azd provision`",
		)
	}

	aspireDashboard := apphost.AspireDashboardUrl(ctx, m.env, m.alphaFeaturesManager)
	if aspireDashboard != nil {
		openWithDefaultBrowser(ctx, m.console, aspireDashboard.Link)
		return nil, nil
	}

	resourceGroups, err := m.resourceManager.GetResourceGroupsForEnvironment(ctx, m.env.GetSubscriptionId(), m.env.Name())
	if err != nil {
		return nil, fmt.Errorf("discovering resource groups from deployment: %w", err)
	}

	var insightsResources []*azapi.ResourceExtended
	var portalResources []*azapi.ResourceExtended

	for _, resourceGroup := range resourceGroups {
		resources, err := m.resourceService.ListResourceGroupResources(
			ctx, azure.SubscriptionFromRID(resourceGroup.Id), resourceGroup.Name, nil)
		if err != nil {
			return nil, fmt.Errorf("listing resources: %w", err)
		}

		for _, resource := range resources {
			switch resource.Type {
			case string(azapi.AzureResourceTypePortalDashboard):
				portalResources = append(portalResources, resource)
			case string(azapi.AzureResourceTypeAppInsightComponent):
				insightsResources = append(insightsResources, resource)
			}
		}
	}

	if len(insightsResources) == 0 && (m.flags.monitorLive || m.flags.monitorLogs) {
		return nil, fmt.Errorf("application does not contain an Application Insights resource")
	}

	if len(portalResources) == 0 && m.flags.monitorOverview {
		return nil, fmt.Errorf("application does not contain an Application Insights dashboard")
	}

	tenantId, err := m.subResolver.LookupTenant(ctx, m.env.GetSubscriptionId())
	if err != nil {
		return nil, err
	}

	for _, insightsResource := range insightsResources {
		if m.flags.monitorLive {
			openWithDefaultBrowser(ctx, m.console,
				fmt.Sprintf("%s/#@%s/resource%s/quickPulse", m.portalUrlBase, tenantId, insightsResource.Id))
		}

		if m.flags.monitorLogs {
			openWithDefaultBrowser(ctx, m.console,
				fmt.Sprintf("%s/#@%s/resource%s/logs", m.portalUrlBase, tenantId, insightsResource.Id))
		}
	}

	for _, portalResource := range portalResources {
		if m.flags.monitorOverview {
			openWithDefaultBrowser(ctx, m.console,
				fmt.Sprintf("%s/#@%s/dashboard/arm%s", m.portalUrlBase, tenantId, portalResource.Id))
		}
	}

	return nil, nil
}

func getCmdMonitorHelpDescription(*cobra.Command) string {
	return generateCmdHelpDescription(
		fmt.Sprintf("Monitor a deployed application %s. For more information, go to: %s.",
			output.WithWarningFormat("(Beta)"),
			output.WithLinkFormat("https://aka.ms/azure-dev/monitor")), nil)
}

func getCmdMonitorHelpFooter(c *cobra.Command) string {
	return generateCmdHelpSamplesBlock(map[string]string{
		"Open Application Insights Overview Dashboard.": output.WithHighLightFormat("azd monitor --overview"),
		"Open Application Insights Live Metrics.":       output.WithHighLightFormat("azd monitor --live"),
		"Open Application Insights Logs.":               output.WithHighLightFormat("azd monitor --logs"),
	})
}
