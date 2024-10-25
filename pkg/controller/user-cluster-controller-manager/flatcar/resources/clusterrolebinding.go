/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package resources

import (
	"k8c.io/reconciler/pkg/reconciling"

	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	operatorClusterRoleBindingName = "flatcar-linux-update-operator"
	agentClusterRoleBindingName    = "flatcar-linux-update-agent"
)

func OperatorClusterRoleBindingReconciler(operatorNamespace string) reconciling.NamedClusterRoleBindingReconcilerFactory {
	return func() (string, reconciling.ClusterRoleBindingReconciler) {
		return operatorClusterRoleBindingName, func(crb *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			crb.RoleRef = rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     operatorClusterRoleName,
			}
			crb.Subjects = []rbacv1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      operatorServiceAccountName,
					Namespace: operatorNamespace,
				},
			}
			return crb, nil
		}
	}
}

func AgentClusterRoleBindingReconciler(operatorNamespace string) reconciling.NamedClusterRoleBindingReconcilerFactory {
	return func() (string, reconciling.ClusterRoleBindingReconciler) {
		return agentClusterRoleBindingName, func(crb *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			crb.RoleRef = rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     agentClusterRoleName,
			}
			crb.Subjects = []rbacv1.Subject{
				{
					Kind:      "ServiceAccount",
					Name:      agentServiceAccountName,
					Namespace: operatorNamespace,
				},
			}
			return crb, nil
		}
	}
}
