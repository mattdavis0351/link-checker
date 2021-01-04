package actions

import "os"

var (
	GithubWorkspace    string = os.Getenv("GITHUB_WORKSPACE")
	IgnoreExtenstions  string = GetInput("ignored_extensions")
	IncludeExtenstions string = GetInput("included_extensions")
)
