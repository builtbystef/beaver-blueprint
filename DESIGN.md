# Beaver Blueprint — Design

A living architectural model of a software application that lives inside the
repository it describes. One directory of markdown files, a CLI, and a local
web UI, built for developers and their coding agents to maintain together.
No external service, no database, and no model calls of its own.

This document records the decisions made before writing code. It is the
reference for what v1 is and is not. It is the sibling of Beaver Backlog's
`DESIGN.md` and deliberately shares its bones.

## Why

Coding agents already do most implementation work and will do more. The
developer's job shifts toward architecture: system design, requirements,
boundaries, data, and the reasoning behind decisions. That work needs a
representation that is higher level than code, stays accurate as the code
changes, and is cheap for agents to read and update.

An earlier attempt derived this representation from source with static
analysis. It was fragile across languages and frameworks, and even when
correct the output was noisy and needed an agent pass to become useful.
Beaver Blueprint inverts that: the agent writes and maintains the model
explicitly, and the tool provides the schema, validation, drift detection,
and a viewer. Tokens are spent on judgment, not on parsing.

## Goals

- A single, structured, high-level model of one application: problem,
  use cases, data, modules, behaviour, interfaces, cross-cutting aspects,
  decisions, and operational context.
- Plain markdown files in the repo, version-controlled and reviewed in PRs
  alongside the code that motivated the change.
- Maintained by whatever coding agent the developer already uses, through
  the files themselves, a CLI, and a set of skills. No agent-specific
  integration.
- Deterministic staleness signals so the agent knows exactly what to
  revisit after code changes.
- A local web UI for understanding the system: graphs, traces, and
  editable detail pages.
- One static binary, no runtime dependencies, clean and small.

## Non-goals for v1

- Calling an LLM from the tool itself, for any purpose.
- JSON output mode, MCP server, or any programmatic contract beyond the
  internal HTTP API. Earlier experience with Beaver Backlog showed agents
  read markdown output more easily and JSON was noisier.
- Multiple blueprints per repository, monorepo aggregation, or references
  across repositories. Scale comes from module hierarchy. Splitting along
  deployment or ownership lines is a later feature, and nothing is
  reserved for it now; keywords will be chosen when needed.
- SQLite or any persistent index. If a query layer is ever wanted it is an
  ephemeral, gitignored derivative of the markdown.
- Custom element types or custom relationship types.
- Generated sequence or state diagrams. Flow steps are prose; agents write
  Mermaid where a diagram helps.
- Non-git projects. Drift detection requires a git work tree.
- Size or count lints ("too many modules"). Abstraction heuristics live in
  the skills.
- Project configuration file.
- npm or Homebrew packaging.
- Any coupling with Beaver Backlog at runtime.

## Concepts

- **Blueprint**: the `blueprint/` directory and everything in it. One
  blueprint describes one application.
- **Element**: a graph node with its own file. Seven types: `use-case`,
  `entity`, `module`, `flow`, `interface`, `aspect`, `decision`.
- **Singleton document**: one-per-blueprint prose files for the parts of
  the spec that are not repeatable: `overview.md`, `security.md`,
  `guidelines.md`, `operations.md`, `glossary.md`.
- **Relationship**: a typed, directed edge stored in exactly one element's
  frontmatter. Inverses are derived, never written.
- **Anchor**: a list of paths or globs on an element or singleton naming
  the code it describes. Anchors drive drift detection.

## Storage

### Location

- Directory `blueprint/` at the project root. Visible, not hidden: the
  blueprint is a first-class artifact.
- Every command except `init` locates it by walking up from the current
  directory, like git does. If none is found the command fails with
  `no blueprint found, run beaver-blueprint init`.
- The project root (the directory containing `blueprint/`) must be inside a
  git work tree. Anchors are relative to the project root.
- The directory is the only source of truth. No cache, index, or database.

### Layout

The layout mirrors the software design specification the tool grew out of,
section by section:

```
blueprint/
  overview.md          general information, problem, solution, assumptions, risks
  use-cases/           one file per use case
  entities/            README.md (data rules, database details) + one file per entity
  modules/             README.md (architectural approach, technologies) + one file per module
  flows/               one file per flow
  interfaces/          one file per interface, internal and external
  aspects/             one file per cross-cutting aspect
  decisions/           one file per decision; open questions are decisions with status open
  security.md          sensitive data, access, retention, regulations, threats
  guidelines.md        development conventions
  operations.md        environments, deployment, migrations, monitoring, recovery
  glossary.md          project terms
```

A `README.md` is permitted in any type directory and ignored as an element.
`init` scaffolds the two listed above with fixed headings; any other README
is free-form.

### Element files

- One file per element: `blueprint/<type-dir>/<id>.md`.
- The **id** is the filename without `.md`. The **type** is the directory.
  The id is not duplicated in frontmatter.
- Ids are slugs: lowercase ASCII letters, digits, and `-`, starting with a
  letter, max 60 characters. Ids are unique across the whole blueprint,
  not just within a type, so references are bare ids.
- Random ids (as in Beaver Backlog) were rejected here. Blueprint elements
  are referenced constantly; `depends_on: [auth, billing]` is readable by
  an agent and a reviewer without a lookup, and architectural names are
  stable. Renames go through `beaver-blueprint mv`, which rewrites
  references.
- `title` is independent of the id. Changing a title never renames a file.
- File content: YAML frontmatter followed by a markdown body with fixed
  `##` headings per type (see Vocabulary).

### Common frontmatter

Every element has:

| Field      | Type           | Required | Notes                                          |
|------------|----------------|----------|------------------------------------------------|
| `title`    | string         | yes      | display name                                   |
| `summary`  | string         | no       | one line, used in lists and graph tooltips     |
| `verified` | git SHA        | no       | explicit "reviewed, nothing to change" marker  |
| `code`     | list of globs  | no       | anchors; only on types that allow them         |

Type-specific fields follow in the Vocabulary section.

### Write rules

- The tool writes fields in a fixed order per type: common fields first in
  the order above, then type fields in the order listed in Vocabulary.
- Optional fields are omitted when empty.
- Unknown keys are preserved on every write. Keys prefixed `x-` are
  extension keys and never warned about. Any other unknown key is a
  `check` warning, so typos like `depends-on` are caught.
- Every write re-serializes frontmatter. Key order and formatting are
  normalized; YAML comments in frontmatter are not preserved.
- Body bytes are written back untouched unless the body itself was edited.

### Example

```markdown
---
title: Authentication
summary: Issues and validates sessions for every entry point.
code:
  - internal/auth/**
  - internal/server/middleware/auth.go
kind: internal
depends_on: [user-store, mailer]
---

## Responsibility

Verifies credentials, issues sessions, and exposes the current principal
to other modules. Owns password hashing and session expiry policy.

## Dependencies

`user-store` for credential lookup. `mailer` for password reset links.

## Not Responsible For

Authorization decisions; those are made by each module against the
principal, following the `authorization` aspect.
```

## Vocabulary

Seven element types. The field tables list type-specific frontmatter; the
heading lists are the body's fixed `##` sections, scaffolded by `new` and
checked by `check`. The field lists are lifted from sections 4 through 10
and 14 of the originating spec.

### use-case (§4)

| Field    | Type            | Required | Notes                   |
|----------|-----------------|----------|-------------------------|
| `actors` | list of strings | no       | people or other systems |

Headings: `Goal`, `Preconditions`, `Main Steps`, `Alternative Flows`,
`Expected Result`, `Business Rules`.

No anchors. A use case describes intent, not code.

### entity (§5)

| Field       | Type                          | Required | Notes                                           |
|-------------|-------------------------------|----------|-------------------------------------------------|
| `owned_by`  | module id                     | yes      | the module responsible for this data            |
| `relations` | list of `{to, cardinality, label}` | no  | `to` is an entity id; cardinality `1:1`, `1:n`, `n:m` |

Headings: `Meaning`, `Attributes`, `Rules and Constraints`.

Anchors allowed (schema files, migrations, model code).

### module (§6)

| Field        | Type            | Required | Notes                                   |
|--------------|-----------------|----------|-----------------------------------------|
| `kind`       | enum            | no       | `internal` (default), `external`        |
| `parent`     | module id       | no       | hierarchy; guidance is three levels max |
| `depends_on` | list of module ids | no    |                                         |

Headings: `Responsibility`, `Dependencies`, `Not Responsible For`.

Anchors allowed and expected for internal modules; `check` warns on an
internal module without anchors. External modules (Stripe, Postgres, a
partner API) have no anchors, never go stale, and render distinctly in the
graph. "Data owned" from the spec is derived from entities' `owned_by`.

### flow (§7)

| Field      | Type               | Required | Notes                      |
|------------|--------------------|----------|----------------------------|
| `realizes` | use-case id        | no       |                            |
| `modules`  | list of module ids | no       | participating modules      |

Headings: `Trigger`, `Steps`, `State Changes`, `Events and Background Work`,
`Failures and Recovery`.

No anchors. Sequence or state diagrams go in the body as Mermaid.

### interface (§8 and §9, unified)

| Field       | Type               | Required | Notes                                                        |
|-------------|--------------------|----------|--------------------------------------------------------------|
| `kind`      | enum               | yes      | `internal`, `gui`, `cli`, `api`, `mcp`, `webhook`, `file`     |
| `provider`  | module id          | yes      |                                                              |
| `consumers` | list of module ids | no       | internal interfaces                                          |
| `audience`  | list of strings    | no       | external kinds: who uses it                                  |

Headings: `Purpose`, `Structure`, `Inputs`, `Outputs`, `Errors`,
`Rules and Guarantees`.

`Structure` holds operations and events for an internal interface, and
screens, commands, endpoints, or tools for an external one. Conventions
(naming, versioning, accessibility) belong under `Rules and Guarantees`.
Anchors allowed.

### aspect (§10)

| Field        | Type               | Required | Notes                      |
|--------------|--------------------|----------|----------------------------|
| `applies_to` | list of module ids | no       | empty means all modules    |

Headings: `Approach`, `Rules`.

A cross-cutting aspect: authentication, authorization, logging and audit,
error handling, configuration and secrets, transactions and consistency.
Anchors allowed. The name was chosen over `concern`, `policy`,
`convention`, and `mechanism` as the one term that means exactly "a thing
that cuts across modules". If it proves to be overhead in practice it
collapses into a singleton `aspects.md` without touching anything else.

### decision (§14)

| Field        | Type            | Required | Notes                                                      |
|--------------|-----------------|----------|------------------------------------------------------------|
| `status`     | enum            | yes      | `open`, `proposed`, `accepted`, `superseded`, `rejected`   |
| `date`       | ISO date        | no       |                                                            |
| `affects`    | list of ids     | no       | any element type                                           |
| `supersedes` | decision id     | no       |                                                            |

Headings: `Context`, `Decision`, `Trade-offs`.

An open question is a decision with `status: open` and an empty `Decision`
section. No separate type.

### Singleton documents

Singletons follow the same fixed-heading rule. All accept an optional
`code` list and the same `verified` marker, so `operations.md` can be
anchored to deployment configuration and `guidelines.md` to linter and
formatter configuration.

- `overview.md` (§1 to §3, plus assumptions and risks from §14).
  Frontmatter: `schema: 1` (required; the only version marker in the
  blueprint), `name` (required), `description`, and optional project
  fields `client`, `team`, `milestones`, `owner`, `links`. "Last updated"
  is never stored; git knows. Headings: `Problem`, `People Affected`,
  `Current Approach`, `Desired Outcome`, `Solution`, `Capabilities`,
  `Out of Scope`, `Constraints`, `Transition`, `Assumptions`, `Risks`.
- `security.md` (§11). Headings: `Sensitive Data`, `Access`,
  `Retention and Deletion`, `Regulations`, `Threats`.
- `guidelines.md` (§12). Headings: `Code Organization`, `Naming`,
  `Formatting`, `Function Size and Responsibility`, `Testing`,
  `Required Coverage`, `Code Review`, `Links`.
- `operations.md` (§13). Headings: `Environments and Hosting`,
  `Build and Deployment`, `Database Migrations`, `Monitoring and Alerts`,
  `Backups and Recovery`, `Rollback`, `Responsibilities`.
- `glossary.md` (§15). No fixed headings; a list of `- **Term**: meaning`.
- `modules/README.md`. Headings: `Approach`, `Technologies and Restrictions`.
- `entities/README.md`. Headings: `Data Rules`, `Database Details`.

## Relationships

- Each relationship is stored once, on the side that naturally owns it:
  `module.parent`, `module.depends_on`, `entity.owned_by`,
  `entity.relations`, `flow.realizes`, `flow.modules`,
  `interface.provider`, `interface.consumers`, `aspect.applies_to`,
  `decision.affects`, `decision.supersedes`.
- The core resolves inverses (a module's children, owned entities,
  dependents, provided interfaces, flows it participates in, aspects that
  apply, decisions that affect it) for `show`, the graph, and the UI.
- Every reference field has an allowed target type. `check` errors on a
  dangling id and on an id of the wrong type.
- Module `parent` chains must be acyclic. `depends_on` cycles are
  permitted and reported as a warning, since real systems have them.

## Drift

Drift detection is deterministic and git-based. It tells the agent what to
look at; the agent decides what to change.

### Rule

An anchored element or singleton is **stale** when any file matching its
anchors changed more recently than the element's own file.

Precisely, for each anchored file `F` with anchors `A`:

1. If `F` has uncommitted modifications, it is being updated and is
   current. Stop.
2. Let `R` be the last commit that touched `F`. If `verified` is set and
   names a commit that is a descendant of `R`, use that instead.
3. Committed drift: `git diff --name-only R..HEAD -- A` is non-empty.
4. Working-tree drift: any tracked file matching `A` is modified, or any
   untracked file matches `A`.

Because an element edited in the same commit as its code has `R = HEAD`,
the normal workflow (change code, change blueprint, commit together) never
produces false drift. The `verified` field exists only for "reviewed,
nothing to change": `verify` writes it, which touches the file and makes
the element current, and the SHA makes the intent visible in a PR.

A `verified` SHA that is not in the current history (after a rebase) is a
warning and the field is ignored.

### Commands and severity

- `status` lists stale elements in two groups, committed drift and
  working-tree drift, with the changed files under each element.
  `--committed` hides the second group.
- `check` reports drift as a warning. `check --strict` promotes it to an
  error. CI runs `check --strict`, so a merged PR can never leave anchored
  code changed without its element touched. Locally, a feature in progress
  across several commits does not fail `check`.
- `verify <id>...` sets `verified` to HEAD on each named file.

### Anchors

- Globs use `**` doublestar syntax, relative to the project root.
- An anchor matching no files is a `check` error. It usually means the
  code moved and the blueprint did not follow.

## Validation

`beaver-blueprint check` validates the whole blueprint.

Errors (exit 1):

- Unparseable frontmatter or body.
- File in a type directory whose name is not a valid id.
- Unknown directory inside `blueprint/`.
- Missing required field, invalid enum value, wrong field type.
- Duplicate id across directories.
- Reference to a nonexistent id, or to an id of the wrong type.
- `parent` cycle.
- Anchor matching no files.
- Missing or unsupported `schema` in `overview.md`.
- Under `--strict`: any warning.

Warnings (exit 0):

- Unknown frontmatter key without `x-` prefix.
- Missing fixed heading, or a `##` heading not in the type's list.
- Internal module without anchors.
- Element with no relationships in either direction.
- `depends_on` cycle.
- `verified` SHA not in history.
- Drift.

`check --fix` performs mechanical repairs only: canonical frontmatter
formatting and field order, and reordering of fixed headings. It never
creates content and never deletes anything.

## CLI

- Binary name: `beaver-blueprint`. Users who want something shorter add a
  shell alias. No second binary is shipped.
- Human-readable output only, terse and aligned, in the style of Beaver
  Backlog. No JSON mode in v1.
- Honors `NO_COLOR`. No spinners or interactive prompts.
- Exit code 0 on success, 1 on any failure.
- Root `--help` opens with a short workflow section for agents (read the
  blueprint before designing, update it in the same change as the code,
  run `check` and `status` before finishing, use the skills) followed by
  the command reference. Help is the single entry point.

### Commands

```
beaver-blueprint init [--agent <name>]       scaffold blueprint/, write the AGENTS.md pointer, install skills
beaver-blueprint new <type> <id> [--title]   create an element with correct frontmatter and headings
beaver-blueprint mv <old-id> <new-id>        rename an element and rewrite every reference to it
beaver-blueprint rm <id> [--force]           delete an element; fails if referenced unless --force, which strips references
beaver-blueprint list [--type <type>]        one line per element: type, id, title, stale marker
beaver-blueprint show <id>                   print the file, then resolved relationships in both directions
beaver-blueprint check [--strict] [--fix]    validate; see Validation
beaver-blueprint status [--committed]        list stale elements and the files that changed
beaver-blueprint verify <id>...              stamp elements as reviewed at HEAD
beaver-blueprint schema [<type>]             print the frontmatter JSON Schema
beaver-blueprint serve [--port] [--open]     start the local web UI
```

`init` on a project that already has `blueprint/` fails. `new` fails on a
duplicate id or an unknown type. `init` creates the complete empty
skeleton: every type directory, the two READMEs, the five singletons with
headings, `overview.md` with `schema: 1` and `name` taken from the
directory name. No example elements. The skills fill the skeleton.

## Agent integration

Files are the API. Agents read and write the markdown directly with their
own file tools. The CLI exists for what agents do badly by hand:
scaffolding the schema, validating, computing drift, renaming with
reference rewriting, and serving the UI.

### Pointer

`init` writes to `AGENTS.md`, creating it if absent and appending if
present:

```
This project keeps its architecture in a Beaver Blueprint (`blueprint/`).
Run `beaver-blueprint --help` before reading or changing it. Skills:
blueprint-greenfield, blueprint-brownfield, blueprint-maintain.
```

### Skills

Three skills in the Agent Skills `SKILL.md` format, embedded in the binary
so they version with the CLI, and installed by `init` into every agent
skills directory found in the project (`.claude/skills/`,
`.agents/skills/`, `.cursor/skills/`, and others verified at
implementation), falling back to `.agents/skills/` when none exist.
`--agent <name>` forces a single target.

- `blueprint-greenfield`: runs the design conversation for a new
  application and fills the skeleton section by section, in spec order,
  before any code is written.
- `blueprint-brownfield`: explores an existing codebase (structure, entry
  points, data, running it where possible) and produces a first
  blueprint, then guides the review-and-refine pass with the developer.
- `blueprint-maintain`: how and when to update the blueprint during
  normal work, driven by `status` output, and what to do before finishing
  a task.

### Abstraction heuristics

These live in the skills, not in `check`:

- A module is something you would assign ownership of or deploy
  separately. Never a file or a class. A typical application has five to
  fifteen top-level modules.
- An entity is a concept the business would name, not a database table.
- A use case is a user goal, not a screen.
- A flow is only written when the coordination between modules is not
  obvious from the module descriptions.
- An interface is written for every external entry point, and for
  internal boundaries only where the contract matters.
- A decision is recorded when it would be expensive to reverse or when a
  future reader would ask "why".
- The blueprint describes what and why. How lives in the code.

## Web UI

Served by `beaver-blueprint serve`.

- Binds `127.0.0.1` on port 4243 by default (Beaver Backlog uses 4242 so
  both can run side by side), falls to the next free port if taken.
  `--port` overrides. Prints the URL; `--open` launches the browser.
  No authentication.
- Views:
  - Overview: `overview.md` rendered, plus a drift summary.
  - Modules: generated dependency graph. Hierarchy as nesting, external
    modules styled distinctly, stale modules highlighted. React Flow with
    ELK layout. Adding an edge on the graph appends to `depends_on`.
  - Entities: generated ER-style diagram from `relations`, ownership shown
    by colour or grouping.
  - Use cases: list, and per use case a trace view (use case to flows to
    modules to interfaces).
  - Flows, interfaces, aspects, decisions: list and detail. Decisions
    filterable by status.
  - Status: the drift list with a verify action per element.
  - Singletons: rendered and editable.
- Element detail, one shape for all types: frontmatter as a type-aware
  form (selects for enums, pickers with autocomplete for references,
  tag input for lists of strings), body in a markdown WYSIWYG editor with
  a raw-markdown toggle, Mermaid blocks rendered.
- Full create, edit, delete, and relationship editing. Creating an element
  asks for type, id, and title. There is no chat or model feature of any
  kind in the UI, and no raw-file editor.
- Body round-tripping: if the body was not edited the original bytes are
  written back untouched; if it was, normalized formatting is accepted.
- Live updates: the server watches `blueprint/` with fsnotify and sends a
  "changed" ping over Server-Sent Events. The UI re-fetches.
- Conflicts: every element response carries a content hash. A save sends
  it back; the server rejects the save if the file changed since and the
  UI shows "changed on disk, reload" with the newer version. No merging.
- Filters and selections live in the URL query string. Nothing is stored
  server-side.
- Theme follows the system light/dark preference.

## HTTP API

- Used by the SPA only. Internal and unversioned in v1.
- Built with Huma on the standard library `net/http`. Huma generates the
  OpenAPI 3.1 document from the Go handler types and validates requests
  against it. The Go structs are the single source of truth for the
  schema; the element frontmatter schemas appear in the document's
  components, and `beaver-blueprint schema` prints them from the embedded
  document.
- Generic element and document endpoints, plus dedicated derived views,
  because every type shares one envelope (id, type, frontmatter, body,
  hash) and type-specific behaviour is read-side and derived:

```
GET    /api/blueprint                  everything: elements, documents, resolved graph
GET    /api/elements                   list; ?type= filter
POST   /api/elements                   create
GET    /api/elements/{id}
PUT    /api/elements/{id}              full replace; requires hash
DELETE /api/elements/{id}
GET    /api/documents/{name}           singletons and READMEs
PUT    /api/documents/{name}
GET    /api/graph                      nodes and edges with inverses
GET    /api/views/modules              module graph with hierarchy, ready for layout
GET    /api/views/entities             entity diagram
GET    /api/views/use-cases/{id}/trace
GET    /api/status                     drift
POST   /api/verify                     body: list of ids
GET    /api/events                     SSE
GET    /api/openapi.json
```

- Codegen: `just openapi` writes `api/openapi.json` from the Go code;
  `openapi-typescript` turns it into `web/src/api/types.ts`;
  `openapi-fetch` provides the typed client. Both generated files are
  committed and CI fails if regenerating them produces a diff, so the
  frontend builds without a Go toolchain.
- All logic lives in the core package; the API and the CLI are thin
  interfaces over it.

## Stack

Mirrors Beaver Backlog wherever the blueprint has no reason to differ.

- Go 1.27, Cobra for the CLI, Huma mounted on standard library `net/http`
  through its `humago` adapter. Viper is not used: there is no
  configuration to load.
- YAML through a library whose node API round-trips unknown keys and
  preserves order (`goccy/go-yaml` or `yaml.v3`, chosen at
  implementation).
- `fsnotify` for watching, `doublestar` for anchor globs, `go-git` or
  shelling out to `git` for history (chosen at implementation; shelling
  out is acceptable since git is required anyway).
- Frontend: Vite, React 19, TypeScript, Tailwind 4, Base UI for
  components, MDXEditor for bodies, React Flow with ELK for generated
  graphs, Mermaid for authored diagrams. Bundle size is irrelevant for a
  localhost tool.
- Frontend built to `web/dist` and embedded with `go:embed`. End users
  never need Node.

### Repository layout

```
cmd/beaver-blueprint/   main package
internal/blueprint/     core: discovery, parsing, writing, schema, validation, graph, drift
internal/cli/           Cobra commands (thin layer over internal/blueprint)
internal/server/        Huma API, SSE, embedded static files
skills/                 the three SKILL.md directories, embedded
api/openapi.json        generated, committed
web/                    Vite + React app; web/dist is embedded
DESIGN.md               this file
justfile                build tasks
.goreleaser.yaml        release config
LICENSE                 MIT
```

### Development workflow

- `just web` builds the frontend; `just openapi` regenerates the OpenAPI
  document and TypeScript types; `just build` runs both then builds the
  Go binary. The Go build fails loudly if `web/dist` is missing.
- During UI development the Vite dev server proxies `/api` to a running
  `beaver-blueprint serve`.
- Releases via goreleaser: linux, macOS, Windows on amd64 and arm64.
  Linux is the day-to-day tested platform. `go install` also works.

## Testing

- `internal/blueprint`: golden-file tests for parsing and writing, table
  tests for id validation, reference resolution, and heading checks,
  validation tests mirroring `check`, and drift tests against temporary
  git repositories with scripted histories.
- `internal/cli`: command tests against a temporary directory.
- `internal/server`: handler tests through Huma's test helpers; a
  staleness check that `api/openapi.json` matches the code.
- Frontend: type checking and lint only; no component tests in v1.

## Build order

1. Core: parse, write, schema, `check`, graph resolution, drift.
2. CLI: `init`, `new`, `mv`, `rm`, `check`, `status`, `verify`, `list`,
   `show`, `schema`.
3. Skills and the `init` pointer. Then dogfood: write Beaver Blueprint's
   own blueprint with `blueprint-greenfield`, and run
   `blueprint-brownfield` on the Beaver Backlog repository. Expect the
   vocabulary and heading lists to change here.
4. Server, Huma API, OpenAPI codegen, SSE.
5. Web UI: list and detail pages first, graphs second, editing third.

The UI comes last because the CLI and skills already deliver the core
value, and dogfooding at step 3 will reshape the vocabulary before any UI
is built on it.

## Sibling project

Beaver Backlog is the sibling tool, shipped as its own binary from its own
repository. The two share conventions (one directory of markdown at the
project root, walk-up discovery, `AGENTS.md` pointer, help as the agent
entry point, same stack, same layout, same conflict model) but no code
and no runtime coupling. If duplication in markdown and frontmatter
handling, discovery, or embedded-UI plumbing becomes painful, it is
extracted into a shared module later.

## Related approaches

- Deriving architecture from source by static analysis: tried and
  rejected, see Why.
- arc42 and the C4 model: informed the vocabulary (module sits roughly at
  C4's container and component levels; the singleton documents cover
  arc42's context, constraints, and deployment views). Rejected as
  formats because both are heavier than needed and neither is designed
  for agent maintenance.
- Architecture decision records: adopted directly as the `decision` type.
- Structurizr DSL and similar modelling languages: rejected as too formal
  and too far from prose for agents to maintain alongside code.
- Backlog.md and Beaver Backlog: borrowed the one-file-per-item layout,
  the `init` pointer, and help as the entry point.
