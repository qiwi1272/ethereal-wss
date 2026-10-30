# Ethereal → Meridian rebrand checklist

Ethereal is rebranding to **Meridian** ([meridian.xyz](https://meridian.xyz)).
This branch makes the changes that are safe now. The rest is blocked on
information from Meridian, on the upstream protos, or on decisions that break
things for current users.

## Done on this branch

- README and the package doc comment now say Meridian.

## Blocked: needs info from Meridian

- [ ] **Stream hosts.** `Testnet` / `Mainnet` in `client.go` still point to
      `ws2.etherealtest.net` / `ws2.ethereal.trade` (marked `TODO(meridian)`).
      Swap them for the hosts in the Perpetuals API section of
      https://docs.meridian.xyz, and update `TestEnvironment` in
      `ethereal_test.go` to match. Also check that the Market Data docs still use
      the same subscribe format and event names.

## Blocked: upstream protos (`roundinternet.money/protos`)

- [ ] **Message types.** Every message type comes from
      `roundinternet.money/protos/gen/dex/ethereal/v1` (`etherealv1.*`).
      Renaming these means a new `dex/meridian/v1` package in
      [roundinternetmoney/protos](https://github.com/roundinternetmoney/protos)
      (Buf module `buf.build/round-internet-money/dex`), then a dependency bump
      here and an `etherealv1` → `meridianv1` change across `client.go`,
      `subscription.go`, the tests and the examples. Wire names such as
      `"Ticker"` and `"L2Book"` stay the same unless Meridian changes the
      protocol.
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
- [ ] **Tag a release** that marks the rename.
