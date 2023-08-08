// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package guardduty

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	guardduty_sdkv1 "github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceFindingIds,
			Name:    "Finding Ids",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDetector,
			TypeName: "aws_guardduty_detector",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDetector,
			TypeName: "aws_guardduty_detector",
			Name:     "Detector",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceFilter,
			TypeName: "aws_guardduty_filter",
			Name:     "Filter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceInviteAccepter,
			TypeName: "aws_guardduty_invite_accepter",
		},
		{
			Factory:  ResourceIPSet,
			TypeName: "aws_guardduty_ipset",
			Name:     "IP Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceMember,
			TypeName: "aws_guardduty_member",
		},
		{
			Factory:  ResourceOrganizationAdminAccount,
			TypeName: "aws_guardduty_organization_admin_account",
		},
		{
			Factory:  ResourceOrganizationConfiguration,
			TypeName: "aws_guardduty_organization_configuration",
		},
		{
			Factory:  ResourcePublishingDestination,
			TypeName: "aws_guardduty_publishing_destination",
		},
		{
			Factory:  ResourceThreatIntelSet,
			TypeName: "aws_guardduty_threatintelset",
			Name:     "Threat Intel Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GuardDuty
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*guardduty_sdkv1.GuardDuty, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return guardduty_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
