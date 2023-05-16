package app

import (
	"github.com/exfury/grbchain/x/common/monitor"
)

var (
	// init monitor prometheus metrics
	orderMetrics  = monitor.DefaultOrderMetrics(monitor.DefaultPrometheusConfig())
	streamMetrics = monitor.DefaultStreamMetrics(monitor.DefaultPrometheusConfig())
)
