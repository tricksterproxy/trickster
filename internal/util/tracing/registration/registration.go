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
	"errors"
	"fmt"
	"sync"

	"go.opentelemetry.io/otel/api/trace"

	"github.com/Comcast/trickster/internal/config"
	"github.com/Comcast/trickster/internal/util/log"
	"github.com/Comcast/trickster/internal/util/tracing"
)

var initialize sync.Once

// Flushers represents a slice of Flusher functions for the configured Tracers
type Flushers []func()

// RegisterAll registers all Tracers in the provided configuration, and returns
// their Flushers
func RegisterAll(cfg *config.TricksterConfig) (Flushers, error) {

	if cfg == nil {
		return nil, errors.New("no config provided")
	}

	if cfg.Origins == nil {
		return nil, errors.New("no origins provided")
	}

	flushers := make(Flushers, 0, len(cfg.Origins))
	activeTracers := make(map[string]*config.TracingConfig)

	for _, oc := range cfg.Origins {
		if oc != nil {

			if oc.TracingConfigName == "" {
				continue
			}

			tc, ok := cfg.TracingConfigs[oc.TracingConfigName]
			if !ok {
				return nil, fmt.Errorf("invalid tracing config name [%s] for origin [%s]", oc.TracingConfigName, oc.Name)
			}

			if tc.Implementation == "" {
				continue
			}

			if _, ok := tracing.TracerImplementations[tc.Implementation]; !ok {
				return nil, fmt.Errorf("invalid tracing implementation [%s] for tracing config [%s]", tc.Implementation, oc.TracingConfigName)
			}

			if _, ok = activeTracers[oc.TracingConfigName]; !ok {
				tracer, flusher, err := Init(tc)
				if err != nil {
					return nil, err
				}

				flushers = append(flushers, flusher)
				activeTracers[oc.TracingConfigName] = tc
				oc.Tracer = tracer
			}

		}
	}

	return flushers, nil

}

// Init initializes tracing and returns a function to flush the tracer. Flush should be called on server shutdown.
func Init(cfg *config.TracingConfig) (trace.Tracer, func(), error) {

	var tracer trace.Tracer
	var fl, flusher func()
	var err error

	if cfg == nil {
		log.Info(
			"Nil Tracing Config, using noop tracer", nil,
		)
		return nil, func() {}, nil
	}
	log.Debug(
		"Trace Init",
		log.Pairs{
			"Implementation": cfg.Implementation,
			"Collector":      cfg.CollectorEndpoint,
			"Type":           tracing.TracerImplementations[cfg.Implementation],
		},
	)
	f := func() {
		tracer, fl, err = tracing.SetTracer(
			tracing.TracerImplementations[cfg.Implementation],
			cfg.CollectorEndpoint,
			cfg.SampleRate,
		)
		flusher = fl
	}
	initialize.Do(f)
	return tracer, flusher, err
}
