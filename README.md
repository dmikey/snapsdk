# SnapSDK

> ⚠️ **Early Development** - This project is not production ready. APIs and features may change significantly.

**Experience the next level of SDK generation with SnapSDK!** Think of it as "Apollo for SDKs" — a schema-first approach to generating type-safe, multi-language SDK libraries from a single specification file.

SnapSDK provides a streamlined process for generating SDK Stub Files across multiple languages, using a spec format inspired by OpenAPI but optimized for library/SDK development rather than HTTP APIs.

## Key Features

- **Schema-First Development**: Define your SDK interface once, generate implementations for all supported languages
- **Type-Safe Code Generation**: Full support for primitives, arrays, and custom object types with proper language-specific mappings
- **Receiver Pattern**: Map generated methods to concrete implementations (similar to GraphQL resolvers), keeping your business logic decoupled from the generated interface
- **Auto-Generated Documentation**: Markdown docs generated directly from your specification
- **Multi-Language Support**: JavaScript, Python, Rust, and Go from a single YAML spec
- **Built-in MCP Server**: Expose your SDK as an MCP server for AI agent integration

## Usage

Getting started with SnapSDK:

1. Ensure you have Go installed on your system
2. Clone the repository: `git clone https://github.com/dmikey/snapsdk.git`
3. Navigate to the snapsdk directory: `cd snapsdk/snapsdk`
4. Build the binary: `go build -o snapsdk`

### Generate SDKs

```bash
./snapsdk -o <output_dir> <spec_file> <language>
```

| Argument | Description |
|----------|-------------|
| `-o <output_dir>` | Output directory for generated files (default: current directory) |
| `<spec_file>` | Path to the YAML specification file |
| `<language>` | Target language: `js`, `py`, `rs`, `go`, or `*` for all |

### Examples

```bash
# Generate all SDKs to the sdk/ directory
./snapsdk -o sdk/ spec.yaml "*"

# Generate only the Go SDK
./snapsdk -o sdk/go spec.yaml go

# Generate Python SDK to current directory
./snapsdk spec.yaml py
```

### MCP Server Mode

SnapSDK can run as an MCP (Model Context Protocol) server, allowing AI agents to discover and call your SDK methods as tools.

```bash
./snapsdk serve <spec_file>
```

This starts an MCP server on stdio that:
- Exposes each SDK method as an MCP tool
- Forwards tool calls to your HTTP receiver
- Returns results back to the AI agent

#### Architecture

```
┌─────────────┐     ┌─────────────────────┐     ┌─────────────────┐
│  AI Agent   │────▶│  snapsdk serve      │────▶│  Your Receiver  │
│  (Claude)   │stdio│  (MCP protocol)     │http │  (any language) │
└─────────────┘     └─────────────────────┘     └─────────────────┘
```

#### MCP Configuration

Add an `mcp` section to your spec:

```yaml
mcp:
  enabled: true
  transport: stdio
  description: A tool for managing dog data
  receiver:
    type: http
    url: http://localhost:8080/api
```

#### Running with MCP

1. Start your receiver (implements the actual SDK logic):
```bash
node receiver.js  # or python receiver.py, etc.
```

2. Start the MCP server:
```bash
./snapsdk serve spec.yaml
```

3. Connect your AI agent to the MCP server via stdio

## Observability

SnapSDK includes built-in observability with OpenTelemetry tracing and Prometheus metrics.

### Configuration

Add an `observability` section to your MCP config:

```yaml
mcp:
  enabled: true
  receiver:
    type: http
    url: http://localhost:8080/api
  observability:
    enabled: true
    service_name: my-sdk
    prometheus_port: 9090        # Exposes /metrics endpoint
    otlp_endpoint: localhost:4317  # OTel collector endpoint
```

### Metrics

When enabled, SnapSDK exposes these Prometheus metrics at `http://localhost:<prometheus_port>/metrics`:

| Metric | Type | Description |
|--------|------|-------------|
| `snapsdk_tool_calls_total` | Counter | Total tool calls by tool name and status |
| `snapsdk_tool_duration_seconds` | Histogram | Tool call latency |
| `snapsdk_tool_errors_total` | Counter | Failed tool calls |
| `snapsdk_receiver_latency_seconds` | Histogram | HTTP receiver response time |

### Tracing

When `otlp_endpoint` is configured, SnapSDK sends traces to your OTel collector:

```
snapsdk.tool.get_dog
├── receiver.http.post
│   └── duration: 45ms
└── status: ok
```

## Specification File

Define your SDK interface in YAML format following the SnapSDK schema. See [SPEC.md](SPEC.md) for the complete specification.

### Example Specification

```yaml
snap: '1.0'
info:
  version: 1.0.0
  title: DogsApp SDK
namespace: DogsAppSDK
objects:
  dog:
    methods:
      getDog:
        operationId: getDog
        summary: Returns a dog by ID.
        description: Returns detailed information about a dog.
        parameters:
          - name: id
            type: integer
        returnType:
          $ref: '#/definitions/Dog'
        receiver:
          go: DogsAppReceiver.GetDog
          python: dogsapp_receiver.get_dog
          javascript: dogsAppReceiver.getDog
          rust: 'dogsapp_receiver::get_dog'
    definitions:
      Dog:
        type: object
        properties:
          id:
            type: integer
          name:
            type: string
```

### Type Mappings

| SnapSDK Type | JavaScript | Python | Go | Rust |
|--------------|------------|--------|-----|------|
| `string` | `string` | `str` | `string` | `String` |
| `integer` | `number` | `int` | `int` | `i32` |
| `boolean` | `boolean` | `bool` | `bool` | `bool` |
| `number` | `number` | `float` | `float64` | `f64` |
| `array` | `Array` | `list` | `[]T` | `Vec<T>` |
| `object` | `Object` | `dict` | `struct` | `struct` |

## Supported Languages

- **JavaScript** (`js`) - ES6 classes with CommonJS exports
- **Python** (`py`) - Classes with type hints ready
- **Rust** (`rs`) - Modules with structs and impl blocks
- **Go** (`go`) - Structs with receiver methods

## The Receiver Pattern

Similar to how Apollo GraphQL uses resolvers, SnapSDK uses **receivers** to map generated SDK methods to your actual implementation:

```yaml
receiver:
  go: MyReceiver.DoSomething
  python: my_receiver.do_something
  javascript: myReceiver.doSomething
  rust: 'my_receiver::do_something'
```

This allows you to regenerate SDK definitions without losing your implementation code.

## Contributing

Contributions are welcome! Open an issue or submit a PR on the [GitHub repository](https://github.com/dmikey/snapsdk).

## License

This project is licensed under the MIT License.
