# zlogwrap

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

zlogwrap is a logging library, which wrapping all common use in My "Go App".

## Table of Contents

- [zlogwrap](#zlogwrap)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Signatures](#signatures)
  - [Examples](#examples)
  - [Config](#config)
  - [Default Config](#default-config)
  - [Dependencies](#dependencies)
  - [More Example Usage](#more-example-usage)

## Installation

```bash
  go get github.com/buildingwatsize/zlogwrap
```

## Signatures

```go
func New(conf ...Config) zerologWrapper
```

## Examples

Import the package

```go
import (
  "github.com/buildingwatsize/zlogwrap"
)
```

Using with these

```go
logger := zlogwrap.New()
logger.Debug("A Debug Log")

// {"level":"debug","time":"2021-07-27T16:49:59+07:00","message":"A Debug Log"}
```

Optional: With configuration

```go
logger := zlogwrap.New(zlogwrap.Config{
  ServiceName: "zlogwrap",
})
logger.Debug("A Debug Log")

// {"level":"debug","service":"zlogwrap","time":"2021-07-27T16:49:59+07:00","message":"A Debug Log"}
```

or customize the "Logger" and this library support middleware [RequestID](https://github.com/gofiber/docs/blob/master/api/middleware/requestid.md)

```go
myLogger := zerolog.New(os.Stdout).
  With().
  Str("foo", "bar").
  Float64("money", 10.99).
  Logger()
logger := zlogwrap.New(zlogwrap.Config{
  ServiceName: serviceName,
  Logger:      myLogger,
  Context:     c, // Which type is `*fiber.Ctx`
})
logger.Debug("A Debug Log")

// ps. see the real usage in `example/main.go`
// {"level":"debug","foo":"bar","money":10.99,"service":"zlogwrap","transaction_id":"b91f468c-c608-4729-88d8-9a12c951c31a","url":"/","message":"A Debug Log"}
```

## Config

```go
// Config defines the config for library.
type Config struct {

  // Optional. Default: false
  Hidden bool

  // Optional. Default: ""
  ServiceName string

  // Optional. Default: &fiber.Ctx{}
  Context *fiber.Ctx

  // Optional. Default: log.Logger
  Logger zerolog.Logger
}
```

## Default Config

```go
var ConfigDefault = Config{
  Hidden:      false,
  ServiceName: "",
  Context:     nil,
  Logger:      log.Logger,
}
```

## Dependencies

- [Zerolog](https://github.com/rs/zerolog)
- [Fiber](https://github.com/gofiber/fiber)

## More Example Usage

Please go to [example/main.go](./example/main.go) and [example/more_usage.go](./example/more_usage.go)
