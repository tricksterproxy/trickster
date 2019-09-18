/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package reverseproxycache

import (
	"net/http"

	"github.com/Comcast/trickster/internal/config"
)

var handlers = make(map[string]http.Handler)
var handlersRegistered = false

func (c *Client) registerHandlers() {
	handlersRegistered = true
	// This is the registry of handlers that Trickster supports for Prometheus,
	// and are able to be referenced by name (map key) in Config Files
	handlers["health"] = http.HandlerFunc(c.HealthHandler)
	handlers["proxy"] = http.HandlerFunc(c.ProxyHandler)
	handlers["proxycache"] = http.HandlerFunc(c.ProxyCacheHandler)
	handlers["localresponse"] = http.HandlerFunc(c.LocalResponseHandler)
}

// Handlers returns a map of the HTTP Handlers the client has registered
func (c *Client) Handlers() map[string]http.Handler {
	if !handlersRegistered {
		c.registerHandlers()
	}
	return handlers
}

// DefaultPathConfigs returns the default PathConfigs for the given OriginType
func (c *Client) DefaultPathConfigs() (map[string]*config.ProxyPathConfig, []string) {
	paths := map[string]*config.ProxyPathConfig{
		"/": &config.ProxyPathConfig{
			Path:        "/",
			HandlerName: "proxy",
			Methods:     []string{http.MethodGet, http.MethodPost},
		},
	}
	orderedPaths := []string{"/"}
	return paths, orderedPaths
}
