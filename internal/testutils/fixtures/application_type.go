package fixtures

import (
	m "github.com/RedHatInsights/sources-api-go/model"
)

var userOwnership = "user"

var TestApplicationTypeData = []m.ApplicationType{
	{
		Id:          1,
		Name:        "app type name",
		DisplayName: "test app type",
	},
	{
		Id:          2,
		Name:        "second app type name",
		DisplayName: "second test app type",
	},
	{
		Id:                   3,
		DisplayName:          "app-studio",
		Name:                 "/insights/platform/app-studio",
		ResourceOwnership:    &userOwnership,
		SupportedSourceTypes: []byte(`["bitbucket", "dockerhub", "github", "gitlab", "quay"]`),
	},
	{
		Id:                   4,
		DisplayName:          "Cost Management",
		Name:                 "/insights/platform/cost-management",
		SupportedSourceTypes: []byte(`["amazon", "azure", "google", "oracle-cloud-infrastructure", "openshift", "ibm"]`),
	},
}
