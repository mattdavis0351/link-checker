package actions

import "os"

var (
	GithubWorkspace    string = os.Getenv("GITHUB_WORKSPACE")
	IgnoreExtenstions  string = actions.GetInput("ignored_extensions")
	IncludeExtenstions string = actions.GetInput("included_extensions")
)
