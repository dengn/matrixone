// Copyright 2023 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	S3ConnectCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "cn",
			Subsystem: "fs",
			Name:      "s3_connect_total",
			Help:      "Total number of s3 connect count.",
		})

	S3DNSResolveCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "cn",
			Subsystem: "fs",
			Name:      "s3_dns_resolve_total",
			Help:      "Total number of s3 dns resolve count.",
		})
)