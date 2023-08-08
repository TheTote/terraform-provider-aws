// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package networkfirewall

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	networkfirewall_sdkv1 "github.com/aws/aws-sdk-go/service/networkfirewall"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceFirewall,
			TypeName: "aws_networkfirewall_firewall",
		},
		{
			Factory:  DataSourceFirewallPolicy,
			TypeName: "aws_networkfirewall_firewall_policy",
		},
		{
			Factory:  DataSourceFirewallResourcePolicy,
			TypeName: "aws_networkfirewall_resource_policy",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceFirewall,
			TypeName: "aws_networkfirewall_firewall",
			Name:     "Firewall",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceFirewallPolicy,
			TypeName: "aws_networkfirewall_firewall_policy",
			Name:     "Firewall Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceLoggingConfiguration,
			TypeName: "aws_networkfirewall_logging_configuration",
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_networkfirewall_resource_policy",
		},
		{
			Factory:  ResourceRuleGroup,
			TypeName: "aws_networkfirewall_rule_group",
			Name:     "Rule Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.NetworkFirewall
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*networkfirewall_sdkv1.NetworkFirewall, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return networkfirewall_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
