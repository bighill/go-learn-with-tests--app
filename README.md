# Tutorial

## Learn Go with Tests

[https://github.com/quii/learn-go-with-tests/blob/master/README.md](https://github.com/quii/learn-go-with-tests/blob/master/README.md)

This repo holds samples from the `Build an application` sections

## Usage

_TODO_

## General getting started

[Getting started on go.dev](https://go.dev/doc/tutorial/getting-started)

```
$ go mod init example/hello
go: creating new go.mod: module example/hello
```

**External package**

After importing a new external package, add as requirement to go.sum

```
# Assume rsc.io/quote was added to project like this...
# import "rsc.io/quote"

# Then ...
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```
