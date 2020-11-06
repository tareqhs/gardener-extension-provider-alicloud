// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/gardener/gardener-extension-provider-alicloud/pkg/apis/config"
	healthcheckconfig "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config"
	healthcheckconfigv1alpha1 "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config/v1alpha1"
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentbaseconfig "k8s.io/component-base/config"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ControllerConfiguration)(nil), (*config.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(a.(*ControllerConfiguration), b.(*config.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerConfiguration)(nil), (*ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(a.(*config.ControllerConfiguration), b.(*ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCD)(nil), (*config.ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCD_To_config_ETCD(a.(*ETCD), b.(*config.ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCD)(nil), (*ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCD_To_v1alpha1_ETCD(a.(*config.ETCD), b.(*ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCDBackup)(nil), (*config.ETCDBackup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(a.(*ETCDBackup), b.(*config.ETCDBackup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCDBackup)(nil), (*ETCDBackup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(a.(*config.ETCDBackup), b.(*ETCDBackup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCDStorage)(nil), (*config.ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(a.(*ETCDStorage), b.(*config.ETCDStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCDStorage)(nil), (*ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(a.(*config.ETCDStorage), b.(*ETCDStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeAPIServer)(nil), (*config.KubeAPIServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_KubeAPIServer_To_config_KubeAPIServer(a.(*KubeAPIServer), b.(*config.KubeAPIServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeAPIServer)(nil), (*KubeAPIServer)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeAPIServer_To_v1alpha1_KubeAPIServer(a.(*config.KubeAPIServer), b.(*KubeAPIServer), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Service)(nil), (*config.Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Service_To_config_Service(a.(*Service), b.(*config.Service), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Service)(nil), (*Service)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Service_To_v1alpha1_Service(a.(*config.Service), b.(*Service), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	out.ClientConnection = (*componentbaseconfig.ClientConnectionConfiguration)(unsafe.Pointer(in.ClientConnection))
	out.MachineImageOwnerSecretRef = (*v1.SecretReference)(unsafe.Pointer(in.MachineImageOwnerSecretRef))
	out.WhitelistedImageIDs = *(*[]string)(unsafe.Pointer(&in.WhitelistedImageIDs))
	out.KubeAPIServer = (*config.KubeAPIServer)(unsafe.Pointer(in.KubeAPIServer))
	if err := Convert_v1alpha1_Service_To_config_Service(&in.Service, &out.Service, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ETCD_To_config_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*healthcheckconfig.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in, out, s)
}

func autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	out.ClientConnection = (*configv1alpha1.ClientConnectionConfiguration)(unsafe.Pointer(in.ClientConnection))
	out.MachineImageOwnerSecretRef = (*v1.SecretReference)(unsafe.Pointer(in.MachineImageOwnerSecretRef))
	out.WhitelistedImageIDs = *(*[]string)(unsafe.Pointer(&in.WhitelistedImageIDs))
	out.KubeAPIServer = (*KubeAPIServer)(unsafe.Pointer(in.KubeAPIServer))
	if err := Convert_config_Service_To_v1alpha1_Service(&in.Service, &out.Service, s); err != nil {
		return err
	}
	if err := Convert_config_ETCD_To_v1alpha1_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*healthcheckconfigv1alpha1.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	return nil
}

// Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration is an autogenerated conversion function.
func Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	if err := Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ETCD_To_config_ETCD is an autogenerated conversion function.
func Convert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCD_To_config_ETCD(in, out, s)
}

func autoConvert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	if err := Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ETCD_To_v1alpha1_ETCD is an autogenerated conversion function.
func Convert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	return autoConvert_config_ETCD_To_v1alpha1_ETCD(in, out, s)
}

func autoConvert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in *ETCDBackup, out *config.ETCDBackup, s conversion.Scope) error {
	out.Schedule = (*string)(unsafe.Pointer(in.Schedule))
	return nil
}

// Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup is an autogenerated conversion function.
func Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in *ETCDBackup, out *config.ETCDBackup, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in, out, s)
}

func autoConvert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in *config.ETCDBackup, out *ETCDBackup, s conversion.Scope) error {
	out.Schedule = (*string)(unsafe.Pointer(in.Schedule))
	return nil
}

// Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup is an autogenerated conversion function.
func Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in *config.ETCDBackup, out *ETCDBackup, s conversion.Scope) error {
	return autoConvert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in, out, s)
}

func autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage is an autogenerated conversion function.
func Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in, out, s)
}

func autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage is an autogenerated conversion function.
func Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	return autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in, out, s)
}

func autoConvert_v1alpha1_KubeAPIServer_To_config_KubeAPIServer(in *KubeAPIServer, out *config.KubeAPIServer, s conversion.Scope) error {
	out.MutateExternalTrafficPolicy = in.MutateExternalTrafficPolicy
	return nil
}

// Convert_v1alpha1_KubeAPIServer_To_config_KubeAPIServer is an autogenerated conversion function.
func Convert_v1alpha1_KubeAPIServer_To_config_KubeAPIServer(in *KubeAPIServer, out *config.KubeAPIServer, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeAPIServer_To_config_KubeAPIServer(in, out, s)
}

func autoConvert_config_KubeAPIServer_To_v1alpha1_KubeAPIServer(in *config.KubeAPIServer, out *KubeAPIServer, s conversion.Scope) error {
	out.MutateExternalTrafficPolicy = in.MutateExternalTrafficPolicy
	return nil
}

// Convert_config_KubeAPIServer_To_v1alpha1_KubeAPIServer is an autogenerated conversion function.
func Convert_config_KubeAPIServer_To_v1alpha1_KubeAPIServer(in *config.KubeAPIServer, out *KubeAPIServer, s conversion.Scope) error {
	return autoConvert_config_KubeAPIServer_To_v1alpha1_KubeAPIServer(in, out, s)
}

func autoConvert_v1alpha1_Service_To_config_Service(in *Service, out *config.Service, s conversion.Scope) error {
	out.BackendLoadBalancerSpec = in.BackendLoadBalancerSpec
	return nil
}

// Convert_v1alpha1_Service_To_config_Service is an autogenerated conversion function.
func Convert_v1alpha1_Service_To_config_Service(in *Service, out *config.Service, s conversion.Scope) error {
	return autoConvert_v1alpha1_Service_To_config_Service(in, out, s)
}

func autoConvert_config_Service_To_v1alpha1_Service(in *config.Service, out *Service, s conversion.Scope) error {
	out.BackendLoadBalancerSpec = in.BackendLoadBalancerSpec
	return nil
}

// Convert_config_Service_To_v1alpha1_Service is an autogenerated conversion function.
func Convert_config_Service_To_v1alpha1_Service(in *config.Service, out *Service, s conversion.Scope) error {
	return autoConvert_config_Service_To_v1alpha1_Service(in, out, s)
}
