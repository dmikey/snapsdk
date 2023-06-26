package main

import (
	"fmt"
	"strings"
)

type GoGenerator struct{}

func (GoGenerator) Name() string { return "go" }
func (g GoGenerator) Generate(stubby Stubby) (string, error) {
	var output strings.Builder

	fmt.Fprintf(&output, "// Generated by Stubby version: %s\n", stubby.Stubby)
	fmt.Fprintf(&output, "package %s\n\n", strings.ToLower(stubby.Namespace))

	// Generate struct definitions.
	for name, def := range stubby.Defs {
		fmt.Fprintf(&output, "type %s struct {\n", name)
		for propName, prop := range def.Properties {
			fmt.Fprintf(&output, "\t%s %s\n", propName, strings.Title(prop.Type))
		}
		fmt.Fprintf(&output, "}\n\n")
	}

	// Generate code for each object and its methods.
	for objName, obj := range stubby.Objects {
		for name, method := range obj.Methods {
			// Create a list of parameter names and types.
			paramNames := []string{}
			params := []string{}
			for _, param := range method.Parameters {
				// Map the STUBBY type to a Go type.
				goType := ""
				switch param.Type {
				case "string":
					goType = "string"
				case "integer":
					goType = "int"
				case "boolean":
					goType = "bool"
				case "array":
					goType = "[]interface{}"
				default:
					// If the type is not a primitive, it must be a reference to a definition.
					if param.Ref != "" {
						refName := strings.Split(param.Ref, "/")[2] // Assuming '$ref' is like '#/definitions/User'
						goType = refName
					} else {
						goType = param.Type
					}
				}

				paramNames = append(paramNames, param.Name)
				params = append(params, fmt.Sprintf("%s %s", param.Name, goType))
			}

			// Write the function signature and doc comment.
			fmt.Fprintf(&output, "// %s\n", method.Description)
			fmt.Fprintf(&output, "func (o *%s) %s(%s) ", strings.Title(objName), strings.Title(name), strings.Join(params, ", "))

			// Handle return type if any
			if method.ReturnType != nil {
				// Check if the ReturnType is a Ref to a definition
				if method.ReturnType.Ref != "" {
					refName := strings.Split(method.ReturnType.Ref, "/")[2] // Assuming '$ref' is like '#/definitions/User'
					fmt.Fprintf(&output, "(%s, error) {\n", refName)
				} else {
					fmt.Fprintf(&output, "() {\n")
				}
			} else {
				fmt.Fprintf(&output, "() {\n")
			}

			// Generate a receiver call if receiver exists
			if receiver, ok := method.Receiver["go"]; ok {
				fmt.Fprintf(&output, "\treturn o.%s(%s)\n", receiver, strings.Join(paramNames, ", "))
			} else {
				// For now, just return.
				fmt.Fprintf(&output, "\t// return ...\n")
			}

			fmt.Fprintf(&output, "}\n\n")
		}
	}

	return output.String(), nil
}