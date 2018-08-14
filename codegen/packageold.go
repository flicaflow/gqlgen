// +build !go1.11

package codegen

import (
	"go/build"
)

// pkgName retrieves the package name given
// the import path and and srcDir.
// This implementation uses go/build.
func pkgName(path, srcDir string) string {
	pkg, err := build.Default.Import(path, srcDir, 0)
	if err != nil {
		panic(err)
	}

	return pkg.Name
}

func resolvePkg(pkgName string) (string, error) {
	cwd, _ := os.Getwd()

	pkg, err := build.Default.Import(pkgName, cwd, build.FindOnly)
	if err != nil {
		return "", err
	}

	return pkg.ImportPath, nil
}
