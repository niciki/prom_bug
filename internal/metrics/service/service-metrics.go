package quizservicemetrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ServiceMetrics struct {
	httpNbReq *prometheus.CounterVec
}

func NewServiceMetrics(ns string) (qs *ServiceMetrics) {
	httpNbReq := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: ns,
			Subsystem: "http",
			Name:      "nnnnnn",
		}, []string{"ID"},
	)
	return &ServiceMetrics{
		httpNbReq: httpNbReq,
	}
}

func (qs *ServiceMetrics) Info(ctx context.Context, id string) (err error) {
	id1 := ""
	for _, v := range id {
		id1 += string(v)
	}
	qs.httpNbReq.WithLabelValues(id).Add(1)
	return nil
}

// curl --location 'http://localhost:8080/v1/info/79'
// curl --location 'http://localhost:8080/v1/info/711'
// curl --location 'http://localhost:8080/v1/info/7'
// curl --location 'http://localhost:8080/v1/info/171'

// curl --location 'http://localhost:9090/'

// An error has occurred while serving metrics:
// collected metric "ns_http_nnnnnn" { label:{name:"ID"  value:"171"}  counter:{value:1  created_timestamp:{seconds:1704806774  nanos:95948000}}} was collected before with the same name and label values
