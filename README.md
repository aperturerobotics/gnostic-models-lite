# Gnostic models lite

This fork retains the upstream module path `github.com/google/gnostic-models`.
It keeps OpenAPI v2 and v3 model types, their document readers, and the compiler
and extension helpers those readers use. Protobuf messages use
`protobuf-go-lite` without protobuf runtime reflection. Decoding always copies
strings, including through the legacy `UnmarshalVTUnsafe` entry point.

The discovery and JSON Schema models, extension executable entry point,
code generators, generator inputs, build scripts and
upstream automation are omitted. This is a focused subset rather than a drop-in
replacement for every upstream package.

Copyright 2017-2025 Google LLC. Licensed under the Apache License 2.0.
