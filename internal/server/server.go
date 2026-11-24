// Package server serves the Huma HTTP API, Server-Sent Events, and the
// embedded web UI for `beaver-blueprint serve`.
package server

import (
	"context"
	"fmt"
	"io/fs"
	"net"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

// DefaultPort is 4243; Beaver Backlog uses 4242 so both can run side by side.
const DefaultPort = 4243

// Options configures Run.
type Options struct {
	Port   int
	Open   bool
	Static fs.FS // built web UI, the contents of web/dist
}

// NewAPI registers every operation on a fresh mux. It does not depend on
// the embedded UI, so the OpenAPI generator can build without web/dist.
func NewAPI(mux *http.ServeMux, bp *core.Blueprint) huma.API {
	config := huma.DefaultConfig("Beaver Blueprint", "v1")
	config.OpenAPIPath = "/api/openapi"
	config.DocsPath = ""
	config.SchemasPath = ""
	config.CreateHooks = nil // drop the $schema link transformer
	api := humago.New(mux, config)
	h := &handlers{bp: bp}
	h.register(api)
	return api
}

// Run serves the API and UI on 127.0.0.1 until ctx is cancelled.
func Run(ctx context.Context, bp *core.Blueprint, opts Options) error {
	mux := http.NewServeMux()
	NewAPI(mux, bp)
	mux.Handle("/", spa(opts.Static))

	w, err := watch(bp.Dir())
	if err != nil {
		return err
	}
	defer func() { _ = w.Close() }()

	// TODO: fall through to the next free port when taken; --open.
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", opts.Port))
	if err != nil {
		return err
	}
	fmt.Printf("http://%s\n", ln.Addr())

	srv := &http.Server{Handler: mux}
	go func() {
		<-ctx.Done()
		_ = srv.Close()
	}()
	if err := srv.Serve(ln); err != http.ErrServerClosed {
		return err
	}
	return nil
}
