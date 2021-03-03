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

package options

import (
	"github.com/tricksterproxy/trickster/cmd/trickster/config/defaults"
	jaegeropts "github.com/tricksterproxy/trickster/pkg/observability/tracing/exporters/jaeger/options"
	stdoutopts "github.com/tricksterproxy/trickster/pkg/observability/tracing/exporters/stdout/options"

	"github.com/tricksterproxy/trickster/pkg/util/strings"
	"github.com/tricksterproxy/trickster/pkg/util/yamlx"
)

// Options is a Tracing Options collection
type Options struct {
	Name          string            `yaml:"-"`
	Provider      string            `yaml:"provider,omitempty"`
	ServiceName   string            `yaml:"service_name,omitempty"`
	CollectorURL  string            `yaml:"collector_url,omitempty"`
	CollectorUser string            `yaml:"collector_user,omitempty"`
	CollectorPass string            `yaml:"collector_pass,omitempty"`
	SampleRate    float64           `yaml:"sample_rate,omitempty"`
	Tags          map[string]string `yaml:"tags,omitempty"`
	OmitTagsList  []string          `yaml:"omit_tags,omitempty"`

	StdOutOptions *stdoutopts.Options `yaml:"stdout,omitempty"`
	JaegerOptions *jaegeropts.Options `yaml:"jaeger,omitempty"`

	OmitTags map[string]bool `yaml:"-"`
	// for tracers that don't support WithProcess (e.g., Zipkin)
	attachTagsToSpan bool
}

// New returns a new *Options with the default values
func New() *Options {
	return &Options{
		Provider:      defaults.DefaultTracerProvider,
		ServiceName:   defaults.DefaultTracerServiceName,
		StdOutOptions: &stdoutopts.Options{},
		JaegerOptions: &jaegeropts.Options{},
	}
}

// Clone returns an exact copy of a tracing config
func (o *Options) Clone() *Options {
	var so *stdoutopts.Options
	if o.StdOutOptions != nil {
		so = o.StdOutOptions.Clone()
	}
	var jo *jaegeropts.Options
	if o.JaegerOptions != nil {
		jo = o.JaegerOptions.Clone()
	}
	return &Options{
		Name:             o.Name,
		Provider:         o.Provider,
		ServiceName:      o.ServiceName,
		CollectorURL:     o.CollectorURL,
		CollectorUser:    o.CollectorUser,
		CollectorPass:    o.CollectorPass,
		SampleRate:       o.SampleRate,
		Tags:             strings.CloneMap(o.Tags),
		OmitTags:         strings.CloneBoolMap(o.OmitTags),
		OmitTagsList:     strings.CloneList(o.OmitTagsList),
		StdOutOptions:    so,
		JaegerOptions:    jo,
		attachTagsToSpan: o.attachTagsToSpan,
	}
}

// ProcessTracingOptions enriches the configuration data of the provided Tracing Options collection
func ProcessTracingOptions(mo map[string]*Options, metadata yamlx.KeyLookup) {
	if len(mo) == 0 {
		return
	}
	for k, v := range mo {
		if metadata != nil {
			if !metadata.IsDefined("tracing", k, "sample_rate") {
				v.SampleRate = 1
			}
			if !metadata.IsDefined("tracing", k, "service_name") {
				v.ServiceName = defaults.DefaultTracerServiceName
			}
			if !metadata.IsDefined("tracing", k, "provider") {
				v.Provider = defaults.DefaultTracerProvider
			}
		}
		v.generateOmitTags()
		v.setAttachTags()
	}
}

func (o *Options) generateOmitTags() {
	o.OmitTags = make(map[string]bool)
	if len(o.OmitTagsList) == 0 {
		return
	}
	for _, v := range o.OmitTagsList {
		o.OmitTags[v] = true
	}
}

// AttachTagsToSpan indicates that Tags should be attached to the span
func (o *Options) AttachTagsToSpan() bool {
	return o.attachTagsToSpan
}

func (o *Options) setAttachTags() {
	if o.Provider == "zipkin" && o.Tags != nil && len(o.Tags) > 0 {
		o.attachTagsToSpan = true
	}
}
