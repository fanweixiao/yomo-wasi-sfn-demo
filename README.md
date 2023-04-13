# yomo-wasi-sfn-demo

This demo shows:

- Implement a [YoMo](https://github.com/yomorun/yomo) Stream Serverless Function in Go.
- Compiled to WebAssembly by [tinygo](https://tinygo.org/docs/guides/webassembly/) with `--target=wasi`.
- parse JSON while [tinygo's limited support for reflection](https://wazero.io/languages/tinygo/#unsupported-standard-libraries).

```bash
# Install YoMo
$ curl -fsSL https://get.yomo.run | sh

# Build
$ yomo build --target wasm -m go.mod app.go

# Test Run
$ yomo dev sfn.wasm
```
