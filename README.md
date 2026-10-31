# Golang Websocket Client for Meridian API

[![Go Reference](https://pkg.go.dev/badge/github.com/qiwi1272/meridian-wss.svg)](https://pkg.go.dev/github.com/qiwi1272/meridian-wss)
[![Go Report Card](https://goreportcard.com/badge/github.com/qiwi1272/meridian-wss)](https://goreportcard.com/report/github.com/qiwi1272/meridian-wss)
[![CI](https://github.com/qiwi1272/meridian-wss/actions/workflows/ci.yml/badge.svg)](https://github.com/qiwi1272/meridian-wss/actions/workflows/ci.yml)

A proto-based websocket client for the [Meridian](https://meridian.xyz) exchange (formerly Ethereal).

## Migrating from Ethereal

- The module moved from `github.com/roundinternetmoney/ethereal-wss/v2` to `github.com/qiwi1272/meridian-wss` (no `/v2`; versions restart at `v1.1.0`), and the package is now `meridian`. Update your imports. Versions up to `v2.0.6` stay available at the old path.
- `Testnet` / `Mainnet` now point at `wss://ws.meridiantest.net/v1/stream` / `wss://ws.meridian.xyz/v1/stream` (note `ws`, not `ws2`). The old hosts remain as deprecated `EtherealTestnet` / `EtherealMainnet`.
- The message protocol and event names are unchanged. Account streams need your new Meridian subaccount ID; Ethereal IDs do not carry over.
- The protobuf types are still named `etherealv1`. New Meridian fields (such as `PositionUpdate.pfee`) are dropped until the protos add them.

## Features
- Protobuf support.
- Minimal dependencies

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

## Releasing

Releases are manual and GitHub marks them immutable, so a published tag can't be moved or deleted.

1. Merge to `main`. Merging does not publish anything.
2. Go to **Actions → Release → Run workflow**, enter the version (for example `1.1.0`) and run it.

The workflow runs the checks, then tags `v<version>` and publishes the GitHub release. Locally, `make release NEW_VERSION=1.1.0` does the same. Only `0.x.y` / `1.x.y` versions that aren't already tagged are accepted, because the module path has no `/vN` suffix.

Contributing
-------------
Contributions are welcome! Please open issues or pull requests as needed.


## Todo
- Add a `resubscribe` intent helper (Meridian closes connections after 239 minutes).
- Move to Meridian protos (`meridianv1`, `Dex_DEX_MERIDIAN`) once [roundinternetmoney/protos](https://github.com/roundinternetmoney/protos) publishes them.
  - Fix field names that don't match the wire and are silently dropped: `OrderFill` (proto uses long names; wire sends `oid`, `px`, `sz`, `m`, ...), `sd` on `OrderUpdate`/`PositionUpdate`, and `OrderUpdate.otyp` (string, not enum).
  - Add new Meridian fields: `PositionUpdate.pfee`, `OrderUpdate.rr`, `TokenTransfer.toTName`/`toTAddr`.
