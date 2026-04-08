package doctor

// Rule is a single doctor check (ALC-224).
type Rule interface {
	ID() string
	Run(ctx *RuleContext) ([]Finding, error)
}
