set shell := ["bash", "-euo", "pipefail", "-c"]

# Regenerate everything and build the binary.
default: build

# Write web/openapi.json from the Go handler types, then the TypeScript
# client in web/src/api/generated from it. Both outputs are committed.
openapi:
    go run ./internal/server/openapigen
    cd web && vp install --frozen-lockfile && vp run openapi

# Build the frontend into web/dist, which the Go binary embeds.
web:
    cd web && vp install --frozen-lockfile && vp build

# Codegen, frontend, then the Go binary.
build: openapi web
    go build -trimpath -o bin/beaver-blueprint ./cmd/beaver-blueprint

# The Go packages embed web/dist, so tests need the frontend built.
test: web
    go test ./...

# Static checks: Go vet and lint, then vp check (format, lint, types).
check: web
    go vet ./...
    golangci-lint run
    cd web && vp check

# Fail if regenerating the OpenAPI document or the client changes them.
openapi-check: openapi
    git diff --exit-code -- web/openapi.json web/src/api/generated

# Everything CI runs.
ci: openapi-check check test

# Vite dev server; proxies /api to a running `beaver-blueprint serve`.
dev:
    cd web && vp dev
