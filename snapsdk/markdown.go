package main

import (
	"fmt"
	"strings"
)

type MarkdownGenerator struct{}

func (MarkdownGenerator) Name() string { return "md" }
func (m MarkdownGenerator) Generate(snap Snap) (string, error) {
	var output strings.Builder

	fmt.Fprintf(&output, "# %s\n\n", snap.Info.Title)

	// Generate Definitions
	for objName, obj := range snap.Objects {
		fmt.Fprintf(&output, "## Object: %s\n", objName)
		if def, ok := obj.Definitions["Dog"]; ok {
			fmt.Fprintf(&output, "### Definitions\n")
			for propName, prop := range def.Properties {
				fmt.Fprintf(&output, "- %s: %s\n", propName, prop.Type)
			}
			fmt.Fprintf(&output, "\n")
		}

		// Generate code for each object and its methods.
		fmt.Fprintf(&output, "### Methods\n")
		for name, method := range obj.Methods {
			// Create a list of parameter names.
			params := []string{}
			for _, param := range method.Parameters {
				params = append(params, fmt.Sprintf("%s (%s)", param.Name, param.Type))
			}

			// Write the method.
			fmt.Fprintf(&output, "#### %s\n", name)
			fmt.Fprintf(&output, "- Summary: %s\n", method.Summary)
			fmt.Fprintf(&output, "- Description: %s\n", method.Description)
			if len(params) > 0 {
				fmt.Fprintf(&output, "- Parameters: %s\n", strings.Join(params, ", "))
			}

			// Generate usage examples for each language.
			fmt.Fprintf(&output, "- Usage Examples:\n")
			for lang, _ := range method.Receiver {
				// Generate the method call example
				switch lang {
				case "go":
					fmt.Fprintf(&output, "  - Go: `%s.%s(%s)`\n", strings.Title(objName), strings.Title(name), strings.Join(params, ", "))
				case "python":
					fmt.Fprintf(&output, "  - Python: `%s.%s(%s)`\n", objName, name, strings.Join(params, ", "))
				case "javascript":
					fmt.Fprintf(&output, "  - JavaScript: `%s.%s(%s)`\n", objName, name, strings.Join(params, ", "))
				case "rust":
					fmt.Fprintf(&output, "  - Rust: `%s::%s(%s)`\n", objName, name, strings.Join(params, ", "))
				}
			}

			// Handle return type if any
			if method.ReturnType != nil {
				// Check if the ReturnType is a Ref to a definition
				if method.ReturnType.Ref != "" {
					refName := strings.Split(method.ReturnType.Ref, "/")[2] // Assuming '$ref' is like '#/definitions/User'
					fmt.Fprintf(&output, "- Returns: %s\n\n", refName)
				} else if method.ReturnType.Type != "" {
					fmt.Fprintf(&output, "- Returns: %s\n\n", method.ReturnType.Type)
				}
			}
		}
		fmt.Fprintf(&output, "\n")
	}

	fmt.Fprintf(&output, "##### Generated with Snap version: %s\n", snap.Snap)

	return output.String(), nil
}
