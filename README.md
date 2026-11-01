# Golang Websocket Client for Meridian API

[![Go Reference](https://pkg.go.dev/badge/github.com/qiwi1272/meridian-wss.svg)](https://pkg.go.dev/github.com/qiwi1272/meridian-wss)
[![Go Report Card](https://goreportcard.com/badge/github.com/qiwi1272/meridian-wss)](https://goreportcard.com/report/github.com/qiwi1272/meridian-wss)
[![CI](https://github.com/qiwi1272/meridian-wss/actions/workflows/ci.yml/badge.svg)](https://github.com/qiwi1272/meridian-wss/actions/workflows/ci.yml)

A proto-based websocket client for the [Meridian](https://meridian.xyz) exchange (formerly Ethereal).

## Migrating from Ethereal

- The module moved from `ethereal-wss/v2` to `meridian-wss`.
- `Testnet` / `Mainnet` now point at `wss://ws.meridiantest.net/v1/stream` / `wss://ws.meridian.xyz/v1/stream` (note `ws`, not `ws2`). The old hosts remain as deprecated `EtherealTestnet` / `EtherealMainnet`.
- The message protocol and event names are unchanged. Account streams need your new Meridian subaccount ID; Ethereal IDs do not carry over.
- The protobuf types moved from `ethereal/v1` to `meridian/v1`.

## Getting started

- Requires Go 1.25+.
- Install from GitHub: `go get github.com/qiwi1272/meridian-wss`

## Example Usage

```go
package main

import (
    "context"
    "log"

    "github.com/qiwi1272/meridian-wss"
    meridianv1 "roundinternet.money/protos/gen/dex/meridian/v1"
)

func main() {
    client := meridian.NewClient(context.Background(), meridian.Mainnet)
    defer client.Close()

    // Subscribe to ticker updates for BTCUSD
    err := client.Subscribe(context.Background(), meridian.EventTypeTicker, "BTCUSD")
    if err != nil {
        log.Fatal(err)
    }

    // typed callback for events
    client.OnTicker(func(msg *meridianv1.Ticker) {
        // message as typed struct
        log.Printf("Received ticker: %v", msg)
    })

    // Start listening
    log.Fatal(client.Listen(context.Background()))
}
```

For more examples, see the `examples/` directory.
Some proto stubs are built for use with NATS, but integration is optional. 
examples using these are prefixed with "nats_".

## Modifying the package
- This client depends on protobuf wrappers from [pkg.go.dev/github.com/qiwi1272/protos](https://pkg.go.dev/github.com/qiwi1272/protos)
- If you want to extend the `.proto` files directly, see the Buf module at [github.com/qiwi1272/protos](https://github.com/qiwi1272/protos)

Contributing
-------------
Contributions are welcome! Please open issues or pull requests as needed.

## Todo
- Add a `resubscribe` intent helper (Meridian closes connections after 239 minutes).
