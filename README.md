# mackerel-plugin-bitcoin

[![Test Status](https://github.com/dekokun/mackerel-plugin-bitcoin/workflows/test/badge.svg?branch=main)][actions]
[![Coverage Status](https://codecov.io/gh/dekokun/mackerel-plugin-bitcoin/branch/main/graph/badge.svg)][codecov]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dekokun/mackerel-plugin-bitcoin)][pkggodev]

[actions]: https://github.com/dekokun/mackerel-plugin-bitcoin/actions?workflow=test
[codecov]: https://codecov.io/gh/dekokun/mackerel-plugin-bitcoin
[license]: https://github.com/dekokun/mackerel-plugin-bitcoin/blob/master/LICENSE
[pkggodev]: https://pkg.go.dev/github.com/dekokun/mackerel-plugin-bitcoin

mackerel plugin to monitor bitcoin.

## Synopsis

```go
$ mackerel-plugin-bitcoin -host localhost -port 8332 -user <your bitcoind rpc user> -password <your bitcoind rpc password>
```

## Description

## Installation

```console
# go get
% go install github.com/dekokun/mackerel-plugin-bitcoin@latest

# mkr plugin install
% mkr plugin install dekokun/mackerel-plugin-bitcoin
```

## Author

[dekokun](https://github.com/dekokun)
