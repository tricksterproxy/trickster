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

package timeseries

import (
	"time"
)

// TimeRangeQuery represents a timeseries database query parsed from an inbound HTTP request
type TimeRangeQuery struct {
	// Statement is the timeseries database query (with tokenized timeranges where present) requested by the user
	Statement string
	// Extent provides the start and end times for the request from a timeseries database
	Extent Extent
	// Step indicates the amount of time in seconds between each datapoint in a TimeRangeQuery's resulting timeseries
	Step time.Duration
	// IsOffset is true if the query uses a relative offset modifier
	IsOffset bool
	// TimestampFieldName indicates the database field name for the timestamp field
	TimestampFieldName string
}

// Copy returns an exact copy of a TimeRangeQuery
func (trq *TimeRangeQuery) Copy() *TimeRangeQuery {
	t2 := &TimeRangeQuery{
		Statement:          trq.Statement,
		Step:               trq.Step,
		Extent:             Extent{Start: trq.Extent.Start, End: trq.Extent.End},
		IsOffset:           trq.IsOffset,
		TimestampFieldName: trq.TimestampFieldName,
	}
	return t2
}

// NormalizeExtent adjusts the Start and End of a TimeRangeQuery's Extent to align against normalized boundaries.
func (trq *TimeRangeQuery) NormalizeExtent() {
	if trq.Step.Seconds() > 0 {
		if !trq.IsOffset && trq.Extent.End.After(time.Now()) {
			trq.Extent.End = time.Now()
		}
		trq.Extent.Start = trq.Extent.Start.Truncate(trq.Step)
		trq.Extent.End = trq.Extent.End.Truncate(trq.Step)
	}
}

// CalculateDeltas provides a list of extents that are not in a cached timeseries, when provided a list of extents that are cached.
func (trq *TimeRangeQuery) CalculateDeltas(have ExtentList) ExtentList {
	if len(have) == 0 {
		return ExtentList{trq.Extent}
	}
	misCap := trq.Extent.End.Sub(trq.Extent.Start) / trq.Step
	if misCap < 0 {
		misCap = 0
	}
	misses := make([]time.Time, 0, misCap)
	for i := trq.Extent.Start; !trq.Extent.End.Before(i); i = i.Add(trq.Step) {
		found := false
		for j := range have {
			if j == 0 && i.Before(have[j].Start) {
				// our earliest datapoint in cache is after the first point the user wants
				break
			}
			if i.Equal(have[j].Start) || i.Equal(have[j].End) || (i.After(have[j].Start) && have[j].End.After(i)) {
				found = true
				break
			}
		}
		if !found {
			misses = append(misses, i)
		}
	}
	// Find the fill and gap ranges
	ins := ExtentList{}
	var inStart = time.Time{}
	l := len(misses)
	for i := range misses {
		if inStart.IsZero() {
			inStart = misses[i]
		}
		if i+1 == l || !misses[i+1].Equal(misses[i].Add(trq.Step)) {
			ins = append(ins, Extent{Start: inStart, End: misses[i]})
			inStart = time.Time{}
		}
	}
	return ins
}
