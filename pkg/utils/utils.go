// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetNamespacedName(obj client.Object) types.NamespacedName {
	return types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
}

func GetObjName(obj client.Object) string {
	return obj.GetName()
}

func IsDelete(obj client.Object) bool {
	return obj != nil && !obj.GetDeletionTimestamp().IsZero()
}

func IsServiceReady(svc *corev1.Service) bool {
	if svc == nil {
		return false
	}

	return true
}
