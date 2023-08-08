/*
Copyright 2023.

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

package controller

import (
	"context"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demov1 "project/api/v1"
)

// DemoResourceReconciler reconciles a DemoResource object
type DemoResourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.demo.kcd.io,resources=demoresources,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.demo.kcd.io,resources=demoresources/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.demo.kcd.io,resources=demoresources/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DemoResource object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *DemoResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("Enter Reconcile", "req", req)

	testres := &demov1.DemoResource{}
	r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, testres)
	l.Info("Enter Reconcile", "spec", testres.Spec, "status", testres.Status)

	if testres.Spec.Name != testres.Status.Name {
		testres.Status.Name = testres.Spec.Name
		r.Status().Update(ctx, testres)
	}

	r.reconcilePOD(ctx, testres, l)

	return ctrl.Result{}, nil
}
func (r *DemoResourceReconciler) reconcilePOD(ctx context.Context, testres *demov1.DemoResource, l logr.Logger) error {
	name, namespace := testres.Name, testres.Namespace

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []v1.Container{
				v1.Container{
					Name:  "nginx",
					Image: "nginx:latest",
				},
			},
		},
	}

	l.Info("Creating POD")
	return r.Create(ctx, pod)

}

// SetupWithManager sets up the controller with the Manager.
func (r *DemoResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.DemoResource{}).
		Complete(r)
}
