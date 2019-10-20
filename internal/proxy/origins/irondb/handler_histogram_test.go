package irondb

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Comcast/trickster/internal/config"
	"github.com/Comcast/trickster/internal/proxy/errors"
	"github.com/Comcast/trickster/internal/proxy/model"
	"github.com/Comcast/trickster/internal/timeseries"
	tc "github.com/Comcast/trickster/internal/util/context"
	tu "github.com/Comcast/trickster/internal/util/testing"
)

func TestHistogramHandler(t *testing.T) {

	client := &Client{name: "test"}
	ts, w, r, hc, err := tu.NewTestInstance("", client.DefaultPathConfigs, 200, "{}", nil, "irondb", "/histogram/0/900/300/00112233-4455-6677-8899-aabbccddeeff/"+
		"metric", "debug")
	client.config = tc.OriginConfig(r.Context())
	client.webClient = hc
	defer ts.Close()
	if err != nil {
		t.Error(err)
	}

	client.HistogramHandler(w, r)
	resp := w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d.", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != "{}" {
		t.Errorf("expected '{}' got %s.", bodyBytes)
	}

	oc := tc.OriginConfig(r.Context())
	cc := tc.CacheClient(r.Context())
	pc := tc.PathConfig(r.Context())

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET",
		"http://0/irondb/histogram/0/900/300/"+
			"00112233-4455-6677-8899-aabbccddeeff/"+
			"metric", nil)
	r = r.WithContext(tc.WithConfigs(r.Context(), oc, cc, pc))

	client.HistogramHandler(w, r)
	resp = w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d.", resp.StatusCode)
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != "{}" {
		t.Errorf("expected '{}' got %s.", bodyBytes)
	}
}

func TestHistogramHandlerDeriveCacheKey(t *testing.T) {

	client := &Client{name: "test"}
	path := "/histogram/0/900/00112233-4455-6677-8899-aabbccddeeff/metric"
	_, _, r, _, err := tu.NewTestInstance("", client.DefaultPathConfigs, 200, "{}", nil, "irondb", path, "debug")
	if err != nil {
		t.Error(err)
	}

	expected := "11cc1b20a869f6ff0559b08b014c3ca6"
	result := client.histogramHandlerDeriveCacheKey(path, r.URL.Query(), r.Header, r.Body, "extra")
	if result != expected {
		t.Errorf("expected %s got %s", expected, result)
	}

	expected = "c70681051e3af3de12f37686b6a4224f"
	path = "/irondb/0/900/00112233-4455-6677-8899-aabbccddeeff/metric"
	result = client.histogramHandlerDeriveCacheKey(path, r.URL.Query(), r.Header, r.Body, "extra")
	if result != expected {
		t.Errorf("expected %s got %s", expected, result)
	}

}

func TestHistogramHandlerParseTimeRangeQuery(t *testing.T) {

	path := "/histogram/0/900/300/00112233-4455-6677-8899-aabbccddeeff/metric"
	r, err := http.NewRequest(http.MethodGet, "http://0"+path, nil)
	if err != nil {
		t.Error(err)
	}

	// provide bad URL with no TimeRange query params
	client := &Client{name: "test"}
	hc := tu.NewTestWebClient()
	cfg := config.NewOriginConfig()
	cfg.Paths, _ = client.DefaultPathConfigs(cfg)

	tr := model.NewRequest("HistogramHandler", r.Method, r.URL, r.Header, cfg.Timeout, r, hc)

	// case where everthing is good
	_, err = client.histogramHandlerParseTimeRangeQuery(tr)
	if err != nil {
		t.Error(err)
	}

	// case where the path is not long enough
	r.URL.Path = "/histogram/0/900/"
	expected := errors.NotTimeRangeQuery().Error()
	_, err = client.histogramHandlerParseTimeRangeQuery(tr)
	if err == nil || err.Error() != expected {
		t.Errorf("expected %s got %s", expected, err.Error())
	}

	// case where the start can't be parsed
	r.URL.Path = "/histogram/z/900/300/00112233-4455-6677-8899-aabbccddeeff/metric"
	expected = `unable to parse timestamp z: strconv.ParseInt: parsing "z": invalid syntax`
	_, err = client.histogramHandlerParseTimeRangeQuery(tr)
	if err == nil || err.Error() != expected {
		t.Errorf("expected %s got %s", expected, err.Error())
	}

	// case where the end can't be parsed
	r.URL.Path = "/histogram/0/z/300/00112233-4455-6677-8899-aabbccddeeff/metric"
	_, err = client.histogramHandlerParseTimeRangeQuery(tr)
	if err == nil || err.Error() != expected {
		t.Errorf("expected %s got %s", expected, err.Error())
	}

	// case where the period can't be parsed
	r.URL.Path = "/histogram/0/900/z/00112233-4455-6677-8899-aabbccddeeff/metric"
	expected = `unable to parse duration zs: time: invalid duration zs`
	_, err = client.histogramHandlerParseTimeRangeQuery(tr)
	if err == nil || err.Error() != expected {
		t.Errorf("expected %s got %s", expected, err.Error())
	}

}

func TestHistogramHandlerSetExtent(t *testing.T) {

	// provide bad URL with no TimeRange query params
	client := &Client{name: "test"}
	hc := tu.NewTestWebClient()
	cfg := config.NewOriginConfig()
	cfg.Paths, _ = client.DefaultPathConfigs(cfg)
	r, err := http.NewRequest(http.MethodGet, "http://0/", nil)
	if err != nil {
		t.Error(err)
	}
	tr := model.NewRequest("HistogramHandler", r.Method, r.URL, r.Header, cfg.Timeout, r, hc)

	now := time.Now()
	then := now.Add(-5 * time.Hour)

	client.histogramHandlerSetExtent(tr, &timeseries.Extent{Start: then, End: now})
	if r.URL.Path != "/" {
		t.Errorf("expected %s got %s", "/", r.URL.Path)
	}

	// although SetExtent does not return a value to test, these lines exercise all coverage areas
	r.URL.Path = "/histogram/900/900/300/00112233-4455-6677-8899-aabbccddeeff/metric"
	client.histogramHandlerSetExtent(tr, &timeseries.Extent{Start: now, End: now})

	r.URL.Path = "/histogram/900/900/300"
	tr.TimeRangeQuery = &timeseries.TimeRangeQuery{Step: 300 * time.Second}
	client.histogramHandlerSetExtent(tr, &timeseries.Extent{Start: then, End: now})

}

func TestHistogramHandlerFastForwardURLError(t *testing.T) {

	// provide bad URL with no TimeRange query params
	client := &Client{name: "test"}
	hc := tu.NewTestWebClient()
	cfg := config.NewOriginConfig()
	cfg.Paths, _ = client.DefaultPathConfigs(cfg)
	r, err := http.NewRequest(http.MethodGet,
		"http://0/", nil)
	if err != nil {
		t.Error(err)
	}
	tr := model.NewRequest("HistogramHandler", r.Method, r.URL, r.Header, cfg.Timeout, r, hc)

	r.URL.Path = "/histogram/x/900/300/00112233-4455-6677-8899-aabbccddeeff/metric"
	_, err = client.histogramHandlerFastForwardURL(tr)
	if err == nil {
		t.Errorf("expected error: %s", "invalid parameters")
	}

	r.URL.Path = "/a/900/900/300/00112233-4455-6677-8899-aabbccddeeff/metric"
	tr.TimeRangeQuery = &timeseries.TimeRangeQuery{Step: 300 * time.Second}
	_, err = client.histogramHandlerFastForwardURL(tr)
	if err == nil {
		t.Errorf("expected error: %s", "invalid parameters")
	}

}
