package spec

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// TPRKindMetadata metadata
	TPRKindMetadata = "metadata"
	// TPRKindTidb tidb
	TPRKindTidb = "tidb"
	// TPRKindEvent event type
	TPRKindEvent = "event"
	// TPRGroup all resources group
	TPRGroup = "tidb.ffan.com"
	// TPRVersion current version is beta
	TPRVersion = "v1beta1"
	// TPRDescription a trp desc
	TPRDescription = "Manage tidb cluster"
)

// Resource tpr
type Resource struct {
	metav1.TypeMeta `json:",inline"`
	Metadata        metav1.ObjectMeta `json:"metadata,omitempty"`
}