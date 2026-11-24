package core

// Type is an element type. The type of an element is its directory.
type Type string

const (
	UseCase   Type = "use-case"
	Entity    Type = "entity"
	Module    Type = "module"
	Flow      Type = "flow"
	Interface Type = "interface"
	Aspect    Type = "aspect"
	Decision  Type = "decision"
)

// TypeSpec describes one element type: where it lives, its fixed body
// headings, and whether it may carry code anchors.
type TypeSpec struct {
	Type     Type
	Dir      string
	Headings []string
	Anchors  bool
}

// Types lists every element type in spec order.
var Types = []TypeSpec{
	{UseCase, "use-cases", []string{"Goal", "Preconditions", "Main Steps", "Alternative Flows", "Expected Result", "Business Rules"}, false},
	{Entity, "entities", []string{"Meaning", "Attributes", "Rules and Constraints"}, true},
	{Module, "modules", []string{"Responsibility", "Dependencies", "Not Responsible For"}, true},
	{Flow, "flows", []string{"Trigger", "Steps", "State Changes", "Events and Background Work", "Failures and Recovery"}, false},
	{Interface, "interfaces", []string{"Purpose", "Structure", "Inputs", "Outputs", "Errors", "Rules and Guarantees"}, true},
	{Aspect, "aspects", []string{"Approach", "Rules"}, true},
	{Decision, "decisions", []string{"Context", "Decision", "Trade-offs"}, true},
}

// LookupType returns the spec for a type name.
func LookupType(name string) (TypeSpec, bool) {
	for _, t := range Types {
		if string(t.Type) == name {
			return t, true
		}
	}
	return TypeSpec{}, false
}

// DocumentSpec describes a singleton document or a type-directory README.
type DocumentSpec struct {
	Name     string // path relative to blueprint/, e.g. "overview.md"
	Headings []string
}

// Documents lists the singleton documents and the scaffolded READMEs.
var Documents = []DocumentSpec{
	{"overview.md", []string{"Problem", "People Affected", "Current Approach", "Desired Outcome", "Solution", "Capabilities", "Out of Scope", "Constraints", "Transition", "Assumptions", "Risks"}},
	{"security.md", []string{"Sensitive Data", "Access", "Retention and Deletion", "Regulations", "Threats"}},
	{"guidelines.md", []string{"Code Organization", "Naming", "Formatting", "Function Size and Responsibility", "Testing", "Required Coverage", "Code Review", "Links"}},
	{"operations.md", []string{"Environments and Hosting", "Build and Deployment", "Database Migrations", "Monitoring and Alerts", "Backups and Recovery", "Rollback", "Responsibilities"}},
	{"glossary.md", nil},
	{"modules/README.md", []string{"Approach", "Technologies and Restrictions"}},
	{"entities/README.md", []string{"Data Rules", "Database Details"}},
}
