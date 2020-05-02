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

package engines

import (
	"net/http"
	"strings"
	"testing"

	txe "github.com/tricksterproxy/trickster/pkg/proxy/errors"
	"github.com/tricksterproxy/trickster/pkg/proxy/headers"
	"github.com/tricksterproxy/trickster/pkg/proxy/ranges/byterange"
)

func TestDocumentFromHTTPResponse(t *testing.T) {

	expected := []byte("1234")

	resp := &http.Response{}
	resp.Header = http.Header{headers.NameContentRange: []string{"bytes 1-4/8"}}
	resp.StatusCode = 206
	d := DocumentFromHTTPResponse(resp, []byte("1234"), nil, testLogger)

	if len(d.Ranges) != 1 {
		t.Errorf("expected 1 got %d", len(d.Ranges))
	} else if string(d.RangeParts[d.Ranges[0]].Content) != string(expected) {
		t.Errorf("expected %s got %s", string(expected), string(d.Body))
	}

	if d.StatusCode != 206 {
		t.Errorf("expected %d got %d", 206, d.StatusCode)
	}

}

func TestCachingPolicyString(t *testing.T) {

	cp := &CachingPolicy{NoTransform: true}
	s := cp.String()

	i := strings.Index(s, `"no_transform":true`)
	if i < 1 {
		t.Errorf("expected value > 1, got %d", i)
	}

}

func TestSetBody(t *testing.T) {

	r := byterange.Range{Start: 0, End: 10}
	d := &HTTPDocument{ContentLength: -1,
		RangeParts: byterange.MultipartByteRanges{r: &byterange.MultipartByteRange{Range: r,
			Content: []byte("01234567890")}}}
	d.SetBody([]byte("testing"))

	if d.ContentLength < 0 {
		t.Errorf("expected value > 0, got %d", d.ContentLength)
	}
}

func TestSize(t *testing.T) {
	r := byterange.Range{Start: 0, End: 10}
	d := &HTTPDocument{ContentLength: -1,
		RangeParts: byterange.MultipartByteRanges{r: &byterange.MultipartByteRange{Range: r,
			Content: []byte("01234567890")}}}

	i := d.Size()

	if i != 62 {
		t.Errorf("expected %d got %d", 62, i)
	}

}

func TestFulfillContentBody(t *testing.T) {
	d := &HTTPDocument{}
	err := d.FulfillContentBody()
	if err != txe.ErrNoRanges {
		if err != nil {
			t.Error(err)
		} else {
			t.Errorf("expected error: %s", txe.ErrNoRanges.Error())
		}
	}
}

func TestParsePartialContentBodyNoRanges(t *testing.T) {

	d := &HTTPDocument{}
	resp := &http.Response{Header: make(http.Header)}
	d.ParsePartialContentBody(resp, []byte("test"), testLogger)

	if string(d.Body) != "test" {
		t.Errorf("expected %s got %s", "test", string(d.Body))
	}

}

func TestParsePartialContentBodySingleRange(t *testing.T) {
	d := &HTTPDocument{}
	d.Ranges = make(byterange.Ranges, 0)
	d.RangeParts = make(byterange.MultipartByteRanges)
	d.StoredRangeParts = make(map[string]*byterange.MultipartByteRange)

	resp := &http.Response{Header: http.Header{
		headers.NameContentRange: []string{"bytes 0-10/1222"},
	}}

	d.ParsePartialContentBody(resp, []byte("Lorem ipsum"), testLogger)

	if string(d.Body) != "" {
		t.Errorf("expected %s got %s", "", string(d.Body))
	}

	if len(d.RangeParts) != 1 {
		t.Errorf("expected %d got %d", 1, len(d.RangeParts))
	}
}

func TestParsePartialContentBodyMultipart(t *testing.T) {
	d := &HTTPDocument{}
	d.Ranges = make(byterange.Ranges, 0)
	d.RangeParts = make(byterange.MultipartByteRanges)
	d.StoredRangeParts = make(map[string]*byterange.MultipartByteRange)

	resp := &http.Response{
		StatusCode: http.StatusPartialContent,
		Header:     http.Header{},
	}

	resp.Header.Set(headers.NameContentType, "multipart/byteranges; boundary=c4fb8e6049a6fdb126d32fa0b15c21e3")
	resp.Header.Set(headers.NameContentLength, "271")

	d.ParsePartialContentBody(resp, []byte(`
--c4fb8e6049a6fdb126d32fa0b15c21e3
Content-Range: bytes 0-6/1222
Content-Type: text/plain; charset=utf-8

Lorem i
--c4fb8e6049a6fdb126d32fa0b15c21e3
Content-Range: bytes 10-20/1222
Content-Type: text/plain; charset=utf-8

m dolor sit
--c4fb8e6049a6fdb126d32fa0b15c21e3--`), testLogger)

	if string(d.Body) != "" {
		t.Errorf("expected %s got %s", "", string(d.Body))
	}

	if len(d.RangeParts) != 2 {
		t.Errorf("expected %d got %d", 2, len(d.RangeParts))
	}
}

func TestParsePartialContentBodyMultipartBadBody(t *testing.T) {
	d := &HTTPDocument{}
	d.Ranges = make(byterange.Ranges, 0)
	d.RangeParts = make(byterange.MultipartByteRanges)
	d.StoredRangeParts = make(map[string]*byterange.MultipartByteRange)

	resp := &http.Response{
		StatusCode: http.StatusPartialContent,
		Header:     http.Header{},
	}

	resp.Header.Set(headers.NameContentType, "multipart/byteranges; boundary=c4fb8e6049a6fdb126d32fa0b15c21e3")
	resp.Header.Set(headers.NameContentLength, "271")

	d.ParsePartialContentBody(resp, []byte(`
--c4fb8e6049a6fdb126d32fa0b15c21e3
Content-Range: bytes 0-6/1222
Content-Type: text/plain; charset=utf-8

Lorem i
--c4fb8e6049a6fdb126d32fa0b15c21e3
Content-Range: baytes 1s0-20/12s22x
Content-Type: text/plain; charset=utf-8

m dolor sit
--c4fb8e6049a6fdb126d32fa0b15c21e3--`), testLogger)

	if string(d.Body) != "" {
		t.Errorf("expected %s got %s", "", string(d.Body))
	}

	if len(d.RangeParts) != 0 {
		t.Errorf("expected %d got %d", 0, len(d.RangeParts))
	}

}

func TestLoadRangeParts(t *testing.T) {

	d := &HTTPDocument{
		rangePartsLoaded: true,
		StoredRangeParts: map[string]*byterange.MultipartByteRange{
			"range": {
				Range:   byterange.Range{Start: 0, End: 8},
				Content: []byte("trickster"),
			},
		},
	}

	// test the short circuit
	d.LoadRangeParts()
	if d.Ranges != nil {
		t.Errorf("expected nil got %s", d.Ranges.String())
	}

	// and now the main functionality
	d.rangePartsLoaded = false
	d.LoadRangeParts()
	if len(d.Ranges) != 1 {
		t.Errorf("expected %d got %d", 1, len(d.Ranges))
	}

}
