package main

type Snap struct {
	Snap    string            `yaml:"snap"`
	Info      Info              `yaml:"info"`
	Namespace string            `yaml:"namespace"`
	Objects   map[string]Object `yaml:"objects"`
	Defs      map[string]Def    `yaml:"definitions"`
}

type Info struct {
	Version string `yaml:"version"`
	Title   string `yaml:"title"`
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
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Ref  string `yaml:"$ref"` // Add Ref for referencing definitions
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
