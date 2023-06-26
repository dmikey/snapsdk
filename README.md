# SnapSDK

**Experience the next level of SDK generation with SnapSDK!** This interactive introspection spec is not just another tool, but a revolution in the way SDKs are created. Providing a streamlined, intuitive process for generating Stub Files for a wide range of languages, SnapSDK takes the lead by adopting a spec closely resembling the OpenAPI Spec.

Why stop at SDK creation? SnapSDK also offers generation of Markdown Docs straight from your Specification File! No more time-consuming documentation process - create, implement, and share with ease.

SnapSDK boasts a unique feature - it accepts \`recievers\` which are mapped to \`Namespace.Method\` calls. The arguments and returns are cross-mapped, allowing you to regenerate new SDK definitions while keeping your implementation code intact. This degree of flexibility and control is what makes SnapSDK a standout.

## Usage

Getting started with SnapSDK is as simple as 1, 2, 3, 4, 5!

1.  Ensure you have Go installed on your system.
2.  Clone the snapsdk repository: `git clone https://github.com/your-username/snapsdk.git`
3.  Navigate to the snapsdk directory: `cd snapsdk`
4.  Build the snapsdk binary: `go build -o snapsdk`
5.  Run snapsdk by providing the necessary arguments: `./snapsdk <spec_file> <language> [-o <output_dir>]`

**<spec_file>**: The path to the specification file in YAML format.

**<language>**: The target language for generating the SDK. Use "\*" to generate SDKs for all supported languages.

**\[-o <output_dir>\]** (optional): The output directory for the generated SDK files. Defaults to the current directory.

Example usage: `./snapsdk spec.yaml go -o sdk/go`

This command generates the Go SDK based on the \`spec.yaml\` specification file and places the generated files in the \`sdk/go\` directory.

## Specification File

With SnapSDK, your specification file is your canvas. Provide it in YAML format and let SnapSDK take care of the rest. The file should adhere to the SnapSDK specification schema, defining the structure and details of the API for which the SDKs will be generated. Dive deeper into the specification schema in [SPEC.md](SPEC.md).

SnapSDK supports a broad spectrum of primitive data types and special types for defining the structure and data types of the API endpoints and models within the specification file. Discover the power of customization with our unique support for custom receiver specifications. With SnapSDK, you get to design receiver methods for your generated stubs, decoupling the generator map for the SDK from your actual SDK implementation.

## Supported Languages

SnapSDK is a polyglot's delight! We currently support generating SDKs for a variety of programming languages:

- JavaScript (js)
- Python (py)
- Rust (rs)
- Go (go)

## Contributing

We believe in the power of collective growth. Contributions to SnapSDK are always welcome! If you encounter any issues or have suggestions for improvements, don't hesitate to open an issue on the GitHub repository. Together, we can make SnapSDK better.

## License

Freedom is fundamental to development. This project is licensed under the MIT License, empowering you to innovate and create with SnapSDK
