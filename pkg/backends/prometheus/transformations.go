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
	"net/http"

	"github.com/tricksterproxy/trickster/pkg/backends/prometheus/model"
	"github.com/tricksterproxy/trickster/pkg/observability/logging"
	"github.com/tricksterproxy/trickster/pkg/proxy/headers"
	"github.com/tricksterproxy/trickster/pkg/proxy/response/merge"
	"github.com/tricksterproxy/trickster/pkg/timeseries"
	"github.com/tricksterproxy/trickster/pkg/timeseries/dataset"
)

func (c *Client) ProcessTransformations(ts timeseries.Timeseries) {
	if !c.hasTransformations {
		return
	}
	ds, ok := ts.(*dataset.DataSet)
	if !ok {
		return
	}
	ds.InjectTags(c.injectLabels)
}

func (c *Client) processVectorTransformations(w http.ResponseWriter, rg *merge.ResponseGate, r *http.Request) {
	var trq *timeseries.TimeRangeQuery
	if rg.Resources.TimeRangeQuery != nil {
		trq = rg.Resources.TimeRangeQuery
	}
	bytes := rg.Body()
	h := w.Header()
	headers.Merge(h, rg.Header())
	t2, err := model.UnmarshalTimeseries(bytes, trq)
	if err != nil {
		logging.Error(rg.Resources.Logger, "vector unmarshaling error",
			logging.Pairs{"provider": "prometheus", "detail": err.Error()})
		defaultWrite(rg.Header(), rg.Response.StatusCode, w, bytes)
		return
	}
	ds, ok := t2.(*dataset.DataSet)
	if !ok {
		logging.Error(rg.Resources.Logger, "vector unmarshaling error",
			logging.Pairs{"provider": "prometheus", "detail": err.Error()})
		w.Write(bytes)
		return
	}
	ds.InjectTags(c.injectLabels)
	model.MarshalTSOrVectorWriter(ds, rg.Resources.TSReqestOptions, rg.Response.StatusCode, w, true)
}

func defaultWrite(eh http.Header, statusCode int, w http.ResponseWriter, b []byte) {
	w.WriteHeader(statusCode)
	w.Write(b)
}
