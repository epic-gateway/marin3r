/*


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

package controllers

import (
	"context"

	envoyv1alpha1 "github.com/3scale/marin3r/apis/envoy/v1alpha1"

	"github.com/go-logr/logr"
	"github.com/operator-framework/operator-lib/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type SecretReconciler struct {
	Client client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=envoy.marin3r.3scale.net,resources=envoyconfigs,verbs=get;list;watch;patch

func (r *SecretReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {

	ctx := context.Background()

	// Fetch the Secret instance
	secret := &corev1.Secret{}
	err := r.Client.Get(ctx, req.NamespacedName, secret)
	if err != nil {
		// Error reading the object - requeue the request.
		// NOTE: We skip the IsNotFound error because we want to trigger EnvoyConfig
		// reconciles when referred secrets are deleted so the envoy control-plane
		// stops publishing them. This might cause errors if the reference hasn't been
		// removed from the EnvoyConfig, but that's ok as we do want to surface this
		// inconsistency instead of silently failing
		if !errors.IsNotFound(err) {
			return reconcile.Result{}, err
		}
	}

	_ = r.Log.WithValues("Namespace", req.Namespace, "Name", req.Name)
	r.Log.Info("Reconciling from 'kubernetes.io/tls' Secret")

	// Get the list of EnvoyConfigRevisions published and
	// check which of them contain refs to this secret
	list := &envoyv1alpha1.EnvoyConfigRevisionList{}
	if err := r.Client.List(ctx, list); err != nil {
		return reconcile.Result{}, err
	}

	for _, ecr := range list.Items {

		if ecr.Status.Conditions.IsTrueFor(envoyv1alpha1.RevisionPublishedCondition) {

			for _, secret := range ecr.Spec.EnvoyResources.Secrets {
				if secret.Ref.Name == req.Name && secret.Ref.Namespace == req.Namespace {
					r.Log.Info("Triggered EnvoyConfigRevision reconcile",
						"EnvoyConfigRevision_Name", ecr.ObjectMeta.Name, "EnvoyConfigRevision_Namespace", ecr.ObjectMeta.Namespace)
					if err != nil {
						return reconcile.Result{}, err
					}

					if !ecr.Status.Conditions.IsTrueFor(envoyv1alpha1.ResourcesOutOfSyncCondition) {
						// patch operation to update Spec.Version in the cache
						patch := client.MergeFrom(ecr.DeepCopy())
						ecr.Status.Conditions.SetCondition(status.Condition{
							Type:    envoyv1alpha1.ResourcesOutOfSyncCondition,
							Reason:  "SecretChanged",
							Message: "A secret relevant to this envoyconfigrevision changed",
							Status:  corev1.ConditionTrue,
						})
						if err := r.Client.Status().Patch(ctx, &ecr, patch); err != nil {
							return reconcile.Result{}, err
						}
						r.Log.V(1).Info("Condition should have been added ...")
					}
				}
			}
		}
	}

	return reconcile.Result{}, nil
}

func (r *SecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Secret{}).
		Complete(r)
}
