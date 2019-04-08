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

package proxy

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Comcast/trickster/internal/config"
	"github.com/Comcast/trickster/internal/util/metrics"
	tu "github.com/Comcast/trickster/internal/util/testing"
)

func init() {
	metrics.Init()
}

func TestProxyRequest(t *testing.T) {

	es := tu.NewTestServer(200, "test")
	defer es.Close()

	err := config.Load("trickster", "test", []string{"-origin", es.URL, "-origin-type", "test", "-log-level", "debug"})
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", es.URL, nil)

	// get URL

	req := NewRequest("default", "test", "TestProxyRequest", "GET", r.URL, http.Header{"testHeaderName": []string{"testHeaderValue"}}, time.Duration(30)*time.Second, r)
	ProxyRequest(req, w)
	resp := w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d.", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != "test" {
		t.Errorf("expected 'test' got '%s'.", bodyBytes)
	}
}

func TestProxyRequestBadGateway(t *testing.T) {

	const badUpstream = "http://127.0.0.1:64389"

	// assume nothing listens on badUpstream, so this should force the proxy to generate a 502 Bad Gateway
	err := config.Load("trickster", "test", []string{"-origin", badUpstream, "-origin-type", "test", "-log-level", "debug"})
	if err != nil {
		t.Errorf("Could not load configuration: %s", err.Error())
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", badUpstream, nil)

	// get URL

	req := NewRequest("default", "test", "TestProxyRequest", "GET", r.URL, make(http.Header), time.Duration(30)*time.Second, r)
	ProxyRequest(req, w)
	resp := w.Result()

	// it should return 502 Bad Gateway
	if resp.StatusCode != 502 {
		t.Errorf("expected 502 got %d.", resp.StatusCode)
	}

}
