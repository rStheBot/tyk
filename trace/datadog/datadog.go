package datadog

import (
	opentracing "github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Name is the name of this tracer.
const Name = "datadog"

type Trace struct {
	opentracing.Tracer
}

func (Trace) Name() string {
	return Name
}

type Logger interface {
	Log(msg string)
	Info(msg string)
}

type wrapLogger struct {
	Logger
}

func (w wrapLogger) Log(msg string) {
	w.Info(msg)
}

// Init returns a implementation of tyk.Tracer using datadog client.
func Init(service string, opts map[string]interface{}, log Logger) (*Trace, error) {
	c, err := Load(opts)
	if err != nil {
		return nil, err
	}

	var ddopts []tracer.StartOption

	ddopts = append(ddopts, tracer.WithLogger(&wrapLogger{Logger: log}))

	if service != "" {
		ddopts = append(ddopts, tracer.WithServiceName(service))
	}

	if c.AgentAddress != "" {
		ddopts = append(ddopts, tracer.WithAgentAddr(c.AgentAddress))
	}

	if c.DogstatsdAddress != "" {
		ddopts = append(ddopts, tracer.WithDogstatsdAddress(c.DogstatsdAddress))
	}

	if c.Analytics.Enabled {
		ddopts = append(ddopts, tracer.WithAnalytics(true))
		ddopts = append(ddopts, tracer.WithAnalyticsRate(c.Analytics.Rate))
	}

	if c.Sampler != nil {
		ddopts = append(ddopts, tracer.WithSampler(tracer.NewRateSampler(c.Sampler.Rate)))
	}

	if c.Propagator != nil {
		ddopts = append(ddopts, tracer.WithPropagator(tracer.NewPropagator(&tracer.PropagatorConfig{
			TraceHeader:    c.Propagator.TraceHeader,
			ParentHeader:   c.Propagator.ParentHeader,
			PriorityHeader: c.Propagator.PriorityHeader,
			BaggagePrefix:  c.Propagator.BaggagePrefix,
		})))
	}

	tr := opentracer.New(ddopts...)

	opentracing.SetGlobalTracer(tr)

	return &Trace{Tracer: tr}, nil
}
