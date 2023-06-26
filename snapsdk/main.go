package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	// Define the flag.
	outputDir := flag.String("o", ".", "Output directory")

	// Parse the flags.
	flag.Parse()

	// Use flag.Args() to get non-flag arguments.
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: ./stubby <spec_file> <language> [-o <output_dir>]")
		os.Exit(1)
	}
	specFile := flag.Arg(0)
	language := flag.Arg(1)

	// Read the input file.
	data, err := ioutil.ReadFile(specFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmarshal the YAML to our Stubby struct.
	var stubby Stubby
	err = yaml.Unmarshal(data, &stubby)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Always generate markdown
	generateForLanguage(MarkdownGenerator{}, *outputDir, stubby)

	// Check the language argument to decide which generator(s) to use.
	if language == "*" {
		// Generate SDKs for all languages.
		generateForLanguage(JavascriptGenerator{}, *outputDir, stubby)
		generateForLanguage(PythonGenerator{}, *outputDir, stubby)
		generateForLanguage(RustGenerator{}, *outputDir, stubby)
		generateForLanguage(GoGenerator{}, *outputDir, stubby)
	} else {
		// Generate SDK for a single language.
		switch language {
		case "js":
			generateForLanguage(JavascriptGenerator{}, *outputDir, stubby)
		case "py":
			generateForLanguage(PythonGenerator{}, *outputDir, stubby)
		case "rs":
			generateForLanguage(RustGenerator{}, *outputDir, stubby)
		case "go":
			generateForLanguage(GoGenerator{}, *outputDir, stubby)
		default:
			fmt.Printf("Unsupported language: %s\n", language)
			os.Exit(1)
		}
	}
}

func generateForLanguage(generator Generator, dir string, stubby Stubby) {
	// Make sure the output directory exists.
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open the output file.
	filename := fmt.Sprintf("%s/%s.%s", dir, stubby.Namespace, generator.Name())
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Use the appropriate generator.
	genOutput, err := generator.Generate(stubby)
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
