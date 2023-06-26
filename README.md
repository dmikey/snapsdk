Stubby
======

**Experience the next level of SDK generation with Stubby!** This interactive introspection spec is not just another tool, but a revolution in the way SDKs are created. Providing a streamlined, intuitive process for generating Stub Files for a wide range of languages, Stubby takes the lead by adopting a spec closely resembling the OpenAPI Spec.

Why stop at SDK creation? Stubby also offers generation of Markdown Docs straight from your Specification File! No more time-consuming documentation process - create, implement, and share with ease.

Stubby boasts a unique feature - it accepts \`recievers\` which are mapped to \`Namespace.Method\` calls. The arguments and returns are cross-mapped, allowing you to regenerate new SDK definitions while keeping your implementation code intact. This degree of flexibility and control is what makes Stubby a standout.

Usage
-----

Getting started with Stubby is as simple as 1, 2, 3, 4, 5!

1.  Ensure you have Go installed on your system.
2.  Clone the Stubby repository: `git clone https://github.com/your-username/stubby.git`
3.  Navigate to the Stubby directory: `cd stubby`
4.  Build the Stubby binary: `go build -o stubby`
5.  Run Stubby by providing the necessary arguments: `./stubby <spec_file> <language> [-o <output_dir>]`

**<spec\_file>**: The path to the specification file in YAML format.

**<language>**: The target language for generating the SDK. Use "\*" to generate SDKs for all supported languages.

**\[-o <output\_dir>\]** (optional): The output directory for the generated SDK files. Defaults to the current directory.

Example usage: `./stubby spec.yaml go -o sdk/go`

This command generates the Go SDK based on the \`spec.yaml\` specification file and places the generated files in the \`sdk/go\` directory.

Specification File
------------------

With Stubby, your specification file is your canvas. Provide it in YAML format and let Stubby take care of the rest. The file should adhere to the Stubby specification schema, defining the structure and details of the API for which the SDKs will be generated. Dive deeper into the specification schema in [SPEC.md](SPEC.md).

Stubby supports a broad spectrum of primitive data types and special types for defining the structure and data types of the API endpoints and models within the specification file. Discover the power of customization with our unique support for custom receiver specifications. With Stubby, you get to design receiver methods for your generated stubs, decoupling the generator map for the SDK from your actual SDK implementation.

Supported Languages
-------------------

Stubby is a polyglot's delight! We currently support generating SDKs for a variety of programming languages:

*   JavaScript (js)
*   Python (py)
*   Rust (rs)
*   Go (go)

Contributing
------------

We believe in the power of collective growth. Contributions to Stubby are always welcome! If you encounter any issues or have suggestions for improvements, don't hesitate to open an issue on the GitHub repository. Together, we can make Stubby better.

License
-------

Freedom is fundamental to development. This project is licensed under the MIT License, empowering you to innovate and create with Stubby.