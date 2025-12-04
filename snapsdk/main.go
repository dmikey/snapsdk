package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	// Check for subcommands first
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
		serveCmd.Parse(os.Args[2:])

		if serveCmd.NArg() < 1 {
			fmt.Println("Usage: snapsdk serve <spec_file>")
			os.Exit(1)
		}

		specFile := serveCmd.Arg(0)
		snap, err := loadSpec(specFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Start MCP server
		server, err := NewMCPServer(snap)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "Starting MCP server for", snap.Namespace)
		if err := server.Serve(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// Original generate command
	// Define the flag.
	outputDir := flag.String("o", ".", "Output directory")

	// Parse the flags.
	flag.Parse()

	// Use flag.Args() to get non-flag arguments.
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: snapsdk [-o <output_dir>] <spec_file> <language>")
		fmt.Println("       snapsdk serve <spec_file>")
		os.Exit(1)
	}
	specFile := flag.Arg(0)
	language := flag.Arg(1)

	snap, err := loadSpec(specFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Always generate markdown
	generateForLanguage(MarkdownGenerator{}, *outputDir, snap)

	// Check the language argument to decide which generator(s) to use.
	if language == "*" {
		// Generate SDKs for all languages.
		generateForLanguage(JavascriptGenerator{}, *outputDir, snap)
		generateForLanguage(PythonGenerator{}, *outputDir, snap)
		generateForLanguage(RustGenerator{}, *outputDir, snap)
		generateForLanguage(GoGenerator{}, *outputDir, snap)
	} else {
		// Generate SDK for a single language.
		switch language {
		case "js":
			generateForLanguage(JavascriptGenerator{}, *outputDir, snap)
		case "py":
			generateForLanguage(PythonGenerator{}, *outputDir, snap)
		case "rs":
			generateForLanguage(RustGenerator{}, *outputDir, snap)
		case "go":
			generateForLanguage(GoGenerator{}, *outputDir, snap)
		default:
			fmt.Printf("Unsupported language: %s\n", language)
			os.Exit(1)
		}
	}
}

func loadSpec(specFile string) (Snap, error) {
	data, err := ioutil.ReadFile(specFile)
	if err != nil {
		return Snap{}, err
	}

	var snap Snap
	err = yaml.Unmarshal(data, &snap)
	if err != nil {
		return Snap{}, err
	}

	return snap, nil
}

func generateForLanguage(generator Generator, dir string, snap Snap) {
	// Make sure the output directory exists.
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open the output file.
	filename := fmt.Sprintf("%s/%s.%s", dir, snap.Namespace, generator.Name())
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Use the appropriate generator.
	genOutput, err := generator.Generate(snap)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Write the generated content to the file.
	_, err = f.WriteString(genOutput)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
