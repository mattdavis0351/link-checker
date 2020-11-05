package files

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattdavis0351/link-checker/actions"
)

func ReadWorkspaceDir() []string {

	var filesToParse []string
	ign := strings.Split(actions.IgnoreExtenstions, ",")
	inc := strings.Split(actions.IncludeExtenstions, ",")

	err := filepath.Walk(actions.GithubWorkspace,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
			}
			if info.IsDir() && (info.Name() == ".git" || info.Name() == ".github") {
				return filepath.SkipDir
			}

			for _, i := range inc {
				if !info.IsDir() && strings.Contains(info.Name(), i) {
					filesToParse = append(filesToParse, path)
					return nil
				}
			}

			for _, e := range ign {
				if !info.IsDir() && !strings.Contains(info.Name(), e) {
					filesToParse = append(filesToParse, path)
				}
			}

			return nil
		})
	if err != nil {
		log.Panicln(err)
	}
	return filesToParse
}
