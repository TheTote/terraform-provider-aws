// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package devicefarm

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	devicefarm_sdkv1 "github.com/aws/aws-sdk-go/service/devicefarm"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDevicePool,
			TypeName: "aws_devicefarm_device_pool",
			Name:     "Device Pool",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceInstanceProfile,
			TypeName: "aws_devicefarm_instance_profile",
			Name:     "Instance Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceNetworkProfile,
			TypeName: "aws_devicefarm_network_profile",
			Name:     "Network Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceProject,
			TypeName: "aws_devicefarm_project",
			Name:     "Project",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceTestGridProject,
			TypeName: "aws_devicefarm_test_grid_project",
			Name:     "Test Grid Project",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceUpload,
			TypeName: "aws_devicefarm_upload",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DeviceFarm
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*devicefarm_sdkv1.DeviceFarm, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return devicefarm_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
