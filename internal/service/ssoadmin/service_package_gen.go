// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ssoadmin

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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceInstances,
			TypeName: "aws_ssoadmin_instances",
		},
		{
			Factory:  DataSourcePermissionSet,
			TypeName: "aws_ssoadmin_permission_set",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccountAssignment,
			TypeName: "aws_ssoadmin_account_assignment",
		},
		{
			Factory:  ResourceCustomerManagedPolicyAttachment,
			TypeName: "aws_ssoadmin_customer_managed_policy_attachment",
		},
		{
			Factory:  ResourceAccessControlAttributes,
			TypeName: "aws_ssoadmin_instance_access_control_attributes",
		},
		{
			Factory:  ResourceManagedPolicyAttachment,
			TypeName: "aws_ssoadmin_managed_policy_attachment",
		},
		{
			Factory:  ResourcePermissionSet,
			TypeName: "aws_ssoadmin_permission_set",
		},
		{
			Factory:  ResourcePermissionSetInlinePolicy,
			TypeName: "aws_ssoadmin_permission_set_inline_policy",
		},
		{
			Factory:  ResourcePermissionsBoundaryAttachment,
			TypeName: "aws_ssoadmin_permissions_boundary_attachment",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSOAdmin
}

var ServicePackage = &servicePackage{}