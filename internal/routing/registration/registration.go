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

package registration

import (
	"fmt"
	"strings"

	"github.com/Comcast/trickster/internal/cache"
	"github.com/Comcast/trickster/internal/cache/registration"
	"github.com/Comcast/trickster/internal/config"
	"github.com/Comcast/trickster/internal/proxy/model"
	"github.com/Comcast/trickster/internal/proxy/origins/influxdb"
	"github.com/Comcast/trickster/internal/proxy/origins/irondb"
	"github.com/Comcast/trickster/internal/proxy/origins/prometheus"
	"github.com/Comcast/trickster/internal/util/log"
	kitlog "github.com/go-kit/kit/log"
)

// ProxyClients maintains a list of proxy clients configured for use by Trickster
var ProxyClients = make(map[string]model.Client)

// RegisterProxyRoutes iterates the Trickster Configuration and registers the routes for the configured origins
func RegisterProxyRoutes(l kitlog.Logger) error {

	defaultOrigin := ""

	// Iterate our origins from the config and register their path handlers into the mux.
	for k, o := range config.Origins {

		// Ensure only one default origin exists
		if o.IsDefault {
			if defaultOrigin != "" {
				return fmt.Errorf("only one origin can be marked as default. Found both %s and %s", defaultOrigin, k)
			}
			defaultOrigin = k
		}

		var client model.Client
		var c cache.Cache
		var err error

		c, err = registration.GetCache(o.CacheName)
		if err != nil {
			return err
		}
		switch strings.ToLower(o.Type) {
		case "prometheus", "":
			log.Info(l, "Registering Prometheus Route Paths", log.Pairs{"originName": k, "upstreamHost": o.Host})
			client = prometheus.NewClient(k, o, c, l)
		case "influxdb":
			log.Info(l, "Registering Influxdb Route Paths", log.Pairs{"originName": k, "upstreamHost": o.Host})
			client = influxdb.NewClient(k, o, c, l)
		case "irondb":
			log.Info(l, "Registering IRONdb Route Paths", log.Pairs{"originName": k, "upstreamHost": o.Host})
			client = irondb.NewClient(k, o, c, l)
		}
		if client != nil {
			ProxyClients[k] = client

			// If it's the default origin, register it last
			if o.IsDefault {
				defer client.RegisterRoutes(k, o)
			} else {
				client.RegisterRoutes(k, o)
			}
		}
	}
	return nil
}
