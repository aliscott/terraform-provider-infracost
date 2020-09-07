package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsDynamoDBTable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":                           resourcesSchema(),
			"monthly_million_write_request_units": usageSchema(),
			"monthly_million_read_request_units":  usageSchema(),
			"monthly_gb_data_storage":             usageSchema(),
		},
	}
}
