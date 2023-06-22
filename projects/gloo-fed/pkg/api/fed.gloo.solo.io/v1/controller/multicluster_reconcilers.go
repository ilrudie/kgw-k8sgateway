// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	fed_gloo_solo_io_v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gloo.solo.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the FederatedUpstream Resource across clusters.
// implemented by the user
type MulticlusterFederatedUpstreamReconciler interface {
	ReconcileFederatedUpstream(clusterName string, obj *fed_gloo_solo_io_v1.FederatedUpstream) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedUpstream Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedUpstreamDeletionReconciler interface {
	ReconcileFederatedUpstreamDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedUpstreamReconcilerFuncs struct {
	OnReconcileFederatedUpstream         func(clusterName string, obj *fed_gloo_solo_io_v1.FederatedUpstream) (reconcile.Result, error)
	OnReconcileFederatedUpstreamDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedUpstreamReconcilerFuncs) ReconcileFederatedUpstream(clusterName string, obj *fed_gloo_solo_io_v1.FederatedUpstream) (reconcile.Result, error) {
	if f.OnReconcileFederatedUpstream == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedUpstream(clusterName, obj)
}

func (f *MulticlusterFederatedUpstreamReconcilerFuncs) ReconcileFederatedUpstreamDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedUpstreamDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedUpstreamDeletion(clusterName, req)
}

type MulticlusterFederatedUpstreamReconcileLoop interface {
	// AddMulticlusterFederatedUpstreamReconciler adds a MulticlusterFederatedUpstreamReconciler to the MulticlusterFederatedUpstreamReconcileLoop.
	AddMulticlusterFederatedUpstreamReconciler(ctx context.Context, rec MulticlusterFederatedUpstreamReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedUpstreamReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedUpstreamReconcileLoop) AddMulticlusterFederatedUpstreamReconciler(ctx context.Context, rec MulticlusterFederatedUpstreamReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedUpstreamMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedUpstreamReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedUpstreamReconcileLoop {
	return &multiclusterFederatedUpstreamReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gloo_solo_io_v1.FederatedUpstream{}, options)}
}

type genericFederatedUpstreamMulticlusterReconciler struct {
	reconciler MulticlusterFederatedUpstreamReconciler
}

func (g genericFederatedUpstreamMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedUpstreamDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedUpstreamDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedUpstreamMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstream)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedUpstream handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedUpstream(cluster, obj)
}

// Reconcile Upsert events for the FederatedUpstreamGroup Resource across clusters.
// implemented by the user
type MulticlusterFederatedUpstreamGroupReconciler interface {
	ReconcileFederatedUpstreamGroup(clusterName string, obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedUpstreamGroup Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedUpstreamGroupDeletionReconciler interface {
	ReconcileFederatedUpstreamGroupDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedUpstreamGroupReconcilerFuncs struct {
	OnReconcileFederatedUpstreamGroup         func(clusterName string, obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) (reconcile.Result, error)
	OnReconcileFederatedUpstreamGroupDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedUpstreamGroupReconcilerFuncs) ReconcileFederatedUpstreamGroup(clusterName string, obj *fed_gloo_solo_io_v1.FederatedUpstreamGroup) (reconcile.Result, error) {
	if f.OnReconcileFederatedUpstreamGroup == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedUpstreamGroup(clusterName, obj)
}

func (f *MulticlusterFederatedUpstreamGroupReconcilerFuncs) ReconcileFederatedUpstreamGroupDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedUpstreamGroupDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedUpstreamGroupDeletion(clusterName, req)
}

type MulticlusterFederatedUpstreamGroupReconcileLoop interface {
	// AddMulticlusterFederatedUpstreamGroupReconciler adds a MulticlusterFederatedUpstreamGroupReconciler to the MulticlusterFederatedUpstreamGroupReconcileLoop.
	AddMulticlusterFederatedUpstreamGroupReconciler(ctx context.Context, rec MulticlusterFederatedUpstreamGroupReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedUpstreamGroupReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedUpstreamGroupReconcileLoop) AddMulticlusterFederatedUpstreamGroupReconciler(ctx context.Context, rec MulticlusterFederatedUpstreamGroupReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedUpstreamGroupMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedUpstreamGroupReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedUpstreamGroupReconcileLoop {
	return &multiclusterFederatedUpstreamGroupReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gloo_solo_io_v1.FederatedUpstreamGroup{}, options)}
}

type genericFederatedUpstreamGroupMulticlusterReconciler struct {
	reconciler MulticlusterFederatedUpstreamGroupReconciler
}

func (g genericFederatedUpstreamGroupMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedUpstreamGroupDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedUpstreamGroupDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedUpstreamGroupMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedUpstreamGroup)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedUpstreamGroup handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedUpstreamGroup(cluster, obj)
}

// Reconcile Upsert events for the FederatedSettings Resource across clusters.
// implemented by the user
type MulticlusterFederatedSettingsReconciler interface {
	ReconcileFederatedSettings(clusterName string, obj *fed_gloo_solo_io_v1.FederatedSettings) (reconcile.Result, error)
}

// Reconcile deletion events for the FederatedSettings Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterFederatedSettingsDeletionReconciler interface {
	ReconcileFederatedSettingsDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterFederatedSettingsReconcilerFuncs struct {
	OnReconcileFederatedSettings         func(clusterName string, obj *fed_gloo_solo_io_v1.FederatedSettings) (reconcile.Result, error)
	OnReconcileFederatedSettingsDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterFederatedSettingsReconcilerFuncs) ReconcileFederatedSettings(clusterName string, obj *fed_gloo_solo_io_v1.FederatedSettings) (reconcile.Result, error) {
	if f.OnReconcileFederatedSettings == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileFederatedSettings(clusterName, obj)
}

func (f *MulticlusterFederatedSettingsReconcilerFuncs) ReconcileFederatedSettingsDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileFederatedSettingsDeletion == nil {
		return nil
	}
	return f.OnReconcileFederatedSettingsDeletion(clusterName, req)
}

type MulticlusterFederatedSettingsReconcileLoop interface {
	// AddMulticlusterFederatedSettingsReconciler adds a MulticlusterFederatedSettingsReconciler to the MulticlusterFederatedSettingsReconcileLoop.
	AddMulticlusterFederatedSettingsReconciler(ctx context.Context, rec MulticlusterFederatedSettingsReconciler, predicates ...predicate.Predicate)
}

type multiclusterFederatedSettingsReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterFederatedSettingsReconcileLoop) AddMulticlusterFederatedSettingsReconciler(ctx context.Context, rec MulticlusterFederatedSettingsReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericFederatedSettingsMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterFederatedSettingsReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterFederatedSettingsReconcileLoop {
	return &multiclusterFederatedSettingsReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &fed_gloo_solo_io_v1.FederatedSettings{}, options)}
}

type genericFederatedSettingsMulticlusterReconciler struct {
	reconciler MulticlusterFederatedSettingsReconciler
}

func (g genericFederatedSettingsMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterFederatedSettingsDeletionReconciler); ok {
		return deletionReconciler.ReconcileFederatedSettingsDeletion(cluster, req)
	}
	return nil
}

func (g genericFederatedSettingsMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*fed_gloo_solo_io_v1.FederatedSettings)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: FederatedSettings handler received event for %T", object)
	}
	return g.reconciler.ReconcileFederatedSettings(cluster, obj)
}