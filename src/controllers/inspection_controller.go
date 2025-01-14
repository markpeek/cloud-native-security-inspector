// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	goharborv1alpha1 "goharbor.io/k8s-security-inspector/api/v1alpha1"
)

// InspectionReconciler reconciles a Inspection object
type InspectionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=goharbor.goharbor.io,resources=inspections,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=goharbor.goharbor.io,resources=inspections/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=goharbor.goharbor.io,resources=inspections/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Inspection object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *InspectionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *InspectionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&goharborv1alpha1.Inspection{}).
		Complete(r)
}
