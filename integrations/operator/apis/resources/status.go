package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Status defines the observed state of the Teleport resource
type Status struct {
	// Conditions represent the latest available observations of an object's state
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
	// +optional
	TeleportResourceID int64 `json:"teleportResourceID"`
}

// DeepCopyInto deep-copies one resource status into another.
// Required to satisfy runtime.Object interface.
func (status *Status) DeepCopyInto(out *Status) {
	*out = Status{}
	out.Conditions = make([]metav1.Condition, len(status.Conditions))
	for i, cond := range status.Conditions {
		out.Conditions[i] = cond
	}
	out.TeleportResourceID = status.TeleportResourceID
}
