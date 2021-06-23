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

package durationmetrics

import (
	"os"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// OperationLatencyKey is the key for stargz operation metrics.
	OperationLatencyKey = "operation_duration"

	// Keep namespace as stargz and subsystem as fs.
	namespace = "stargz"
	subsystem = "fs"
)

// Lists all metric labels.
const (
	Mount = "mount"
	RemoteRegistryGet = "remote_registry_get"
	NodeReaddir = "node_readdir"
)

var (
	// Buckets for OperationLatency metric in milliseconds.
	latencyBuckets = []float64{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384} // in milliseconds

	// OperationLatency collects operation latency numbers by operation
	// type.
	OperationLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      OperationLatencyKey,
			Help:      "Latency in milliseconds of stargz snapshotter operations. Broken down by operation type.",
			Buckets:   latencyBuckets,
			ConstLabels: prometheus.Labels {"component": "file_system"},
		},
		[]string{"operation_type", "host"},
	)
	
)

var registerMetrics sync.Once

var hostname string

// sinceInMilliseconds gets the time since the specified start in microseconds.
func sinceInMilliseconds(start time.Time) float64 {
	return float64(time.Since(start).Nanoseconds()/1e6)
}

func getHostName() string {
	sync.Once.Do(func() {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = ""
		}
	})
	
	return hostname
}

// Register metrics. This is always called only once.
func Register() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(OperationLatency)
	})
}

func MeasureLatency(operation string, start time.Time) {
	OperationLatency.WithLabelValues(operation, getHostName()).Observe(sinceInMilliseconds(start))
}
