# Developers Guide

Modified: 2024-11

## Protobuf

The protobuf schemas are managed using the `buf` ecosystem. First [install buf cli](https://buf.build/docs/installation/) using whatever method you prefer then checkout the source code.

There are four basic tests that are performed in our ci/cd pipelines that can be verified locally:
 1. Linting - check for syntax errors and standards enforcement
 2. Formatting - check for formatting correctness
 3. Breaking Change Detection - check for backwards compatibility violations against another version of the protobuf module
 3. Code generation - check that code can be generated from protobuf schemas

To lint protobuf modules:
```sh
buf lint
```

To show formatting diffs:
```sh
buf format -d
```

To format protobuf modules files in place:
```sh
buf format -w
```

To test whether your local changes have backwards compatibility issues with target source control:
```sh
buf breaking --against https://github.com/dronectl/protocols.git#branch=master
```

To generate code:
```sh
buf generate
```
This will generate code for the target modules under `gen` in the repository root.

### Tooling

Install the VSCode extension `bufbuild.vscode-buf` for syntax highlighting, and inline static analysis. For vim users, `bufls` is no longer supported and does not interpret a `v2` `buf.yaml` template. They are working on a new LSP which will work with `buf`: https://github.com/bufbuild/buf-language-server.