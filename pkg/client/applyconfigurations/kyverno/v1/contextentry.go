/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ContextEntryApplyConfiguration represents an declarative configuration of the ContextEntry type for use
// with apply.
type ContextEntryApplyConfiguration struct {
	Name            *string                                        `json:"name,omitempty"`
	ConfigMap       *ConfigMapReferenceApplyConfiguration          `json:"configMap,omitempty"`
	APICall         *ContextAPICallApplyConfiguration              `json:"apiCall,omitempty"`
	ImageRegistry   *ImageRegistryApplyConfiguration               `json:"imageRegistry,omitempty"`
	Variable        *VariableApplyConfiguration                    `json:"variable,omitempty"`
	GlobalReference *GlobalContextEntryReferenceApplyConfiguration `json:"globalReference,omitempty"`
}

// ContextEntryApplyConfiguration constructs an declarative configuration of the ContextEntry type for use with
// apply.
func ContextEntry() *ContextEntryApplyConfiguration {
	return &ContextEntryApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ContextEntryApplyConfiguration) WithName(value string) *ContextEntryApplyConfiguration {
	b.Name = &value
	return b
}

// WithConfigMap sets the ConfigMap field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMap field is set to the value of the last call.
func (b *ContextEntryApplyConfiguration) WithConfigMap(value *ConfigMapReferenceApplyConfiguration) *ContextEntryApplyConfiguration {
	b.ConfigMap = value
	return b
}

// WithAPICall sets the APICall field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APICall field is set to the value of the last call.
func (b *ContextEntryApplyConfiguration) WithAPICall(value *ContextAPICallApplyConfiguration) *ContextEntryApplyConfiguration {
	b.APICall = value
	return b
}

// WithImageRegistry sets the ImageRegistry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImageRegistry field is set to the value of the last call.
func (b *ContextEntryApplyConfiguration) WithImageRegistry(value *ImageRegistryApplyConfiguration) *ContextEntryApplyConfiguration {
	b.ImageRegistry = value
	return b
}

// WithVariable sets the Variable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Variable field is set to the value of the last call.
func (b *ContextEntryApplyConfiguration) WithVariable(value *VariableApplyConfiguration) *ContextEntryApplyConfiguration {
	b.Variable = value
	return b
}

// WithGlobalReference sets the GlobalReference field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GlobalReference field is set to the value of the last call.
func (b *ContextEntryApplyConfiguration) WithGlobalReference(value *GlobalContextEntryReferenceApplyConfiguration) *ContextEntryApplyConfiguration {
	b.GlobalReference = value
	return b
}