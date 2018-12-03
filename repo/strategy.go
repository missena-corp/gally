package repo

type StrategyOpts struct {
	Branch string
	Only   string
}

func CompareTo(opts StrategyOpts) ([]string, error) {
	return UpdatedFiles(opts.Branch)
}

func PreviousCommitDiff(opts StrategyOpts) ([]string, error) {
	if !IsOnBranch(opts.Only) {
		return make([]string, 0), nil
	}
	return UpdatedFiles("HEAD^1")
}
