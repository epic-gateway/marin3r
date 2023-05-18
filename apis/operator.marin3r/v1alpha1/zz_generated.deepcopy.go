//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/3scale-ops/marin3r/pkg/envoy/container/defaults"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/autoscaling/v2"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASignedConfig) DeepCopyInto(out *CASignedConfig) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASignedConfig.
func (in *CASignedConfig) DeepCopy() *CASignedConfig {
	if in == nil {
		return nil
	}
	out := new(CASignedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateOptions) DeepCopyInto(out *CertificateOptions) {
	*out = *in
	out.Duration = in.Duration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateOptions.
func (in *CertificateOptions) DeepCopy() *CertificateOptions {
	if in == nil {
		return nil
	}
	out := new(CertificateOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRenewalConfig) DeepCopyInto(out *CertificateRenewalConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRenewalConfig.
func (in *CertificateRenewalConfig) DeepCopy() *CertificateRenewalConfig {
	if in == nil {
		return nil
	}
	out := new(CertificateRenewalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerPort) DeepCopyInto(out *ContainerPort) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(v1.Protocol)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerPort.
func (in *ContainerPort) DeepCopy() *ContainerPort {
	if in == nil {
		return nil
	}
	out := new(ContainerPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryService) DeepCopyInto(out *DiscoveryService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryService.
func (in *DiscoveryService) DeepCopy() *DiscoveryService {
	if in == nil {
		return nil
	}
	out := new(DiscoveryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceCertificate) DeepCopyInto(out *DiscoveryServiceCertificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceCertificate.
func (in *DiscoveryServiceCertificate) DeepCopy() *DiscoveryServiceCertificate {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceCertificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryServiceCertificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceCertificateList) DeepCopyInto(out *DiscoveryServiceCertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveryServiceCertificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceCertificateList.
func (in *DiscoveryServiceCertificateList) DeepCopy() *DiscoveryServiceCertificateList {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceCertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryServiceCertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceCertificateSigner) DeepCopyInto(out *DiscoveryServiceCertificateSigner) {
	*out = *in
	if in.SelfSigned != nil {
		in, out := &in.SelfSigned, &out.SelfSigned
		*out = new(SelfSignedConfig)
		**out = **in
	}
	if in.CASigned != nil {
		in, out := &in.CASigned, &out.CASigned
		*out = new(CASignedConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceCertificateSigner.
func (in *DiscoveryServiceCertificateSigner) DeepCopy() *DiscoveryServiceCertificateSigner {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceCertificateSigner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceCertificateSpec) DeepCopyInto(out *DiscoveryServiceCertificateSpec) {
	*out = *in
	if in.IsServerCertificate != nil {
		in, out := &in.IsServerCertificate, &out.IsServerCertificate
		*out = new(bool)
		**out = **in
	}
	if in.IsCA != nil {
		in, out := &in.IsCA, &out.IsCA
		*out = new(bool)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Signer.DeepCopyInto(&out.Signer)
	out.SecretRef = in.SecretRef
	if in.CertificateRenewalConfig != nil {
		in, out := &in.CertificateRenewalConfig, &out.CertificateRenewalConfig
		*out = new(CertificateRenewalConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceCertificateSpec.
func (in *DiscoveryServiceCertificateSpec) DeepCopy() *DiscoveryServiceCertificateSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceCertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceCertificateStatus) DeepCopyInto(out *DiscoveryServiceCertificateStatus) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	if in.NotBefore != nil {
		in, out := &in.NotBefore, &out.NotBefore
		*out = (*in).DeepCopy()
	}
	if in.NotAfter != nil {
		in, out := &in.NotAfter, &out.NotAfter
		*out = (*in).DeepCopy()
	}
	if in.CertificateHash != nil {
		in, out := &in.CertificateHash, &out.CertificateHash
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceCertificateStatus.
func (in *DiscoveryServiceCertificateStatus) DeepCopy() *DiscoveryServiceCertificateStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceCertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceList) DeepCopyInto(out *DiscoveryServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveryService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceList.
func (in *DiscoveryServiceList) DeepCopy() *DiscoveryServiceList {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveryServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceSpec) DeepCopyInto(out *DiscoveryServiceSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(bool)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.PKIConfig != nil {
		in, out := &in.PKIConfig, &out.PKIConfig
		*out = new(PKIConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.XdsServerPort != nil {
		in, out := &in.XdsServerPort, &out.XdsServerPort
		*out = new(uint32)
		**out = **in
	}
	if in.MetricsPort != nil {
		in, out := &in.MetricsPort, &out.MetricsPort
		*out = new(uint32)
		**out = **in
	}
	if in.ServiceConfig != nil {
		in, out := &in.ServiceConfig, &out.ServiceConfig
		*out = new(ServiceConfig)
		**out = **in
	}
	if in.PodPriorityClass != nil {
		in, out := &in.PodPriorityClass, &out.PodPriorityClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceSpec.
func (in *DiscoveryServiceSpec) DeepCopy() *DiscoveryServiceSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryServiceStatus) DeepCopyInto(out *DiscoveryServiceStatus) {
	*out = *in
	if in.DeploymentName != nil {
		in, out := &in.DeploymentName, &out.DeploymentName
		*out = new(string)
		**out = **in
	}
	if in.DeploymentStatus != nil {
		in, out := &in.DeploymentStatus, &out.DeploymentStatus
		*out = new(appsv1.DeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
	out.UnimplementedStatefulSetStatus = in.UnimplementedStatefulSetStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryServiceStatus.
func (in *DiscoveryServiceStatus) DeepCopy() *DiscoveryServiceStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveryServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicReplicasSpec) DeepCopyInto(out *DynamicReplicasSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(v2.HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicReplicasSpec.
func (in *DynamicReplicasSpec) DeepCopy() *DynamicReplicasSpec {
	if in == nil {
		return nil
	}
	out := new(DynamicReplicasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyDeployment) DeepCopyInto(out *EnvoyDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyDeployment.
func (in *EnvoyDeployment) DeepCopy() *EnvoyDeployment {
	if in == nil {
		return nil
	}
	out := new(EnvoyDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyDeploymentList) DeepCopyInto(out *EnvoyDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvoyDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyDeploymentList.
func (in *EnvoyDeploymentList) DeepCopy() *EnvoyDeploymentList {
	if in == nil {
		return nil
	}
	out := new(EnvoyDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyDeploymentSpec) DeepCopyInto(out *EnvoyDeploymentSpec) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ContainerPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientCertificateDuration != nil {
		in, out := &in.ClientCertificateDuration, &out.ClientCertificateDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdminPort != nil {
		in, out := &in.AdminPort, &out.AdminPort
		*out = new(uint32)
		**out = **in
	}
	if in.AdminAccessLogPath != nil {
		in, out := &in.AdminAccessLogPath, &out.AdminAccessLogPath
		*out = new(string)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(ReplicasSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(ProbeSpec)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ProbeSpec)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ShutdownManager != nil {
		in, out := &in.ShutdownManager, &out.ShutdownManager
		*out = new(ShutdownManager)
		(*in).DeepCopyInto(*out)
	}
	if in.InitManager != nil {
		in, out := &in.InitManager, &out.InitManager
		*out = new(InitManager)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyDeploymentSpec.
func (in *EnvoyDeploymentSpec) DeepCopy() *EnvoyDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvoyDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyDeploymentStatus) DeepCopyInto(out *EnvoyDeploymentStatus) {
	*out = *in
	if in.DeploymentName != nil {
		in, out := &in.DeploymentName, &out.DeploymentName
		*out = new(string)
		**out = **in
	}
	if in.DeploymentStatus != nil {
		in, out := &in.DeploymentStatus, &out.DeploymentStatus
		*out = new(appsv1.DeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
	out.UnimplementedStatefulSetStatus = in.UnimplementedStatefulSetStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyDeploymentStatus.
func (in *EnvoyDeploymentStatus) DeepCopy() *EnvoyDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvoyDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitManager) DeepCopyInto(out *InitManager) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitManager.
func (in *InitManager) DeepCopy() *InitManager {
	if in == nil {
		return nil
	}
	out := new(InitManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PKIConfig) DeepCopyInto(out *PKIConfig) {
	*out = *in
	if in.RootCertificateAuthority != nil {
		in, out := &in.RootCertificateAuthority, &out.RootCertificateAuthority
		*out = new(CertificateOptions)
		**out = **in
	}
	if in.ServerCertificate != nil {
		in, out := &in.ServerCertificate, &out.ServerCertificate
		*out = new(CertificateOptions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PKIConfig.
func (in *PKIConfig) DeepCopy() *PKIConfig {
	if in == nil {
		return nil
	}
	out := new(PKIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeSpec) DeepCopyInto(out *ProbeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeSpec.
func (in *ProbeSpec) DeepCopy() *ProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicasSpec) DeepCopyInto(out *ReplicasSpec) {
	*out = *in
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(int32)
		**out = **in
	}
	if in.Dynamic != nil {
		in, out := &in.Dynamic, &out.Dynamic
		*out = new(DynamicReplicasSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicasSpec.
func (in *ReplicasSpec) DeepCopy() *ReplicasSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedConfig) DeepCopyInto(out *SelfSignedConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedConfig.
func (in *SelfSignedConfig) DeepCopy() *SelfSignedConfig {
	if in == nil {
		return nil
	}
	out := new(SelfSignedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfig) DeepCopyInto(out *ServiceConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfig.
func (in *ServiceConfig) DeepCopy() *ServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShutdownManager) DeepCopyInto(out *ShutdownManager) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ServerPort != nil {
		in, out := &in.ServerPort, &out.ServerPort
		*out = new(uint32)
		**out = **in
	}
	if in.DrainTime != nil {
		in, out := &in.DrainTime, &out.DrainTime
		*out = new(int64)
		**out = **in
	}
	if in.DrainStrategy != nil {
		in, out := &in.DrainStrategy, &out.DrainStrategy
		*out = new(defaults.DrainStrategy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShutdownManager.
func (in *ShutdownManager) DeepCopy() *ShutdownManager {
	if in == nil {
		return nil
	}
	out := new(ShutdownManager)
	in.DeepCopyInto(out)
	return out
}
