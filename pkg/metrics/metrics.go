package metrics

import (
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
)

// ReconcileAttempts is a counter that indicates the number of attempts to reconcile the operator configuration.
var ReconcileAttempts = metrics.NewCounter(&metrics.CounterOpts{
	Name:           "cluster_monitoring_operator_reconcile_attempts_total",
	Help:           "Number of attempts to reconcile the operator configuration.",
	StabilityLevel: metrics.ALPHA,
})

// ReconcileStatus is a gauge that indicates the latest reconciliation state.
var ReconcileStatus = metrics.NewGauge(&metrics.GaugeOpts{
	Name:           "cluster_monitoring_operator_last_reconciliation_successful",
	Help:           "Latest reconciliation state. Set to 1 if last reconciliation succeeded, else 0.",
	StabilityLevel: metrics.ALPHA,
})

func init() {
	// The API (metrics) server is instrumented to work with component-base.
	// Refer: https://github.com/kubernetes/kubernetes/blob/ec87834bae787ab6687921d65c3bcfde8a6e01b9/staging/src/k8s.io/apiserver/pkg/server/routes/metrics.go#L44.
	legacyregistry.MustRegister(ReconcileAttempts)
	legacyregistry.MustRegister(ReconcileStatus)
}
