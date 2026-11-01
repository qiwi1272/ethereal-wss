# Golang Websocket Client for Meridian API

[![Go Reference](https://pkg.go.dev/badge/github.com/roundinternetmoney/ethereal-wss/v2.svg)](https://pkg.go.dev/github.com/roundinternetmoney/ethereal-wss/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/roundinternetmoney/ethereal-wss/v2)](https://goreportcard.com/report/github.com/roundinternetmoney/ethereal-wss/v2)
[![CI](https://github.com/roundinternetmoney/ethereal-wss/actions/workflows/ci.yml/badge.svg)](https://github.com/roundinternetmoney/ethereal-wss/actions/workflows/ci.yml)

A proto-based websocket client for the [Meridian](https://meridian.xyz) exchange (formerly Ethereal).

## Migrating from Ethereal

- `Testnet` / `Mainnet` now point at `wss://ws.meridiantest.net/v1/stream` / `wss://ws.meridian.xyz/v1/stream` (note `ws`, not `ws2`). The old hosts remain as deprecated `EtherealTestnet` / `EtherealMainnet`.
- The message protocol and event names are unchanged. Account streams need your new Meridian subaccount ID; Ethereal IDs do not carry over.
- The protobuf types are still named `etherealv1`. New Meridian fields (such as `PositionUpdate.pfee`) are dropped until the protos add them.

## Features
- Protobuf support.
- Minimal dependencies

## Getting started

- Requires Go 1.25+.
- Install from GitHub: `go get github.com/roundinternetmoney/ethereal-wss/v2`

## Example Usage

```go
package main

import (
    "context"
    "log"

    "github.com/roundinternetmoney/ethereal-wss/v2"
    etherealv1 "roundinternet.money/protos/gen/dex/ethereal/v1"
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
    client.OnTicker(func(msg *etherealv1.Ticker) {
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
- This client depends on protobuf wrappers from [pkg.go.dev/roundinternet.money/protos](https://pkg.go.dev/roundinternet.money/protos)
- If you want to extend the `.proto` files directly, see the Buf module at [buf.build/round-internet-money/dex](https://buf.build/round-internet-money/dex)
- Otherwise, use or fork [github.com/roundinternetmoney/protos](https://github.com/roundinternetmoney/protos)

Contributing
-------------
Contributions are welcome! Please open issues or pull requests as needed.

## Modifying the package
- This client depends on protobuf wrappers from [pkg.go.dev/qiwi1272/protos](https://pkg.go.dev/qiwi1272/protos)

## Todo
- Add a `resubscribe` intent helper (Meridian closes connections after 239 minutes).
- Move to Meridian protos (`meridianv1`, `Dex_DEX_MERIDIAN`) once [roundinternetmoney/protos](https://github.com/roundinternetmoney/protos) publishes them.
  - Fix field names that don't match the wire and are silently dropped: `OrderFill` (proto uses long names; wire sends `oid`, `px`, `sz`, `m`, ...), `sd` on `OrderUpdate`/`PositionUpdate`, and `OrderUpdate.otyp` (string, not enum).
  - Add new Meridian fields: `PositionUpdate.pfee`, `OrderUpdate.rr`, `TokenTransfer.toTName`/`toTAddr`.
- Rename the module path and package from `ethereal` to `meridian`.
