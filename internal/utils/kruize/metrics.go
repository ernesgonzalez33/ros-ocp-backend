package kruize

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	kruizeAPIException = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "rosocp_kruize_api_exception_total",
		Help: "The total number of exception got while calling kruize API",
	},
		[]string{"path"},
	)
)
