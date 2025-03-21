//go:build azure
// +build azure

// This file contains unit tests for the client factory implementation(s).

package azure

import (
	"os"
	"reflect"
	"testing"

	autorest "github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Local consts for this file only
const govCloudEnvName = "AzureUSGovernmentCloud"
const publicCloudEnvName = "AzurePublicCloud"
const chinaCloudEnvName = "AzureChinaCloud"
const germanyCloudEnvName = "AzureGermanCloud"

func TestDefaultEnvIsPublicWhenNotSet(t *testing.T) {
	// save any current env value and restore on exit
	originalEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, originalEnv)

	// Set env var to missing value
	os.Setenv(AzureEnvironmentEnvName, "")

	// get the default
	env := getDefaultEnvironmentName()

	// Make sure it's public cloud
	assert.Equal(t, autorest.PublicCloud.Name, env)
}

func TestDefaultEnvSetToGov(t *testing.T) {
	// save any current env value and restore on exit
	originalEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, originalEnv)

	// Set env var to gov
	os.Setenv(AzureEnvironmentEnvName, govCloudEnvName)

	// get the default
	env := getDefaultEnvironmentName()

	// Make sure it's public cloud
	assert.Equal(t, autorest.USGovernmentCloud.Name, env)
}

func TestSubscriptionClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/SubscriptionClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/SubscriptionClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/SubscriptionClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/SubscriptionClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a VM client
			client, err := CreateSubscriptionsClientE()
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}

// snippet-tag-start::client_factory_example.UnitTest

func TestVMClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/VMClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/VMClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/VMClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/VMClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a VM client
			client, err := CreateVirtualMachinesClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}

// snippet-tag-end::client_factory_example.UnitTest

func TestManagedClustersClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/ManagedClustersClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/ManagedClustersClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/ManagedClustersClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/ManagedClustersClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a VM client
			client, err := CreateManagedClustersClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}

func TestCosmosDBAccountClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/CosmosDBAccountClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/CosmosDBAccountClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/CosmosDBAccountClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/CosmosDBAccountClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a VM client
			client, err := CreateCosmosDBAccountClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}

func TestCosmosDBSQLClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/CosmosDBAccountClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/CosmosDBAccountClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/CosmosDBAccountClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/CosmosDBAccountClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a VM client
			client, err := CreateCosmosDBSQLClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}
func TestPublicIPAddressesClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/CosmosDBAccountClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/CosmosDBAccountClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/CosmosDBAccountClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/CosmosDBAccountClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a VM client
			client, err := CreatePublicIPAddressesClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}
func TestLoadBalancerClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/CosmosDBAccountClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/CosmosDBAccountClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/CosmosDBAccountClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/CosmosDBAccountClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a VM client
			client, err := CreateLoadBalancerClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}

func TestFrontDoorClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/FrontDoorClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/FrontDoorClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/FrontDoorClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/FrontDoorClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a Front Door client
			client, err := CreateFrontDoorClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}

func TestFrontDoorFrontendEndpointClientBaseURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
	}{
		{"GovCloud/FrontDoorClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint},
		{"PublicCloud/FrontDoorClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint},
		{"ChinaCloud/FrontDoorClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint},
		{"GermanCloud/FrontDoorClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint},
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)

			// Get a AFD frontend endpoint client
			client, err := CreateFrontDoorFrontendEndpointClientE("")
			require.NoError(t, err)

			// Check for correct ARM URI
			assert.Equal(t, tt.ExpectedBaseURI, client.BaseURI)
		})
	}
}

func TestCreateManagedEnvironmentsClientEEndpointURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
		ExpectErr       bool
	}{
		{"Default/ManagedEnvironmentsClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint, false},
		{"PublicCloud/ManagedEnvironmentsClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint, false},
		{"GovCloud/ManagedEnvironmentsClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint, false},
		{"ChinaCloud/ManagedEnvironmentsClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint, false},
		{"GermanCloud/ManagedEnvironmentsClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint, true}, //GermanCloud is deleted as of 2021-10-21 https://learn.microsoft.com/en-us/previous-versions/azure/germany/germany-welcome
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			if tt.EnvironmentName != "" {
				os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)
			} else {
				os.Unsetenv(AzureEnvironmentEnvName)
			}

			// Get a ManagedEnvironmentsClient client
			client, err := CreateManagedEnvironmentsClientE("")
			if tt.ExpectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
				// Not ideal, but to get the base URI we need to access the internal field
				field := reflect.ValueOf(client).Elem().FieldByName("internal").Elem().FieldByName("ep")
				assert.Equal(t, field.String()+"/", tt.ExpectedBaseURI)
			}
		})
	}
}

func TestCreateContainerAppsClientEEndpointURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
		ExpectErr       bool
	}{
		{"Default/ContainerAppsClient", "", autorest.PublicCloud.ResourceManagerEndpoint, false},
		{"PublicCloud/ContainerAppsClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint, false},
		{"GovCloud/ContainerAppsClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint, false},
		{"ChinaCloud/ContainerAppsClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint, false},
		{"GermanCloud/ContainerAppsClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint, true}, //GermanCloud is deleted as of 2021-10-21 https://learn.microsoft.com/en-us/previous-versions/azure/germany/germany-welcome
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			if tt.EnvironmentName != "" {
				os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)
			} else {
				os.Unsetenv(AzureEnvironmentEnvName)
			}

			// Get a ManagedEnvironmentsClient client
			client, err := CreateContainerAppsClientE("")
			if tt.ExpectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
				// Not ideal, but to get the base URI we need to access the internal field
				field := reflect.ValueOf(client).Elem().FieldByName("internal").Elem().FieldByName("ep")
				assert.Equal(t, field.String()+"/", tt.ExpectedBaseURI)
			}
		})
	}
}

func TestCreateContainerAppJobsClientEEndpointURISetCorrectly(t *testing.T) {
	var cases = []struct {
		CaseName        string
		EnvironmentName string
		ExpectedBaseURI string
		ExpectErr       bool
	}{
		{"Default/ContainerAppsClient", "", autorest.PublicCloud.ResourceManagerEndpoint, false},
		{"PublicCloud/ContainerAppsClient", publicCloudEnvName, autorest.PublicCloud.ResourceManagerEndpoint, false},
		{"GovCloud/ContainerAppsClient", govCloudEnvName, autorest.USGovernmentCloud.ResourceManagerEndpoint, false},
		{"ChinaCloud/ContainerAppsClient", chinaCloudEnvName, autorest.ChinaCloud.ResourceManagerEndpoint, false},
		{"GermanCloud/ContainerAppsClient", germanyCloudEnvName, autorest.GermanCloud.ResourceManagerEndpoint, true}, //GermanCloud is deleted as of 2021-10-21 https://learn.microsoft.com/en-us/previous-versions/azure/germany/germany-welcome
	}

	// save any current env value and restore on exit
	currentEnv := os.Getenv(AzureEnvironmentEnvName)
	defer os.Setenv(AzureEnvironmentEnvName, currentEnv)

	for _, tt := range cases {
		// The following is necessary to make sure testCase's values don't
		// get updated due to concurrency within the scope of t.Run(..) below
		tt := tt
		t.Run(tt.CaseName, func(t *testing.T) {
			// Override env setting
			if tt.EnvironmentName != "" {
				os.Setenv(AzureEnvironmentEnvName, tt.EnvironmentName)
			} else {
				os.Unsetenv(AzureEnvironmentEnvName)
			}

			// Get a ManagedEnvironmentsClient client
			client, err := CreateContainerAppJobsClientE("")
			if tt.ExpectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
				// Not ideal, but to get the base URI we need to access the internal field
				field := reflect.ValueOf(client).Elem().FieldByName("internal").Elem().FieldByName("ep")
				assert.Equal(t, field.String()+"/", tt.ExpectedBaseURI)
			}
		})
	}
}
