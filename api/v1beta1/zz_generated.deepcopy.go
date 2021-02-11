// +build !ignore_autogenerated

/*


Copyright 2021 Red Hat, Inc.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationStatus) DeepCopyInto(out *AuthenticationStatus) {
	*out = *in
	if in.AuthenticationCredentialsFound != nil {
		in, out := &in.AuthenticationCredentialsFound, &out.AuthenticationCredentialsFound
		*out = new(bool)
		**out = **in
	}
	if in.ValidBasicAuth != nil {
		in, out := &in.ValidBasicAuth, &out.ValidBasicAuth
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationStatus.
func (in *AuthenticationStatus) DeepCopy() *AuthenticationStatus {
	if in == nil {
		return nil
	}
	out := new(AuthenticationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudDotRedHatSourceSpec) DeepCopyInto(out *CloudDotRedHatSourceSpec) {
	*out = *in
	if in.CreateSource != nil {
		in, out := &in.CreateSource, &out.CreateSource
		*out = new(bool)
		**out = **in
	}
	if in.CheckCycle != nil {
		in, out := &in.CheckCycle, &out.CheckCycle
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudDotRedHatSourceSpec.
func (in *CloudDotRedHatSourceSpec) DeepCopy() *CloudDotRedHatSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudDotRedHatSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudDotRedHatSourceStatus) DeepCopyInto(out *CloudDotRedHatSourceStatus) {
	*out = *in
	if in.SourceDefined != nil {
		in, out := &in.SourceDefined, &out.SourceDefined
		*out = new(bool)
		**out = **in
	}
	if in.CreateSource != nil {
		in, out := &in.CreateSource, &out.CreateSource
		*out = new(bool)
		**out = **in
	}
	in.LastSourceCheckTime.DeepCopyInto(&out.LastSourceCheckTime)
	if in.CheckCycle != nil {
		in, out := &in.CheckCycle, &out.CheckCycle
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudDotRedHatSourceStatus.
func (in *CloudDotRedHatSourceStatus) DeepCopy() *CloudDotRedHatSourceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudDotRedHatSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedObjectMetadata) DeepCopyInto(out *EmbeddedObjectMetadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedObjectMetadata.
func (in *EmbeddedObjectMetadata) DeepCopy() *EmbeddedObjectMetadata {
	if in == nil {
		return nil
	}
	out := new(EmbeddedObjectMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedPersistentVolumeClaim) DeepCopyInto(out *EmbeddedPersistentVolumeClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.EmbeddedObjectMetadata.DeepCopyInto(&out.EmbeddedObjectMetadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedPersistentVolumeClaim.
func (in *EmbeddedPersistentVolumeClaim) DeepCopy() *EmbeddedPersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(EmbeddedPersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KokuMetricsConfig) DeepCopyInto(out *KokuMetricsConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KokuMetricsConfig.
func (in *KokuMetricsConfig) DeepCopy() *KokuMetricsConfig {
	if in == nil {
		return nil
	}
	out := new(KokuMetricsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KokuMetricsConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KokuMetricsConfigList) DeepCopyInto(out *KokuMetricsConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KokuMetricsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KokuMetricsConfigList.
func (in *KokuMetricsConfigList) DeepCopy() *KokuMetricsConfigList {
	if in == nil {
		return nil
	}
	out := new(KokuMetricsConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KokuMetricsConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KokuMetricsConfigSpec) DeepCopyInto(out *KokuMetricsConfigSpec) {
	*out = *in
	out.Authentication = in.Authentication
	out.Packaging = in.Packaging
	in.Upload.DeepCopyInto(&out.Upload)
	in.PrometheusConfig.DeepCopyInto(&out.PrometheusConfig)
	in.Source.DeepCopyInto(&out.Source)
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(EmbeddedPersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KokuMetricsConfigSpec.
func (in *KokuMetricsConfigSpec) DeepCopy() *KokuMetricsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KokuMetricsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KokuMetricsConfigStatus) DeepCopyInto(out *KokuMetricsConfigStatus) {
	*out = *in
	in.Authentication.DeepCopyInto(&out.Authentication)
	in.Packaging.DeepCopyInto(&out.Packaging)
	in.Upload.DeepCopyInto(&out.Upload)
	in.Prometheus.DeepCopyInto(&out.Prometheus)
	out.Reports = in.Reports
	in.Source.DeepCopyInto(&out.Source)
	out.Storage = in.Storage
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(EmbeddedPersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KokuMetricsConfigStatus.
func (in *KokuMetricsConfigStatus) DeepCopy() *KokuMetricsConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KokuMetricsConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackagingSpec) DeepCopyInto(out *PackagingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackagingSpec.
func (in *PackagingSpec) DeepCopy() *PackagingSpec {
	if in == nil {
		return nil
	}
	out := new(PackagingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackagingStatus) DeepCopyInto(out *PackagingStatus) {
	*out = *in
	in.LastSuccessfulPackagingTime.DeepCopyInto(&out.LastSuccessfulPackagingTime)
	if in.MaxReports != nil {
		in, out := &in.MaxReports, &out.MaxReports
		*out = new(int64)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int64)
		**out = **in
	}
	if in.PackagedFiles != nil {
		in, out := &in.PackagedFiles, &out.PackagedFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReportCount != nil {
		in, out := &in.ReportCount, &out.ReportCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackagingStatus.
func (in *PackagingStatus) DeepCopy() *PackagingStatus {
	if in == nil {
		return nil
	}
	out := new(PackagingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSpec) DeepCopyInto(out *PrometheusSpec) {
	*out = *in
	if in.SkipTLSVerification != nil {
		in, out := &in.SkipTLSVerification, &out.SkipTLSVerification
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSpec.
func (in *PrometheusSpec) DeepCopy() *PrometheusSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusStatus) DeepCopyInto(out *PrometheusStatus) {
	*out = *in
	in.LastQueryStartTime.DeepCopyInto(&out.LastQueryStartTime)
	in.LastQuerySuccessTime.DeepCopyInto(&out.LastQuerySuccessTime)
	if in.SkipTLSVerification != nil {
		in, out := &in.SkipTLSVerification, &out.SkipTLSVerification
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusStatus.
func (in *PrometheusStatus) DeepCopy() *PrometheusStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportsStatus) DeepCopyInto(out *ReportsStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportsStatus.
func (in *ReportsStatus) DeepCopy() *ReportsStatus {
	if in == nil {
		return nil
	}
	out := new(ReportsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageStatus) DeepCopyInto(out *StorageStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageStatus.
func (in *StorageStatus) DeepCopy() *StorageStatus {
	if in == nil {
		return nil
	}
	out := new(StorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UploadSpec) DeepCopyInto(out *UploadSpec) {
	*out = *in
	if in.UploadWait != nil {
		in, out := &in.UploadWait, &out.UploadWait
		*out = new(int64)
		**out = **in
	}
	if in.UploadCycle != nil {
		in, out := &in.UploadCycle, &out.UploadCycle
		*out = new(int64)
		**out = **in
	}
	if in.UploadToggle != nil {
		in, out := &in.UploadToggle, &out.UploadToggle
		*out = new(bool)
		**out = **in
	}
	if in.ValidateCert != nil {
		in, out := &in.ValidateCert, &out.ValidateCert
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UploadSpec.
func (in *UploadSpec) DeepCopy() *UploadSpec {
	if in == nil {
		return nil
	}
	out := new(UploadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UploadStatus) DeepCopyInto(out *UploadStatus) {
	*out = *in
	if in.UploadToggle != nil {
		in, out := &in.UploadToggle, &out.UploadToggle
		*out = new(bool)
		**out = **in
	}
	if in.UploadWait != nil {
		in, out := &in.UploadWait, &out.UploadWait
		*out = new(int64)
		**out = **in
	}
	if in.UploadCycle != nil {
		in, out := &in.UploadCycle, &out.UploadCycle
		*out = new(int64)
		**out = **in
	}
	in.LastUploadTime.DeepCopyInto(&out.LastUploadTime)
	in.LastSuccessfulUploadTime.DeepCopyInto(&out.LastSuccessfulUploadTime)
	if in.ValidateCert != nil {
		in, out := &in.ValidateCert, &out.ValidateCert
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UploadStatus.
func (in *UploadStatus) DeepCopy() *UploadStatus {
	if in == nil {
		return nil
	}
	out := new(UploadStatus)
	in.DeepCopyInto(out)
	return out
}
