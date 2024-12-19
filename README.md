# rwlocker

[![Go Reference](https://pkg.go.dev/badge/badge/github.com/linhns/rwlocker.svg)](https://pkg.go.dev/github.com/linhns/rwlocker)
[![Go Report Card](https://goreportcard.com/badge/github.com/linhns/rwlocker)](https://goreportcard.com/report/github.com/linhns/rwlocker)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Package rwlocker supplements the [sync](https://pkg.go.dev/sync) package
from the Go standard library. It introduces a convenient `RWLocker` interface
to represent an object that can be locked by multiple writers and a single reader.

## Features

- `rwlocker.RWLocker`: Embeds [sync.Locker](https://pkg.go.dev/sync#Locker)
   and adds `RLock()` and `RUnlock()` methods for read locks.
   By default, [sync.RWMutex](https://pkg.go.dev/sync#RWMutex) implements this interface.

- **Null implementation**: `rwlocker.NullRWLocker` is a null implementation of
  `rwlocker.RWLocker`. This can be used as a substitute where concurrency safety
  is not required.

## Installation

Install via `go get`:

```shell
go get github.com/linhns/rwlocker
```

## License

This package is licensed under the MIT License. See the [LICENSE](./LICENSE)
file for more details.
