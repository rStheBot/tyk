package datadog

import (
	"encoding/json"
)

type Config struct {
	ServiceName      string      `json:"service_name"`
	AgentAddress     string      `json:"agent_address"`
	DogstatsdAddress string      `json:"dogstatsd_address"`
	Analytics        *Analytics  `json:"analytics,omitempty"`
	Sampler          *Sampler    `json:"sampler,omitempty"`
	Propagator       *Propagator `json:"propagator,omitempty"`
	DebugMode        bool        `json:"debug_mode"`
	RuntimeMetrics   bool        `json:"runtime_metrics_enabled"`
}

type Analytics struct {
	Enabled bool    `json:"enabled"`
	Rate    float64 `json:"rate"`
}

type Sampler struct {
	Rate float64 `json:"rate"`
}

type Propagator struct {
	BaggagePrefix  string `json:"baggage_prefix"`
	TraceHeader    string `json:"trace_header"`
	ParentHeader   string `json:"parent_header"`
	PriorityHeader string `json:"priority_header"`
}

func Load(opts map[string]interface{}) (*Config, error) {
	b, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(b, &c)
	return &c, nil
}
