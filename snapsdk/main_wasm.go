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

	// Convert YAML string to our Snap struct.
	var snap Snap
	err := yaml.Unmarshal([]byte(specYaml), &snap)
	if err != nil {
		return "Error in YAML unmarshal: " + err.Error()
	}

	// Create a map to hold the sources of all languages and the snap version.
	sources := make(map[string]interface{})

	// Generate sources for all languages and store in the map.
	sources["go"], _ = generateForLanguage("go", snap)
	sources["js"], _ = generateForLanguage("js", snap)
	sources["py"], _ = generateForLanguage("py", snap)
	sources["rs"], _ = generateForLanguage("rs", snap)
	sources["md"], _ = generateForLanguage("md", snap)

	// Include the snap version.
	sources["snap"] = snap.Snap

	// Convert the map to JSON.
	jsonSources, err := json.Marshal(sources)
	if err != nil {
		return "Error in JSON marshal: " + err.Error()
	}

	// Return the JSON string.
	return string(jsonSources)
}

func generateForLanguage(lang string, snap Snap) (string, error) {
	// Use the appropriate generator.
	var genOutput string
	var err error
	switch lang {
	case "js":
		genOutput, err = WriteSDK(JavascriptGenerator{}, snap)
	case "py":
		genOutput, err = WriteSDK(PythonGenerator{}, snap)
	case "rs":
		genOutput, err = WriteSDK(RustGenerator{}, snap)
	case "go":
		genOutput, err = WriteSDK(GoGenerator{}, snap)
	case "md":
		genOutput, err = WriteSDK(MarkdownGenerator{}, snap)
	default:
		return "", fmt.Errorf("Unsupported language: %s", lang)
	}

	if err != nil {
		return "", err
	}

	return genOutput, nil
}
