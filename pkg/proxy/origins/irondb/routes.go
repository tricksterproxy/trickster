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

package irondb

import (
	"net/http"

	"github.com/tricksterproxy/trickster/pkg/cache/key"
	oo "github.com/tricksterproxy/trickster/pkg/proxy/origins/options"
	"github.com/tricksterproxy/trickster/pkg/proxy/paths/matching"
	po "github.com/tricksterproxy/trickster/pkg/proxy/paths/options"
)

func (c *Client) registerHandlers() {
	c.handlersRegistered = true
	c.handlers = make(map[string]http.Handler)
	// This is the registry of handlers that Trickster supports for IRONdb,
	// and are able to be referenced by name (map key) in Config Files
	c.handlers["health"] = http.HandlerFunc(c.HealthHandler)
	c.handlers[mnRaw] = http.HandlerFunc(c.RawHandler)
	c.handlers[mnRollup] = http.HandlerFunc(c.RollupHandler)
	c.handlers[mnFetch] = http.HandlerFunc(c.FetchHandler)
	c.handlers[mnRead] = http.HandlerFunc(c.TextHandler)
	c.handlers[mnHistogram] = http.HandlerFunc(c.HistogramHandler)
	c.handlers[mnFind] = http.HandlerFunc(c.FindHandler)
	c.handlers[mnState] = http.HandlerFunc(c.StateHandler)
	c.handlers[mnCAQL] = http.HandlerFunc(c.CAQLHandler)
	c.handlers["proxy"] = http.HandlerFunc(c.ProxyHandler)
}

// Handlers returns a map of the HTTP Handlers the client has registered
func (c *Client) Handlers() map[string]http.Handler {
	if !c.handlersRegistered {
		c.registerHandlers()
	}
	return c.handlers
}

func populateHeathCheckRequestValues(oc *oo.Options) {
	if oc.HealthCheckUpstreamPath == "-" {
		oc.HealthCheckUpstreamPath = "/" + mnState
	}
	if oc.HealthCheckVerb == "-" {
		oc.HealthCheckVerb = http.MethodGet
	}
	if oc.HealthCheckQuery == "-" {
		oc.HealthCheckQuery = ""
	}
}

// DefaultPathConfigs returns the default PathConfigs for the given OriginType
func (c *Client) DefaultPathConfigs(oc *oo.Options) map[string]*po.Options {

	populateHeathCheckRequestValues(oc)

	paths := map[string]*po.Options{

		"/" + mnRaw + "/": {
			Path:            "/" + mnRaw + "/",
			HandlerName:     "RawHandler",
			Methods:         []string{http.MethodGet},
			CacheKeyParams:  []string{},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnRollup + "/": {
			Path:            "/" + mnRollup + "/",
			HandlerName:     "RollupHandler",
			Methods:         []string{http.MethodGet},
			CacheKeyParams:  []string{upSpan, upEngine, upType},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnFetch: {
			Path:            "/" + mnFetch,
			HandlerName:     "FetchHandler",
			KeyHasher:       []key.HasherFunc{c.fetchHandlerDeriveCacheKey},
			Methods:         []string{http.MethodPost},
			CacheKeyParams:  []string{},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnRead + "/": {
			Path:            "/" + mnRead + "/",
			HandlerName:     "TextHandler",
			KeyHasher:       []key.HasherFunc{c.textHandlerDeriveCacheKey},
			Methods:         []string{http.MethodGet},
			CacheKeyParams:  []string{"*"},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnHistogram + "/": {
			Path:            "/" + mnHistogram + "/",
			HandlerName:     "HistogramHandler",
			Methods:         []string{http.MethodGet},
			KeyHasher:       []key.HasherFunc{c.histogramHandlerDeriveCacheKey},
			CacheKeyParams:  []string{},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnFind + "/": {
			Path:            "/" + mnFind + "/",
			HandlerName:     "FindHandler",
			Methods:         []string{http.MethodGet},
			CacheKeyParams:  []string{upQuery},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnState + "/": {
			Path:            "/" + mnState + "/",
			HandlerName:     "StateHandler",
			Methods:         []string{http.MethodGet},
			CacheKeyParams:  []string{"*"},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnCAQL: {
			Path:            "/" + mnCAQL,
			HandlerName:     "CAQLHandler",
			Methods:         []string{http.MethodGet},
			CacheKeyParams:  []string{upQuery, upCAQLQuery, upCAQLPeriod},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/" + mnCAQLPub + "/": {
			Path:            "/" + mnCAQLPub + "/",
			HandlerName:     "CAQLPubHandler",
			Methods:         []string{http.MethodGet},
			CacheKeyParams:  []string{upQuery, upCAQLQuery, upCAQLPeriod},
			CacheKeyHeaders: []string{},
			MatchType:       matching.PathMatchTypePrefix,
			MatchTypeName:   "prefix",
		},

		"/": {
			Path:          "/",
			HandlerName:   "ProxyHandler",
			Methods:       []string{http.MethodGet},
			MatchType:     matching.PathMatchTypePrefix,
			MatchTypeName: "prefix",
		},
	}

	return paths

}
