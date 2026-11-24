// Command openapigen writes the OpenAPI document generated from the Huma
// handler types. Run through `just openapi`.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/builtbystef/beaver-blueprint/internal/server"
)

func main() {
	api := server.NewAPI(http.NewServeMux(), nil)
	out, err := json.MarshalIndent(api.OpenAPI(), "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.WriteFile("web/openapi.json", append(out, '\n'), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
