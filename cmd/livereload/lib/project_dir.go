package lib

import (
	"os"
	"path"
	"strings"
)

type ProjectDir string

func GetProjectDir() ProjectDir {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for {
		if d == "" {
			panic("must run live-reload under site-demo")
		}
		if strings.HasSuffix(d, "site-demo") {
			return ProjectDir(d)
		}
		d = path.Dir(d)
	}
}

func (d ProjectDir) String() string {
	return string(d)
}

func (d ProjectDir) TmpDir() string {
	return path.Join(d.String(), "tmp", "entrypoint")
}
