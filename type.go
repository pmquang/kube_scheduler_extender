package main

import (
	"k8s.io/apimachinery/pkg/api/resource"
	k8sapi "k8s.io/kubernetes/pkg/api"
)

type NodeResource struct {
	Name          string
	Schedule      bool
	CpuReqs       resource.Quantity
	CpuLimits     resource.Quantity
	MemoryReqs    resource.Quantity
	MemoryLimits  resource.Quantity
	Allocate      k8sapi.ResourceList
	Load          float64
	FailedMessage string
	Threshold     ThreshodInfo
}

type ThreshodInfo struct {
	Cpu    int64
	Memory int64
	Load   float64
	CpuIdle float64
}

type Resp_PrometheusMetrics struct {
	Status string `json:"status"`
	Data   struct {
		Result []struct {
			Metric struct {
				Instance string `json:"instance"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
