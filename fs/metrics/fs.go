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

var once sync.Once

type FileSystemMetrics struct {
	MountOperationDuration prometheus.Summary
	FetchRoundtripDuration prometheus.Summary
}

var instance *FileSystemMetrics

func GetMetrics() *FileSystemMetrics {
	once.Do(func() {
		instance = &FileSystemMetrics {
			MountOperationDuration: prometheus.NewSummary(
				prometheus.SummaryOpts{
					Name:       "fs_mount_request_duration",
					Help:       "fs mount latency in seconds",
					Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
				}),
			FetchRoundtripDuration: prometheus.NewSummary(
				prometheus.SummaryOpts{
					Name:       "fetch_request_roundtrip_duration",
					Help:       "fetch request roundtrip latency in seconds",
					Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
				}),
		}

		isntance.register()
	})

	return instance
}

// we can potentially utilize options
func (m *FileSystemMetrics) register() {
	prometheus.MustRegister(m.MountOperationDuration)
	prometheus.MustRegister(m.FetchRoundtripDuration)
}