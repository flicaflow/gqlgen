// +build go1.11

package codegen

import (
	"golang.org/x/tools/go/packages"
	"strconv"
)

// pkgName retrieves the package name given
// the import path and and srcDir.
// This implementation uses go/packages (requires 1.11).
func pkgName(path, destDir string) string {
	pkgs, err := packages.Load(nil, path)
	if err != nil {
		panic(err)
	}

	if len(pkgs) != 1 {
		panic("expected exactly on package for path " + path + ", got " + strconv.Itoa(len(pkgs)))
	}

	return pkgs[0].Name
}
