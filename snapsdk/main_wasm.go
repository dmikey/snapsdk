//go:build wasm
// +build wasm

package main

import (
	"encoding/json"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"syscall/js"
)

func main() {
	// Expose the generateAll function to JavaScript.
	js.Global().Set("generateAll", js.FuncOf(generateAll))
	select {} // keep running
}

func generateAll(this js.Value, args []js.Value) interface{} {
	// Convert JavaScript value to Go string.
	specYaml := args[0].String()

	// Convert YAML string to our Stubby struct.
	var stubby Stubby
	err := yaml.Unmarshal([]byte(specYaml), &stubby)
	if err != nil {
		return "Error in YAML unmarshal: " + err.Error()
	}

	// Create a map to hold the sources of all languages and the stubby version.
	sources := make(map[string]interface{})

	// Generate sources for all languages and store in the map.
	sources["go"], _ = generateForLanguage("go", stubby)
	sources["js"], _ = generateForLanguage("js", stubby)
	sources["py"], _ = generateForLanguage("py", stubby)
	sources["rs"], _ = generateForLanguage("rs", stubby)
	sources["md"], _ = generateForLanguage("md", stubby)

	// Include the stubby version.
	sources["stubby"] = stubby.Stubby

	// Convert the map to JSON.
	jsonSources, err := json.Marshal(sources)
	if err != nil {
		return "Error in JSON marshal: " + err.Error()
	}

	// Return the JSON string.
	return string(jsonSources)
}

func generateForLanguage(lang string, stubby Stubby) (string, error) {
	// Use the appropriate generator.
	var genOutput string
	var err error
	switch lang {
	case "js":
		genOutput, err = WriteSDK(JavascriptGenerator{}, stubby)
	case "py":
		genOutput, err = WriteSDK(PythonGenerator{}, stubby)
	case "rs":
		genOutput, err = WriteSDK(RustGenerator{}, stubby)
	case "go":
		genOutput, err = WriteSDK(GoGenerator{}, stubby)
	case "md":
		genOutput, err = WriteSDK(MarkdownGenerator{}, stubby)
	default:
		return "", fmt.Errorf("Unsupported language: %s", lang)
	}

	if err != nil {
		return "", err
	}

	return genOutput, nil
}
