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
	"time"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// DockerOperationsKey is the key for docker operation metrics.
	OperationsLatencyKey = "operation_duration"
	

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

// We can potentially utilize options to granularly allow different metrics
func Register() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(OperationsLatency)
	})
}

// SinceInSeconds gets the time since the specified start in seconds.
func SinceInSeconds(start time.Time) float64 {
	return time.Since(start).Seconds()
}