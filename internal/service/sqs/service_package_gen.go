// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sqs

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	sqs_sdkv1 "github.com/aws/aws-sdk-go/service/sqs"
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
			Factory:  DataSourceQueue,
			TypeName: "aws_sqs_queue",
		},
		{
			Factory:  DataSourceQueues,
			TypeName: "aws_sqs_queues",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceQueue,
			TypeName: "aws_sqs_queue",
			Name:     "Queue",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            types.ListTagsFunc(listTags_),
				UpdateTags:          types.UpdateTagsFunc(updateTags_),
			},
		},
		{
			Factory:  ResourceQueuePolicy,
			TypeName: "aws_sqs_queue_policy",
		},
		{
			Factory:  ResourceQueueRedriveAllowPolicy,
			TypeName: "aws_sqs_queue_redrive_allow_policy",
		},
		{
			Factory:  ResourceQueueRedrivePolicy,
			TypeName: "aws_sqs_queue_redrive_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SQS
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*sqs_sdkv1.SQS, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return sqs_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
