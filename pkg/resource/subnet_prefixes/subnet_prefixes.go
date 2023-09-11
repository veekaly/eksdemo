package subnet_prefixes

import (
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/resource"
)

func NewResource() *resource.Resource {
	res := &resource.Resource{
		Command: cmd.Command{
			Name:        "subnet-prefixes",
			Description: "VPC Subnet /28 Prefixes",
			Aliases:     []string{"subnet-prefix"},
			Args:        []string{"Subnet ID"},
		},

		Getter: &Getter{},

		Options: &resource.CommonOptions{
			ClusterFlagOptional: true,
		},
	}

	return res
}
