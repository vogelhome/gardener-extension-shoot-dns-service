//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	service "github.com/gardener/gardener-extension-shoot-dns-service/pkg/apis/service"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*DNSConfig)(nil), (*service.DNSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSConfig_To_service_DNSConfig(a.(*DNSConfig), b.(*service.DNSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*service.DNSConfig)(nil), (*DNSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_service_DNSConfig_To_v1alpha1_DNSConfig(a.(*service.DNSConfig), b.(*DNSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DNSIncludeExclude)(nil), (*service.DNSIncludeExclude)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSIncludeExclude_To_service_DNSIncludeExclude(a.(*DNSIncludeExclude), b.(*service.DNSIncludeExclude), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*service.DNSIncludeExclude)(nil), (*DNSIncludeExclude)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_service_DNSIncludeExclude_To_v1alpha1_DNSIncludeExclude(a.(*service.DNSIncludeExclude), b.(*DNSIncludeExclude), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DNSProvider)(nil), (*service.DNSProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSProvider_To_service_DNSProvider(a.(*DNSProvider), b.(*service.DNSProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*service.DNSProvider)(nil), (*DNSProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_service_DNSProvider_To_v1alpha1_DNSProvider(a.(*service.DNSProvider), b.(*DNSProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DNSProviderReplication)(nil), (*service.DNSProviderReplication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSProviderReplication_To_service_DNSProviderReplication(a.(*DNSProviderReplication), b.(*service.DNSProviderReplication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*service.DNSProviderReplication)(nil), (*DNSProviderReplication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_service_DNSProviderReplication_To_v1alpha1_DNSProviderReplication(a.(*service.DNSProviderReplication), b.(*DNSProviderReplication), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_DNSConfig_To_service_DNSConfig(in *DNSConfig, out *service.DNSConfig, s conversion.Scope) error {
	out.DNSProviderReplication = (*service.DNSProviderReplication)(unsafe.Pointer(in.DNSProviderReplication))
	out.Providers = *(*[]service.DNSProvider)(unsafe.Pointer(&in.Providers))
	out.SyncProvidersFromShootSpecDNS = (*bool)(unsafe.Pointer(in.SyncProvidersFromShootSpecDNS))
	return nil
}

// Convert_v1alpha1_DNSConfig_To_service_DNSConfig is an autogenerated conversion function.
func Convert_v1alpha1_DNSConfig_To_service_DNSConfig(in *DNSConfig, out *service.DNSConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSConfig_To_service_DNSConfig(in, out, s)
}

func autoConvert_service_DNSConfig_To_v1alpha1_DNSConfig(in *service.DNSConfig, out *DNSConfig, s conversion.Scope) error {
	out.DNSProviderReplication = (*DNSProviderReplication)(unsafe.Pointer(in.DNSProviderReplication))
	out.Providers = *(*[]DNSProvider)(unsafe.Pointer(&in.Providers))
	out.SyncProvidersFromShootSpecDNS = (*bool)(unsafe.Pointer(in.SyncProvidersFromShootSpecDNS))
	return nil
}

// Convert_service_DNSConfig_To_v1alpha1_DNSConfig is an autogenerated conversion function.
func Convert_service_DNSConfig_To_v1alpha1_DNSConfig(in *service.DNSConfig, out *DNSConfig, s conversion.Scope) error {
	return autoConvert_service_DNSConfig_To_v1alpha1_DNSConfig(in, out, s)
}

func autoConvert_v1alpha1_DNSIncludeExclude_To_service_DNSIncludeExclude(in *DNSIncludeExclude, out *service.DNSIncludeExclude, s conversion.Scope) error {
	out.Include = *(*[]string)(unsafe.Pointer(&in.Include))
	out.Exclude = *(*[]string)(unsafe.Pointer(&in.Exclude))
	return nil
}

// Convert_v1alpha1_DNSIncludeExclude_To_service_DNSIncludeExclude is an autogenerated conversion function.
func Convert_v1alpha1_DNSIncludeExclude_To_service_DNSIncludeExclude(in *DNSIncludeExclude, out *service.DNSIncludeExclude, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSIncludeExclude_To_service_DNSIncludeExclude(in, out, s)
}

func autoConvert_service_DNSIncludeExclude_To_v1alpha1_DNSIncludeExclude(in *service.DNSIncludeExclude, out *DNSIncludeExclude, s conversion.Scope) error {
	out.Include = *(*[]string)(unsafe.Pointer(&in.Include))
	out.Exclude = *(*[]string)(unsafe.Pointer(&in.Exclude))
	return nil
}

// Convert_service_DNSIncludeExclude_To_v1alpha1_DNSIncludeExclude is an autogenerated conversion function.
func Convert_service_DNSIncludeExclude_To_v1alpha1_DNSIncludeExclude(in *service.DNSIncludeExclude, out *DNSIncludeExclude, s conversion.Scope) error {
	return autoConvert_service_DNSIncludeExclude_To_v1alpha1_DNSIncludeExclude(in, out, s)
}

func autoConvert_v1alpha1_DNSProvider_To_service_DNSProvider(in *DNSProvider, out *service.DNSProvider, s conversion.Scope) error {
	out.Domains = (*service.DNSIncludeExclude)(unsafe.Pointer(in.Domains))
	out.Primary = (*bool)(unsafe.Pointer(in.Primary))
	out.SecretName = (*string)(unsafe.Pointer(in.SecretName))
	out.Type = (*string)(unsafe.Pointer(in.Type))
	out.Zones = (*service.DNSIncludeExclude)(unsafe.Pointer(in.Zones))
	return nil
}

// Convert_v1alpha1_DNSProvider_To_service_DNSProvider is an autogenerated conversion function.
func Convert_v1alpha1_DNSProvider_To_service_DNSProvider(in *DNSProvider, out *service.DNSProvider, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSProvider_To_service_DNSProvider(in, out, s)
}

func autoConvert_service_DNSProvider_To_v1alpha1_DNSProvider(in *service.DNSProvider, out *DNSProvider, s conversion.Scope) error {
	out.Domains = (*DNSIncludeExclude)(unsafe.Pointer(in.Domains))
	out.Primary = (*bool)(unsafe.Pointer(in.Primary))
	out.SecretName = (*string)(unsafe.Pointer(in.SecretName))
	out.Type = (*string)(unsafe.Pointer(in.Type))
	out.Zones = (*DNSIncludeExclude)(unsafe.Pointer(in.Zones))
	return nil
}

// Convert_service_DNSProvider_To_v1alpha1_DNSProvider is an autogenerated conversion function.
func Convert_service_DNSProvider_To_v1alpha1_DNSProvider(in *service.DNSProvider, out *DNSProvider, s conversion.Scope) error {
	return autoConvert_service_DNSProvider_To_v1alpha1_DNSProvider(in, out, s)
}

func autoConvert_v1alpha1_DNSProviderReplication_To_service_DNSProviderReplication(in *DNSProviderReplication, out *service.DNSProviderReplication, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_DNSProviderReplication_To_service_DNSProviderReplication is an autogenerated conversion function.
func Convert_v1alpha1_DNSProviderReplication_To_service_DNSProviderReplication(in *DNSProviderReplication, out *service.DNSProviderReplication, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSProviderReplication_To_service_DNSProviderReplication(in, out, s)
}

func autoConvert_service_DNSProviderReplication_To_v1alpha1_DNSProviderReplication(in *service.DNSProviderReplication, out *DNSProviderReplication, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_service_DNSProviderReplication_To_v1alpha1_DNSProviderReplication is an autogenerated conversion function.
func Convert_service_DNSProviderReplication_To_v1alpha1_DNSProviderReplication(in *service.DNSProviderReplication, out *DNSProviderReplication, s conversion.Scope) error {
	return autoConvert_service_DNSProviderReplication_To_v1alpha1_DNSProviderReplication(in, out, s)
}
