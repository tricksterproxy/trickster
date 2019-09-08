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
	"github.com/Comcast/trickster/internal/proxy/origins/reverseproxycache"
	"github.com/Comcast/trickster/internal/util/log"
)


// ProxyClients maintains a list of proxy clients configured for use by Trickster
var ProxyClients = make(map[string]model.Client)

// RegisterProxyRoutes iterates the Trickster Configuration and registers the routes for the configured origins
func RegisterProxyRoutes() error {

	defaultOrigin := ""
	var ndo *config.OriginConfig // points to the origin config named "default"
	var cdo *config.OriginConfig // points to the origin config with IsDefault set to true

	// This iteration will ensure default origins are handled properly
	for k, o := range config.Origins {

		// Ensure only one default origin exists
		if o.IsDefault {
			if cdo != nil {
				return fmt.Errorf("only one origin can be marked as default. Found both %s and %s", defaultOrigin, k)
			}
			log.Debug("default origin identified", log.Pairs{"name": k})
			defaultOrigin = k
			cdo = o
			continue
		}

		// handle origin named "default" last as it needs special handling based on a full pass over the range
		if k == "default" {
			ndo = o
			continue
		}

		err := registerOriginRoutes(k, o)
		if err != nil {
			return err
		}
	}

	if ndo != nil {
		if cdo == nil {
			ndo.IsDefault = true
			cdo = ndo
			defaultOrigin = "default"
		} else {
			err := registerOriginRoutes("default", ndo)
			if err != nil {
				return err
			}
		}
	}

	if cdo != nil {
		return registerOriginRoutes(defaultOrigin, cdo)
	}

	return nil
}


func registerOriginRoutes(k string, o *config.OriginConfig) error {

	var client model.Client
	var c cache.Cache
	var err error

	c, err = registration.GetCache(o.CacheName)
	if err != nil {
		return err
	}
	switch strings.ToLower(o.OriginType) {
	case "prometheus", "":
		log.Info("registering Prometheus route paths", log.Pairs{"originName": k, "upstreamHost": o.Host})
		client = prometheus.NewClient(k, o, c)
	case "influxdb":
		log.Info("registering Influxdb route paths", log.Pairs{"originName": k, "upstreamHost": o.Host})
		client = influxdb.NewClient(k, o, c)
	case "irondb":
		log.Info("registering IRONdb route paths", log.Pairs{"originName": k, "upstreamHost": o.Host})
		client = irondb.NewClient(k, o, c)
	case "rpc", "reverseproxycache":
		log.Info("Registering ReverseProxyCache Route Paths", log.Pairs{"originName": k, "upstreamHost": o.Host})
		client = reverseproxycache.NewClient(k, o, c)
	default:
		log.Error("unknown origin type", log.Pairs{"originName": k, "originType": o.OriginType})
		return fmt.Errorf("unknown origin type in origin config. originName: %s, originType: %s", k, o.OriginType)
	}
	if client != nil {
		ProxyClients[k] = client
		client.RegisterRoutes(k, o)
	}

	return nil
}
