go-wren
=======

[![GoDoc](https://godoc.org/github.com/xyproto/go-wren?status.png)](https://godoc.org/github.com/xyproto/go-wren)

This package provides Go bindings to the [Wren](http://wren.io/index.html) scripting language, version 0.4.0.

This is a fork of https://github.com/dradtke/go-wren

Includes a simple REPL in cmd/repl

Getting Started
---------------

Run these commands to download the package, build the library, and run tests:

```
$ go get -d github.com/xyproto/go-wren
$ cd ${GOPATH}/src/github.com/xyproto/go-wren
$ (cd wren && make static)
$ go test
```
