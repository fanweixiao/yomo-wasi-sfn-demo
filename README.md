# yomo-wasi-sfn-demo

This demo shows:

- Implement a [YoMo](https://github.com/yomorun/yomo) Stream Serverless Function in Go.
- Compiled to WebAssembly by tinygo with `--target=wasi`.
- parse JSON while tinygo's limited support for reflection.

```bash
$yomo build --target wasm -m go.mod app.go

$yomo dev sfn.wasm
```
