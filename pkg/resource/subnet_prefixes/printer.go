package subnet_prefixes

import (
	"io"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/eksdemo/pkg/printer"
)

type SubnetPrinter struct {
	subnets       []types.Subnet
	multipleCidrs bool
}

func NewPrinter(subnets []types.Subnet) *SubnetPrinter {
	return &SubnetPrinter{subnets: subnets}
}

func (p *SubnetPrinter) PrintTable(writer io.Writer) error {
	table := printer.NewTablePrinter()
	table.SetHeader([]string{"Id", "Zone", "CIDR Reservation", "CIDR", "Prefixes (Total/Assigned)"})

	for _, subnet := range p.subnets {
		table.AppendRow([]string{
			aws.ToString(subnet.SubnetId),
			aws.ToString(subnet.AvailabilityZone),
			aws.ToString(subnet.SubnetId),
			aws.ToString(subnet.CidrBlock),
			strconv.Itoa(int(aws.ToInt32(subnet.AvailableIpAddressCount))),
		})
	}

	table.Print(writer)

	return nil
}

func (p *SubnetPrinter) PrintJSON(writer io.Writer) error {
	return printer.EncodeJSON(writer, p.subnets)
}

func (p *SubnetPrinter) PrintYAML(writer io.Writer) error {
	return printer.EncodeYAML(writer, p.subnets)
}
