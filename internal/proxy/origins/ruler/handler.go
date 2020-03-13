/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
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

package ruler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Comcast/trickster/internal/cache"
	"github.com/Comcast/trickster/internal/config"
	"github.com/Comcast/trickster/internal/proxy/methods"
)

// Handler processes the HTTP request through the rules engine
func (c *Client) Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("!!!!!!!!")

	// TODO: Connect the logic dots that actually determine nextRoute
	nextRoute := c.rule.nextRoute

	if oc, ok := c.originConfigs[c.rule.nextRoute]; ok {
		fmt.Println("???????")
		if strings.HasPrefix(r.URL.Path, c.pathPrefix) {
			fmt.Println("********")
			r.URL.Path = strings.Replace(r.URL.Path, c.pathPrefix, "/"+nextRoute+"/", 1)
		} else {
			fmt.Println("<<<<<<<<<")
			r.URL.Path = "/" + nextRoute + "/" + r.URL.Path
		}
		fmt.Println("======", r.URL.Path)

		oc.Router.ServeHTTP(w, r)
		return
	}

	fmt.Println(":(:(:(")

	c.options.Router.NotFoundHandler.ServeHTTP(w, r)

}

// Client Implements the Proxy Client Interface
type Client struct {
	name               string
	originConfigs      map[string]*config.OriginConfig
	options            *config.OriginConfig
	handlers           map[string]http.Handler
	handlersRegistered bool
	rule               *rule
	pathPrefix         string
}

// NewClient returns a new Rules Router client reference
func NewClient(name string, options *config.OriginConfig, ocm map[string]*config.OriginConfig) (*Client, error) {
	return &Client{
		name:          name,
		options:       options,
		originConfigs: ocm,
		pathPrefix:    "/" + name + "/",
		rule:          &rule{nextRoute: "sim2"},
	}, nil
}

// Configuration returns the Client Configuration
func (c *Client) Configuration() *config.OriginConfig {
	return c.options
}

// DefaultPathConfigs returns the default PathConfigs for the given OriginType
func (c *Client) DefaultPathConfigs(oc *config.OriginConfig) map[string]*config.PathConfig {

	m := methods.CacheableHTTPMethods()
	m = append(m, methods.CacheableHTTPMethods()...)

	paths := map[string]*config.PathConfig{
		"/" + strings.Join(m, "-"): {
			Path:          "/",
			HandlerName:   "ruler",
			Methods:       m,
			OriginConfig:  oc,
			MatchType:     config.PathMatchTypePrefix,
			MatchTypeName: "prefix",
		},
	}
	return paths
}

func (c *Client) registerHandlers() {
	c.handlersRegistered = true
	c.handlers = make(map[string]http.Handler)
	// This is the registry of handlers that Trickster supports for the Reverse Proxy Cache,
	// and are able to be referenced by name (map key) in Config Files
	c.handlers["ruler"] = http.HandlerFunc(c.Handler)
}

// Handlers returns a map of the HTTP Handlers the client has registered
func (c *Client) Handlers() map[string]http.Handler {
	if !c.handlersRegistered {
		c.registerHandlers()
	}
	return c.handlers
}

// HTTPClient is not used by the Ruler, and is present to conform to the Client interface
func (c *Client) HTTPClient() *http.Client {
	return nil
}

// Cache is not used by the Ruler, and is present to conform to the Client interface
func (c *Client) Cache() cache.Cache {
	return nil
}

// Name returns the name of the upstream Configuration proxied by the Client
func (c *Client) Name() string {
	return c.name
}

// SetCache is not used by the Ruler, and is present to conform to the Client interface
func (c *Client) SetCache(cc cache.Cache) {}