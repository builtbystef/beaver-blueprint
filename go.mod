module github.com/builtbystef/beaver-blueprint

go 1.27.0

// pnpm installs into web/node_modules; it is not Go code.
ignore ./web/node_modules

require (
	github.com/bmatcuk/doublestar/v4 v4.10.4
	github.com/danielgtaylor/huma/v2 v2.41.0
	github.com/fsnotify/fsnotify v1.10.1
	github.com/spf13/cobra v1.10.2
	go.yaml.in/yaml/v4 v4.0.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.1.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
)
