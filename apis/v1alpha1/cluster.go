// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterSpec defines the desired state of Cluster.
//
// An object representing an Amazon EKS cluster.
type ClusterSpec struct {

	// The access configuration for the cluster.

	AccessConfig *CreateAccessConfigRequest `json:"accessConfig,omitempty"`
	// If you set this value to False when creating a cluster, the default networking
	// add-ons will not be installed.
	//
	// The default networking addons include vpc-cni, coredns, and kube-proxy.
	//
	// Use this option when you plan to install third-party alternative add-ons
	// or self-manage the default networking add-ons.

	BootstrapSelfManagedAddons *bool `json:"bootstrapSelfManagedAddons,omitempty"`
	// A unique, case-sensitive identifier that you provide to ensurethe idempotency
	// of the request.

	ClientRequestToken *string `json:"clientRequestToken,omitempty"`
	// Enable or disable the compute capability of EKS Auto Mode when creating your
	// EKS Auto Mode cluster. If the compute capability is enabled, EKS Auto Mode
	// will create and delete EC2 Managed Instances in your Amazon Web Services
	// account

	ComputeConfig *ComputeConfigRequest `json:"computeConfig,omitempty"`
	// The encryption configuration for the cluster.

	EncryptionConfig []*EncryptionConfig `json:"encryptionConfig,omitempty"`
	// The Kubernetes network configuration for the cluster.

	KubernetesNetworkConfig *KubernetesNetworkConfigRequest `json:"kubernetesNetworkConfig,omitempty"`
	// Enable or disable exporting the Kubernetes control plane logs for your cluster
	// to CloudWatch Logs. By default, cluster control plane logs aren't exported
	// to CloudWatch Logs. For more information, see Amazon EKS Cluster control
	// plane logs (https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
	// in the Amazon EKS User Guide .
	//
	// CloudWatch Logs ingestion, archive storage, and data scanning rates apply
	// to exported control plane logs. For more information, see CloudWatch Pricing
	// (http://aws.amazon.com/cloudwatch/pricing/).

	Logging *Logging `json:"logging,omitempty"`
	// The unique name to give to your cluster. The name can contain only alphanumeric
	// characters (case-sensitive),hyphens, and underscores. It must start with
	// an alphanumeric character and can't be longer than100 characters. The name
	// must be unique within the Amazon Web Services Region and Amazon Web Services
	// account that you're creating the cluster in.

	// +kubebuilder:validation:Required

	Name *string `json:"name"`
	// An object representing the configuration of your local Amazon EKS cluster
	// on an Amazon Web Services Outpost. Before creating a local cluster on an
	// Outpost, review Local clusters for Amazon EKS on Amazon Web Services Outposts
	// (https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-local-cluster-overview.html)
	// in the Amazon EKS User Guide. This object isn't available for creating Amazon
	// EKS clusters on the Amazon Web Services cloud.

	OutpostConfig *OutpostConfigRequest `json:"outpostConfig,omitempty"`
	// The configuration in the cluster for EKS Hybrid Nodes. You can't change or
	// update this configuration after the cluster is created.

	RemoteNetworkConfig *RemoteNetworkConfigRequest `json:"remoteNetworkConfig,omitempty"`
	// The VPC configuration that's used by the cluster control plane. Amazon EKS
	// VPC resources have specific requirements to work properly with Kubernetes.
	// For more information, see Cluster VPC Considerations (https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html)
	// and Cluster Security Group Considerations (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html)
	// in the Amazon EKS User Guide. You must specify at least two subnets. You
	// can specify up to five security groups. However, we recommend that you use
	// a dedicated security group for your cluster control plane.

	// +kubebuilder:validation:Required

	ResourcesVPCConfig *VPCConfigRequest `json:"resourcesVPCConfig"`
	// The Amazon Resource Name (ARN) of the IAM role that provides permissions
	// for the Kubernetes control plane to make calls to Amazon Web Services API
	// operations on your behalf. For more information, see Amazon EKS Service IAM
	// Role (https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html)
	// in the Amazon EKS User Guide .

	RoleARN *string `json:"roleARN,omitempty"`

	RoleRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"roleRef,omitempty"`
	// Enable or disable the block storage capability of EKS Auto Mode when creating
	// your EKS Auto Mode cluster. If the block storage capability is enabled, EKS
	// Auto Mode will create and delete EBS volumes in your Amazon Web Services
	// account.

	StorageConfig *StorageConfigRequest `json:"storageConfig,omitempty"`
	// Metadata that assists with categorization and organization. Each tag consists
	// of a key and an optional value. You define both. Tags don't propagate to
	// any other cluster or Amazon Web Services resources.

	Tags map[string]*string `json:"tags,omitempty"`
	// New clusters, by default, have extended support enabled. You can disable
	// extended support when creating a cluster by setting this value to STANDARD.

	UpgradePolicy *UpgradePolicyRequest `json:"upgradePolicy,omitempty"`
	// The desired Kubernetes version for your cluster. If you don't specify a value
	// here, the default version available in Amazon EKS is used.
	//
	// The default version might not be the latest version available.

	Version *string `json:"version,omitempty"`
	// Enable or disable ARC zonal shift for the cluster. If zonal shift is enabled,
	// Amazon Web Services configures zonal autoshift for the cluster.
	//
	// Zonal shift is a feature of Amazon Application Recovery Controller (ARC).
	// ARC zonal shift is designed to be a temporary measure that allows you to
	// move traffic for a resource away from an impaired AZ until the zonal shift
	// expires or you cancel it. You can extend the zonal shift if necessary.
	//
	// You can start a zonal shift for an EKS cluster, or you can allow Amazon Web
	// Services to do it for you by enabling zonal autoshift. This shift updates
	// the flow of east-to-west network traffic in your cluster to only consider
	// network endpoints for Pods running on worker nodes in healthy AZs. Additionally,
	// any ALB or NLB handling ingress traffic for applications in your EKS cluster
	// will automatically route traffic to targets in the healthy AZs. For more
	// information about zonal shift in EKS, see Learn about Amazon Application
	// Recovery Controller (ARC) Zonal Shift in Amazon EKS (https://docs.aws.amazon.com/eks/latest/userguide/zone-shift.html)
	// in the Amazon EKS User Guide .

	ZonalShiftConfig *ZonalShiftConfigRequest `json:"zonalShiftConfig,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The certificate-authority-data for your cluster.
	// +kubebuilder:validation:Optional
	CertificateAuthority *Certificate `json:"certificateAuthority,omitempty"`
	// The configuration used to connect to a cluster for registration.
	// +kubebuilder:validation:Optional
	ConnectorConfig *ConnectorConfigResponse `json:"connectorConfig,omitempty"`
	// The Unix epoch timestamp at object creation.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// The endpoint for your Kubernetes API server.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty"`
	// An object representing the health of your Amazon EKS cluster.
	// +kubebuilder:validation:Optional
	Health *ClusterHealth `json:"health,omitempty"`
	// The ID of your local Amazon EKS cluster on an Amazon Web Services Outpost.
	// This property isn't available for an Amazon EKS cluster on the Amazon Web
	// Services cloud.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty"`
	// The identity provider information for the cluster.
	// +kubebuilder:validation:Optional
	Identity *Identity `json:"identity,omitempty"`
	// The platform version of your Amazon EKS cluster. For more information about
	// clusters deployed on the Amazon Web Services Cloud, see Platform versions
	// (https://docs.aws.amazon.com/eks/latest/userguide/platform-versions.html)
	// in the Amazon EKS User Guide . For more information about local clusters
	// deployed on an Outpost, see Amazon EKS local cluster platform versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-outposts-platform-versions.html)
	// in the Amazon EKS User Guide .
	// +kubebuilder:validation:Optional
	PlatformVersion *string `json:"platformVersion,omitempty"`
	// The current status of the cluster.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// Cluster is the Schema for the Clusters API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="VERSION",type=string,priority=0,JSONPath=`.spec.version`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="PLATFORMVERSION",type=string,priority=1,JSONPath=`.status.platformVersion`
// +kubebuilder:printcolumn:name="ENDPOINT",type=string,priority=1,JSONPath=`.status.endpoint`
// +kubebuilder:printcolumn:name="Synced",type="string",priority=0,JSONPath=".status.conditions[?(@.type==\"ACK.ResourceSynced\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",priority=0,JSONPath=".metadata.creationTimestamp"
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// ClusterList contains a list of Cluster
// +kubebuilder:object:root=true
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
