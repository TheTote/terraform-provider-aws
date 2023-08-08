// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	ecs_sdkv1 "github.com/aws/aws-sdk-go/service/ecs"
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
			Factory:  DataSourceCluster,
			TypeName: "aws_ecs_cluster",
		},
		{
			Factory:  DataSourceContainerDefinition,
			TypeName: "aws_ecs_container_definition",
		},
		{
			Factory:  DataSourceService,
			TypeName: "aws_ecs_service",
		},
		{
			Factory:  DataSourceTaskDefinition,
			TypeName: "aws_ecs_task_definition",
		},
		{
			Factory:  DataSourceTaskExecution,
			TypeName: "aws_ecs_task_execution",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccountSettingDefault,
			TypeName: "aws_ecs_account_setting_default",
		},
		{
			Factory:  ResourceCapacityProvider,
			TypeName: "aws_ecs_capacity_provider",
			Name:     "Capacity Provider",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceCluster,
			TypeName: "aws_ecs_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceClusterCapacityProviders,
			TypeName: "aws_ecs_cluster_capacity_providers",
		},
		{
			Factory:  ResourceService,
			TypeName: "aws_ecs_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceTag,
			TypeName: "aws_ecs_tag",
		},
		{
			Factory:  ResourceTaskDefinition,
			TypeName: "aws_ecs_task_definition",
			Name:     "Task Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceTaskSet,
			TypeName: "aws_ecs_task_set",
			Name:     "Task Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ECS
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*ecs_sdkv1.ECS, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return ecs_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
