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

package tracing

import (
	"context"

	"github.com/Comcast/trickster/internal/util/log"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
)

const (
	// Trace implementation enum
	NoopTracer TracerImplementation = iota
	RecorderTracer
	StdoutTracer

	// TODO New Implementations go here

	JaegerTracer
)

type TracerImplementation int

var (
	tracerImplementationStrings = []string{
		"noop",
		"recorder",
		"stdout",
		"jaeger",
	}
	TracerImplementations = map[string]TracerImplementation{
		tracerImplementationStrings[NoopTracer]:     NoopTracer,
		tracerImplementationStrings[RecorderTracer]: RecorderTracer,
		tracerImplementationStrings[StdoutTracer]:   StdoutTracer,
		// TODO New Implementations go here
		tracerImplementationStrings[JaegerTracer]: JaegerTracer,
	}
)

func GlobalTracer(ctx context.Context) trace.Tracer {
	tracerName, ok := ctx.Value(tracerNameKey).(string)
	if !ok {
		return trace.NoopTracer{}

	}

	return global.TraceProvider().Tracer(tracerName)

}

func (t TracerImplementation) String() string {
	if t < NoopTracer || t > JaegerTracer {
		return "unknown-tracer"
	}
	return tracerImplementationStrings[t]
}

func SetTracer(t TracerImplementation, collectorURL string, sampleRate float64) (func(), error) {

	switch t {
	case StdoutTracer:

		return setStdOutTracer(sampleRate)
	case JaegerTracer:

		return setJaegerTracer(collectorURL, sampleRate)
	case RecorderTracer:
		// TODO make recorder available at runtime
		flush, _, err := setRecorderTracer(
			// Only called if there is an error so the log message won't be evaluated otherwise
			func(err error) {
				pairs := log.Pairs{
					"Error":                err,
					"TracerImplementation": tracerImplementationStrings[t],
					"Collector":            collectorURL,
					"SampleRate":           sampleRate,
				}
				log.Error(
					"Trace Recorder Error",
					pairs,
				)
			},
			sampleRate,
		)
		return flush, err
	default:

		return setNoopTracer()
	}

}
