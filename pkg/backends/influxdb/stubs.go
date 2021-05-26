/*
 * Copyright 2018 The Trickster Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package influxdb

import (
	"net/http"

	"github.com/trickstercache/trickster/pkg/timeseries"
)

// This file holds funcs required by the Proxy Client or Timeseries interfaces,
// but are (currently) unused by the InfluxDB implementation.

// Series (timeseries.Timeseries Interface) stub funcs

// FastForwardRequest is not used for InfluxDB and is here to conform to the Proxy Client interface
func (c *Client) FastForwardRequest(r *http.Request) (*http.Request, error) {
	return nil, nil
}

// InfluxDb Client (proxy.Client Interface) stub funcs

// UnmarshalInstantaneous is not used for InfluxDB and is here to conform to the Proxy Client interface
func (c *Client) UnmarshalInstantaneous(data []byte) (timeseries.Timeseries, error) {
	return nil, nil
}
