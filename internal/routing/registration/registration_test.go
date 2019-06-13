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
	"testing"

	"github.com/Comcast/trickster/internal/cache/registration"
	"github.com/Comcast/trickster/internal/config"
	"github.com/Comcast/trickster/internal/util/metrics"
)

func init() {
	metrics.Init()
}

func TestRegisterProxyRoutes(t *testing.T) {

	err := config.Load("trickster", "test", []string{"-log-level", "debug", "-origin-url", "http://1", "-origin-type", "prometheus"})
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}
	registration.LoadCachesFromConfig()
	RegisterProxyRoutes()

	if len(ProxyClients) == 0 {
		t.Errorf("expected %d got %d", 1, 0)
	}
}

func TestRegisterProxyRoutesInflux(t *testing.T) {

	err := config.Load("trickster", "test", []string{"-log-level", "debug", "-origin-url", "http://1", "-origin-type", "influxdb"})
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}

	registration.LoadCachesFromConfig()
	RegisterProxyRoutes()

	if len(ProxyClients) == 0 {
		t.Errorf("expected %d got %d", 1, 0)
	}
}

func TestRegisterProxyRoutesMultipleDefaults(t *testing.T) {
	expected := "too many default origins"
	a := []string{"-config", "../../../testdata/test.too_many_defaults.conf"}
	err := config.Load("trickster", "test", a)
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}
	registration.LoadCachesFromConfig()
	err = RegisterProxyRoutes()
	if err == nil {
		t.Errorf("expected error `%s` got nothing", expected)
	} else if err.Error() != expected {
		t.Errorf("expected error `%s` got `%s`", expected, err.Error())
	}
}

func TestRegisterProxyRoutesBadCacheName(t *testing.T) {
	expected := "invalid cache name in origin config. originName: test, cacheName: test2"
	a := []string{"-config", "../../../testdata/test.bad_cache_name.conf"}
	err := config.Load("trickster", "test", a)
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}
	registration.LoadCachesFromConfig()
	err = RegisterProxyRoutes()
	if err == nil {
		t.Errorf("expected error `%s` got nothing", expected)
	} else if err.Error() != expected {
		t.Errorf("expected error `%s` got `%s`", expected, err.Error())
	}
}

func TestRegisterProxyRoutesBadOriginType(t *testing.T) {
	expected := "unknown origin type in origin config. originName: test, originType: foo"
	a := []string{"-config", "../../../testdata/test.unknown_origin_type.conf"}
	err := config.Load("trickster", "test", a)
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}
	registration.LoadCachesFromConfig()
	err = RegisterProxyRoutes()
	if err == nil {
		t.Errorf("expected error `%s` got nothing", expected)
	} else if err.Error() != expected {
		t.Errorf("expected error `%s` got `%s`", expected, err.Error())
	}
}

func TestRegisterMultipleOrigins(t *testing.T) {
	a := []string{"-config", "../../../testdata/test.multiple_origins.conf"}
	err := config.Load("trickster", "test", a)
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}
	registration.LoadCachesFromConfig()
	err = RegisterProxyRoutes()
	if err != nil {
		t.Error(err)
	}
}

func TestRegisterMultipleOriginsPlusDefault(t *testing.T) {
	a := []string{"-config", "../../../testdata/test.multiple_origins_plus_default.conf"}
	err := config.Load("trickster", "test", a)
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}
	registration.LoadCachesFromConfig()
	err = RegisterProxyRoutes()
	if err != nil {
		t.Error(err)
	}
	if !config.Origins["default"].IsDefault {
		t.Errorf("expected origin %s.IsDefault to be true", "default")
	}
}
