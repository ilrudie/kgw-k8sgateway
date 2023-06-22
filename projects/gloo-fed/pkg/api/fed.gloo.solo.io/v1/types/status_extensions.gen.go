// Code generated by skv2. DO NOT EDIT.

// Generated status setting functions

package types

import (
	v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/core/v1"
)

// SetPlacementStatus assigns the PlacementStatus for FederatedUpstreamStatus
func (x *FederatedUpstreamStatus) SetPlacementStatus(status *v1.PlacementStatus) {
	if x != nil {
		x.PlacementStatus = status
	}
}

// SetNamespacedPlacementStatuses assigns the PlacementStatuses for FederatedUpstreamStatus
func (x *FederatedUpstreamStatus) SetNamespacedPlacementStatuses(statuses map[string]*v1.PlacementStatus) {
	if x != nil {
		x.NamespacedPlacementStatuses = statuses
	}
}

// SetPlacementStatus assigns the PlacementStatus for FederatedUpstreamGroupStatus
func (x *FederatedUpstreamGroupStatus) SetPlacementStatus(status *v1.PlacementStatus) {
	if x != nil {
		x.PlacementStatus = status
	}
}

// SetNamespacedPlacementStatuses assigns the PlacementStatuses for FederatedUpstreamGroupStatus
func (x *FederatedUpstreamGroupStatus) SetNamespacedPlacementStatuses(statuses map[string]*v1.PlacementStatus) {
	if x != nil {
		x.NamespacedPlacementStatuses = statuses
	}
}

// SetPlacementStatus assigns the PlacementStatus for FederatedSettingsStatus
func (x *FederatedSettingsStatus) SetPlacementStatus(status *v1.PlacementStatus) {
	if x != nil {
		x.PlacementStatus = status
	}
}

// SetNamespacedPlacementStatuses assigns the PlacementStatuses for FederatedSettingsStatus
func (x *FederatedSettingsStatus) SetNamespacedPlacementStatuses(statuses map[string]*v1.PlacementStatus) {
	if x != nil {
		x.NamespacedPlacementStatuses = statuses
	}
}