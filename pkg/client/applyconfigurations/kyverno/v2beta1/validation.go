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

package v2beta1

import (
	kyverno "github.com/kyverno/kyverno/api/kyverno"
	v1 "github.com/kyverno/kyverno/api/kyverno/v1"
	kyvernov1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v1"
)

// ValidationApplyConfiguration represents an declarative configuration of the Validation type for use
// with apply.
type ValidationApplyConfiguration struct {
	ValidationFailureAction          *v1.ValidationFailureAction                                   `json:"validationFailureAction,omitempty"`
	ValidationFailureActionOverrides []kyvernov1.ValidationFailureActionOverrideApplyConfiguration `json:"validationFailureActionOverrides,omitempty"`
	Message                          *string                                                       `json:"message,omitempty"`
	Manifests                        *kyvernov1.ManifestsApplyConfiguration                        `json:"manifests,omitempty"`
	ForEachValidation                []kyvernov1.ForEachValidationApplyConfiguration               `json:"foreach,omitempty"`
	RawPattern                       *kyverno.Any                                                  `json:"pattern,omitempty"`
	RawAnyPattern                    *kyverno.Any                                                  `json:"anyPattern,omitempty"`
	Deny                             *DenyApplyConfiguration                                       `json:"deny,omitempty"`
	PodSecurity                      *kyvernov1.PodSecurityApplyConfiguration                      `json:"podSecurity,omitempty"`
	CEL                              *kyvernov1.CELApplyConfiguration                              `json:"cel,omitempty"`
}

// ValidationApplyConfiguration constructs an declarative configuration of the Validation type for use with
// apply.
func Validation() *ValidationApplyConfiguration {
	return &ValidationApplyConfiguration{}
}

// WithValidationFailureAction sets the ValidationFailureAction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ValidationFailureAction field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithValidationFailureAction(value v1.ValidationFailureAction) *ValidationApplyConfiguration {
	b.ValidationFailureAction = &value
	return b
}

// WithValidationFailureActionOverrides adds the given value to the ValidationFailureActionOverrides field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ValidationFailureActionOverrides field.
func (b *ValidationApplyConfiguration) WithValidationFailureActionOverrides(values ...*kyvernov1.ValidationFailureActionOverrideApplyConfiguration) *ValidationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithValidationFailureActionOverrides")
		}
		b.ValidationFailureActionOverrides = append(b.ValidationFailureActionOverrides, *values[i])
	}
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithMessage(value string) *ValidationApplyConfiguration {
	b.Message = &value
	return b
}

// WithManifests sets the Manifests field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Manifests field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithManifests(value *kyvernov1.ManifestsApplyConfiguration) *ValidationApplyConfiguration {
	b.Manifests = value
	return b
}

// WithForEachValidation adds the given value to the ForEachValidation field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ForEachValidation field.
func (b *ValidationApplyConfiguration) WithForEachValidation(values ...*kyvernov1.ForEachValidationApplyConfiguration) *ValidationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithForEachValidation")
		}
		b.ForEachValidation = append(b.ForEachValidation, *values[i])
	}
	return b
}

// WithRawPattern sets the RawPattern field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RawPattern field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithRawPattern(value kyverno.Any) *ValidationApplyConfiguration {
	b.RawPattern = &value
	return b
}

// WithRawAnyPattern sets the RawAnyPattern field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RawAnyPattern field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithRawAnyPattern(value kyverno.Any) *ValidationApplyConfiguration {
	b.RawAnyPattern = &value
	return b
}

// WithDeny sets the Deny field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Deny field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithDeny(value *DenyApplyConfiguration) *ValidationApplyConfiguration {
	b.Deny = value
	return b
}

// WithPodSecurity sets the PodSecurity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodSecurity field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithPodSecurity(value *kyvernov1.PodSecurityApplyConfiguration) *ValidationApplyConfiguration {
	b.PodSecurity = value
	return b
}

// WithCEL sets the CEL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CEL field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithCEL(value *kyvernov1.CELApplyConfiguration) *ValidationApplyConfiguration {
	b.CEL = value
	return b
}
