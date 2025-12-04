package main

type Snap struct {
	Snap      string            `yaml:"snap"`
	Info      Info              `yaml:"info"`
	Namespace string            `yaml:"namespace"`
	Objects   map[string]Object `yaml:"objects"`
	Defs      map[string]Def    `yaml:"definitions"`
	MCP       *MCPConfig        `yaml:"mcp,omitempty"`
}

type Info struct {
	Version string `yaml:"version"`
	Title   string `yaml:"title"`
}

type MCPConfig struct {
	Enabled       bool                 `yaml:"enabled"`
	Transport     string               `yaml:"transport,omitempty"` // stdio, sse
	Description   string               `yaml:"description,omitempty"`
	Receiver      MCPReceiver          `yaml:"receiver,omitempty"`
	Observability *ObservabilityConfig `yaml:"observability,omitempty"`
}

type MCPReceiver struct {
	Type string `yaml:"type"` // http, command
	URL  string `yaml:"url,omitempty"`
	Cmd  string `yaml:"cmd,omitempty"`
}

type ObservabilityConfig struct {
	Enabled        bool   `yaml:"enabled"`
	ServiceName    string `yaml:"service_name,omitempty"`
	OTLPEndpoint   string `yaml:"otlp_endpoint,omitempty"`   // e.g., localhost:4317
	PrometheusPort int    `yaml:"prometheus_port,omitempty"` // e.g., 9090
}

type Object struct {
	Methods     map[string]Method `yaml:"methods"`
	Definitions map[string]Def    `yaml:"definitions"`
}

type Method struct {
	OperationId string            `yaml:"operationId"`
	Summary     string            `yaml:"summary"`
	Description string            `yaml:"description"`
	Parameters  []Parameter       `yaml:"parameters"`
	ReturnType  *ReturnType       `yaml:"returnType"` // Allow nil
	Receiver    map[string]string `yaml:"receiver"`   // Receiver added here
}

type Parameter struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Ref         string `yaml:"$ref"` // Add Ref for referencing definitions
	Description string `yaml:"description,omitempty"`
	Required    bool   `yaml:"required,omitempty"`
}

type ReturnType struct {
	Type  string `yaml:"type"`
	Ref   string `yaml:"$ref"`  // Add Ref for referencing definitions
	Items *Ref   `yaml:"items"` // Allow nil
}

type Ref struct {
	Ref string `yaml:"$ref"`
}

type Def struct {
	Type       string              `yaml:"type"`
	Properties map[string]Property `yaml:"properties"`
}

type Property struct {
	Type string `yaml:"type"`
}
