package server

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/sse"
	"github.com/fsnotify/fsnotify"
)

// ChangedEvent tells the UI that something under blueprint/ changed and it
// should re-fetch.
type ChangedEvent struct{}

func (h *handlers) registerEvents(api huma.API) {
	sse.Register(api, huma.Operation{
		OperationID: "events", Method: http.MethodGet, Path: "/api/events",
		Summary: "Change notifications", Tags: []string{"events"},
	}, map[string]any{"changed": ChangedEvent{}}, func(ctx context.Context, in *struct{}, send sse.Sender) {
		// TODO: fan out from a single watcher started in Run.
		<-ctx.Done()
	})
}

// watch reports changes under dir. Placeholder for the fsnotify watcher.
func watch(dir string) (*fsnotify.Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	if err := w.Add(dir); err != nil {
		_ = w.Close()
		return nil, err
	}
	return w, nil
}
