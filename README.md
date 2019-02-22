# go-material-components-web

Go language WASM bindings to the material-components-web library.

## Development

The Go module should install the ``wasm-server`` included with the
``dennwc/dom`` library.  If it's not available try running the following
command:

``` shell
go install github.com/dennwc/dom/cmd/wasm-server
```

Running the following command in the root of the project should result
in server starting:

``` shell
wasm-server
```

Navigate to <http://localhost:8080> and the the catalog should be compiled,
loaded into the browser and run.

## Links

1. [Go Web Assembly Wiki](https://github.com/golang/go/wiki/WebAssembly)
2. [DOM Library for Go and WASM](https://github.com/dennwc/dom)
3. [Block-Element-Modifier](http://getbem.com/)
