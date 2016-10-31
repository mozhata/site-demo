package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

// Provide a homebrew module support.
// Reference: https://nodejs.org/api/modules.html
// (Not fully implemented yet, add features as needs)

var jsRequirePattern = regexp.MustCompile(`require\("[^"]+"\)`)
var jsRequireExtractPattern = regexp.MustCompile(`^require\("([^"]+)"\)$`)
var paths = []string{"./node_modules", "./tpl"}

func loadJS(module string, imported map[string]string) string {
	modulePath := ""
	for _, prefix := range paths {
		p := path.Join(prefix, module)
		stat, err := os.Stat(p)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			panic(err)
		}
		if stat.IsDir() {
			modulePath = path.Join(modulePath, "index.js")
			break
		}
		modulePath = p
		break
	}
	if modulePath == "" {
		panic(fmt.Sprintf("not found: %s", module))
	}
	js, err := ioutil.ReadFile(modulePath)
	if err != nil {
		panic(fmt.Sprintf("%v: %s", err, modulePath))
	}
	return fmt.Sprintf(`(function(){
var module = {};
var exports = module.exports;
%s
return exports
})()`, resolveRequires(string(js), imported))
}

func resolveRequires(js string, imported map[string]string) string {
	return jsRequirePattern.ReplaceAllStringFunc(js, func(s string) string {
		modulePath := jsRequireExtractPattern.FindStringSubmatch(s)[1]
		return loadJS(modulePath, imported)
	})
}
