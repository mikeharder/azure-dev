// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package infra

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/azure/azure-dev/cli/azd/pkg/environment"
)

// ResourceId returns the resource ID for the corresponding name.
//
// If the name is a resource ID string, it is immediately parsed without translation.
func ResourceId(name string, env *environment.Environment) (resId *arm.ResourceID, err error) {
	resId, err = arm.ParseResourceID(name)
	if err == nil {
		return resId, nil
	}

	key := ResourceIdName(name)
	resourceId, ok := env.LookupEnv(key)
	if !ok {
		return resId, fmt.Errorf("%s is not set as an output variable", key)
	}

	if resourceId == "" {
		return resId, fmt.Errorf("%s is empty", key)
	}

	resId, err = arm.ParseResourceID(resourceId)
	if err != nil {
		return resId, fmt.Errorf("parsing %s: %w", key, err)
	}

	return resId, nil
}

// ResourceIdName returns the environment variable name for the resource ID of the given name.
func ResourceIdName(name string) string {
	return fmt.Sprintf("AZURE_RESOURCE_%s_ID", environment.Key(name))
}

// KeyVaultName returns the name of the "canonical" key vault to use for secrets.
func KeyVaultName(env *environment.Environment) string {
	return env.Getenv("AZURE_KEY_VAULT_NAME")
}
