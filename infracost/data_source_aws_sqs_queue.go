package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsSQSQueue() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":        resourcesSchema(),
			"monthly_requests": usageSchema(),
			"request_size_kb":  usageSchema(),
		},
	}
}
