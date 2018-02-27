package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BreakerConfig represents a Breaker configuration
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BreakerConfig struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents the desired BreakerConfig specification
	Spec BreakerConfigSpec `json:"spec,omitempty"`

	// Status represents the current BreakerConfig status
	Status BreakerConfigStatus `json:"status,omitempty"`
}

// BreakerConfigList implements list of BreakerConfig.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BreakerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is the list of BreakerConfig
	Items []BreakerConfig `json:"items"`
}

// BreakerConfigSpec contains BreakerConfig specification
type BreakerConfigSpec struct {
	Breaker  BreakerStrategy `json:"breaker"`
	Retry    RetryStrategy   `json:"retry"`
	Selector labels.Selector `json:"selector,omitempty"`
}

// BreakerConfigStatus contains BreakerConfig status
type BreakerConfigStatus struct {
	CurrentStatus string `json:"status"`
}

// BreakerStrategy contains BreakerStrategy definition
type BreakerStrategy struct {
	EvaluationPeriod      time.Duration `json:"evaluationPeriod,omitempty"`
	MinPodsAvailableCount *uint         `json:"minPodsAvailableCount,omitempty"`
	MinPodsAvailableRatio *uint         `json:"minPodsAvailableRatio,omitempty"`

	DiscreteValueOutOfList *DiscreteValueOutOfList `json:"discreteValueOutOfList,omitempty"`
}

// DiscreteValueOutOfList detect anomaly when the a value is not in the list with a ratio that exceed the tolerance
// The promQL should return counter that are grouped by:
// 1-the key of the value to monitor
// 2-the podname
type DiscreteValueOutOfList struct {
	PrometheusService string   `json:"PrometheusService"`
	PromQL            string   `json:"promQL"`               // example: sum(delta(ms_rpc_count{job=\"kubernetes-pods\",run=\"foo\"}[10s])) by (code,kubernetes_pod_name)
	Key               string   `json:"key"`                  // Key for the metrics. For the previous example it will be "code"
	PodNameKey        string   `json:"podNamekey"`           // Key to access the podName
	GoodValues        []string `json:"goodValues,omitempty"` // Good Values ["200","201"]. If empty means that BadValues should be used to do exclusion instead of inclusion.
	BadValues         []string `json:"badValues,omitempty"`  // Bad Values ["500","404"].
	TolerancePercent  uint     `json:"tolerance"`            // % of Bad values tolerated until the pod is considered out of SLA
}

// RetryStrategy contains RetryStrategy definition
type RetryStrategy struct {
	Mode     RetryStrategyMode `json:"mode"`
	Period   time.Duration     `json:"period,omitempty"`
	MaxRetry time.Duration     `json:"maxRetry,omitempty"`
}

// RetryStrategyMode represent the breaker Strategy Mode
type RetryStrategyMode string

// RetryStrategyModeDisabled represent the default retry strategy
const (
	RetryStrategyModeDisabled      RetryStrategyMode = "disabled"
	RetryStrategyModePeriodic      RetryStrategyMode = "periodic"
	RetryStrategyModeRetryAndKill  RetryStrategyMode = "retryAndKill"
	RetryStrategyModeRetryAndPause RetryStrategyMode = "retryAndPause"
)