# Ethereal → Meridian rebrand checklist

Ethereal is now **Meridian** ([docs.meridian.xyz](https://docs.meridian.xyz)).
Meridian is a separate deployment on Robinhood Chain: Ethereal subaccount IDs
do not carry over. See the official
[Migrate from Ethereal](https://docs.meridian.xyz/developer-guides/trading-api/migrate-from-ethereal)
guide.

## Done on this branch

- `Testnet` / `Mainnet` now point to `wss://ws.meridiantest.net/v1/stream` /
  `wss://ws.meridian.xyz/v1/stream` (from
  [API Hosts](https://docs.meridian.xyz/protocol-reference/api-hosts)). Note the
  host is `ws`, not `ws2`. The old hosts remain as deprecated
  `EtherealTestnet` / `EtherealMainnet`.
- README and the package doc comment now say Meridian.

## Checked against the docs, no change needed

- The native `ws2` JSON protocol carries over: same `subscribe` /
  `unsubscribe` envelope, `type` + `symbol` / `subaccountId`, event names
  (`L2Book`, `Ticker`, `TradeFill`, `SubaccountLiquidation`, `PositionUpdate`,
  `OrderUpdate`, `OrderFill`, `TokenTransfer`) and short field names (see
  [WebSockets](https://docs.meridian.xyz/developer-guides/trading-api/websockets)).
  Socket.IO was removed, but this client never used it.

## Still to do

- [ ] **Live testnet check.** Run `examples/log_bitcoin_events` against
      `ws.meridiantest.net`. This sandbox could reach the docs but not the
      stream hosts, so nothing here has been verified against the live API.
- [ ] **New payload fields.** Meridian added `PositionUpdate.pfee`,
      `OrderUpdate.rr`, `TokenTransfer.toTName` / `toTAddr` and the `CONVERT`
      transfer type. The client parses with `DiscardUnknown`, so it won't break,
      but these fields are dropped until the protos add them.
- [ ] **Reconnects.** Connections now close after 239 minutes (code `1000`), and
      a successful subscribe returns `{"ok": true}`. Consider automatic
      reconnect and resubscribe (this fits the README's `resubscribe` TODO).

## Blocked: upstream protos (`roundinternet.money/protos`)

- [ ] **Message types.** Every message type comes from
      `roundinternet.money/protos/gen/dex/ethereal/v1` (`etherealv1.*`).
      Renaming these means a new `dex/meridian/v1` package in
      [roundinternetmoney/protos](https://github.com/roundinternetmoney/protos)
      (Buf module `buf.build/round-internet-money/dex`), including the new
      fields above. Then bump the dependency here and change
      `etherealv1` → `meridianv1` across `client.go`, `subscription.go`, the
      tests and the examples.
- [ ] **`basev1.Dex_DEX_ETHEREAL`** is used in `examples/nats_ethereal_*`.
      Once protos adds a Meridian enum value, rename the examples to
      `nats_meridian_*` (and the `natsetherealserver` package).

## Breaking changes: decide and release together

- [ ] **Repo / module path.** The module is still
      `github.com/roundinternetmoney/ethereal-wss/v2`, while this repo lives at
      `qiwi1272/ethereal-wss`. Pick the final home (e.g. `.../meridian-wss`),
      rename the GitHub repo, then update `go.mod`, the example imports and the
      README badges and snippet.
- [ ] **Package name.** `ethereal` → `meridian`. The README snippet imports the
      package without an alias (`ethereal.NewClient`), so this change breaks
      callers. Make it in the same release as the module path, together with a
      rename of `ethereal_test.go`.
- [ ] **Tag a release** that marks the rename. Its notes should say that
      `Testnet` / `Mainnet` now point to Meridian.
