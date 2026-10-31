.PHONY: examples log_bitcoin_events wss version check-version tag release

# Releases are manual: `make release NEW_VERSION=1.1.0` (or the Release workflow).
# The module path has no /vN suffix, so only v0/v1 versions are valid.
LATEST_TAG := $(shell git describe --tags --abbrev=0 --match 'v[01].*' 2>/dev/null || echo none)
RELEASE := $(patsubst v%,%,$(NEW_VERSION))

examples: log_bitcoin_events

log_bitcoin_events:
	go build -o bin/log_bitcoin_events examples/log_bitcoin_events/main.go

wss: 
	go vet .
	go test .
	make examples

version:
	@echo "Latest v0/v1 tag: $(LATEST_TAG)"

check-version:
	@echo "$(RELEASE)" | grep -Eq '^[01]\.[0-9]+\.[0-9]+$$' || \
		{ echo "NEW_VERSION must be 0.x.y or 1.x.y, e.g. make release NEW_VERSION=1.1.0"; exit 1; }
	@! git rev-parse -q --verify "refs/tags/v$(RELEASE)" >/dev/null || \
		{ echo "tag v$(RELEASE) already exists"; exit 1; }

tag:
	git tag -a v$(RELEASE) -m "Release v$(RELEASE)"
	git push origin v$(RELEASE)

release: check-version wss tag
	gh release create v$(RELEASE) \
		--latest \
		--title "v$(RELEASE)" \
		--notes "Release v$(RELEASE)"
