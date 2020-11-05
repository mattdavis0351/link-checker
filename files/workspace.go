package files

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattdavis0351/link-checker/actions"
)

var (
	githubWorkspace    string = os.Getenv("GITHUB_WORKSPACE")
	ignoreExtenstions  string = actions.GetInput("ignored_extensions")
	includeExtenstions string = actions.GetInput("included_extensions")
)

func ReadWorkspaceDir() []string {
	// files, err := ioutil.ReadDir(os.Getenv("GITHUB_WORKSPACE"))
	// var fileNames []string
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // TODO check if current item is file or dir
	// // if dir then walk the tree getting out all the filenames to append
	// // if file, then append the filename
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		f, err := ioutil.ReadDir(file.Name())
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		for _, x := range f {
	// 			if x.IsDir() {
	// 				a, err := ioutil.ReadDir(x.Name())
	// 				if err != nil {
	// 					log.Fatal()
	// 				}
	// 				for _, y := range a {
	// 					fmt.Println("use recursion")
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fileNames = append(fileNames, file.Name())

	// 	fmt.Println(file.Name())

	// }
	// return fileNames
	var filesToParse []string
	ign := strings.Split(ignoreExtenstions, ",")
	inc := strings.Split(includeExtenstions, ",")

	err := filepath.Walk(githubWorkspace,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
			}
			if info.IsDir() && (info.Name() == ".git" || info.Name() == ".github") {
				return filepath.SkipDir
			}
			// filesToParse = append(filesToParse, path)
			// for _, i := range inc {
			// 	for _, e := range ign {
			// 		if !info.IsDir() && (strings.Contains(info.Name(), i) || !strings.Contains(info.Name(), e)) {
			// 			filesToParse = append(filesToParse, path)
			// 		}
			// 	}
			// }
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

func PrepareToParse(p string, inc, exc []string) []string {
	var filesToParse []string

	for _, i := range inc {
		for _, e := range exc {
			if strings.Contains(p, i) || !strings.Contains(p, e) {
				filesToParse = append(filesToParse, p)
			}
		}

	}
	// if the file.Name() includes extension or does not exclude extension

	return filesToParse

}
