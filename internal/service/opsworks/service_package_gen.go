// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package opsworks

import (
	"context"

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
			Factory:  ResourceApplication,
			TypeName: "aws_opsworks_application",
		},
		{
			Factory:  ResourceCustomLayer,
			TypeName: "aws_opsworks_custom_layer",
		},
		{
			Factory:  ResourceECSClusterLayer,
			TypeName: "aws_opsworks_ecs_cluster_layer",
		},
		{
			Factory:  ResourceGangliaLayer,
			TypeName: "aws_opsworks_ganglia_layer",
		},
		{
			Factory:  ResourceHAProxyLayer,
			TypeName: "aws_opsworks_haproxy_layer",
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_opsworks_instance",
		},
		{
			Factory:  ResourceJavaAppLayer,
			TypeName: "aws_opsworks_java_app_layer",
		},
		{
			Factory:  ResourceMemcachedLayer,
			TypeName: "aws_opsworks_memcached_layer",
		},
		{
			Factory:  ResourceMySQLLayer,
			TypeName: "aws_opsworks_mysql_layer",
		},
		{
			Factory:  ResourceNodejsAppLayer,
			TypeName: "aws_opsworks_nodejs_app_layer",
		},
		{
			Factory:  ResourcePermission,
			TypeName: "aws_opsworks_permission",
		},
		{
			Factory:  ResourcePHPAppLayer,
			TypeName: "aws_opsworks_php_app_layer",
		},
		{
			Factory:  ResourceRailsAppLayer,
			TypeName: "aws_opsworks_rails_app_layer",
		},
		{
			Factory:  ResourceRDSDBInstance,
			TypeName: "aws_opsworks_rds_db_instance",
		},
		{
			Factory:  ResourceStack,
			TypeName: "aws_opsworks_stack",
		},
		{
			Factory:  ResourceStaticWebLayer,
			TypeName: "aws_opsworks_static_web_layer",
		},
		{
			Factory:  ResourceUserProfile,
			TypeName: "aws_opsworks_user_profile",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpsWorks
}

var ServicePackage = &servicePackage{}