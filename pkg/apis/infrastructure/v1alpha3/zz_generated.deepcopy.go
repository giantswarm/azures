// +build !ignore_autogenerated

/*
Copyright 2020 Giant Swarm GmbH.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCluster) DeepCopyInto(out *AzureCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCluster.
func (in *AzureCluster) DeepCopy() *AzureCluster {
	if in == nil {
		return nil
	}
	out := new(AzureCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterList) DeepCopyInto(out *AzureClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterList.
func (in *AzureClusterList) DeepCopy() *AzureClusterList {
	if in == nil {
		return nil
	}
	out := new(AzureClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpec) DeepCopyInto(out *AzureClusterSpec) {
	*out = *in
	out.Cluster = in.Cluster
	out.Provider = in.Provider
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpec.
func (in *AzureClusterSpec) DeepCopy() *AzureClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecCluster) DeepCopyInto(out *AzureClusterSpecCluster) {
	*out = *in
	out.DNS = in.DNS
	out.KubeProxy = in.KubeProxy
	out.OIDC = in.OIDC
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecCluster.
func (in *AzureClusterSpecCluster) DeepCopy() *AzureClusterSpecCluster {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecClusterDNS) DeepCopyInto(out *AzureClusterSpecClusterDNS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecClusterDNS.
func (in *AzureClusterSpecClusterDNS) DeepCopy() *AzureClusterSpecClusterDNS {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecClusterDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecClusterKubeProxy) DeepCopyInto(out *AzureClusterSpecClusterKubeProxy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecClusterKubeProxy.
func (in *AzureClusterSpecClusterKubeProxy) DeepCopy() *AzureClusterSpecClusterKubeProxy {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecClusterKubeProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecClusterOIDC) DeepCopyInto(out *AzureClusterSpecClusterOIDC) {
	*out = *in
	out.Claims = in.Claims
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecClusterOIDC.
func (in *AzureClusterSpecClusterOIDC) DeepCopy() *AzureClusterSpecClusterOIDC {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecClusterOIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecClusterOIDCClaims) DeepCopyInto(out *AzureClusterSpecClusterOIDCClaims) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecClusterOIDCClaims.
func (in *AzureClusterSpecClusterOIDCClaims) DeepCopy() *AzureClusterSpecClusterOIDCClaims {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecClusterOIDCClaims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecProvider) DeepCopyInto(out *AzureClusterSpecProvider) {
	*out = *in
	out.CredentialSecret = in.CredentialSecret
	out.Master = in.Master
	out.Pods = in.Pods
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecProvider.
func (in *AzureClusterSpecProvider) DeepCopy() *AzureClusterSpecProvider {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecProviderCredentialSecret) DeepCopyInto(out *AzureClusterSpecProviderCredentialSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecProviderCredentialSecret.
func (in *AzureClusterSpecProviderCredentialSecret) DeepCopy() *AzureClusterSpecProviderCredentialSecret {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecProviderCredentialSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecProviderMaster) DeepCopyInto(out *AzureClusterSpecProviderMaster) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecProviderMaster.
func (in *AzureClusterSpecProviderMaster) DeepCopy() *AzureClusterSpecProviderMaster {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecProviderMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpecProviderPods) DeepCopyInto(out *AzureClusterSpecProviderPods) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpecProviderPods.
func (in *AzureClusterSpecProviderPods) DeepCopy() *AzureClusterSpecProviderPods {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpecProviderPods)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterStatus) DeepCopyInto(out *AzureClusterStatus) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	out.Provider = in.Provider
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterStatus.
func (in *AzureClusterStatus) DeepCopy() *AzureClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AzureClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterStatusProvider) DeepCopyInto(out *AzureClusterStatusProvider) {
	*out = *in
	out.Network = in.Network
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterStatusProvider.
func (in *AzureClusterStatusProvider) DeepCopy() *AzureClusterStatusProvider {
	if in == nil {
		return nil
	}
	out := new(AzureClusterStatusProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterStatusProviderNetwork) DeepCopyInto(out *AzureClusterStatusProviderNetwork) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterStatusProviderNetwork.
func (in *AzureClusterStatusProviderNetwork) DeepCopy() *AzureClusterStatusProviderNetwork {
	if in == nil {
		return nil
	}
	out := new(AzureClusterStatusProviderNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureControlPlane) DeepCopyInto(out *AzureControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureControlPlane.
func (in *AzureControlPlane) DeepCopy() *AzureControlPlane {
	if in == nil {
		return nil
	}
	out := new(AzureControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureControlPlaneList) DeepCopyInto(out *AzureControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureControlPlaneList.
func (in *AzureControlPlaneList) DeepCopy() *AzureControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(AzureControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureControlPlaneSpec) DeepCopyInto(out *AzureControlPlaneSpec) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureControlPlaneSpec.
func (in *AzureControlPlaneSpec) DeepCopy() *AzureControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(AzureControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeployment) DeepCopyInto(out *AzureMachineDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeployment.
func (in *AzureMachineDeployment) DeepCopy() *AzureMachineDeployment {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentList) DeepCopyInto(out *AzureMachineDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureMachineDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentList.
func (in *AzureMachineDeploymentList) DeepCopy() *AzureMachineDeploymentList {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentSpec) DeepCopyInto(out *AzureMachineDeploymentSpec) {
	*out = *in
	out.NodePool = in.NodePool
	in.Provider.DeepCopyInto(&out.Provider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentSpec.
func (in *AzureMachineDeploymentSpec) DeepCopy() *AzureMachineDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentSpecInstanceDistribution) DeepCopyInto(out *AzureMachineDeploymentSpecInstanceDistribution) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentSpecInstanceDistribution.
func (in *AzureMachineDeploymentSpecInstanceDistribution) DeepCopy() *AzureMachineDeploymentSpecInstanceDistribution {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentSpecInstanceDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentSpecNodePool) DeepCopyInto(out *AzureMachineDeploymentSpecNodePool) {
	*out = *in
	out.Machine = in.Machine
	out.Scaling = in.Scaling
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentSpecNodePool.
func (in *AzureMachineDeploymentSpecNodePool) DeepCopy() *AzureMachineDeploymentSpecNodePool {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentSpecNodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentSpecNodePoolMachine) DeepCopyInto(out *AzureMachineDeploymentSpecNodePoolMachine) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentSpecNodePoolMachine.
func (in *AzureMachineDeploymentSpecNodePoolMachine) DeepCopy() *AzureMachineDeploymentSpecNodePoolMachine {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentSpecNodePoolMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentSpecNodePoolScaling) DeepCopyInto(out *AzureMachineDeploymentSpecNodePoolScaling) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentSpecNodePoolScaling.
func (in *AzureMachineDeploymentSpecNodePoolScaling) DeepCopy() *AzureMachineDeploymentSpecNodePoolScaling {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentSpecNodePoolScaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentSpecProvider) DeepCopyInto(out *AzureMachineDeploymentSpecProvider) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.InstanceDistribution = in.InstanceDistribution
	out.Worker = in.Worker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentSpecProvider.
func (in *AzureMachineDeploymentSpecProvider) DeepCopy() *AzureMachineDeploymentSpecProvider {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentSpecProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentSpecProviderWorker) DeepCopyInto(out *AzureMachineDeploymentSpecProviderWorker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentSpecProviderWorker.
func (in *AzureMachineDeploymentSpecProviderWorker) DeepCopy() *AzureMachineDeploymentSpecProviderWorker {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentSpecProviderWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentStatus) DeepCopyInto(out *AzureMachineDeploymentStatus) {
	*out = *in
	in.Provider.DeepCopyInto(&out.Provider)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentStatus.
func (in *AzureMachineDeploymentStatus) DeepCopy() *AzureMachineDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentStatusProvider) DeepCopyInto(out *AzureMachineDeploymentStatusProvider) {
	*out = *in
	in.Worker.DeepCopyInto(&out.Worker)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentStatusProvider.
func (in *AzureMachineDeploymentStatusProvider) DeepCopy() *AzureMachineDeploymentStatusProvider {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentStatusProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineDeploymentStatusProviderWorker) DeepCopyInto(out *AzureMachineDeploymentStatusProviderWorker) {
	*out = *in
	if in.InstanceTypes != nil {
		in, out := &in.InstanceTypes, &out.InstanceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineDeploymentStatusProviderWorker.
func (in *AzureMachineDeploymentStatusProviderWorker) DeepCopy() *AzureMachineDeploymentStatusProviderWorker {
	if in == nil {
		return nil
	}
	out := new(AzureMachineDeploymentStatusProviderWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonClusterStatus) DeepCopyInto(out *CommonClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CommonClusterStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]CommonClusterStatusVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonClusterStatus.
func (in *CommonClusterStatus) DeepCopy() *CommonClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CommonClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonClusterStatusCondition) DeepCopyInto(out *CommonClusterStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonClusterStatusCondition.
func (in *CommonClusterStatusCondition) DeepCopy() *CommonClusterStatusCondition {
	if in == nil {
		return nil
	}
	out := new(CommonClusterStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonClusterStatusVersion) DeepCopyInto(out *CommonClusterStatusVersion) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonClusterStatusVersion.
func (in *CommonClusterStatusVersion) DeepCopy() *CommonClusterStatusVersion {
	if in == nil {
		return nil
	}
	out := new(CommonClusterStatusVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlane) DeepCopyInto(out *G8sControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlane.
func (in *G8sControlPlane) DeepCopy() *G8sControlPlane {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *G8sControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlaneList) DeepCopyInto(out *G8sControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]G8sControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlaneList.
func (in *G8sControlPlaneList) DeepCopy() *G8sControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *G8sControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlaneSpec) DeepCopyInto(out *G8sControlPlaneSpec) {
	*out = *in
	out.InfrastructureRef = in.InfrastructureRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlaneSpec.
func (in *G8sControlPlaneSpec) DeepCopy() *G8sControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *G8sControlPlaneStatus) DeepCopyInto(out *G8sControlPlaneStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new G8sControlPlaneStatus.
func (in *G8sControlPlaneStatus) DeepCopy() *G8sControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(G8sControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}
