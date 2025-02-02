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

type AMITypes string

const (
	AMITypes_AL2023_ARM_64_STANDARD     AMITypes = "AL2023_ARM_64_STANDARD"
	AMITypes_AL2023_x86_64_NEURON       AMITypes = "AL2023_x86_64_NEURON"
	AMITypes_AL2023_x86_64_NVIDIA       AMITypes = "AL2023_x86_64_NVIDIA"
	AMITypes_AL2023_x86_64_STANDARD     AMITypes = "AL2023_x86_64_STANDARD"
	AMITypes_AL2_ARM_64                 AMITypes = "AL2_ARM_64"
	AMITypes_AL2_x86_64                 AMITypes = "AL2_x86_64"
	AMITypes_AL2_x86_64_GPU             AMITypes = "AL2_x86_64_GPU"
	AMITypes_BOTTLEROCKET_ARM_64        AMITypes = "BOTTLEROCKET_ARM_64"
	AMITypes_BOTTLEROCKET_ARM_64_NVIDIA AMITypes = "BOTTLEROCKET_ARM_64_NVIDIA"
	AMITypes_BOTTLEROCKET_x86_64        AMITypes = "BOTTLEROCKET_x86_64"
	AMITypes_BOTTLEROCKET_x86_64_NVIDIA AMITypes = "BOTTLEROCKET_x86_64_NVIDIA"
	AMITypes_CUSTOM                     AMITypes = "CUSTOM"
	AMITypes_WINDOWS_CORE_2019_x86_64   AMITypes = "WINDOWS_CORE_2019_x86_64"
	AMITypes_WINDOWS_CORE_2022_x86_64   AMITypes = "WINDOWS_CORE_2022_x86_64"
	AMITypes_WINDOWS_FULL_2019_x86_64   AMITypes = "WINDOWS_FULL_2019_x86_64"
	AMITypes_WINDOWS_FULL_2022_x86_64   AMITypes = "WINDOWS_FULL_2022_x86_64"
)

type AccessScopeType string

const (
	AccessScopeType_cluster   AccessScopeType = "cluster"
	AccessScopeType_namespace AccessScopeType = "namespace"
)

type AddonIssueCode string

const (
	AddonIssueCode_ACCESS_DENIED                   AddonIssueCode = "ACCESS_DENIED"
	AddonIssueCode_ADDON_PERMISSION_FAILURE        AddonIssueCode = "ADDON_PERMISSION_FAILURE"
	AddonIssueCode_ADDON_SUBSCRIPTION_NEEDED       AddonIssueCode = "ADDON_SUBSCRIPTION_NEEDED"
	AddonIssueCode_ADMISSION_REQUEST_DENIED        AddonIssueCode = "ADMISSION_REQUEST_DENIED"
	AddonIssueCode_CLUSTER_UNREACHABLE             AddonIssueCode = "CLUSTER_UNREACHABLE"
	AddonIssueCode_CONFIGURATION_CONFLICT          AddonIssueCode = "CONFIGURATION_CONFLICT"
	AddonIssueCode_INSUFFICIENT_NUMBER_OF_REPLICAS AddonIssueCode = "INSUFFICIENT_NUMBER_OF_REPLICAS"
	AddonIssueCode_INTERNAL_FAILURE                AddonIssueCode = "INTERNAL_FAILURE"
	AddonIssueCode_K8S_RESOURCE_NOT_FOUND          AddonIssueCode = "K8S_RESOURCE_NOT_FOUND"
	AddonIssueCode_UNSUPPORTED_ADDON_MODIFICATION  AddonIssueCode = "UNSUPPORTED_ADDON_MODIFICATION"
)

type AddonStatus_SDK string

const (
	AddonStatus_SDK_ACTIVE        AddonStatus_SDK = "ACTIVE"
	AddonStatus_SDK_CREATE_FAILED AddonStatus_SDK = "CREATE_FAILED"
	AddonStatus_SDK_CREATING      AddonStatus_SDK = "CREATING"
	AddonStatus_SDK_DEGRADED      AddonStatus_SDK = "DEGRADED"
	AddonStatus_SDK_DELETE_FAILED AddonStatus_SDK = "DELETE_FAILED"
	AddonStatus_SDK_DELETING      AddonStatus_SDK = "DELETING"
	AddonStatus_SDK_UPDATE_FAILED AddonStatus_SDK = "UPDATE_FAILED"
	AddonStatus_SDK_UPDATING      AddonStatus_SDK = "UPDATING"
)

type AuthenticationMode string

const (
	AuthenticationMode_API                AuthenticationMode = "API"
	AuthenticationMode_API_AND_CONFIG_MAP AuthenticationMode = "API_AND_CONFIG_MAP"
	AuthenticationMode_CONFIG_MAP         AuthenticationMode = "CONFIG_MAP"
)

type CapacityTypes string

const (
	CapacityTypes_CAPACITY_BLOCK CapacityTypes = "CAPACITY_BLOCK"
	CapacityTypes_ON_DEMAND      CapacityTypes = "ON_DEMAND"
	CapacityTypes_SPOT           CapacityTypes = "SPOT"
)

type Category string

const (
	Category_UPGRADE_READINESS Category = "UPGRADE_READINESS"
)

type ClusterIssueCode string

const (
	ClusterIssueCode_ACCESS_DENIED                  ClusterIssueCode = "ACCESS_DENIED"
	ClusterIssueCode_CLUSTER_UNREACHABLE            ClusterIssueCode = "CLUSTER_UNREACHABLE"
	ClusterIssueCode_CONFIGURATION_CONFLICT         ClusterIssueCode = "CONFIGURATION_CONFLICT"
	ClusterIssueCode_EC2_SECURITY_GROUP_NOT_FOUND   ClusterIssueCode = "EC2_SECURITY_GROUP_NOT_FOUND"
	ClusterIssueCode_EC2_SERVICE_NOT_SUBSCRIBED     ClusterIssueCode = "EC2_SERVICE_NOT_SUBSCRIBED"
	ClusterIssueCode_EC2_SUBNET_NOT_FOUND           ClusterIssueCode = "EC2_SUBNET_NOT_FOUND"
	ClusterIssueCode_IAM_ROLE_NOT_FOUND             ClusterIssueCode = "IAM_ROLE_NOT_FOUND"
	ClusterIssueCode_INSUFFICIENT_FREE_ADDRESSES    ClusterIssueCode = "INSUFFICIENT_FREE_ADDRESSES"
	ClusterIssueCode_INTERNAL_FAILURE               ClusterIssueCode = "INTERNAL_FAILURE"
	ClusterIssueCode_KMS_GRANT_REVOKED              ClusterIssueCode = "KMS_GRANT_REVOKED"
	ClusterIssueCode_KMS_KEY_DISABLED               ClusterIssueCode = "KMS_KEY_DISABLED"
	ClusterIssueCode_KMS_KEY_MARKED_FOR_DELETION    ClusterIssueCode = "KMS_KEY_MARKED_FOR_DELETION"
	ClusterIssueCode_KMS_KEY_NOT_FOUND              ClusterIssueCode = "KMS_KEY_NOT_FOUND"
	ClusterIssueCode_OTHER                          ClusterIssueCode = "OTHER"
	ClusterIssueCode_RESOURCE_LIMIT_EXCEEDED        ClusterIssueCode = "RESOURCE_LIMIT_EXCEEDED"
	ClusterIssueCode_RESOURCE_NOT_FOUND             ClusterIssueCode = "RESOURCE_NOT_FOUND"
	ClusterIssueCode_STS_REGIONAL_ENDPOINT_DISABLED ClusterIssueCode = "STS_REGIONAL_ENDPOINT_DISABLED"
	ClusterIssueCode_UNSUPPORTED_VERSION            ClusterIssueCode = "UNSUPPORTED_VERSION"
	ClusterIssueCode_VPC_NOT_FOUND                  ClusterIssueCode = "VPC_NOT_FOUND"
)

type ClusterStatus_SDK string

const (
	ClusterStatus_SDK_ACTIVE   ClusterStatus_SDK = "ACTIVE"
	ClusterStatus_SDK_CREATING ClusterStatus_SDK = "CREATING"
	ClusterStatus_SDK_DELETING ClusterStatus_SDK = "DELETING"
	ClusterStatus_SDK_FAILED   ClusterStatus_SDK = "FAILED"
	ClusterStatus_SDK_PENDING  ClusterStatus_SDK = "PENDING"
	ClusterStatus_SDK_UPDATING ClusterStatus_SDK = "UPDATING"
)

type ConfigStatus string

const (
	ConfigStatus_ACTIVE   ConfigStatus = "ACTIVE"
	ConfigStatus_CREATING ConfigStatus = "CREATING"
	ConfigStatus_DELETING ConfigStatus = "DELETING"
)

type ConnectorConfigProvider string

const (
	ConnectorConfigProvider_AKS          ConnectorConfigProvider = "AKS"
	ConnectorConfigProvider_ANTHOS       ConnectorConfigProvider = "ANTHOS"
	ConnectorConfigProvider_EC2          ConnectorConfigProvider = "EC2"
	ConnectorConfigProvider_EKS_ANYWHERE ConnectorConfigProvider = "EKS_ANYWHERE"
	ConnectorConfigProvider_GKE          ConnectorConfigProvider = "GKE"
	ConnectorConfigProvider_OPENSHIFT    ConnectorConfigProvider = "OPENSHIFT"
	ConnectorConfigProvider_OTHER        ConnectorConfigProvider = "OTHER"
	ConnectorConfigProvider_RANCHER      ConnectorConfigProvider = "RANCHER"
	ConnectorConfigProvider_TANZU        ConnectorConfigProvider = "TANZU"
)

type EKSAnywhereSubscriptionLicenseType string

const (
	EKSAnywhereSubscriptionLicenseType_Cluster EKSAnywhereSubscriptionLicenseType = "Cluster"
)

type EKSAnywhereSubscriptionStatus string

const (
	EKSAnywhereSubscriptionStatus_ACTIVE   EKSAnywhereSubscriptionStatus = "ACTIVE"
	EKSAnywhereSubscriptionStatus_CREATING EKSAnywhereSubscriptionStatus = "CREATING"
	EKSAnywhereSubscriptionStatus_DELETING EKSAnywhereSubscriptionStatus = "DELETING"
	EKSAnywhereSubscriptionStatus_EXPIRED  EKSAnywhereSubscriptionStatus = "EXPIRED"
	EKSAnywhereSubscriptionStatus_EXPIRING EKSAnywhereSubscriptionStatus = "EXPIRING"
	EKSAnywhereSubscriptionStatus_UPDATING EKSAnywhereSubscriptionStatus = "UPDATING"
)

type EKSAnywhereSubscriptionTermUnit string

const (
	EKSAnywhereSubscriptionTermUnit_MONTHS EKSAnywhereSubscriptionTermUnit = "MONTHS"
)

type ErrorCode string

const (
	ErrorCode_ACCESS_DENIED                   ErrorCode = "ACCESS_DENIED"
	ErrorCode_ADMISSION_REQUEST_DENIED        ErrorCode = "ADMISSION_REQUEST_DENIED"
	ErrorCode_CLUSTER_UNREACHABLE             ErrorCode = "CLUSTER_UNREACHABLE"
	ErrorCode_CONFIGURATION_CONFLICT          ErrorCode = "CONFIGURATION_CONFLICT"
	ErrorCode_ENI_LIMIT_REACHED               ErrorCode = "ENI_LIMIT_REACHED"
	ErrorCode_INSUFFICIENT_FREE_ADDRESSES     ErrorCode = "INSUFFICIENT_FREE_ADDRESSES"
	ErrorCode_INSUFFICIENT_NUMBER_OF_REPLICAS ErrorCode = "INSUFFICIENT_NUMBER_OF_REPLICAS"
	ErrorCode_IP_NOT_AVAILABLE                ErrorCode = "IP_NOT_AVAILABLE"
	ErrorCode_K8S_RESOURCE_NOT_FOUND          ErrorCode = "K8S_RESOURCE_NOT_FOUND"
	ErrorCode_NODE_CREATION_FAILURE           ErrorCode = "NODE_CREATION_FAILURE"
	ErrorCode_OPERATION_NOT_PERMITTED         ErrorCode = "OPERATION_NOT_PERMITTED"
	ErrorCode_POD_EVICTION_FAILURE            ErrorCode = "POD_EVICTION_FAILURE"
	ErrorCode_SECURITY_GROUP_NOT_FOUND        ErrorCode = "SECURITY_GROUP_NOT_FOUND"
	ErrorCode_SUBNET_NOT_FOUND                ErrorCode = "SUBNET_NOT_FOUND"
	ErrorCode_UNKNOWN                         ErrorCode = "UNKNOWN"
	ErrorCode_UNSUPPORTED_ADDON_MODIFICATION  ErrorCode = "UNSUPPORTED_ADDON_MODIFICATION"
	ErrorCode_VPC_ID_NOT_FOUND                ErrorCode = "VPC_ID_NOT_FOUND"
)

type FargateProfileIssueCode string

const (
	FargateProfileIssueCode_ACCESS_DENIED                     FargateProfileIssueCode = "ACCESS_DENIED"
	FargateProfileIssueCode_CLUSTER_UNREACHABLE               FargateProfileIssueCode = "CLUSTER_UNREACHABLE"
	FargateProfileIssueCode_INTERNAL_FAILURE                  FargateProfileIssueCode = "INTERNAL_FAILURE"
	FargateProfileIssueCode_POD_EXECUTION_ROLE_ALREADY_IN_USE FargateProfileIssueCode = "POD_EXECUTION_ROLE_ALREADY_IN_USE"
)

type FargateProfileStatus_SDK string

const (
	FargateProfileStatus_SDK_ACTIVE        FargateProfileStatus_SDK = "ACTIVE"
	FargateProfileStatus_SDK_CREATE_FAILED FargateProfileStatus_SDK = "CREATE_FAILED"
	FargateProfileStatus_SDK_CREATING      FargateProfileStatus_SDK = "CREATING"
	FargateProfileStatus_SDK_DELETE_FAILED FargateProfileStatus_SDK = "DELETE_FAILED"
	FargateProfileStatus_SDK_DELETING      FargateProfileStatus_SDK = "DELETING"
)

type IPFamily string

const (
	IPFamily_IPV4 IPFamily = "IPV4"
	IPFamily_IPV6 IPFamily = "IPV6"
)

type InsightStatusValue string

const (
	InsightStatusValue_ERROR   InsightStatusValue = "ERROR"
	InsightStatusValue_PASSING InsightStatusValue = "PASSING"
	InsightStatusValue_UNKNOWN InsightStatusValue = "UNKNOWN"
	InsightStatusValue_WARNING InsightStatusValue = "WARNING"
)

type LogType string

const (
	LogType_API                LogType = "API"
	LogType_AUDIT              LogType = "AUDIT"
	LogType_AUTHENTICATOR      LogType = "AUTHENTICATOR"
	LogType_CONTROLLER_MANAGER LogType = "CONTROLLER_MANAGER"
	LogType_SCHEDULER          LogType = "SCHEDULER"
)

type NodegroupIssueCode string

const (
	NodegroupIssueCode_ACCESS_DENIED                                  NodegroupIssueCode = "ACCESS_DENIED"
	NodegroupIssueCode_AMI_ID_NOT_FOUND                               NodegroupIssueCode = "AMI_ID_NOT_FOUND"
	NodegroupIssueCode_ASG_INSTANCE_LAUNCH_FAILURES                   NodegroupIssueCode = "ASG_INSTANCE_LAUNCH_FAILURES"
	NodegroupIssueCode_AUTO_SCALING_GROUP_INSTANCE_REFRESH_ACTIVE     NodegroupIssueCode = "AUTO_SCALING_GROUP_INSTANCE_REFRESH_ACTIVE"
	NodegroupIssueCode_AUTO_SCALING_GROUP_INVALID_CONFIGURATION       NodegroupIssueCode = "AUTO_SCALING_GROUP_INVALID_CONFIGURATION"
	NodegroupIssueCode_AUTO_SCALING_GROUP_NOT_FOUND                   NodegroupIssueCode = "AUTO_SCALING_GROUP_NOT_FOUND"
	NodegroupIssueCode_AUTO_SCALING_GROUP_OPT_IN_REQUIRED             NodegroupIssueCode = "AUTO_SCALING_GROUP_OPT_IN_REQUIRED"
	NodegroupIssueCode_AUTO_SCALING_GROUP_RATE_LIMIT_EXCEEDED         NodegroupIssueCode = "AUTO_SCALING_GROUP_RATE_LIMIT_EXCEEDED"
	NodegroupIssueCode_CLUSTER_UNREACHABLE                            NodegroupIssueCode = "CLUSTER_UNREACHABLE"
	NodegroupIssueCode_EC2_INSTANCE_TYPE_DOES_NOT_EXIST               NodegroupIssueCode = "EC2_INSTANCE_TYPE_DOES_NOT_EXIST"
	NodegroupIssueCode_EC2_LAUNCH_TEMPLATE_DELETION_FAILURE           NodegroupIssueCode = "EC2_LAUNCH_TEMPLATE_DELETION_FAILURE"
	NodegroupIssueCode_EC2_LAUNCH_TEMPLATE_INVALID_CONFIGURATION      NodegroupIssueCode = "EC2_LAUNCH_TEMPLATE_INVALID_CONFIGURATION"
	NodegroupIssueCode_EC2_LAUNCH_TEMPLATE_MAX_LIMIT_EXCEEDED         NodegroupIssueCode = "EC2_LAUNCH_TEMPLATE_MAX_LIMIT_EXCEEDED"
	NodegroupIssueCode_EC2_LAUNCH_TEMPLATE_NOT_FOUND                  NodegroupIssueCode = "EC2_LAUNCH_TEMPLATE_NOT_FOUND"
	NodegroupIssueCode_EC2_LAUNCH_TEMPLATE_VERSION_MAX_LIMIT_EXCEEDED NodegroupIssueCode = "EC2_LAUNCH_TEMPLATE_VERSION_MAX_LIMIT_EXCEEDED"
	NodegroupIssueCode_EC2_LAUNCH_TEMPLATE_VERSION_MISMATCH           NodegroupIssueCode = "EC2_LAUNCH_TEMPLATE_VERSION_MISMATCH"
	NodegroupIssueCode_EC2_SECURITY_GROUP_DELETION_FAILURE            NodegroupIssueCode = "EC2_SECURITY_GROUP_DELETION_FAILURE"
	NodegroupIssueCode_EC2_SECURITY_GROUP_NOT_FOUND                   NodegroupIssueCode = "EC2_SECURITY_GROUP_NOT_FOUND"
	NodegroupIssueCode_EC2_SUBNET_INVALID_CONFIGURATION               NodegroupIssueCode = "EC2_SUBNET_INVALID_CONFIGURATION"
	NodegroupIssueCode_EC2_SUBNET_LIST_TOO_LONG                       NodegroupIssueCode = "EC2_SUBNET_LIST_TOO_LONG"
	NodegroupIssueCode_EC2_SUBNET_MISSING_IPV6_ASSIGNMENT             NodegroupIssueCode = "EC2_SUBNET_MISSING_IPV6_ASSIGNMENT"
	NodegroupIssueCode_EC2_SUBNET_NOT_FOUND                           NodegroupIssueCode = "EC2_SUBNET_NOT_FOUND"
	NodegroupIssueCode_IAM_INSTANCE_PROFILE_NOT_FOUND                 NodegroupIssueCode = "IAM_INSTANCE_PROFILE_NOT_FOUND"
	NodegroupIssueCode_IAM_LIMIT_EXCEEDED                             NodegroupIssueCode = "IAM_LIMIT_EXCEEDED"
	NodegroupIssueCode_IAM_NODE_ROLE_NOT_FOUND                        NodegroupIssueCode = "IAM_NODE_ROLE_NOT_FOUND"
	NodegroupIssueCode_IAM_THROTTLING                                 NodegroupIssueCode = "IAM_THROTTLING"
	NodegroupIssueCode_INSTANCE_LIMIT_EXCEEDED                        NodegroupIssueCode = "INSTANCE_LIMIT_EXCEEDED"
	NodegroupIssueCode_INSUFFICIENT_FREE_ADDRESSES                    NodegroupIssueCode = "INSUFFICIENT_FREE_ADDRESSES"
	NodegroupIssueCode_INTERNAL_FAILURE                               NodegroupIssueCode = "INTERNAL_FAILURE"
	NodegroupIssueCode_KUBERNETES_LABEL_INVALID                       NodegroupIssueCode = "KUBERNETES_LABEL_INVALID"
	NodegroupIssueCode_LIMIT_EXCEEDED                                 NodegroupIssueCode = "LIMIT_EXCEEDED"
	NodegroupIssueCode_NODE_CREATION_FAILURE                          NodegroupIssueCode = "NODE_CREATION_FAILURE"
	NodegroupIssueCode_NODE_TERMINATION_FAILURE                       NodegroupIssueCode = "NODE_TERMINATION_FAILURE"
	NodegroupIssueCode_POD_EVICTION_FAILURE                           NodegroupIssueCode = "POD_EVICTION_FAILURE"
	NodegroupIssueCode_SOURCE_EC2_LAUNCH_TEMPLATE_NOT_FOUND           NodegroupIssueCode = "SOURCE_EC2_LAUNCH_TEMPLATE_NOT_FOUND"
	NodegroupIssueCode_UNKNOWN                                        NodegroupIssueCode = "UNKNOWN"
)

type NodegroupStatus_SDK string

const (
	NodegroupStatus_SDK_ACTIVE        NodegroupStatus_SDK = "ACTIVE"
	NodegroupStatus_SDK_CREATE_FAILED NodegroupStatus_SDK = "CREATE_FAILED"
	NodegroupStatus_SDK_CREATING      NodegroupStatus_SDK = "CREATING"
	NodegroupStatus_SDK_DEGRADED      NodegroupStatus_SDK = "DEGRADED"
	NodegroupStatus_SDK_DELETE_FAILED NodegroupStatus_SDK = "DELETE_FAILED"
	NodegroupStatus_SDK_DELETING      NodegroupStatus_SDK = "DELETING"
	NodegroupStatus_SDK_UPDATING      NodegroupStatus_SDK = "UPDATING"
)

type ResolveConflicts string

const (
	ResolveConflicts_NONE      ResolveConflicts = "NONE"
	ResolveConflicts_OVERWRITE ResolveConflicts = "OVERWRITE"
	ResolveConflicts_PRESERVE  ResolveConflicts = "PRESERVE"
)

type SupportType string

const (
	SupportType_EXTENDED SupportType = "EXTENDED"
	SupportType_STANDARD SupportType = "STANDARD"
)

type TaintEffect string

const (
	TaintEffect_NO_EXECUTE         TaintEffect = "NO_EXECUTE"
	TaintEffect_NO_SCHEDULE        TaintEffect = "NO_SCHEDULE"
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
)

type UpdateParamType string

const (
	UpdateParamType_ADDON_VERSION              UpdateParamType = "ADDON_VERSION"
	UpdateParamType_AUTHENTICATION_MODE        UpdateParamType = "AUTHENTICATION_MODE"
	UpdateParamType_CLUSTER_LOGGING            UpdateParamType = "CLUSTER_LOGGING"
	UpdateParamType_COMPUTE_CONFIG             UpdateParamType = "COMPUTE_CONFIG"
	UpdateParamType_CONFIGURATION_VALUES       UpdateParamType = "CONFIGURATION_VALUES"
	UpdateParamType_DESIRED_SIZE               UpdateParamType = "DESIRED_SIZE"
	UpdateParamType_ENCRYPTION_CONFIG          UpdateParamType = "ENCRYPTION_CONFIG"
	UpdateParamType_ENDPOINT_PRIVATE_ACCESS    UpdateParamType = "ENDPOINT_PRIVATE_ACCESS"
	UpdateParamType_ENDPOINT_PUBLIC_ACCESS     UpdateParamType = "ENDPOINT_PUBLIC_ACCESS"
	UpdateParamType_IDENTITY_PROVIDER_CONFIG   UpdateParamType = "IDENTITY_PROVIDER_CONFIG"
	UpdateParamType_KUBERNETES_NETWORK_CONFIG  UpdateParamType = "KUBERNETES_NETWORK_CONFIG"
	UpdateParamType_LABELS_TO_ADD              UpdateParamType = "LABELS_TO_ADD"
	UpdateParamType_LABELS_TO_REMOVE           UpdateParamType = "LABELS_TO_REMOVE"
	UpdateParamType_LAUNCH_TEMPLATE_NAME       UpdateParamType = "LAUNCH_TEMPLATE_NAME"
	UpdateParamType_LAUNCH_TEMPLATE_VERSION    UpdateParamType = "LAUNCH_TEMPLATE_VERSION"
	UpdateParamType_MAX_SIZE                   UpdateParamType = "MAX_SIZE"
	UpdateParamType_MAX_UNAVAILABLE            UpdateParamType = "MAX_UNAVAILABLE"
	UpdateParamType_MAX_UNAVAILABLE_PERCENTAGE UpdateParamType = "MAX_UNAVAILABLE_PERCENTAGE"
	UpdateParamType_MIN_SIZE                   UpdateParamType = "MIN_SIZE"
	UpdateParamType_PLATFORM_VERSION           UpdateParamType = "PLATFORM_VERSION"
	UpdateParamType_POD_IDENTITY_ASSOCIATIONS  UpdateParamType = "POD_IDENTITY_ASSOCIATIONS"
	UpdateParamType_PUBLIC_ACCESS_CIDRS        UpdateParamType = "PUBLIC_ACCESS_CIDRS"
	UpdateParamType_RELEASE_VERSION            UpdateParamType = "RELEASE_VERSION"
	UpdateParamType_RESOLVE_CONFLICTS          UpdateParamType = "RESOLVE_CONFLICTS"
	UpdateParamType_SECURITY_GROUPS            UpdateParamType = "SECURITY_GROUPS"
	UpdateParamType_SERVICE_ACCOUNT_ROLE_ARN   UpdateParamType = "SERVICE_ACCOUNT_ROLE_ARN"
	UpdateParamType_STORAGE_CONFIG             UpdateParamType = "STORAGE_CONFIG"
	UpdateParamType_SUBNETS                    UpdateParamType = "SUBNETS"
	UpdateParamType_TAINTS_TO_ADD              UpdateParamType = "TAINTS_TO_ADD"
	UpdateParamType_TAINTS_TO_REMOVE           UpdateParamType = "TAINTS_TO_REMOVE"
	UpdateParamType_UPGRADE_POLICY             UpdateParamType = "UPGRADE_POLICY"
	UpdateParamType_VERSION                    UpdateParamType = "VERSION"
	UpdateParamType_ZONAL_SHIFT_CONFIG         UpdateParamType = "ZONAL_SHIFT_CONFIG"
)

type UpdateStatus string

const (
	UpdateStatus_CANCELLED   UpdateStatus = "CANCELLED"
	UpdateStatus_FAILED      UpdateStatus = "FAILED"
	UpdateStatus_IN_PROGRESS UpdateStatus = "IN_PROGRESS"
	UpdateStatus_SUCCESSFUL  UpdateStatus = "SUCCESSFUL"
)

type UpdateType string

const (
	UpdateType_ACCESS_CONFIG_UPDATE                  UpdateType = "ACCESS_CONFIG_UPDATE"
	UpdateType_ADDON_UPDATE                          UpdateType = "ADDON_UPDATE"
	UpdateType_ASSOCIATE_ENCRYPTION_CONFIG           UpdateType = "ASSOCIATE_ENCRYPTION_CONFIG"
	UpdateType_ASSOCIATE_IDENTITY_PROVIDER_CONFIG    UpdateType = "ASSOCIATE_IDENTITY_PROVIDER_CONFIG"
	UpdateType_AUTO_MODE_UPDATE                      UpdateType = "AUTO_MODE_UPDATE"
	UpdateType_CONFIG_UPDATE                         UpdateType = "CONFIG_UPDATE"
	UpdateType_DISASSOCIATE_IDENTITY_PROVIDER_CONFIG UpdateType = "DISASSOCIATE_IDENTITY_PROVIDER_CONFIG"
	UpdateType_ENDPOINT_ACCESS_UPDATE                UpdateType = "ENDPOINT_ACCESS_UPDATE"
	UpdateType_LOGGING_UPDATE                        UpdateType = "LOGGING_UPDATE"
	UpdateType_UPGRADE_POLICY_UPDATE                 UpdateType = "UPGRADE_POLICY_UPDATE"
	UpdateType_VERSION_UPDATE                        UpdateType = "VERSION_UPDATE"
	UpdateType_VPC_CONFIG_UPDATE                     UpdateType = "VPC_CONFIG_UPDATE"
	UpdateType_ZONAL_SHIFT_CONFIG_UPDATE             UpdateType = "ZONAL_SHIFT_CONFIG_UPDATE"
)
