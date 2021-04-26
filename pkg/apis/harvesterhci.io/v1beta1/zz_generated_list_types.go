/*
Copyright 2021 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

// +k8s:deepcopy-gen=package
// +groupName=harvesterhci.io
package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeyPairList is a list of KeyPair resources
type KeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KeyPair `json:"items"`
}

func NewKeyPair(namespace, name string, obj KeyPair) *KeyPair {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("KeyPair").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PreferenceList is a list of Preference resources
type PreferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Preference `json:"items"`
}

func NewPreference(namespace, name string, obj Preference) *Preference {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Preference").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SettingList is a list of Setting resources
type SettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Setting `json:"items"`
}

func NewSetting(namespace, name string, obj Setting) *Setting {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Setting").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradeList is a list of Upgrade resources
type UpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Upgrade `json:"items"`
}

func NewUpgrade(namespace, name string, obj Upgrade) *Upgrade {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("Upgrade").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserList is a list of User resources
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []User `json:"items"`
}

func NewUser(namespace, name string, obj User) *User {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("User").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineBackupList is a list of VirtualMachineBackup resources
type VirtualMachineBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualMachineBackup `json:"items"`
}

func NewVirtualMachineBackup(namespace, name string, obj VirtualMachineBackup) *VirtualMachineBackup {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineBackup").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineBackupContentList is a list of VirtualMachineBackupContent resources
type VirtualMachineBackupContentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualMachineBackupContent `json:"items"`
}

func NewVirtualMachineBackupContent(namespace, name string, obj VirtualMachineBackupContent) *VirtualMachineBackupContent {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineBackupContent").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineRestoreList is a list of VirtualMachineRestore resources
type VirtualMachineRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualMachineRestore `json:"items"`
}

func NewVirtualMachineRestore(namespace, name string, obj VirtualMachineRestore) *VirtualMachineRestore {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineRestore").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineImageList is a list of VirtualMachineImage resources
type VirtualMachineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualMachineImage `json:"items"`
}

func NewVirtualMachineImage(namespace, name string, obj VirtualMachineImage) *VirtualMachineImage {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineImage").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineTemplateList is a list of VirtualMachineTemplate resources
type VirtualMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualMachineTemplate `json:"items"`
}

func NewVirtualMachineTemplate(namespace, name string, obj VirtualMachineTemplate) *VirtualMachineTemplate {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineTemplate").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineTemplateVersionList is a list of VirtualMachineTemplateVersion resources
type VirtualMachineTemplateVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VirtualMachineTemplateVersion `json:"items"`
}

func NewVirtualMachineTemplateVersion(namespace, name string, obj VirtualMachineTemplateVersion) *VirtualMachineTemplateVersion {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("VirtualMachineTemplateVersion").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SupportBundleList is a list of SupportBundle resources
type SupportBundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SupportBundle `json:"items"`
}

func NewSupportBundle(namespace, name string, obj SupportBundle) *SupportBundle {
	obj.APIVersion, obj.Kind = SchemeGroupVersion.WithKind("SupportBundle").ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}
