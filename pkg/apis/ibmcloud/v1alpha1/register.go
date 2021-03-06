/*
Copyright 2019 The Kubernetes Authors.

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

// NOTE: Boilerplate only.  Ignore this file.

// Package v1alpha1 contains API Schema definitions for the ibmcloud v1alpha1 API group
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=sigs.k8s.io/cluster-api-provider-ibmcloud/pkg/apis/ibmcloud
// +k8s:defaulter-gen=TypeMeta
// +groupName=ibmcloud.cluster.k8s.io
package v1alpha1

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"sigs.k8s.io/yaml"

	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "ibmcloud.cluster.k8s.io", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is required by pkg/client/...
	AddToScheme = SchemeBuilder.AddToScheme
)

// Resource is required by pkg/client/listers/...
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// ClusterConfigFromProviderSpec unmarshals a provider config into an IBMCloud Cluster type
func ClusterSpecFromProviderSpec(providerSpec clusterv1.ProviderSpec) (*IBMCloudClusterProviderSpec, error) {
	if providerSpec.Value == nil {
		return nil, errors.New("no such providerSpec found in manifest")
	}

	var config IBMCloudClusterProviderSpec
	if err := yaml.Unmarshal(providerSpec.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// ClusterStatusFromProviderStatus unmarshals a provider status into an IBMCloud Cluster Status type
func ClusterStatusFromProviderStatus(extension *runtime.RawExtension) (*IBMCloudClusterProviderStatus, error) {
	if extension == nil {
		return &IBMCloudClusterProviderStatus{}, nil
	}

	status := new(IBMCloudClusterProviderStatus)
	if err := yaml.Unmarshal(extension.Raw, status); err != nil {
		return nil, err
	}

	return status, nil
}

func MachineSpecFromProviderSpec(providerSpec clusterv1.ProviderSpec) (*IBMCloudMachineProviderSpec, error) {
	if providerSpec.Value == nil {
		return nil, errors.New("no such providerSpec found in manifest")
	}

	var config IBMCloudMachineProviderSpec
	if err := yaml.Unmarshal(providerSpec.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func EncodeClusterStatus(status *IBMCloudClusterProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	var rawBytes []byte
	var err error

	//  TODO: use apimachinery conversion https://godoc.org/k8s.io/apimachinery/pkg/runtime#Convert_runtime_Object_To_runtime_RawExtension
	if rawBytes, err = json.Marshal(status); err != nil {
		return nil, err
	}

	return &runtime.RawExtension{
		Raw: rawBytes,
	}, nil
}

func EncodeMachineStatus(status *IBMCloudMachineProviderStatus) (*runtime.RawExtension, error) {
	if status == nil {
		return &runtime.RawExtension{}, nil
	}

	var rawBytes []byte
	var err error

	//  TODO: use apimachinery conversion https://godoc.org/k8s.io/apimachinery/pkg/runtime#Convert_runtime_Object_To_runtime_RawExtension
	if rawBytes, err = json.Marshal(status); err != nil {
		return nil, err
	}

	return &runtime.RawExtension{
		Raw: rawBytes,
	}, nil
}
