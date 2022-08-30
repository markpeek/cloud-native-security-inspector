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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Assessment) DeepCopyInto(out *Assessment) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Assessment.
func (in *Assessment) DeepCopy() *Assessment {
	if in == nil {
		return nil
	}
	out := new(Assessment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentError) DeepCopyInto(out *AssessmentError) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentError.
func (in *AssessmentError) DeepCopy() *AssessmentError {
	if in == nil {
		return nil
	}
	out := new(AssessmentError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentFailure) DeepCopyInto(out *AssessmentFailure) {
	*out = *in
	if in.Baseline != nil {
		in, out := &in.Baseline, &out.Baseline
		*out = new(ComplianceBaseline)
		**out = **in
	}
	out.Container = in.Container
	out.AssessmentError = in.AssessmentError
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentFailure.
func (in *AssessmentFailure) DeepCopy() *AssessmentFailure {
	if in == nil {
		return nil
	}
	out := new(AssessmentFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentReport) DeepCopyInto(out *AssessmentReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentReport.
func (in *AssessmentReport) DeepCopy() *AssessmentReport {
	if in == nil {
		return nil
	}
	out := new(AssessmentReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssessmentReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentReportList) DeepCopyInto(out *AssessmentReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AssessmentReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentReportList.
func (in *AssessmentReportList) DeepCopy() *AssessmentReportList {
	if in == nil {
		return nil
	}
	out := new(AssessmentReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssessmentReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentReportSpec) DeepCopyInto(out *AssessmentReportSpec) {
	*out = *in
	in.InspectionConfiguration.DeepCopyInto(&out.InspectionConfiguration)
	if in.NamespaceAssessments != nil {
		in, out := &in.NamespaceAssessments, &out.NamespaceAssessments
		*out = make([]*NamespaceAssessment, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NamespaceAssessment)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentReportSpec.
func (in *AssessmentReportSpec) DeepCopy() *AssessmentReportSpec {
	if in == nil {
		return nil
	}
	out := new(AssessmentReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssessmentReportStatus) DeepCopyInto(out *AssessmentReportStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssessmentReportStatus.
func (in *AssessmentReportStatus) DeepCopy() *AssessmentReportStatus {
	if in == nil {
		return nil
	}
	out := new(AssessmentReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cache) DeepCopyInto(out *Cache) {
	*out = *in
	if in.CredentialRef != nil {
		in, out := &in.CredentialRef, &out.CredentialRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(int)
		**out = **in
	}
	in.Settings.DeepCopyInto(&out.Settings)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cache.
func (in *Cache) DeepCopy() *Cache {
	if in == nil {
		return nil
	}
	out := new(Cache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSettings) DeepCopyInto(out *CacheSettings) {
	*out = *in
	if in.SkipTLSVerify != nil {
		in, out := &in.SkipTLSVerify, &out.SkipTLSVerify
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSettings.
func (in *CacheSettings) DeepCopy() *CacheSettings {
	if in == nil {
		return nil
	}
	out := new(CacheSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComplianceBaseline) DeepCopyInto(out *ComplianceBaseline) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComplianceBaseline.
func (in *ComplianceBaseline) DeepCopy() *ComplianceBaseline {
	if in == nil {
		return nil
	}
	out := new(ComplianceBaseline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Connection) DeepCopyInto(out *Connection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Connection.
func (in *Connection) DeepCopy() *Connection {
	if in == nil {
		return nil
	}
	out := new(Connection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataProvider) DeepCopyInto(out *DataProvider) {
	*out = *in
	if in.Credential != nil {
		in, out := &in.Credential, &out.Credential
		*out = new(Credential)
		**out = **in
	}
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(Cache)
		(*in).DeepCopyInto(*out)
	}
	out.Connection = in.Connection
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataProvider.
func (in *DataProvider) DeepCopy() *DataProvider {
	if in == nil {
		return nil
	}
	out := new(DataProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSource) DeepCopyInto(out *DataSource) {
	*out = *in
	in.Registry.DeepCopyInto(&out.Registry)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSource.
func (in *DataSource) DeepCopy() *DataSource {
	if in == nil {
		return nil
	}
	out := new(DataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnforcementResult) DeepCopyInto(out *EnforcementResult) {
	*out = *in
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnforcementResult.
func (in *EnforcementResult) DeepCopy() *EnforcementResult {
	if in == nil {
		return nil
	}
	out := new(EnforcementResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FollowupAction) DeepCopyInto(out *FollowupAction) {
	*out = *in
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FollowupAction.
func (in *FollowupAction) DeepCopy() *FollowupAction {
	if in == nil {
		return nil
	}
	out := new(FollowupAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FollowupActionEnforcement) DeepCopyInto(out *FollowupActionEnforcement) {
	*out = *in
	in.Action.DeepCopyInto(&out.Action)
	in.Result.DeepCopyInto(&out.Result)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FollowupActionEnforcement.
func (in *FollowupActionEnforcement) DeepCopy() *FollowupActionEnforcement {
	if in == nil {
		return nil
	}
	out := new(FollowupActionEnforcement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inspection) DeepCopyInto(out *Inspection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inspection.
func (in *Inspection) DeepCopy() *Inspection {
	if in == nil {
		return nil
	}
	out := new(Inspection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Inspection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionConfiguration) DeepCopyInto(out *InspectionConfiguration) {
	*out = *in
	out.Assessment = in.Assessment
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*FollowupAction, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FollowupAction)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Baselines != nil {
		in, out := &in.Baselines, &out.Baselines
		*out = make([]*ComplianceBaseline, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ComplianceBaseline)
				**out = **in
			}
		}
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkloadSelector != nil {
		in, out := &in.WorkloadSelector, &out.WorkloadSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionConfiguration.
func (in *InspectionConfiguration) DeepCopy() *InspectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(InspectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionList) DeepCopyInto(out *InspectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Inspection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionList.
func (in *InspectionList) DeepCopy() *InspectionList {
	if in == nil {
		return nil
	}
	out := new(InspectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicy) DeepCopyInto(out *InspectionPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicy.
func (in *InspectionPolicy) DeepCopy() *InspectionPolicy {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectionPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicyList) DeepCopyInto(out *InspectionPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InspectionPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicyList.
func (in *InspectionPolicyList) DeepCopy() *InspectionPolicyList {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectionPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicySpec) DeepCopyInto(out *InspectionPolicySpec) {
	*out = *in
	if in.WorkNamespace != nil {
		in, out := &in.WorkNamespace, &out.WorkNamespace
		*out = new(string)
		**out = **in
	}
	in.Inspection.DeepCopyInto(&out.Inspection)
	in.Strategy.DeepCopyInto(&out.Strategy)
	if in.Inspector != nil {
		in, out := &in.Inspector, &out.Inspector
		*out = new(Inspector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicySpec.
func (in *InspectionPolicySpec) DeepCopy() *InspectionPolicySpec {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionPolicyStatus) DeepCopyInto(out *InspectionPolicyStatus) {
	*out = *in
	if in.Executor != nil {
		in, out := &in.Executor, &out.Executor
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionPolicyStatus.
func (in *InspectionPolicyStatus) DeepCopy() *InspectionPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(InspectionPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionSpec) DeepCopyInto(out *InspectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionSpec.
func (in *InspectionSpec) DeepCopy() *InspectionSpec {
	if in == nil {
		return nil
	}
	out := new(InspectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectionStatus) DeepCopyInto(out *InspectionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*v1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectionStatus.
func (in *InspectionStatus) DeepCopy() *InspectionStatus {
	if in == nil {
		return nil
	}
	out := new(InspectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Inspector) DeepCopyInto(out *Inspector) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Inspector.
func (in *Inspector) DeepCopy() *Inspector {
	if in == nil {
		return nil
	}
	out := new(Inspector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnownRegistry) DeepCopyInto(out *KnownRegistry) {
	*out = *in
	in.Registry.DeepCopyInto(&out.Registry)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnownRegistry.
func (in *KnownRegistry) DeepCopy() *KnownRegistry {
	if in == nil {
		return nil
	}
	out := new(KnownRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAssessment) DeepCopyInto(out *NamespaceAssessment) {
	*out = *in
	out.Namespace = in.Namespace
	if in.WorkloadAssessments != nil {
		in, out := &in.WorkloadAssessments, &out.WorkloadAssessments
		*out = make([]*WorkloadAssessment, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(WorkloadAssessment)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAssessment.
func (in *NamespaceAssessment) DeepCopy() *NamespaceAssessment {
	if in == nil {
		return nil
	}
	out := new(NamespaceAssessment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]*Container, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Container)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.CredentialRef != nil {
		in, out := &in.CredentialRef, &out.CredentialRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Setting) DeepCopyInto(out *Setting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Setting.
func (in *Setting) DeepCopy() *Setting {
	if in == nil {
		return nil
	}
	out := new(Setting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Setting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingList) DeepCopyInto(out *SettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Setting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingList.
func (in *SettingList) DeepCopy() *SettingList {
	if in == nil {
		return nil
	}
	out := new(SettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingSpec) DeepCopyInto(out *SettingSpec) {
	*out = *in
	if in.KnownRegistries != nil {
		in, out := &in.KnownRegistries, &out.KnownRegistries
		*out = make([]KnownRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.DataSource.DeepCopyInto(&out.DataSource)
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(Cache)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingSpec.
func (in *SettingSpec) DeepCopy() *SettingSpec {
	if in == nil {
		return nil
	}
	out := new(SettingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingStatus) DeepCopyInto(out *SettingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingStatus.
func (in *SettingStatus) DeepCopy() *SettingStatus {
	if in == nil {
		return nil
	}
	out := new(SettingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Strategy) DeepCopyInto(out *Strategy) {
	*out = *in
	if in.HistoryLimit != nil {
		in, out := &in.HistoryLimit, &out.HistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Strategy.
func (in *Strategy) DeepCopy() *Strategy {
	if in == nil {
		return nil
	}
	out := new(Strategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	out.ObjectReference = in.ObjectReference
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]*Pod, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Pod)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadAssessment) DeepCopyInto(out *WorkloadAssessment) {
	*out = *in
	in.Workload.DeepCopyInto(&out.Workload)
	if in.Failures != nil {
		in, out := &in.Failures, &out.Failures
		*out = make([]*AssessmentFailure, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AssessmentFailure)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ActionEnforcements != nil {
		in, out := &in.ActionEnforcements, &out.ActionEnforcements
		*out = make([]*FollowupActionEnforcement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FollowupActionEnforcement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadAssessment.
func (in *WorkloadAssessment) DeepCopy() *WorkloadAssessment {
	if in == nil {
		return nil
	}
	out := new(WorkloadAssessment)
	in.DeepCopyInto(out)
	return out
}