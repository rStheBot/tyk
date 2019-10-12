package datadog

import (
	"encoding/json"
	"reflect"
	"testing"
)

const sampleConfig = `{
  "service_name": "your_service_name",
  "agent_address": "host:port",
  "dogstatsd_address": "host:port",
  "debug_mode": false,
  "runtime_metrics_enabled": true,
  "analytics": {
    "enabled": false,
    "rate": 0.5
  },
  "sampler": {
    "rate": 0.5
  },
  "propagator": {
    "baggage_prefix": "prefix",
    "trace_header": "x-trace-header",
    "parent_header": "x-parent-header",
    "priority_header": "x-priority-header"
  }
}`

func TestLoad(t *testing.T) {
	cfg := Config{
		ServiceName:      "your_service_name",
		AgentAddress:     "host:port",
		DogstatsdAddress: "host:port",
		DebugMode:        false,
		RuntimeMetrics:   true,
		Analytics: &Analytics{
			Enabled: false,
			Rate:    0.5,
		},
		Sampler: &Sampler{Rate: 0.5},
		Propagator: &Propagator{
			BaggagePrefix:  "prefix",
			TraceHeader:    "x-trace-header",
			ParentHeader:   "x-parent-header",
			PriorityHeader: "x-priority-header",
		},
	}

	var o map[string]interface{}
	err := json.Unmarshal([]byte(sampleConfig), &o)
	if err != nil {
		t.Fatal(err)
	}
	loadedConfig, err := Load(o)
	if err != nil {
		t.Fatal(err)
	}
	a := []interface{}{
		cfg.ServiceName, cfg.AgentAddress,
		cfg.DogstatsdAddress, cfg.DebugMode, cfg.RuntimeMetrics,
		cfg.Analytics, cfg.Sampler, cfg.Propagator,
	}
	b := []interface{}{
		loadedConfig.ServiceName, loadedConfig.AgentAddress,
		loadedConfig.DogstatsdAddress, loadedConfig.DebugMode, loadedConfig.RuntimeMetrics,
		loadedConfig.Analytics, loadedConfig.Sampler, loadedConfig.Propagator,
	}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("expected %v\n got  %v\n", cfg, loadedConfig)
	}
}
