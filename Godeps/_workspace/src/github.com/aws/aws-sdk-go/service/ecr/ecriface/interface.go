// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ecriface provides an interface for the Amazon EC2 Container Registry.
package ecriface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecr"
)

// ECRAPI is the interface type for ecr.ECR.
type ECRAPI interface {
	BatchCheckLayerAvailabilityRequest(*ecr.BatchCheckLayerAvailabilityInput) (*request.Request, *ecr.BatchCheckLayerAvailabilityOutput)

	BatchCheckLayerAvailability(*ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error)

	BatchDeleteImageRequest(*ecr.BatchDeleteImageInput) (*request.Request, *ecr.BatchDeleteImageOutput)

	BatchDeleteImage(*ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error)

	BatchGetImageRequest(*ecr.BatchGetImageInput) (*request.Request, *ecr.BatchGetImageOutput)

	BatchGetImage(*ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error)

	CompleteLayerUploadRequest(*ecr.CompleteLayerUploadInput) (*request.Request, *ecr.CompleteLayerUploadOutput)

	CompleteLayerUpload(*ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error)

	CreateRepositoryRequest(*ecr.CreateRepositoryInput) (*request.Request, *ecr.CreateRepositoryOutput)

	CreateRepository(*ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error)

	DeleteRepositoryRequest(*ecr.DeleteRepositoryInput) (*request.Request, *ecr.DeleteRepositoryOutput)

	DeleteRepository(*ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error)

	DeleteRepositoryPolicyRequest(*ecr.DeleteRepositoryPolicyInput) (*request.Request, *ecr.DeleteRepositoryPolicyOutput)

	DeleteRepositoryPolicy(*ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error)

	DescribeRepositoriesRequest(*ecr.DescribeRepositoriesInput) (*request.Request, *ecr.DescribeRepositoriesOutput)

	DescribeRepositories(*ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error)

	GetAuthorizationTokenRequest(*ecr.GetAuthorizationTokenInput) (*request.Request, *ecr.GetAuthorizationTokenOutput)

	GetAuthorizationToken(*ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error)

	GetDownloadUrlForLayerRequest(*ecr.GetDownloadUrlForLayerInput) (*request.Request, *ecr.GetDownloadUrlForLayerOutput)

	GetDownloadUrlForLayer(*ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error)

	GetRepositoryPolicyRequest(*ecr.GetRepositoryPolicyInput) (*request.Request, *ecr.GetRepositoryPolicyOutput)

	GetRepositoryPolicy(*ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error)

	InitiateLayerUploadRequest(*ecr.InitiateLayerUploadInput) (*request.Request, *ecr.InitiateLayerUploadOutput)

	InitiateLayerUpload(*ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error)

	ListImagesRequest(*ecr.ListImagesInput) (*request.Request, *ecr.ListImagesOutput)

	ListImages(*ecr.ListImagesInput) (*ecr.ListImagesOutput, error)

	PutImageRequest(*ecr.PutImageInput) (*request.Request, *ecr.PutImageOutput)

	PutImage(*ecr.PutImageInput) (*ecr.PutImageOutput, error)

	SetRepositoryPolicyRequest(*ecr.SetRepositoryPolicyInput) (*request.Request, *ecr.SetRepositoryPolicyOutput)

	SetRepositoryPolicy(*ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error)

	UploadLayerPartRequest(*ecr.UploadLayerPartInput) (*request.Request, *ecr.UploadLayerPartOutput)

	UploadLayerPart(*ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error)
}

var _ ECRAPI = (*ecr.ECR)(nil)