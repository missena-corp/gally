package project

type Strategy struct {
	Branch string
	Only   string
}

type Strategies map[string]Strategy

const (
	COMPARE_TO      = "compare-to"
	PREVIOUS_COMMIT = "previous-commit"
)

var (
	defaultStrategies = Strategies{
		COMPARE_TO: {
			Branch: "master",
		},
		PREVIOUS_COMMIT: {
			Only: "master",
		},
	}
)
