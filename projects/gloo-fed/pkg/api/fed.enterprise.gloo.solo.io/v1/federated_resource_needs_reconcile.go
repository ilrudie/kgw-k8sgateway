// Code generated by skv2. DO NOT EDIT.

package v1

import (
	mc_types "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/core/v1"
)

// NeedsReconcile returns true if the object has not been observed or is in some state where a retry is needed
// This implementation is not ideal, as it relies on using the Status of the resource, which should
// be a read-only field. Instead, we could compare the resource hash to the last observed hash
func (obj *FederatedAuthConfig) NeedsReconcile(placementStatus *mc_types.PlacementStatus) bool {
	// If the FederatedAuthConfig has not been observed or is in some state where a retry is needed, it needs reconcile
	return obj.Generation != placementStatus.GetObservedGeneration() ||
		placementStatus.GetState() != mc_types.PlacementStatus_PLACED ||
		placementStatus.GetState() != mc_types.PlacementStatus_FAILED
}