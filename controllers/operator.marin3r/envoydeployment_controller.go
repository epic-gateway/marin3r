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
	"fmt"
	"reflect"
	"time"

	marin3rv1alpha1 "github.com/3scale-ops/marin3r/apis/marin3r/v1alpha1"
	operatorv1alpha1 "github.com/3scale-ops/marin3r/apis/operator.marin3r/v1alpha1"
	"github.com/3scale-ops/marin3r/pkg/envoy/container/defaults"
	"github.com/3scale-ops/marin3r/pkg/reconcilers/lockedresources"
	"github.com/3scale-ops/marin3r/pkg/reconcilers/operator/envoydeployment/generators"
	"github.com/go-logr/logr"
	operatorutil "github.com/redhat-cop/operator-utils/pkg/util"
	"github.com/redhat-cop/operator-utils/pkg/util/lockedresourcecontroller/lockedpatch"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// EnvoyDeploymentReconciler reconciles a EnvoyDeployment object
type EnvoyDeploymentReconciler struct {
	lockedresources.Reconciler
	Log logr.Logger
}

//+kubebuilder:rbac:groups=operator.marin3r.3scale.net,namespace=placeholder,resources=envoydeployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=operator.marin3r.3scale.net,namespace=placeholder,resources=envoydeployments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=operator.marin3r.3scale.net,namespace=placeholder,resources=envoydeployments/finalizers,verbs=update
//+kubebuilder:rbac:groups="apps",namespace=placeholder,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="autoscaling",namespace=placeholder,resources=horizontalpodautoscalers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="policy",namespace=placeholder,resources=poddisruptionbudgets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=marin3r.3scale.net,namespace=placeholder,resources=envoyconfigs,verbs=get;list;watch
//+kubebuilder:rbac:groups=operator.marin3r.3scale.net,namespace=placeholder,resources=discoveryservices,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *EnvoyDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("envoydeployment", req.NamespacedName)

	ed := &operatorv1alpha1.EnvoyDeployment{}
	key := types.NamespacedName{Name: req.Name, Namespace: req.Namespace}
	err := r.GetClient().Get(ctx, key, ed)
	if err != nil {
		if errors.IsNotFound(err) {
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if ok := r.IsInitialized(ed, operatorv1alpha1.Finalizer); !ok {
		err := r.GetClient().Update(ctx, ed)
		if err != nil {
			log.Error(err, "unable to initialize instance")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	if operatorutil.IsBeingDeleted(ed) {
		if !operatorutil.HasFinalizer(ed, operatorv1alpha1.Finalizer) {
			return ctrl.Result{}, nil
		}
		err := r.ManageCleanUpLogic(ed, log)
		if err != nil {
			log.Error(err, "unable to delete instance")
			return ctrl.Result{}, err
		}
		operatorutil.RemoveFinalizer(ed, operatorv1alpha1.Finalizer)
		err = r.GetClient().Update(ctx, ed)
		if err != nil {
			log.Error(err, "unable to update instance")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// Get the address of the DiscoveryService instance
	ds := &operatorv1alpha1.DiscoveryService{}
	dsKey := types.NamespacedName{Name: ed.Spec.DiscoveryServiceRef, Namespace: ed.GetNamespace()}
	if err := r.GetClient().Get(ctx, dsKey, ds); err != nil {
		if errors.IsNotFound(err) {
			log.Error(err, "DiscoveryService does not exist", "DiscoveryService", ed.Spec.DiscoveryServiceRef)
		}
		return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, err
	}

	// Get the EnvoyConfig for additional data (like the envoy API version in use)
	ec, err := r.getEnvoyConfig(ctx, types.NamespacedName{Name: ed.Spec.EnvoyConfigRef, Namespace: ed.GetNamespace()})
	if err != nil {
		log.Error(err, "unable to get EnvoyConfig", "EnvoyConfig", ed.Spec.EnvoyConfigRef)
		return ctrl.Result{}, err
	}

	generate := generators.GeneratorOptions{
		InstanceName:         ed.GetName(),
		Namespace:            ed.GetNamespace(),
		DiscoveryServiceName: ed.Spec.DiscoveryServiceRef,
		XdssAdress:           fmt.Sprintf("%s.%s.%s", ds.GetServiceConfig().Name, ds.GetNamespace(), "svc"),
		XdssPort:             int(ds.GetXdsServerPort()),
		EnvoyAPIVersion:      ec.GetEnvoyAPIVersion(),
		EnvoyNodeID:          ec.Spec.NodeID,
		EnvoyClusterID: func() string {
			if ed.Spec.ClusterID != nil {
				return *ed.Spec.ClusterID
			}
			return ec.Spec.NodeID
		}(),
		ClientCertificateName:     fmt.Sprintf("%s-%s", defaults.DeploymentClientCertificate, ed.GetName()),
		ClientCertificateDuration: ed.ClientCertificateDuration(),
		SigningCertificateName:    ds.GetRootCertificateAuthorityOptions().SecretName,
		DeploymentImage:           ed.Image(),
		DeploymentResources:       ed.Resources(),
		ExposedPorts:              ed.Spec.Ports,
		ExtraArgs:                 ed.Spec.ExtraArgs,
		AdminPort:                 int32(ed.AdminPort()),
		AdminAccessLogPath:        ed.AdminAccessLogPath(),
		Replicas:                  ed.Replicas(),
		LivenessProbe:             ed.LivenessProbe(),
		ReadinessProbe:            ed.ReadinessProbe(),
		Affinity:                  ed.Affinity(),
		PodDisruptionBudget:       ed.PodDisruptionBudget(),
		ShutdownManager:           ed.Spec.ShutdownManager,
		InitManager:               ed.Spec.InitManager,
	}

	replicas, err := r.getDeploymentReplicas(ctx, generate.OwnedResourceKey())
	if err != nil {
		log.Error(err, "unable to get Deployment", "DeploymentName", key.Name)
		return ctrl.Result{}, err
	}

	lr := []lockedresources.LockedResource{
		{GeneratorFn: generate.ClientCertificate(), ExcludePaths: defaultExcludedPaths},
		{GeneratorFn: generate.Deployment(replicas), ExcludePaths: defaultExcludedPaths},
	}

	if ed.Replicas().Dynamic != nil {
		lr = append(lr, lockedresources.LockedResource{GeneratorFn: generate.HPA(), ExcludePaths: defaultExcludedPaths})
	}
	if !reflect.DeepEqual(ed.PodDisruptionBudget(), operatorv1alpha1.PodDisruptionBudgetSpec{}) {
		lr = append(lr, lockedresources.LockedResource{GeneratorFn: generate.PDB(), ExcludePaths: defaultExcludedPaths})
	}

	resources, err := r.NewLockedResources(lr, ed)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.UpdateLockedResources(ctx, ed, resources, []lockedpatch.LockedPatch{})
	if err != nil {
		log.Error(err, "unable to update locked resources")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *EnvoyDeploymentReconciler) getEnvoyConfig(ctx context.Context, key types.NamespacedName) (*marin3rv1alpha1.EnvoyConfig, error) {
	ec := &marin3rv1alpha1.EnvoyConfig{}
	err := r.GetClient().Get(ctx, key, ec)

	if err != nil {
		return nil, err
	}

	return ec, nil
}

// reconcileDeploymentReplicas: this is required when using dynamic number of replicas to avoid the controller from
// overriding the dynamic replica value set by the HPA
func (r *EnvoyDeploymentReconciler) getDeploymentReplicas(ctx context.Context, key types.NamespacedName) (*int32, error) {
	dep := &appsv1.Deployment{}
	err := r.GetClient().Get(ctx, key, dep)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return dep.Spec.Replicas, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *EnvoyDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1alpha1.EnvoyDeployment{}).
		Watches(&source.Channel{Source: r.GetStatusChangeChannel()}, &handler.EnqueueRequestForObject{}).
		Watches(&source.Kind{Type: &marin3rv1alpha1.EnvoyConfig{TypeMeta: metav1.TypeMeta{Kind: "EnvoyConfig"}}},
			r.EnvoyConfigHandler()).
		Complete(r)
}

// EnvoyConfigHandler returns an EventHandler to watch for EnvoyConfigs
func (r *EnvoyDeploymentReconciler) EnvoyConfigHandler() handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(
		func(o client.Object) []reconcile.Request {
			edList := &operatorv1alpha1.EnvoyDeploymentList{}
			if err := r.GetClient().List(context.TODO(), edList, client.InNamespace(o.GetNamespace())); err != nil {
				r.Log.Error(err, "unable to retrieve the list of EnvoyDeployment resources in the namespace",
					"Type", "EnvoyConfig", "Name", o.GetName(), "Namespace", o.GetNamespace())
				return []reconcile.Request{}
			}

			// Return a reconcile event for all EnvoyDeployments that have a reference to this EnvoyConfig
			req := []reconcile.Request{}
			for _, ed := range edList.Items {
				if ed.Spec.EnvoyConfigRef == o.GetName() {
					req = append(req, reconcile.Request{
						NamespacedName: types.NamespacedName{Name: ed.GetName(), Namespace: ed.GetNamespace()},
					})
				}
			}

			return req
		},
	)
}
