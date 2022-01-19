mackerel-plugin-bitcoin-core
=======

[![Test Status](https://github.com/dekokun/mackerel-plugin-bitcoin-core/workflows/test/badge.svg?branch=main)][actions]
[![Coverage Status](https://codecov.io/gh/dekokun/mackerel-plugin-bitcoin-core/branch/main/graph/badge.svg)][codecov]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dekokun/mackerel-plugin-bitcoin-core)][PkgGoDev]

[actions]: https://github.com/dekokun/mackerel-plugin-bitcoin-core/actions?workflow=test
[codecov]: https://codecov.io/gh/dekokun/mackerel-plugin-bitcoin-core
[license]: https://github.com/dekokun/mackerel-plugin-bitcoin-core/blob/master/LICENSE
[PkgGoDev]: https://pkg.go.dev/github.com/dekokun/mackerel-plugin-bitcoin-core

mackerel plugin to monitor bitcoin-core.

## Synopsis

```go
$ mackerel-plugin-bitcoin-core -host localhost -port 8332 -user <your bitcoind rpc user> -password <your bitcoind rpc password>
```

## Description

## Installation

```console
# go get
% go install github.com/dekokun/mackerel-plugin-bitcoin-core/cmd/mackerel-plugin-bitcoin-core@latest

# mkr plugin install
% mkr plugin install dekokun/mackerel-plugin-bitcoin-core
```

## Author

[dekokun](https://github.com/dekokun)
