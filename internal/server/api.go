package server

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

type handlers struct {
	bp *core.Blueprint
}

// ElementBody is the envelope every element type shares.
type ElementBody struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	Frontmatter map[string]any `json:"frontmatter"`
	Body        string         `json:"body"`
	Hash        string         `json:"hash"`
}

// DocumentBody is a singleton document or README.
type DocumentBody struct {
	Name        string         `json:"name"`
	Frontmatter map[string]any `json:"frontmatter"`
	Body        string         `json:"body"`
	Hash        string         `json:"hash"`
}

type elementOutput struct{ Body ElementBody }
type elementListOutput struct{ Body []ElementBody }
type documentOutput struct{ Body DocumentBody }

type idPath struct {
	ID string `path:"id"`
}

type namePath struct {
	Name string `path:"name"`
}

func (h *handlers) register(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "listElements", Method: http.MethodGet, Path: "/api/elements",
		Summary: "List elements", Tags: []string{"elements"},
	}, func(ctx context.Context, in *struct {
		Type string `query:"type"`
	}) (*elementListOutput, error) {
		return nil, huma.Error501NotImplemented("not implemented")
	})

	huma.Register(api, huma.Operation{
		OperationID: "getElement", Method: http.MethodGet, Path: "/api/elements/{id}",
		Summary: "Get an element", Tags: []string{"elements"},
	}, func(ctx context.Context, in *idPath) (*elementOutput, error) {
		return nil, huma.Error501NotImplemented("not implemented")
	})

	huma.Register(api, huma.Operation{
		OperationID: "getDocument", Method: http.MethodGet, Path: "/api/documents/{name}",
		Summary: "Get a singleton document or README", Tags: []string{"documents"},
	}, func(ctx context.Context, in *namePath) (*documentOutput, error) {
		return nil, huma.Error501NotImplemented("not implemented")
	})

	// TODO: GET /api/blueprint, POST/PUT/DELETE /api/elements,
	// PUT /api/documents/{name}, /api/graph, /api/views/*, /api/status,
	// POST /api/verify.

	h.registerEvents(api)
}
