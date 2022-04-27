package fixtures

import m "github.com/RedHatInsights/sources-api-go/model"

var TestTenantData = []m.Tenant{
	{
		Id:             1,
		ExternalTenant: "12345",
		OrgID:          "9876543210",
	},
	{
		Id:             2,
		ExternalTenant: "67890",
	},
}