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

package prometheus

import (
	ho "github.com/trickstercache/trickster/pkg/backends/healthcheck/options"
)

// DefaultHealthCheckConfig returns the default HealthCheck Config for this backend provider
func (c *Client) DefaultHealthCheckConfig() *ho.Options {
	o := ho.New()
	u := c.BaseUpstreamURL()
	o.Scheme = u.Scheme
	o.Host = u.Host
	o.Path = u.Path + "/api/v1/query"
	o.Query = "query=up"
	return o
}
