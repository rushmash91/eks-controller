// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package addon

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customPreCompare(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AddonVersion, b.ko.Spec.AddonVersion) {
		delta.Add("Spec.AddonVersion", a.ko.Spec.AddonVersion, b.ko.Spec.AddonVersion)
	} else if a.ko.Spec.AddonVersion != nil && b.ko.Spec.AddonVersion != nil {
		if *a.ko.Spec.AddonVersion != *b.ko.Spec.AddonVersion {
			delta.Add("Spec.AddonVersion", a.ko.Spec.AddonVersion, b.ko.Spec.AddonVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ClientRequestToken, b.ko.Spec.ClientRequestToken) {
		delta.Add("Spec.ClientRequestToken", a.ko.Spec.ClientRequestToken, b.ko.Spec.ClientRequestToken)
	} else if a.ko.Spec.ClientRequestToken != nil && b.ko.Spec.ClientRequestToken != nil {
		if *a.ko.Spec.ClientRequestToken != *b.ko.Spec.ClientRequestToken {
			delta.Add("Spec.ClientRequestToken", a.ko.Spec.ClientRequestToken, b.ko.Spec.ClientRequestToken)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ClusterName, b.ko.Spec.ClusterName) {
		delta.Add("Spec.ClusterName", a.ko.Spec.ClusterName, b.ko.Spec.ClusterName)
	} else if a.ko.Spec.ClusterName != nil && b.ko.Spec.ClusterName != nil {
		if *a.ko.Spec.ClusterName != *b.ko.Spec.ClusterName {
			delta.Add("Spec.ClusterName", a.ko.Spec.ClusterName, b.ko.Spec.ClusterName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ClusterRef, b.ko.Spec.ClusterRef) {
		delta.Add("Spec.ClusterRef", a.ko.Spec.ClusterRef, b.ko.Spec.ClusterRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ConfigurationValues, b.ko.Spec.ConfigurationValues) {
		delta.Add("Spec.ConfigurationValues", a.ko.Spec.ConfigurationValues, b.ko.Spec.ConfigurationValues)
	} else if a.ko.Spec.ConfigurationValues != nil && b.ko.Spec.ConfigurationValues != nil {
		if *a.ko.Spec.ConfigurationValues != *b.ko.Spec.ConfigurationValues {
			delta.Add("Spec.ConfigurationValues", a.ko.Spec.ConfigurationValues, b.ko.Spec.ConfigurationValues)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResolveConflicts, b.ko.Spec.ResolveConflicts) {
		delta.Add("Spec.ResolveConflicts", a.ko.Spec.ResolveConflicts, b.ko.Spec.ResolveConflicts)
	} else if a.ko.Spec.ResolveConflicts != nil && b.ko.Spec.ResolveConflicts != nil {
		if *a.ko.Spec.ResolveConflicts != *b.ko.Spec.ResolveConflicts {
			delta.Add("Spec.ResolveConflicts", a.ko.Spec.ResolveConflicts, b.ko.Spec.ResolveConflicts)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ServiceAccountRoleARN, b.ko.Spec.ServiceAccountRoleARN) {
		delta.Add("Spec.ServiceAccountRoleARN", a.ko.Spec.ServiceAccountRoleARN, b.ko.Spec.ServiceAccountRoleARN)
	} else if a.ko.Spec.ServiceAccountRoleARN != nil && b.ko.Spec.ServiceAccountRoleARN != nil {
		if *a.ko.Spec.ServiceAccountRoleARN != *b.ko.Spec.ServiceAccountRoleARN {
			delta.Add("Spec.ServiceAccountRoleARN", a.ko.Spec.ServiceAccountRoleARN, b.ko.Spec.ServiceAccountRoleARN)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ServiceAccountRoleRef, b.ko.Spec.ServiceAccountRoleRef) {
		delta.Add("Spec.ServiceAccountRoleRef", a.ko.Spec.ServiceAccountRoleRef, b.ko.Spec.ServiceAccountRoleRef)
	}
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
