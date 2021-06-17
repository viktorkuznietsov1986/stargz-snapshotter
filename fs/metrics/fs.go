/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// DockerOperationsKey is the key for docker operation metrics.
	OperationsLatencyKey = "request_duration"
	

	// Keep the "kubelet" subsystem for backward compatibility.
	stargz = "stargz"

)

var (
	// DockerOperationsLatency collects operation latency numbers by operation
	// type.
	OperationsLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: stargz,
			Name:      OperationsLatencyKey,
			Help:      "Latency in seconds of stargz snapshotter operations. Broken down by operation type.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"operation_type"},
	)
	
)

var registerMetrics sync.Once

// we can potentially utilize options
func Register() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(OperationsLatency)
	})
}